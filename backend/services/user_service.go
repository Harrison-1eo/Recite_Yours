package services

import (
	"backend/models"
	"backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var JwtSecret = []byte("X2PP7eAXSj")

// Register 新用户注册
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	// 校验用户名是否已存在
	if models.IsUsernameExist(user.Username) {
		utils.Error(c, "用户名已存在")
		return
	}

	// 保存用户到数据库
	if err := models.DB.Create(&user).Error; err != nil {
		utils.Error(c, "注册失败")
		return
	}

	utils.Success(c, "注册成功", nil)
}

// Login 用户登录
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	// 校验用户信息
	if err := models.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; err != nil {
		utils.Error(c, "用户名或密码错误")
		return
	}

	// 如果用户信息正确，生成 JWT
	token, err := generateToken(user.ID)
	if err != nil {
		utils.Error(c, "生成 token 失败")
		return
	}

	utils.Success(c, "登录成功", gin.H{"token": token})
}

func generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

// JWTMiddleware JWT 验证中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			utils.NLI(c, "未登录状态")
			c.Abort()
			return
		}

		// 检查Bearer token格式
		if len(token) < 7 || token[:7] != "Bearer " {
			utils.NLI(c, "token 格式错误")
			c.Abort()
			return
		}

		token = token[7:]

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})

		if err != nil {
			utils.NLI(c, "验证 token 失败")
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
