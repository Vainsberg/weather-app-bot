package request

import (
	"fmt"
	"log"
	"strings"
	weatherinfo "weatherbottelegram/weatherInfo"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleRequest(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	var latitudetext, longitudetext string
	if message.Text != "" {
		words := strings.Split(message.Text, ",")
		for i, word := range words {
			words[i] = strings.TrimSpace(word)
		}
		if len(words) != 2 {
			errorMsg := "Неправильная форма заполнения. Пожалуйста, введите координаты в формате 'широта,долгота'."
			reply := tgbotapi.NewMessage(message.Chat.ID, errorMsg)
			_, err := bot.Send(reply)
			if err != nil {
				log.Println("Ошибка при отправке сообщения боту:", err)
			}
			return
		}
		latitudetext = words[0]
		longitudetext = words[1]

	} else if message.Location != nil {
		latitude := message.Location.Latitude
		longitude := message.Location.Longitude
		latitudetext = fmt.Sprintf("%f", latitude)
		longitudetext = fmt.Sprintf("%f", longitude)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, weatherinfo.GetWeatherInfo(latitudetext, longitudetext))
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Ошибка при отправке сообщения боту:", err)
	}
}
