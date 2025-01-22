package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

type BlockData struct {
	Index        int
	Timestamp    int64
	PrevHash     string
	Transactions string
	Reward       int
}

type TransactionData struct {
	Sender    string
	Recipient string
	Amount    int
}

func CalculateHash(data BlockData) string {
	hashInput := fmt.Sprintf("%d%d%s%s%d", data.Index, data.Timestamp, data.PrevHash, data.Transactions, data.Reward)
	hash := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hash[:])
}

func CalculateTransactionHash(data TransactionData) string {
	hashInput := fmt.Sprintf("%s%s%d", data.Sender, data.Recipient, data.Amount)
	hash := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hash[:])
}

func GenerateUniqueAddress() string {
    randSource := rand.New(rand.NewSource(time.Now().Unix()))
    randomBytes := make([]byte, 32)
    for i := range randomBytes {
        randomBytes[i] = byte(randSource.Intn(256))
    }
    return hex.EncodeToString(randomBytes)
}