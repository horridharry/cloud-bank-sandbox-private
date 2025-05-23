package main

import "github.com/gorilla/mux"

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/transactions", CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions", ListTransactions).Methods("GET")
	r.HandleFunc("/healthz", Healthz).Methods("GET")
	return r
}
