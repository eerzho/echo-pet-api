package main

import (
	"auth-service/src/application"
	"auth-service/src/application/database"
	"auth-service/src/application/routes"
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

	application.InitializeApp(log.DEBUG)
	application.InitializeDB(logger.Info)
	routes.RegisterRoute("/api/v1")
	database.RunFixtures()

	err := application.GlobalApp.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("App start exception: %v", err)
	}
}
