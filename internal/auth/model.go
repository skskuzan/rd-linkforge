// Package auth holds the identity model and the contracts for issuing and
// verifying access tokens.
package auth

import "time"

// User is an account that owns links. PasswordHash is never serialised onto a
// transport.
type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

// Credentials is a login attempt.
type Credentials struct {
	Email    string
	Password string
}

// Claims is the verified content of an access token.
type Claims struct {
	UserID    string
	Email     string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

// Token is an issued access token and its expiry.
type Token struct {
	AccessToken string
	ExpiresAt   time.Time
}
