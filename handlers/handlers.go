package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/trotelalexandre/proto/blockchain"
	"github.com/trotelalexandre/proto/stats"
)

func GetBlockchain(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bc.Mux.Lock()
		defer bc.Mux.Unlock()
		json.NewEncoder(w).Encode(bc.Blocks)
	}
}

func SendTransaction(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction blockchain.Transaction
		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			http.Error(w, "Invalid transaction data", http.StatusBadRequest)
			return
		}
		bc.AddTransaction(transaction)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Transaction added successfully")
	}
}

func AddWallet(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wallet := blockchain.AddWallet()
		bc.AddWallet(wallet)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(wallet)
	}
}

func GetBlockCount(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, stats.GetBlockCount(bc))
	}
}

func GetTransactionCount(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, stats.GetTransactionCount(bc))
	}
}

func GetWalletCount(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, stats.GetWalletCount(bc))
	}
}

func GetBlockReward(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, stats.GetBlockReward(bc))
	}
}

func GetAllBlocks(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stats.GetAllBlocks(bc))
	}
}

func GetAllTransactions(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stats.GetAllTransactions(bc))
	}
}

func GetWalletBalance(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")
		if address == "" {
			http.Error(w, "Address is required", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, stats.GetWalletBalance(bc, address))
	}
}