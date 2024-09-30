package sejarah

import (
	"html/template"
	"time"

	"gorm.io/gorm"
)

type Sejarah struct {
	ID        uint `gorm:"primaryKey"`
	Sejarah   string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type SejarahView struct {
	ID        uint
	Index     int
	Sejarah   template.HTML
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
