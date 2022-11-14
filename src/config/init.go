package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("src/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置异常：", err)
	}

	fmt.Println("config settings:", viper.Get("settings"))
}

var DB *gorm.DB

func InitDB() {
	DB, _ = gorm.Open(mysql.Open(viper.GetString("settings.database.source")), &gorm.Config{})
}
