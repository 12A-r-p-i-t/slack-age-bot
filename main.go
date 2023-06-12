package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	err := godotenv.Load()
	BOT_TOKEN := os.Getenv("BOT_TOKEN")
	APP_TOKEN := os.Getenv("SLACK_APP_TOKEN")
	if err != nil {
		log.Fatal("Error loading the env file")
	}

	bot := slacker.NewClient(BOT_TOKEN, APP_TOKEN)
	bot.Command("my yob is {year}", &slacker.CommandDefinition{
		Description: "Calculate the age",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			word := request.Param("year")
			yob, err := strconv.Atoi(word)
			if err != nil {
				log.Fatal("Error in converting to int from string")
			}
			curr_year := time.Now().Year()
			age := curr_year - yob
			curr_age := strconv.Itoa(age)
			response.Reply("Your age is " + curr_age)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
