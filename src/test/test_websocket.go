package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// webSocket请求ping 返回pong
func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			println("异常:", err)
		}
	}(ws)
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		println("接受到的消息：", string(message))
		if err != nil {
			break
		}

		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func main() {
	bindAddress := "localhost:2303"
	r := gin.Default()
	r.GET("/ping", ping)
	err := r.Run(bindAddress)
	if err != nil {
		println("异常:", err)
	}
}

/**
JS 代码：
<script>
var ws = new WebSocket("ws://localhost:2303/ping");
//连接打开时触发
ws.onopen = function(evt) {
    console.log("Connection open ...");
    ws.send("Hello WebSockets!");
};
//接收到消息时触发
ws.onmessage = function(evt) {
    console.log("Received Message: " + evt.data);
};
//连接关闭时触发
ws.onclose = function(evt) {
    console.log("Connection closed.");
};

</script>
*/
