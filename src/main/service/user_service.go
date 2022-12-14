package service

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
	"star-im/src/main/models"
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
// @Param		 data body string true "请求体" SchemaExample({\n "username": "张三", \n "password": "123456"\n})
// @Success      200  {string}  json{"code", "msg", "data"}
// @Router       /user/create [post]
func CreateUser(context *gin.Context) {
	// 获取请求体
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "请输入用户名密码",
		})
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "转换实体异常，请排查",
		})
		return
	}

	dbUser := models.FindUserByUsername(user.Username)
	if dbUser.ID != 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "用户名已注册",
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

	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	user.Password = models.EncryptPassword(user.Password, salt)
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
// @Param		 data body string true "请求体" SchemaExample({\n "id": "1", \n "username": "张三", \n "password": "123456"\n})
// @Success      200  {string}  json{"code", "msg", "data"}
// @Router       /user/update [post]
func UpdateUser(context *gin.Context) {
	// 获取请求体
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "请求体为空",
		})
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	_, validError := govalidator.ValidateStruct(user)
	if validError != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "修改参数不匹配",
		})
		return
	}

	models.UpdateUser(user)
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "修改用户成功",
		"data": user,
	})
}
