package services

import (
	"backend/models"
	"backend/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Autocomplete 通过前缀查找匹配的单词
func Autocomplete(c *gin.Context) {
	prefix := c.Query("prefix")
	if prefix == "" {
		utils.Error(c, "prefix 参数不能为空")
		return
	}

	var words []models.EnChDict
	// 按照priority降序排列，取前10个
	if err := models.DB.Where("word LIKE ?", prefix+"%").Order("priority DESC").Limit(20).Find(&words).Error; err != nil {
		utils.Error(c, "查询失败")
		return
	}

	// 返回匹配的单词，需要将单词去重
	suggestions := make(map[string]models.EnChDict)
	for _, word := range words {
		suggestions[word.Word] = word
	}

	s := make([]string, 0, len(suggestions))
	for _, v := range suggestions {
		s = append(s, v.Word)
	}

	utils.Success(c, "查询成功", s)
}

// GetMeaning 获取单词的含义
func GetMeaning(c *gin.Context) {
	word := c.Query("word")
	if word == "" {
		utils.Error(c, "word 参数不能为空")
		return
	}

	var entry models.EnChDict
	if err := models.DB.Where("word = ?", word).Order("priority DESC").First(&entry).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Error(c, "未找到该单词")
		} else {
			utils.Error(c, "查询失败")
		}
		return
	}

	utils.Success(c, "查询成功", entry)
}
