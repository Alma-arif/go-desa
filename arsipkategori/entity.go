package arsipkategori

import (
	"time"

	"gorm.io/gorm"
)

type KategoriArsip struct {
	ID        uint `gorm:"primaryKey"`
	Nama      string
	Deskripsi string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type KategoriArsipView struct {
	ID        uint
	Index     int
	Nama      string
	Deskripsi string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
