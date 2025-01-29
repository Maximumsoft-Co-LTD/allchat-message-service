package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Type         string             `bson:"type" json:"type"`
	Datetime     time.Time          `bson:"datetime" json:"datetime"`
	Nickname     string             `bson:"nickname" json:"nickname"`
	FirstName    string             `bson:"first_name" json:"first_name"`
	LastName     string             `bson:"last_name" json:"last_name"`
	PhoneNumber  string             `bson:"phone_number" json:"phone_number"`
	Email        string             `bson:"email" json:"email"`
	PhofileUrl   string             `bson:"profile_url" json:"profile_url"`
	Prefix       string             `bson:"prefix" json:"prefix"`
	BusinessCode string             `bson:"business_code" json:"business_code"`
	Platform     string             `bson:"platform" json:"platform"`
	SocialID     string             `bson:"social_id" json:"social_id"`
	PageID       string             `bson:"page_id" json:"page_id"`
	RoomID       primitive.ObjectID `bson:"room_id" json:"room_id"`
}
