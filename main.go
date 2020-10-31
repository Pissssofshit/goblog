package main

import (
	"goblog/model"
	"goblog/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	router := gin.Default()

	dsn := "root:root@tcp(127.0.0.1:8889)/blog_service?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.Comment{})

	api := router.Group("/api")
	{
		var userService service.UserService
		var articleService service.ArticleService
		var commentService service.CommentService

		api.POST("/user", userService.Create)
		api.DELETE("/user/:id", userService.Delete)
		api.PATCH("/user/:id", userService.Edit)
		api.GET("/user/:id", userService.Get)

		api.POST("/article", articleService.Create)
		api.DELETE("/article/:id", articleService.Delete)
		api.PATCH("/article/:id", articleService.Edit)
		api.GET("/article/:id", articleService.Get)

		api.POST("/comment", commentService.Create)
		api.DELETE("/comment/:id", commentService.Delete)
		api.PATCH("/comment/:id", commentService.Edit)
		api.GET("/comment/:id", commentService.Get)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
