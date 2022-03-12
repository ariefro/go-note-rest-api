package router

import (
	authRoutes "github.com/ariefro/notes-server/router/auth"
	noteRoutes "github.com/ariefro/notes-server/router/note"
	userRoutes "github.com/ariefro/notes-server/router/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	noteRoutes.SetupNoteRoutes(api)
	userRoutes.SetupUserRoutes(api)
	authRoutes.SetupAuthRoutes(api)
}