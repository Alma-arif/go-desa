package beritakategori

type KategoriBeritaInput struct {
	Nama      string `form:"nama" validate:"required"`
	Deskripsi string `form:"deskripsi" `
	// Status    int    `form:"status" `
}

type KategoriBeritaUpdate struct {
	ID        uint   `form:"id" validate:"required"`
	Nama      string `form:"nama" validate:"required"`
	Deskripsi string `form:"deskripsi" `
	// Status    int    `form:"status" `
}
