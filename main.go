package main

import (
	"fmt"

	"github.com/trotelalexandre/proto/blockchain"
	"github.com/trotelalexandre/proto/node"
)

func main() {
    coin := blockchain.Coin{Name: "ProtoCoin", Symbol: "PRT", Decimals: 18, TotalSupply: 10000000}
    genesisBlock := blockchain.CreateGenesisBlock(coin)
    bc := &blockchain.Blockchain{Blocks: []blockchain.Block{*genesisBlock}}
    bcWrapper := &node.BlockchainWrapper{Blockchain: bc}

    err := bcWrapper.SaveToFile("protochain.json")
    if err != nil {
        fmt.Println("Error saving blockchain to file:", err)
    }
}
