package router

import (
	"github.com/gin-api-demo/api"
	"github.com/gin-api-demo/middleware"
	"github.com/gin-gonic/gin"
)

func Init_router(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		// 注册
		apiV1.POST("/register", api.Register)
		// 登录
		apiV1.POST("/login", api.Login)
	}
	apiV2 := r.Group("/api/v2")
	apiV2.Use(middleware.JWTUserMiddleware())
	{

		// 创建文章
		apiV2.POST("/createPost", api.CreatePost)
		// 获取文章列表
		apiV2.GET("/getPostList", api.ListPost)
		// 获取文章详情
		apiV2.GET("/getPostDetail", api.PostDetail)
		// 更新文章
		apiV2.POST("/updatePost", api.UpdatePost)
		// 删除文章
		apiV2.DELETE("/deletePost", api.DeletePost)
		// 创建评论
		apiV2.POST("/createComment", api.CreateComment)
		// 获取评论列表
		apiV2.GET("/GetComments", api.GetComments)

	}

}
