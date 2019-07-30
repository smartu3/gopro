package router

import (
	"net/http"
	"myapi/service"
	"./middleware"
	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter(g *gin.Engine) {
	middlewares := []gin.HandlerFunc{}
	// 中间件
	g.Use(gin.Recovery)
	g.Use(middleware.NoCache)
    g.Use(middleware.Options)
    g.Use(middleware.Secure)
	g.Use(middlewares...)
	// 404
	g.NoRoute(func(c *gin.Context) {
        c.String(http.StatusNotFound, "The incorrect API route.")
	})
	// registered router
	router := g.Group("/user")
	{
		router.POST("/addUser", service.AddUser) // 添加用户
		router.GET("/selectUser", service.Select) // 查询用户
	}
})