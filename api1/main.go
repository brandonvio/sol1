package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("vim-go")
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/albums", apiGetAlbums)
	router.Run()
}

func apiGetAlbums(c *gin.Context) {
	_albums := getAlbums()
	c.IndentedJSON(http.StatusOK, _albums)
}
