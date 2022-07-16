package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"blog/config"
	"blog/controller/param"
	"blog/model"
	"blog/util/context"
)

func GetCommentsOfPost(c echo.Context) (err error) {
	pid := c.Param("pid")

	m := model.GetModel()
	defer m.Close()

	comments, err := m.GetCommentsByPid(pid)
	if err != nil {
		return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}

	var data []param.ResponseGetComments

	for i := range comments {
		cid, err := comments[i].ObjectID.MarshalText()
		if err != nil {
			return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
		}
		c := param.ResponseGetComments{
			Cid:       string(cid),
			ParentCid: comments[i].ParentCid,
			Time:      comments[i].ObjectID.Timestamp().Format(config.C.App.TimeFormat),
			From:      comments[i].From,
			FromUrl:   comments[i].FromUrl,
			Content:   comments[i].Content,
		}
		data = append(data, c)
	}

	return context.SuccessResponse(c, &data)
}

func NewComment(c echo.Context) (err error) {
	pid := c.Param("pid")

	req := new(param.RequestComment)
	err = c.Bind(req)
	if err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}

	if err := c.Echo().Validator.Validate(req); err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}

	m := model.GetModel()
	defer m.Close()

	comment := &model.Comment{
		Pid:       pid,
		ParentCid: req.ParentCid,
		From:      req.From,
		Email:     req.Email,
		FromUrl:   req.FromUrl,
		To:        req.To,
		Content:   req.Content,
	}
	cid, err := m.NewComment(comment)
	if err != nil {
		return context.ErrorResponse(c, http.StatusBadRequest, "", err)
	}
	data := param.ResponseNewComment{
		Cid: cid,
	}

	return context.SuccessResponse(c, data)
}

func DeleteComment(c echo.Context) (err error) {
	cid := c.Param("cid")

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

		del := &model.Comment{
			IsDeleted: true,
		}
		err = m.UpdateComment(cid, del)
		if err != nil {
			return context.ErrorResponse(c, http.StatusInternalServerError, "", err)
		}
	}

	return context.SuccessResponse(c, nil)
}
