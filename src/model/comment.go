package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	collectionName string `collection:"comment" `

	ObjectID  primitive.ObjectID `bson:"_id,omitempty"`
	ParentCid string             `bson:"parent_cid,omitempty"`
	Pid       string             `bson:"post,omitempty"`

	From    string `bson:"from,omitempty"` //0-Public | 1-Private | 2-Script
	Email   string `bson:"email,omitempty"`
	FromUrl string `bson:"from_url,omitempty"`
	To      string `bson:"to,omitempty"` //0-PlainText | 1-Markdown | 2-HTML

	Content string `bson:"content,omitempty"`

	IsDeleted bool `bson:"is_deleted"`
}

func (m *Model) NewComment(c *Comment) (cid string, err error) {
	objectID, err := m.CreateDocument(c)
	if err != nil {
		return "", err
	}
	id, err := objectID.MarshalText()
	if err != nil {
		return "", err
	}
	return string(id), nil
}

func (m *Model) UpdateComment(cid string, c *Comment) (err error) {
	objectID, err := stringToObjectID(cid)
	if err != nil {
		return err
	}
	res, err := m.UpdateDocument(objectID, c)
	if res != 1 || err != nil {
		return err
	}
	return nil
}

func (m *Model) GetCommentsByPid(pid string) (comments []Comment, err error) {
	res, err := m.GetDocuments(
		&Comment{Pid: pid},
	)
	if err != nil {
		return nil, err
	}
	for i := range res {
		c := &Comment{}
		doc, _ := bson.Marshal(res[i])
		err = bson.Unmarshal(doc, c)
		if err != nil {
			return nil, err
		}
		comments = append(comments, *c)
	}
	return comments, nil
}
