package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"myGo/common/result"
	"myGo/wrBlog/pojo"
	"myGo/wrBlog/service"
	"net/http"
)

func Login(c *gin.Context) {
	b, _ := c.GetRawData()
	var loginUserVo pojo.LoginUserVo
	// 反序列化
	_ = json.Unmarshal(b, &loginUserVo)
	token, err := service.Login(loginUserVo.UserName, loginUserVo.Password)
	var res result.Result
	if err != nil {
		res = result.ResultErr("登录失败：" + err.Error())
	} else {
		res = result.ResultSuc("登录成功", token)
	}
	c.JSON(http.StatusOK, res)
}

func Register(c *gin.Context) {
	b, _ := c.GetRawData()
	var registerUser pojo.RegisterUser
	// 反序列化
	_ = json.Unmarshal(b, &registerUser)
	res := result.ResultSuc("注册成功", service.Register(registerUser))
	c.JSON(http.StatusOK, res)
}
