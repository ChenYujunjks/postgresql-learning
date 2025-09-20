package main

import (
	"fmt"

	"github.com/ChenYujunjks/gorm-postgres/config"
	database "github.com/ChenYujunjks/gorm-postgres/db"
	"github.com/ChenYujunjks/gorm-postgres/models"
)

func main() {
	// 读取配置并初始化数据库
	cfg := config.LoadConfig()
	database.InitDB(cfg)

	// === 业务逻辑直接写在 main ===

	// 创建用户
	database.DB.Create(&models.User{Name: "Alice", Email: "alice@example.com"})
	database.DB.Create(&models.User{Name: "Bob", Email: "bob@abc.com"})

	// 查询用户
	var user1 models.User
	database.DB.First(&user1, 1)
	fmt.Println("📌 ORM 查询:", user1)

	// Raw SQL 查询
	var user2 models.User
	database.DB.Raw("SELECT id, name, email FROM users WHERE email LIKE ?", "%@example.com").Scan(&user2)
	fmt.Println("📌 Raw SQL 查询:", user2)

	// 更新用户
	database.DB.Exec("UPDATE users SET email = ? WHERE name = ?", "bob@example.com", "Bob")

	// 查询更新结果
	var bob models.User
	database.DB.Where("name = ?", "Bob").First(&bob)
	fmt.Println("📌 更新后的 Bob:", bob)
}
