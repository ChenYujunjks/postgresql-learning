package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

func main() {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{}) // 内存数据库 test.db

	if err != nil {
		log.Fatal("❌ 连接数据库失败:", err)
	}
	fmt.Println("✅ 成功连接 PostgreSQL!")

	// 自动迁移
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal("❌ 自动迁移失败:", err)
	}

	// 实验 1：故意写错 SQL，看看回滚效果
	fmt.Println("\n🚨 实验 1：制造错误，看看事务会不会回滚")

	err = db.Transaction(func(tx *gorm.DB) error {
		// ORM 插入 Charlie
		user := User{Name: "Charlie", Email: "charlie@temp.com"}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		fmt.Println("📌 插入用户:", user)

		// 故意写错 SQL (tablename 错误 "userrs")
		if err := tx.Exec("UPDATE userrs SET email=? WHERE name=?", "charlie@example.com", "Charlie").Error; err != nil {
			return err // 出错 → 回滚
		}

		return nil
	})

	if err != nil {
		fmt.Println("❌ 实验 1 事务失败，已回滚:", err)
	} else {
		fmt.Println("✅ 实验 1 事务成功提交")
	}

	// 查看数据库里有没有 Charlie
	var check1 User
	db.Where("name = ?", "Charlie").First(&check1)
	fmt.Println("📌 实验 1 结束后，数据库里的 Charlie:", check1)

	// 实验 2：修正 SQL，事务成功提交
	fmt.Println("\n✅ 实验 2：修正 SQL，事务应该成功提交")

	err = db.Transaction(func(tx *gorm.DB) error {
		// ORM 插入 Charlie
		user := User{Name: "Charlie", Email: "charlie@temp.com"}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		fmt.Println("📌 插入用户:", user)

		// 正确 SQL
		if err := tx.Exec("UPDATE users SET email=? WHERE name=?", "charlie@example.com", "Charlie").Error; err != nil {
			return err
		}
		fmt.Println("📌 更新 Charlie 邮箱成功")

		return nil
	})

	if err != nil {
		fmt.Println("❌ 实验 2 事务失败，已回滚:", err)
	} else {
		fmt.Println("✅ 实验 2 事务成功提交")
	}

	var check2 User
	db.Where("name = ?", "Charlie").First(&check2)
	fmt.Println("📌 实验 2 结束后，数据库里的 Charlie:", check2)
}
