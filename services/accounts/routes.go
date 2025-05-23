package main

import (
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/accounts", CreateAccount).Methods("POST")
	r.HandleFunc("/accounts", ListAccounts).Methods("GET")
	r.HandleFunc("/accounts/update-balance", UpdateBalance).Methods("POST")
	r.HandleFunc("/healthz", Healthz).Methods("GET")
	
	return r
}
