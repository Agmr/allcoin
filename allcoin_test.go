package allcoin

import (
	"testing"
)

const TestFile string = "coins.json"

func TestNewFromAPI(t *testing.T) {
	coins, err := NewFromAPI()

	if err != nil {
		t.Fatalf("Error occured while parsing coins: %v\n", err)
	}

	if len(coins) <= 1 {
		t.Fatal("There are should be more then 1 coin")
	}

	err = WriteToFile(coins, TestFile)

	if err != nil {
		t.Fatalf("Error while writing to a file: %v\n", err)
	}

}

func TestWriteToBuildFromFile(t *testing.T) {
	coins, err := NewFromJSON(TestFile)
	if err != nil {
		t.Fatalf("Error while building coins from JSON: %v\n", err)
	}

	if !coins.Exist("BTC") || !coins.Exist("LTC") || !coins.Exist("NEO") {
		t.Fatal("BTC, LTC, NEO are existing coins")
	}
}
