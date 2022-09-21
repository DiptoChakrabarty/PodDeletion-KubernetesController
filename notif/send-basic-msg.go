package notif

import (
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

func (client *SlackModel) SendMessage(s string) (string, string, error) {
	channelID, timestamp, err := client.Connection.PostMessage(
		loadEnv("CHANNEL_ID", "not provided"),
		slack.MsgOptionText(s, false),
	)
	return channelID, timestamp, err
}
