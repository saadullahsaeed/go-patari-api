package patari

const (
	TypePlaylist = "playlist"
)

// Item is the generic type for most detailed types
type Item struct {
	ID           string `json:"_id"`
	ItemIndex    int    `json:"itemIndex"`
	ThumbBig     string `json:"thumbBig"`
	Thumb        string `json:"thumb"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	Name         string `json:"name"`
	DiscoverSlug string `json:"discoverSlug"`
}

// Playlist ...
type Playlist struct {
	ID    string `json:"playlistID"`
	Name  string `json:"playlistName"`
	Tag   string `json:"playlistTag"`
	Thumb string `json:"playlistThumb"`
}

// Artist ...
type Artist struct {
	ID string `json:"artistID"`
}

// Album ...
type Album struct {
	ID         string `json:"albumID"`
	AlbumName  string `json:"albumName"`
	AlbumThumb string `json:"albumThumb"`
	ArtistID   string `json:"artistID"`
	ArtistName string `json:"artistName"`
}

// Song ...
type Song struct {
	ID             string `json:"songID"`
	ShortName      string `json:"songShortName"`
	Name           string `json:"songName"`
	StreamingReady bool   `json:"streamingReady"`
	Stream         string `json:"stream"`
	SecureStream   string `json:"secureStream"`
	AlbumID        string `json:"albumID"`
	AlbumName      string `json:"albumName"`
	AlbumThumb     string `json:"albumThumb"`
	AlbumThumbBig  string `json:"albumThumbBig"`
	ArtistID       string `json:"artistID"`
	ArtistThumb    string `json:"artistThumb"`
	ArtistName     string `json:"artistName"`
	Plays          int    `json:"songPlays"`
	Slug           string `json:"songSlug"`
	Genre          string `json:"genre"`
}
