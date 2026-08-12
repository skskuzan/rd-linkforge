package auth

import "context"

// Verifier checks an access token and returns what it asserts.
//
// This is the interface the authorisation middleware depends on; it has no
// reason to know how tokens are minted.
type Verifier interface {
	// Verify validates the token's signature and expiry.
	Verify(ctx context.Context, token string) (Claims, error)
}

// NoopVerifier satisfies Verifier and accepts everything.
//
// It must never be wired into a running server. Implemented in milestone M9.
type NoopVerifier struct{}

var _ Verifier = NoopVerifier{}

// Verify returns empty Claims and no error.
func (NoopVerifier) Verify(context.Context, string) (Claims, error) { return Claims{}, nil }
