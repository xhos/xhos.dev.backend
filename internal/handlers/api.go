package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/xhos/xhos.dev.backend/internal/spotify"
)

func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /spotify/{userID}/name", http.HandlerFunc(GetSpotifyUserName))
	router.Handle("GET /spotify/{userID}/url", http.HandlerFunc(GetSpotifyUserProfileURL))
	router.Handle("GET /spotify/{userID}/icon/{size}", http.HandlerFunc(GetSpotifyUserIcon))
	router.Handle("GET /spotify/{userID}/playlists", http.HandlerFunc(GetSpotifyUserPlaylists))

	return router
}

func GetSpotifyUserName(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")

	userData, err := spotify.GetUserData(userID)
	if err != nil {
		handleSpotifyError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"name": userData.DisplayName})
}

func GetSpotifyUserProfileURL(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")

	userData, err := spotify.GetUserData(userID)
	if err != nil {
		handleSpotifyError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"url": userData.ProfileURL})
}

func GetSpotifyUserIcon(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	size := r.PathValue("size")

	userData, err := spotify.GetUserData(userID)
	if err != nil {
		handleSpotifyError(w, err)
		return
	}

	var imageURL string
	for _, image := range userData.Images {
		if size == "64" && image.Height == 64 {
			imageURL = image.URL
			break
		} else if size == "300" && image.Height == 300 {
			imageURL = image.URL
			break
		}
	}

	if imageURL == "" {
		http.Error(w, "Image size not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"url": imageURL})
}

func GetSpotifyUserPlaylists(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")

	// Get limit from query parameters, default to 20 if not specified
	limit := 20
	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		if val, err := strconv.Atoi(limitStr); err == nil {
			limit = val
		}
	}

	playlists, err := spotify.GetUserPlaylists(userID, limit)
	if err != nil {
		handleSpotifyError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playlists)
}

func handleSpotifyError(w http.ResponseWriter, err error) {
	log.Error("Error fetching Spotify data: ", err)
	switch err {
	case spotify.ErrMissingUserID:
		http.Error(w, "Missing user ID", http.StatusBadRequest)
	case spotify.ErrSpotifyCredentials:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	case spotify.ErrUserNotFound:
		http.Error(w, "User not found", http.StatusNotFound)
	case spotify.ErrPlaylistsNotFound:
		http.Error(w, "Playlists not found", http.StatusNotFound)
	default:
		http.Error(w, "Error fetching Spotify data", http.StatusInternalServerError)
	}
}
