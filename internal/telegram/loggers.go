package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func logMessage(message tgbotapi.Message) {
	log.Printf("[%d][%s] %s", message.From.ID, message.From.UserName, message.Text)
}

func logError(err error) {
	log.Printf("Error: %s", err.Error())
}
