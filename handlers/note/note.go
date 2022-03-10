package noteHandler

import (
	"github.com/ariefro/notes-server/database"
	"github.com/ariefro/notes-server/model"
	"github.com/gofiber/fiber/v2"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	db.Find(&notes)

	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "No notes present",
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Notes Found",
		"data": notes,
	})
}