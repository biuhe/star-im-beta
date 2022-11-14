package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"star-im/src/models"
)

// GetUserList
// @Summary      用户列表
// @Description  查询用户列表
// @Tags         用户列表
// @Accept       json
// @Produce      json
// @Success      200  {string}  json:{"code", "msg", "data"}
// @Router       /user/list [get]
func GetUserList(context *gin.Context) {
	data := make([]*models.User, 10)
	data = models.GetUserList()
	context.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": data,
	})
}
