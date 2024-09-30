package surat

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Surat struct {
	ID           uint `gorm:"primaryKey"`
	KodeSurat    uint
	NoSurat      int
	Nama         string
	Perihal      string
	Keterangan   string
	Data         datatypes.JSON
	FileLocation string
	FilePDF      string
	Path         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type SuratView struct {
	ID                  uint `gorm:"primaryKey"`
	Unix                string
	Index               int
	KodeSurat           uint
	KeteranagnKodeSurat string
	KodeSuratString     string
	NoSurat             int
	NoSuratString       string
	Nama                string
	Perihal             string
	Keterangan          string
	Data                datatypes.JSON
	FileLocation        string
	FilePDF             string
	Path                string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

type SuratKeteranganUsaha struct {
	ID                 uint `gorm:"primaryKey"`
	IdUSer             uint
	KodeSurat          string
	IDSurat            int
	Nama               string
	JenisKelamin       string
	Agama              string
	Status             string
	TempatTanggalLahir string
	Pekerjaan          string
	Keperluan          string
	JenisUsaha         string
	Keterangan         string
	TanggalMulai       time.Time
	TanggalSelesai     time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

type SuratPengantarNikah struct {
}
