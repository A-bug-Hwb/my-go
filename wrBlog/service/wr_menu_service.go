package service

import (
	"myGo/models"
	"myGo/utils/boToVoUtil"
	"myGo/wrBlog/pojo"
)

func GetList(wrMenu *pojo.WrMenuBo) ([]*pojo.WrMenuBo, error) {
	var wrMenuBoListBo []*pojo.WrMenuBo
	db := models.MysqlDb
	if wrMenu.MenuName != "" {
		db = db.Where("menu_name LIKE ?", "%"+wrMenu.MenuName+"%")
	}

	err := db.Find(&wrMenuBoListBo).Error
	if err != nil {
		return nil, err
	}

	return wrMenuBoListBo, nil
}

func GetById(id string) (*pojo.WrMenuVo, error) {
	//用bo去查
	var bo *pojo.WrMenuBo
	err := models.MysqlDb.First(&bo, id).Error
	if err != nil {
		return nil, err
	}
	var vo pojo.WrMenuVo
	boToVoUtil.CopyFields(&vo, bo)
	return &vo, nil
}

func Update(wrMenu *pojo.WrMenuBo) (int, error) {
	//models.MysqlDb.Save(wrMenuBo)//更新所有字段
	update := models.MysqlDb.Model(&wrMenu).Updates(wrMenu) //更新非空字段
	err := update.Error
	if err != nil {
		return 0, err
	}
	return int(update.RowsAffected), nil
}
