package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/neko-neko/echo-logrus/v2/log"
)

// GetLogger returns a middleware that logs HTTP requests.
func GetLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()

			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			time := stop.UnixMilli() - start.UnixMilli()

			id := req.Header.Get(echo.HeaderXRequestID) + " "
			if id == " " {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
			reqSize := req.Header.Get(echo.HeaderContentLength)
			if reqSize == "" {
				reqSize = "0"
			}
			errmsg := ""
			if err != nil {
				errmsg = "\n" + err.Error()
			}
			log.Debugf("%s%s %s %s %dms"+errmsg,
				id,
				c.RealIP(),
				req.Method,
				req.RequestURI,
				time,
			)
			return err
		}
	}
}
