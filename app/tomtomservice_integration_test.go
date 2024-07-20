package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"

	gotdotenv "github.com/joho/godotenv"
)

const integrationTestApiKeyEnvVar = "INTEGRATION_TEST_API_KEY"

// calls the api and uses either the .env
func Test_Integration_TomTomService_GetRouteDistance(t *testing.T) {
	// try to load .env file
	err := gotdotenv.Load() // load .env file
	if err != nil {
		log.Print(".env file not found")
	}
	integrationTestApiKey := os.Getenv(integrationTestApiKeyEnvVar)
	if integrationTestApiKey == "" {
		t.Skip("Integration test skipped because there was no API key in .env file")
	}

	tomTomService := NewTomTomService(integrationTestApiKey, http.DefaultClient)

	type args struct {
		ctx   context.Context
		start GPSCoordinates
		end   GPSCoordinates
	}
	tests := []struct {
		name    string
		tr      *TomTomService
		args    args
		wantErr bool
	}{
		{
			name: "integration test",
			tr:   tomTomService,
			args: args{
				ctx: context.Background(),
				// TomTom Amsterdam
				start: GPSCoordinates{Latitude: 52.3764134, Longitude: 4.908321},
				// TomTom Eindhoven
				end: GPSCoordinates{Latitude: 51.4125186, Longitude: 5.4505796},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTravelTimeInSeconds, err := tt.tr.GetRouteDistance(tt.args.ctx, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("TomTomService.GetRouteDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTravelTimeInSeconds <= 0 {
				t.Errorf("TomTomService.GetRouteDistance() = %v, time greater then 0", gotTravelTimeInSeconds)
			}
		})
	}
}
