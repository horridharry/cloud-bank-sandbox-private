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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)

	_, err := db.Exec(
		"INSERT INTO accounts (id, name, balance) VALUES ($1, $2, $3)",
		acc.ID, acc.Name, acc.Balance,
	)
	if err != nil {
		http.Error(w, "Failed to insert account", http.StatusInternalServerError)
		return
	}

}

func ListAccounts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, balance FROM accounts")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var list []Account
	for rows.Next() {
		var acc Account
		if err := rows.Scan(&acc.ID, &acc.Name, &acc.Balance); err != nil {
			continue
		}
		list = append(list, acc)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func UpdateBalance(w http.ResponseWriter, r *http.Request) {
	var input struct {
		AccountID string  `json:"account_id"`
		Amount    float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Use a single UPDATE with RETURNING to get new balance
	row := db.QueryRow(`
        UPDATE accounts
        SET balance = balance + $1
        WHERE id = $2
        RETURNING id, name, balance
    `, input.Amount, input.AccountID)

	var acc Account
	if err := row.Scan(&acc.ID, &acc.Name, &acc.Balance); err != nil {
		http.Error(w, "Account not found or DB error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
