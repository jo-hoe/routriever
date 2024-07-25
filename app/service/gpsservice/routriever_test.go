package gpsservice

import (
	"os"
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
	tempfile := CreateTempFile(t)
	defer os.Remove(tempfile)

	service, err := NewRoutrieverService(tempfile)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if service == nil {
		t.Error("expected non-nil service, got nil")
	}
}

func CreateTempFile(t *testing.T) string {
	file, err := os.CreateTemp("", "temp-*.txt")
	if err != nil {
		t.Error(err)
	}

	_, err = file.WriteString("demo content")
	if err != nil {
		t.Error(err)
	}

	err = file.Close()
	if err != nil {
		t.Error(err)
	}

	return file.Name()
}

