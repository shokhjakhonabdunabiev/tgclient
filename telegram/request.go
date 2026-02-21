package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) get(method string, params any) (*Response, error) {
	u, err := url.Parse(c.baseURL + "/" + method)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	q, err := structToQuery(params)
	if err != nil {
		return nil, fmt.Errorf("failed to parse params: %w", err)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiRes Response
	if err := json.Unmarshal(body, &apiRes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !apiRes.OK {
		return nil, fmt.Errorf("telegram BOT API error: %s (code %d)", apiRes.Description, apiRes.ErrorCode)
	}

	return &apiRes, nil

}

func (c *Client) post(method string, payload any) (*Response, error) {
	u, err := url.Parse(c.baseURL + "/" + method)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiRes Response
	if err := json.Unmarshal(body, &apiRes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !apiRes.OK {
		return nil, fmt.Errorf("telegram BOT API error: %s (code %d)", apiRes.Description, apiRes.ErrorCode)
	}

	return &apiRes, nil
}

func structToQuery(params any) (url.Values, error) {
	values := url.Values{}

	if params == nil {
		return values, nil
	}

	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	for k, v := range m {
		switch val := v.(type) {
		case string:
			values.Set(k, val)
		case float64:
			values.Set(k, fmt.Sprintf("%.0f", val))
		case bool:
			values.Set(k, fmt.Sprintf("%t", val))
		default:
			nested, err := json.Marshal(val)
			if err != nil {
				return nil, err
			}
			values.Set(k, string(nested))
		}
	}

	return values, nil
}
