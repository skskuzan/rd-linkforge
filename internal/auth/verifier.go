package auth

import "context"

// Verifier checks an access token and returns what it asserts. This is what
// the authorisation middleware depends on.
type Verifier interface {
	Verify(ctx context.Context, token string) (Claims, error)
}

// NoopVerifier satisfies Verifier and accepts everything, so it must never be
// wired into a running server. Implemented in M9.
type NoopVerifier struct{}

var _ Verifier = NoopVerifier{}

func (NoopVerifier) Verify(context.Context, string) (Claims, error) { return Claims{}, nil }
