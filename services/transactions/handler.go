package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FromAccountID string  `json:"from_account_id"`
		ToAccountID   string  `json:"to_account_id"`
		Amount        float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("DB insert failed:", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	txID := uuid.New().String()

	tx, err := db.Begin()
	if err != nil {
		log.Println("DB insert failed:", err)
		http.Error(w, "Failed to begin transaction", http.StatusInternalServerError)
		return
	}

	// Deduct from sender
	_, err = tx.Exec(
		"UPDATE accounts SET balance = balance - $1 WHERE id = $2 AND balance >= $1",
		input.Amount, input.FromAccountID,
	)
	if err != nil {
		tx.Rollback()
		log.Println("DB insert failed:", err)
		http.Error(w, "Failed to deduct balance", http.StatusBadRequest)
		return
	}

	// Credit to receiver
	_, err = tx.Exec(
		"UPDATE accounts SET balance = balance + $1 WHERE id = $2",
		input.Amount, input.ToAccountID,
	)
	if err != nil {
		tx.Rollback()
		log.Println("DB insert failed:", err)
		http.Error(w, "Failed to credit balance", http.StatusBadRequest)
		return
	}

	// Insert transaction
	_, err = tx.Exec(
		"INSERT INTO transactions (id, from_account_id, to_account_id, amount) VALUES ($1, $2, $3, $4)",
		txID, input.FromAccountID, input.ToAccountID, input.Amount,
	)
	if err != nil {
		tx.Rollback()
		log.Println("DB insert failed:", err)
		http.Error(w, "Failed to save transaction", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		log.Println("DB insert failed:", err)
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	t := Transaction{
		ID:            txID,
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Amount:        input.Amount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)

	log.Printf("Transaction request: %+v", input)

}

func ListTransactions(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, from_account_id, to_account_id, amount FROM transactions")
	if err != nil {
		http.Error(w, "Failed to read transactions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var list []Transaction
	for rows.Next() {
		var t Transaction
		if err := rows.Scan(&t.ID, &t.FromAccountID, &t.ToAccountID, &t.Amount); err == nil {
			list = append(list, t)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
