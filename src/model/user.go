package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"blog/util/bcrypt"
)

type User struct {
	// collectionName get the collection name by tag
	collectionName string `collection:"user" `

	ObjectID primitive.ObjectID `bson:"_id,omitempty" `
	Email    string             `bson:"email,omitempty" `
	PwdHash  string             `bson:"password,omitempty" `

	Avatar        string `bson:"avatar,omitempty"`
	Intro         string `bson:"intro,omitempty"`
	GithubAddress string `bson:"github_address,omitempty"`
}

func (m *model) ValidateUser(email string, pwd string) (*User, error) {
	u := User{
		Email: email,
	}
	doc, err := m.GetDocument(u)
	if err != nil {
		return nil,err
	}
	err = bson.Unmarshal(doc, &u)
	if err != nil {
		return nil,err
	}
	err = bcrypt.ValidatePwd(pwd, u.PwdHash)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (u *User)GetDocument()error{
	m := GetModel()
	defer m.Close()

	doc, err := m.GetDocument(*u)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(doc, u)
	if err != nil {
		return err
	}
	return nil
}

func (p *Post)GetDocument()error{
	m := GetModel()
	defer m.Close()

	doc, err := m.GetDocument(*p)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(doc, p)
	if err != nil {
		return err
	}
	return nil
}

type MongoModel interface{
	GetDocument()error
}
