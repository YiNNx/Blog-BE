package router

import (
	"github.com/labstack/echo/v4"

	"blog/controller"
)

// InitRouter 初始化所有路由，可以每个路由分函数分文件写，方便之后维护
func InitRouter(g *echo.Group) {
	// TODO: 完成router
	initIndexRouter(g)
	g.GET("/", controller.HelloWorldHandler)
	grp := g.Group("/user")
	initUserRouter(grp)
}

func initIndexRouter(g *echo.Group) {

}
