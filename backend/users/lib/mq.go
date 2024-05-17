package lib

import (
	"github.com/pranay999000/users/configs"
	"github.com/pranay999000/users/utils"
	"github.com/rabbitmq/amqp091-go"
)

var RabbitChannel *amqp091.Channel
var RabbitConn *amqp091.Connection

func SetUpRabbitMQConnectionChannel() (*amqp091.Connection, *amqp091.Channel) {

	url, err := configs.EnvMap("mq")

	utils.FailOnError(err, "MQ url not found")

	conn, err := amqp091.Dial(url)

	utils.FailOnError(err, "Failed to connect to MQ")

	ch, err := conn.Channel()

	utils.FailOnError(err, "Failed to open a channel")

	RabbitChannel = ch
	RabbitConn = conn

	return RabbitConn, RabbitChannel

}