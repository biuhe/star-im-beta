package models

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/set"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
)

// Message 消息
type Message struct {
	gorm.Model
	FormId   uint   // 发送者
	TargetId uint   // 接受者
	Type     string // 发送类型（群聊、私聊、广播）
	Media    int    // 消息类型（文字、图片、音频）
	Content  string // 消息内容
	Pic      string // 图片
	Url      string // 链接
	Desc     string // 描述
	Amount   int    // 其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

// 设置websocket CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 需要 ：发送者ID ，接受者ID ，消息类型，发送的内容，发送类型
// 服务端接收到消息后，转发给chan，内部的udp客户端，监听到chan有消息，
// 就将chan的消息发送给udp的服务端，然后udp的服务端收到消息后，去执行了业务逻辑，根据消息类型走了私聊的业务代码

func Chat(writer http.ResponseWriter, request *http.Request) {
	// 获取参数，并校验 token 合法性
	// token := query.Get("token")
	query := request.URL.Query()
	userId, _ := strconv.ParseInt(query.Get("userId"), 10, 64)

	//msgType := query.Get("msgType")
	//targetId := query.Get("targetId")
	//context := query.Get("context")
	//boolValid := true // todo 验证token 从数据库中查询用户
	conn, err := upGrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取 conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	// 获取用户关系

	// userId 跟 node绑定并加锁
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	//5.完成发送逻辑
	go sendProc(node)
	//6.完成接受逻辑
	go recvProc(node)
	sendMsg(userId, []byte("欢迎进入聊天系统"))
}

// 发送消息
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}

// 接受消息
func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws] <<<<<", string(data))
	}

}

var udpSendChan = make(chan []byte, 1024)

// 广播
func broadMsg(data []byte) {
	udpSendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

// 完成 udp 数据发送协程
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 255),
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case data := <-udpSendChan:
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// 完成 udp 数据接收协程
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n])
	}
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case "1": //私信
		sendMsg(int64(msg.TargetId), data)
		// case 2: //群发
		// sendGroupMsg()
		// case 3://广播
		// sendAllMsg()
		//case 4:
		//
	}
}

func sendMsg(userId int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
