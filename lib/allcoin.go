package allcoin

import (
	"encoding/json"
	"gopkg.in/resty.v1"
)

const AllCoinsUrl = "https://min-api.cryptocompare.com/data/all/coinlist"

func NewFromAPI() (Coins, error) {

	type ApiCoins struct {
		Data Coins `json:Data`
	}

	var allCoins ApiCoins

	resp, err := resty.R().Get(AllCoinsUrl)

	if err != nil {
		return allCoins.Data, err
	}

	err = json.Unmarshal(resp.Body(), &allCoins)

	if err != nil {
		return allCoins.Data, err
	}

	return allCoins.Data, nil
}

/*
func NewFromJSON(fileName string) {
	_, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
}

func WriteToFile(w Writer) {
}

*/
