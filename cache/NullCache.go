package cache

/*
Dummy cache implementation that doesn't do anything.

It can be used in testing or in situations when cache is required but shall be disabled.
*/
type NullCache struct{}

// Creates a new instance of the cache.
// Returns *NullCache
func NewNullCache() *NullCache {
	return &NullCache{}
}

// Retrieves cached value from the cache using its key. If value is missing in the cache or expired it returns null.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - key string
//   a unique value key.
// Returns interface{}, error
func (c *NullCache) Retrieve(correlationId string, key string) (interface{}, error) {
	return nil, nil
}

// Retrieve cached value from the cache using its key and restore into reference object. If value is missing in the cache or expired it returns false.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - key string   a unique value key.
//   - refObj       pointer to object for restore
// Returns bool, error
func (c *NullCache) RetrieveAs(correlationId string, key string, refObj interface{}) (bool, error) {
	return false, nil
}

// Stores value in the cache with expiration time, if success return stored value.
// Parameters:
//   - correlationId string
//    transaction id to trace execution through call chain.
//   - key string
//   a unique value key.
//   - value interface{}
//   a value to store.
//   - timeout int64
//   expiration timeout in milliseconds.
// Returns interface{}, error
func (c *NullCache) Store(correlationId string, key string, value interface{}, timeout int64) (interface{}, error) {
	return value, nil
}

// Removes a value from the cache by its key.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.\
//   - key string
//   a unique value key.
// Returns error
func (c *NullCache) Remove(correlationId string, key string) error {
	return nil
}
