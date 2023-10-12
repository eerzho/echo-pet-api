package main

import (
	"auth-service/src/config/application"
	"auth-service/src/config/database"
	"auth-service/src/config/routes"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Env exception: %v", err)
	}

	app := application.CreateApp()
	routes.RegisterRoutes(app)

	db := database.Connection()
	database.AutoMigrate(db)

	err := app.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("App start exception: %v", err)
	}
}
