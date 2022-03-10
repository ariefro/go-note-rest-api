package main

import (
	"github.com/ariefro/notes-server/database"
	"github.com/ariefro/notes-server/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)
	
	app.Listen(":3000")
}