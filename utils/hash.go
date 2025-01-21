package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type BlockData struct {
	Index        int
	Timestamp    string
	PrevHash     string
	Transactions string
	Reward       int
}

func CalculateHash(data BlockData) string {
	hashInput := fmt.Sprintf("%d%s%s%s%d", data.Index, data.Timestamp, data.PrevHash, data.Transactions, data.Reward)
	hash := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hash[:])
}
