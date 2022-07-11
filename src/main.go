package main

import (
	_ "github.com/inconshreveable/log15"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"blog/config"
	"blog/controller"
	mymiddleware "blog/middleware"
	_ "blog/model"
	"blog/router"
	"blog/util/log"
)

func initEcho(e *echo.Echo) *echo.Echo {
	log.SetLoggerOfEcho(e)
	e.Validator = mymiddleware.GetValidator()
	e.HTTPErrorHandler = controller.HTTPErrorHandler
	e.Use(
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.DefaultCORSConfig),
	)
	return e
}

func main() {
	e := initEcho(echo.New())

	g := e.Group(config.C.App.Prefix)
	router.InitRouter(g)

	log.Logger.Fatal(e.Start(config.C.App.Addr))
}
