package auth

import "context"

// UserStore is the persistence contract the auth service requires.
type UserStore interface {
	// Create persists a new user.
	Create(ctx context.Context, u User) (User, error)

	// ByEmail returns the user registered under an email address.
	ByEmail(ctx context.Context, email string) (User, error)
}

// NoopUserStore satisfies UserStore and persists nothing.
//
// Implemented in milestone M7.
type NoopUserStore struct{}

var _ UserStore = NoopUserStore{}

// Create returns the user unchanged.
func (NoopUserStore) Create(_ context.Context, u User) (User, error) { return u, nil }

// ByEmail returns the zero User and no error.
func (NoopUserStore) ByEmail(context.Context, string) (User, error) { return User{}, nil }
