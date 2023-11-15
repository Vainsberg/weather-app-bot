package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	var apiToken string
	flag.StringVar(&apiToken, "api-token", "", "API токен")
	flag.Parse()
	if apiToken == "" {
		fmt.Println("Необходимо указать API токен")
		return
	}
	fmt.Printf("API токен: %s\n", apiToken)
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Location != nil {
				latitude := update.Message.Location.Latitude
				longitude := update.Message.Location.Longitude
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Попучены координаты: %f,%f", latitude, longitude))
				bot.Send(msg)
				latitudetext := fmt.Sprintf("%f", latitude)
				longitudetext := fmt.Sprintf("%f", longitude)
				backendURL := "https://localhost:8080/get_weather?latitude=" + latitudetext + "&longitude=" + longitudetext
				resp, err := http.Get(backendURL)
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()
			}

		}

	}
}
