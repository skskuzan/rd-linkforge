package link

import "context"

// Shortener is the domain service that every transport sits in front of.
//
// The HTTP handlers, the TCP resolver and the CLI all reach the system through
// this one interface; none of them may talk to a Store directly.
type Shortener interface {
	// Shorten creates a link for the requested target.
	Shorten(ctx context.Context, req ShortenRequest) (Link, error)

	// Resolve returns the link a code points at, if it exists and is live.
	Resolve(ctx context.Context, code string) (Link, error)

	// List returns one page of an owner's links.
	List(ctx context.Context, ownerID string, page Page) ([]Link, error)

	// Delete removes an owner's link by code.
	Delete(ctx context.Context, ownerID, code string) error
}

// NoopShortener satisfies Shortener and does nothing.
//
// Implemented in milestone M2.
type NoopShortener struct{}

var _ Shortener = NoopShortener{}

// Shorten returns the zero Link and no error.
func (NoopShortener) Shorten(context.Context, ShortenRequest) (Link, error) { return Link{}, nil }

// Resolve returns the zero Link and no error.
func (NoopShortener) Resolve(context.Context, string) (Link, error) { return Link{}, nil }

// List returns no links.
func (NoopShortener) List(context.Context, string, Page) ([]Link, error) { return nil, nil }

// Delete does nothing.
func (NoopShortener) Delete(context.Context, string, string) error { return nil }
