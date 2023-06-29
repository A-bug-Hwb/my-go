package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"myGo/common/logger"
	"myGo/config"
)

var MysqlDb *gorm.DB

func InitMysql() {
	MysqlDb = DB()
}

// 初始化链接
func DB() *gorm.DB {
	conf := config.MyConf

	//拼接数据库配置
	var MysqlDSN string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Dbname, conf.Charset, conf.Timeout)
	// 打开连接失败
	db, err := gorm.Open(mysql.Open(MysqlDSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tp_",
			SingularTable: true,
		},
	})
	//defer MysqlDb.Close();
	if err != nil {
		logger.Error("数据库连接失败:" + err.Error())
		panic("数据库连接失败:" + err.Error())
	}
	return db
}
