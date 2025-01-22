package node

import (
	"encoding/json"
	"os"

	"github.com/trotelalexandre/proto/blockchain"
)

type BlockchainWrapper struct {
    *blockchain.Blockchain
}

func (bc *BlockchainWrapper) SaveToFile(filename string) error {
    data, err := json.Marshal(bc)
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644)
}

func LoadBlockchainFromFile(filename string) (*blockchain.Blockchain, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var blockchain blockchain.Blockchain
    err = json.Unmarshal(data, &blockchain)
    if err != nil {
        return nil, err
    }
    return &blockchain, nil
}
