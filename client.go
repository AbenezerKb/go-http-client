package client

import "net/http"

type Client struct {
	apiURL     string
	httpClient *http.Client
}

var (
	DefaultURL = "https://pokeapi.co"
)

type Option func(*Client)

func NewClient(opts ...Option) *Client {
	client := &Client{
		apiURL:     DefaultURL,
		httpClient: http.DefaultClient,
	}

	for _, o := range opts {
		o(client)
	}

	return client
}

func WithAPIURL(url string) Option {
	return func(c *Client) {
		c.apiURL = url
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
