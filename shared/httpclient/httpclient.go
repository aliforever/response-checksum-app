package httpclient

import (
	"io"
	"net/http"
	"strings"
)

// HttpClient It's an interface to be passed to the Client making it easier to write mocks in tests and also \
// pass custom Clients
type HttpClient interface {
	Get(address string) (*http.Response, error)
}

// Client is a wrapper for http methods. Currently supporting only Get
type Client struct {
	httpClient HttpClient
}

// New Initializes a new Client
func New(c HttpClient) *Client {
	if c == nil {
		c = http.DefaultClient
	}
	return &Client{httpClient: c}
}

// ParseUrl adds http scheme to urls not containing http or https schemes
func (Client) ParseUrl(address string) string {
	address = strings.TrimSpace(address)

	if !strings.HasPrefix(address, "http://") && !strings.HasPrefix(address, "https://") {
		return "http://" + address
	}

	return address
}

// Get Fires a get request to the provided address and returns the response body or error
func (c Client) Get(address string) ([]byte, error) {
	resp, err := c.httpClient.Get(address)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
