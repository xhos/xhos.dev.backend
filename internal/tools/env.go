package tools

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
)

var (
	apiKey     string
	apiKeyOnce sync.Once
)

// GetAPIKey returns the API key from environment variables or configuration.
// It uses a singleton pattern to only load the key once.
func GetAPIKey() string {
	apiKeyOnce.Do(func() {
		apiKey = os.Getenv("API_KEY")

		if apiKey == "" {
			log.Warn("API_KEY environment variable is not set")
		}
	})

	return apiKey
}
