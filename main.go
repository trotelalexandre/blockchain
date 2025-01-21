package main

import (
	"fmt"
	"net/http"
	"blockchain/blockchain"
	"blockchain/handlers"
)

func main() {
	genesisBlock := blockchain.CreateGenesisBlock()
	bc := &blockchain.Blockchain{Blocks: []*blockchain.Block{genesisBlock}}

	http.HandleFunc("/blockchain", handlers.GetBlockchain(bc))
	http.HandleFunc("/add", handlers.AddTransaction(bc))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}