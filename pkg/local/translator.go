package local

import (
	"context"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"strings"
)

func NewTranslator(defaultLang language.Tag) *Translator {
	bundle := i18n.NewBundle(defaultLang)
	return &Translator{
		bundle:      bundle,
		localizers:  make(map[language.Tag]*i18n.Localizer),
		defaultLang: defaultLang,
	}
}

type Translator struct {
	bundle      *i18n.Bundle
	localizers  map[language.Tag]*i18n.Localizer
	defaultLang language.Tag
}

func (l *Translator) DefaultLang() language.Tag {
	return l.defaultLang
}

func (l *Translator) Languages() []language.Tag {
	return l.bundle.LanguageTags()
}

func (l *Translator) Load(path string, format string, unmarshalFunc i18n.UnmarshalFunc) error {
	l.bundle.RegisterUnmarshalFunc(format, unmarshalFunc)
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, format) {
			return nil
		}
		if _, err := l.bundle.LoadMessageFile(path); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	for _, tag := range l.bundle.LanguageTags() {
		l.localizers[tag] = i18n.NewLocalizer(l.bundle, tag.String())
	}
	return nil
}

func (l *Translator) ByLangWithData(lang language.Tag, id string, data interface{}) string {
	localizer, ok := l.localizers[lang]
	if !ok {
		localizer = l.localizers[l.defaultLang]
	}
	str, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: data,
	})
	if err != nil {
		return id
	}
	return str
}

func (l *Translator) ByLang(lang language.Tag, id string) string {
	return l.ByLangWithData(lang, id, nil)
}

func (l *Translator) ByContextWithData(ctx context.Context, id string, data interface{}) string {
	return l.ByLangWithData(GetLang(ctx), id, data)
}

func (l *Translator) ByContext(ctx context.Context, id string) string {
	return l.ByContextWithData(ctx, id, nil)
}
