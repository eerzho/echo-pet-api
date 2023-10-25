package main

import (
	_ "auth-service/docs"
	"auth-service/src/application"
	"auth-service/src/application/database"
	"auth-service/src/application/routes"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm/logger"
	"os"
)

// @title Auth service
// @version 1.0
// @description Auth service docs
// @host localhost:8081
// @BasePath /api/v1
// @SecurityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Env exception: %v", err)
	}

	application.InitializeApp(log.DEBUG)
	application.InitializeDB(logger.Info)

	application.GlobalApp.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.RegisterRoute("/api/v1")
	database.RunFixtures()

	err := application.GlobalApp.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("App start exception: %v", err)
	}
}
