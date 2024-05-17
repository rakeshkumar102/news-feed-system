package functions

import (
	"context"
	"time"

	"github.com/pranay999000/users/lib"
	"github.com/pranay999000/users/utils"
	"github.com/rabbitmq/amqp091-go"
)


func CreateUserNode(user_id string) {

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	ch := lib.RabbitChannel
	err := ch.PublishWithContext(ctx,
		"",
		"hello",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body: []byte(user_id),
		},
	)
	
	utils.FailOnError(err, "Failed to publish a message")


}
