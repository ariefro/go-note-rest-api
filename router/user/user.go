package userRoute

import (
	userHandler "github.com/ariefro/go-note-rest-api/handlers/user"
	"github.com/ariefro/go-note-rest-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Get("/:userId", userHandler.GetUser)
	user.Post("/", userHandler.CreateUser)
	user.Patch("/:id", middleware.Protected(), userHandler.UpdateUser)
}
