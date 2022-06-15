package betterttv

import (
	"encoding/json"
	"net/http"
)

const apiBase = "https://api.betterttv.net/3"

type id string

type SearchResult struct {
	ID   id     `json:"id"`
	Code string `json:"code"`
	Type string `json:"imageType"`
}

var DefaultClient = NewClient()

type Client struct {
	HttpClient *http.Client
	URL        string
}

// NewClient returns a new Client with the default http.Client.
func NewClient() *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		URL:        apiBase,
	}
}

// Search returns a list of search results for the given query.
func (c *Client) Search(query string) ([]SearchResult, error) {
	req, err := http.NewRequest("GET", c.URL+"/emotes/shared/search?offset=0&limit=50&query="+query, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var results []SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	return results, nil
}

// Url returns the image url for the given id.
func (id id) Url() string {
	return "https://cdn.betterttv.net/emote/" + string(id) + "/3x"
}
