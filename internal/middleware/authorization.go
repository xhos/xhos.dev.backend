package middleware

import (
	"crypto/hmac"
	"errors"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/xhos/xhos.dev.backend/api"
	"github.com/xhos/xhos.dev.backend/internal/tools"
)

var UnAuthorizedError = errors.New("Invalid token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// Verify the API key (token)
		// Format expected: "Bearer YOUR_API_KEY"
		const prefix = "Bearer "
		if !strings.HasPrefix(token, prefix) {
			log.Error("Invalid token format")
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		apiKey := strings.TrimPrefix(token, prefix)

		// Get expected API key from environment or config
		expectedKey := tools.GetAPIKey() // You'll need to implement this function

		// Compare using a timing-safe comparison to prevent timing attacks
		if !hmac.Equal([]byte(apiKey), []byte(expectedKey)) {
			log.Error("Invalid API key")
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// Call the next handler if authentication succeeds
		next.ServeHTTP(w, r)
	})
}
