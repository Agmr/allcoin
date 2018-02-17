package allcoin

import (
	"log"
	"os"
)

type Coins map[string]CoinInfo

type CoinInfo struct {
	Symbol   string `json:Symbol`
	CoinName string `json:CoinName`
}

func (cs Coins) Exist(coin string) bool {
	if _, ok := cs[coin]; ok {
		return true
	} else {
		return false
	}
}

func (cs Coins) GetCoinsFromSymbol(symbol string) ([2]string, error) {
	// check if symbol has delimiter, delimit by delimiter if yes,
	// or check against all of the coins if yes
	return [2]string{"BTC", "LTC"}, nil
}

func (ci CoinInfo) Resembles(coin string) bool {
	if ci.Symbol == coin {
		return true
	} else {
		return false
	}
}
