package database

import (
	"echo-pet-api/config"
	"echo-pet-api/src/model"
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Post{},
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
	env := config.Env()

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		env.GetString("db.host"),
		env.GetString("db.user"),
		env.GetString("db.name"),
		env.GetString("db.password"))

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
