package patari

// SearchResponse represents the structure of response for the Search API Result
// Ignoring Album and playlists for now
type SearchResponse struct {
	Data struct {
		Song     []*Song     `json:"song"`
		Album    []*Album    `json:"album"`
		Artist   []*Artist   `json:"artist"`
		Playlist []*Playlist `json:"playlist"`
	} `json:"data"`
	Success bool `json:"success"`
}

// PlaylistDetail struct represents the response for GetPlaylist request
type PlaylistDetail struct {
	Item
	Songs []Song `json:"songs"`
}
