package demografi

import (
	"html/template"
	"time"

	"gorm.io/gorm"
)

type Demografi struct {
	ID        uint `gorm:"primaryKey"`
	Demografi string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type DemografiView struct {
	ID        uint
	Index     int
	Demografi template.HTML
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
