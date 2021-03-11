package test_cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pip-services3-gox/pip-services3-components-gox/cache"
)

func TestMemoryCache(t *testing.T) {
	cache := cache.NewMemoryCache()
	var str string

	ok, err := cache.RetrieveAs("", "key1", &str)
	assert.False(t, ok)
	assert.Equal(t, "", str)
	assert.Nil(t, err)

	value, err := cache.Retrieve("", "key1")
	assert.Nil(t, value)
	assert.Nil(t, err)

	value, err = cache.Store("", "key1", "value1", 250)
	assert.Equal(t, "value1", value)
	assert.Nil(t, err)

	value, err = cache.Retrieve("", "key1")
	assert.Equal(t, "value1", value)
	assert.Nil(t, err)

	ok, err = cache.RetrieveAs("", "key1", &str)
	assert.True(t, ok)
	assert.Equal(t, "value1", str)
	assert.Nil(t, err)

	time.Sleep(500 * time.Millisecond)

	value, err = cache.Retrieve("", "key1")
	assert.Nil(t, value)
	assert.Nil(t, err)
}
