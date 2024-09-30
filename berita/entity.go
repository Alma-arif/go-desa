package berita

import (
	beritaimage "app-desa-kepuk/beritaImage"
	"html/template"
	"time"

	"gorm.io/gorm"
)

type Berita struct {
	ID         uint `gorm:"primaryKey"`
	Header     string
	Judul      string
	Berita     string
	IDUser     uint
	IdKategori uint
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type BeritaView struct {
	ID         uint
	Index      int
	Header     string
	Judul      string
	Berita     template.HTML
	IDUser     uint
	UserName   string
	Image      string
	IdKategori uint
	Kategori   string
	Status     int
	CreatedAt  time.Time
	DeletedAt  time.Time
}

type BeritaViewWeb struct {
	ID          uint
	Index       int
	Header      string
	Judul       string
	Berita      template.HTML
	Image       string
	IDUser      uint
	UserName    string
	IdKategori  uint
	Kategori    string
	Status      int
	ImageBerita beritaimage.ImageFormatter
	CreatedAt   string
}
