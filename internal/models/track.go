package models

type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URI  string `json:"uri,omitempty"`
}

type Album struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Images []Image `json:"images,omitempty"`
	URI    string  `json:"uri,omitempty"`
}

type Track struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Artists     []Artist `json:"artists"`
	Album       Album    `json:"album,omitempty"`
	Duration    int      `json:"duration_ms"`
	Provider    string   `json:"provider"`
	ExternalURL string   `json:"external_url,omitempty"`
	URI         string   `json:"uri,omitempty"`
	IsPlayable  bool     `json:"is_playable"`
	Explicit    bool     `json:"explicit"`
	TrackNumber int      `json:"track_number,omitempty"`
	DiscNumber  int      `json:"disc_number,omitempty"`
	Popularity  int      `json:"popularity,omitempty"`
	PreviewURL  string   `json:"preview_url,omitempty"`
	IsLocal     bool     `json:"is_local"`
	AddedAt     string   `json:"added_at,omitempty"`
}
