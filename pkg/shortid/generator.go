package shortid

// Generator issues fresh short codes.
//
// Implementations are expected to place a counter in the low bits and CSPRNG
// output in the high bits, so that the keyspace cannot be walked by
// incrementing a known code.
type Generator interface {
	// Next returns a code that has not been issued before by this generator.
	Next() (string, error)
}

// NoopGenerator satisfies Generator and does nothing.
//
// Implemented in milestone M4.
type NoopGenerator struct{}

var _ Generator = NoopGenerator{}

// Next returns the empty string and no error.
func (NoopGenerator) Next() (string, error) { return "", nil }
