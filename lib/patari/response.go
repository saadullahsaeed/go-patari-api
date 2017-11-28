package patari

// SearchResponse represents the structure of response for the Search API Result
// Ignoring Album and playlists for now
type SearchResponse struct {
	Data struct {
		Song []*Song `json:"song"`
	} `json:"data"`
	Success bool `json:"success"`
}
