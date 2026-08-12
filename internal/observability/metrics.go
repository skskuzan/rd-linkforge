// Package observability defines the metrics contract used across the service.
//
// Logging is not abstracted here: slog is the standard library's answer and
// wrapping it buys nothing.
package observability

import "time"

// Metrics records the counters and histograms the service exports.
//
// The interface exists so that domain and pipeline code can be tested without
// a metrics registry, and so a build can run with metrics switched off.
type Metrics interface {
	// ObserveRequest records the outcome and duration of one request.
	ObserveRequest(route string, status int, d time.Duration)

	// ObserveCacheLookup records whether a redirect cache lookup hit.
	ObserveCacheLookup(hit bool)

	// ObserveClickDropped counts one event the pipeline could not accept.
	ObserveClickDropped()

	// SetQueueDepth reports the current depth of the ingestion queue.
	SetQueueDepth(n int)
}

// NoopMetrics satisfies Metrics and records nothing.
//
// Implemented in milestone M10.
type NoopMetrics struct{}

var _ Metrics = NoopMetrics{}

// ObserveRequest does nothing.
func (NoopMetrics) ObserveRequest(string, int, time.Duration) {}

// ObserveCacheLookup does nothing.
func (NoopMetrics) ObserveCacheLookup(bool) {}

// ObserveClickDropped does nothing.
func (NoopMetrics) ObserveClickDropped() {}

// SetQueueDepth does nothing.
func (NoopMetrics) SetQueueDepth(int) {}
