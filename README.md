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

TELEGRAM_API_TOKEN=602736695:AAG4lScHIBg1_OG6_q88LDqnFAPtrgwe0oA;GOOGLE_CLOUD_PROJECT_ID=zetraisonai;SESSION_ID=123;GOOGLE_APPLICATION_CREDENTIALS=/home/zetraison/go/src/dialogflow-telegram-bot/zetraisonai-964f7ad95704.json