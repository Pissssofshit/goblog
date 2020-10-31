package service

import (
	"goblog/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	user model.User
}

func (userService *UserService) Create(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var user model.User
	user.Name = username
	user.Password = password
	user.Create()
	c.JSON(200, gin.H{
		"username": username,
		"password": password,
	})
}

func (user *UserService) Delete(c *gin.Context) {
}
func (user *UserService) Edit(c *gin.Context) {
}
func (user *UserService) Get(c *gin.Context) {
}

type ArticleService struct {
}

func (articleService *ArticleService) Create(c *gin.Context) {
}

func (articleService *ArticleService) Delete(c *gin.Context) {
}
func (articleService *ArticleService) Edit(c *gin.Context) {
}

func (articleService *ArticleService) Get(c *gin.Context) {
	article_id := c.Param("id")

	var article model.Article
	tmp_id, _ := strconv.Atoi(article_id)
	article.ID = uint(tmp_id)
	user := article.GetAuthor()
	c.JSON(200, gin.H{
		"data": user.Name,
		"ds":   article_id,
	})
}

type CommentService struct {
}

func (commentService *CommentService) Create(c *gin.Context) {
}

func (commentService *CommentService) Delete(c *gin.Context) {
}
func (commentService *CommentService) Edit(c *gin.Context) {
}
func (commentService *CommentService) Get(c *gin.Context) {
}