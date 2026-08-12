package auth

import "context"

// Issuer mints access tokens for authenticated users.
type Issuer interface {
	// Issue returns a signed token for u.
	Issue(ctx context.Context, u User) (Token, error)
}

// NoopIssuer satisfies Issuer and mints nothing.
//
// Implemented in milestone M9.
type NoopIssuer struct{}

var _ Issuer = NoopIssuer{}

// Issue returns the zero Token and no error.
func (NoopIssuer) Issue(context.Context, User) (Token, error) { return Token{}, nil }
