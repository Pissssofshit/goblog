package main

import (
	"goblog/middleware"
	"goblog/model"
	"goblog/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	router := gin.Default()

	dsn := "root:root@tcp(127.0.0.1:8889)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.Comment{})

	var userService service.UserService
	var articleService service.ArticleService

	api := router.Group("/api")

	api.Use(middleware.JWTAuth())
	{
		api.DELETE("/user/:id", userService.Delete)
		api.PATCH("/user/:id", userService.Edit)

		api.POST("/article", userService.PostArticle)
		api.DELETE("/article/:id", articleService.Delete)
		api.PATCH("/article/:id", articleService.Edit)

	}

	//todo 这个奇怪的写法
	api2 := router.Group("/api")
	{
		api2.GET("/article/:id", articleService.Get)
		api2.GET("/user/:username", userService.Get)
		api2.POST("/user", userService.Create)
	}

	router.POST("/api/token", userService.Login)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
