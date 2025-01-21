package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func CalculateHash(block Block) string {
	hashInput := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.PrevHash, block.Transactions)
	hash := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hash[:])
}