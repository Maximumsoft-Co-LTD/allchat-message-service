package port

type AmqpPublisher interface {
	Publish(exchange, routingKey string, body []byte) error
}

type MqttProtocol interface {
}
