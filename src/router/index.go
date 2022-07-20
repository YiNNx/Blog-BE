package router

import (
	"github.com/labstack/echo/v4"

	"blog/controller"
	"blog/middleware"
)

// InitRouters 初始化所有路由，可以每个路由分函数分文件写，方便之后维护
func InitRouters(g *echo.Group) {
	g.Use(middleware.JwtCheck)

	groupPost := g.Group("/post")
	initPostRouters(groupPost)

	groupTag := g.Group("/tag")
	initTagRouters(groupTag)

	groupComment := g.Group("/comment")
	initCommentRouters(groupComment)

	groupLike := g.Group("/like")
	initLikeRouters(groupLike)

	groupInfo := g.Group("/info")
	initInfoRouters(groupInfo)

	groupStats := g.Group("/stats")
	initStatsRouters(groupStats)

	groupToken := g.Group("/token")
	initTokenRouters(groupToken)

}

func initPostRouters(g *echo.Group) {
	g.GET("", controller.GetPosts)
	g.GET("/:pid", controller.GetPostByPid)

	g.GET("/deleted", controller.GetDeletedPosts, middleware.JWT...)
	g.POST("", controller.NewPost, middleware.JWT...)
	g.PUT("/:pid", controller.UpdatePost, middleware.JWT...)
	g.DELETE("/:pid", controller.DeletePost, middleware.JWT...)
}

func initCommentRouters(g *echo.Group) {
	g.GET("/:pid", controller.GetCommentsOfPost)
	g.POST("/:pid", controller.NewComment)

	g.DELETE("/:cid", controller.DeleteComment, middleware.JWT...)
}

func initLikeRouters(g *echo.Group) {
	g.PUT("/:pid", controller.LikePost)
}

func initTagRouters(g *echo.Group) {
	g.GET("", controller.GetTags)
}

func initStatsRouters(g *echo.Group) {
	g.GET("", controller.GetStats, middleware.JWT...)
}

func initInfoRouters(g *echo.Group) {
	g.GET("", controller.GetInfo)

	g.PUT("", controller.UpdateInfo, middleware.JWT...)
}

func initTokenRouters(g *echo.Group) {
	g.GET("", controller.LogIn)
}
