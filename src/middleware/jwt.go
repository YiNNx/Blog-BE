package middleware

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"blog/config"
	"blog/util"
	"blog/util/context"
)

var JWT = []echo.MiddlewareFunc{
	middleware.JWTWithConfig(util.Conf),
	CustomJWT,
}

func CustomJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get(config.JWTContextKey).(*jwt.Token).Claims.(*util.JwtUserClaims).Role
		id := c.Get(config.JWTContextKey).(*jwt.Token).Claims.(*util.JwtUserClaims).Id
		if !role || id != 0 {
			err := errors.New("no permission")
			return context.ErrorResponse(c, http.StatusForbidden, err.Error(), err)
		}
		return next(c)
	}
}
