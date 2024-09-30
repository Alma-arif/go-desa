package pengumuman

type PengumumanInput struct {
	Judul      string `form:"judul" validate:"required"`
	Pengumuman string `form:"pengumuman" validate:"required"`
	IDUser     uint
	Status     int `form:"status"`
}

type PengumumanUpdate struct {
	ID         uint   `form:"id" validate:"required"`
	Judul      string `form:"judul" validate:"required"`
	Pengumuman string `form:"pengumuman" validate:"required"`
	IDUser     uint
	Status     int `form:"status"`
}
