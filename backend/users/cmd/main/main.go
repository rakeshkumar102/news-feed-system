package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/users/lib"
	"github.com/pranay999000/users/services"
)

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Users@v1.0.0.0",
	})
}

func InitRouters(r *gin.Engine) {
	r.GET("/", rootHandler)
	services.AuthRouters(r)

}

func main() {
	r := gin.Default()

	InitRouters(r)

	connection, channel := lib.SetUpRabbitMQConnectionChannel()

	defer connection.Close()
	defer channel.Close()

	log.Fatal(r.Run(":8001"))
}