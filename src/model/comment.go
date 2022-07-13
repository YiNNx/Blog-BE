package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	collectionName string `collection:"comment" `

	ObjectID primitive.ObjectID `bson:"_id,omitempty"`

	From    string `bson:"from,omitempty"` //0-Public | 1-Private | 2-Script
	Email   string `bson:"email,omitempty"`
	FromUrl string `bson:"from_url,omitempty"`
	To      string `bson:"to,omitempty"` //0-PlainText | 1-Markdown | 2-HTML

	Time    int64  `bson:"time,omitempty"`
	Content string `bson:"content,omitempty"`

	IsDeleted bool `bson:"is_deleted"`
}
