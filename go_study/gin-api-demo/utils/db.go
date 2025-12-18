package utils

import (
	"github.com/gin-api-demo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 全局DB实例
var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 显示SQL日志
	})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
}
