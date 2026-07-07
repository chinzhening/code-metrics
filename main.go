package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// check if .env file exists and load it
	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		log.Println(".env file does not exist, skipping loading environment variables")
	} else {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: http.HandlerFunc(helloWorld),
	}

	log.Println("Server listening on http://localhost:" + port)

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Error starting server", err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
