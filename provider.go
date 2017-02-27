package main

import (
	"microservice-task/queue/providers"
	"github.com/streadway/amqp"
	"fmt"
	"microservice-task/utils"
	"os"
	"strconv"
)

var queueName = os.Getenv("U2P_QUEUENAME")
var rmqUsername = os.Getenv("U2P_RABBIT_USER")
var rmqPassword = os.Getenv("U2P_RABBIT_PASSWORD")
var rmqHost = os.Getenv("U2P_RABBIT_HOST")
var rmqPort = os.Getenv("U2P_RABBIT_PORT")

func NewRabbitMqProvider() *queue_providers.RabbitMqProvider{

	rmqPortInt, err := strconv.Atoi(rmqPort)
	utils.FailOnError(err, "Port number must be numeric")

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", rmqUsername, rmqPassword, rmqHost, rmqPortInt))
	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		queueName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil, // arguments
	)
	utils.FailOnError(err, "Failed to create queue")

	return &queue_providers.RabbitMqProvider{
		Channel: ch,
		Queue: q,
		Conn: conn,
	}
}
