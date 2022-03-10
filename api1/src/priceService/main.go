package main

import (
	"api1/src/priceService/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("vim-go")
	router := gin.Default()
	router.Use(services.CORSMiddleware())
	router.GET("/prices", apiGetPrices)
	router.Run()
}

func apiGetPrices(c *gin.Context) {
	_prices := services.GetPrices(c)
	c.IndentedJSON(http.StatusOK, _prices)
}
