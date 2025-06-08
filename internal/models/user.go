package models

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Platform   string `json:"platform"`
	ProfileURL string `json:"profile_url,omitempty"`
	ImageURL   string `json:"image_url,omitempty"`
}
