package core

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
	Context context.Context
}

type QueueName string

const (
	QueueAPI       QueueName = "api.event"
	QueueWebSocket QueueName = "ws.notification"
	QueueCron      QueueName = "cron.task"
)

func InitRabbitMQ(name QueueName) *RabbitMQ {
	Conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	Channel, err := Conn.Channel()
	if err != nil {
		panic(err)
	}

	Queue, err := Channel.QueueDeclare(
		string(name),
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

func (r *RabbitMQ) PublishMessage(queue QueueName, body []byte) error {
	err := r.Channel.PublishWithContext(r.Context,
		"",
		string(queue),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	return err
}

func (r *RabbitMQ) ConsumeMessage(queue QueueName, handler func(delivery amqp.Delivery)) error {
	msgs, err := r.Channel.Consume(
		string(queue),
		"gym-partner-consumer",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			handler(d)
		}
	}()

	return nil
}

func (r *RabbitMQ) Close() {
	if r.Channel != nil {
		_ = r.Channel.Close()
	}
	if r.Conn != nil {
		_ = r.Conn.Close()
	}
}
