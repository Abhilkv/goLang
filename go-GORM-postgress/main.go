package main

import (
	"fmt"
	"go-gorm-postgres/models"
	"go-gorm-postgres/storage"
	"log"
	"net/http"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"` // json:"author" is specified to understand the go the json
	Title     string `json:"title"`
	Price     int    `json:"price"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

// CREATE
func (r *Repository) CreateBook(context *fiber.Ctx) error {
	fmt.Println("create")
	book := Book{}

	err := context.BodyParser(&book) // convert the request to json using the struct book

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Couldn't create book"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "New entry created"})
	return nil
}

// GET all
func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}
	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Bad Request"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Fetched successfully", "data": bookModels})

	return nil
}

func (r *Repository) GetBookById(context *fiber.Ctx) error {
	bookModels := &models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Id is empty"})
		return nil
	}
	err := r.DB.Where("id = ?", id).Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Bad Request"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Fetched successfully", "data": bookModels})

	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModels := models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Id is empty"})
		return nil
	}
	err := r.DB.Delete(bookModels, id).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Failed to delete"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Deleted successfully", "data": bookModels})

	return nil
}

func (r *Repository) UpdateBook(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Id is empty"})
		return nil
	}

	book := Book{}
	err := context.BodyParser(&book)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}

	err = r.DB.Model(&models.Books{}).Where("id = ?", id).Updates(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Couldn't update book"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Book updated successfully"})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	fmt.Println("routes")
	api := app.Group("/baseUrl") // all the api should start with /baseUrl string
	api.Post("/create_book", r.CreateBook)
	api.Put("/update_book/:id", r.UpdateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_book/:id", r.GetBookById)
	api.Get("/get_books", r.GetBooks)

}

func main() {
	err := godotenv.Load(".env")

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
	r := Repository{
		DB: db, // initializing the DB
	}

	if err != nil {
		log.Fatal("Failed to load DB")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("Couldn't migrate DB")
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
