package patari

const (
	// CommandSearch performs search for songs, albums etc.
	CommandSearch = "search"

	// CategorySong ...
	CategorySong = "song"

	// CategoryArtist ...
	CategoryArtist = "artist"

	// CategoryAlbum ...
	CategoryAlbum = "album"

	// CategoryPlaylist ...
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
