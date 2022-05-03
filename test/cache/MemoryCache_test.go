package test_cache

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pip-services3-gox/pip-services3-components-gox/cache"
)

func TestMemoryCache(t *testing.T) {
	var _cache cache.ICache[any]
	_cache = cache.NewMemoryCache[any]()

	value, err := _cache.Retrieve(context.Background(), "", "key1")
	assert.Nil(t, value)
	assert.Nil(t, err)

	value, err = _cache.Store(context.Background(), "", "key1", "value1", 250)
	assert.Equal(t, "value1", value)
	assert.Nil(t, err)

	value, err = _cache.Retrieve(context.Background(), "", "key1")
	assert.Equal(t, "value1", value)
	assert.Nil(t, err)

	time.Sleep(500 * time.Millisecond)

	value, err = _cache.Retrieve(context.Background(), "", "key1")
	assert.Nil(t, value)
	assert.Nil(t, err)
}
