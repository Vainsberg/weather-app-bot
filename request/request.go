package request

import (
	"fmt"
	"strings"
	weatherinfo "weatherbottelegram/weatherInfo"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleRequest(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if message.Text != "" {
		words := strings.Split(message.Text, ",")
		for i, word := range words {
			words[i] = strings.TrimSpace(word)
		}
		if len(words) != 2 {
			errorMsg := "Неправильная форма заполнения. Пожалуйста, введите координаты в формате 'широта,долгота'."
			reply := tgbotapi.NewMessage(message.Chat.ID, errorMsg)
			bot.Send(reply)
			return
		}
		latitudetext := words[0]
		longitudetext := words[1]

		msg := tgbotapi.NewMessage(message.Chat.ID, weatherinfo.GetWeatherInfo(latitudetext, longitudetext))
		bot.Send(msg)
	} else if message.Location != nil {
		latitude := message.Location.Latitude
		longitude := message.Location.Longitude
		latitudetext := fmt.Sprintf("%f", latitude)
		longitudetext := fmt.Sprintf("%f", longitude)
		Msg := tgbotapi.NewMessage(message.Chat.ID, weatherinfo.GetWeatherInfo(latitudetext, longitudetext))
		bot.Send(Msg)
	}
}
