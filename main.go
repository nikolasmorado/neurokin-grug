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

  url := os.Getenv("SUPA_URL")
  key := os.Getenv("SUPA_SK")

  store, err := store.InitSupabase(url, key)

  if err != nil {
    log.Fatal("Error initializing Supabase")
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
