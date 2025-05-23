package main

import (
	"log"
	"net/http"
)

func main() {
	if err := InitDB(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	r := SetupRoutes()
	log.Println("Starting accounts-service on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
