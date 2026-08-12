package click

import "context"

// Publisher fans events out to live WebSocket subscribers. Implementations
// drop frames for a subscriber that cannot keep up rather than blocking.
type Publisher interface {
	Publish(c Click)
	Subscribe(ctx context.Context, ownerID string) (<-chan Click, func())
}

// NoopPublisher satisfies Publisher and delivers nothing. Implemented in M9.
type NoopPublisher struct{}

var _ Publisher = NoopPublisher{}

func (NoopPublisher) Publish(Click) {}

func (NoopPublisher) Subscribe(context.Context, string) (<-chan Click, func()) {
	ch := make(chan Click)
	close(ch)
	return ch, func() {}
}
