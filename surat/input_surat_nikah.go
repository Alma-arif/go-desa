package surat

type InputSuratPengantarNikahNSatu struct {
	KodeSurat               uint   `json:"kode_surat" form:"kode_surat" `
	NoSurat                 int    `json:"no_surat" form:"no_surat" `
	NamaPemohon             string `json:"nama_pemohon" form:"nama_pemohon" validate:"required"`
	NikPemohon              int    `json:"nik_pemohon" form:"nik_pemohon" validate:"required"`
	JenisKelaminPemohon     string `json:"jenis_kelamin_pemohon" form:"jenis_kelamin_pemohon" validate:"required"`
	TempatLahirPemohon      string `json:"tempat_lahir_pemohon" form:"tempat_lahir_pemohon" validate:"required"`
	TanggalLahirPemohon     string `json:"tanggal_lahir_pemohon" form:"tanggal_lahir_pemohon" validate:"required"`
	KewarganegaraanPemohon  string `json:"kewarganegaraan_pemohon" form:"kewarganegaraan_pemohon" validate:"required"`
	AgamaPemohon            string `json:"agama_pemohon" form:"agama_pemohon" validate:"required"`
	PekerjaanPemohon        string `json:"pekerjaan_pemohon" form:"pekerjaan_pemohon"`
	AlamatPemohon           string `json:"alamat_pemohon" form:"alamat_pemohon" validate:"required"`
	StatusPernikahanPemohon string `json:"status_pernikahan_pemohon" form:"status_pernikahan_pemohon"`
	BeristrikePemohon       string `json:"beristrike_pemohon" form:"beristrike_pemohon"`

	NamaAyahPemohon string `json:"nama_ayah" form:"nama_ayah" validate:"required"`
	NikAyahPemohon  int    `json:"nik_ayah" form:"nik_ayah" validate:"required"`
	// JenisKelaminAyahPemohon    string `json:"jenis_kelamin_ayah" form:"jenis_kelamin_ayah" validate:"required"`
	TempatLahirAyahPemohon     string `json:"tempat_lahir_ayah" form:"tempat_lahir_ayah" validate:"required"`
	TanggalLahirAyahPemohon    string `json:"tanggal_lahir_ayah" form:"tanggal_lahir_ayah" validate:"required"`
	KewarganegaraanAyahPemohon string `json:"kewarganegaraan_ayah" form:"kewarganegaraan_ayah" validate:"required"`
	AgamaAyahPemohon           string `json:"agama_ayah" form:"agama_ayah"`
	PekerjaanAyahPemohon       string `json:"pekerjaan_ayah" form:"pekerjaan_ayah"`
	AlamatAyahPemohon          string `json:"alamat_ayah" form:"alamat_ayah"`

	NamaIbuPemohon string `json:"nama_ibu" form:"nama_ibu" validate:"required"`
	NikIbuPemohon  int    `json:"nik_ibu" form:"nik_ibu" validate:"required"`
	// JenisKelaminIbuPemohon    string `json:"jenis_kelamin_ibu" form:"jenis_kelamin_ibu" validate:"required"`
	TempatLahirIbuPemohon     string `json:"tempat_lahir_ibu" form:"tempat_lahir_ibu" validate:"required"`
	TanggalLahirIbuPemohon    string `json:"tanggal_lahir_ibu" form:"tanggal_lahir_ibu" validate:"required"`
	KewarganegaraanIbuPemohon string `json:"kewarganegaraan_ibu" form:"kewarganegaraan_ibu" validate:"required"`
	AgamaIbuPemohon           string `json:"agama_ibu" form:"agama_ibu"`
	PekerjaanIbuPemohon       string `json:"pekerjaan_ibu" form:"pekerjaan_ibu"`
	AlamatIbuPemohon          string `json:"alamat_ibu" form:"alamat_ibu"`
}

