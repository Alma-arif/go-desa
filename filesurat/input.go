package filesurat

type InputFileSurat struct {
	KodeSuratFix string `form:"Kode-surat-fix" validate:"required"`
	KodeSurat    string `form:"Kode-surat" validate:"required"`
	Nama         string `form:"nama" validate:"required"`
	FileMain     string
	NamaFileMain string
	File         string
	NamaFile     string
}

type UpdateFileSurat struct {
	ID        uint   `form:"id" validate:"required"`
	KodeSurat string `form:"Kode-surat" validate:"required"`
	Nama      string `form:"nama" validate:"required"`
}

type UpdateFileSuratSecone struct {
	ID       uint `form:"id" validate:"required"`
	File     string
	NamaFile string
}
