// Package shortid defines the contracts for turning link identifiers into the
// base62 codes that appear in short URLs, and back.
package shortid

// Codec converts between a numeric link identifier and its short code.
//
// Base62 is the intended encoding: the codes travel in URL paths, so the
// alphabet must contain no character that needs escaping.
type Codec interface {
	// Encode renders n as a short code.
	Encode(n uint64) string

	// Decode parses a short code back into the value it was encoded from.
	Decode(code string) (uint64, error)
}

// NoopCodec satisfies Codec and does nothing.
//
// Implemented in milestone M1.
type NoopCodec struct{}

var _ Codec = NoopCodec{}

// Encode returns the empty string.
func (NoopCodec) Encode(uint64) string { return "" }

// Decode returns the zero value and no error.
func (NoopCodec) Decode(string) (uint64, error) { return 0, nil }
