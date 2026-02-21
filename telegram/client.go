package telegram

import (
	"net/http"
	"time"
)

const (
	BaseURL = "https://api.telegram.org/bot"
)

type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

func NewClient(token string, timeout time.Duration) *Client {
	return &Client{
		token:   token,
		baseURL: BaseURL + token,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}
