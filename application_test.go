package main

import (
	"testing"
	"microservice-task/queue/providers"
	"microservice-task/queue"
	"fmt"
)

var app *Application

func init(){
	app = new(Application)
	provider := new(queue_providers.MockRabbitMqProvider)
	app.SetQueueProvider(provider)
}


func TestCreateJobWithValidUrl(t *testing.T) {
	url := "http://google.com"
	job, err := app.createJob(url)
	if err != nil {
		t.Error("Expected no error, got "+err.Error())
	}

	if job.Url != url{
		t.Error("url passed to createJob does not match job.Url")
	}
}


func TestCreateJobWithInValidUrl(t *testing.T) {
	url := "http://google.coom"
	_, err := app.createJob(url)
	if err == nil {
		t.Error("Expected error")
	}
}

type MockConsumer struct {

}

func (c MockConsumer) Process(msg *queue.Message) bool {
	fmt.Println(msg.Body)
	return true
}

func TestPublish(t *testing.T){
	app.QueueProvider.Publish("oke")
	consumer := new(MockConsumer)
	app.QueueProvider.Consume(consumer)
}
