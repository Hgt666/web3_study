package main

import (
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDB() *gorm.DB {
	dsn := "root:root123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// 1、模型定义

type User struct {
	gorm.Model
	Name      string
	Age       uint
	Posts     []Post `gorm:"foreignKey:UserID;references:ID"`
	PostCount uint
}

type Post struct {
	gorm.Model
	//PostID       uint
	UserID       uint
	PostTitle    string
	PostContent  string
	Comments     []Comment `gorm:"foreignKey:PostID;references:ID"`
	Status       string
	CommentCount uint
}

type Comment struct {
	gorm.Model
	PostID         uint
	CommentContent string
}

// 3.1 创建文章时自动更新用户的文章数量

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	// 找出是哪个用户创建的
	err = tx.First(&user, "id = ?", p.UserID).Find(&user).Error
	if err != nil {

		tx.Rollback()
		return err
	}
	tx.Model(&user).Update("PostCount", user.PostCount+1)

	return nil
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	var post Post
	err = tx.First(&post, "id = ?", c.PostID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Model(&post).Update("CommentCount", post.CommentCount+1)

	return nil
}

// 3.2 删除评论时检查文章的评论数量，如果评论数量为0，则将文章的评论状态设置为无评论

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var post Post
	// 找出评论对应的文章
	err = tx.Debug().First(&post, "id = ?", c.PostID).Error
	fmt.Println(&post)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 更新文章的评论数量
	tx.Model(&post).Update("CommentCount", post.CommentCount-1)
	if post.CommentCount == 0 {
		tx.Model(&post).Update("Status", "无评论")
	}

	return nil
}

func main() {
	db := getDB()
	_ = db.AutoMigrate(&User{}, &Comment{}, &Post{})

	//初始化数据
	comment1 := Comment{CommentContent: "评论11111"}
	comment2 := Comment{CommentContent: "评论22222"}
	comment3 := Comment{CommentContent: "评论33333"}
	comment4 := Comment{CommentContent: "评论44444"}
	post1 := Post{PostTitle: "文章标题1", PostContent: "文章内容1111", Comments: []Comment{comment1, comment2}}
	post2 := Post{PostTitle: "文章标题2", PostContent: "文章内容2222", Comments: []Comment{comment3}}
	post3 := Post{PostTitle: "文章标题3", PostContent: "文章内容3333", Comments: []Comment{comment4}}
	post4 := Post{PostTitle: "文章标题4", PostContent: "文章内容4444"}
	user1 := User{Name: "user1", Age: 18, Posts: []Post{post1}}
	user2 := User{Name: "user2", Age: 19, Posts: []Post{post3}}
	user3 := User{Name: "user3", Age: 20, Posts: []Post{post2, post4}}
	result := db.Create(&[]User{user1, user2, user3})
	fmt.Println(result.RowsAffected)

	// 查询数据
	//1、查询某个用户发布的所有文章及其对应的评论信息
	var user User
	db.Find(&user, 3)
	db.Preload("Posts").Preload("Posts.Comments").Find(&user)
	for _, post := range user.Posts {
		fmt.Printf("用户%s发布的文章：\n", user.Name)
		fmt.Printf("文章标题：%s\n", post.PostTitle)
		for _, comment := range post.Comments {
			fmt.Printf("评论内容：%s\n", comment.CommentContent)
		}
	}

	// 2、查询评论数量最多的文章信息
	// 1、找出所有用户，遍历文章，统计每个文章的评论数量
	var users []User
	var maxCommentCount int
	var maxPostTitle Post
	db.Preload("Posts").Preload("Posts.Comments").Find(&users)
	postMap := make(map[string]int)
	for _, user := range users {
		for _, post := range user.Posts {
			// 把string转成int
			postMap["文章标题："+post.PostTitle], _ = strconv.Atoi(post.PostTitle)
			postMap[post.PostTitle+"的评论数量_"] = len(post.Comments)
		}
	}
	// 在postMap中找出评论数量最多的文章
	for k, v := range postMap {
		if v > maxCommentCount {
			maxCommentCount = v
			maxPostTitle.PostTitle = k
		}
	}
	fmt.Println("评论数量最多的文章：", maxPostTitle.PostTitle, "评论数量：", maxCommentCount)

	// 3.2 删除Comment一条评论
	var comments Comment
	db.Debug().First(&comments, "post_id = ?", 3).Delete(&comments)

}