type UpdateSuratPengantarNikahNSatu struct {
	ID uint `json:"id" form:"id" validate:"required"`

	KodeSurat               uint   `json:"kode_surat" form:"kode_surat" `
	NoSurat                 int    `json:"no_surat" form:"no_surat" `
	NamaPemohon             string `json:"nama_pemohon" form:"nama_pemohon" validate:"required"`
	NikPemohon              int    `json:"nik_pemohon" form:"nik_pemohon" validate:"required"`
	JenisKelaminPemohon     string `json:"jenis_kelamin_pemohon" form:"jenis_kelamin_pemohon" validate:"required"`
	TempatLahirPemohon      string `json:"tempat_lahir_pemohon" form:"tempat_lahir_pemohon" validate:"required"`
	TanggalLahirPemohon     string `json:"tanggal_lahir_pemohon" form:"tanggal_lahir_pemohon" validate:"required"`
	KewarganegaraanPemohon  string `json:"kewarganegaraan_pemohon" form:"kewarganegaraan_pemohon" validate:"required"`
	AgamaPemohon            string `json:"agama_pemohon" form:"agama_pemohon" validate:"required"`
	PekerjaanPemohon        string `json:"pekerjaan_pemohon" form:"pekerjaan_pemohon"`
	AlamatPemohon           string `json:"alamat_pemohon" form:"alamat_pemohon" validate:"required"`
	StatusPernikahanPemohon string `json:"status_pernikahan_pemohon" form:"status_pernikahan_pemohon"`
	BeristrikePemohon       string `json:"beristrike_pemohon" form:"beristrike_pemohon"`

	NamaAyahPemohon            string `json:"nama_ayah" form:"nama_ayah" validate:"required"`
	NikAyahPemohon             int    `json:"nik_ayah" form:"nik_ayah" validate:"required"`
	TempatLahirAyahPemohon     string `json:"tempat_lahir_ayah" form:"tempat_lahir_ayah" validate:"required"`
	TanggalLahirAyahPemohon    string `json:"tanggal_lahir_ayah" form:"tanggal_lahir_ayah" validate:"required"`
	KewarganegaraanAyahPemohon string `json:"kewarganegaraan_ayah" form:"kewarganegaraan_ayah" validate:"required"`
	AgamaAyahPemohon           string `json:"agama_ayah" form:"agama_ayah"`
	PekerjaanAyahPemohon       string `json:"pekerjaan_ayah" form:"pekerjaan_ayah"`
	AlamatAyahPemohon          string `json:"alamat_ayah" form:"alamat_ayah"`

	NamaIbuPemohon            string `json:"nama_ibu" form:"nama_ibu" validate:"required"`
	NikIbuPemohon             int    `json:"nik_ibu" form:"nik_ibu" validate:"required"`
	TempatLahirIbuPemohon     string `json:"tempat_lahir_ibu" form:"tempat_lahir_ibu" validate:"required"`
	TanggalLahirIbuPemohon    string `json:"tanggal_lahir_ibu" form:"tanggal_lahir_ibu" validate:"required"`
	KewarganegaraanIbuPemohon string `json:"kewarganegaraan_ibu" form:"kewarganegaraan_ibu" validate:"required"`
	AgamaIbuPemohon           string `json:"agama_ibu" form:"agama_ibu"`
	PekerjaanIbuPemohon       string `json:"pekerjaan_ibu" form:"pekerjaan_ibu"`
	AlamatIbuPemohon          string `json:"alamat_ibu" form:"alamat_ibu"`
}

