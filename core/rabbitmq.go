package core

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
	Context context.Context
}

func NewRabbitMQ() *RabbitMQ {
	Conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	Channel, err := Conn.Channel()
	if err != nil {
		panic(err)
	}

	Queue, err := Channel.QueueDeclare(
		"test_queue",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return &RabbitMQ{
		Conn:    Conn,
		Channel: Channel,
		Queue:   Queue,
		Context: ctx,
	}
}

func (r *RabbitMQ) PublishMessage(body string) error {
	err := r.Channel.PublishWithContext(r.Context,
		"",
		r.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	return err
}

func (r *RabbitMQ) ConsumeMessage() {
	msgs, err := r.Channel.Consume(
		r.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
}
