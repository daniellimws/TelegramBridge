package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/yanzay/tbot/v2"
)

func linkgen(repo, event string) string {
	context := map[string]string{
		"issue_comment":               "issues",
		"issues":                      "issues",
		"pull_request":                "pulls",
		"pull_request_review_comment": "pulls",
		"push":                        "commits",
		"project_card":                "projects",
	}

	event = context[strings.ToLower(event)]

	// generates link based on the triggered event
	return fmt.Sprintf("https://github.com/%s/%s/", repo, event)
}

func main() {

	var (
		// inputs provided by Github Actions runtime
		// should be defined in the action.yml
		token = os.Getenv("INPUT_TOKEN")
		chat  = os.Getenv("INPUT_CHAT")

		// github environment context
		workflow = os.Getenv("GITHUB_WORKFLOW")
		repo     = os.Getenv("GITHUB_REPOSITORY")
		// commit   = os.Getenv("GITHUB_SHA")
	)

	// Create Telegram client using token
	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")

	// link to the commit
	link := linkgen(repo, event)

	// Prepare message to send
	msg := os.Getenv("INPUT_MESSAGE")

	// Send to chat using Markdown format
	_, err := c.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
