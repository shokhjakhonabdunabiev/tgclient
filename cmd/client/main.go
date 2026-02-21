package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shokhjakhonabdunabiev/tgclient/telegram"
)

func main() {
	client := telegram.NewClient("", 10*time.Second)

	user, err := client.GetMe()
	if err != nil {
		fmt.Println(err)
		return
	}
	print(user)

	chat, err := client.GetChat(telegram.GetChatRequest{ChatID: "@move_it_move"})
	if err != nil {
		fmt.Println(err)
		return
	}
	print(chat)

}

func print[T any](data T) {
	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(b))
}
