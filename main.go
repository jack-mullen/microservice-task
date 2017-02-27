package main

import (
	"os"
	"fmt"
	"microservice-task/queue"
	"strconv"
	"microservice-task/utils"
)

var QueueProvider queue.Provider

var outputDir = os.Getenv("U2P_OUTPUT_DIR")
var program = os.Getenv("U2P_PROGRAM")
var port = os.Getenv("U2P_SERVICE_PORT")

func main() {
	action := "no-action"
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	QueueProvider = NewRabbitMqProvider()
	defer QueueProvider.Cleanup()

	switch action{
	case "server":
		portInt, err := strconv.Atoi(port)
		utils.FailOnError(err, "Port number must be numeric")
		startServer(portInt)
	case "consume":
		consumer := NewUrl2PdfConsumer(outputDir, program)
		QueueProvider.Consume(consumer)
	default:
		fmt.Println("Valid options are: [server | consume]")
	}
}
