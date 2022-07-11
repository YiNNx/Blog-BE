package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ObjectID primitive.ObjectID `bson:"_id"`
	Status   int                `bson:"status,omitempty"`
	Title    string             `bson:"title,omitempty"`
	Time     int64              `bson:"time,omitempty"`

	Digest  string `bson:"digest,omitempty"`
	Content string `bson:"content,omitempty"`

	TopicID    primitive.ObjectID   `bson:"topic"`
	CategoryID []primitive.ObjectID `bson:"category"`

	Views     int  `bson:"views"`
	Likes     int  `bson:"likes"`
	IsDeleted bool `bson:"is_deleted"`
}

type Tag struct {
}
