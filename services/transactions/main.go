package main

import (
	"log"
	"net/http"
)

func main() {
	r := SetupRoutes()
	log.Println("Starting transactions-service on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
