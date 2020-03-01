package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yanzay/tbot/v2"
)

func main() {

	var (
		// inputs provided by Github Actions runtime
		// should be defined in the action.yml
		token = os.Getenv("INPUT_TOKEN")
		chat  = os.Getenv("INPUT_CHAT")
	)

	// Create Telegram client using token
	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")

	// Prepare message to send
	msg := os.Getenv("INPUT_MESSAGE")

	// Send to chat
	_, err := c.SendMessage(chat, msg)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
