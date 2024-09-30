package berita

type BeritaInput struct {
	Judul      string `form:"nama" validate:"required"`
	Berita     string `form:"berita" validate:"required"`
	IdUSer     uint
	IdKategori uint `form:"Kategori_id" validate:"required"`
	Status     int  `form:"status" `
}

type BeritaUpdate struct {
	ID         uint   `form:"id" validate:"required"`
	Judul      string `form:"nama" validate:"required"`
	Berita     string `form:"berita" validate:"required"`
	IdKategori uint   `form:"Kategori_id" `
	idUSer     uint
	Status     int `form:"status" `
}
