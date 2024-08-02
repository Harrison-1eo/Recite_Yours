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

	if word.SourceID == 0 {
		utils.Error(c, "请选择生词来源")
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

// GetWords 获取生词列表
func GetWords(c *gin.Context) {
	userID := c.GetUint("user_id")
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	sourceID := c.Query("source_id")
	if startTime == "" || endTime == "" || sourceID == "" {
		utils.Error(c, "缺少参数")
		return
	}

	// 将start_time和end_time转换为time.Time类型
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		utils.Error(c, "参数格式错误")
		return
	}
	et, err := time.Parse("2006-01-02", endTime)
	if err != nil {
		utils.Error(c, "参数格式错误")
		return
	}

	// 将start_time和end_time设置为当天的0点和24点
	st = time.Date(st.Year(), st.Month(), st.Day(), 0, 0, 0, 0, time.Local)
	et = time.Date(et.Year(), et.Month(), et.Day(), 23, 59, 59, 0, time.Local)

	var words []models.Word
	if sourceID == "0" {
		models.DB.Where("user_id = ? AND date BETWEEN ? AND ?", userID, st, et).Find(&words)
	} else {
		models.DB.Where("user_id = ? AND date BETWEEN ? AND ? AND source_id = ?", userID, st, et, sourceID).Find(&words)
	}

	utils.Success(c, "获取成功", words)

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
