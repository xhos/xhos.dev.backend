package handlers

import (
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /auth/spotify/login", http.HandlerFunc(InitiateSpotifyAuth))
	router.Handle("GET /auth/spotify/callback", http.HandlerFunc(HandleSpotifyCallback))

	router.Handle("GET /spotify/me", http.HandlerFunc(GetSpotifyProfile))
	// Dashboard endpoints for music management
	// router.Handle("GET /{provider}/{userID}", http.HandlerFunc(GetUserProfile))
	// router.Handle("GET /{provider}/{userID}/playlists", http.HandlerFunc(GetUserPlaylists))
	// router.Handle("GET /{provider}/playlist/{playlistID}/tracks", http.HandlerFunc(GetPlaylistTracks))
	// router.Handle("GET /dashboard/search", http.HandlerFunc(SearchTracks))

	// Playlist modification endpoints
	// router.Handle("POST /dashboard/{provider}/playlist/{playlistID}/tracks", http.HandlerFunc(AddTracks))
	// router.Handle("DELETE /dashboard/{provider}/playlist/{playlistID}/tracks", http.HandlerFunc(RemoveTracks))
	// router.Handle("PUT /dashboard/{provider}/playlist/{playlistID}/tracks", http.HandlerFunc(UpdatePlaylistTracks))

	return router
}
