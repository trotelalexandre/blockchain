package blockchain

import (
	"fmt"
	"time"

	"github.com/trotelalexandre/proto/common"
)

type Blockchain struct {
	Blocks             []Block
	State			   State
}

type State struct {
	Accounts map[string]Account
}

func (bc *Blockchain) AddBlock(transactions []Transaction) error {
    lastBlock := bc.Blocks[len(bc.Blocks)-1]
    
    var senderAccount Account
    for _, transaction := range transactions {
        var exists bool
        senderAccount, exists := bc.State.Accounts[transaction.Sender]
        if !exists {
            return fmt.Errorf("sender account not found")
        }

        if senderAccount.Balance < transaction.Value {
            return fmt.Errorf("insufficient funds for transaction from %s", transaction.Sender)
        }

        for _, block := range bc.Blocks {
            for _, tx := range block.Data {
                if tx.Sender == transaction.Sender && tx.Recipient == transaction.Recipient && tx.Value == transaction.Value {
                    return fmt.Errorf("duplicate transaction detected")
                }
            }
        }
    }

    newBlock := Block{
        Index:        lastBlock.Index + 1,
        Timestamp:    time.Now(),
        Data:         transactions,
        PreviousHash: lastBlock.Hash,
        Hash:         common.HashData(lastBlock.ToBlockData()),
    }
    
    for _, transaction := range transactions {
        senderAccount.Balance -= transaction.Value
        bc.State.Accounts[transaction.Sender] = senderAccount

        recipientAccount, exists := bc.State.Accounts[transaction.Recipient]
        if !exists {
            recipientAccount = Account{Address: transaction.Recipient}
        }
        recipientAccount.Balance += transaction.Value
        bc.State.Accounts[transaction.Recipient] = recipientAccount
    }

    bc.Blocks = append(bc.Blocks, newBlock)
    return nil
}

