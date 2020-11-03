package service

import (
	param "goblog/Param"
	"goblog/database"
	"goblog/middleware"
	"goblog/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	user model.User
	db   *gorm.DB
}

var db *gorm.DB

func init() {
	db = database.GetInstance()
}

func NewUserService() UserService {
	var userService UserService
	userService.res = param.NewResponseStruct()
	userService.db = database.GetInstance()
	return userService
}

func (userService *UserService) PostArticle(c *gin.Context) {

	//TODO
	//从jwt中获取登录用户的信息
	token := c.Request.Header.Get("token")

	claim, _ := middleware.NewJWT().ParserToken(token)

	userName := claim.Name

	title := c.Query("title")
	content := c.Query("content")

	var user model.User

	db.Where("name = ?", userName).First(&user)

	err := db.Model(&user).Where("sdasda = sdasds").Association("Articles").Append(&model.Article{Title: title, Content: content})

	if err != nil {
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "success",
	})

}

type CreateUserParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (userService *UserService) Create(c *gin.Context) {
	var json CreateUserParam

	response := param.NewResponseStruct()
	if err := c.ShouldBindJSON(&json); err != nil {
		response.Code = 1
		response.Msg = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var user model.User
	user.Name = json.Username
	user.Password = json.Password
	err := user.Create()

	if err != nil {
		response.Msg = err.Error()
	} else {
		var token string
		token, err = generateToken(c, user)

		if err != nil {
			response.Msg = err.Error()
			response.Code = 1
		}
		response.Data["token"] = token
		response.Data["username"] = user.Name
	}

	c.JSON(200, response)
}

func (userService *UserService) Delete(c *gin.Context) {
	token := c.Request.Header.Get("token")

	claim, _ := middleware.NewJWT().ParserToken(token)

	userName := claim.Name

	var user model.User
	user.Name = userName
	db.First(&user)
	db.Select("Articles").Delete(&user)
	response := param.NewResponseStruct()
	c.JSON(200, response)
}

func (user *UserService) Edit(c *gin.Context) {
}

type QueryUserParam struct {
	UserName string `uri:"username" binding:"required"`
}

func (userService *UserService) Get(c *gin.Context) {

	var queryUserParam QueryUserParam

	response := param.NewResponseStruct()
	if err := c.ShouldBindUri(&queryUserParam); err != nil {
		response.Code = 1
		response.Msg = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var user model.User
	not_found := (&user).Get(queryUserParam.UserName)
	if !not_found {
		response.Data["user"] = user
	}

	c.JSON(200, response)
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
			token, err := generateToken(c, user)
			if err == nil {
				c.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": err.Error(),
				})
			}
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
			"msg":    "登录参数缺失",
			"data":   nil,
		})
	}
}

// token生成器
// md 为上面定义好的middleware中间件
func generateToken(c *gin.Context, user model.User) (string, error) {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := middleware.NewJWT()

	// 构造用户claims信息(负荷)
	claims := middleware.CustomClaims{
		user.Name,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),    // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600*30), // 签名过期时间
			Issuer:    "goblog",                           // 签名颁发者
		},
	}

	// 根据claims生成token对象
	token, err := j.CreateToken(claims)

	if err != nil {
		return "", err
	}

	log.Println(token)

	return token, nil

}

type LoginResult struct {
	Name  string
	Token string
}

type ArticleService struct {
	model *model.Article
}

func (articleService *ArticleService) Create(c *gin.Context) {
	claims, _ := c.Get("claims")

	c.JSON(200, gin.H{
		"claims": claims,
	})
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
