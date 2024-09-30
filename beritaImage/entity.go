package beritaimage

import (
	"time"
)

type ImageBerita struct {
	ID            uint `gorm:"primaryKey"`
	NamaImageFile string
	IdBerita      uint
	ImageUtama    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	// DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type ImageBeritaView struct {
	ID            uint
	Index         int
	NamaImageFile string
	IdBerita      uint
	ImageUtama    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	// DeletedAt     gorm.DeletedAt
}
