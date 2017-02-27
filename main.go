package main

import (
	"os"
	"microservice-task/queue"
)

var QueueProvider queue.Provider

var outputDir = os.Getenv("U2P_OUTPUT_DIR")
var program = os.Getenv("U2P_PROGRAM")
var port = os.Getenv("U2P_SERVICE_PORT")

func main() {
	QueueProvider = NewRabbitMqProvider()
	defer QueueProvider.Cleanup()
}



