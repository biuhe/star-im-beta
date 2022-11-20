package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"net/http"
	"star-im/src/main/models"
	"star-im/src/main/utils"
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
	index, err := template.ParseFiles("src/resource/index.html", "src/resource/views/chat/head.html")
	if err != nil {
		panic(err)
	}

	err = index.Execute(context.Writer, "index")
	if err != nil {
		panic(err)
	}
}

func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("src/resource/views/user/register.html")
	if err != nil {
		panic(err)
	}
	err = ind.Execute(c.Writer, "register")
	if err != nil {
		panic(err)
	}
}

// Login
// @Summary      登录
// @Description  登录
// @Tags         首页
// @Accept       json
// @Produce      json
// @Param		 data body string true "请求体" SchemaExample({\n "username": "张三", \n "password": "123456"\n})
// @Success      200  {string}  json{"code", "msg", "data"}
// @Router       /login [post]
func Login(context *gin.Context) {
	// 获取请求体
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "请求体为空",
		})
		return
	}

	dto := models.User{}
	// json 解析转换为实体
	err = json.Unmarshal(body, &dto)
	if dto.Username == "" || dto.Password == "" {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "请输入用户名和密码",
		})
		return
	}

	//_, validError := govalidator.ValidateStruct(user)
	//if validError != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"msg": "修改参数不匹配",
	//	})
	//	return
	//}
	// 查找用户
	user := models.FindUserByUsername(dto.Username)
	// 校验密码
	res := models.ValidPassword(dto.Password, user.Salt, user.Password)
	if !res {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "用户名或密码错误",
		})
		return
	}

	// 生成 jwt
	//token加密
	user.Identity, _ = utils.GenToken(user.Username)
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "登录成功",
		"data": user,
	})
}
