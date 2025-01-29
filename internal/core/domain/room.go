package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Type         string             `bson:"type" json:"type"`
	Datetime     time.Time          `bson:"datetime" json:"datetime"`
	LastDate     time.Time          `bson:"last_date" json:"last_date"`
	Prefix       string             `bson:"prefix" json:"prefix"`
	BusinessCode string             `bson:"business_code" json:"business_code"`
	Platform     string             `bson:"platform" json:"platform"`
	PageID       string             `bson:"page_id" json:"page_id"`
	PageName     string             `bson:"page_name" json:"page_name"`
	PageImg      string             `bson:"page_img" json:"page_img"`
	IsRead       bool               `bson:"is_read" json:"is_read"`
	IsReadUser   []string           `bson:"is_read_user" json:"is_read_user"`
	Tags         []TagsChat         `bson:"tags" json:"tags"`
}

type TagsChat struct {
	ID   primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name string             `bson:"name" json:"name"`
}
