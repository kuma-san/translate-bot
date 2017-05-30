package main

import (
	"log"
	"os"

	"cloud.google.com/go/translate"

	"github.com/nlopes/slack"
	"golang.org/x/net/context"
	"golang.org/x/text/language"
)

func googletranslate(text string) (string, error) {
	ctx := context.Background()

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}

	detected, err := client.DetectLanguage(ctx, []string{text})
	if err != nil {
		return "", err
	}
	log.Printf("Detection: %+v\n", detected)
	target, err := language.Parse("ja")

	if detected[0][0].Language == language.Japanese {
		target, err = language.Parse("en")
	}
	if err != nil {
		return "", err
	}

	translation, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		return "", err
	}

	return translation[0].Text, nil
}

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				translated, err := googletranslate(ev.Text)
				if err != nil {
					log.Fatalf("Failed to translate text: %v", err)
				}
				rtm.SendMessage(rtm.NewOutgoingMessage(translated, ev.Channel))

			case *slack.InvalidAuthEvent:
				log.Print("Invalid Credentials")
				return 1
			}
		}
	}
}

func main() {
	slackKey := os.Getenv("SLACKAPIKEY")
	log.Print(slackKey)
	api := slack.New(slackKey)
	os.Exit(run(api))
}
