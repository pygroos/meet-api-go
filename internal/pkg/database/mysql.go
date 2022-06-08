package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DbHandle *gorm.DB

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	DbHandle, err = gorm.Open(os.Getenv("DB_DRIVER"), dsn)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	DbHandle.SingularTable(true)
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	DbHandle.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	DbHandle.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置连接可复用的最大时间。
	DbHandle.DB().SetConnMaxLifetime(1800 * time.Second)
}
