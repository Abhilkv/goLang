package models

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	ID        uint   `gorm:"primary key;autoIncrement" json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Price     int    `json:"price"`
	Publisher string `json:"publisher"`
}

type File struct {
	ID        string `gorm:"primaryKey"`
	FileName  string `gorm:"not null"`
	FileData  []byte `gorm:"type:bytea"`
	FileType  string `gorm:"not null"`
	CreatedAt time.Time
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{}, &File{})
	return err
}
