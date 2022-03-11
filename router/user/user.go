package userRoute

import (
	userHandler "github.com/ariefro/notes-server/handlers/user"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Get("/:userId", userHandler.GetUser)
}