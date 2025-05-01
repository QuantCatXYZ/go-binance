package main

import (
	"context"
	"fmt"
	"github.com/quantcatxyz/go-binance/v2"
	"log"
)

func main() {
	// Ticker()
	// Ohlcv()
	// SpotOrder()
	// FuturesOrder()
	//WalletBalance()

	GetExchangeInfo()
}

func GetExchangeInfo() {
	apiKey := ""
	secret := ""
	client := binance.NewClient(apiKey, secret)
	info, err := client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Exchange Info:")
	for _, symbol := range info.Symbols {
		fmt.Printf("Symbol: %s, Status: %s\n", symbol.Symbol, symbol.Status)
	}

}
