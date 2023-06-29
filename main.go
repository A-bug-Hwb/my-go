package main

import (
	_ "myGo/common/logger" //初始化日志配置
	"myGo/config"
	_ "myGo/config" //初始化配置文件
	"myGo/models"
	"myGo/router" //初始化api
)

func main() {
	//初始化redis配置
	models.InitRedis()
	//初始化数据库配置
	models.InitMysql()
	//启动应用
	r := router.InitRouter()
	r.Run(":" + config.SerConf.Port)

	//开启websocket
	//socket.InitSocket()
}
