package beritaimage

type ImageBeritaInput struct {
	NamaImageFile string
	IdBerita      uint `form:"id_berita" validate:"required"`
	ImageUtama    int  `form:"image_utama"`
}

type ImageBeritaUpdate struct {
	ID            uint `form:"id" validate:"required"`
	NamaImageFile string
	IdBerita      uint `form:"id_berita" validate:"required"`
	ImageUtama    int  `form:"image_utama"`
}
