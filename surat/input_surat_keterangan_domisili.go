package surat

type InputSuratKeteranganDomisiliUsaha struct {
	KodeSurat       string `json:"kode_surat" form:"kode_surat" `
	NoSurat         string `json:"no_surat" form:"no_surat" `
	Nama            string `json:"nama_pemohon" form:"nama_pemohon" validate:"required"`
	NIK             string `json:"nik" form:"nik" validate:"required"`
	JenisKelamin    string `json:"jenis_kelamin" form:"jenis_kelamin" validate:"required"`
	Alamat          string `json:"alamat" form:"alamat" validate:"required"`
	HariMeninggal   string `json:"hari_meninggal" form:"hari_meninggal" validate:"required"`
	TanggalMeinggal string `json:"tanggal_meninggal" form:"tanggal_meninggal" validate:"required"`
	TempatMeninggal string `json:"tempat_meninggal" form:"tempat_meninggal" validate:"required"`
	Penyebab        string `json:"penyebab" form:"penyebab" validate:"required"`
}
