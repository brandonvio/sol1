package main

import (
	"api1/src/priceService/services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("vim-go")
	router := gin.Default()
	router.Use(services.CORSMiddleware())
	router.GET("/prices", apiGetPrices)
	router.Run()
}

func apiGetPrices(c *gin.Context) {
	ctx := context.Background()
	_prices := services.GetPrices(ctx)
	c.IndentedJSON(http.StatusOK, _prices)
}
