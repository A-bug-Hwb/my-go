package router

import (
	"github.com/gin-gonic/gin"
	"myGo/common/intercept"
	"myGo/common/request"
	"myGo/common/swagger"
	"myGo/wrBlog/router"
)

// InitRouter var Router *gin.Engine
func InitRouter() *gin.Engine {
	Router := gin.Default()
	//注册swagger
	swagger.RegisterSwagger(Router)
	//全局拦截
	Router.Use(intercept.CheckIp, intercept.CheckToken, request.SetRequest)
	//接口
	router.WrMenuApi(Router)
	return Router
}
