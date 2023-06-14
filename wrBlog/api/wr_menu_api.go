package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"myGo/Result"
	"myGo/logger"
	"myGo/wrBlog/pojo"
	"myGo/wrBlog/service"
	"net/http"
)

// GetWrMenuList
// @Summary 获取列表
// @Description 获取所有菜单列表
// @Tags WrMenu
// @Accept json
// @Produce json
// @Param json body pojo.WrMenuVo true "实体类信息"
// @Success 0 {string} string "ok"
// @Router /WrMenu/getWrMenuList [post]
func GetWrMenuList(c *gin.Context) {
	b, _ := c.GetRawData()
	var wrMenuBo pojo.WrMenuBo
	// 反序列化
	_ = json.Unmarshal(b, &wrMenuBo)
	list, err := service.GetList(&wrMenuBo)

	if err != nil {
		res := Result.ResultInfo(Result.ERROR)
		c.JSON(http.StatusOK, res)
		return
	}
	res := Result.ResultSuc(list)
	c.JSON(http.StatusOK, res)
}

// GetWrMenuById
// @Summary 获取详情
// @Description 根据菜单id获取菜单详情
// @Tags WrMenu
// @Accept json
// @Produce json
// @Param id query string true "菜单id"
// @Success 0 {string} string "ok"
// @Router /WrMenu/getWrMenuById [get]
func GetWrMenuById(c *gin.Context) {
	info, err := service.GetById(c.Query("id"))
	if err != nil {
		//res := Result.ResultInfo(Result.ERROR)
		logger.Error(err.Error())
		//c.JSON(http.StatusOK, res)
		//return
	}
	res := Result.ResultSuc(info)
	c.JSON(http.StatusOK, res)
}

// 修改
func UpdateWrMenu(c *gin.Context) {
	b, _ := c.GetRawData()
	var wrMenuBo pojo.WrMenuBo
	// 反序列化
	_ = json.Unmarshal(b, &wrMenuBo)

	//models.MysqlDb.Save(wrMenuBo)//更新所有字段
	num, err := service.Update(&wrMenuBo) //更新非空字段
	if err != nil {
		res := Result.ResultInfo(Result.ERROR)
		c.JSON(http.StatusOK, res)
		return
	}
	res := Result.ResultSuc(num)
	c.JSON(http.StatusOK, res)
}
