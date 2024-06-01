package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	a "neurokin/api" 
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	listenAddr := os.Getenv("LISTEN_ADDR")

	slog.Info("HTTP server started", "listenAddr", listenAddr)

	api := a.NewServer(
		listenAddr,
		Public(),
	)

	api.Start()
}
