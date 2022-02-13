package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("vim-go")
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AddAllowMethods("OPTIONS")
	config.AllowCredentials = true
	// router.Use(cors.New(config))
	router.Use(CORSMiddleware())
	router.GET("/albums", apiGetAlbums)
	router.Run()
}

func apiGetAlbums(c *gin.Context) {
	_albums := getAlbums()
	c.IndentedJSON(http.StatusOK, _albums)
}
