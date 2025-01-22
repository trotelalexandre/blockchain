package blockchain

import (
	"sync"
	"time"

	"github.com/trotelalexandre/proto/utils"
)

type Blockchain struct {
	Blocks []*Block
	Mux    sync.Mutex
	Reward int
	Name   string
	Coin   Coin
	PendingTransactions []Transaction
	State map[string]int // Wallet address to balance
}

func (bc *Blockchain) AddBlock(miner string) {
	bc.Mux.Lock()
	defer bc.Mux.Unlock()

	transactions := bc.PendingTransactions

	rewardTransaction := Transaction{
		Hash:      utils.CalculateTransactionHash(utils.TransactionData{Sender: "Proto", Recipient: miner, Amount: bc.Reward}),
		Sender:    "Proto",
		Recipient: miner,
		Amount:    bc.Reward,
	}
	transactions = append(transactions, rewardTransaction)

	prevBlock := bc.GetLastBlock()

	newBlock := &Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().UnixNano(),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
		Reward:       bc.Reward,
	}

	newBlock.Hash = utils.CalculateHash(newBlock.ToBlockData())
	bc.Blocks = append(bc.Blocks, newBlock)

	bc.PendingTransactions = []Transaction{}
}

func (bc *Blockchain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddTransaction(transaction Transaction) {
	bc.Mux.Lock()
	defer bc.Mux.Unlock()
	bc.PendingTransactions = append(bc.PendingTransactions, transaction)
}

func (bc *Blockchain) AddWallet(wallet *Wallet) {
	bc.Mux.Lock()
	defer bc.Mux.Unlock()
	bc.State[wallet.Address] = wallet.Balance
}