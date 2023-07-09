package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type WeatherData struct {
	Hourly struct {
		Time          []string  `json:"time"`
		Temperature2m []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}

func GetWeatherData(latitude, longitude float64) (*WeatherData, error) {
	apiURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=temperature_2m", latitude, longitude)

	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherData
	err = json.Unmarshal(data, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}
