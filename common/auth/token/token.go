package token

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"myGo/common/auth/jwtAuth"
	"myGo/config"
	"myGo/wrBlog/pojo"
)

const LOGIN_USER_KEY = "login_user_key"
const TOKEN_PREFIX = "Bearer "

var jwt = jwtAuth.NewJWT()

func CreateToken(loginUser pojo.LoginUser) (string, error) {
	var token = uuid.NewV4().String()
	loginUser.Token = token
	var claims jwtAuth.CustomClaims
	claims.Uuid = token
	return jwt.CreateToken(claims)
}

func GetToken(c *gin.Context) string {
	tokenInfo := c.GetHeader(config.TokConf.Header)
	return tokenInfo[len(TOKEN_PREFIX):]
}
