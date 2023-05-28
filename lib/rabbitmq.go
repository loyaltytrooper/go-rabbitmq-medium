package lib

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"go-rabbitmq-medium/constants"
	"go-rabbitmq-medium/utils"
)

var RabbitChannel *amqp091.Channel
var rabbitConn *amqp091.Connection

func SetupRabbitMq() (*amqp091.Channel, *amqp091.Connection) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", constants.USERNAME, constants.PASSWORD, constants.HOST, constants.PORT)

	conn, err := amqp091.Dial(url)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	RabbitChannel = ch
	return RabbitChannel, rabbitConn
}
