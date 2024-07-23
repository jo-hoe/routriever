package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Routes []Route `yaml:"routes"`
}

type Route struct {
	Name  string         `yaml:"name"`
	Start GPSCoordinates `yaml:"start"`
	End   GPSCoordinates `yaml:"end"`
}

type GPSCoordinates struct {
	Latitude  float32 `yaml:"latitude"`
	Longitude float32 `yaml:"longitude"`
}

func GetConfig(configPath string) (config Config, err error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		err = fmt.Errorf("could not read config file error: '%v'", err)
		return config, err
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		err = fmt.Errorf("could not unmarshal config file error: '%v'", err)
	}

	return config, err
}

func GetConfigFromFile(configPath string) (config Config, err error) {
	config, err = GetConfig(configPath)
	if err != nil {
		log.Fatal("could not read config")
	}
	return config, err
}
