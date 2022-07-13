package controller

import (
	"github.com/labstack/echo/v4"

	"blog/util/context"
)

func GetStats(c echo.Context) error {
	return context.SuccessResponse(c, nil)
}
