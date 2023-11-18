package weatherinfo

import (
	"io"
	"log"
	"net/http"
	"weatherbottelegram/viper"
)

func GetWeatherInfo(latitude string, longitude string) string {

	backendURL := viper.WeatherApi() + "get_weather?latitude=" + latitude + "&longitude=" + longitude

	resp, err := http.Get(backendURL)
	if err != nil {
		log.Println("Ошибка при выполнении HTTP-запроса:", err)
		return "Ошибка"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка при чтении ответа:", err)
		return "Ошибка"
	}
	return string(body)
}
