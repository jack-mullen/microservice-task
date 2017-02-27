package queue_providers

import (
	"encoding/json"
	"microservice-task/queue"
	"microservice-task/utils"
)

type MockRabbitMqProvider struct {
	Queue []*queue.Message
}

func (p *MockRabbitMqProvider) Publish(data interface{}) {

	dataJson, err := json.Marshal(data)
	utils.FailOnError(err, "Json marhsall failed")

	msg := &queue.Message{
		Body: string(dataJson),
	}

	p.Queue = append(p.Queue, msg)
}

func (p *MockRabbitMqProvider) Consume(c queue.Consumer) {
	for i, msg := range p.Queue {
		c.Process(msg)

		if i+1 < len(p.Queue){
			p.Queue = p.Queue[i+1:]
		}else{
			p.Queue = nil
		}
	}

}

func (p *MockRabbitMqProvider) Cleanup() {

}
