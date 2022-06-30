package genval

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyContext(t *testing.T) {
	type a struct{}

	ctx := context.Background()

	_, err := FromContext[struct{}](ctx, a{})
	require.ErrorIs(t, err, errCtx)
}

func TestContext(t *testing.T) {
	type a struct{}

	type val struct{}

	ctx := context.Background()

	value := &val{}

	ctx = WithContext(ctx, a{}, value)

	gotValue, err := FromContext[*val](ctx, a{})
	require.NoError(t, err)

	require.Equal(t, value, gotValue)
}
