package main

import (
	"os"
)

var outputDir = os.Getenv("U2P_OUTPUT_DIR")
var program = os.Getenv("U2P_PROGRAM")
var port = os.Getenv("U2P_SERVICE_PORT")

func main() {
	action := "no-action"
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	app := new(Application)
	app.SetQueueProvider(NewRabbitMqProvider())
	app.Run(action)
}
