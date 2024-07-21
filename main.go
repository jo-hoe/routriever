package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/jo-hoe/routriever/app/config"
	"github.com/jo-hoe/routriever/app/service"
)

const (
	portEnvVar  = "PORT"
	defaultPort = "8080"

	configPathEnvVar  = "CONFIG_PATH"
	defaultConfigPath = "./config.yaml"
)

var (
// serviceInstance service.RoutrieverService
// serviceConfig   config.Config
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", probeHandler)

	port := os.Getenv(portEnvVar)
	if port == "" {
		port = defaultPort
	}

	// start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func init() {
	configPath := os.Getenv(configPathEnvVar)
	if configPath == "" {
		log.Printf("CONFIG_PATH not set, using default path '%s'", defaultConfigPath)
		configPath = defaultConfigPath
	}

	_, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatal("could not read config")
	}

	_, err = service.NewRoutrieverService()
	if err != nil {
		log.Fatal("could not create routriever service")
	}
}

func probeHandler(ctx echo.Context) (err error) {
	return ctx.NoContent(http.StatusOK)
}
