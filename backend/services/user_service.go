package services

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Register 新用户注册
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, utils.Error(err.Error()))
		return
	}

	// 进行密码哈希处理
	// ...

	// 保存用户到数据库
	// ...

	c.JSON(http.StatusOK, utils.Success("注册成功", nil))
}

// Login 用户登录
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, utils.Error(err.Error()))
		return
	}

	// 校验用户信息
	// ...

	// 如果用户信息正确，生成 JWT
	token, err := generateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusOK, utils.Error("生成token失败"))
		return
	}

	c.JSON(http.StatusOK, utils.Success("登录成功", gin.H{"token": token}))
}

func generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWT 验证中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, utils.Error("授权失败"))
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusOK, utils.Error("授权失败"))
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
