package service

import (
	param "goblog/Param"
	"goblog/middleware"
	"goblog/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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

// 定义登陆逻辑
// model.LoginReq中定义了登陆的请求体(name,passwd)
func (userService *UserService) Login(c *gin.Context) {
	var loginReq param.LoginParam
	if c.BindJSON(&loginReq) == nil {
		// 登陆逻辑校验(查库，验证用户是否存在以及登陆信息是否正确)
		user, isPass := userService.user.LoginCheck(loginReq)
		// 验证通过后为该次请求生成token
		if isPass {
			generateToken(c, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "验证失败",
				"data":   nil,
			})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "用户数据解析失败",
			"data":   nil,
		})
	}
}

// token生成器
// md 为上面定义好的middleware中间件
func generateToken(c *gin.Context, user model.User) {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := middleware.NewJWT()

	// 构造用户claims信息(负荷)
	claims := middleware.CustomClaims{
		user.Name,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 签名过期时间
			Issuer:    "goblog",                        // 签名颁发者
		},
	}

	// 根据claims生成token对象
	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}

	log.Println(token)
	// 封装一个响应数据,返回用户名和token
	data := LoginResult{
		Name:  user.Name,
		Token: token,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登陆成功",
		"data":   data,
	})
	return

}

type LoginResult struct {
	Name  string
	Token string
}

type ArticleService struct {
	model *model.Article
}

func (articleService *ArticleService) Create(c *gin.Context) {
}

func (articleService *ArticleService) Delete(c *gin.Context) {
}
func (articleService *ArticleService) Edit(c *gin.Context) {
}

func (articleService *ArticleService) Get(c *gin.Context) {
	article_id := c.Param("id")

	// var article model.Article
	tmp_id, _ := strconv.Atoi(article_id)

	article := articleService.model.Get(uint(tmp_id))

	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"article": article,
		},
	})
}

type CommentService struct {
	model *model.Comment
}

type CommentCreateParam struct {
	ArticleId int    `form:"article_id" json:"article_id"`
	content   string `form:"content" json:"content"`
}

func (commentService *CommentService) Create(c *gin.Context) {
	var param CommentCreateParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (commentService *CommentService) Delete(c *gin.Context) {
}
func (commentService *CommentService) Edit(c *gin.Context) {
}
func (commentService *CommentService) Get(c *gin.Context) {
}
