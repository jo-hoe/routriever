package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"testing"
)

const envFileDir = "test"
const envFileName = "testconfig.yaml"

func Test_GetConfig(t *testing.T) {
	type args struct {
		configPath string
	}
	tests := []struct {
		name       string
		args       args
		wantConfig Config
		wantErr    bool
	}{
		{
			name: "positive test",
			args: args{
				configPath: GetConfigFilePath(),
			},
			wantConfig: Config{
				Routes: []Route{
					{
						Name: "Route to TomTom HQs",
						Start: GPSCoordinates{
							Latitude:  52.3764134,
							Longitude: 4.908321,
						},
						End: GPSCoordinates{
							Latitude:  51.4125186,
							Longitude: 5.4505796,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "load incorrect file path",
			args: args{
				configPath: "invalid",
			},
			wantConfig: Config{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConfig, err := GetConfig(tt.args.configPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("GetConfig() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}

func GetConfigFilePath() string {
	current_directory, error := os.Getwd()
	if error != nil {
		log.Fatal(error)
	}

	appFolderPath := filepath.Dir(current_directory)
	workingDirectoryPath := filepath.Dir(appFolderPath)
	return path.Join(workingDirectoryPath, envFileDir, envFileName)
}
