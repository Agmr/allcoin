package allcoin

import (
	"fmt"
	"os"
	"testing"
)

var coins Coins
var err error

const testFile string = "test.json"

func init() {
	coins, err = NewFromAPI()
}

func TestNewFromAPI(t *testing.T) {

	if err != nil {
		t.Fatalf("Error occured while parsing coins: %v\n", err)
	}

	fmt.Println(len(coins))
}

func TestWriteToBuildFromFile(t *testing.T) {

	err = WriteToFile(coins, testFile)

	defer os.Remove(testFile)

	if err != nil {
		t.Fatalf("Error while writing to a file: %v\n", err)
	}

	coins2, err := NewFromJSON(testFile)

	if err != nil {
		t.Fatalf("Error while building coins from JSON: %v\n", err)
	}

	for _, coin := range coins2 {
		if !coins.Exist(coin.Symbol) {

			t.Fatalf("Should have the same map, error on: %s - %s\n", coin.Symbol, coin.CoinName)
		}
	}
}
