package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DiptoChakrabarty/podDeletionController/logger"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

type SlackModel struct {
	Connection *slack.Client
}

func loadEnv(key, defaultValue string) string {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Some error occured: %s", err)
		logger.Info("Defaulting to default values")
		return defaultValue
	}
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func getSlackClient() *slack.Client {
	oauthToken := loadEnv("OAUTH_TOKEN", "somth")
	return slack.New(oauthToken)

}

func NewSlackClient() *SlackModel {
	client := getSlackClient()
	return &SlackModel{
		Connection: client,
	}
}

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
