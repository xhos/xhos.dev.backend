package middleware

import (
	"crypto/subtle"
	"errors"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/xhos/xhos.dev.backend/internal/tools"
)

var UnAuthorizedError = errors.New("invalid token")

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			log.Error(UnAuthorizedError)
			http.Error(w, UnAuthorizedError.Error(), http.StatusUnauthorized)
			return
		}

		// Format expected: "Bearer YOUR_API_KEY"
		const prefix = "Bearer "
		if !strings.HasPrefix(token, prefix) {
			log.Error("Invalid token format")
			http.Error(w, UnAuthorizedError.Error(), http.StatusUnauthorized)
			return
		}

		apiKey := strings.TrimPrefix(token, prefix)

		// Get expected API key from environment or config
		expectedKey := tools.GetAPIKey()

		// Compare using a timing-safe comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(apiKey), []byte(expectedKey)) != 1 {
			log.Error("Invalid API key")
			http.Error(w, UnAuthorizedError.Error(), http.StatusUnauthorized)
			return
		}

		// Call the next handler if authentication succeeds
		next.ServeHTTP(w, r)
	})
}
