package main

import (
	"echo-pet-api/src/config"
	"echo-pet-api/src/database"
	"echo-pet-api/src/routes"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Env exception: %v", err)
	}

	app := config.CreateApp()
	routes.RegisterRoutes(app)

	db := database.Connection()
	database.AutoMigrate(db)

	err := app.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("App start exception: %v", err)
	}
}
