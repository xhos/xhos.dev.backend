package api

import (
	"encoding/json"
	"net/http"

type UserInfo struct {
	Username string
}

type userInfoResponse struct {
	Code int
	Name string
	IconURL string 
}

type Error struct {
	Code int
	Message string
}