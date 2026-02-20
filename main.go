package main

import (
	"log"
	"main/Handlers"
	"net/http"
	"os"
)

func main() {
	// Checks if there is a env variable for port
	port := os.Getenv("PORT")

	// Override port with default port (8080) if not provided (e.g. local deployment)
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	router := http.NewServeMux()

	// All the handlers that my API is providing
	router.HandleFunc("/countryinfo/v1/status", Handlers.StatusHandler)
	router.HandleFunc("/countryinfo/v1/exchange/{p1}", Handlers.ExchangeHandler)
	router.HandleFunc("/countryinfo/v1/info/{p1}", Handlers.InfoHandler)

	log.Println("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
