package main

import (
	"echo-pet-api/config"
	"echo-pet-api/database"
	"echo-pet-api/routes"
	"fmt"
	"github.com/labstack/gommon/log"
)

func main() {
	env := config.Env()

	app := config.CreateApp()
	routes.RegisterRoutes(app)

	db := database.Connection()
	database.AutoMigrate(db)

	err := app.Start(fmt.Sprintf(":%s", env.GetString("app.port")))
	if err != nil {
		log.Fatalf("App start error: %v", err)
	}
}
