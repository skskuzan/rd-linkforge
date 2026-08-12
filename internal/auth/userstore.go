package auth

import "context"

// UserStore is the persistence contract the auth service requires.
type UserStore interface {
	Create(ctx context.Context, u User) (User, error)
	ByEmail(ctx context.Context, email string) (User, error)
}

// NoopUserStore satisfies UserStore and persists nothing. Implemented in M7.
type NoopUserStore struct{}

var _ UserStore = NoopUserStore{}

func (NoopUserStore) Create(_ context.Context, u User) (User, error) { return u, nil }

func (NoopUserStore) ByEmail(context.Context, string) (User, error) { return User{}, nil }
