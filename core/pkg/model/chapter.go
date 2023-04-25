package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	// NotStarted Continued
	NotStarted = iota
	Handled
)

type ChapterState int

// Chapter 小说章节内容
type Chapter struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`                  //ID
	NovelId     primitive.ObjectID `bson:"novelId" json:"novelId"`         //小说ID
	ChapterNo   uint32             `bson:"chapterNo" json:"chapterNo"`     //当前章节序号
	Name        string             `bson:"name" json:"name"`               //名称
	Status      ChapterState       `bson:"status" json:"status"`           //状态
	Description string             `bson:"description" json:"description"` //描述
	Link        string             `bson:"link" json:"link"`               //生成的链接
	Source      string             `bson:"source" json:"source"`           //下载URL
	CreatedDate time.Time          `bson:"createdDate" json:"createdDate"` //创建时间
	UpdatedDate time.Time          `bson:"updatedDate" json:"updatedDate"` //更新时间
}

type ChapterStas struct {
	Views     uint32 `bson:"views" json:"views"`         //阅读次数
	TextNum   uint32 `bson:"textNum" json:"textNum"`     //总字数
	Retries   uint32 `bson:"errorNum" json:"errorNum"`   //下载出错次数
	LastError string `bson:"lastError" json:"lastError"` //上次出错信息
}
