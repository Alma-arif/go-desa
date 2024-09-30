package pegawai

type PegawaiInput struct {
	Nama         string `form:"nama" validate:"required"`
	Jabatan      string `form:"jabatan" validate:"required"`
	NoHP         string `form:"no_hp"`
	Alamat       string `form:"alamat"`
	TanggalLahir string `form:"tanggal_lahir" validate:"required"`
}

type PegawaiUpdate struct {
	ID           uint   `form:"id" validate:"required"`
	Nama         string `form:"nama" validate:"required"`
	Jabatan      string `form:"jabatan" validate:"required"`
	NoHP         string `form:"no_hp" `
	Alamat       string `form:"alamat"`
	TanggalLahir string `form:"tanggal_lahir" validate:"required"`
}
