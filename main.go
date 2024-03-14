package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Epiq122/dreampic-ai/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))
	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("application running", "port", port)
	log.Fatal(http.ListenAndServe(os.Getenv("HTTP_LISTEN_ADDR"), router))
}

func initEverything() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}
