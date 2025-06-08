package models

import "context"

// Provider defines operations that any music service must implement
type Provider interface {
	Name() string

	GetUser(ctx context.Context, userID string) (*User, error)
	GetMe(ctx context.Context) (*User, error)
}
