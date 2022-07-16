package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"blog/controller/param"
	"blog/model"
	"blog/util/context"
)

func LogIn(c echo.Context) error {
	email := c.QueryParam("email")
	pwd := c.QueryParam("pwd")

	m := model.GetModel()
	defer m.Close()

	err := m.ValidateUser(email, pwd)
	if err != nil {
		return context.ErrorResponse(c, http.StatusUnauthorized, "", err)
	}
	response := &param.ResponseLogIn{
		Token: context.GenerateToken(0, true),
	}

	return context.SuccessResponse(c, response)
}

func GetInfo(c echo.Context) error {
	m := model.GetModel()
	defer m.Close()

	u, err := m.GetInfo()
	if err != nil {
		return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}

	return context.SuccessResponse(c, u)
}

func UpdateInfo(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}
