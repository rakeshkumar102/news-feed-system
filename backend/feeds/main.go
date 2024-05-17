package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/feeds/services"
)

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Feeds@v1.0.0.0",
	})
}

func InitRouters(r *gin.Engine) {
	r.GET("/", rootHandler)
	services.FeedRouters(r)
}

func main() {
	r := gin.Default()

	InitRouters(r)

	log.Fatal(r.Run(":8031"))
}