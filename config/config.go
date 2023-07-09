package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	City      string  `json:"city"`
}

func LoadOrCreateConfig(filename string) (*Config, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		config := &Config{}

		fmt.Println("Первоначальная настройка:")
		fmt.Print("Введите название города: ")
		_, err = fmt.Scan(&config.City)
		if err != nil {
			return nil, err
		}

		err := SaveConfig(filename, config)
		if err != nil {
			return nil, err
		}
		return config, nil
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(filename string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
