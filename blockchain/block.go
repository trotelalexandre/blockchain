package blockchain

import (
	"fmt"
	"time"

	"github.com/trotelalexandre/proto/utils"
)

type Block struct {
	Index        int
	Timestamp    string
	Hash         string
	Transactions []Transaction
	PrevHash     string
	Reward	     int
}

func CreateGenesisBlock() *Block {
	genesisBlock := &Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: []Transaction{},
		PrevHash:     "0",
		Reward:       100,
	}
	genesisBlock.Hash = utils.CalculateHash(genesisBlock.ToBlockData())
	return genesisBlock
}

func (b *Block) ToBlockData() utils.BlockData {
	return utils.BlockData{
		Index:        b.Index,
		Timestamp:    b.Timestamp,
		PrevHash:     b.PrevHash,
		Transactions: fmt.Sprintf("%v", b.Transactions),
		Reward:       b.Reward,
	}
}