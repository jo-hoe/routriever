package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const updateIntervalDefault = "1m"

type Config struct {
	Routes         []Route `yaml:"routes"`
	UpdateInterval string  `yaml:"updateInterval"`
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
	} else {
		// set default update interval if not set and no failure detected
		if config.UpdateInterval == "" {
			config.UpdateInterval = updateIntervalDefault
		}
	}

	return config, err
}
