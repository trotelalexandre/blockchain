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

	err := bc.SaveToFile("protochain.json")
	if err != nil {
		fmt.Println("Error saving blockchain to file:", err)
	}

	node := node.Node{Address: "localhost:8080", Peers: []string{"localhost:8081"}, Blockchain: bc}
	go node.StartNode(node, bc)
	node.ConnectToPeers(node)
}
