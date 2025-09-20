package database

import (
	"fmt"
	"log"

	"github.com/ChenYujunjks/gorm-postgres/config"
	"github.com/ChenYujunjks/gorm-postgres/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ 数据库连接失败:", err)
	}

	// 自动迁移
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ 自动迁移失败:", err)
	}

	DB = db
	fmt.Println("✅ 成功连接 PostgreSQL!")
}