type InputSuratPengatarNikahNEmpat struct {
	KodeSurat uint `json:"kode_surat" form:"kode_surat" `
	NoSurat   int  `json:"no_surat" form:"no_surat" `

	NamaCalonSuami  string `json:"nama_calon_suami" form:"nama_calon_suami" validate:"required"`
	BintiCalonSuami string `json:"binti_calon_suami" form:"binti_calon_suami" validate:"required"`
	NikCalonSuami   string `json:"nik_calon_suami" form:"nik_calon_suami" validate:"required"`
	// JenisKelaminCalonSuami    string `json:"jenis_kelamin_calon_suami" form:"jenis_kelamin_calon_suami" validate:"required"`
	TempatLahirCalonSuami     string `json:"tempat_lahir_calon_suami" form:"tempat_lahir_calon_suami" validate:"required"`
	TanggalLahirCalonSuami    string `json:"tanggal_lahir_calon_suami" form:"tanggal_lahir_calon_suami" validate:"required"`
	KewarganegaraanCalonSuami string `json:"kewarganegaraan_calon_suami" form:"kewarganegaraan_calon_suami" validate:"required"`
	AgamaCalonSuami           string `json:"agama_calon_suami" form:"agama_calon_suami" validate:"required"`
	PekerjaanCalonSuami       string `json:"pekerjaan_calon_suami" form:"pekerjaan_calon_suami"`
	AlamatCalonSuami          string `json:"alamat_calon_suami" form:"alamat_calon_suami"`

	NamaCalonIstri  string `json:"nama_calon_istri" form:"nama_calon_istri" validate:"required"`
	BintiCalonIstri string `json:"binti_calon_istri" form:"binti_calon_istri" validate:"required"`
	NikCalonIstri   string `json:"nik_calon_istri" form:"nik_calon_istri" validate:"required"`
	// JenisKelaminCalonIstri    string `json:"jenis_kelamin_calon_istri" form:"jenis_kelamin_calon_istri" validate:"required"`
	TempatLahirCalonIstri     string `json:"tempat_lahir_calon_istri" form:"tempat_lahir_calon_istri" validate:"required"`
	TanggalLahirCalonIstri    string `json:"tanggal_lahir_calon_istri" form:"tanggal_lahir_calon_istri" validate:"required"`
	KewarganegaraanCalonIstri string `json:"kewarganegaraan_calon_istri" form:"kewarganegaraan_calon_istri" validate:"required"`
	AgamaCalonIstri           string `json:"agama_calon_istri" form:"agama_calon_istri" validate:"required"`
	PekerjaanCalonIstri       string `json:"pekerjaan_calon_istri" form:"pekerjaan_calon_istri"`
	AlamatCalonIstri          string `json:"alamat_calon_istri" form:"alamat_calon_istri"`
}

type UpdateSuratPengatarNikahNEmpat struct {
	ID uint `json:"id" form:"id" validate:"required"`

	KodeSurat uint `json:"kode_surat" form:"kode_surat" `
	NoSurat   int  `json:"no_surat" form:"no_surat" `

	NamaCalonSuami            string `json:"nama_calon_suami" form:"nama_calon_suami" validate:"required"`
	BintiCalonSuami           string `json:"binti_calon_suami" form:"binti_calon_suami" validate:"required"`
	NikCalonSuami             string `json:"nik_calon_suami" form:"nik_calon_suami" validate:"required"`
	TempatLahirCalonSuami     string `json:"tempat_lahir_calon_suami" form:"tempat_lahir_calon_suami" validate:"required"`
	TanggalLahirCalonSuami    string `json:"tanggal_lahir_calon_suami" form:"tanggal_lahir_calon_suami" validate:"required"`
	KewarganegaraanCalonSuami string `json:"kewarganegaraan_calon_suami" form:"kewarganegaraan_calon_suami" validate:"required"`
	AgamaCalonSuami           string `json:"agama_calon_suami" form:"agama_calon_suami" validate:"required"`
	PekerjaanCalonSuami       string `json:"pekerjaan_calon_suami" form:"pekerjaan_calon_suami"`
	AlamatCalonSuami          string `json:"alamat_calon_suami" form:"alamat_calon_suami"`

	NamaCalonIstri            string `json:"nama_calon_istri" form:"nama_calon_istri" validate:"required"`
	BintiCalonIstri           string `json:"binti_calon_istri" form:"binti_calon_istri" validate:"required"`
	NikCalonIstri             string `json:"nik_calon_istri" form:"nik_calon_istri" validate:"required"`
	TempatLahirCalonIstri     string `json:"tempat_lahir_calon_istri" form:"tempat_lahir_calon_istri" validate:"required"`
	TanggalLahirCalonIstri    string `json:"tanggal_lahir_calon_istri" form:"tanggal_lahir_calon_istri" validate:"required"`
	KewarganegaraanCalonIstri string `json:"kewarganegaraan_calon_istri" form:"kewarganegaraan_calon_istri" validate:"required"`
	AgamaCalonIstri           string `json:"agama_calon_istri" form:"agama_calon_istri" validate:"required"`
	PekerjaanCalonIstri       string `json:"pekerjaan_calon_istri" form:"pekerjaan_calon_istri"`
	AlamatCalonIstri          string `json:"alamat_calon_istri" form:"alamat_calon_istri"`
}

