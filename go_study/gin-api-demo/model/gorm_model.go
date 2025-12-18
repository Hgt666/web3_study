package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string    `gorm:"unique;not null" json:"username" form:"username" binding:"required" validate:"required,min=4,max=16"`
	Password string    `gorm:"not null" json:"password" form:"password"  binding:"required"`
	Email    string    `gorm:"not null" json:"email" form:"email" binding:"required" validate:"email"`
	Posts    []Post    `gorm:"foreignKey:UserID;references:ID"`
	Comments []Comment `gorm:"foreignKey:UserID;references:ID"`
}

type Post struct {
	gorm.Model
	Title    string    `gorm:"not null" json:"title" form:"title" binding:"required"`
	Content  string    `gorm:"not null" json:"content" form:"content" binding:"required"`
	UserID   uint      `gorm:"not null" json:"user_id" form:"user_id" `
	Comments []Comment `gorm:"foreignKey:PostID;references:ID"  json:"comments" form:"comments" `
}

type Comment struct {
	gorm.Model
	UserID  uint
	Content string `gorm:"not null" json:"content" form:"content"`
	PostID  uint
}
