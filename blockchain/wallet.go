package blockchain

import (
	"github.com/trotelalexandre/proto/utils"
)

type Wallet struct {
	Address string
	Balance int
}

func GenerateAddress() string {
	return "proto" + utils.GenerateUniqueAddress()
}

func AddWallet() *Wallet {
	return &Wallet{
		Address: GenerateAddress(),
		Balance: 0,
	}
}