package config

import "fmt"

// 全局配置
var (
	DBHost     = "localhost"
	DBPort     = 3306
	DBUser     = "root"
	DBPassword = "root123456"
	DBName     = "gorm"
	ServerPort = 8081
	JWTSecret  = "secret"
)

func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser, DBPassword, DBHost, DBPort, DBName)
}
