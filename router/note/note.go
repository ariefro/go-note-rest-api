package noteRoutes

import (
	noteHandler "github.com/ariefro/go-note-rest-api/handlers/note"
	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")

	note.Post("/", noteHandler.CreateNote)
	note.Get("/", noteHandler.GetNotes)
	note.Get("/:noteId", noteHandler.GetNote)
	note.Put("/:noteId", noteHandler.UpdateNote)
	note.Delete("/:noteId", noteHandler.DeleteNote)
}
