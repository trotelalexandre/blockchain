package main

import (
	"fmt"
	"net/http"

	"github.com/trotelalexandre/proto/blockchain"
	"github.com/trotelalexandre/proto/handlers"
)

func main() {
	genesisBlock := blockchain.CreateGenesisBlock()
	bc := &blockchain.Blockchain{
		Blocks: []*blockchain.Block{genesisBlock},
		Reward: 50,
		Name: "Proto",
		Coin: blockchain.Coin{
			Name: "Proto",
			Symbol: "PRT",
		},
	}

	http.HandleFunc("/blockchain", handlers.GetBlockchain(bc))
	http.HandleFunc("/transaction", handlers.SendTransaction(bc))

	fmt.Println("proto is running on :8080")
	http.ListenAndServe(":8080", nil)
}