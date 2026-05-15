package database

import (
"fmt"
"log"

"store-app/server/internal/config"

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
