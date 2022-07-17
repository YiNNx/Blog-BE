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

type contextValue map[string]interface{}

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := strings.Replace(c.Request().Header["Authorization"][0], "Bearer ", "", -1)
		claims, err := util.ParseToken(token)
		if err != nil || claims == nil {
			errorMessage(c, err)
			return
		}
		data := contextValue{
			"claims": claims,
		}
		ctx := context.WithValue(r.Context(), "jwt", data)
		// next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
