package blockchain

import (
	"sync"
)

type Blockchain struct {
	Blocks []*Block
	mux    sync.Mutex
}

func (bc *Blockchain) AddBlock(transactions []Transaction) {
	bc.mux.Lock()
	defer bc.mux.Unlock()
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
	}
	newBlock.Hash = CalculateHash(*newBlock)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}