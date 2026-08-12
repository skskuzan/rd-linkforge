package link

import "context"

// Store is the persistence contract that Shortener requires.
//
// It is declared here, in the consuming package, rather than beside any
// implementation. Storage constructors return concrete types; each consumer
// depends only on the narrow slice of behaviour it uses.
type Store interface {
	// Create persists l and returns it with any fields the store assigns.
	Create(ctx context.Context, l Link) (Link, error)

	// ByCode returns the link with the given code.
	ByCode(ctx context.Context, code string) (Link, error)

	// ListByOwner returns one page of the owner's links, newest first.
	ListByOwner(ctx context.Context, ownerID string, page Page) ([]Link, error)

	// Delete removes the link with the given id.
	Delete(ctx context.Context, id int64) error
}

// NoopStore satisfies Store and persists nothing.
//
// Implemented in milestone M1 (memory) and M7 (Postgres).
type NoopStore struct{}

var _ Store = NoopStore{}

// Create returns the link unchanged.
func (NoopStore) Create(_ context.Context, l Link) (Link, error) { return l, nil }

// ByCode returns the zero Link and no error.
func (NoopStore) ByCode(context.Context, string) (Link, error) { return Link{}, nil }

// ListByOwner returns no links.
func (NoopStore) ListByOwner(context.Context, string, Page) ([]Link, error) { return nil, nil }

// Delete does nothing.
func (NoopStore) Delete(context.Context, int64) error { return nil }
