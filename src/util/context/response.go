package context

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"blog/config"
)

// Response 返回值
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Hint    string      `json:"hint,omitempty"`
	Data    interface{} `json:"data"`
}

// SuccessResponse 成功
func SuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "",
		Hint:    "",
		Data:    data,
	})
}

// ErrorResponse 错误
func ErrorResponse(c echo.Context, code int, msg string, err error) error {
	ret := Response{
		Success: false,
		Message: msg,
	}

	if config.C.Debug {
		if err != nil {
			ret.Hint = err.Error()
		}
	}

	return c.JSON(code, ret)
}
