// Package link holds the Linkforge domain model and the contracts its
// consumers require.
//
// Nothing here may import a transport or a storage package: the dependency
// arrow points inwards. Interfaces are declared where they are consumed, and
// implementations return concrete types.
package link

import "time"

// Link is a shortened URL.
//
// A zero ExpiresAt means the link never expires.
type Link struct {
	ID        int64
	Code      string
	TargetURL string
	OwnerID   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// ShortenRequest carries everything needed to create a link.
//
// An empty Alias asks for a generated code.
type ShortenRequest struct {
	TargetURL string
	Alias     string
	OwnerID   string
	ExpiresAt time.Time
}

// Page describes one slice of a cursor-paginated listing.
type Page struct {
	Cursor string
	Limit  int
}
