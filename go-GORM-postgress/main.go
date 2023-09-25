package main

import (
	"go-gorm-postgres/config"
	"go-gorm-postgres/controllers"
	"go-gorm-postgres/middlewares"
	"go-gorm-postgres/models"
	"go-gorm-postgres/storage"
	"log"

	"os"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	Author    string `json:"author"` // json:"author" is specified to understand the go the json
	Title     string `json:"title"`
	Price     int    `json:"price"`
	Publisher string `json:"publisher"`
}

func main() {
	err := config.LoadEnv()

	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSLMode"),
	}

	db, err := storage.NewConnection(config) // initialising the db connection

	if err != nil {
		log.Fatal("Failed to load DB")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("Couldn't migrate DB") // create or update the database schema for a table corresponding to the Books{} struct
	} //  convenient way to ensure that the database schema for the Books table is in sync with the corresponding Go struct.

	app := fiber.New()
	app.Use(middlewares.AuthMiddleware) // will work for the entire application
	r := controllers.Repository{DB: db}
	controllers.SetupRoutes(app, r)
	app.Listen(":8080")
}

// router
