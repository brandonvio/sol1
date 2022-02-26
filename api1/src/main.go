package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("vim-go")
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/prices", apiGetPrices)
	router.Run()
}

func apiGetPrices(c *gin.Context) {
	ctx := context.Background()
	_prices := getPrices(ctx)
	c.IndentedJSON(http.StatusOK, _prices)
}
