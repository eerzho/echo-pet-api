package database

import (
	"auth-service/src/model"
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"sync"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		log.Fatalf("Migration error: %v", err)
	}
}

var (
	dbOnce     sync.Once
	dbInstance *gorm.DB
)

func Connection() *gorm.DB {
	dbOnce.Do(func() {
		dbInstance = newDB()
	})

	return dbInstance
}

func newDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting underlying sql.DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(3)

	return db
}
