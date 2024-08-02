package routers

import (
	"backend/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 跨域问题
	// 跨域中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true 万事大吉
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	}))

	// 添加 Cross-Origin-Resource-Policy 头部，解决资源跨域问题
	r.Use(func(c *gin.Context) {
		c.Header("Cross-Origin-Resource-Policy", "cross-origin")
		c.Next()
		c.Header("Cross-Origin-Resource-Policy", "cross-origin")
	})

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
		words.GET("/autocomplete", services.Autocomplete)
		words.GET("/meaning", services.GetMeaning)
		words.POST("/add", services.AddWord)
		words.GET("/list", services.GetWords)
		words.GET("/export", services.ExportWords)
		words.GET("/stats", services.GetStatistics)
	}

	//
}
