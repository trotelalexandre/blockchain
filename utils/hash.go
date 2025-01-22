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

func GenerateUniqueAddress() string {
    randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
    randomBytes := make([]byte, 32)
    for i := range randomBytes {
        randomBytes[i] = byte(randSource.Intn(256))
    }
    return hex.EncodeToString(randomBytes)
}