type InputSuratPengatarNikahNLima struct {
	KodeSurat uint `json:"kode_surat" form:"kode_surat" `
	NoSurat   int  `json:"no_surat" form:"no_surat" `

	StatusAyahPemohon          string `json:"status_ayah_pemohon" form:"status_ayah_pemohon" validate:"required"`
	NamaAyahPemohon            string `json:"nama_ayah" form:"nama_ayah" validate:"required"`
	BintiAyahPemohon           string `json:"binti_ayah_pemohon" form:"binti_ayah_pemohon" validate:"required"`
	NikAyahPemohon             string `json:"nik_ayah" form:"nik_ayah" validate:"required"`
	JenisKelaminAyahPemohon    string `json:"jenis_kelamin_ayah" form:"jenis_kelamin_ayah" validate:"required"`
	TempatLahirAyahPemohon     string `json:"tempat_lahir_ayah" form:"tempat_lahir_ayah" validate:"required"`
	TanggalLahirAyahPemohon    string `json:"tanggal_lahir_ayah" form:"tanggal_lahir_ayah" validate:"required"`
	KewarganegaraanAyahPemohon string `json:"kewarganegaraan_ayah" form:"kewarganegaraan_ayah" validate:"required"`
	AgamaAyahPemohon           string `json:"agama_ayah" form:"agama_ayah" validate:"required"`
	PekerjaanAyahPemohon       string `json:"pekerjaan_ayah" form:"pekerjaan_ayah"`
	AlamatAyahPemohon          string `json:"alamat_ayah" form:"alamat_ayah"`

	StatusIbuPemohon          string `json:"status_ibu_pemohon" form:"status_ibu_pemohon" validate:"required"`
	NamaIbuPemohon            string `json:"nama_ibu" form:"nama_ibu" validate:"required"`
	BintiIbuPemohon           string `json:"binti_ibu_pemohon" form:"binti_ibu_pemohon" validate:"required"`
	NikIbuPemohon             string `json:"nik_ibu" form:"nik_ibu" validate:"required"`
	JenisKelaminIbuPemohon    string `json:"jenis_kelamin_ibu" form:"jenis_kelamin_ibu" validate:"required"`
	TempatLahirIbuPemohon     string `json:"tempat_lahir_ibu" form:"tempat_lahir_ibu" validate:"required"`
	TanggalLahirIbuPemohon    string `json:"tanggal_lahir_ibu" form:"tanggal_lahir_ibu" validate:"required"`
	KewarganegaraanIbuPemohon string `json:"kewarganegaraan_ibu" form:"kewarganegaraan_ibu" validate:"required"`
	AgamaIbuPemohon           string `json:"agama_ibu" form:"agama_ibu" validate:"required"`
	PekerjaanIbuPemohon       string `json:"pekerjaan_ibu" form:"pekerjaan_ibu"`
	AlamatIbuPemohon          string `json:"alamat_ibu" form:"alamat_ibu"`

	NamaPemohon            string `json:"nama_pemohon" form:"nama_pemohon" validate:"required"`
	BintiPemohon           string `json:"binti_pemohon" form:"binti_pemohon" validate:"required"`
	NikPemohon             string `json:"nik_pemohon" form:"nik_pemohon" validate:"required"`
	JenisKelaminPemohon    string `json:"jenis_kelamin_pemohon" form:"jenis_kelamin_pemohon" validate:"required"`
	TempatLahirPemohon     string `json:"tempat_lahir_pemohon" form:"tempat_lahir_pemohon" validate:"required"`
	TanggalLahirPemohon    string `json:"tanggal_lahir_pemohon" form:"tanggal_lahir_pemohon" validate:"required"`
	KewarganegaraanPemohon string `json:"kewarganegaraan_pemohon" form:"kewarganegaraan_pemohon" validate:"required"`
	AgamaPemohon           string `json:"agama_pemohon" form:"agama_pemohon" validate:"required"`
	PekerjaanPemohon       string `json:"pekerjaan_pemohon" form:"pekerjaan_pemohon"`
	AlamatPemohon          string `json:"alamat_pemohon" form:"alamat_pemohon"`

	NamaPendaping            string `json:"nama_pendamping" form:"nama_pendamping" validate:"required"`
	BintiPendamping          string `json:"binti_pendamping" form:"binti_pendamping" validate:"required"`
	NikPendaping             string `json:"nik_pendamping" form:"nik_pendamping" validate:"required"`
	JenisKelaminPendaping    string `json:"jenis_kelamin_pendamping" form:"jenis_kelamin_pendamping" validate:"required"`
	TempatLahirPendaping     string `json:"tempat_lahir_pendamping" form:"tempat_lahir_pendamping" validate:"required"`
	TanggalLahirPendaping    string `json:"tanggal_lahir_pendamping" form:"tanggal_lahir_pendamping" validate:"required"`
	KewarganegaraanPendaping string `json:"kewarganegaraan_pendamping" form:"kewarganegaraan_pendamping" validate:"required"`
	AgamaPendaping           string `json:"agama_pendamping" form:"agama_pendamping" validate:"required"`
	PekerjaanPendaping       string `json:"pekerjaan_pendamping" form:"pekerjaan_pendamping"`
	AlamatPendaping          string `json:"alamat_pendamping" form:"alamat_pendamping"`
}

