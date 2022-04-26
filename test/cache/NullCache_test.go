package test_cache

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pip-services3-gox/pip-services3-components-gox/cache"
)

func TestNullCache(t *testing.T) {
	_cache := cache.NewNullCache[any]()

	value, err := _cache.Retrieve(context.Background(), "", "key1")
	assert.Nil(t, value)
	assert.Nil(t, err)

	value, err = _cache.Store(context.Background(), "", "key1", "value1", 0)
	assert.Equal(t, "value1", value)
	assert.Nil(t, err)
}
