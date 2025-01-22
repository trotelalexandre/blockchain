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
		count += len(block.Data)
	}
	return count
}

func GetWalletCount(bc *blockchain.Blockchain) int {
	return len(bc.State.Accounts)
}

func GetAllBlocks(bc *blockchain.Blockchain) []blockchain.Block {
	return bc.Blocks
}

func GetAllTransactions(bc *blockchain.Blockchain) []blockchain.Transaction {
	transactions := []blockchain.Transaction{}
	for _, block := range bc.Blocks {
		transactions = append(transactions, block.Data...)
	}
	return transactions
}

func GetWalletBalance(bc *blockchain.Blockchain, address string) int64 {
	account, exists := bc.State.Accounts[address]
	if !exists {
		return 0
	}
	return account.Balance
}

func GetBlockByIndex(bc *blockchain.Blockchain, index int) *blockchain.Block {
	for _, block := range bc.Blocks {
		if block.Index == index {
			return &block
		}
	}
	return nil
}

func GetTransactionByHash(bc *blockchain.Blockchain, hash string) *blockchain.Transaction {
	for _, block := range bc.Blocks {
		for _, transaction := range block.Data {
			if transaction.Hash == hash {
				return &transaction
			}
		}
	}
	return nil
}

func GetWalletByAddress(bc *blockchain.Blockchain, address string) *blockchain.Account {
	account, exists := bc.State.Accounts[address]
	if !exists {
		return nil
	}
	return &account
}

func GetBlockByHash(bc *blockchain.Blockchain, hash string) *blockchain.Block {
	for _, block := range bc.Blocks {
		if block.Hash == hash {
			return &block
		}
	}
	return nil
}