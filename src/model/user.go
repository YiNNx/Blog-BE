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

func (m *Model) ValidateUser(email string, pwd string) error {
	u := &User{
		Email: email,
	}
	doc, err := m.GetDocument(u)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(doc, u)
	if err != nil {
		return err
	}
	err = bcrypt.ValidatePwd(pwd, u.PwdHash)
	if err != nil {
		return err
	}
	return nil
}
