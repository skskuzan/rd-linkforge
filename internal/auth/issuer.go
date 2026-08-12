package auth

import "context"

// Issuer mints access tokens for authenticated users.
type Issuer interface {
	Issue(ctx context.Context, u User) (Token, error)
}

// NoopIssuer satisfies Issuer and mints nothing. Implemented in M9.
type NoopIssuer struct{}

var _ Issuer = NoopIssuer{}

func (NoopIssuer) Issue(context.Context, User) (Token, error) { return Token{}, nil }
