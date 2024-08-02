package services

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// AddWord 添加生词
func AddWord(c *gin.Context) {
	var word models.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	userID := c.GetUint("user_id") // 从上下文获取 user_id
	word.UserID = userID
	word.Date = time.Now()

	// 保存单词到数据库
	if err := models.DB.Create(&word).Error; err != nil {
		utils.Error(c, "添加失败")
		return
	}

	utils.Success(c, "添加成功", word)
}

// ExportWords 导出生词
func ExportWords(c *gin.Context) {

	userID := c.GetUint("user_id")
	// 实现导出逻辑，返回统一格式
	utils.Success(c, "导出成功", userID)
}

// GetStatistics 获取统计数据
func GetStatistics(c *gin.Context) {
	userID := c.GetUint("user_id")
	// 实现统计逻辑，返回统一格式
	utils.Success(c, "获取成功", userID)
}
