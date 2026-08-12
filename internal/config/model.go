// Package config defines how runtime configuration is obtained.
//
// Configuration is environment-only and validated at startup: a malformed or
// missing value fails the process rather than surfacing as a runtime error.
package config

import "time"

// Config is the fully resolved runtime configuration.
type Config struct {
	HTTPAddr  string
	AdminAddr string
	TCPAddr   string

	PostgresDSN string
	MongoURI    string

	JWTSecret   string
	TokenTTL    time.Duration
	CacheSize   int
	LogLevel    string
	ShutdownTTL time.Duration
}
