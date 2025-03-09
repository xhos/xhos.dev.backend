package spotify

import (
	"context"
	"errors"

	"github.com/xhos/xhos.dev.backend/internal/tools"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

// PlaylistData represents simple playlist information from Spotify
type PlaylistData struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Images      []spotify.Image `json:"images"`
	TracksTotal int             `json:"tracks_total"`
	Public      bool            `json:"public"`
	URL         string          `json:"url,omitempty"`
}

// PlaylistsResponse represents the response structure for playlists
type PlaylistsResponse struct {
	Playlists []PlaylistData `json:"playlists"`
	Total     int            `json:"total"`
}

var (
	ErrPlaylistsNotFound = errors.New("playlists not found")
)

// GetUserPlaylists fetches playlists for a Spotify user
func GetUserPlaylists(userID string, limit int) (*PlaylistsResponse, error) {
	if userID == "" {
		return nil, ErrMissingUserID
	}

	ctx := context.Background()

	// Get client credentials from environment
	config := &clientcredentials.Config{
		ClientID:     tools.GetSpotifyID(),
		ClientSecret: tools.GetSpotifySecret(),
		TokenURL:     spotifyauth.TokenURL,
	}

	token, err := config.Token(ctx)
	if err != nil {
		return nil, ErrSpotifyCredentials
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	// Set default limit if not specified or invalid
	if limit <= 0 || limit > 50 {
		limit = 20
	}

	// Get user's playlists with limit
	playlistPage, err := client.GetPlaylistsForUser(ctx, userID, spotify.Limit(limit))
	if err != nil {
		return nil, err
	}

	// Convert to our response format
	response := &PlaylistsResponse{
		Total:     int(playlistPage.Total), // Convert Numeric to int
		Playlists: make([]PlaylistData, len(playlistPage.Playlists)),
	}

	for i, p := range playlistPage.Playlists {
		playlistData := PlaylistData{
			ID:          string(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Images:      p.Images,
			TracksTotal: int(p.Tracks.Total), // Convert Numeric to int
			Public:      p.IsPublic,
		}

		if external, ok := p.ExternalURLs["spotify"]; ok {
			playlistData.URL = external
		}

		response.Playlists[i] = playlistData
	}

	return response, nil
}
