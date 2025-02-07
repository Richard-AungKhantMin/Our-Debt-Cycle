package main

import (
	"encoding/json"
	"net/http"
)

func addTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var transaction struct {
		Payer  string  `json:"payer"`
		Payee  string  `json:"payee"`
		Amount float64 `json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process transaction (adapt your existing logic here)
	amount := calculateAmounts(transaction.Payer, transaction.Amount)
	payers, err := getPayers(transaction.Payer, transaction.Payee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, payer := range payers {
		for _, payee := range payer.ToPay {
			if payee.Username == transaction.Payee {
				payee.Amount += amount
			}
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func getBalancesHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	balances := make(map[string]map[string]float64)
	for _, user := range allUsers {
		balances[user.Username] = make(map[string]float64)
		for _, payee := range user.ToPay {
			balances[user.Username][payee.Username] = payee.Amount
		}
	}

	json.NewEncoder(w).Encode(balances)
}

func cleanupHandler(w http.ResponseWriter, _ *http.Request) {
	cleanUpAmounts()
	w.WriteHeader(http.StatusOK)
}
