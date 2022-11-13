package models

import "gorm.io/gorm"

// User 用户
type User struct {
	// gorm.Model 基础实体定义，包含了id, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	// Name 姓名
	Name string
	// Password 密码
	Password string
	// Phone 手机号
	Phone string
	// Email 邮箱
	Email string
	// Identity 唯一标识
	Identity string
	// ClientIp 客户端 Ip
	ClientIp string
	// ClientPort 客户端端口号
	ClientPort string
	// LoginTime 登录时间
	LoginTime uint64
	// HeartbeatTime 心跳检测时间
	HeartbeatTime uint64
	// LogOutTime 登出时间
	LogOutTime uint64
	// boolLogout 是否登出
	boolLogout bool
	// DeviceInfo 设备信息
	DeviceInfo string
}

// TableName 表名
func (table *User) TableName() string {
	return "User"
}
