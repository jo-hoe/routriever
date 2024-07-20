package app

import "context"

type RoutrieverService interface {
	GetRouteDistance(ctx context.Context, start GPSCoordinates, end GPSCoordinates) (travelTimeInSeconds int, err error)
}

type GPSCoordinates struct {
	Latitude  float32
	Longitude float32
}
