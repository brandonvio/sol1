package main

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

var prices = []price{
	{ID: "1", Title: "Nevermind", Artist: "Nirvana", Price: 56.99},
	{ID: "2", Title: "Ten", Artist: "Pearl Jam", Price: 17.99},
	{ID: "3", Title: "Life Through This", Artist: "Hole", Price: 39.99},
	{ID: "4", Title: "Welcome to the Jungle", Artist: "Guns and Roses", Price: 71.99},
}

func getPrices(ctx context.Context) []price {
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

	return nil
}
