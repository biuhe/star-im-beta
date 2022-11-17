package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"star-im/src/utils"
	"time"
)

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
		mt, message, err := ws.ReadMessage()
		if err != nil {

			return
		}
		fmt.Println(string(message))

		if err != nil {
			fmt.Println("订阅websocket消息异常:", err)
		}

		fmt.Println("发送消息：", msg)
		datetime := time.Now().Format("2006-01-02 15:04:02")
		m := fmt.Sprintf("[ws][%s]:%s", datetime, msg)
		err = ws.WriteMessage(mt, []byte(m))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
