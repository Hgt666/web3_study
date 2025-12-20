package api

import (
	"log"
	"net/http"

	"github.com/gin-api-demo/model"
	"github.com/gin-api-demo/utils"
	"github.com/gin-gonic/gin"
)

// 创建文章
func CreatePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "failed",
			"msg":    "参数错误",
		})
		log.Println(err)
		return
	}
	// 从上下文中获取用户信息
	userID := c.MustGet("user_id").(uint64)
	post.UserID = uint(userID)
	result := utils.DB.Create(&post)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   post,
			"msg":    "创建文章成功",
		})
		return
	}

}

// 获取文章列表
func ListPost(c *gin.Context) {
	var posts []model.Post
	allPosts := utils.DB.Find(&posts)
	if allPosts.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"msg":    "获取文章列表失败",
		})
		log.Println(allPosts.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   posts,
		"msg":    "获取文章列表成功",
		"total":  len(posts),
	})
}

// 获取文章详情
func PostDetail(c *gin.Context) {
	var post model.Post
	ID := c.Query("id")
	if err := utils.DB.First(&post, ID).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"msg":    "文章不存在或已被删除",
		})
		log.Println(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   post,
		"msg":    "获取文章详情成功",
	})
}

// 更新文章
func UpdatePost(c *gin.Context) {
	// 查出用户的所有文章
	userID := c.MustGet("user_id").(uint64)
	var post []model.Post
	utils.DB.Model(&post).Where("user_id = ?", userID).Find(&post)
	title := c.PostForm("title")
	content := c.PostForm("content")
	id := c.PostForm("id")
	result := utils.DB.Model(&post).Where("id = ?", id).First(&post)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"msg":    "文章不存在",
		})
		c.Abort()
		log.Println("文章不存在")
		return
	}
	if err := utils.DB.Model(&post).Where("id = ?", id).Updates(model.Post{Title: title, Content: content}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"msg":    "更新文章失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   post,
		"msg":    "更新文章成功",
	})
	return
}

// 删除文章
func DeletePost(c *gin.Context) {
	userID := c.MustGet("user_id").(uint64)
	var post []model.Post
	id := c.Query("id")
	utils.DB.Model(&post).Where("user_id = ?", userID).Find(&post)
	// 通过id查找文章
	result := utils.DB.Model(&post).Where("id = ?", id).Delete(&post)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"msg":    "文章不存在或已被删除",
		})
		c.Abort()
		return
	}
	//if utils.DB.Debug().Model(&post).Where("id = ?", id).Delete(&post).Error != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": "failed",
	//		"msg":    "文章不存在或已被删除",
	//	})
	//	c.Abort()
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"msg":    "删除文章成功",
	})
	return

}
