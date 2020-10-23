package httpext

import (
	"github.com/gitfort/goraz/pkg/contextext"
	"github.com/google/uuid"
	"golang.org/x/text/language"
	"net/http"
)

func LoadContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		if lang := GetLang(req); lang != language.Und {
			ctx = contextext.SetLang(ctx, lang)
		}
		if token := GetToken(req); token != "" {
			ctx = contextext.SetToken(ctx, token)
		}
		if sessionID := GetSessionID(req); sessionID != uuid.Nil {
			ctx = contextext.SetSessionID(ctx, sessionID)
		}
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
