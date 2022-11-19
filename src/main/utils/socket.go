package utils

import (
	"context"
	"fmt"
	"star-im/src/main/config"
)

const (
	PublishKey = "websocket"
)

// Publish 发布消息到 redis
func Publish(context context.Context, channel string, msg string) error {
	fmt.Println("发布websocket消息：", msg)
	err := config.RDB.Publish(context, channel, msg).Err()
	if err != nil {
		fmt.Println("发布websocket消息异常：", err)
	}
	return err
}

// Subscribe 订阅 redis 消息
func Subscribe(context context.Context, channel string) (string, error) {
	sub := config.RDB.Subscribe(context, channel)
	fmt.Println("订阅websocket消息", context)
	msg, err := sub.ReceiveMessage(context)

	if err != nil {
		fmt.Println("接受websocket信息异常：", err)
		return "", err
	}
	fmt.Println("订阅的消息内容：", msg.Payload)
	return msg.Payload, err
}
