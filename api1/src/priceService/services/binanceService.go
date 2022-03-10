package services

import (
	"context"
	"log"

	"github.com/adshao/go-binance/v2"
)

type price struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func GetPrices(ctx context.Context) []*binance.SymbolPrice {
	log.Println("Getting prices from Binance.")
	var (
		apiKey    = "your api key"
		secretKey = "your secret key"
	)

	client := binance.NewClient(apiKey, secretKey)
	prices, err := client.NewListPricesService().Do(ctx)

	if err != nil {
		log.Fatalln("An error occurred calling Binance.", err)
		return nil
	}

	return prices
}
