package routers

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", services.Register)
	r.POST("/login", services.Login)

	// 需要验证的路由
	api := r.Group("/api")
	api.Use(services.JWTMiddleware())

	sources := api.Group("/sources")
	{
		sources.POST("/add", services.AddSource)
		sources.GET("/list", services.GetSources)
	}

	words := api.Group("/words")
	{
		words.POST("/add", services.AddWord)
		words.GET("/export", services.ExportWords)
		words.GET("/stats", services.GetStatistics)
	}

	//
}
