package services

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

func AddSource(c *gin.Context) {
	var source models.Source
	if err := c.ShouldBindJSON(&source); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	userId := c.GetUint("user_id") // 从上下文获取 user_id

	// 检查单词来源是否已存在
	if models.IsSourceExist(source.Name, userId) {
		utils.Error(c, "来源已存在")
		return
	}

	// 保存单词来源到数据库
	source.UserID = userId
	models.DB.Create(&source)

	utils.Success(c, "添加来源成功", source)
}

func GetSources(c *gin.Context) {
	userId := c.GetUint("user_id") // 从上下文获取 user_id

	var sources []models.Source
	models.DB.Where("user_id = ?", userId).Find(&sources)

	utils.Success(c, "获取来源列表成功", sources)
}
