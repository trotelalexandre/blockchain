package stats

import (
	"github.com/trotelalexandre/proto/blockchain"
)

func GetBlockCount(bc *blockchain.Blockchain) int {
	return len(bc.Blocks)
}

func GetTransactionCount(bc *blockchain.Blockchain) int {
	count := 0
	for _, block := range bc.Blocks {
		count += len(block.Transactions)
	}
	return count
}

func GetWalletCount(bc *blockchain.Blockchain) int {
	return len(bc.State)
}

func GetBlockReward(bc *blockchain.Blockchain) int {
	return bc.Reward
}

func GetAllBlocks(bc *blockchain.Blockchain) []*blockchain.Block {
	return bc.Blocks
}

func GetAllTransactions(bc *blockchain.Blockchain) []blockchain.Transaction {
	transactions := []blockchain.Transaction{}
	for _, block := range bc.Blocks {
		transactions = append(transactions, block.Transactions...)
	}
	return transactions
}

func GetWalletBalance(bc *blockchain.Blockchain, address string) int {
	return bc.State[address]
}