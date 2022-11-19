package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("src/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置异常：", err)
	}

	fmt.Println("初始化 app 配置")
}

var DB *gorm.DB

func InitDB() {
	// 初始化 logger 日志，自定义日志模板，打印 SQL 语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			// 慢SQL阈值（秒）
			SlowThreshold: time.Second,
			// 日志级别
			LogLevel: logger.Info,
			// 彩色
			Colorful: true,
		},
	)

	// 连接数据库，并配置慢SQL日志
	DB, _ = gorm.Open(mysql.Open(viper.GetString("settings.database.source")),
		&gorm.Config{Logger: newLogger})
	fmt.Println("初始化数据库配置")

}

var RDB *redis.Client

func InitCache() {
	RDB = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("settings.cache.redis.addr"),
		Password:     viper.GetString("settings.cache.redis.password"), // no password set
		DB:           viper.GetInt("settings.cache.redis.db"),          // use default DB
		PoolSize:     viper.GetInt("settings.cache.redis.poolSize"),
		MinIdleConns: viper.GetInt("settings.cache.redis.minIdleConn"),
	})
	pong, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("初始化缓存配置失败：", err)
		return
	}
	fmt.Println("初始化缓存配置", pong)
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish 。。。。", msg)
	err = RDB.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := RDB.Subscribe(ctx, channel)
	fmt.Println("Subscribe 。。。。", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe 。。。。", msg.Payload)
	return msg.Payload, err
}
