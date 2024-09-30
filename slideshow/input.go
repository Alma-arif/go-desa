package slideshow

type ImageSlideShowInput struct {
	Judul      string `form:"judul" validate:"required"`
	Keterangan string `form:"keterangan" validate:"required"`
	Link       string `form:"url"`
	Utama      int    `form:"utama"`
}

type ImageSlideShowUpdate struct {
	ID              uint   `form:"id" validate:"required"`
	Judul           string `form:"judul" validate:"required"`
	Keterangan      string `form:"keterangan" validate:"required"`
	SlideShowImages string
	Link            string `form:"url" `
	Utama           int    `form:"utama"`
}
