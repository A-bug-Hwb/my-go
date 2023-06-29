package service

import (
	"myGo/common/utils/converUtil"
	"myGo/models"
	"myGo/wrBlog/pojo"
)

func getRoleList(userId int) []pojo.RoleBo {
	var roleList []pojo.RoleBo = nil
	var userRolePos []pojo.UserRolePo
	err := models.MysqlDb.Where("user_id = ?", userId).Find(&userRolePos).Error
	if err == nil {
		for i := 0; i < len(userRolePos); i++ {
			var rolePo pojo.RolePo
			errRole := models.MysqlDb.Where("role_id = ?", userRolePos[i].RoleId).First(&rolePo).Error
			if errRole == nil {
				var roleBo pojo.RoleBo
				converUtil.CopyFields(rolePo, &roleBo)
				roleList = append(roleList, roleBo)
			}
		}
	}
	return roleList
}
