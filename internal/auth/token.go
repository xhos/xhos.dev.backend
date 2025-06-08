package auth

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

// StoreTokenInCookie stores an OAuth token in a HTTP cookie
func StoreTokenInCookie(w http.ResponseWriter, r *http.Request, name string, token *oauth2.Token) error {
	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    base64.URLEncoding.EncodeToString(tokenBytes),
		Path:     "/",
		MaxAge:   int(token.Expiry.Sub(time.Now()).Seconds()),
		HttpOnly: true,
		Secure:   r.TLS != nil,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}

// GetTokenFromCookie retrieves an OAuth token from a cookie with improved error handling
func GetTokenFromCookie(r *http.Request, name string) (*oauth2.Token, error) {
	// Get cookie
	cookie, err := r.Cookie(name)
	if err != nil {
		return nil, fmt.Errorf("authentication cookie not found: %w", err)
	}

	// Decode
	tokenBytes, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to decode token: %w", err)
	}

	// Unmarshal
	var token oauth2.Token
	if err := json.Unmarshal(tokenBytes, &token); err != nil {
		return nil, fmt.Errorf("invalid token format: %w", err)
	}

	return &token, nil
}
