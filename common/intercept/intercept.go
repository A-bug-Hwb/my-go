package intercept

import (
	"github.com/gin-gonic/gin"
	"myGo/common/auth/jwtAuth"
	token "myGo/common/auth/token"
	"myGo/common/logger"
	"myGo/common/result"
	"net/http"
)

// CheckIp ip拦截
func CheckIp(c *gin.Context) {
	//ip白名单
	if NotIp(c.ClientIP()) {
		c.Next()
	} else {
		logger.Error("该Ip无权访问:", c.ClientIP())
		c.JSON(http.StatusOK, result.ResultDef(result.IP_NOT_AUTHORIZED, "该Ip无权访问", nil))
		c.Abort() //停止执行
	}
}

// CheckToken token拦截
func CheckToken(c *gin.Context) {
	//判断拦截的路由
	if !NotIntercept(c.FullPath()) {
		token := token.GetToken(c)
		if token == "" {
			res := result.ResultDef(result.UNAUTHORIZED, "请求访问："+c.FullPath()+"，认证失败，未登录", nil)
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}
		_, err := jwtAuth.NewJWT().ParseToken(token)
		if err != nil {
			logger.Error("Token验证失败:", err.Error())
			res := result.ResultDef(result.FORBIDDEN, err.Error(), nil)
			c.JSON(http.StatusOK, res)
			c.Abort() //停止执行
			return
		}
	}
	c.Next() //继续执行
}
