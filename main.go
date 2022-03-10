package main

import (
	"github.com/ariefro/notes-server/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()
	
	app.Listen(":3000")
}