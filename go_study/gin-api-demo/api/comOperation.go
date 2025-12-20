package api

import (
	"net/http"
	"strconv"

	"github.com/gin-api-demo/model"
	"github.com/gin-api-demo/utils"
	"github.com/gin-gonic/gin"
)

// 创建评论
func CreateComment(c *gin.Context) {
	userID := c.MustGet("user_id").(uint64)
	postID := c.PostForm("post_id")
	content := c.PostForm("content")
	var comment model.Comment
	// 查找文章是否存在
	isExist := utils.DB.Model(&model.Post{}).Where("id = ?", postID).First(&model.Post{}).RowsAffected > 0
	if !isExist {
		c.JSON(200, gin.H{
			"status": "failed",
			"msg":    "文章不存在",
		})
		c.Abort()
		return
	}
	comment.Content = content

	comment.PostID, _ = strconv.Atoi(postID)
	comment.UserID = uint(userID)
	if err := utils.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"msg":    "创建评论失败",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"msg":    "创建评论成功",
	})
	return
}

// 获取文章的所有评论
func GetComments(c *gin.Context) {
	var comments []model.Comment
	postID := c.Query("post_id")
	// 查找文章是否存在
	result := utils.DB.Model(&model.Post{}).Where("id = ?", postID).Find(&model.Post{})
	if result.RowsAffected == 0 {
		c.JSON(200, gin.H{
			"status": "failed",
			"msg":    "文章不存在",
		})
		c.Abort()
		return
	}
	// 获取文章的评论列
	utils.DB.Model(&comments).Where("post_id = ?", postID).Find(&comments)
	c.JSON(200, gin.H{
		"status": "success",
		"msg":    "获取评论成功",
		"data":   comments,
	})
	return

}
