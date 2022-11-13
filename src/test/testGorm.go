package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"star-im/src/models"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/star-im?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema 如果没有则创建实体
	db.AutoMigrate(&models.User{})

	// Create
	user := &models.User{}
	user.Name = "张三"

	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1))   // 根据整型主键查找
	db.First(user, "Name = ?", "张三") // 查找 Name 字段值为 张三 的记录

	// Update - 将 user 的 Password 更新为 123
	db.Model(user).Update("Password", "123")
	// Update - 更新多个字段
	//db.Model(&product).Updates(models.User{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
