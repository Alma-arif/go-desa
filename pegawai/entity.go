package pegawai

import (
	"time"

	"gorm.io/gorm"
)

type Pegawai struct {
	ID           uint `gorm:"primaryKey"`
	Nama         string
	Jabatan      string
	NoHP         string
	Alamat       string
	Image        string
	TanggalLahir time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type PegawaiView struct {
	ID           uint
	Index        int
	Nama         string
	Jabatan      string
	NoHP         string
	Alamat       string
	Image        string
	TanggalLahir time.Time
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
