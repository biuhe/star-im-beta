package models

import (
	"gorm.io/gorm"
	"star-im/src/main/config"
)

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

// SearchFriend 查找好友
func SearchFriend(userId uint) []User {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	// 根据所述用户id 查询用户关系
	config.DB.Where("user_id = ? and type= 1", userId).Find(&contacts)
	for _, v := range contacts {
		// 将对应人员 id 组装成数组
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]User, 0)
	// 查询对应人员
	config.DB.Where("id in ?", objIds).Find(&users)
	return users
}

func SearchUserByGroupId(communityId uint) []uint {
	contacts := make([]Contact, 0)
	objIds := make([]uint, 0)
	config.DB.Where("target_id = ? and type=2", communityId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, v.UserId)
	}
	return objIds
}
