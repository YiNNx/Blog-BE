package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"blog/model"
	"blog/util/context"
)

func LogIn(c echo.Context) error {
	email := c.QueryParam("email")
	pwd := c.QueryParam("pwd")

	m:=model.GetModel()
	defer m.Close()
	u, err := model.ValidateUser(email, pwd)
	if err != nil {
		tx.Rollback()
		return context.ErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	response := &responseUserToken{
		Uid:   u.Uid,
		Token: context.GenerateToken(u.Uid, u.Role),
	}

	return context.SuccessResponse(c, nil)
}

func GetInfo(c echo.Context) error {
	m := model.GetModel()
	defer m.Close()

	u := model.User{
		Email: "1826630034@qq.com",
	}
	doc, err := m.GetDocument(u)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(doc, &u)
	if err != nil {
		return err
	}
	return context.SuccessResponse(c, nil)
}

func UpdateInfo(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}
