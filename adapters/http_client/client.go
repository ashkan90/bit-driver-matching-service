package http_client

import (
	"io"
	"net/http"
	"time"
)

type Client struct {
	Timeout time.Duration
	client  *http.Client
}

type ClientImplementation interface {
	Get(url string) (*http.Response, error)
	Do(method, url string, body io.Reader) (*http.Response, error)
}

const (
	DefaultTimeout = time.Second * 3
)

func NewClient() *Client {
	return &Client{
		client: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

func (c *Client) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}

func (c *Client) Do(method, url string, body io.Reader) (*http.Response, error) {
	var req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.client.Do(req)
}
