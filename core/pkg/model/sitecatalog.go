package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SiteCatalog struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	ParentId    primitive.ObjectID `bson:"parentId" json:"parentId"`
	Name        string             `bson:"name" json:"name"`
	Order       int32              `bson:"order" json:"order"`
	CreatedDate time.Time          `bson:"createdDate" json:"createdDate"`
	UpdatedDate time.Time          `bson:"updatedDate" json:"updatedDate"`
	DeletedDate time.Time          `bson:"deletedDate" json:"deletedDate"`
}
