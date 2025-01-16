package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type TelegramRepository struct {
	DB *mongo.Database
}

func NewTelegramRepository(DB *mongo.Database) *TelegramRepository {
	return &TelegramRepository{
		DB,
	}
}
