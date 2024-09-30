package file

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID           uint `gorm:"primaryKey"`
	ArsipID      uint
	NamaFile     string
	Deskripsi    string
	FileLocation string
	FileSize     float64
	Status       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type FileView struct {
	ID           uint
	Index        int
	ArsipID      uint
	NamaFile     string
	Deskripsi    string
	FileLocation string
	FileSize     float64
	Status       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
