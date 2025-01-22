package blockchain

import "math"

type Coin struct {
	Name string
	Symbol string
	Decimals int
	TotalSupply int
}

func ToDecimals(amount int, coin Coin) int {
	return amount * int(math.Pow10(coin.Decimals))
}