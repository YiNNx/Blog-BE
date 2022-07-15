package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"blog/config"
	"blog/controller"
	"blog/model"
	"blog/router"
	"blog/util"
	"blog/util/log"
)

func initEcho(e *echo.Echo) *echo.Echo {
	// Set Logger
	log.SetLoggerOfEcho(e)
	// Set custom validator and HTTPErrorHandler
	e.Validator = util.GetValidator()
	e.HTTPErrorHandler = controller.HTTPErrorHandler
	// Use middleware
	e.Use(
		middleware.Recover(),
		middleware.CORS(),
	)

	// Set prefix and init routers
	g := e.Group(config.C.App.Prefix)
	router.InitRouters(g)

	return e
}

func testUnit() {
	m := model.GetModel()
	defer m.Close()

	data, err := m.DeleteDocument(&model.Post{
		Title: "test2",
	})
	log.Logger.Info((data))
	log.Logger.Info(err)

}

func main() {
	model.ConnectMongo()
	defer model.Disconnect()
	testUnit()

	e := initEcho(echo.New())

	log.Logger.Fatal(e.Start(config.C.App.Addr))
}
