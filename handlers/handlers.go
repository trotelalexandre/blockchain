package handlers

import (
	"blockchain/blockchain"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetBlockchain(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bc.Mux.Lock()
		defer bc.Mux.Unlock()
		json.NewEncoder(w).Encode(bc.Blocks)
	}
}

func AddTransaction(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transactions []blockchain.Transaction
		miner := r.URL.Query().Get("miner")
		if miner == "" {
			http.Error(w, "Miner address is required", http.StatusBadRequest)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&transactions); err != nil {
			http.Error(w, "Invalid transaction data", http.StatusBadRequest)
			return
		}
		bc.AddBlock(transactions, miner)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Block added successfully")
	}
}