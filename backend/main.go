package main

import (
	"backend/routers"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("your_secret_key") // 请换成复杂的随机字符串

func main() {
	r := gin.Default()
	routers.SetupRoutes(r)
	r.Run(":8080") // 设置监听端口
}
