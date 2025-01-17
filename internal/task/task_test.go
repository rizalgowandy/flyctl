package task

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromContextPanics(t *testing.T) {
	assert.Panics(t, func() { _ = FromContext(context.Background()) })
}

func TestClient(t *testing.T) {
	exp := new(manager)

	ctx := WithContext(context.Background(), exp)
	assert.Same(t, exp, FromContext(ctx))
}
