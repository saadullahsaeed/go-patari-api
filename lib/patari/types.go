package patari

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
	ID            string `json:"songID"`
	ShortName     string `json:"songShortName"`
	Name          string `json:"songName"`
	Stream        string `json:"stream"`
	AlbumID       string `json:"albumID"`
	AlbumName     string `json:"albumName"`
	AlbumThumb    string `json:"albumThumb"`
	AlbumThumbBig string `json:"albumThumbBig"`
	ArtistID      string `json:"artistID"`
	ArtistThumb   string `json:"artistThumb"`
	ArtistName    string `json:"artistName"`
	Plays         int    `json:"songPlays"`
	Slug          string `json:"songSlug"`
	Genre         string `json:"genre"`
}
