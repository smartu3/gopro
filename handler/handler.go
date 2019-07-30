package handler

import (
	"net/http"
	"../pkg/errno"
	"github.com/gin-gonic/gin" // Gin web框架
)

type Resopnse struct {
	Code int 	`json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// 返回http 状态码
	c.JSON(http.StatusOK, Resopnese{
		Code: code
		Message: message
		Data: data
	})
}