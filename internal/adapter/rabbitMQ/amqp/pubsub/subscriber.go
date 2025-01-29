package pubsub

import (
	rabbitmq "allchat-message-service/internal/adapter/rabbitMQ/amqp"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type Subscriber struct {
	conn               *rabbitmq.AmqpRabbitClient
	telegramSubscriber TelegramSubscriber
}

func NewSubscriber(conn *rabbitmq.AmqpRabbitClient, telegramSubscriber TelegramSubscriber) *Subscriber {
	return &Subscriber{
		conn,
		telegramSubscriber,
	}
}

type MessagePublish struct {
	Type string
	Data any
}

func (s *Subscriber) RawDataWebhookSubscribe() error {
	ch, err := s.conn.CreateChannel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	queue := "raw_data"
	ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	ch.Qos(
		10,    // prefetch count
		0,     // prefetch size
		false, // global
	)
	msgs, err := ch.Consume(
		queue, // queue
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err)
	}
	var forever chan struct{}

	go func() {
		tmpD := amqp.Delivery{}
		// telegramMessages := []any{}
		for d := range msgs {
			var message MessagePublish
			if err := json.Unmarshal(d.Body, &message); err != nil {
				fmt.Printf("Err Received a message: %s\n", d.Body)
				fmt.Println(err)
			} else {
				fmt.Printf("Received a message type : %s\n", message.Type)

				switch v := message.Type; v {
				case "telegram":
					// telegramMessages = append(telegramMessages, message.Data)
					s.telegramSubscriber.InsertWebhookRawData(message)
				}
				// case "line":
				// 	lineSub.InsertRawData(message.Data)
			}
			tmpD = d
		}
		// s.telegramSubscriber.InsertWebhookRawData(telegramMessages)
		tmpD.Ack(true) // Manually acknowledge the message
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	fmt.Println("end")
	return nil
}
