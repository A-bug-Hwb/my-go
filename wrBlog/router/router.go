package router

import (
	"github.com/gin-gonic/gin"
	"myGo/wrBlog/api"
)

func WrMenuApi(Router *gin.Engine) {

	//获取列表
	apiV1 := Router.Group("/api/v1")
	{
		UserApi := apiV1.Group("/user")
		{
			UserApi.POST("/login", api.Login)
			UserApi.POST("/register", api.Register)
			//UserApi.GET("/getUserInfo", api.GetUserInfo)
		}
	}
}
