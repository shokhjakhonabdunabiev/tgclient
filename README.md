# tgclient

A simple and lightweight Telegram Bot API client library for Go.

## Features

- 🚀 Simple and easy-to-use API
- 📦 Minimal dependencies (only standard library)
- ✅ Type-safe methods and responses
- 🔄 Support for long polling (getUpdates)
- 💬 Essential bot operations (getMe, sendMessage, etc.)
- 🧪 Well-tested codebase

## Installation

```bash
go get github.com/shokhjakhonabdunabiev/tgclient
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/shokhjakhonabdunabiev/tgclient"
)

func main() {
    // Create a new client
    client := tgclient.NewClient("YOUR_BOT_TOKEN")
    
    // Get bot information
    me, err := client.GetMe()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Bot username: @%s\n", me.Username)
    
    // Send a message
    msg, err := client.SendMessage(tgclient.SendMessageParams{
        ChatID: 123456789,
        Text:   "Hello from tgclient!",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Message sent with ID: %d\n", msg.MessageID)
}
```

## Usage

### Creating a Client

```go
client := tgclient.NewClient("YOUR_BOT_TOKEN")

// Optionally set custom timeout
client.SetTimeout(60 * time.Second)
```

### Getting Bot Information

```go
me, err := client.GetMe()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Bot ID: %d\n", me.ID)
fmt.Printf("Bot Username: @%s\n", me.Username)
```

### Sending Messages

```go
msg, err := client.SendMessage(tgclient.SendMessageParams{
    ChatID:    123456789,
    Text:      "Hello, World!",
    ParseMode: "Markdown", // Optional
})
if err != nil {
    log.Fatal(err)
}
```

### Receiving Updates (Long Polling)

```go
offset := int64(0)
for {
    updates, err := client.GetUpdates(tgclient.GetUpdatesParams{
        Offset:  offset,
        Limit:   100,
        Timeout: 30,
    })
    if err != nil {
        log.Printf("Error: %v", err)
        continue
    }
    
    for _, update := range updates {
        if update.Message != nil {
            fmt.Printf("Received: %s\n", update.Message.Text)
            // Process message...
        }
        offset = update.UpdateID + 1
    }
}
```

## Available Methods

- `GetMe()` - Get basic information about the bot
- `SendMessage(params SendMessageParams)` - Send text messages
- `GetUpdates(params GetUpdatesParams)` - Receive incoming updates via long polling

## Examples

Check the [examples](examples/) directory for complete working examples:

- [Basic Example](examples/basic/main.go) - Demonstrates GetMe and GetUpdates

To run an example:

```bash
cd examples/basic
export TELEGRAM_BOT_TOKEN="your_bot_token"
go run main.go
```

## Testing

Run the test suite:

```bash
go test -v
```

Run tests with coverage:

```bash
go test -v -cover
```

## API Reference

For detailed information about the Telegram Bot API, refer to the official documentation:
https://core.telegram.org/bots/api

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Roadmap

Future enhancements may include:
- Additional API methods (deleteMessage, editMessage, etc.)
- Support for file uploads and multimedia
- Webhook support
- Inline keyboard support
- More comprehensive type definitions
- Rate limiting and retry logic