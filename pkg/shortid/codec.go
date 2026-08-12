// Package shortid converts link identifiers to the base62 codes used in short
// URLs, and back.
package shortid

// Codec converts between a numeric link identifier and its short code.
type Codec interface {
	Encode(n uint64) string
	Decode(code string) (uint64, error)
}

// NoopCodec satisfies Codec and does nothing. Implemented in M1.
type NoopCodec struct{}

var _ Codec = NoopCodec{}

func (NoopCodec) Encode(uint64) string { return "" }

func (NoopCodec) Decode(string) (uint64, error) { return 0, nil }
