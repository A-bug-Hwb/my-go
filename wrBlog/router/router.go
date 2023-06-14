package router

import (
	"github.com/gin-gonic/gin"
	"myGo/wrBlog/api"
)

func WrMenuApi(Router *gin.Engine) {

	//获取列表
	apiV1 := Router.Group("/api/v1")
	{
		WrMenuApi := apiV1.Group("/WrMenu")
		{
			WrMenuApi.POST("/getWrMenuList", api.GetWrMenuList)
			WrMenuApi.GET("/getWrMenuById", api.GetWrMenuById)
			WrMenuApi.POST("/updateWrMenu", api.UpdateWrMenu)
		}
	}
}
