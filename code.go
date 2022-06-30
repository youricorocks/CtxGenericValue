package genval

import (
	"context"

	"github.com/bdandy/go-errors"
)

var (
	errCtx = serrors.String("value missing in the context")
)

func WithContext(ctx context.Context, key any, value any) context.Context {
	return context.WithValue(ctx, key, value)
}

func FromContext[T any, PtrT *T](ctx context.Context, key any) (PtrT, error) {
	value, ok := ctx.Value(key).(PtrT)
	if !ok {
		return nil, errCtx
	}

	return value, nil
}
