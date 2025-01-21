package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"blockchain_project/blockchain"
)

func GetBlockchain(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bc.mux.Lock()
		defer bc.mux.Unlock()
		json.NewEncoder(w).Encode(bc.Blocks)
	}
}

func AddTransaction(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transactions []blockchain.Transaction
		if err := json.NewDecoder(r.Body).Decode(&transactions); err != nil {
			http.Error(w, "Invalid transaction data", http.StatusBadRequest)
			return
		}
		bc.AddBlock(transactions)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Block added successfully")
	}
}