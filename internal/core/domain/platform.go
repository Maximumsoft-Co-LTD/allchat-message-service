package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Platform struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Datetime     time.Time          `bson:"datetime" json:"datetime"`
	Prefix       string             `bson:"prefix" json:"prefix"`
	BusinessID   string             `bson:"business_id" json:"business_id"`
	BusinessCode string             `bson:"business_code" json:"business_code"`
	Name         string             `bson:"name" json:"name"`
	AccessToken  string             `bson:"access_token" json:"access_token"`
	Secret       string             `bson:"secret" json:"secret"`
	ImgUrl       string             `bson:"img_url" json:"img_url"`
	Category     string             `bson:"category" json:"category"`
	PageID       string             `bson:"page_id" json:"page_id"`
	Platform     string             `bson:"platform" json:"platform"`
	Active       bool               `bson:"active" json:"active"`
	Details      string             `bson:"details" json:"details"`

	//USER
	UpdateAt         time.Time `bson:"update_at" json:"update_at"`
	ActionByID       string    `bson:"action_by_id" json:"action_by_id"`
	ActionByUsername string    `bson:"action_by_username" json:"action_by_username"`
	ActionByName     string    `bson:"action_by_name" json:"action_by_name"`
}

type PlatformConnection struct {
	Name     string `bson:"name" json:"name"`
	ImgUrl   string `bson:"img_url" json:"img_url"`
	Platform string `bson:"platform" json:"platform"`
}
