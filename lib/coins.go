package allcoin

import (
	"fmt"
	"strings"
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

func (cs Coins) Add(symbol, coinName string) error {
	if cs.Exist(symbol) {
		return fmt.Errorf("%s already exist! Please, use Set instead!\n", symbol)
	}

	cs.Set(symbol, coinName)

	return nil
}

func (cs Coins) Set(symbol, coinName string) {
	cs[symbol] = CoinInfo{
		Symbol:   symbol,
		CoinName: coinName,
	}
}

func (cs Coins) GetCoinsFromSymbol(symbol string) ([2]string, error) {
	symbol = strings.ToUpper(symbol)

	if len(symbol)%2 == 0 { // attempt to optimize

		mid := len(symbol) / 2
		if cs.ValidCoins(symbol[mid:], symbol[:mid]) {
			return [2]string{symbol[mid:], symbol[:mid]}, nil
		}
	}

	return cs.TryAllCombinations(symbol, len(symbol)-2)
}

func (cs Coins) TryAllCombinations(symbol string, cut int) ([2]string, error) {
	if cut <= 2 {
		err := fmt.Errorf("Could not find coins for symbol %s", symbol)
		return [2]string{}, err
	}

	c1 := symbol[cut:]
	c2 := symbol[:cut]

	if !cs.ValidCoins(c1, c2) {
		return cs.TryAllCombinations(symbol, cut-1)
	} else {
		return [2]string{c1, c2}, nil
	}
}

func (cs Coins) ValidCoins(coins ...string) bool {
	for _, coin := range coins {
		if _, ok := cs[coin]; !ok {
			return false
		}
	}
	return true
}

func (ci CoinInfo) Resembles(coin string) bool {
	if ci.Symbol == coin {
		return true
	} else {
		return false
	}
}
