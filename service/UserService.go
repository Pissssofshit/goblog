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
}

var db *gorm.DB

func init() {
	db = database.GetInstance()
}

func (userService *UserService) PostArticle(c *gin.Context) {

	// claims, _ := c.Get("claims")

	// intType := reflect.TypeOf(claims).Elem()
	// intPtr2 := reflect.New(intType)
	// // Same as above
	// item := intPtr2.Elem().Interface().(middleware.CustomClaims)

	// for i := 0; i < claims.Elem().NumField(); i++ {
	// 	p := alias.Elem().Field(i)
	// 	p.SetString(params[i])
	// }

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

func (userService *UserService) Create(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var user model.User
	user.Name = username
	user.Password = password
	err := user.Create()

	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}

	var token string
	token, err = generateToken(c, user)

	if err != nil {
		c.JSON(200, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"username": username,
		"password": password,
		"token":    token,
	})
}

func (user *UserService) Delete(c *gin.Context) {
}
func (user *UserService) Edit(c *gin.Context) {
}
func (userService *UserService) Get(c *gin.Context) {
	userName := c.Param("username")
	var user model.User
	not_found := (&user).Get(userName)
	if not_found {
		c.JSON(200, gin.H{
			"user": nil,
		})
	}

	c.JSON(200, gin.H{
		"user": user,
	})
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
