package arsipkategori

type KategoriArsipInput struct {
	Nama      string `form:"nama" validate:"required"`
	Deskripsi string `form:"deskripsi" `
}

type KategoriArsipUpdate struct {
	ID        uint   `form:"id" validate:"required"`
	Nama      string `form:"nama" validate:"required"`
	Deskripsi string `form:"deskripsi" `
}
