package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"star-im/src/models"
	"strconv"
)

// GetUserList
// @Summary      用户列表
// @Description  查询用户列表
// @Tags         用户
// @Accept       json
// @Produce      json
// @Success      200  {string}  json:{"code", "msg", "data"}
// @Router       /user/list [get]
func GetUserList(context *gin.Context) {
	data := make([]*models.User, 10)
	data = models.GetUserList()
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询用户列表成功",
		"data": data,
	})
}

// CreateUser
// @Summary      新增用户
// @Description  新增用户
// @Tags         用户
// @Accept       json
// @Produce      json
// @Param		 data body string true "请求体" SchemaExample({\n "username": "张三", \n "password": "123"\n})
// @Success      200  {string}  json{"code", "msg", "data"}
// @Router       /user/create [post]
func CreateUser(context *gin.Context) {
	// 获取请求体
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "请输入用户名密码",
		})
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	fmt.Println("请求内容：", user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "转换实体异常，请排查",
		})
		return
	}

	//user.Username = context.Param("username")
	//user.Password = context.Param("password")
	//confirmPassword := context.Param("confirmPassword")
	//if user.Password != user.Password {
	//	context.JSON(http.StatusOK, gin.H{
	//		"msg": "两次密码不一致",
	//	})
	//}

	user.LoginTime = nil
	user.HeartbeatTime = nil
	user.LogoutTime = nil

	models.CreateUser(user)
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "新增用户成功",
		"data": user,
	})
}

// DeleteUser
// @Summary      删除用户
// @Description  删除用户
// @Tags         用户
// @Accept       json
// @Produce      json
// @Param		 id query string false "id"
// @Success      200  {string}  json {"code", "msg"}
// @Router       /user/delete [get]
func DeleteUser(context *gin.Context) {
	user := models.User{}
	id, _ := strconv.Atoi(context.Query("id"))
	user.ID = uint(id)

	models.DeleteUser(user)
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "删除用户成功",
	})
}

// UpdateUser
// @Summary      修改用户
// @Description  修改用户
// @Tags         用户
// @Accept       json
// @Produce      json
// @Param		 data body string true "请求体" SchemaExample({\n "id": "1", \n "username": "张三", \n "password": "123"\n})
// @Success      200  {string}  json{"code", "msg", "data"}
// @Router       /user/update [post]
func UpdateUser(context *gin.Context) {
	// 获取请求体
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "请求体为空",
		})
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	models.UpdateUser(user)
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "修改用户成功",
		"data": user,
	})
}
