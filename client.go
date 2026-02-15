package tgclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// BaseURL is the base URL for the Telegram Bot API
	BaseURL = "https://api.telegram.org/bot"
)

// Client represents a Telegram Bot API client
type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Telegram Bot API client
func NewClient(token string) *Client {
	return &Client{
		token:   token,
		baseURL: BaseURL + token,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SetTimeout sets the HTTP client timeout
func (c *Client) SetTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}

// makeRequest makes an HTTP request to the Telegram API
func (c *Client) makeRequest(method string, params interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, method)

	var reqBody io.Reader
	if params != nil {
		jsonData, err := json.Marshal(params)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal params: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !apiResp.Ok {
		return nil, fmt.Errorf("telegram API error: %s (code: %d)", apiResp.Description, apiResp.ErrorCode)
	}

	return apiResp.Result, nil
}
