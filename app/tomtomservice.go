package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	base_url     = "https://api.tomtom.com"
	routing_path = "/routing/1/calculateRoute/"
	routing_url  = base_url + routing_path
)

type TomTomService struct {
	apiKey     string
	httpClient *http.Client
}

type tomTomCalcRouteResponse struct {
	Routes []struct {
		Summary struct {
			TravelTimeInSeconds int `json:"travelTimeInSeconds"`
		} `json:"summary"`
	} `json:"routes"`
}

func NewTomTomService(apiKey string, httpClient *http.Client) *TomTomService {
	return &TomTomService{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (t *TomTomService) GetRouteDistance(ctx context.Context, start GPSCoordinates, end GPSCoordinates) (travelTimeInSeconds int, err error) {
	travelTimeInSeconds = -1

	response, err := t.httpClient.Get(t.generateURL(start, end))
	if err != nil {
		return travelTimeInSeconds, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return travelTimeInSeconds, fmt.Errorf("status code: %d", response.StatusCode)
	}
	var responseData tomTomCalcRouteResponse
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return travelTimeInSeconds, err
	}

	return responseData.Routes[0].Summary.TravelTimeInSeconds, nil
}

func (t *TomTomService) generateURL(start GPSCoordinates, end GPSCoordinates) string {
	return fmt.Sprintf("%s%.7f,%.7f:%.7f,%.7f/json?&key=%s", routing_url, start.Latitude, start.Longitude, end.Latitude, end.Longitude, t.apiKey)
}