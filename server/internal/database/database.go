package database

import (
	"fmt"
	"log"

	"store-app/server/internal/config"
	"store-app/server/internal/config/business"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Init(cfg *config.Config) {
	dsn := cfg.DSN()
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connected")
}

// Seed 根据业务类型插入种子数据，幂等设计
func Seed(businessType string) {
	seedData := business.Get(businessType)
	if seedData == nil {
		fmt.Printf("Business type %q not found, skipping seed\n", businessType)
		return
	}

	// 幂等：检查店铺是否已有数据
	var count int64
	DB.Model(&struct{ ID int }{}).Table("store").Count(&count)
	if count > 0 {
		fmt.Printf("Seed data already exists (store has %d row(s)), skipping\n", count)
		return
	}

	fmt.Printf("Seeding data for business type: %s\n", businessType)

	// 店铺
	if err := DB.Create(&seedData.Store).Error; err != nil {
		log.Printf("Seed store error: %v", err)
		return
	}

	// 分类
	for _, cat := range seedData.Categories {
		if err := DB.Create(&cat).Error; err != nil {
			log.Printf("Seed category error: %v", err)
			return
		}
	}

	// 商品
	for _, p := range seedData.Products {
		if err := DB.Create(&p).Error; err != nil {
			log.Printf("Seed product error: %v", err)
			return
		}
	}

	fmt.Printf("Seed complete: 1 store, %d categories, %d products\n",
		len(seedData.Categories), len(seedData.Products))
}
