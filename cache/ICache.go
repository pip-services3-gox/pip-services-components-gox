package cache

/*
Interface for caches that are used to cache values to improve performance.
*/
type ICache interface {
	// Retrieves cached value from the cache using its key. If value is missing in the cache or expired it returns nil.
	Retrieve(correlationId string, key string) (interface{}, error)
	// Retrieves cached value from the cache using its key into reference object. If value is missing in the cache or expired it returns false.
	RetrieveAs(correlationId string, key string, refObj interface{}) (bool, error)
	// Stores value in the cache with expiration time.
	Store(correlationId string, key string, value interface{}, timeout int64) (interface{}, error)
	// Removes a value from the cache by its key.
	Remove(correlationId string, key string) error
}
