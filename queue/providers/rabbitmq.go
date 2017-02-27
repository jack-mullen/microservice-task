package queue_providers

import (
	"github.com/streadway/amqp"
	"log"
	"encoding/json"
	"microservice-task/queue"
	"microservice-task/utils"
)

type RabbitMqProvider struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
	Conn    *amqp.Connection
}

func (p *RabbitMqProvider) Publish(data interface{}) {

	dataJson, err := json.Marshal(data)
	utils.FailOnError(err, "Json marhsall failed")

	msg := &queue.Message{
		Body: string(dataJson),
	}

	dataBytes, err := json.Marshal(msg)
	utils.FailOnError(err, "Failed to publish a message")

	err = p.Channel.Publish(
		"", // exchange
		p.Queue.Name, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        dataBytes,
		})
	utils.FailOnError(err, "Failed to publish a message")
}

func (p *RabbitMqProvider) Consume(c queue.Consumer) {
	msgs, err := p.Channel.Consume(
		p.Queue.Name, // queue
		"", // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil, // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	wait := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			msg := &queue.Message{}
			err := json.Unmarshal(d.Body, msg)
			if err == nil {
				c.Process(msg)
			}
			d.Ack(true)
		}
	}()

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-wait
}

func (p *RabbitMqProvider) Cleanup(){
	p.Channel.Close()
	p.Conn.Close()
}
