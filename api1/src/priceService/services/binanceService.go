package services

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

type price struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func GetPrices(ctx context.Context) []*binance.SymbolPrice {
	var (
		apiKey    = "your api key"
		secretKey = "your secret key"
	)

	client := binance.NewClient(apiKey, secretKey)
	prices, err := client.NewListPricesService().Do(ctx)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, p := range prices {
		fmt.Println(p)
	}

	return prices
}
