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
		Name: "ProtoChain",
		Coin: blockchain.Coin{
			Name: "ProtoCoin",
			Symbol: "PRT",
		},
	}

	http.HandleFunc("/blockchain", handlers.GetBlockchain(bc))
	http.HandleFunc("/transaction", handlers.SendTransaction(bc))

	fmt.Println("ProtoChain is running on :8080")
	http.ListenAndServe(":8080", nil)
}