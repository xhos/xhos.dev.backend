package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/xhos/xhos.dev.backend/internal/auth"
	"github.com/xhos/xhos.dev.backend/internal/factory"
	"github.com/xhos/xhos.dev.backend/internal/music/spotify"
)

// GetSpotifyProfile returns the current user's Spotify profile
func GetSpotifyProfile(w http.ResponseWriter, r *http.Request) {
	// Get provider instance
	providerInstance, err := factory.GetProvider("spotify")
	if err != nil {
		sendErrorResponse(w, errors.New("Provider not available"), http.StatusInternalServerError)
		return
	}

	spotifyProvider := providerInstance.(*spotify.Provider)

	// Try to use GetMe directly first (simpler approach)
	userData, err := spotifyProvider.GetMe(r.Context())

	// If client credentials approach fails (as expected), fall back to token-based auth
	if err != nil {
		// Use auth module to get token from cookie
		token, err := auth.GetTokenFromCookie(r, "spotify_token")
		if err != nil {
			sendErrorResponse(w, errors.New("Not authenticated with Spotify"), http.StatusUnauthorized)
			return
		}

		// Use getUserClient from the provider with the user's token
		client, err := spotifyProvider.GetUserClient(r.Context(), token)
		if err != nil {
			sendErrorResponse(w, errors.New("Authentication failed"), http.StatusUnauthorized)
			return
		}

		// Get current user profile
		user, err := client.CurrentUser(r.Context())
		if err != nil {
			sendErrorResponse(w, errors.New("Failed to fetch profile"), http.StatusUnauthorized)
			return
		}

		// Use GetUser method with the ID we retrieved
		userData, err = spotifyProvider.GetUser(r.Context(), user.ID)
		if err != nil {
			sendErrorResponse(w, errors.New("Failed to process user data"), http.StatusInternalServerError)
			return
		}
	}

	// Return the user profile data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}
