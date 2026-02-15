package tgclient

import (
	"encoding/json"
	"fmt"
)

// GetMe returns basic information about the bot
func (c *Client) GetMe() (*User, error) {
	data, err := c.makeRequest("getMe", nil)
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return &user, nil
}

// SendMessage sends a text message to a chat
func (c *Client) SendMessage(params SendMessageParams) (*Message, error) {
	data, err := c.makeRequest("sendMessage", params)
	if err != nil {
		return nil, err
	}

	var message Message
	if err := json.Unmarshal(data, &message); err != nil {
		return nil, fmt.Errorf("failed to unmarshal message: %w", err)
	}

	return &message, nil
}

// GetUpdates receives incoming updates using long polling
func (c *Client) GetUpdates(params GetUpdatesParams) ([]Update, error) {
	data, err := c.makeRequest("getUpdates", params)
	if err != nil {
		return nil, err
	}

	var updates []Update
	if err := json.Unmarshal(data, &updates); err != nil {
		return nil, fmt.Errorf("failed to unmarshal updates: %w", err)
	}

	return updates, nil
}
