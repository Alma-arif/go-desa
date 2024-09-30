package slideshow

import (
	"time"

	"gorm.io/gorm"
)

type ImageSlideShow struct {
	ID              uint `gorm:"primaryKey"`
	Judul           string
	Keterangan      string
	SlideShowImages string
	Link            string
	Utama           int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type ImageSlideShowView struct {
	ID              uint
	Index           int
	Judul           string
	Keterangan      string
	SlideShowImages string
	Link            string
	Utama           int
	CreatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
