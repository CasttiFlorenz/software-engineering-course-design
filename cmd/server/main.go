package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CasttiFlorenz/software-engineering-course-design/internal/httpapi"
)

func main() {
	address := os.Getenv("HTTP_ADDR")
	if address == "" {
		address = "127.0.0.1:8080"
	}

	server := &http.Server{
		Addr:              address,
		Handler:           httpapi.NewHandler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       time.Minute,
	}

	log.Printf("研改 API listening on %s", address)
	log.Fatal(server.ListenAndServe())
}
