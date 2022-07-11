// controller 包写web应用的主逻辑，请尽量让它与model层解耦
// 不同router下的controller也建议分文件写
package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"blog/config"
	"blog/util/context"
)

func HelloWorldHandler(c echo.Context) error {
	if config.C.Debug{
		return c.String(200, "yes")
	}
	return c.String(200, "Hello World!")
}

// HTTPErrorHandler 替换默认的错误处理，统一成目前使用的格式
func HTTPErrorHandler(err error, c echo.Context) {
	httpError, ok := err.(*echo.HTTPError)
	if ok {
		context.ErrorResponse(c, httpError.Code, fmt.Sprintf("%s", httpError.Message), err)
		return
	}

	context.ErrorResponse(c, http.StatusInternalServerError, "Unhandled internal server error", err)
}
