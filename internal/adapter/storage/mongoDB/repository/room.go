package repository

import (
	"allchat-message-service/internal/core/domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	roomCollection = "room"
)

type RoomRepository struct {
	DB *mongo.Database
}

func NewRoomRepository(DB *mongo.Database) *RoomRepository {
	return &RoomRepository{
		DB,
	}
}

func (r *RoomRepository) CreateRoom(c context.Context, room domain.Room) error {
	// Create room
	// ctx, cancel := db.InitContext()
	// defer cancel()
	fmt.Println("repo create room")
	room.ID = primitive.NewObjectID()
	_, err := r.DB.Collection(roomCollection).InsertOne(c, room)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	return nil
}
