// Package transport defines the lifecycle contract shared by every listener:
// HTTP, WebSocket, the TCP resolver and the admin endpoint.
package transport

import "context"

// Server is one network listener with a managed lifecycle.
//
// Uniform Start/Shutdown across all four listeners is what makes a single
// coordinated graceful shutdown possible: cancel the root context, then wait
// on each server's Shutdown within one deadline.
type Server interface {
	// Start begins serving and blocks until the listener stops.
	Start(ctx context.Context) error

	// Shutdown stops accepting new work and drains what is in flight,
	// returning early if ctx expires first.
	Shutdown(ctx context.Context) error

	// Addr reports the address the server listens on.
	Addr() string
}

// NoopServer satisfies Server and listens on nothing.
//
// Implemented in milestones M6 (TCP) and M8 (HTTP).
type NoopServer struct{}

var _ Server = NoopServer{}

// Start returns immediately.
func (NoopServer) Start(context.Context) error { return nil }

// Shutdown returns immediately.
func (NoopServer) Shutdown(context.Context) error { return nil }

// Addr returns the empty string.
func (NoopServer) Addr() string { return "" }
