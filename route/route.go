package route

import (
	"api/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRoute(noteController *controller.NoteController) *fiber.App {
	route := fiber.New()

	route.Route("/notes", func(router fiber.Router) {
		router.Post("/", noteController.Create)
		router.Get("/", noteController.FindAll)
	})

	route.Route("/notes/:noteId", func(router fiber.Router) {
		router.Put("", noteController.Update)
		router.Get("", noteController.FindById)
		router.Delete("", noteController.Delete)
	})

	return route
}
