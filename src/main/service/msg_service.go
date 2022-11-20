package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"star-im/src/main/models"
	"star-im/src/main/utils"
	"strconv"
	"time"
)

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles("src/resource/views/chat/index.html",
		"src/resource/views/chat/head.html",
		"src/resource/views/chat/foot.html",
		"src/resource/views/chat/tabmenu.html",
		"src/resource/views/chat/concat.html",
		"src/resource/views/chat/group.html",
		"src/resource/views/chat/profile.html",
		"src/resource/views/chat/createcom.html",
		"src/resource/views/chat/userinfo.html",
		"src/resource/views/chat/main.html")
	if err != nil {
		panic(err)
	}

	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.User{}
	user.ID = uint(userId)
	user.Identity = token
	//fmt.Println("ToChat>>>>>>>>", user)
	err = ind.Execute(c.Writer, user)
	if err != nil {
		panic(err)
	}
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}

// 设置websocket CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(context *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		fmt.Println("异常：", err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			println("异常:", err)
		}
	}(ws)

	// 消息处理
	MsgHandler(context, ws)
}

func MsgHandler(context *gin.Context, ws *websocket.Conn) {
	for {
		msg, err := utils.Subscribe(context, utils.PublishKey)

		if err != nil {
			fmt.Println("订阅websocket消息异常:", err)
		}

		fmt.Println("发送消息：", msg)
		datetime := time.Now().Format("2006-01-02 15:04:02")
		m := fmt.Sprintf("[ws][%s]:%s", datetime, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func Chat(context *gin.Context) {
	models.Chat(context.Writer, context.Request)
}
