package surat

type InputSuratKeteranganUsaha struct {
	KodeSurat       uint   `json:"kode_surat" form:"kode_surat" validate:"required"`
	NoSurat         int    `json:"no_surat" form:"no_surat" `
	Nama            string `json:"nama_pemohon" form:"nama_pemohon" validate:"required"`
	NIK             string `json:"nik" form:"nik" validate:"required"`
	TempatLahir     string `json:"tempat_lahir" form:"tempat_lahir" `
	TanggalLahir    string `json:"tanggal_lahir" form:"tanggal_lahir" `
	KewargaNegaraan string `json:"kewarga_negaraan" form:"kewarga_negaraan" `
	Agama           string `json:"agama" form:"agama" `
	StatusPenikahan string `json:"status_nikah" form:"status_nikah" `
	Pekerjaan       string `json:"pekerjaan" form:"pekerjaan" `
	Alamat          string `json:"alamat" form:"alamat" `
	JenisUsaha      string `json:"jenis_usaha" form:"jenis_usaha" `
}

type UpdateSuratKeteranganUsaha struct {
	ID              uint   `json:"id" form:"id" validate:"required"`
	KodeSurat       uint   `json:"kode_surat" form:"kode_surat" validate:"required"`
	NoSurat         int    `json:"no_surat" form:"no_surat"`
	Nama            string `json:"nama_pemohon" form:"nama_pemohon" validate:"required"`
	NIK             string `json:"nik" form:"nik" validate:"required"`
	TempatLahir     string `json:"tempat_lahir" form:"tempat_lahir" `
	TanggalLahir    string `json:"tanggal_lahir" form:"tanggal_lahir" `
	KewargaNegaraan string `json:"kewarga_negaraan" form:"kewarga_negaraan" `
	Agama           string `json:"agama" form:"agama" `
	StatusPenikahan string `json:"status_nikah" form:"status_nikah" `
	Pekerjaan       string `json:"pekerjaan" form:"pekerjaan" `
	Alamat          string `json:"alamat" form:"alamat" `
	JenisUsaha      string `json:"jenis_usaha" form:"jenis_usaha" `
}

// type InputSuratKeteranganAhliwaris struct {
// 	// IdUSer             uint
// 	KodeSurat       string `json:"kode_surat" form:"kode_surat" `
// 	NoSurat         string `json:"no_surat" form:"no_surat" `
// 	NamaAlihWaris   string `json:"nama_alih_waris" form:"nama_alih_waris"`
// 	NamaPemohon     string `json:"nama_pemohon" form:"nama_pemohon" validate:"required"`
// 	TempatLahir     string `json:"tempat_lahir" form:"tempat_lahir" validate:"required"`
// 	TanggalLahir    string `json:"tanggal_lahir" form:"tanggal_lahir" validate:"required"`
// 	TanggalMeinggal string `json:"tanggal_meninggal" form:"tanggal_meninggal" validate:"required"`
// 	TempatMeninggal string `json:"tempat_meninggal" form:"tempat_meninggal" validate:"required"`
// 	Keterangan      string `json:"keterangan" form:"Keterangan" validate:"required"`
// }

// type InputSuratKeteranganDomisili struct {
// 	KodeSurat string `json:"kode_surat" form:"kode_surat" `
// 	NoSurat   string `json:"no_surat" form:"no_surat" `
// 	NamaPemohon            string `json:"nama_pemohon" form:"nama_pemohon" validate:"required"`
// 	JenisKelamin           string `json:"jenis_kelamin" form:"jenis_kelamin" validate:"requirTime
// 	Agama                  string `json:"agama" form:"agama" validate:"required"`
// 	AlamatPemohon          string `json:"alamat_pemohon" form:"alamat_pemohon" validate:"required"`
// 	NIK                    string `json:"nik" form:"nik" validate:"required"`
// 	TempatLahir            string `json:"tempat_lahir" form:"tempat_lahir" validate:"required"`
// 	TanggalLahir           string `json:"tanggal_lahir" form:"tanggal_lahir" validate:"required"`
// 	Pekerjaan              string `json:"pekerjaan" form:"pekerjaan" validate:"required"`
// 	StatusPenikahan        string `json:"status_nikah" form:"status_nikah" validate:"required"`
// 	KewarganegaraanPemohon string `json:"kewarganegaraan_pemohon" form:"kewarganegaraan_pemohon" validate:"required"`

// 	RT int `json:"rt" form:"rt" validate:"required"`
// 	RW int `json:"rw" form:"rw" validate:"required"`

// 	Keterangan string `json:"keterangan" form:"Keterangan" validate:"required"`
// }

type InputSuratKeteranganKTPSementara struct {
	// IdUSer             uint
	KodeSurat string `json:"kode_surat" form:"kode_surat" `
	NoSurat   string `json:"no_surat" form:"no_surat" `

	NamaPemohon   string `json:"nama_pemohon" form:"nama_pemohon" validate:"required"`
	JenisKelamin  string `json:"jenis_kelamin" form:"jenis_kelamin" validate:"required"`
	Agama         string `json:"agama" form:"agama" validate:"required"`
	AlamatPemohon string `json:"alamat_pemohon" form:"alamat_pemohon" validate:"required"`
	NIK           string `json:"nik" form:"nik" validate:"required"`
	TempatLahir   string `json:"tempat_lahir" form:"tempat_lahir" validate:"required"`
	TanggalLahir  string `json:"tanggal_lahir" form:"tanggal_lahir" validate:"required"`
	Pekerjaan     string `json:"pekerjaan" form:"pekerjaan" validate:"required"`
	// StatusPenikahan        string `json:"status_nikah" form:"status_nikah" validate:"required"`
	KewarganegaraanPemohon string `json:"kewarganegaraan_pemohon" form:"kewarganegaraan_pemohon" validate:"required"`
	Keterangan             string `json:"keterangan" form:"Keterangan" validate:"required"`
}

type InputSuratKeteranganKelahiran struct {
	// IdUSer             uint
	KodeSurat string `json:"kode_surat" form:"kode_surat" `
	NoSurat   string `json:"no_surat" form:"no_surat" `

	Nama         string `json:"nama" form:"nama" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin" validate:"required"`
	Agama        string `json:"agama" form:"agama" validate:"required"`
	Alamat       string `json:"alamat" form:"alamat" validate:"required"`
	TempatLahir  string `json:"tempat_lahir" form:"tempat_lahir" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" form:"tanggal_lahir" validate:"required"`

	NamaAyah string `json:"nama_ayah" form:"nama_ayah" validate:"required"`
	NamaIbu  string `json:"nama_ibu" form:"nama_ibu" validate:"required"`

	Keterangan string `json:"keterangan" form:"Keterangan" validate:"required"`
}
