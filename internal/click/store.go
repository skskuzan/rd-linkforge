package click

import "context"

// Store is the append-only persistence contract for raw events. Writes arrive
// in batches from the pipeline, never one event at a time.
type Store interface {
	Append(ctx context.Context, batch []Click) error
	Range(ctx context.Context, code string, w Window) ([]Click, error)
}

// NoopStore satisfies Store and persists nothing. Implemented in M7.
type NoopStore struct{}

var _ Store = NoopStore{}

func (NoopStore) Append(context.Context, []Click) error { return nil }

func (NoopStore) Range(context.Context, string, Window) ([]Click, error) { return nil, nil }
