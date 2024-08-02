package main

import (
	"backend/models"
	"backend/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	models.InitDB()

	r := gin.Default()
	routers.SetupRoutes(r)
	err := r.Run("0.0.0.0:8010")
	if err != nil {
		panic("启动失败" + err.Error())
	} // 设置监听端口
}
