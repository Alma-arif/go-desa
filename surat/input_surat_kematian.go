package surat

type InputSuratKeteranganMeninggal struct {
	KodeSurat        uint   `json:"kode_surat" form:"kode_surat" `
	NoSurat          int    `json:"no_surat" form:"no_surat" `
	Nama             string `json:"nama" form:"nama" validate:"required"`
	NIK              string `json:"nik" form:"nik" validate:"required"`
	JenisKelamin     string `json:"jenis_kelamin" form:"jenis_kelamin"`
	Alamat           string `json:"alamat" form:"alamat"`
	HariMeninggal    string `json:"hari_meninggal" form:"hari_meninggal"`
	TanggalMeninggal string `json:"tanggal_meninggal" form:"tanggal_meninggal"`
	TempatMeninggal  string `json:"tempat_meninggal" form:"tempat_meninggal"`
	Penyebab         string `json:"penyebab" form:"penyebab"`
}

type UpdateSuratKeteranganMeninggal struct {
	ID               uint   `json:"id" form:"id" validate:"required"`
	KodeSurat        uint   `json:"kode_surat" form:"kode_surat" validate:"required"`
	NoSurat          int    `json:"no_surat" form:"no_surat" validate:"required"`
	Nama             string `json:"nama" form:"nama" validate:"required"`
	NIK              string `json:"nik" form:"nik" validate:"required"`
	JenisKelamin     string `json:"jenis_kelamin" form:"jenis_kelamin"`
	Alamat           string `json:"alamat" form:"alamat"`
	HariMeninggal    string `json:"hari_meninggal" form:"hari_meninggal"`
	TanggalMeninggal string `json:"tanggal_meninggal" form:"tanggal_meninggal"`
	TempatMeninggal  string `json:"tempat_meninggal" form:"tempat_meninggal"`
	Penyebab         string `json:"penyebab" form:"penyebab"`
}
