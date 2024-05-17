package services

import (
	"fmt"
	"log"

	"github.com/pranay999000/follows/functions"
	"github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s : %s", msg, err)
	}
}

func CreateUserNode() {
	mq, err := amqp091.Dial("amqp://mq:password@localhost:5672/")
	FailOnError(err, "Failed to open channel")
	defer mq.Close()

	ch, err := mq.Channel()
	FailOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	FailOnError(err, "Failed to register a consumer")
	
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
			_, err := functions.CreateVertex(string(d.Body))

			FailOnError(err, "Unable to create vertex")
		}
	}()

	fmt.Println(" [*] - Waiting for messages")
	<- forever
}