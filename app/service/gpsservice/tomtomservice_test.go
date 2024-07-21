package gpsservice

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/jo-hoe/routriever/app/config"
	"github.com/jo-hoe/routriever/test"
)

func TestInterfaceImplementation(t *testing.T) {
	var _ RoutrieverService = (*TomTomService)(nil)
}

func TestTomTomService_generateURL(t *testing.T) {
	testApiKey := "testApiKey"

	tomTomService := NewTomTomService(testApiKey, nil)

	type args struct {
		start config.GPSCoordinates
		end   config.GPSCoordinates
	}
	tests := []struct {
		name string
		tr   *TomTomService
		args args
		want string
	}{
		{
			name: "positive test",
			tr:   tomTomService,
			args: args{
				start: config.GPSCoordinates{
					Latitude:  1.0,
					Longitude: 2.0,
				},
				end: config.GPSCoordinates{
					Latitude:  3.0,
					Longitude: 4.0,
				},
			},
			want: "https://api.tomtom.com/routing/1/calculateRoute/1.0000000,2.0000000:3.0000000,4.0000000/json?&key=testApiKey",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.generateURL(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("TomTomService.generateURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTomTomService_GetRouteDistance(t *testing.T) {
	client := test.NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: io.NopCloser(bytes.NewBufferString(`{
				"formatVersion": "0.0.12",
				"routes": [
					{
						"summary": {
							"lengthInMeters": 1350,
							"travelTimeInSeconds": 218
						}
					}
				]
			}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	tomTomService := NewTomTomService("", client)

	type args struct {
		ctx   context.Context
		start config.GPSCoordinates
		end   config.GPSCoordinates
	}
	tests := []struct {
		name                    string
		tr                      *TomTomService
		args                    args
		wantTravelTimeInSeconds int
		wantErr                 bool
	}{
		{
			name: "positive test",
			tr:   tomTomService,
			args: args{
				ctx:   context.Background(),
				start: config.GPSCoordinates{Latitude: 1.0, Longitude: 2.0},
				end:   config.GPSCoordinates{Latitude: 3.0, Longitude: 4.0},
			},
			wantTravelTimeInSeconds: 218,
			wantErr:                 false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTravelTimeInSeconds, err := tt.tr.GetRouteDistance(tt.args.ctx, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("TomTomService.GetRouteDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTravelTimeInSeconds != tt.wantTravelTimeInSeconds {
				t.Errorf("TomTomService.GetRouteDistance() = %v, want %v", gotTravelTimeInSeconds, tt.wantTravelTimeInSeconds)
			}
		})
	}
}
