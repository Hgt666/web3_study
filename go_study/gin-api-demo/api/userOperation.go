package api

import (
	"log"
	"net/http"

	"github.com/gin-api-demo/model"
	"github.com/gin-api-demo/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// Register 用户注册
func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "参数错误",
		})
		return
	}
	// 初始化验证器
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证用户名是否已存在
	result := utils.DB.Model(&user).Where("user_name =?", user.UserName).Find(&user)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "用户名已存在",
		})
		return
	}

	// 获取用户名和密码，并对密码进行加密
	password := user.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	// 将加密后的密码保存到数据库
	user.Password = string(hashedPassword)
	createResult := utils.DB.Create(&user)
	if createResult.RowsAffected == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "注册用户成功",
		})
	}

}

// Login 用户登录
func Login(c *gin.Context) {
	var user model.User
	// 从请求中获取用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 验证用户名和密码是否为空
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "用户名或密码不能为空",
		})
		return
	}
	// 从数据库中查询用户
	dbUser := utils.DB.Model(&user).Where("user_name = ?", username).Find(&user)
	if dbUser.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "用户名或密码错误",
		})
		return
	}
	// 从数据库中获取用户的密码，并验证密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "用户名或密码错误",
		})
	}
	accessToken, refreshToken, err := utils.GenerateToken(uint64(user.ID), username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		log.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "登录成功",
		"data": gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
	// 记录登录日志
	log.Println("用户", username, "登录成功")
}
