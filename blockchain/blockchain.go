package blockchain

import (
	"blockchain/utils"
	"sync"
	"time"
)

type Blockchain struct {
	Blocks []*Block
	Mux    sync.Mutex
	Reward int
}

func (bc *Blockchain) AddBlock(transactions []Transaction, miner string) {
	bc.Mux.Lock()
	defer bc.Mux.Unlock()
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	rewardTransaction := Transaction{
		Sender:    "System",
		Recipient: miner,
		Amount:    bc.Reward,
	}
	transactions = append(transactions, rewardTransaction)

	newBlock := &Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
		Reward:       bc.Reward,
	}

	newBlock.Hash = utils.CalculateHash(newBlock.ToBlockData())
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}