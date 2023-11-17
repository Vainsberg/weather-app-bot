package request

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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

		latitudeText := words[0]
		longitudeText := words[1]

		backendURL := "http://localhost:8080/get_weather?latitude=" + latitudeText + "&longitude=" + longitudeText

		resp, err := http.Get(backendURL)
		if err != nil {
			log.Println("Ошибка при выполнении HTTP-запроса:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Ошибка при чтении ответа:", err)
			return
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, string(body))
		bot.Send(msg)
	}
	if message.Location != nil {
		latitude := message.Location.Latitude
		longitude := message.Location.Longitude
		latitudetext := fmt.Sprintf("%f", latitude)
		longitudetext := fmt.Sprintf("%f", longitude)
		backendURL := "http://localhost:8080/get_weather?latitude=" + latitudetext + "&longitude=" + longitudetext
		resp, err := http.Get(backendURL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		Msg := tgbotapi.NewMessage(message.Chat.ID, string(body))
		bot.Send(Msg)
	}
}
