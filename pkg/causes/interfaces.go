package causes

import "context"

type Translator interface {
	ByContext(ctx context.Context, id string) string
}
