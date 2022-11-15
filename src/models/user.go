package models

import (
	"fmt"
	"gorm.io/gorm"
	"star-im/src/config"
	"time"
)

// User 用户
type User struct {
	// gorm.Model 基础实体定义，包含了id, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	// Name 姓名
	Username string `json:"username"`
	// Password 密码
	Password string `json:"password"`
	// Phone 手机号
	Phone string `json:"phone"`
	// Email 邮箱
	Email string `json:"email"`
	// Identity 唯一标识
	Identity string `json:"identity"`
	// ClientIp 客户端 Ip
	ClientIp string `json:"clientIp"`
	// ClientPort 客户端端口号
	ClientPort string `json:"clientPort"`
	// LoginTime 登录时间
	LoginTime *time.Time `json:"loginTime"`
	// HeartbeatTime 心跳检测时间
	HeartbeatTime *time.Time `json:"heartbeatTime"`
	// LogOutTime 登出时间
	LogoutTime *time.Time `json:"logoutTime"`
	// BoolLogout 是否登出
	BoolLogout bool `json:"boolLogout"`
	// DeviceInfo 设备信息
	DeviceInfo string `json:"deviceInfo"`
}

// TableName 表名
func (table *User) TableName() string {
	return "User"
}

// GetUserList 获取用户列表
func GetUserList() []*User {
	data := make([]*User, 10)
	config.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

// CreateUser 创建用户
func CreateUser(user User) *gorm.DB {
	return config.DB.Create(&user)
}

// DeleteUser 删除用户
func DeleteUser(user User) *gorm.DB {
	return config.DB.Delete(&user)
}

// UpdateUser 修改用户
func UpdateUser(user User) *gorm.DB {
	return config.DB.Model(user).Updates(&user)
}
