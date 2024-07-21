package app

import (
	"context"
	"log"
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

func NewRoutrieverService() RoutrieverService {
	tomTomApiKey := os.Getenv(TomTomApiKeyEnvVar)
	if tomTomApiKey == "" {
		log.Printf(TomTomApiKeyEnvVar + " not set")
		return nil
	} else {
		return NewTomTomService(tomTomApiKey, http.DefaultClient)
	}
}