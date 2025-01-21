package rabbitmq

import (
	"allchat-message-service/internal/adapter/config"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttRabbitClient struct {
	client mqtt.Client
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	// Receive message
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected to MQTT broker")
	// Subscribe to the necessary topics here if needed
	sub(client)
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v\n", err)
	reconnect(client) // Attempt to reconnect
}

func NewMQTT(config *config.RabbitMQ) (*MqttRabbitClient, error) {
	opts := mqtt.NewClientOptions()
	mqttURL := fmt.Sprintf("%s://%s:%s", config.MqttProtocal, config.Addr, config.MqttPort)
	opts.AddBroker(mqttURL)
	opts.SetUsername(config.User)
	opts.SetPassword(config.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Connection failed err : ", token.Error())
		return nil, token.Error()
	}

	return &MqttRabbitClient{
		client,
	}, nil
}

func reconnect(client mqtt.Client) {
	for !client.IsConnected() {
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			fmt.Printf("Reconnection failed: %v\n", token.Error())
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("Reconnected successfully")
			break
		}
	}
}

func sub(client mqtt.Client) {
	topic := "chatAA"
	token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	})
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Subscription error: %s\n", token.Error())
		return
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)
}

func (rb *MqttRabbitClient) Reconnect() {
	reconnect(rb.client)
}
