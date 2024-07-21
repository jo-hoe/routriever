package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/jo-hoe/routriever/app"
)

const defaultPort = "8080"

func main() {
	_, err := app.NewRoutrieverService()
	if err != nil {
		log.Fatal("could not create routriever service")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", probeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func probeHandler(ctx echo.Context) (err error) {
	return ctx.NoContent(http.StatusOK)
}
