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
		return nil, fmt.Errorf("failed to unmarshal User: %w", err)
	}

	return &user, nil
}

type GetChatRequest struct {
	ChatID string `json:"chat_id"`
}

func (c *Client) GetChat(params GetChatRequest) (*ChatFullInfo, error) {
	res, err := c.get("getChat", params)
	if err != nil {
		return nil, err
	}

	var chatInfo ChatFullInfo
	if err := json.Unmarshal(res.Result, &chatInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ChatFullInfo: %w", err)
	}

	return &chatInfo, nil
}

type ParseMode string

const (
	HTML       ParseMode = "HTML"
	MarkdownV2 ParseMode = "MarkdownV2"
)

type SendMessageRequest struct {
	ChatID    string    `json:"chat_id"`
	Text      string    `json:"text"`
	ParseMode ParseMode `json:"parse_mode"`
}

func (c *Client) SendMessage(body SendMessageRequest) (*Message, error) {
	res, err := c.post("sendMessage", body)
	if err != nil {
		return nil, err
	}

	var msg Message
	if err := json.Unmarshal(res.Result, &msg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Message: %w", err)
	}

	return &msg, nil
}
