package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"star-im/src/models"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/star-im?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema 如果没有则创建实体
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to create schema")
	}

	// Create
	user := &models.User{}
	user.Username = "张三"
	user.Password = "123456"
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	user.Password = models.EncryptPassword(user.Password, salt)
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1))       // 根据整型主键查找
	db.First(user, "Username = ?", "张三") // 查找 Name 字段值为 张三 的记录

	// Update - 将 user 的 Phone 更新为 13888888888
	db.Model(user).Update("Phone", "13888888888")
	// Update - 更新多个字段
	//db.Model(&product).Updates(models.User{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
