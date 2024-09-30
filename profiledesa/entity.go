package profiledesa

import (
	"html/template"
	"time"

	"gorm.io/gorm"
)

type ProfileDesa struct {
	ID          uint `gorm:"primaryKey"`
	ProfileDesa string
	ImageDesa   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ProfileDesaView struct {
	ID          uint
	Index       int
	ProfileDesa template.HTML
	ImageDesa   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
