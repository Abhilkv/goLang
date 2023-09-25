package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Check if the user is authenticated (you can implement your logic here)
	// token := c.Get("Authorization")
	// log.Println("Auth middlware")
	// userAgent := c.Get("User-Agent")
	// paramValue := c.Query("paramKey")
	// username := c.FormValue("username")
	// id = c.Params("id")
	// if err := c.BodyParser(&data); err != nil {
	//     return err
	// }
	authenticated := true // Replace with your authentication logic
	if !authenticated {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	return c.Next() // Continue to the next middleware or route handler
}
