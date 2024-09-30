package arsip

import (
	"time"

	"gorm.io/gorm"
)

type ArsipDesa struct {
	ID         uint `gorm:"primaryKey"`
	Nama       string
	KategoriID uint
	Deskripsi  string
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type ArsipList struct {
	ID         uint
	Index      int
	Nama       string
	Kategori   string
	KategoriID uint
	Deskripsi  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type arsipFile struct {
	ID           uint
	Index        int
	ArsipID      uint
	NamaFile     string
	FileLocation string
	FileSize     float64
	Status       int
}
