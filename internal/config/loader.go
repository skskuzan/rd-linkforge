package config

import "context"

// Loader resolves configuration from its source.
//
// Keeping this behind an interface lets tests supply a fixed Config without
// mutating the process environment.
type Loader interface {
	// Load reads and validates the configuration.
	Load(ctx context.Context) (Config, error)
}

// NoopLoader satisfies Loader and returns an empty configuration.
//
// Implemented in milestone M3.
type NoopLoader struct{}

var _ Loader = NoopLoader{}

// Load returns the zero Config and no error.
func (NoopLoader) Load(context.Context) (Config, error) { return Config{}, nil }
