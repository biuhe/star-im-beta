package models

import "gorm.io/gorm"

// 人员关系
type Contact struct {
	gorm.Model
	UserId   uint   // 谁的关系信息（owner）
	TargetId uint   // 对应谁
	Type     int    // 对应类型
	Desc     string // 描述
}

func (table *Contact) TableName() string {
	return "contact"
}
