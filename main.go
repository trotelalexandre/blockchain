package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/trotelalexandre/proto/blockchain"
	"github.com/trotelalexandre/proto/handlers"
)

func main() {
	coin := blockchain.Coin{
		Name: "Proto",
		Symbol: "PRT",
		Decimals: 18,
		TotalSupply: 10000000 * int(math.Pow10(18)), // 10,000,000 PRT
	}

	genesisBlock := blockchain.CreateGenesisBlock(coin)
	bc := &blockchain.Blockchain{
		Blocks: []*blockchain.Block{genesisBlock},
		Reward: 50,
		Name: "Proto",
		Coin: coin,
	}

	http.HandleFunc("/blockchain", handlers.GetBlockchain(bc))
	http.HandleFunc("/transaction", handlers.SendTransaction(bc))
	http.HandleFunc("/add-wallet", handlers.AddWallet(bc))

	fmt.Println("Proto is running on :8080")
	http.ListenAndServe(":8080", nil)
}