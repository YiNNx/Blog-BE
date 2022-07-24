package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"blog/config"
	"blog/controller/param"
	"blog/model"
	"blog/util/context"
)

func GetPosts(c echo.Context) (err error) {
	// year := c.QueryParam("year")
	// keyword := c.QueryParam("keyword")
	// tag := c.QueryParam("tag")
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	skip, err := strconv.Atoi(c.QueryParam("skip"))

	m := model.GetModel()
	defer m.Close()

	posts, err := m.GetAllPost(int64(limit), int64(skip))
	if err != nil {
		return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}

	var data []param.PostOutline

	for i := range posts {
		pid, err := posts[i].ObjectID.MarshalText()
		if err != nil {
			return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
		}
		p := param.PostOutline{
			Pid:     string(pid),
			Status:  posts[i].Status,
			Title:   posts[i].Title,
			Time:    posts[i].ObjectID.Timestamp().Format(config.C.App.TimeFormat),
			Tags:    posts[i].Tags,
			Excerpt: posts[i].Excerpt,
			Stats: param.Stats{
				Likes:    posts[i].Likes,
				Views:    posts[i].Views,
				Comments: posts[i].Comments,
			},
			IsDeleted: posts[i].IsDeleted,
		}
		if c.Get("authorized").(bool) || p.Status == 0 {
			data = append(data, p)
		}
	}

	return context.SuccessResponse(c, data)
}

func GetPostByPid(c echo.Context) (err error) {
	pid := c.Param("pid")

	m := model.GetModel()
	defer m.Close()

	p, err := m.GetPostByPid(pid)
	if err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}

	if !c.Get("authorized").(bool) && (p.Status != 0 || p.IsDeleted) {
		return context.ErrorResponse(c, http.StatusUnauthorized, "no permission", errors.New("no permission"))
	}

	data := &param.ResponseGetPostByPid{
		Pid:       pid,
		IsDeleted: p.IsDeleted,
		Status:    p.Status,
		Type:      p.Type,
		Title:     p.Title,
		Time:      p.ObjectID.Timestamp().Format(config.C.App.TimeFormat),
		Tags:      p.Tags,
		Content:   p.Content,
		Stats: param.Stats{
			Likes:    p.Likes,
			Views:    p.Views,
			Comments: p.Comments,
		},
	}
	return context.SuccessResponse(c, data)
}

func GetTags(c echo.Context) (err error) {
	return context.SuccessResponse(c, nil)
}

func LikePost(c echo.Context) (err error) {
	pid := c.Param("pid")

	req := new(param.RequestStatus)
	err = c.Bind(req)
	if err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}

	if err := c.Echo().Validator.Validate(req); err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}

	if req.Status {
		m := model.GetModel()
		defer m.Close()

		p, err := m.GetPostByPid(pid)
		if err != nil {
			return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
		}
		p.Likes += 1
		err = m.UpdatePost(pid, p)
		if err != nil {
			return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
		}
	}

	return context.SuccessResponse(c, nil)
}

func NewPost(c echo.Context) (err error) {
	req := new(param.RequestNewPost)
	err = c.Bind(req)
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

func UpdatePost(c echo.Context) (err error) {
	pid := c.Param("pid")

	req := new(param.RequestUpdatePost)
	err = c.Bind(req)
	if err != nil {
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
	err = m.UpdatePost(pid, p)
	if err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}
	return context.SuccessResponse(c, nil)
}

func DeletePost(c echo.Context) (err error) {
	pid := c.Param("pid")

	req := new(param.RequestStatus)
	err = c.Bind(req)
	if err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}

	if err := c.Echo().Validator.Validate(req); err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}

	if req.Status {
		m := model.GetModel()
		defer m.Close()

		p := &model.Post{
			IsDeleted: true,
		}
		err = m.UpdatePost(pid, p)
		if err != nil {
			return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
		}
	}

	return context.SuccessResponse(c, nil)
}

// func DeletePost(c echo.Context) (err error) {
// 	pid := c.Param("pid")

// 	req := new(param.RequestStatus)
// 	err = c.Bind(req)
// 	if err != nil {
// 		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
// 	}

// 	if err := c.Echo().Validator.Validate(req); err != nil {
// 		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
// 	}

// 	if req.Status {
// 		m := model.GetModel()
// 		defer m.Close()

// 		err = m.DeletePost(pid)
// 		if err != nil {
// 			return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
// 		}
// 	}

// 	return context.SuccessResponse(c, nil)
// }

func GetDeletedPosts(c echo.Context) (err error) {
	m := model.GetModel()
	defer m.Close()

	posts, err := m.GetDeletedPost()
	if err != nil {
		return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}

	var data []param.PostOutline

	for i := range posts {
		pid, err := posts[i].ObjectID.MarshalText()
		if err != nil {
			return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
		}
		p := param.PostOutline{
			Pid:     string(pid),
			Status:  posts[i].Status,
			Title:   posts[i].Title,
			Time:    posts[i].ObjectID.Timestamp().Format(config.C.App.TimeFormat),
			Tags:    posts[i].Tags,
			Excerpt: posts[i].Excerpt,
			Stats: param.Stats{
				Likes:    posts[i].Likes,
				Views:    posts[i].Views,
				Comments: posts[i].Comments,
			},
			IsDeleted: posts[i].IsDeleted,
		}
		data = append(data, p)

	}

	return context.SuccessResponse(c, data)
}
