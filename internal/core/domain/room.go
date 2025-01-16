package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}
