package routers

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", services.Register)
	r.POST("/login", services.Login)

	// 需要验证的路由
	authorized := r.Group("/", services.JWTMiddleware())
	{
		authorized.POST("/words", services.AddWord)
		authorized.GET("/words/export", services.ExportWords)
		authorized.GET("/words/stats", services.GetStatistics)
	}
}
