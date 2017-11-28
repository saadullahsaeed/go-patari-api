package patari

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

const (
	defaultAPIHost = "wss://search-0-0.patari.pk/realtime"

	// CDNHost to serve the static assets
	CDNHost = "https://patarimedia.blob.core.windows.net/patari/"
)

// Client ...
type Client struct {
	Host       string
	Connection *websocket.Conn
}

// Connect establishes a websocket connection with the Host
func (c *Client) Connect() error {
	if c.Host == "" {
		c.Host = defaultAPIHost
	}

	var err error
	c.Connection, _, err = websocket.DefaultDialer.Dial(c.Host, nil)
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
		Categories:     []string{CategorySong},
		CategoryLimits: []int{100},
		Message:        term,
	}

	jrm, err := c.RequestAPI(r)
	if err != nil {
		return nil, err
	}

	res := &SearchResponse{}
	err = json.Unmarshal(jrm, res)
	return res, err
}

// RequestAPI sends a websocket command to the Patari API
func (c *Client) RequestAPI(r *RequestData) (json.RawMessage, error) {
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

// NewClient returns an instance of Patari API client
func NewClient() *Client {
	return &Client{}
}
