SHELL:=/bin/bash

.EXPORT_ALL_VARIABLES:

install:
	go mod tidy

clean:
	rm -rf bin

build:
	GO111MODULE=on go build -o bin/tgbot ./cmd/tgbot

run:
	GOOGLE_APPLICATION_CREDENTIALS=./configs/google-console-credentials.json ./bin/tgbot --config-path configs/tgbot_prod.toml

up: clean build run

go-run:
	go run cmd/tgbot/main.go