package main

import (
	"allchat-message-service/internal/adapter/config"
	"allchat-message-service/internal/adapter/handler/http"
	rabbitmq "allchat-message-service/internal/adapter/rabbitMQ/amqp"
	"allchat-message-service/internal/adapter/rabbitMQ/amqp/pubsub"
	"allchat-message-service/internal/adapter/storage/mongoDB"
	"allchat-message-service/internal/adapter/storage/mongoDB/repository"
	"allchat-message-service/internal/adapter/storage/redis"
	"allchat-message-service/internal/core/service"
	"context"
	"fmt"
	"log"

	"github.com/fatih/color"
)

func main() {
	color.Green("Server starting...")
	// _ = godotenv.Load()

	// Load environment variables
	config, err := config.New()
	if err != nil {
		log.Fatalf("Error loading environment variables", err)
	}

	// database
	resource, err := mongoDB.New(config.DB)
	if err != nil {
		log.Fatalf("Connection database failure, Please check connection", err)
	}
	defer resource.Close()

	// Init cache service
	ctx := context.Background()
	cache, err := redis.New(ctx, config.Redis)
	if err != nil {
		log.Fatalf("Error initializing cache connection", err)
	}
	defer cache.Close()

	// // Init rabbitMQ service
	// mqtt, err := rabbitmq.NewMQTT(config.RabbitMQ)
	// if err != nil {
	// 	log.Fatalf("Error initializing rabbitMQ connection", err)
	// }
	// fmt.Println("rabbitMQ mqtt ", mqtt)

	// Init amqp service
	amqp, err := rabbitmq.NewAMQP(config.RabbitMQ)
	if err != nil {
		log.Fatalf("Error initializing rabbitMQ connection", err)
	}
	defer amqp.Close()

	// amqp Publish
	publisher := pubsub.NewPublisher(amqp)

	// Dependency injection
	//Room
	roomRepo := repository.NewRoomRepository(resource.DB)

	// //Telegram
	telegramRepo := repository.NewTelegramRepository(resource.DB)
	telegramService := service.NewTelegramService(telegramRepo, roomRepo, cache, publisher)
	telegramHandler := http.NewTelegramHandler(telegramService)
	telegramSub := pubsub.NewTelegramSubscriber(telegramService)

	// amqp Subscribe
	subscriber := pubsub.NewSubscriber(amqp, *telegramSub)
	go func() {
		if err = subscriber.RawDataWebhookSubscribe(); err != nil {
			log.Fatalf("Error subscribing to raw_data", err)
		}
	}()

	// Router
	router, err := http.NewRouter(
		config.HTTP,
		*telegramHandler,
	)
	if err != nil {
		log.Fatalf("Error initializing router", err)
	}

	// server
	fmt.Printf("Server listening on port %s\n", config.HTTP.Port)
	listenAddr := fmt.Sprintf(":%s", config.HTTP.Port)

	err = router.Serve(listenAddr)
	if err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}
