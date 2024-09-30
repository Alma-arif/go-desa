package arsip

type CreateArsipInput struct {
	Nama       string `form:"nama" validate:"required"`
	KategoriID uint   `form:"kategori_id"`
	Deskripsi  string `form:"deskripsi"`
}

type UpdateArsipInput struct {
	ID         int    `form:"id" validate:"required"`
	Nama       string `form:"nama" validate:"required"`
	KategoriID uint   `form:"kategori_id"`
	Deskripsi  string `form:"deskripsi"`
}
