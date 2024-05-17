package lib

import (
	"github.com/pranay999000/follows/configs"
	"github.com/pranay999000/follows/functions"
	"github.com/rabbitmq/amqp091-go"
)

var RabbitChannel *amqp091.Channel
var RabbitConn *amqp091.Connection

func SetUpRabbitMQConnectionChannel() (*amqp091.Connection, *amqp091.Channel) {

	url, err := configs.EnvMap("mq")

	functions.FailOnError(err, "MQ url not found")

	conn, err := amqp091.Dial(url)

	functions.FailOnError(err, "Failed to connect to MQ")

	ch, err := conn.Channel()

	functions.FailOnError(err, "Failed to open a channel")

	RabbitChannel = ch
	RabbitConn = conn

	return RabbitConn, RabbitChannel

}