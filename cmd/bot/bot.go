package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
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

		switch update.Message.Command() {
		case "help":
			msg.Text = "Write /status command"
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		case "status":
			msg.Text = "OK"
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		case "image":
			URL := "http://127.0.0.1:9000/img.png"

			photoBytes, err := GetImageFromURL(URL)
			if err != nil {
				log.Fatal(err)
			}

			photoFileBytes := tgbotapi.FileBytes{
				Name:  "picture",
				Bytes: photoBytes,
			}
			bot.Send(tgbotapi.NewPhotoUpload(update.Message.Chat.ID, photoFileBytes))
		default:
			msg.Text = "Try /help command"
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

	}

}

func GetImageFromURL(URL string) ([]byte, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("Received non 200 response code")
	}

	photoBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return photoBytes, nil
}
