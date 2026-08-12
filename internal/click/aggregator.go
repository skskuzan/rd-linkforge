package click

import "context"

// Aggregator answers the analytics queries the API exposes.
//
// Implementations own their counter state in a single goroutine and serve
// reads over a request channel, so no lock guards the multi-field structure.
type Aggregator interface {
	// Series returns the click time series for one code.
	Series(ctx context.Context, code string, w Window, g Granularity) ([]Bucket, error)

	// Top returns the n most-clicked links within w, highest first.
	Top(ctx context.Context, w Window, n int) ([]Total, error)
}

// NoopAggregator satisfies Aggregator and counts nothing.
//
// Implemented in milestone M5.
type NoopAggregator struct{}

var _ Aggregator = NoopAggregator{}

// Series returns no buckets.
func (NoopAggregator) Series(context.Context, string, Window, Granularity) ([]Bucket, error) {
	return nil, nil
}

// Top returns no totals.
func (NoopAggregator) Top(context.Context, Window, int) ([]Total, error) { return nil, nil }
