package test_cache

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pip-services3-gox/pip-services3-components-gox/cache"
)

func TestNullCache(t *testing.T) {
	cache := cache.NewNullCache()

	value, err := cache.Retrieve("", "key1")
	assert.Nil(t, value)
	assert.Nil(t, err)

	var str string
	ok, err := cache.RetrieveAs("", "key1", &str)
	assert.False(t, ok)
	assert.Equal(t, str, "")
	assert.Nil(t, err)

	value, err = cache.Store("", "key1", "value1", 0)
	assert.Equal(t, "value1", value)
	assert.Nil(t, err)
}
