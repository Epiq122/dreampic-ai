package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Epiq122/dreampic-ai/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	router.Use(handler.WithUser)

	// render config file - static file server ( FILE SYSTEM )
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))

	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))
	router.Get("/login", handler.MakeHandler(handler.HandleLoginIndex))

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
