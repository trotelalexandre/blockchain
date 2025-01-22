package blockchain

import (
	"fmt"
	"time"

	"github.com/trotelalexandre/proto/utils"
)

type Block struct {
	Index        int
	Timestamp    int64
	Hash         string
	Transactions []Transaction
	PrevHash     string
	Reward	     int
}

func CreateGenesisBlock(coin Coin) *Block {
	genesisBlock := &Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Transactions: []Transaction{},
		PrevHash:     "0",
		Reward:       ToDecimals(100, coin),
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