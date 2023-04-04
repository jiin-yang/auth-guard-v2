package main

import (
	"github.com/jiin-yang/auth-guard-v2/api-gateway/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	e := echo.New()
	client := http.Client{Timeout: 10 * time.Second}

	handler.NewHandler(e, &client)

	if err := e.Start(os.Getenv("API_GATEWAY_PORT")); err != http.ErrServerClosed {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
