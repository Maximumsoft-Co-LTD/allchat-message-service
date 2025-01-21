package rabbitmq

import (
	"allchat-message-service/internal/adapter/config"
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

const maxRetry = 99
const retryInterval = 500 * time.Millisecond

type AmqpRabbitClient struct {
	conn    *amqp.Connection
	amqpURL string
}

func NewAMQP(config *config.RabbitMQ) (*AmqpRabbitClient, error) {

	amqpURL := fmt.Sprintf("%s://%s:%s", config.AmqpProtocal, config.Addr, config.AmqpPort)
	client := &AmqpRabbitClient{}
	client.amqpURL = amqpURL

	err := client.ConnectToRabbitMQ(amqpURL)
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ", err)
		return nil, err
	}

	go client.monitorConnection()

	return client, nil
}

func (a *AmqpRabbitClient) ConnectToRabbitMQ(amqpURL string) error {
	var err error
	a.conn, err = amqp.Dial(amqpURL)
	if err != nil {
		fmt.Printf("Failed to connect to RabbitMQ: %v\n", err)
		return err
	}
	return nil
}

func (a *AmqpRabbitClient) Reconnect() error {

	for retry := 0; retry <= maxRetry; retry++ {
		fmt.Printf("RabbitMQ retry attempt: %d\n", retry)
		if retry > 0 {
			time.Sleep(retryInterval)
		}

		if !a.conn.IsClosed() {
			fmt.Println("Connection is still active")
			return nil
		}

		err := a.ConnectToRabbitMQ(a.amqpURL)
		if err != nil {
			fmt.Println("Failed to reconnect to RabbitMQ", err)
			continue
		}

		fmt.Println("Successfully connected to RabbitMQ")
		return nil
	}

	return fmt.Errorf("failed to connect to RabbitMQ after %d retries", maxRetry)
}

func (a *AmqpRabbitClient) monitorConnection() {
	for {
		if a.conn.IsClosed() {
			fmt.Println("Connection lost, attempting to reconnect...")
			a.Reconnect()
		}
		time.Sleep(3 * time.Second) // Check connection status every 3 seconds
	}
}

func (a *AmqpRabbitClient) Publish(exchange, routingKey string, body []byte) error {
	ch, err := a.conn.Channel()
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
		exchange,   // exchange
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

type MessagePublish struct {
	Type     string
	Username string
	UserID   string
	Data     any
}

func (a *AmqpRabbitClient) RawDataWebhookSubscribe(telegramSub TelegramSubscribe) error {
	ch, err := a.conn.Channel()
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
		1,     // prefetch count
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

	go func() {
		for d := range msgs {
			var message MessagePublish
			if err := json.Unmarshal(d.Body, &message); err != nil {
				fmt.Printf("Err Received a message: %s\n", d.Body)
				fmt.Println(err)
			} else {
				fmt.Printf("Received a message type : %s username : %s\n", message.Type, message.Username)
				switch v := message.Type; v {
				case "telegram":
					telegramSub.InsertWebhookRawData(message.Data)
				}
				// case "line":
				// 	lineSub.InsertRawData(message.Data)
			}

			d.Ack(false) // Manually acknowledge the message
		}
	}()

	return nil
}

func (a *AmqpRabbitClient) Webhook(d amqp.Delivery) {
	fmt.Println("Received a message: ", string(d.Body))

}

func (a *AmqpRabbitClient) Close() error {
	return a.conn.Close()
}
