package contextext

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/text/language"
	"google.golang.org/grpc/metadata"
)

func SetValue(ctx context.Context, key, value string) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, key, value)
	ctx = context.WithValue(ctx, key, value)
	return ctx
}

func GetValue(ctx context.Context, key string) (string, bool) {
	value, ok := ctx.Value(key).(string)
	if ok {
		return value, true
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	if len(md[key]) > 0 {
		return md[key][0], true
	}
	return "", false
}

func SetToken(ctx context.Context, token string) context.Context {
	if token == "" {
		return ctx
	}
	return SetValue(ctx, "token", token)
}
func GetToken(ctx context.Context) (string, bool) {
	return GetValue(ctx, "token")
}

func SetLang(ctx context.Context, lang language.Tag) context.Context {
	if lang == language.Und {
		return ctx
	}
	return SetValue(ctx, "lang", lang.String())
}
func GetLang(ctx context.Context) (language.Tag, bool) {
	value, ok := GetValue(ctx, "lang")
	if !ok {
		return language.Und, false
	}
	tag, err := language.Parse(value)
	if err != nil {
		return language.Und, false
	}
	return tag, true
}

func SetSessionID(ctx context.Context, sessionID uuid.UUID) context.Context {
	if sessionID == uuid.Nil {
		return ctx
	}
	return SetValue(ctx, "session_id", sessionID.String())
}
func GetSessionID(ctx context.Context) (uuid.UUID, bool) {
	value, ok := GetValue(ctx, "session_id")
	if !ok {
		return uuid.Nil, false
	}
	id, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, false
	}
	return id, true
}
