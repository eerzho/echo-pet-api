package main

import (
	"auth-service/src/application"
	"auth-service/src/application/database"
	"auth-service/src/application/routes"
	"auth-service/src/model"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm/logger"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Env exception: %v", err)
	}

	models := []interface{}{
		&model.User{},
		&model.Role{},
		&model.Permission{},
	}

	application.InitializeApp(log.DEBUG)
	application.InitializeDB(logger.Info, models)
	routes.RegisterRoute("/api/v1")
	database.RunFixtures()

	err := application.GlobalApp.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("App start exception: %v", err)
	}
}
