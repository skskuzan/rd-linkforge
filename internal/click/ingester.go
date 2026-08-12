package click

// Ingester accepts click events from the redirect path. Submit must never
// block; it drops under saturation and reports false rather than applying
// backpressure to a redirect.
type Ingester interface {
	Submit(c Click) bool
}

// NoopIngester satisfies Ingester and discards every event. Implemented in M5.
type NoopIngester struct{}

var _ Ingester = NoopIngester{}

func (NoopIngester) Submit(Click) bool { return true }
