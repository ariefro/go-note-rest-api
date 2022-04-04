package userRoute

import (
	userHandler "github.com/ariefro/notes-server/handlers/user"
	"github.com/ariefro/notes-server/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Get("/:userId", userHandler.GetUser)
	user.Post("/", userHandler.CreateUser)
	user.Patch("/:id", middleware.Protected(), userHandler.UpdateUser)
}