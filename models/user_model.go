package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Links struct {
	Name string `bson:"name" json:"name"`
	URL  string `bson:"url" json:"url"`
	Icon string `bson:"icon" json:"icon"`
}

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"`
	Links     []Links            `bson:"links,omitempty" json:"links,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}