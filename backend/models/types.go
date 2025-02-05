package models

type UserProfile struct {
	DisplayName string  `json:"display_name"`
	Followers   int     `json:"followers"`
	Images      []Image `json:"images"`
}

type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type TopTracks struct {
	Items []Track `json:"items"`
}

type Track struct {
	Album struct {
		Images []Image `json:"images"`
	} `json:"album"`
	Name    string `json:"name"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
	Duration     int `json:"duration_ms"`
	ExternalURLs struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
}

type Artist struct {
	Name         string `json:"name"`
	ExternalURLs struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Genres []string `json:"genres"`
	Images []Image  `json:"images"`
}

type TopArtists struct {
	Items []Artist `json:"items"`
}
