package pengumuman

import (
	"html/template"
	"time"

	"gorm.io/gorm"
)

type Pengumuman struct {
	ID         uint `gorm:"primaryKey"`
	Header     string
	Judul      string
	Pengumuman string
	IDUser     uint
	Image      string
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type PengumumanView struct {
	ID         uint
	Index      int
	Header     string
	Judul      string
	Pengumuman template.HTML
	IDUser     uint
	Username   string
	Image      string
	Status     int
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type PengumumanViewWeb struct {
	ID         uint
	Index      int
	Header     string
	Judul      string
	Pengumuman template.HTML
	IDUser     uint
	Username   string
	Image      string
	Status     int
	CreatedAt  string
}
