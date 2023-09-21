package middlewares


import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)


func AuthMiddleware(c *fiber.Ctx) error {
	// Check if the user is authenticated (you can implement your logic here)
	authenticated := true // Replace with your authentication logic
	if !authenticated {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	return c.Next() // Continue to the next middleware or route handler
}