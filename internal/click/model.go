// Package click holds the click-event model and the contracts of the
// analytics path.
//
// Events are append-only and best-effort: losing some of them degrades
// analytics but must never break or slow a redirect.
package click

import "time"

// Click records one follow of a short link.
//
// The raw client address is deliberately absent. Only a salted hash is
// carried, and the hashing happens at the transport edge where the address is
// first seen.
type Click struct {
	Code      string
	At        time.Time
	IPHash    string
	Referer   string
	UserAgent string
	Country   string
}

// Granularity is the bucket width of a time series.
type Granularity string

// Supported series granularities.
const (
	GranularityHour Granularity = "hour"
	GranularityDay  Granularity = "day"
)

// Window is a half-open time range [From, To).
type Window struct {
	From time.Time
	To   time.Time
}

// Bucket is one point of a click time series.
type Bucket struct {
	Start  time.Time
	Clicks int64
}

// Total is a link's aggregate click count, used for top-N rankings.
type Total struct {
	Code   string
	Clicks int64
}
