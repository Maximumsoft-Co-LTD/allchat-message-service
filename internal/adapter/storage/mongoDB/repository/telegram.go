package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	rawDataCollection = "raw_data"
)

type TelegramRepository struct {
	DB *mongo.Database
}

func NewTelegramRepository(DB *mongo.Database) *TelegramRepository {
	return &TelegramRepository{
		DB,
	}
}

func (r *TelegramRepository) InsertWebhookRawData(data any) error {
	// Insert webhook raw data
	fmt.Println("repo create webhook telegram data")
	c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := r.DB.Collection(rawDataCollection).InsertOne(c, data)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}
