package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	// Continued 连载中；Finished 完结
	Continued = iota
	Finished
)

type NovelState int

// Novel 小说
type Novel struct {
	Id              primitive.ObjectID `bson:"_id" json:"id"`                      //ID
	Order           int32              `bson:"order" json:"order"`                 //排序
	SiteCatalogId   primitive.ObjectID `bson:"siteCatalogId" json:"siteCatalogId"` //目录ID
	SiteCatalogName string             `bson:"parentId" json:"parentId"`           //目录名
	Name            string             `bson:"name" json:"name"`                   //名称
	PicUrl          string             `bson:"picUrl" json:"picUrl"`               //图片地址
	AuthorId        primitive.ObjectID `bson:"authorId" json:"authorId"`           //作者URL
	Author          string             `bson:"author" json:"author"`               //作者姓名
	Status          NovelState         `bson:"status" json:"status"`               //状态
	Description     string             `bson:"description" json:"description"`     //简要描述

	Hot              bool `bson:"hot" json:"hot"`                             //是否热门
	VipReward        bool `bson:"vipReward" json:"vipReward"`                 //vip打赏
	VipUpdate        bool `bson:"vipUpdate" json:"vipUpdate"`                 //vip更新
	Original         bool `bson:"original" json:"original"`                   //是否原创
	Recommend        bool `bson:"recommend" json:"recommend"`                 //是否推荐
	RecommendToday   bool `bson:"recommendToday" json:"RecommendToday"`       //是否今日推荐
	RecommendCollect bool `bson:"recommend_collect" json:"recommend_collect"` //推荐收藏

	CreatedDate time.Time `bson:"createdDate" json:"createdDate"` //创建时间
	UpdatedDate time.Time `bson:"updatedDate" json:"updatedDate"` //更新时间
}

// NovelStats 小说统计信息
type NovelStats struct {
	CollectNum       uint32 `bson:"collectNum" json:"createdDate"`            //收藏次数
	Views            uint32 `bson:"views" json:"views"`                       //阅读次数
	RecNum           uint32 `bson:"recNum" json:"recNum"`                     //推荐数
	TextNum          uint32 `bson:"textNum" json:"textNum"`                   //总字数
	ChapterNum       uint32 `bson:"chapterNum" json:"chapterNum"`             //章节数
	ChapterUpdatedAt uint32 `bson:"chapterUpdatedAt" json:"chapterUpdatedAt"` //最新章节更新时间
	ChapterId        string `bson:"chapterId" json:"chapterId"`               //最新章节
	ChapterTitle     string `bson:"chapterTitle" json:"chapterTitle"`         //最新章节标题
}
