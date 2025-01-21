package main

import (
	"blockchain/blockchain"
	"blockchain/handlers"
	"fmt"
	"net/http"
)

func main() {
	genesisBlock := blockchain.CreateGenesisBlock()
	bc := &blockchain.Blockchain{
		Blocks: []*blockchain.Block{genesisBlock},
		Reward: 50,
	}

	http.HandleFunc("/blockchain", handlers.GetBlockchain(bc))
	http.HandleFunc("/add", handlers.AddTransaction(bc))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}