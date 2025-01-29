package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		// App   *App
		// Token *Token
		Redis    *Redis
		DB       *DB
		HTTP     *HTTP
		RabbitMQ *RabbitMQ
		AWS      *AWS
	}
	// // App contains all the environment variables for the application
	// App struct {
	// 	Name string
	// 	Env  string
	// }
	// // Token contains all the environment variables for the token service
	// Token struct {
	// 	Duration string
	// }
	// Redis contains all the environment variables for the cache service
	Redis struct {
		Addr     string
		Password string
	}
	// Database contains all the environment variables for the database
	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}
	// HTTP contains all the environment variables for the http server
	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
	}
	RabbitMQ struct {
		Addr         string
		User         string
		Password     string
		MqttProtocal string
		AmqpProtocal string
		MqttPort     string
		AmqpPort     string
	}

	AWS struct {
		AwsS3Region        string
		AwsS3Bucket        string
		AwsAccessKeyID     string
		AwsSecretAccessKey string
		DomainUrlS3        string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	// app := &App{
	// 	Name: os.Getenv("APP_NAME"),
	// 	Env:  os.Getenv("APP_ENV"),
	// }

	// token := &Token{
	// 	Duration: os.Getenv("TOKEN_DURATION"),
	// }

	redis := &Redis{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	http := &HTTP{
		Env:            os.Getenv("APP_ENV"),
		URL:            os.Getenv("HTTP_URL"),
		Port:           os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	rabbitMQ := &RabbitMQ{
		Addr:         os.Getenv("RABBITMQ_ADDR"),
		User:         os.Getenv("RABBITMQ_USER"),
		Password:     os.Getenv("RABBITMQ_PASSWORD"),
		MqttProtocal: os.Getenv("RABBITMQ_MQTT_PROTOCOL"),
		AmqpProtocal: os.Getenv("RABBITMQ_AMQP_PROTOCOL"),
		MqttPort:     os.Getenv("RABBITMQ_MQTT_PORT"),
		AmqpPort:     os.Getenv("RABBITMQ_AMQP_PORT"),
	}

	aws := &AWS{
		AwsS3Region:        os.Getenv("AWS_S3_REGION"),
		AwsS3Bucket:        os.Getenv("AWS_S3_BUCKET"),
		AwsAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		DomainUrlS3:        os.Getenv("DOMAIN_URL_S3"),
	}

	return &Container{
		// app,
		// token,
		redis,
		db,
		http,
		rabbitMQ,
		aws,
	}, nil
}
