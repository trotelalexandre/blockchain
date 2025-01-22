package main

import (
	"fmt"
	"log"
	"os"

	"github.com/trotelalexandre/proto/blockchain"
	"github.com/trotelalexandre/proto/config"
	"github.com/trotelalexandre/proto/node"
)

func main() {
	coin := blockchain.Coin{Name: "ProtoCoin", Symbol: "PRT", Decimals: 18, TotalSupply: 10000000}
	var bc *blockchain.Blockchain

	_, err := os.Stat("protochain.json")
	if err == nil {
		bc, err = blockchain.LoadBlockchainFromFile("protochain.json")
		if err != nil {
			fmt.Println("Error loading Protochain from file:", err)
			return
		}
	} else {
		genesisBlock := blockchain.CreateGenesisBlock(coin)
		bc = &blockchain.Blockchain{Blocks: []blockchain.Block{*genesisBlock}}

		err := bc.SaveToFile("protochain.json")
		if err != nil {
			fmt.Println("Error saving Protochain to file:", err)
			return
		}
	}
	log.Println("Protochain loaded successfully")

	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	node := node.Node{
		Config:     *config,
		Blockchain: bc,
	}

	go node.StartNode()

	go node.SyncBlockchain()

	select {}
}
