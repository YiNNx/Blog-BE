package main

import (
	"blog-1.0/config"
	"blog-1.0/controller"
	middleware2 "blog-1.0/middleware"
	"blog-1.0/router"
	"blog-1.0/util"
	. "blog-1.0/util/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ok, err := util.ParseFlag()
	if err != nil {
		Logger.Fatal(err)
	}

	if !ok {
		return
	}

	e := echo.New()

	// 自定义未处理错误的handler
	e.HTTPErrorHandler = controller.HTTPErrorHandler

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Validator = middleware2.GetValidator()
	err = middleware2.InitBeforeStart(e)
	if err != nil {
		Logger.Fatal(err)
	}

	gAPI := e.Group(config.C.App.Prefix)
	router.InitRouter(gAPI)

	Logger.Fatal(e.Start(config.C.App.Addr))
}
