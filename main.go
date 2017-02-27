package main

import (
	"microservice-task/queue"
)

var QueueProvider queue.Provider

func main() {
	QueueProvider = NewRabbitMqProvider()
	defer QueueProvider.Cleanup()
}



