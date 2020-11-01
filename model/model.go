package model

import (
	param "goblog/Param"
	"goblog/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	Articles []Article `gorm:"foreignKey:AuthorId"`
}

var db_init *gorm.DB

func init() {
	db_init = database.GetInstance()
}

func (user *User) LoginCheck(loginParam param.LoginParam) (User, bool) {
	var result User
	db_init.Where("name = ? and pass_word = ?", loginParam.Username, loginParam.Password).First(&result)

	if result.Name == loginParam.Username {
		return result, true
	}
	return User{}, false
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

func (article *Article) Get(id uint) Article {
	db := database.GetInstance()
	var res Article
	db.Where("id = ?", id).First(&res)
	db.Model(&res).Association("User").Find(&res.User)
	db.Model(&res).Association("Comments").Find(&res.Comments)
	return res
}

func (article *Article) GetAuthor() User {
	db := database.GetInstance()
	var user User
	db.Model(article).Association("User").Find(&user)
	return user
}

type Comment struct {
	gorm.Model
	Content     string
	ArticleId   int
	CommentorId int
	User        User    `gorm:"foreignKey:CommentorId"`
	Article     Article `gorm:"foreignKey:ArticleId"`
}

func (comment *Comment) Create() {
	db := database.GetInstance()
	db.Create(comment)
}
