package factory

import (
	"errors"

	"github.com/xhos/xhos.dev.backend/internal/models"
	"github.com/xhos/xhos.dev.backend/internal/music/spotify"
)

var (
	ErrProviderNotSupported = errors.New("music provider not supported")
)

// GetProvider returns the appropriate provider implementation
func GetProvider(name string) (models.Provider, error) {
	switch name {
	case "spotify":
		return spotify.New(), nil
	// Future providers would be added here:
	// case "soundcloud":
	//    return soundcloud.New(), nil
	default:
		return nil, ErrProviderNotSupported
	}
}
