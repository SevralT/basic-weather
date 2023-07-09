package main

import (
	"fmt"
	"log"
	"time"
  "os"

	"github.com/SevralT/basic-weather/config"
	"github.com/SevralT/basic-weather/location"
	"github.com/SevralT/basic-weather/weather"
)

func main() {
	configData, err := config.LoadOrCreateConfig("config.json")
	if err != nil {
		log.Fatal("Ошибка при загрузке/создании файла конфигурации:", err)
	}

	if configData.Latitude == 0 || configData.Longitude == 0 || len(os.Args) > 1 && os.Args[1] == "edit" {
		locationData, err := location.GetLocationFromAPI(configData.City)
		if err != nil {
			log.Fatal("Ошибка при получении координат города:", err)
		}

		configData.Latitude = locationData.Latitude
		configData.Longitude = locationData.Longitude

		err = config.SaveConfig("config.json", configData)
		if err != nil {
			log.Fatal("Ошибка при сохранении файла конфигурации:", err)
		}
	}

	weatherData, err := weather.GetWeatherData(configData.Latitude, configData.Longitude)
	if err != nil {
		log.Fatal("Ошибка при получении данных о погоде:", err)
	}

	currentTime := time.Now().Format("2006-01-02T15:00") // Форматирование текущего времени в формате "YYYY-MM-DDTHH:00"

	fmt.Printf("Текущая температура для города %s:\n", configData.City)

	for i, time := range weatherData.Hourly.Time {
		if time == currentTime {
			temp := weatherData.Hourly.Temperature2m[i]
			fmt.Printf("%.1f°C\n", temp)
			break
		}
	}
}
