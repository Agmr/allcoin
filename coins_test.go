package allcoin

import (
	"testing"
)

// Was written in a rush, needs to be broken on smaller tests
func TestCoins(t *testing.T) {
	allcoin, err := NewFromJSON(TestFile)

	if err != nil {
		t.Fatalf("Error occured while parsing coins: %v\n", err)
	}

	err = allcoin.Add("ShitCoin", "SHIT COIN")

	if err != nil {
		t.Fatal(err)
	}

	if !allcoin.Exist("ShitCoin") {
		t.Fatal("ShitCoin was added in previous step, but it cannot find it")
	}

	err = allcoin.Add("ShitCoin", "SC")

	if err == nil {
		t.Fatal("Error should be not nil as we already added shitcoin")
	}

	allcoin.Remove("ShitCoin")
	if allcoin.Exist("ShitCoin") {
		t.Fatal("ShitCoin should not exist, we removed it during a previous step")
	}

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
