package patari

const (
	// CommandSearch performs search for songs, albums etc.
	CommandSearch = "search"

	// Categories
	CategorySong     = "song"
	CategoryArtist   = "artist"
	CategoryAlbum    = "album"
	CategoryPlaylist = "playlist"
)

// RequestData ...
type RequestData struct {
	Command        string   `json:"command"`
	Message        string   `json:"message"`
	Categories     []string `json:"categories"`
	CategoryLimits []int    `json:"category_limits"`
	Callback       int      `json:"callback"`
}
