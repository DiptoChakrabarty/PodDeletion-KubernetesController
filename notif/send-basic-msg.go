package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	//"github.com/slack-go/slack"
)

func getConnection() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured: %s", err)
	}
	fmt.Println(os.Getenv("OAUTH_TOKEN"))
	//api := slack.New(os.Getenv("OAUTH_TOKEN"))
	// channelID, timestamp, err := api.PostMessage(

	//)
}

func main() {
	getConnection()
}
