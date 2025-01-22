package blockchain

import (
	"crypto/ecdsa"

	"github.com/trotelalexandre/proto/common"
)

type Account struct {
	Address string
	Balance int64
}

func CreateAccount(privateKey *ecdsa.PrivateKey) *Account {
	publicKey := common.GetPublicKeyFromPrivateKey(privateKey)
	return &Account{
		Address: common.GenerateAddress(publicKey),
		Balance: 0,
	}
}