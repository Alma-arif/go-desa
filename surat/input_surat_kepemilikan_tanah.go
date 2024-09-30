package surat

type InputSuratKepemilikanTanah struct {
	KodeSurat              uint   `json:"kode_surat" form:"kode_surat" validate:"required"`
	NoSurat                int    `json:"no_surat" form:"no_surat" `
	NamaPemilik            string `json:"nama_pemohon" form:"nama_pemohon" validate:"required"`
	NIKPemilik             string `json:"nik" form:"nik" validate:"required"`
	TempatLahirPemilik     string `json:"tempat_lahir" form:"tempat_lahir" `
	TanggalLahirPemilik    string `json:"tanggal_lahir" form:"tanggal_lahir" `
	UmurPemilik            string `json:"umur_negaraan" form:"umur_negaraan" `
	AgamaPemilik           string `json:"agama" form:"agama" `
	StatusPenikahanPemilik string `json:"status_nikah" form:"status_nikah" `
	PekerjaanPemilik       string `json:"pekerjaan" form:"pekerjaan" `
	AlamatPemilik          string `json:"alamat" form:"alamat" `

	LuasTanah              string `json:"luas_tanah" form:"luas_tanah" `
	AlamatTanah            string `json:"alamat_tanah" form:"alamat_tanah" `
	TanahHasil             string `json:"tanah_hasil" form:"tanah_hasil" `
	NomerSertifikatTanah   string `json:"nomer_sertifikat_tanah" form:"nomer_sertifikat_tanah" `
	TanggalSertifikatTanah string `json:"tanggal_sertifikat_tanah" form:"tanggal_sertifikat_tanah" `
	InstansiPeresmi        string `json:"insyansi_pemerintah" form:"insyansi_pemerintah" `
	Wilayah                string `json:"wilayah" form:"wilayah" `

	BatasTanahUtara   string `json:"batas_tanah_utara" form:"batas_tanah_utara" `
	BatasTanahSeletan string `json:"batas_tanah_seletan" form:"batas_tanah_seletan" `
	BatasTanahTimur   string `json:"batas_tanah_timur" form:"batas_tanah_timur" `
	BatasTanahBarat   string `json:"batas_tanah_barat" form:"batas_tanah_barat" `

	NomerSertifikatTanahInstansi   string `json:"nomer_sertifikat_tanah_instansi" form:"nomer_sertifikat_tanah_instansi" `
	TanggalSertifikatTanahInstansi string `json:"tanggal_sertifikat_tanah_instansi" form:"tanggal_sertifikat_tanah_instansi" `

	SaksiBatasTanahUtara   string `json:"saksi_batas_tanah_utara" form:"saksi_batas_tanah_utara" `
	SaksiBatasTanahSeletan string `json:"saksi_batas_tanah_seletan" form:"saksi_batas_tanah_seletan" `
	SaksiBatasTanahTimur   string `json:"saksi_batas_tanah_timur" form:"saksi_batas_tanah_timur" `
	SaksiBatasTanahBarat   string `json:"saksi_batas_tanah_barat" form:"saksi_batas_tanah_barat" `
}
