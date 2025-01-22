package blockchain

import (
	"time"

	"github.com/trotelalexandre/proto/common"
)

type Block struct {
	Index        int
	Timestamp    time.Time
	Data 		 []Transaction
	PreviousHash string
	Hash         string
}

func CreateGenesisBlock(coin Coin) *Block {
	genesisBlock := &Block{
		Index:        0,
		Timestamp:    time.Now(),
		Data:         []Transaction{},
		PreviousHash: "0",
		Hash:         "",
	}
	genesisBlock.Hash = common.HashData(genesisBlock.ToBlockData())
	return genesisBlock
}

func (b *Block) ToBlockData() []byte {
	data := []byte{}
	for _, transaction := range b.Data {
		data = append(data, transaction.ToTransactionData()...)
	}
	return append(data, []byte(b.PreviousHash)...)
}