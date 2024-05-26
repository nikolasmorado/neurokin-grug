package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/nikolasmorado/neurokin-grug/api"
)

func main() {

  if err := godotenv.Load(); err != nil {
    log.Fatal("Error loading .env file")
  }

  listenAddr := os.Getenv("LISTEN_ADDR")

  slog.Info("HTTP server started", "listenAddr", listenAddr)

  api := api.NewServer(listenAddr)

  api.Start()
}
