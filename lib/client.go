package patari

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	defaultWSEndpoint = "wss://search-0-0.patari.pk/realtime"
	defaultAPIHost    = "http://api.patari.pk/site"

	// CDNHost to serve the static assets
	CDNHost = "https://patarimedia.blob.core.windows.net/patari/"
)

// Client ...
type Client struct {
	WebsocketEndpoint string
	APIHost           string
	Connection        *websocket.Conn
}

// Connect establishes a websocket connection with the Host
func (c *Client) Connect() error {
	if c.WebsocketEndpoint == "" {
		c.WebsocketEndpoint = defaultWSEndpoint
	}

	var err error
	c.Connection, _, err = websocket.DefaultDialer.Dial(c.WebsocketEndpoint, nil)
	if err != nil {
		return err
	}
	return nil
}

// Close will disconenct from the websocket
func (c *Client) Close() error {
	return c.Connection.Close()
}

// Search takes a term and returns the result, only supports songs for now
func (c *Client) Search(term string) (*SearchResponse, error) {
	r := &RequestData{
		Command:        CommandSearch,
		Categories:     []string{CategorySong, CategoryArtist, CategoryAlbum, CategoryPlaylist},
		CategoryLimits: []int{100, 100, 100, 100},
		Message:        term,
	}

	jrm, err := c.RequestWebsocket(r)
	if err != nil {
		return nil, err
	}

	res := &SearchResponse{}
	err = json.Unmarshal(jrm, res)
	return res, err
}

// GetPlaylist can accept a Playlist ID or a discovery slug
func (c *Client) GetPlaylist(pid string) (*PlaylistDetail, error) {
	jstr, err := c.RequestHTTPAPI("GET", fmt.Sprintf("/getPlaylist/%s", pid), nil)
	if err != nil {
		return nil, err
	}

	pd := &PlaylistDetail{}
	err = json.Unmarshal(jstr, pd)
	return pd, err
}

// RequestWebsocket sends a websocket command to the Patari API
func (c *Client) RequestWebsocket(r *RequestData) (json.RawMessage, error) {
	jdata, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	err = c.Connection.WriteMessage(websocket.TextMessage, jdata)
	if err != nil {
		return nil, err
	}

	_, message, err := c.Connection.ReadMessage()
	if err != nil {
		return nil, err
	}
	return message, nil
}

// RequestHTTPAPI sends a websocket command to the Patari API
func (c *Client) RequestHTTPAPI(method, url string, data interface{}) (json.RawMessage, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", defaultAPIHost, url), nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

// NewClient returns an instance of Patari API client
func NewClient() *Client {
	return &Client{}
}
