package main

import (
	"goblog/service"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	api := router.Group("/api")
	{
		var tagService service.TagService
		api.POST("/tags/:id", tagService.Create)
		api.DELETE("/tags/:id", tagService.Delete)
		api.PATCH("/tags/:id", tagService.Edit)
		api.GET("/tags", tagService.Get)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
