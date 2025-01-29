package repository

import (
	"allchat-message-service/internal/core/domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	platformCollection = "platform"
)

type PlatformRepository struct {
	DB *mongo.Database
}

func NewPlatformRepository(DB *mongo.Database) *PlatformRepository {
	return &PlatformRepository{
		DB,
	}
}

func (p *PlatformRepository) GetInfoPlatformByToken(c context.Context, Token string) (domain.Platform, error) {
	entity := domain.Platform{}
	filter := bson.M{"$or": bson.A{
		bson.M{"access_token": Token},
		bson.M{"page_id": Token},
	}}
	if err := p.DB.Collection(platformCollection).FindOne(c, filter).Decode(&entity); err != nil {
		fmt.Println("err", err)
		return entity, err
	}
	return entity, nil
}
