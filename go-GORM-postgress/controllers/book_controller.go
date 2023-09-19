package controllers

import (
	"fmt"
	"go-gorm-postgres/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
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
		&fiber.Map{"message": "Deleted successfully"})

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
