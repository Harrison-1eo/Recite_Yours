package services

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddWord 添加生词
func AddWord(c *gin.Context) {
	var word models.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		c.JSON(http.StatusOK, utils.Error(err.Error()))
		return
	}

	userID := c.GetUint("user_id") // 从上下文获取 user_id
	word.UserID = userID

	// 保存单词到数据库
	// 如果需要，处理错误并响应
	// ...

	c.JSON(http.StatusOK, utils.Success("添加成功", word))
}

// ExportWords 导出生词
func ExportWords(c *gin.Context) {
	userID := c.GetUint("user_id")
	// 实现导出逻辑，返回统一格式
	c.JSON(http.StatusOK, utils.Success("导出成功", nil))
}

// GetStatistics 获取统计数据
func GetStatistics(c *gin.Context) {
	userID := c.GetUint("user_id")
	// 实现统计逻辑，返回统一格式
	c.JSON(http.StatusOK, utils.Success("统计成功", nil))
}
