package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func LogMiddleware(c *fiber.Ctx) error {
	// Log information about the request (you can customize this)
	log.Printf("Request: %s %s", c.Method(), c.OriginalURL())
	return c.Next() // Continue to the next middleware or route handler
}

func RouteMiddleWare(c *fiber.Ctx) error {
	// Log information about the request (you can customize this)
	log.Printf("Request: %s %s", c.Method(), c.OriginalURL())
	return c.Next() // Continue to the next middleware or route handler
}



func SetupRoutes(app *fiber.App, r Repository) {
	api := app.Group("/baseUrl")
	api.Use(LogMiddleware)  // will work for all api under /baseUrl
	api.Post("/create_book", r.CreateBook)
	api.Put("/update_book/:id", r.UpdateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_book/:id", RouteMiddleWare,  r.GetBookById)  // will work only for the specific route 
	api.Get("/get_books", r.GetBooks)
}
