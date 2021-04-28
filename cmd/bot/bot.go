package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const telegramTokenEnv string = "TELEGRAM_TOKEN"

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv(telegramTokenEnv))
	if err != nil {
		log.Fatalf("Can't connect to telegram API: %s", err.Error())
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		// msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, nil)

		switch update.Message.Command() {
		case "help":
			msg.Text = "Write /status command"
		case "status":
			msg.Text = "OK"
		default:
			msg.Text = "Try /help command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

}
