package location

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LocationData struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Elevation   float64 `json:"elevation"`
	FeatureCode string  `json:"feature_code"`
	CountryCode string  `json:"country_code"`
	Timezone    string  `json:"timezone"`
	Population  int     `json:"population"`
}

func GetLocationFromAPI(city string) (*LocationData, error) {
	apiURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", city)

	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Results []LocationData `json:"results"`
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Results) == 0 {
		return nil, fmt.Errorf("Город не найден: %s", city)
	}

	return &result.Results[0], nil
}
