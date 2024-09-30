package visimisidesa

import (
	"html/template"
	"time"

	"gorm.io/gorm"
)

type VisiMisi struct {
	ID        uint `gorm:"primaryKey"`
	VisiMisi  string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type VisiMisiView struct {
	ID        uint
	Index     int
	VisiMisi  template.HTML
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
