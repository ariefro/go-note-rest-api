package router

import (
	authRoutes "github.com/ariefro/go-note-rest-api/router/auth"
	noteRoutes "github.com/ariefro/go-note-rest-api/router/note"
	userRoutes "github.com/ariefro/go-note-rest-api/router/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	noteRoutes.SetupNoteRoutes(api)
	userRoutes.SetupUserRoutes(api)
	authRoutes.SetupAuthRoutes(api)
}
