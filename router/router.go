package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"myGo/Result"
	"myGo/config"
	"myGo/docs"
	"myGo/jwtAuth"
	"myGo/logger"
	"myGo/wrBlog/router"
	"net/http"
	"strings"
)

//var Router *gin.Engine

func InitRouter() *gin.Engine {
	Router := gin.Default()
	//注册swagger
	registerSwagger(Router)
	//全局拦截
	Router.Use(CheckIp, CheckToken)
	//接口
	router.WrMenuApi(Router)
	return Router
}

func registerSwagger(r gin.IRouter) {
	// API文档访问地址: http://host/swagger/index.html
	// 注解定义可参考 https://github.com/swaggo/swag#declarative-comments-format
	// 样例 https://github.com/swaggo/swag/blob/master/example/basic/api/api.go
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "管理后台接口"
	docs.SwaggerInfo.Description = "实现一个管理系统的后端API服务"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.SerConf.Ip + ":" + config.SerConf.Port
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

// ip拦截
func CheckIp(c *gin.Context) {
	//ip白名单
	if !NotIp(c.ClientIP()) {
		logger.Error("该Ip无权访问:", c.ClientIP())
		c.JSON(http.StatusOK, Result.ResultInfo(Result.IP_NOT_AUTHORIZED))
		c.Abort() //停止执行
	}
	c.Next()
}

// 不拦截的ip(ip白名单)
func NotIp(ip string) bool {
	arr := strings.Split(config.IConf.IpWhite, ",")
	if len(arr) == 0 {
		return true
	}
	//默认本机不拦截
	arr = append(arr, "127.0.0.1", "::1")
	for _, element := range arr {
		if ip == element {
			return true
		}
	}
	return false
}

// token拦截
func CheckToken(c *gin.Context) {
	token := c.GetHeader(config.JwtConf.Header)
	//判断拦截的路由
	if !NotIntercept(c.FullPath()) {
		_, err := jwtAuth.NewJWT().ParseToken(token)
		if err != nil {
			logger.Error("Token验证失败:", err)
			c.JSON(http.StatusOK, Result.ResultInfo(Result.TOKEN_ERR))
			c.Abort() //停止执行
		}
	}
	c.Next() //继续执行
}

// 判断路由是否拦截
func NotIntercept(path string) bool {
	//不拦截的地址
	noPath := []string{"/api/v1/WrMenu/updateWrMenu", "/api/v1/WrMenu/getWrMenuById"}
	for _, element := range noPath {
		if path == element {
			return true
		}
	}
	return false
}
