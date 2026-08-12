// Package transport defines the lifecycle contract shared by every listener:
// HTTP, WebSocket, TCP and admin.
package transport

import "context"

// Server is one network listener with a managed lifecycle. A uniform contract
// across listeners is what makes a single coordinated graceful shutdown work.
type Server interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
	Addr() string
}

// NoopServer satisfies Server and listens on nothing. Implemented in M6 (TCP)
// and M8 (HTTP).
type NoopServer struct{}

var _ Server = NoopServer{}

func (NoopServer) Start(context.Context) error { return nil }

func (NoopServer) Shutdown(context.Context) error { return nil }

func (NoopServer) Addr() string { return "" }
