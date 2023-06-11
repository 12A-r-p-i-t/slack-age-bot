package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load()
	Auth_Token := os.Getenv("AUTH_TOKEN")
	if err != nil {
		log.Fatal("Error loading the env file")
	}
	api := slack.New(os.Getenv("AUTH_TOKEN"))
	api.PostMessage()
}
