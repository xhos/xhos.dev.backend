package api

import (
	"encoding/json"
	"net/http"
)

type UserInfo struct {
	Username string
}

type userInfoResponse struct {
	Code    int
	Name    string
	IconURL string
}

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, messege string, code int) {
	resp := Error{
		Code:    code,
		Message: messege,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, "An unexpected error occured", http.StatusInternalServerError)
	}
)
