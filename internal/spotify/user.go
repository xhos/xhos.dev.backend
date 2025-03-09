package spotify

import (
	"context"
	"errors"
	"os"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

// UserData represents the public profile information of a Spotify user
type UserData struct {
	UserID      string          `json:"user_id"`
	DisplayName string          `json:"display_name"`
	SpotifyURI  string          `json:"spotify_uri"`
	Endpoint    string          `json:"endpoint"`
	Followers   int             `json:"followers"`
	ProfileURL  string          `json:"profile_url,omitempty"`
	Images      []spotify.Image `json:"images,omitempty"`
}

var (
	ErrMissingUserID      = errors.New("missing user ID")
	ErrSpotifyCredentials = errors.New("couldn't get spotify token")
	ErrUserNotFound       = errors.New("spotify user not found")
)

// GetUserData fetches a Spotify user's public profile data
func GetUserData(userID string) (*UserData, error) {
	if userID == "" {
		return nil, ErrMissingUserID
	}

	ctx := context.Background()

	// Get client credentials from environment
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}

	token, err := config.Token(ctx)
	if err != nil {
		return nil, ErrSpotifyCredentials
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	user, err := client.GetUsersPublicProfile(ctx, spotify.ID(userID))
	if err != nil {
		return nil, err
	}

	// Create UserData object from Spotify user data
	userData := &UserData{
		UserID:      user.ID,
		DisplayName: user.DisplayName,
		SpotifyURI:  string(user.URI),
		Endpoint:    user.Endpoint,
		Followers:   int(user.Followers.Count),
		Images:      user.Images,
	}

	if external, ok := user.ExternalURLs["spotify"]; ok {
		userData.ProfileURL = external
	}

	return userData, nil
}
