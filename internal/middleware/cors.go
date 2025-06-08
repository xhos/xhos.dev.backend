package middleware

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/xhos/xhos.dev.backend/internal/tools"
)

func CORS(next http.Handler) http.Handler {
	websiteURL := tools.GetWebsiteURL()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{websiteURL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            false,
	})

	return c.Handler(next)
}
