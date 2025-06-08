package spotify

import (
	"context"
	"errors"

	"github.com/xhos/xhos.dev.backend/internal/models"
	"github.com/zmb3/spotify/v2"
)

var (
	ErrMissingUserID      = errors.New("missing user ID")
	ErrSpotifyCredentials = errors.New("couldn't get spotify token")
	ErrUserNotFound       = errors.New("spotify user not found")
)

func (p *Provider) ExtractUserData(id string, name string, urls map[string]string, images []spotify.Image) *models.User {
	userData := &models.User{
		ID:       id,
		Name:     name,
		Platform: "spotify",
	}

	// Extract profile URL
	if external, ok := urls["spotify"]; ok {
		userData.ProfileURL = external
	}

	// Extract image
	if len(images) > 0 {
		userData.ImageURL = images[0].URL

		for _, img := range images {
			if img.Height == 300 && img.Width == 300 {
				userData.ImageURL = img.URL
				break
			}
		}
	}

	return userData
}

// GetUser fetches a Spotify user's public profile data with simplified response
func (p *Provider) GetUser(ctx context.Context, userID string) (*models.User, error) {
	if userID == "" {
		return nil, ErrInvalidID
	}

	client, err := p.getClient(ctx)
	if err != nil {
		return nil, ErrSpotifyAPI
	}

	// Get user profile from Spotify API
	user, err := client.GetUsersPublicProfile(ctx, spotify.ID(userID))
	if err != nil {
		return nil, err
	}

	return p.ExtractUserData(user.ID, user.DisplayName, user.ExternalURLs, user.Images), nil
}

// GetMe fetches the current Spotify user's profile data with simplified response
func (p *Provider) GetMe(ctx context.Context) (*models.User, error) {
	client, err := p.getClient(ctx)
	if err != nil {
		return nil, ErrSpotifyAPI
	}

	// Get current user profile from Spotify API
	user, err := client.CurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	return p.ExtractUserData(user.ID, user.DisplayName, user.ExternalURLs, user.Images), nil
}
