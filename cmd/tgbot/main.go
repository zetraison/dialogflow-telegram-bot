package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/zetraison/dialogflow-telegram-bot/internal/app/tgbot"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/tgbot.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := tgbot.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	err = config.LoadFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	bot := tgbot.New(config)
	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
