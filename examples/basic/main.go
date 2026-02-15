package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shokhjakhonabdunabiev/tgclient"
)

func main() {
	// Get bot token from environment variable
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	// Create a new client
	client := tgclient.NewClient(token)

	// Get bot information
	me, err := client.GetMe()
	if err != nil {
		log.Fatalf("Failed to get bot info: %v", err)
	}

	fmt.Printf("Bot Info:\n")
	fmt.Printf("  ID: %d\n", me.ID)
	fmt.Printf("  Username: @%s\n", me.Username)
	fmt.Printf("  First Name: %s\n", me.FirstName)
	fmt.Printf("  Is Bot: %t\n", me.IsBot)

	// Example: Send a message (uncomment and set chat ID to use)
	// chatID := int64(123456789) // Replace with actual chat ID
	// msg, err := client.SendMessage(tgclient.SendMessageParams{
	// 	ChatID: chatID,
	// 	Text:   "Hello from tgclient!",
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to send message: %v", err)
	// }
	// fmt.Printf("Message sent with ID: %d\n", msg.MessageID)

	// Example: Get updates (long polling)
	fmt.Println("\nListening for updates... (Press Ctrl+C to stop)")
	offset := int64(0)
	for {
		updates, err := client.GetUpdates(tgclient.GetUpdatesParams{
			Offset:  offset,
			Limit:   100,
			Timeout: 30,
		})
		if err != nil {
			log.Printf("Failed to get updates: %v", err)
			continue
		}

		for _, update := range updates {
			if update.Message != nil {
				fmt.Printf("\nReceived message:\n")
				fmt.Printf("  From: @%s (%s)\n", update.Message.From.Username, update.Message.From.FirstName)
				fmt.Printf("  Text: %s\n", update.Message.Text)
				fmt.Printf("  Chat ID: %d\n", update.Message.Chat.ID)
			}
			offset = update.UpdateID + 1
		}
	}
}
