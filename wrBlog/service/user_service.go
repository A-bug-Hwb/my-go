package service

import (
	"errors"
	"fmt"
	"myGo/common/auth/token"
	"myGo/common/utils/bcrypt"
	"myGo/common/utils/converUtil"
	"myGo/models"
	"myGo/wrBlog/pojo"
)

func Login(username string, password string) (string, error) {
	var userPo *pojo.UserPo
	err := models.MysqlDb.Where("user_name = ?", username).First(&userPo).Error
	if err != nil {
		return "", errors.New("没有此用户")
	}
	//转为Bo
	var userBo pojo.UserBo
	converUtil.CopyFields(&userBo, userPo)
	//判断密码对不对
	if !bcrypt.CheckPasswordHash(password, userBo.Password) {
		return "", errors.New("账号或密码错误")
	}
	return token.CreateToken(getLoginUser(&userBo))
}

func Register(registerUser pojo.RegisterUser) bool {
	if registerUser.Password != registerUser.ConfirmPassword {
		fmt.Print("两次密码不一致")
	}
	newPassword, _ := bcrypt.HashPassword(registerUser.Password)
	registerUser.Password = newPassword
	err := models.MysqlDb.Create(registerUser).Error
	if err != nil {
		return false
	}
	return true
}
