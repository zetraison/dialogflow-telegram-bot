package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	ruLanguage   = "ru"
	errorMessage = "Повторите, пожалуйста, вопрос."
)

func DetectIntentText(ctx context.Context, projectID, sessionID, text, languageCode string) (string, error) {
	sessionClient, err := dialogflow.NewSessionsClient(ctx)
	if err != nil {
		return "", err
	}
	defer sessionClient.Close()

	if projectID == "" || sessionID == "" {
		return "", errors.New(fmt.Sprintf("Received empty project (%s) or session (%s)", projectID, sessionID))
	}

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", projectID, sessionID)
	textInput := dialogflowpb.TextInput{Text: text, LanguageCode: languageCode}
	queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
	queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
	request := dialogflowpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return "", err
	}

	queryResult := response.GetQueryResult()
	fulfillmentText := queryResult.GetFulfillmentText()
	return fulfillmentText, nil
}

func main() {
	tgApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	if len(tgApiToken) == 0 {
		panic("Telegram API token not set!")
	}

	projectID := os.Getenv("PROJECT_ID")
	if len(projectID) == 0 {
		panic("Project ID not set!")
	}

	sessionID := os.Getenv("SESSION_ID")
	if len(sessionID) == 0 {
		panic("Session ID not set!")
	}

	bot, err := tgbotapi.NewBotAPI(tgApiToken)
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	ctx, _ := context.WithCancel(context.Background())

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		text, err := DetectIntentText(ctx, projectID, sessionID, update.Message.Text, ruLanguage)
		if err != nil {
			text = errorMessage
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
