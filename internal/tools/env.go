package tools

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
)

type envCache struct {
	value string
	once  sync.Once
}

var envVars = map[string]*envCache{
	"API_KEY":        &envCache{},
	"SPOTIFY_ID":     &envCache{},
	"SPOTIFY_SECRET": &envCache{},
}

func getEnvVar(name string) string {
	cache := envVars[name]
	cache.once.Do(func() {
		cache.value = os.Getenv(name)
		if cache.value == "" {
			log.Warnf("%s environment variable is not set", name)
		}
	})
	return cache.value
}

func GetAPIKey() string {
	return getEnvVar("API_KEY")
}

func GetSpotifyID() string {
	return getEnvVar("SPOTIFY_ID")
}

func GetSpotifySecret() string {
	return getEnvVar("SPOTIFY_SECRET")
}
