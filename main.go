package main

import (
	"log"
	"log/slog"
	"os"

	a "neurokin/api"
	"neurokin/store"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

  store, err := store.InitFirebase("./ServiceAccount.json")

  if err != nil {
    log.Fatal("Error initializing Firebase")
  }

	listenAddr := os.Getenv("LISTEN_ADDR")

	slog.Info("HTTP server started", "listenAddr", listenAddr)

	api := a.NewServer(
		listenAddr,
    store,
		Public(),
	)

	api.Start()
}
