// Methods to alter our structure manually or change coins are needed because
// some exchanges are using modified name of coins in a symbols to keep them
// of a certain length

// This source file holds all the methods relative to Coins struct
package allcoin

import (
	"fmt"
	"strings"
)

type Coins map[string]CoinInfo

type CoinInfo struct {
	Symbol   string `json:Symbol`   // will become private with getters later
	CoinName string `json:CoinName` /// will become private with getters later
}

// Add symbol to the map. If it is already in the map - returns an errork
func (cs Coins) Add(symbol, coinName string) error {
	if cs.Exist(symbol) {
		return fmt.Errorf("%s already exist! Please, use Set instead!\n", symbol)
	}

	cs.Set(symbol, coinName)

	return nil
}

// Set a name for existing symbol
func (cs Coins) Set(symbol, coinName string) {
	cs[symbol] = CoinInfo{
		Symbol:   symbol,
		CoinName: coinName,
	}
}

// Remove a coin from the map
func (cs Coins) Remove(symbol string) {
	delete(cs, symbol)
}

// Try to find valid coins that compose given exchange symbol.
// As optimization, checks if length of symbol is even, and if yes,
// then tries to slice it on two equal length parts and check if both parts are
// valid coins, otherwise sends symbol to bruteforce algo which will go through
// all possible solutions
func (cs Coins) SliceSymbolOnCoins(symbol string) ([2]string, error) {
	symbol = strings.ToUpper(symbol)

	if len(symbol)%2 == 0 { // attempt to optimize

		mid := len(symbol) / 2
		if cs.Exist(symbol[mid:], symbol[:mid]) {
			return [2]string{symbol[mid:], symbol[:mid]}, nil
		}
	}

	return cs.tryAllCombinations(symbol, len(symbol)-2)
}

// Slices symbol on different chunks(coins) and tries them until it finds a valid one
// Valid pair condition - when both coins are found in a coin map.
// If valid coins are found - returns those coins, otherwise will return error
func (cs Coins) tryAllCombinations(symbol string, cut int) ([2]string, error) {
	if float64(cut) == float64(len(symbol))/2.0 {
		return cs.tryAllCombinations(symbol, cut-1)
	}

	if cut <= 2 {
		err := fmt.Errorf("Could not find coins for symbol %s", symbol)
		return [2]string{}, err
	}

	c1 := symbol[cut:]
	c2 := symbol[:cut]

	if !cs.Exist(c1, c2) {
		return cs.tryAllCombinations(symbol, cut-1)
	} else {
		return [2]string{c1, c2}, nil
	}
}

// Takes in variable number of coins and return true if all of them are valid,
// otherwise false
func (cs Coins) Exist(coins ...string) bool {
	for _, coin := range coins {
		if _, ok := cs[coin]; !ok {
			return false
		}
	}
	return true
}
