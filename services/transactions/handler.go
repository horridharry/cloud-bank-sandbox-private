package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var transactions []Transaction // In-memory for now

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var input struct {
		From string  `json:"from_account_id"`
		To   string  `json:"to_account_id"`
		Amt  float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	txn := Transaction{
		ID:            uuid.New().String(),
		FromAccountID: input.From,
		ToAccountID:   input.To,
		Amount:        input.Amt,
	}

	transactions = append(transactions, txn)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(txn)

	if !accountExists(input.From) || !accountExists(input.To) {
		http.Error(w, "One or both accounts do not exist", http.StatusBadRequest)
		// Debit sender
		http.Post("http://accounts-service:8080/accounts/update-balance", "application/json",
			strings.NewReader(fmt.Sprintf(`{"account_id":"%s", "amount":%f}`, input.From, -input.Amt)))

		// Credit receiver
		http.Post("http://accounts-service:8080/accounts/update-balance", "application/json",
			strings.NewReader(fmt.Sprintf(`{"account_id":"%s", "amount":%f}`, input.To, input.Amt)))

		return
	}
}

func ListTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func accountExists(accountID string) bool {
	resp, err := http.Get("http://accounts-service:8080/accounts")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var accounts []Account
	if err := json.NewDecoder(resp.Body).Decode(&accounts); err != nil {
		return false
	}

	for _, acc := range accounts {
		if acc.ID == accountID {
			return true
		}
	}
	return false
}
