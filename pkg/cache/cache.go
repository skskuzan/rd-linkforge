// Package cache defines the bounded in-process cache used on the redirect hot
// path.
package cache

// Cache is a bounded key/value store with an eviction policy of the
// implementation's choosing. Implementations must be safe for concurrent use.
type Cache[K comparable, V any] interface {
	Get(key K) (V, bool)
	Put(key K, value V)
	Delete(key K)
	Len() int
}

// NoopCache satisfies Cache and stores nothing: every Get misses. Implemented
// in M4.
type NoopCache[K comparable, V any] struct{}

var _ Cache[string, int] = NoopCache[string, int]{}

func (NoopCache[K, V]) Get(K) (V, bool) {
	var zero V
	return zero, false
}

func (NoopCache[K, V]) Put(K, V) {}

func (NoopCache[K, V]) Delete(K) {}

func (NoopCache[K, V]) Len() int { return 0 }
