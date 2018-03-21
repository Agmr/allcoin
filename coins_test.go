package allcoin

import (
	"testing"
)

var allcoin Coins

func init() {
	allcoin, _ = NewFromJSON(TestFile)
}

func SystemTest(t *testing.T) {
	if !allcoin.Exist("BTC", "LTC", "NEO") {
		t.Fatal("Error. BTC, LTC should exist")
	}

	coins, err := allcoin.SliceSymbolOnCoins("USDTETH")
	if err != nil {
		t.Fatal(err)
	}

	for _, coin := range coins {
		if coin != "USDT" && coin != "ETH" {
			t.Fatal("Coins should be USDT and ETH")
		}
	}

	coins, err = allcoin.SliceSymbolOnCoins("btceth")
	if err != nil {
		t.Fatal(err)
	}

	for _, coin := range coins {
		if coin != "BTC" && coin != "ETH" {
			t.Fatal("Coins shoud be BTC and ETH")
		}
	}

	_, err = allcoin.SliceSymbolOnCoins("KOKOAUG")
	if err == nil {
		t.Fatalf("There are no coins for symbol KOKOAUG\n")
	}
}

func TestSetter(t *testing.T) {
	allcoin.Set("ShitCoin", "SHIT COIN")

	if !allcoin.Exist("ShitCoin") {
		t.Fatal("ShitCoin was added in previous step, but it cannot find it")
	}
}

func TestRemove(t *testing.T) {
	allcoin.Set("ShitCoin", "SC")

	allcoin.Remove("ShitCoin")

	if allcoin.Exist("ShitCoin") {
		t.Fatal("ShitCoin should not exist, we removed it during a previous step")
	}
}

func TestEncodeJSON(t *testing.T) {
	marshaled, err := allcoin.Marshal()

	if err != nil {
		t.Fatal(err)
	}

	jsonEncoded, err := allcoin.EncodeJSON()

	if err != nil {
		t.Fatal(err)
	}

	if jsonEncoded != string(marshaled) {
		t.Fatalf("Should have the same value: EncodeJSON: %s, Marshaled and casted to string: %s\n", jsonEncoded, string(marshaled))
	}
}
