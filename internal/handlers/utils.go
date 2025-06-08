package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

// Helper functions
func getProvidersFromQuery(r *http.Request) []string {
	providersParam := r.URL.Query().Get("providers")
	providers := strings.Split(providersParam, ",")
	if len(providers) == 1 && providers[0] == "" {
		return []string{"spotify"} // Default provider
	}
	return providers
}

func getLimitFromQuery(r *http.Request, defaultLimit int) int {
	limit := defaultLimit
	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		if parsedLimit, err := strconv.Atoi(limitParam); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}
	return limit
}

func sendResponse(w http.ResponseWriter, data interface{}, errors map[string]string) {
	response := map[string]interface{}{
		"data":   data,
		"errors": errors,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Error("Failed to encode response:", err)
	}
}

func sendErrorResponse(w http.ResponseWriter, err error, status int) {
	logrus.Error(err)
	response := map[string]interface{}{
		"error": err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
