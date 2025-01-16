package mongoDB

import (
	"allchat-message-service/internal/adapter/config"
	"context"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout           = 30
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

type Resource struct {
	DB *mongo.Database
}

func New(config *config.DB) (*Resource, error) {
	_ = godotenv.Load()
	var err error
	var client *mongo.Client
	client, err = mongo.NewClient(
		options.Client().ApplyURI(config.Connection),
		options.Client().SetMinPoolSize(1),
		options.Client().SetMaxPoolSize(2),
	)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	_ = client.Connect(ctx)
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	color.Green("Connect database successfully")
	color.Green(config.Connection)
	return &Resource{DB: client.Database(config.Name)}, nil
}

func (r *Resource) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	if err := r.DB.Client().Disconnect(ctx); err != nil {
		color.Red("Close connection falure, Something wrong...")
		return
	}

	color.Cyan("Close connection successfully")
}
