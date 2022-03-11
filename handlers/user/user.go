package userHandler

import (
	"github.com/ariefro/notes-server/database"
	"github.com/ariefro/notes-server/model"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	id := c.Params("userId")
	db.Find(&user, "id = ?", id)
	if user.Username == "" {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "No user found with ID" + id,
			"data": user,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "User found",
		"data": user,
	})
}