package cache

import "context"

// ICache interface for caches that are used to cache
// values to improve performance.
type ICache[T any] interface {

	// Retrieve cached value from the cache using its key.
	// If value is missing in the cache or expired it returns nil.
	Retrieve(ctx context.Context, correlationId string, key string) (T, error)

	// Store value in the cache with expiration time.
	Store(ctx context.Context, correlationId string, key string, value T, timeout int64) (T, error)

	// Remove a value from the cache by its key.
	Remove(ctx context.Context, correlationId string, key string) error

	// Contains check is value stores
	Contains(ctx context.Context, correlationId string, key string) bool
}
