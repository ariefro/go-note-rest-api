package authRouter

import (
	authHandler "github.com/ariefro/go-note-rest-api/handlers/auth"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/login", authHandler.Login)
}
