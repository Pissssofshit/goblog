package main

import (
	"fmt"
	"goblog/middleware"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var model middleware.CustomClaims
	model.Name = "黄枭帅"
	var inter interface{}
	inter = model

	dsn := "root:root@tcp(127.0.0.1:8889)/blog_service?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.User{})

	// user := model.User{
	// 	Name:     "黄枭帅",
	// 	Articles: []model.Article{{Title: "123"}, {Title: "234"}},
	// }

	// db.Create(&user)
	// db.Save(&user)

	var article model.Article

	//db.Where("author_id = 5").First(&article)

	var user model.User

	fmt.Println(user)
	article.AuthorId = 5

	str := db.Model(&article).Association("User")

	fmt.Println(str)

	db.Model(&article).Association("User").Find(&user)

	fmt.Println(user)
}
