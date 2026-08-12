package click

import "context"

// Aggregator answers the analytics queries the API exposes. Implementations
// own their counters in a single goroutine and serve reads over a channel.
type Aggregator interface {
	Series(ctx context.Context, code string, w Window, g Granularity) ([]Bucket, error)
	Top(ctx context.Context, w Window, n int) ([]Total, error)
}

// NoopAggregator satisfies Aggregator and counts nothing. Implemented in M5.
type NoopAggregator struct{}

var _ Aggregator = NoopAggregator{}

func (NoopAggregator) Series(context.Context, string, Window, Granularity) ([]Bucket, error) {
	return nil, nil
}

func (NoopAggregator) Top(context.Context, Window, int) ([]Total, error) { return nil, nil }
