package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DiptoChakrabarty/podDeletionController/logger"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func getConnection() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured: %s", err)
	}
	api := slack.New(os.Getenv("OAUTH_TOKEN"))
	channelID, timestamp, err := api.PostMessage(
		os.Getenv("CHANNEL_ID"),
		slack.MsgOptionText("Hello World", false),
	)

	if err != nil {
		logger.Error("Some error occured: %s", err)
	}
	fmt.Println("Message sent successfully at %s to channel ID %s", timestamp, channelID)
}

func main() {
	getConnection()
}
