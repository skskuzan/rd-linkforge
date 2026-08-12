package click

import "context"

// Store is the append-only persistence contract for raw click events.
//
// Writes arrive in batches from the pipeline's batcher, never one event at a
// time.
type Store interface {
	// Append writes a batch of events.
	Append(ctx context.Context, batch []Click) error

	// Range returns the events recorded for one code within w, oldest first.
	Range(ctx context.Context, code string, w Window) ([]Click, error)
}

// NoopStore satisfies Store and persists nothing.
//
// Implemented in milestone M7.
type NoopStore struct{}

var _ Store = NoopStore{}

// Append discards the batch.
func (NoopStore) Append(context.Context, []Click) error { return nil }

// Range returns no events.
func (NoopStore) Range(context.Context, string, Window) ([]Click, error) { return nil, nil }
