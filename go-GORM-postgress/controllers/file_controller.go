package controllers

import (
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Define a struct to represent the file model in the database
type File struct {
	ID        string `gorm:"primaryKey"`
	FileName  string `gorm:"not null"`
	FileData  []byte `gorm:"type:bytea"`
	FileType  string `gorm:"not null"`
	CreatedAt time.Time
}

func (r *Repository) uploadFile(c *fiber.Ctx) error {
	// Parse the multipart form data
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse form data",
		})
	}

	// Process each uploaded file
	files := form.File["files"]
	fileID := generateFileID()

	for _, file := range files {
		// Generate a unique file name
		fileName := filepath.Base(file.Filename)
		fileType := filepath.Ext(fileName)

		// Read the file into memory
		fileContent, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to read file",
				"file":    fileName,
			})
		}
		defer fileContent.Close()

		// Convert the file content to a byte slice
		fileData, err := ioutil.ReadAll(fileContent)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to read file content",
				"file":    fileName,
			})
		}

		// Create a new file record in the database with the file content
		newFile := File{
			ID:       fileID,
			FileName: fileName,
			FileData: fileData,
			FileType: fileType,
		}
		if err := r.DB.Create(&newFile).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to insert file into the database",
				"file":    fileName,
			})
		}
	}

	return c.JSON(fiber.Map{
		"message": "Files uploaded successfully",
		"fileId":  fileID,
	})
}

func generateFileID() string {
	// Generate a unique file ID based on the current time and a random string
	currentTime := time.Now()
	uniqueID := currentTime.Format("20060102150405") + "_" + randomString(6)
	return uniqueID
}

// Function to generate a random string of a given length
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (r *Repository) getFile(c *fiber.Ctx) error {
	// Get the fileId from the request parameters
	fileID := c.Params("fileId")

	// Query the database for the file with the specified fileId
	var file File
	result := r.DB.Where("id = ?", fileID).First(&file)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "File not found",
		})
	}

	// Set the response headers to indicate the file download
	c.Set("Content-Disposition", "attachment; filename="+file.FileName)
	c.Set("Content-Type", file.FileType)
	c.Send(file.FileData)

	return nil
}
