package gpsservice

import (
	"net/http"
	"os"

	"github.com/jo-hoe/routriever/app/config"
)

type RoutrieverService interface {
	GetRouteDistance(start config.GPSCoordinates, end config.GPSCoordinates) (travelTimeInSeconds int, err error)
}

func NewRoutrieverService(secretPath string) (result RoutrieverService, err error) {
	apiKey, err := readSecretFromFile(secretPath)
	if err != nil {
		return result, err
	}

	result = NewTomTomService(apiKey, http.DefaultClient)

	return result, err
}

func readSecretFromFile(secretPath string) (secret string, err error) {
	data, err := os.ReadFile(secretPath)
	if err != nil {
		return secret, err
	}

	secret = string(data)
	return secret, err
}
