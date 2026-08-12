// Package observability defines the metrics contract used across the service.
// Logging is not abstracted: slog is the standard answer and wrapping it buys
// nothing.
package observability

import "time"

// Metrics records the counters and histograms the service exports. The
// interface exists so domain code can be tested without a registry.
type Metrics interface {
	ObserveRequest(route string, status int, d time.Duration)
	ObserveCacheLookup(hit bool)
	ObserveClickDropped()
	SetQueueDepth(n int)
}

// NoopMetrics satisfies Metrics and records nothing. Implemented in M10.
type NoopMetrics struct{}

var _ Metrics = NoopMetrics{}

func (NoopMetrics) ObserveRequest(string, int, time.Duration) {}

func (NoopMetrics) ObserveCacheLookup(bool) {}

func (NoopMetrics) ObserveClickDropped() {}

func (NoopMetrics) SetQueueDepth(int) {}
