package app

import (
	"context"
	"errors"
	"net/http"
	"os"
)

type RoutrieverService interface {
	GetRouteDistance(ctx context.Context, start GPSCoordinates, end GPSCoordinates) (travelTimeInSeconds int, err error)
}

type GPSCoordinates struct {
	Latitude  float32
	Longitude float32
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
