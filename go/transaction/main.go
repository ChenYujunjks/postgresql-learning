package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 模型
type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

func main() {
	// DSN 替换成你自己的数据库信息
	dsn := "host=localhost user=postgres password=你的密码 dbname=你的数据库名 port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ 连接数据库失败:", err)
	}
	fmt.Println("✅ 成功连接 PostgreSQL!")

	// 自动迁移
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal("❌ 自动迁移失败:", err)
	}

	// -------------------------
	// 事务开始
	// -------------------------
	err = db.Transaction(func(tx *gorm.DB) error {
		// 1. ORM 插入
		user := User{Name: "Charlie", Email: "charlie@temp.com"}
		if err := tx.Create(&user).Error; err != nil {
			return err // 出错 → 回滚
		}
		fmt.Println("📌 插入用户:", user)

		// 2. Raw SQL 更新（这里故意写个错误示例，你可以改对）
		if err := tx.Exec("UPDATE users SET email=? WHERE name=?", "charlie@example.com", "Charlie").Error; err != nil {
			return err // 出错 → 回滚
		}
		fmt.Println("📌 更新 Charlie 邮箱成功")

		// 如果都成功，return nil → 提交事务
		return nil
	})

	// -------------------------
	// 事务结束
	// -------------------------
	if err != nil {
		fmt.Println("❌ 事务失败，已回滚:", err)
	} else {
		fmt.Println("✅ 事务提交成功")
	}

	// 检查 Charlie 的数据
	var charlie User
	db.Where("name = ?", "Charlie").First(&charlie)
	fmt.Println("📌 最终数据库里的 Charlie:", charlie)
}
