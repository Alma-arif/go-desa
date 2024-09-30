package filesurat

import (
	"time"

	"gorm.io/gorm"
)

type FileSurat struct {
	ID           uint `gorm:"primaryKey"`
	KodeSuratFix string
	KodeSurat    string
	Nama         string
	FileMain     string
	NamaFileMain string
	File         string
	NamaFile     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type FileSuratView struct {
	ID           uint
	Index        int
	KodeSuratFix string
	KodeSurat    string
	Nama         string
	FileMain     string
	NamaFileMain string
	File         string
	NamaFile     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
