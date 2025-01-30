package pubsub

import (
	rabbitmq "allchat-message-service/internal/adapter/messageBroker/rabbitMQ/amqp"
	"fmt"

	"github.com/streadway/amqp"
)

type Publisher struct {
	conn *rabbitmq.AmqpRabbitClient
}

func NewPublisher(conn *rabbitmq.AmqpRabbitClient) *Publisher {
	return &Publisher{conn: conn}
}

func (a *Publisher) Publish(exchange, routingKey string, body []byte) error {
	//start send queue
	fmt.Println("Publishing message to queue")
	ch, err := a.conn.CreateChannel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	if _, err := ch.QueueDeclare(
		exchange, // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	); err != nil {
		return fmt.Errorf("failed to declare a queue")
	}

	err = ch.Publish(
		"",         // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         body,
			DeliveryMode: amqp.Persistent, // Mark message as persistent
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}

	return nil
}
