package gpsservice

import (
	"os"
	"reflect"
	"testing"
)

func TestNewRoutrieverServiceWithoutEnvVar(t *testing.T) {
	service, err := NewRoutrieverService()
	if err == nil {
		t.Error("expected no error, got nil")
	}
	if service != nil {
		t.Error("expected nil service, got", reflect.TypeOf(service))
	}
}

func TestNewRoutrieverService(t *testing.T) {
	os.Setenv(TomTomApiKeyEnvVar, "test")
	defer os.Unsetenv(TomTomApiKeyEnvVar)

	service, err := NewRoutrieverService()
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if service == nil {
		t.Error("expected non-nil service, got nil")
	}
}
