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
	// m := model.GetModel()
	// defer m.Close()

	// objectID, err := model.StringToObjectID("62d2b7d3bf59f9da6c5fd155")
	// p := &model.Post{
	// 	ObjectID: objectID,
	// }
	// doc, err := m.GetOneDocument(p)
	// log.Logger.Info(doc)
	// log.Logger.Info(err)
}

func main() {
	model.ConnectMongo()
	defer model.Disconnect()
	testUnit()

	e := initEcho(echo.New())

	log.Logger.Fatal(e.Start(config.C.App.Addr))
}
