package click

// Publisher fans processed events out to live WebSocket subscribers.
//
// A slow subscriber must not stall the pipeline: implementations drop frames
// for a consumer that cannot keep up rather than blocking the publisher.
type Publisher interface {
	// Publish offers an event to every current subscriber.
	Publish(c Click)

	// Subscribe returns a channel of events for one owner's links, and a
	// cancel function that closes it and releases the subscription.
	Subscribe(ownerID string) (<-chan Click, func())
}

// NoopPublisher satisfies Publisher and delivers nothing.
//
// Implemented in milestone M9.
type NoopPublisher struct{}

var _ Publisher = NoopPublisher{}

// Publish discards the event.
func (NoopPublisher) Publish(Click) {}

// Subscribe returns a closed channel and a cancel function that does nothing.
func (NoopPublisher) Subscribe(string) (<-chan Click, func()) {
	ch := make(chan Click)
	close(ch)
	return ch, func() {}
}
