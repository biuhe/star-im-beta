package config

import (
	"fmt"
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
