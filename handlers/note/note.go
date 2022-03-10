package noteHandler

import (
	"github.com/ariefro/notes-server/database"
	"github.com/ariefro/notes-server/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Review your input",
			"data": err,
		})
	}

	note.ID = uuid.New()
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Could not create note",
			"data": err,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Created note",
		"data": note,
	})
}