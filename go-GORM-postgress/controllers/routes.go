package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, r Repository) {
	api := app.Group("/baseUrl")
	api.Post("/create_book", r.CreateBook)
	api.Put("/update_book/:id", r.UpdateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_book/:id", r.GetBookById)
	api.Get("/get_books", r.GetBooks)
}
