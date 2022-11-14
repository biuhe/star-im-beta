package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndex(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "hello world",
	})
}
