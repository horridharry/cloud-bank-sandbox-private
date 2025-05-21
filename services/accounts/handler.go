package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	acc := Account{
		ID:      uuid.New().String(),
		Name:    input.Name,
		Balance: 0.0,
	}

	mu.Lock()
	accounts[acc.ID] = acc
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}

func ListAccounts(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	list := []Account{}
	for _, acc := range accounts {
		list = append(list, acc)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
