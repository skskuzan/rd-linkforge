package click

// Ingester accepts click events from the redirect path.
//
// Submit must never block. When the pipeline is saturated the event is
// dropped and reported, because analytics may not apply backpressure to a
// redirect.
type Ingester interface {
	// Submit hands an event to the pipeline. The boolean reports whether it
	// was accepted; false means the event was dropped.
	Submit(c Click) bool
}

// NoopIngester satisfies Ingester and discards every event.
//
// Implemented in milestone M5.
type NoopIngester struct{}

var _ Ingester = NoopIngester{}

// Submit discards the event and reports acceptance.
func (NoopIngester) Submit(Click) bool { return true }
