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
	log.Println("Starting transactions-service on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
