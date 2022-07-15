package model

import (
	"go.mongodb.org/mongo-driver/bson"
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

func (m *Model) GetAllPost() (posts []Post, err error) {
	res, err := m.GetAllDocuments(&Post{})
	for i := range res {
		p := &Post{}
		doc, _ := bson.Marshal(res[i])
		err = bson.Unmarshal(doc, p)
		if err != nil {
			return nil, err
		}
		posts = append(posts, *p)
	}
	return posts, err
}

func (m *Model) GetPostByPid(pid string) (p *Post, err error) {
	var objectID primitive.ObjectID
	objectID.UnmarshalJSON([]byte(pid))
	p = &Post{
		ObjectID: objectID,
	}
	doc, err := m.GetDocument(p)
	if err != nil {
		return nil, err
	}
	err = bson.Unmarshal(doc, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (m *Model) NewPost(p *Post) (pid string, err error) {
	objectID, err := m.CreateDocument(p)
	if err != nil {
		return "", err
	}
	id, err := objectID.MarshalText()
	if err != nil {
		return "", err
	}
	return string(id), nil
}

func (m *Model) UpdatePost(pid string, p *Post) (err error) {
	var objectID primitive.ObjectID
	err = objectID.UnmarshalText([]byte(pid))
	if err != nil {
		return err
	}
	res, err := m.UpdateDocument(objectID, p)
	if res != 1 || err != nil {
		return err
	}
	return nil
}
