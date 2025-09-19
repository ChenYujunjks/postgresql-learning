package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 定义模型
type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

func main() {
	// 连接信息
	dsn := "host=localhost user=postgres password=你的密码 dbname=你的数据库名 port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// 打开连接
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ 连接数据库失败:", err)
	}
	fmt.Println("✅ 成功连接 PostgreSQL!")

	// 自动迁移（如果没有表会自动创建）
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("❌ 自动迁移失败:", err)
	}

	// -------------------------
	// 1. ORM 插入
	// -------------------------
	db.Create(&User{Name: "Alice", Email: "alice@example.com"})
	db.Create(&User{Name: "Bob", Email: "bob@abc.com"})

	// -------------------------
	// 2. 用 ORM 查询
	// -------------------------
	var user1 User
	db.First(&user1, 1) // 主键=1
	fmt.Println("📌 ORM 查询:", user1)

	// -------------------------
	// 3. 用 Raw SQL 查询
	// -------------------------
	var user2 User
	db.Raw("SELECT id, name, email FROM users WHERE email LIKE ?", "%@example.com").Scan(&user2)
	fmt.Println("📌 Raw SQL 查询:", user2)

	// -------------------------
	// 4. 用 Raw SQL 更新
	// -------------------------
	db.Exec("UPDATE users SET email = ? WHERE name = ?", "bob@example.com", "Bob")

	// 再次 ORM 查询
	var bob User
	db.Where("name = ?", "Bob").First(&bob)
	fmt.Println("📌 更新后的 Bob:", bob)
}
