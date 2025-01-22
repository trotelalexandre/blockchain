package blockchain

import "github.com/trotelalexandre/proto/utils"

type Transaction struct {
	Hash	  string
	Sender    string
	Recipient string
	Amount    int
}

func GenerateTransactionHash(sender, recipient string, amount int) string {
	return utils.CalculateTransactionHash(ToTransactionData(sender, recipient, amount))
}

func ToTransactionData(sender, recipient string, amount int) utils.TransactionData {
	return utils.TransactionData{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
}