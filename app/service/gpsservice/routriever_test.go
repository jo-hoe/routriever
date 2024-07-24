package gpsservice

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"testing"
)

const secretFileDir = "dev"
const secretFileName = "secret.txt"

func TestNewRoutrieverServiceWithoutPath(t *testing.T) {
	service, err := NewRoutrieverService("")
	if err == nil {
		t.Error("expected no error, got nil")
	}
	if service != nil {
		t.Error("expected nil service, got", reflect.TypeOf(service))
	}
}

func TestNewRoutrieverService(t *testing.T) {
	service, err := NewRoutrieverService(GetSecretFilePath())
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if service == nil {
		t.Error("expected non-nil service, got nil")
	}
}

func GetSecretFilePath() string {
	current_directory, error := os.Getwd()
	if error != nil {
		log.Fatal(error)
	}

	serviceFolder := filepath.Dir(current_directory)
	appFolderPath := filepath.Dir(serviceFolder)
	workingDirectoryPath := filepath.Dir(appFolderPath)
	return path.Join(workingDirectoryPath, secretFileDir, secretFileName)
}
