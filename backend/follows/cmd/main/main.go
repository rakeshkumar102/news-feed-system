package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/follows/functions"
	"github.com/pranay999000/follows/lib"
	"github.com/pranay999000/follows/services"
)

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Follows@v1.0.0.0",
	})
}

func InitRouters(r *gin.Engine) {
	r.GET("/", rootHandler)
	services.FollowRouters(r)
}

func main() {
	r := gin.Default()

	InitRouters(r)
	
	connection, channel := lib.SetUpRabbitMQConnectionChannel()

	defer connection.Close()
	defer channel.Close()

	requestQueue, err := channel.QueueDeclare(
		"hello",
		true,
		false,
		false,
		false,
		nil,
	)

	functions.FailOnError(err, "Failed to register a queue")

	request, err := channel.Consume(
		requestQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	functions.FailOnError(err, "Failed to register a listener in queue")

	go func() {
		for d := range request {
			fmt.Printf("Recieved Message: %s\n", d.Body)
			_, err := functions.CreateVertex(string(d.Body))

			functions.FailOnError(err, "Unable to create vertex")
			d.Ack(false)
		}
	}()

	serverErr := r.Run(":8002")
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}