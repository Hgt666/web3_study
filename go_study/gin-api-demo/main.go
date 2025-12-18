package main

import (
	"fmt"

	"github.com/gin-api-demo/config"
	"github.com/gin-api-demo/model"
	"github.com/gin-api-demo/router"
	"github.com/gin-api-demo/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitDB()
	db := utils.DB
	// 迁移数据库
	_ = db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	// 初始化gin
	r := gin.Default()
	// 注册中间件
	//r.Use(middleware.JWTAuth())
	// 注册路由
	router.Init_router(r)
	// 初始化验证器

	// 启动服务
	err := r.Run(fmt.Sprintf(":%d", config.ServerPort))
	if err != nil {
		panic(err)
	}

}
