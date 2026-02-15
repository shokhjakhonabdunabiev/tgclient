package tgclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	token := "test_token"
	client := NewClient(token)

	if client == nil {
		t.Fatal("NewClient returned nil")
	}

	if client.token != token {
		t.Errorf("Expected token %s, got %s", token, client.token)
	}

	expectedBaseURL := BaseURL + token
	if client.baseURL != expectedBaseURL {
		t.Errorf("Expected baseURL %s, got %s", expectedBaseURL, client.baseURL)
	}

	if client.httpClient == nil {
		t.Error("httpClient should not be nil")
	}
}

func TestSetTimeout(t *testing.T) {
	client := NewClient("test_token")
	timeout := 10 * time.Second
	client.SetTimeout(timeout)

	if client.httpClient.Timeout != timeout {
		t.Errorf("Expected timeout %v, got %v", timeout, client.httpClient.Timeout)
	}
}

func TestGetMe(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/bottest_token/getMe" {
			t.Errorf("Expected path /bottest_token/getMe, got %s", r.URL.Path)
		}

		response := APIResponse{
			Ok: true,
			Result: json.RawMessage(`{
				"id": 123456789,
				"is_bot": true,
				"first_name": "TestBot",
				"username": "test_bot"
			}`),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient("test_token")
	client.baseURL = server.URL + "/bottest_token"

	user, err := client.GetMe()
	if err != nil {
		t.Fatalf("GetMe failed: %v", err)
	}

	if user.ID != 123456789 {
		t.Errorf("Expected user ID 123456789, got %d", user.ID)
	}

	if !user.IsBot {
		t.Error("Expected IsBot to be true")
	}

	if user.FirstName != "TestBot" {
		t.Errorf("Expected first name TestBot, got %s", user.FirstName)
	}

	if user.Username != "test_bot" {
		t.Errorf("Expected username test_bot, got %s", user.Username)
	}
}

func TestSendMessage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/bottest_token/sendMessage" {
			t.Errorf("Expected path /bottest_token/sendMessage, got %s", r.URL.Path)
		}

		response := APIResponse{
			Ok: true,
			Result: json.RawMessage(`{
				"message_id": 123,
				"chat": {
					"id": 987654321,
					"type": "private"
				},
				"date": 1234567890,
				"text": "Hello, World!"
			}`),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test_token")
	client.baseURL = server.URL + "/bottest_token"

	params := SendMessageParams{
		ChatID: 987654321,
		Text:   "Hello, World!",
	}

	message, err := client.SendMessage(params)
	if err != nil {
		t.Fatalf("SendMessage failed: %v", err)
	}

	if message.MessageID != 123 {
		t.Errorf("Expected message ID 123, got %d", message.MessageID)
	}

	if message.Chat.ID != 987654321 {
		t.Errorf("Expected chat ID 987654321, got %d", message.Chat.ID)
	}

	if message.Text != "Hello, World!" {
		t.Errorf("Expected text 'Hello, World!', got %s", message.Text)
	}
}

func TestGetUpdates(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/bottest_token/getUpdates" {
			t.Errorf("Expected path /bottest_token/getUpdates, got %s", r.URL.Path)
		}

		response := APIResponse{
			Ok: true,
			Result: json.RawMessage(`[
				{
					"update_id": 12345,
					"message": {
						"message_id": 100,
						"chat": {
							"id": 111222333,
							"type": "private"
						},
						"date": 1234567890,
						"text": "Test message"
					}
				}
			]`),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test_token")
	client.baseURL = server.URL + "/bottest_token"

	params := GetUpdatesParams{
		Offset: 0,
		Limit:  100,
	}

	updates, err := client.GetUpdates(params)
	if err != nil {
		t.Fatalf("GetUpdates failed: %v", err)
	}

	if len(updates) != 1 {
		t.Errorf("Expected 1 update, got %d", len(updates))
	}

	if updates[0].UpdateID != 12345 {
		t.Errorf("Expected update ID 12345, got %d", updates[0].UpdateID)
	}

	if updates[0].Message == nil {
		t.Fatal("Expected message to not be nil")
	}

	if updates[0].Message.Text != "Test message" {
		t.Errorf("Expected text 'Test message', got %s", updates[0].Message.Text)
	}
}

func TestAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := APIResponse{
			Ok:          false,
			Description: "Bad Request: chat not found",
			ErrorCode:   400,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test_token")
	client.baseURL = server.URL + "/bottest_token"

	_, err := client.GetMe()
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	expectedError := "telegram API error: Bad Request: chat not found (code: 400)"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}
