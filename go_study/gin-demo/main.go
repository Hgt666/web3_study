package main

import (
	"github.com/gin-gonic/gin"
	"github.com/htt/gin-demo/GORM"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "sayHello", "status": "success"})
}

func test(c *gin.Context) {
	name := c.Param("name")
	c.String(200, name)
}

func testPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.JSON(200, gin.H{"username": username, "password": password})

}

func testPut(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.JSON(200, gin.H{"username": username, "password": password})
}

func testDelete(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.JSON(200, gin.H{"username": username, "password": password})
}

func main() {
	////r := gin.Default()
	//r := router.InitRouter()
	//r.LoadHTMLGlob("templates/users/*")
	//apiV1 := r.Group("/api/v1")
	//{
	//	apiV1.GET("/hello", sayHello)
	//}
	//
	//err := r.Run(":8081").Error
	//if err != nil {
	//	panic(err)
	//}

	db := GORM.GetDB()
	_ = db.AutoMigrate(&GORM.User{}, &GORM.Profile{})

}
