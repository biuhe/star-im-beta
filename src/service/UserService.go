package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"star-im/src/models"
)

func GetUserList(context *gin.Context) {
	data := make([]*models.User, 10)
	data = models.GetUserList()
	context.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": data,
	})
}
