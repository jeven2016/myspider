package model

import (
	"core/pkg/common/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Site struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Enabled     bool               `bson:"enabled" json:"enabled"`
	HomeUrl     string             `bson:"homeUrl" json:"homeUrl"`
	Order       int                `bson:"order" json:"order"`
	Description string             `bson:"description" json:"description"`
	CreatedTime time.Time          `bson:"createdTime" json:"createdTime"`
	Jobs        []SiteJob          `bson:"jobs" json:"jobs"`
}

func (s *Site) BaseUrl() (string, error) {
	if len(s.HomeUrl) == 0 {
		return "", nil
	}
	return utils.GetBaseUrl(s.HomeUrl)
}

type SiteCatalog struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	ParentId    primitive.ObjectID `bson:"parentId" json:"parentId"`
	Name        string             `bson:"name" json:"name"`
	Order       int32              `bson:"order" json:"order"`
	CreatedDate time.Time          `bson:"createdDate" json:"createdDate"`
	UpdatedDate time.Time          `bson:"updatedDate" json:"updatedDate"`
	DeletedDate time.Time          `bson:"deletedDate" json:"deletedDate"`
}

type SiteJob struct {
	Enabled     bool   `bson:"enabled" json:"enabled"`
	Name        string `bson:"name" json:"name"`
	Type        string `bson:"type" json:"type"`
	Parser      string `bson:"parser" json:"parser"`
	SubPageUrl  string `bson:"subPageUrl" json:"subPageUrl"`
	Source      string `bson:"source" json:"source"`
	Destination string `bson:"destination" json:"destination"`
}
