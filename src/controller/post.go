package controller

import (
	"github.com/labstack/echo/v4"

	"blog/util"
	"blog/util/context"
)

func GetPosts(c echo.Context) error {
	token := util.GenerateToken(5, false)
	return context.SuccessResponse(c, token)
}

func GetPostByPid(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}

func GetTags(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}

func LikePost(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}

func NewPost(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}

func UpdatePost(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}

func DeletePost(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}
