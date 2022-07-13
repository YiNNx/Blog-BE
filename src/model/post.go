package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	collectionName string `collection:"post" `

	ObjectID primitive.ObjectID `bson:"_id,omitempty"`
	Status   int                `bson:"status,omitempty"` //0-Public | 1-Private | 2-Script
	Type     int                `bson:"type,omitempty"`   //0-PlainText | 1-Markdown | 2-HTML
	Title    string             `bson:"title,omitempty"`
	Time     int64              `bson:"time,omitempty"`

	Excerpt string   `bson:"excerpt,omitempty"`
	Content string   `bson:"content,omitempty"`
	Tags    []string `bson:"tags,omitempty"`

	Views    int `bson:"views,omitempty"`
	Likes    int `bson:"likes,omitempty"`
	Comments int `bson:"comments,omitempty"`

	IsDeleted bool `bson:"is_deleted,omitempty"`
}

