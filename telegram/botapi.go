package telegram

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetMe() (*User, error) {
	res, err := c.get("getMe", nil)
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal(res.Result, &user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return &user, nil
}

type GetChatRequest struct {
	ChatID string `json:"chat_id"`
}

func (c *Client) GetChat(req GetChatRequest) (*ChatFullInfo, error) {
	data, err := c.get("getChat", req)
	if err != nil {
		return nil, err
	}

	var chatInfo ChatFullInfo
	if err := json.Unmarshal(data.Result, &chatInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return &chatInfo, nil
}
