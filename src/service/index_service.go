package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIndex
// @Summary      首页
// @Description  应用正常访问
// @Tags         首页
// @Accept       json
// @Produce      json
// @Success      200  {string}  hello world
// @Failure      400  {string}  http.StatusBadRequest
// @Failure      404  {string}  http.StatusNotFound
// @Failure      500  {string}  http.StatusInternalServerError
// @Router       /index [get]
func GetIndex(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "hello world",
	})
}