type UpdateSuratPengatarNikahNLima struct {
	ID uint `json:"id" form:"id" validate:"required"`

	KodeSurat uint `json:"kode_surat" form:"kode_surat" `
	NoSurat   int  `json:"no_surat" form:"no_surat" `

	StatusAyahPemohon          string `json:"status_ayah_pemohon" form:"status_ayah_pemohon" validate:"required"`
	NamaAyahPemohon            string `json:"nama_ayah" form:"nama_ayah" validate:"required"`
	BintiAyahPemohon           string `json:"binti_ayah_pemohon" form:"binti_ayah_pemohon" validate:"required"`
	NikAyahPemohon             string `json:"nik_ayah" form:"nik_ayah" validate:"required"`
	JenisKelaminAyahPemohon    string `json:"jenis_kelamin_ayah" form:"jenis_kelamin_ayah" validate:"required"`
	TempatLahirAyahPemohon     string `json:"tempat_lahir_ayah" form:"tempat_lahir_ayah" validate:"required"`
	TanggalLahirAyahPemohon    string `json:"tanggal_lahir_ayah" form:"tanggal_lahir_ayah" validate:"required"`
	KewarganegaraanAyahPemohon string `json:"kewarganegaraan_ayah" form:"kewarganegaraan_ayah" validate:"required"`
	AgamaAyahPemohon           string `json:"agama_ayah" form:"agama_ayah" validate:"required"`
	PekerjaanAyahPemohon       string `json:"pekerjaan_ayah" form:"pekerjaan_ayah"`
	AlamatAyahPemohon          string `json:"alamat_ayah" form:"alamat_ayah"`

	StatusIbuPemohon          string `json:"status_ibu_pemohon" form:"status_ibu_pemohon" validate:"required"`
	NamaIbuPemohon            string `json:"nama_ibu" form:"nama_ibu" validate:"required"`
	BintiIbuPemohon           string `json:"binti_ibu_pemohon" form:"binti_ibu_pemohon" validate:"required"`
	NikIbuPemohon             string `json:"nik_ibu" form:"nik_ibu" validate:"required"`
	JenisKelaminIbuPemohon    string `json:"jenis_kelamin_ibu" form:"jenis_kelamin_ibu" validate:"required"`
	TempatLahirIbuPemohon     string `json:"tempat_lahir_ibu" form:"tempat_lahir_ibu" validate:"required"`
	TanggalLahirIbuPemohon    string `json:"tanggal_lahir_ibu" form:"tanggal_lahir_ibu" validate:"required"`
	KewarganegaraanIbuPemohon string `json:"kewarganegaraan_ibu" form:"kewarganegaraan_ibu" validate:"required"`
	AgamaIbuPemohon           string `json:"agama_ibu" form:"agama_ibu" validate:"required"`
	PekerjaanIbuPemohon       string `json:"pekerjaan_ibu" form:"pekerjaan_ibu"`
	AlamatIbuPemohon          string `json:"alamat_ibu" form:"alamat_ibu"`

	NamaPemohon            string `json:"nama_pemohon" form:"nama_pemohon" validate:"required"`
	BintiPemohon           string `json:"binti_pemohon" form:"binti_pemohon" validate:"required"`
	NikPemohon             string `json:"nik_pemohon" form:"nik_pemohon" validate:"required"`
	JenisKelaminPemohon    string `json:"jenis_kelamin_pemohon" form:"jenis_kelamin_pemohon" validate:"required"`
	TempatLahirPemohon     string `json:"tempat_lahir_pemohon" form:"tempat_lahir_pemohon" validate:"required"`
	TanggalLahirPemohon    string `json:"tanggal_lahir_pemohon" form:"tanggal_lahir_pemohon" validate:"required"`
	KewarganegaraanPemohon string `json:"kewarganegaraan_pemohon" form:"kewarganegaraan_pemohon" validate:"required"`
	AgamaPemohon           string `json:"agama_pemohon" form:"agama_pemohon" validate:"required"`
	PekerjaanPemohon       string `json:"pekerjaan_pemohon" form:"pekerjaan_pemohon"`
	AlamatPemohon          string `json:"alamat_pemohon" form:"alamat_pemohon"`

	NamaPendaping            string `json:"nama_pendamping" form:"nama_pendamping" validate:"required"`
	BintiPendamping          string `json:"binti_pendamping" form:"binti_pendamping" validate:"required"`
	NikPendaping             string `json:"nik_pendamping" form:"nik_pendamping" validate:"required"`
	JenisKelaminPendaping    string `json:"jenis_kelamin_pendamping" form:"jenis_kelamin_pendamping" validate:"required"`
	TempatLahirPendaping     string `json:"tempat_lahir_pendamping" form:"tempat_lahir_pendamping" validate:"required"`
	TanggalLahirPendaping    string `json:"tanggal_lahir_pendamping" form:"tanggal_lahir_pendamping" validate:"required"`
	KewarganegaraanPendaping string `json:"kewarganegaraan_pendamping" form:"kewarganegaraan_pendamping" validate:"required"`
	AgamaPendaping           string `json:"agama_pendamping" form:"agama_pendamping" validate:"required"`
	PekerjaanPendaping       string `json:"pekerjaan_pendamping" form:"pekerjaan_pendamping"`
	AlamatPendaping          string `json:"alamat_pendamping" form:"alamat_pendamping"`
}
