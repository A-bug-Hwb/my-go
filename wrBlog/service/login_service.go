package service

import (
	"myGo/common/request"
	"myGo/wrBlog/pojo"
)

func getLoginUser(userBo *pojo.UserBo) pojo.LoginUser {
	var roleList = getRoleList(userBo.UserId)
	if roleList != nil {
		userBo.RoleBos = roleList
	}
	var loginUser pojo.LoginUser
	loginUser.UserBo = *userBo
	loginUser.Permissions = []string{"user:list"}
	loginUser.Ipaddr = request.GetRequest().ClientIP()
	return loginUser
}
