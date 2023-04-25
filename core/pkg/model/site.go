package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Site struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Order       int                `bson:"order" json:"order"`
	Description string             `bson:"description" json:"description"`
	CreatedTime time.Time          `bson:"createdTime" json:"createdTime"`
}
