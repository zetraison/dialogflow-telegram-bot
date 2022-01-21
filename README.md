# dialogflow-telegram-bot
Telegram Bot based on [Google Dialogflow ES](https://cloud.google.com/dialogflow/es/docs/quick/api).

![Go](https://github.com/zetraison/dialogflow-telegram-bot/workflows/Go/badge.svg)

## Environment

- GOOGLE_APPLICATION_CREDENTIALS - Google service account key file
- GOOGLE_CLOUD_PROJECT_ID - Google Cloud Project identifier
- TELEGRAM_API_TOKEN - Telegram bot API token

## Install modules

```bash
go mod tidy
```

## Build

```bash
go build -o bot
```

## Run

```bash 
export TELEGRAM_API_TOKEN={TELEGRAM_API_TOKEN} GOOGLE_CLOUD_PROJECT_ID={GOOGLE_CLOUD_PROJECT_ID} GOOGLE_APPLICATION_CREDENTIALS={GOOGLE_APPLICATION_CREDENTIALS} && ./bot
```