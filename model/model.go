package model

import (
	"goblog/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	Articles []Article `gorm:"foreignKey:AuthorId"`
}

func (user *User) Create() {
	db := database.GetInstance()
	db.Create(user)
}

type Article struct {
	gorm.Model
	Title    string
	Content  string
	IsDel    int
	State    int
	AuthorId int
	User     User      `gorm:"foreignKey:AuthorId"`
	Comments []Comment `gorm:"foreignKey:ArticleId"`
}

func (article *Article) Create() {
	db := database.GetInstance()
	db.Create(article)
}

func (article *Article) GetAuthor() User {
	db := database.GetInstance()
	var user User

	db.Model(article).Where("ds = dasd").Association("User").Find(&user)
	return user
}

type Comment struct {
	gorm.Model
	Content   string
	ArticleId int
	Article   Article `gorm:"foreignKey:ArticleId"`
}

func (comment *Comment) Create() {
	db := database.GetInstance()
	db.Create(comment)
}
