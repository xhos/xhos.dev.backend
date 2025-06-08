package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xhos/xhos.dev.backend/internal/auth"
	"github.com/xhos/xhos.dev.backend/internal/tools"
	spotifyapi "github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const redirectURI = "http://localhost:8080/auth/spotify/callback"

// generateRandomState creates a secure random state to protect against CSRF
func generateRandomState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// InitiateSpotifyAuth starts the Spotify OAuth flow
func InitiateSpotifyAuth(w http.ResponseWriter, r *http.Request) {
	authenticator := spotifyauth.New(
		spotifyauth.WithClientID(tools.GetSpotifyID()),
		spotifyauth.WithClientSecret(tools.GetSpotifySecret()),
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopePlaylistReadPrivate,
			spotifyauth.ScopePlaylistModifyPrivate,
			spotifyauth.ScopePlaylistModifyPublic,
		),
	)

	// Generate a random state value for CSRF protection
	state, err := generateRandomState()
	if err != nil {
		logrus.Error("Failed to generate state:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Store state in a cookie for verification in the callback
	http.SetCookie(w, &http.Cookie{
		Name:     "spotify_auth_state",
		Value:    state,
		Path:     "/",
		MaxAge:   600, // 10 minutes
		HttpOnly: true,
		Secure:   r.TLS != nil,
		SameSite: http.SameSiteLaxMode,
	})

	// Redirect to Spotify's authorization page
	authURL := authenticator.AuthURL(state)
	http.Redirect(w, r, authURL, http.StatusFound)
}

// HandleSpotifyCallback processes the callback from Spotify after authorization
func HandleSpotifyCallback(w http.ResponseWriter, r *http.Request) {
	// Get the state from cookie
	cookie, err := r.Cookie("spotify_auth_state")
	if err != nil {
		logrus.Error("No state cookie:", err)
		http.Error(w, "State verification failed", http.StatusBadRequest)
		return
	}

	// Compare the state parameter
	if r.URL.Query().Get("state") != cookie.Value {
		logrus.Error("State mismatch")
		http.Error(w, "State verification failed", http.StatusBadRequest)
		return
	}

	// Create the authenticator with the same parameters
	authenticator := spotifyauth.New( // Renamed from "auth" to "authenticator"
		spotifyauth.WithClientID(tools.GetSpotifyID()),
		spotifyauth.WithClientSecret(tools.GetSpotifySecret()),
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeUserReadEmail,
			spotifyauth.ScopePlaylistReadPrivate,
			spotifyauth.ScopePlaylistModifyPrivate,
			spotifyauth.ScopePlaylistModifyPublic,
		),
	)

	// Get token
	token, err := authenticator.Token(r.Context(), cookie.Value, r) // Updated variable name
	if err != nil {
		logrus.Error("Couldn't get token:", err)
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}

	// Store token using the auth module
	if err := auth.StoreTokenInCookie(w, r, "spotify_token", token); err != nil { // Fixed if statement
		logrus.Error("Failed to store token:", err)
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}

	// Also store the user ID for easy access
	client := spotifyapi.New(authenticator.Client(r.Context(), token)) // Updated variable name
	user, err := client.CurrentUser(r.Context())
	if err == nil {
		http.SetCookie(w, &http.Cookie{
			Name:     "spotify_user_id",
			Value:    user.ID,
			Path:     "/",
			MaxAge:   int(token.Expiry.Sub(time.Now()).Seconds()),
			HttpOnly: true,
			Secure:   r.TLS != nil,
			SameSite: http.SameSiteLaxMode,
		})
	}

	// Redirect as before
	frontendURL := "http://localhost:3000/sort"
	http.Redirect(w, r, frontendURL, http.StatusFound)
}
