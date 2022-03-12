package authRouter

import (
	authHandler "github.com/ariefro/notes-server/handlers/auth"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/login", authHandler.Login)
}