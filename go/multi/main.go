package main

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 账户模型
type Account struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Balance int
}

func main() {
	// 替换成你自己的数据库信息
	dsn := "host=localhost user=postgres password=你的密码 dbname=你的数据库名 port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ 连接数据库失败:", err)
	}
	fmt.Println("✅ 成功连接 PostgreSQL!")

	// 自动迁移
	if err := db.AutoMigrate(&Account{}); err != nil {
		log.Fatal("❌ 自动迁移失败:", err)
	}

	// 初始化账户（重置余额）
	db.Save(&Account{Name: "Alice", Balance: 1000})
	db.Save(&Account{Name: "Bob", Balance: 500})

	// 模拟并发转账
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ { // 开 5 个 goroutine
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			err := transfer(db, "Alice", "Bob", 100)
			if err != nil {
				fmt.Printf("❌ Goroutine %d 转账失败: %v\n", id, err)
			} else {
				fmt.Printf("✅ Goroutine %d 转账成功\n", id)
			}
		}(i)
	}
	wg.Wait()

	// 查看最终余额
	var alice, bob Account
	db.Where("name = ?", "Alice").First(&alice)
	db.Where("name = ?", "Bob").First(&bob)

	fmt.Printf("\n📌 最终余额 → Alice: %d, Bob: %d (总和 = %d)\n",
		alice.Balance, bob.Balance, alice.Balance+bob.Balance)
}

// 转账函数：from 给 to 转 amount
func transfer(db *gorm.DB, from string, to string, amount int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var sender, receiver Account

		// 查 sender (加 FOR UPDATE 锁，避免并发读写冲突)
		if err := tx.Raw("SELECT * FROM accounts WHERE name = ? FOR UPDATE", from).Scan(&sender).Error; err != nil {
			return err
		}
		if err := tx.Raw("SELECT * FROM accounts WHERE name = ? FOR UPDATE", to).Scan(&receiver).Error; err != nil {
			return err
		}

		// 检查余额
		if sender.Balance < amount {
			return fmt.Errorf("%s 余额不足 (只有 %d)", sender.Name, sender.Balance)
		}

		// 扣钱 + 加钱
		if err := tx.Exec("UPDATE accounts SET balance = ? WHERE id = ?", sender.Balance-amount, sender.ID).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE accounts SET balance = ? WHERE id = ?", receiver.Balance+amount, receiver.ID).Error; err != nil {
			return err
		}

		return nil
	})
}
