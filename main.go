package main

import (
	"fmt"
	"log"
	request "weatherbottelegram/request"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Ошибка чтения конфигурации:", err)
		return
	}

	value := viper.GetString("API-KEY")
	bot, err := tgbotapi.NewBotAPI(value)
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
			request.HandleRequest(bot, update.Message)

		}
	}
}
