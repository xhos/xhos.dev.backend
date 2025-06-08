package spotify

import (
	"context"
	"errors"

	"github.com/xhos/xhos.dev.backend/internal/tools"
	spotifyapi "github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2" // Add this import
	"golang.org/x/oauth2/clientcredentials"
)

var (
	ErrInvalidID  = errors.New("invalid ID")
	ErrSpotifyAPI = errors.New("spotify API error")
)

// Provider implements the models.Provider interface for Spotify
type Provider struct {
	clientID     string
	clientSecret string
}

// New creates a new Spotify provider
func New() *Provider {
	return &Provider{
		clientID:     tools.GetSpotifyID(),
		clientSecret: tools.GetSpotifySecret(),
	}
}

// Name returns the provider name
func (p *Provider) Name() string {
	return "spotify"
}

// getClient creates a Spotify client using client credentials
func (p *Provider) getClient(ctx context.Context) (*spotifyapi.Client, error) {
	config := &clientcredentials.Config{
		ClientID:     p.clientID,
		ClientSecret: p.clientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}

	token, err := config.Token(ctx)
	if err != nil {
		return nil, errors.New("failed to get Spotify auth token")
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotifyapi.New(httpClient)

	return client, nil
}

// getUserClient creates a Spotify client using a user's OAuth token
func (p *Provider) GetUserClient(ctx context.Context, token *oauth2.Token) (*spotifyapi.Client, error) {
	if token == nil {
		return nil, errors.New("no valid token provided")
	}

	auth := spotifyauth.New(
		spotifyauth.WithClientID(p.clientID),
		spotifyauth.WithClientSecret(p.clientSecret),
	)

	client := spotifyapi.New(auth.Client(ctx, token))
	return client, nil
}
