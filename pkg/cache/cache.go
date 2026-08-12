// Package cache defines the bounded in-process cache used on the redirect
// hot path.
package cache

// Cache is a bounded key/value store with an eviction policy of the
// implementation's choosing.
//
// Implementations are expected to be safe for concurrent use; the redirect
// path reads from many goroutines at once.
type Cache[K comparable, V any] interface {
	// Get returns the value stored under key. The boolean reports presence.
	Get(key K) (V, bool)

	// Put stores value under key, evicting another entry if the cache is full.
	Put(key K, value V)

	// Delete removes key. Deleting an absent key is not an error.
	Delete(key K)

	// Len reports how many entries are currently held.
	Len() int
}

// NoopCache satisfies Cache and stores nothing: every Get is a miss.
//
// Implemented in milestone M4.
type NoopCache[K comparable, V any] struct{}

var _ Cache[string, int] = NoopCache[string, int]{}

// Get always reports a miss.
func (NoopCache[K, V]) Get(K) (V, bool) {
	var zero V
	return zero, false
}

// Put discards the entry.
func (NoopCache[K, V]) Put(K, V) {}

// Delete does nothing.
func (NoopCache[K, V]) Delete(K) {}

// Len always returns zero.
func (NoopCache[K, V]) Len() int { return 0 }
