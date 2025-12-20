package main

import (
	"fmt"
	"io"
	"log"
	"os"

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
	// 创建输出的日志文件
	gin.SetMode(gin.ReleaseMode)
	// 文件不存在则创建
	file, _ := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	//f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)
	gin.DefaultErrorWriter = io.MultiWriter(file)
	log.SetOutput(file)
	// 自定义日志格式
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 自定义输出格式，字段
		return fmt.Sprintf("%s - [%s] %s %s %d \n",
			param.ClientIP,
			param.Path,
			param.Request.Proto,
			param.Method,
			param.StatusCode,
		)
	}))
	r.Use(gin.Recovery())
	log.Println("日志输出成功")
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
