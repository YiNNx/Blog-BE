package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"blog/config"
	"blog/controller/param"
	"blog/model"
	"blog/util/context"
)

func GetPosts(c echo.Context) error {
	m := model.GetModel()
	defer m.Close()

	posts, err := m.GetAllPost()
	if err != nil {
		return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}

	return context.SuccessResponse(c, posts)
}

func GetPostByPid(c echo.Context) error {
	pid := c.Param("pid")

	m := model.GetModel()
	defer m.Close()

	p, err := m.GetPostByPid(pid)
	if err != nil {
		return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}

	data := &param.ResponseGetPostByPid{

		Pid:     pid,
		Status:  p.Status,
		Type:    p.Type,
		Title:   p.Title,
		Time:    p.ObjectID.Timestamp().Format(config.C.App.TimeFormat),
		Tags:    p.Tags,
		Content: p.Content,
		Stats: param.Stats{
			Likes:    p.Likes,
			Views:    p.Views,
			Comments: p.Comments,
		},
	}
	return context.SuccessResponse(c, data)
}

func GetTags(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}

func LikePost(c echo.Context) error {
	//pid := c.Param("pid")

	req := new(param.RequestStatus)
	err := c.Bind(req)
	if err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}

	if req.Status{

	}

	return context.SuccessResponse(c, nil)
}

func NewPost(c echo.Context) error {
	req := new(param.RequestNewPost)
	err := c.Bind(req)
	if err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}

	if err := c.Echo().Validator.Validate(req); err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}

	m := model.GetModel()
	defer m.Close()

	p := &model.Post{
		Status:  req.Status,
		Type:    req.Type,
		Title:   req.Title,
		Excerpt: req.Excerpt,
		Content: req.Content,
		Tags:    req.Tags,
	}
	pid, err := m.NewPost(p)
	if err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}
	data := param.ResponseNewPost{
		Pid: pid,
	}
	return context.SuccessResponse(c, data)
}

func UpdatePost(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}

func DeletePost(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}
