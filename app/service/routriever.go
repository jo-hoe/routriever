package service

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/jo-hoe/routriever/app/config"
)

type RoutrieverService interface {
	GetRouteDistance(ctx context.Context, start config.GPSCoordinates, end config.GPSCoordinates) (travelTimeInSeconds int, err error)
}

func NewRoutrieverService() (result RoutrieverService, err error) {
	tomTomApiKey := os.Getenv(TomTomApiKeyEnvVar)
	if tomTomApiKey != "" {
		result = NewTomTomService(tomTomApiKey, http.DefaultClient)
	} else {
		err = errors.New(TomTomApiKeyEnvVar + " not set")
	}

	return result, err
}
