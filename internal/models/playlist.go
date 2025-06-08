package models

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type Owner struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name,omitempty"`
	ProfileURL  string `json:"profile_url,omitempty"`
}

type Playlist struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Provider      string  `json:"provider"` // "spotify", "soundcloud", etc.
	ExternalURL   string  `json:"external_url,omitempty"`
	Images        []Image `json:"images,omitempty"`
	TracksCount   int     `json:"tracks_count"`
	Public        bool    `json:"public"`
	Collaborative bool    `json:"collaborative"`
	Owner         Owner   `json:"owner"`
	URI           string  `json:"uri,omitempty"`
	Href          string  `json:"href,omitempty"`
}
