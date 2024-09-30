package beritakategori

import (
	"time"

	"gorm.io/gorm"
)

type KategoriBerita struct {
	ID        uint `gorm:"primaryKey"`
	Nama      string
	Deskripsi string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type KategoriBeritaView struct {
	ID        uint
	Index     int
	Nama      string
	Deskripsi string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
