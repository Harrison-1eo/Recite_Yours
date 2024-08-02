package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(200, Response{
		Code:    0,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, message string) {
	c.JSON(200, Response{
		Code:    1,
		Message: message,
		Data:    nil,
	})
}

func NLI(c *gin.Context, message string) {
	c.JSON(200, Response{
		Code:    2,
		Message: message,
		Data:    nil,
	})
}
