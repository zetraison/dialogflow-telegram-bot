package tgbot

import (
	"context"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type tgBot struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *tgBot {
	return &tgBot{
		config: config,
		logger: logrus.New(),
	}
}

func (t *tgBot) Start() error {
	err := t.configureLogger()
	if err != nil {
		log.Fatal(err)
	}

	api, err := tgbotapi.NewBotAPI(t.config.TelegramApiToken)
	if err != nil {
		log.Fatal(err)
	}

	t.logger.Infof("Authorized on tg account \"%s\"", api.Self.UserName)

	updates, err := t.getUpdates(api)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		t.logger.Infof("[%d][%s] %s", update.Message.From.ID, update.Message.From.UserName, update.Message.Text)

		answer := t.handleUpdate(ctx, update)

		message, err := api.Send(answer)
		if err != nil {
			log.Fatal(err)
		}

		t.logger.Infof("[%d][%s] %s", message.From.ID, message.From.UserName, message.Text)
	}

	return nil
}

func (t *tgBot) configureLogger() error {
	level, err := logrus.ParseLevel(t.config.LogLevel)
	if err != nil {
		return err
	}

	t.logger.SetLevel(level)

	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FullTimestamp = true
	t.logger.SetFormatter(formatter)

	return nil
}

func (t *tgBot) getUpdates(api *tgbotapi.BotAPI) (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := api.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	return updates, nil
}

func (t *tgBot) handleUpdate(ctx context.Context, update tgbotapi.Update) *tgbotapi.MessageConfig {
	sessionID := sha256HashFromInt(update.Message.From.ID)

	text, err := DetectIntentText(ctx, t.config.GoogleCloudProjectID, sessionID, update.Message.Text, RuLanguageCode)
	if err != nil {
		t.logger.Error(err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID

	return &msg
}
