package main

import (
    "encoding/json"
    "gopkg.in/resty.v1"
    "fmt"
)

const AllCoinsUrl = "https://min-api.cryptocompare.com/data/all/coinlist"

type AllCoins struct {
    Data CoinEntry `json:Data`
}

type CoinEntry map[string]CoinInfo

type CoinInfo struct {
    Symbol string `json:Symbol`
    CoinName string `json:CoinName`
}

func GetAllCoins() (map[string]string, error) {
    resp, err := resty.R().Get(AllCoinsUrl)

    var coins map[string]string = map[string]string{}

    if err != nil {
        return coins, err
    }

    var allCoins AllCoins

    err = json.Unmarshal(resp.Body(), &allCoins)

    if err != nil {
        return coins, err
    }

    for _, v :=  range allCoins.Data {
        coins[v.Symbol] = v.CoinName
    }

    return coins, nil
}


func main () {
    coins, err := GetAllCoins()

    if err != nil {
        fmt.Printf("Error occured while parsing coins: %v\n", err)
    }

    fmt.Println(len(coins))
}
