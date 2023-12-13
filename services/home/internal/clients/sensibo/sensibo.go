package sensibo

import (
	"net/http"
	"net/url"
	"time"
)

var baseUrl = url.URL{
	Scheme: "https",
	Host:   "home.sensibo.com",
	Path:   "/api/v2/",
}

func New(apiKey string) *Client {
	c := &http.Client{Timeout: 10 * time.Second}
	return &Client{
		apiKey:     apiKey,
		httpClient: c,
	}
}

type Client struct {
	apiKey     string
	httpClient *http.Client
}

type Response[T any] struct {
	Status string
	Result []T
}
