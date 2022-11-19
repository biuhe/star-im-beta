package models

import "gorm.io/gorm"

// 群
type Group struct {
	gorm.Model
	Name    string // 群名
	OwnerId uint   // 群主用户id
	Type    int    // 类型
	Icon    string // 图标
	Desc    string // 描述
}

func (table *Group) TableName() string {
	return "group"
}
