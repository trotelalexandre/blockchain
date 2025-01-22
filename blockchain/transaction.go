package blockchain

import (
	"fmt"

	"github.com/trotelalexandre/proto/common"
)

type Transaction struct {
	Hash	  string
	Sender    string
	Recipient string
	Value     int64
}

func GenerateTransactionHash(transaction Transaction) string {
	data := transaction.ToTransactionData()
	return common.HashData(data)
}

func (t *Transaction) ToTransactionData() []byte {
	return []byte(t.Sender + t.Recipient + fmt.Sprintf("%d", t.Value))
}