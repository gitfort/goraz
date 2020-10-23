package local

import (
	"context"
	"github.com/gitfort/goraz/pkg/contextext"
	"golang.org/x/text/language"
)

func GetLang(ctx context.Context) language.Tag {
	value, ok := contextext.GetValue(ctx, "lang")
	if !ok {
		return language.Und
	}
	tag, err := language.Parse(value)
	if err != nil {
		return language.Und
	}
	return tag
}
