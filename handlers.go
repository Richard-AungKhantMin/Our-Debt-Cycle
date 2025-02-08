package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Define the path to the HTML file (e.g., "./templates/index.html")
		filePath := filepath.Join("templates", "index.html")

		// Open the HTML file
		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "Unable to load the homepage", http.StatusInternalServerError)
			log.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()

		// Serve the HTML file
		http.ServeFile(w, r, filePath)
	} else {
		// For non-GET methods, send a 405 Method Not Allowed response
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
