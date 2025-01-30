package rabbitmq

import (
	"allchat-message-service/internal/adapter/config"
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

func (a *AmqpRabbitClient) monitorConnection() {
	for {
		if a.conn.IsClosed() {
			fmt.Println("Connection lost, attempting to reconnect...")
			a.Reconnect()
		}
		time.Sleep(3 * time.Second) // Check connection status every 3 seconds
	}
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

func (a *AmqpRabbitClient) CreateChannel() (*amqp.Channel, error) {
	ch, err := a.conn.Channel()
	if err != nil {
		if a.conn.IsClosed() {
			if err := a.Reconnect(); err != nil {
				return nil, err
			} else {
				ch, err = a.conn.Channel()
				if err != nil {
					return nil, fmt.Errorf("failed to open a channel: %v", err)
				}
			}
		} else {
			return nil, fmt.Errorf("failed to open a channel: %v", err)
		}
	}
	return ch, nil
}

func (a *AmqpRabbitClient) Close() error {
	return a.conn.Close()
}
