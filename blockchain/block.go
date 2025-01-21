package blockchain

import "time"

type Block struct {
	Index        int
	Timestamp    string
	Hash         string
	Transactions []Transaction
	PrevHash     string
}

func CreateGenesisBlock() *Block {
	genesisBlock := &Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: []Transaction{},
		PrevHash:     "0",
	}
	genesisBlock.Hash = CalculateHash(*genesisBlock)
	return genesisBlock
}