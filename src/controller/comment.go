package controller

import (
	"github.com/labstack/echo/v4"

	"blog/util/context"
)

func GetCommentsOfPost(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}

func Comment(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}

func DeleteComment(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}
