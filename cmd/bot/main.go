package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zetraison/dialogflow-telegram-bot/pkg/google"
)

const (
	errorMessage = "Повторите, пожалуйста, вопрос."
)

func HandleMessageUpdate(ctx context.Context, projectID string, update tgbotapi.Update) *tgbotapi.MessageConfig {
	sessionID := SessionFromUserID(update.Message.From.ID)
	text, err := DetectIntentText(ctx, projectID, sessionID, update.Message.Text, RuLanguageCode)
	if err != nil {
		logError(err)
		text = errorMessage
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID
	return &msg
}

func main() {
	token := os.Getenv("TELEGRAM_API_TOKEN")
	if len(token) == 0 {
		panic("Telegram API token not set!")
	}

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT_ID")
	if len(projectID) == 0 {
		panic("Project ID not set!")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	ctx := context.Background()

	for update := range updates {
		if update.Message == nil {
			continue
		}
		logMessage(*update.Message)
		answer := HandleMessageUpdate(ctx, projectID, update)
		message, err := bot.Send(answer)
		if err != nil {
			logError(err)
			continue
		}
		logMessage(message)
	}
}
