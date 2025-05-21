package main

import (
    "log"
    "net/http"
)

func main() {
    r := SetupRoutes()
    log.Println("Starting accounts-service on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
