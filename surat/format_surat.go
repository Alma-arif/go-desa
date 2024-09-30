package surat

import (
	"app-desa-kepuk/helper"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lukasjarosch/go-docx"
	DocxReplace "github.com/nguyenthenguyen/docx"
)

func GenerateSuratName(data Surat) (string, string) {
	unixtime := time.Now().UnixNano()
	unixtimeStr := strconv.FormatInt(unixtime, 10)
	formattedDate := data.CreatedAt.Format("02-01-2006")
	fileName := fmt.Sprint(data.NoSurat, "-", data.KodeSurat, "-", data.Nama, "-", formattedDate, "-", unixtimeStr, ".docx")
	filefix := helper.StringWithoutSpaces(fileName)
	derektory := fmt.Sprint("./derektori/surat/file_surat/", filefix)
	return filefix, derektory
}

func SuratUsaha(data Surat, kodeSurat string, fileTempateMain string) (string, error) {
	var person InputSuratKeteranganUsaha
	fmt.Println("c1 :")

	if fileTempateMain == "" {
		fmt.Println("c2 :")
		return "", errors.New("File tempate Surat tidak di temukan!")
	}

	err := json.Unmarshal(data.Data, &person)
	if err != nil {
		fmt.Println("c3 :", err)

		return "", err
	}

	dateTimeNow, _ := helper.StringToDateSepesific(person.TanggalLahir)
	ttl, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahir, dateTimeNow)
	date := time.Now()
	timeindo, _ := helper.IndonesiaFormat(date)

	replaceMap := docx.PlaceholderMap{
		"tahun":                date.Year(),
		"kode-surat":           kodeSurat,
		"no-surat":             helper.NomerSurat(data.NoSurat),
		"nama":                 person.Nama,
		"agama":                person.Agama,
		"status":               person.StatusPenikahan,
		"nik":                  person.NIK,
		"tempat-tanggal-lahir": strings.TrimSpace(ttl),
		"pekerjaan":            person.Pekerjaan,
		"kewarganegaraan":      person.KewargaNegaraan,
		"jenis-usaha":          person.JenisUsaha,
		"alamat":               person.Alamat,
		"tanggal-dibuat":       timeindo,
	}

	fileTempate := fmt.Sprintf("derektori/surat/template/%s", fileTempateMain)
	doc, err := docx.Open(fileTempate)
	if err != nil {
		fmt.Println("c4 :", err)

		return "", err
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		fmt.Println("c5 :", err)

		return "", err
	}

	NameSurat, path := GenerateSuratName(data)
	err = doc.WriteToFile(path)
	if err != nil {
		fmt.Println("c6 :", err)

		return "", err
	}

	fmt.Println("c7 : end")

	return NameSurat, nil
}

func SuratUsahaRiplace(data Surat, kodeSurat string, fileTempateMain string) (string, error) {
	var person InputSuratKeteranganUsaha
	fmt.Println("c1 :")

	if fileTempateMain == "" {
		fmt.Println("c2 :")
		return "", errors.New("File tempate Surat tidak di temukan!")
	}

	err := json.Unmarshal(data.Data, &person)
	if err != nil {
		fmt.Println("c3 :", err)

		return "", err
	}

	dateTimeNow, _ := helper.StringToDateSepesific(person.TanggalLahir)
	ttl, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahir, dateTimeNow)
	date := time.Now()
	timeindo, _ := helper.IndonesiaFormat(date)

	fileTempate := fmt.Sprintf("./derektori/surat/template/%s", fileTempateMain)
	r, err := DocxReplace.ReadDocxFile(fileTempate)
	if err != nil {
		return "", err
	}
	docx1 := r.Editable()
	docx1.Replace("no-surat", helper.NomerSurat(data.NoSurat), -1)
	docx1.Replace("kode-surat", kodeSurat, -1)
	docx1.Replace("tahun", fmt.Sprint(date.Year()), -1)
	docx1.Replace("tanggal-dibuat", timeindo, -1)

	docx1.Replace("nama", person.Nama, -1)
	docx1.Replace("agama", person.Agama, -1)
	docx1.Replace("status-pernikahan", person.StatusPenikahan, -1)
	docx1.Replace("nik", person.NIK, -1)
	docx1.Replace("tempat-tanggal-lahir", strings.TrimSpace(ttl), -1)
	docx1.Replace("pekerjaan", person.Pekerjaan, -1)
	docx1.Replace("kewarganegaraan", person.KewargaNegaraan, -1)
	docx1.Replace("jenis-usaha", person.JenisUsaha, -1)
	docx1.Replace("alamat", person.Alamat, -1)

	NameSurat, path := GenerateSuratName(data)
	err = docx1.WriteToFile(path)
	if err != nil {
		return "", err
	}

	r.Close()

	return NameSurat, nil
}

func SuratKeteranagnKematian(data Surat, kodeSurat string, fileTempateMain string) (string, error) {
	var person InputSuratKeteranganMeninggal

	if fileTempateMain == "" {
		return "", errors.New("File tempate Surat tidak di temukan!")
	}

	err := json.Unmarshal(data.Data, &person)
	if err != nil {
		return "", err
	}

	dateTimeNow, err := helper.StringToDateSepesific(person.TanggalMeninggal)
	tanggalIndo, _ := helper.DateToDateIndoFormat(dateTimeNow)
	date := time.Now()
	timeindo, err := helper.IndonesiaFormat(date)

	replaceMap := docx.PlaceholderMap{
		"tahun":             date.Year(),
		"kode-surat":        kodeSurat,
		"no-surat":          helper.NomerSurat(data.NoSurat),
		"nama":              person.Nama,
		"nik":               person.NIK,
		"alamat":            person.Alamat,
		"jenis-kelamin":     person.JenisKelamin,
		"hari-meninggal":    person.HariMeninggal,
		"tanggal-meninggal": tanggalIndo,
		"tempat-meninggal":  person.TempatMeninggal,
		"penyebab":          person.Penyebab,
		"tanggal-dibuat":    timeindo,
	}

	fileTempate := fmt.Sprintf("./derektori/surat/template/%s", fileTempateMain)
	doc, err := docx.Open(fileTempate)
	if err != nil {
		return "", err
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		return "", err
	}

	NameSurat, path := GenerateSuratName(data)
	err = doc.WriteToFile(path)
	if err != nil {
		return "", err
	}
	return NameSurat, nil
}

func SuratKeteranagnKematianRiplace(data Surat, kodeSurat string, fileTempateMain string) (string, error) {
	var person InputSuratKeteranganMeninggal

	if fileTempateMain == "" {
		return "", errors.New("File tempate Surat tidak di temukan!")
	}

	err := json.Unmarshal(data.Data, &person)
	if err != nil {
		return "", err
	}

	dateTimeNow, err := helper.StringToDateSepesific(person.TanggalMeninggal)
	tanggalIndo, _ := helper.DateToDateIndoFormat(dateTimeNow)
	date := time.Now()
	timeindo, err := helper.IndonesiaFormat(date)

	fileTempate := fmt.Sprintf("./derektori/surat/template/%s", fileTempateMain)
	r, err := DocxReplace.ReadDocxFile(fileTempate)
	if err != nil {
		return "", err
	}
	docx1 := r.Editable()
	docx1.Replace("no-surat", helper.NomerSurat(data.NoSurat), -1)
	docx1.Replace("kode-surat", kodeSurat, -1)
	docx1.Replace("tahun", fmt.Sprint(date.Year()), -1)

	docx1.Replace("tanggal-dibuat", timeindo, -1)

	docx1.Replace("nama", person.Nama, -1)
	docx1.Replace("nik", person.NIK, -1)
	docx1.Replace("jenis-kelamin", person.JenisKelamin, -1)
	docx1.Replace("alamat", person.Alamat, -1)

	docx1.Replace("hari-meningal", person.HariMeninggal, -1)
	docx1.Replace("tanggal-meninggal", fmt.Sprint(tanggalIndo), -1)
	docx1.Replace("tempat-meninggal", person.TempatMeninggal, -1)
	docx1.Replace("penyebab", person.Penyebab, -1)

	NameSurat, path := GenerateSuratName(data)
	err = docx1.WriteToFile(path)
	if err != nil {
		return "", err
	}

	r.Close()

	return NameSurat, nil
}

func SuratNikahNSatu(data Surat, kodeSurat string, fileTempateMain string) (string, error) {
	var person InputSuratPengantarNikahNSatu

	if fileTempateMain == "" {
		return "", errors.New("File tempate Surat tidak di temukan!")
	}

	err := json.Unmarshal(data.Data, &person)
	if err != nil {
		return "", err
	}

	dateTimeNowPemohon, _ := helper.StringToDateSepesific(person.TanggalLahirPemohon)
	ttlPemohon, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirPemohon, dateTimeNowPemohon)

	dateTimeNowAyahPemohon, _ := helper.StringToDateSepesific(person.TanggalLahirAyahPemohon)
	ttlAyahPemohon, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirAyahPemohon, dateTimeNowAyahPemohon)

	dateTimeNowIbuPemohon, _ := helper.StringToDateSepesific(person.TanggalLahirIbuPemohon)
	ttlIbuPemohon, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirIbuPemohon, dateTimeNowIbuPemohon)

	date := time.Now()
	timeindo, _ := helper.IndonesiaFormat(date)

	replaceMap := docx.PlaceholderMap{
		"tahun":                     date.Year(),
		"id-surat":                  "",
		"kode-surat":                kodeSurat,
		"no-surat":                  helper.NomerSurat(data.NoSurat),
		"nama-pemohon":              person.NamaPemohon,
		"nik-pemohon":               person.NikPemohon,
		"jenis-kelamin-pemohon":     person.JenisKelaminPemohon,
		"ttl-pemohon":               strings.TrimSpace(ttlPemohon),
		"kewarganegaraan-pemohon":   person.KewarganegaraanPemohon,
		"agama-pemohon":             person.AgamaPemohon,
		"pekerjaan-pemohon":         person.PekerjaanPemohon,
		"alamat-pemohon":            person.AlamatPemohon,
		"status-pernikahan-pemohon": person.StatusPernikahanPemohon,
		// "beristrike-pemohon":        person.BeristrikePemohon,

		"nama-ayah-pemohon": person.NamaAyahPemohon,
		"nik-ayah-pemohon":  person.NikAyahPemohon,
		// "jenis-kelamin-ayah-pemohon":   person.JenisKelaminAyahPemohon,
		"ttl-ayah-pemohon":             strings.TrimSpace(ttlAyahPemohon),
		"kewarganegaraan-ayah-pemohon": person.KewarganegaraanAyahPemohon,
		"agama-ayah-pemohon":           person.AgamaAyahPemohon,
		"pekerjaan-ayah-pemohon":       person.PekerjaanAyahPemohon,
		"alamat-ayah-pemohon":          person.AlamatAyahPemohon,

		"nama-ibu-pemohon": person.NamaIbuPemohon,
		"nik-ibu-pemohon":  person.NikIbuPemohon,
		// "jenis-kelamin-ibu-pemohon":   person.JenisKelaminIbuPemohon,
		"ttl-ibu-pemohon":             strings.TrimSpace(ttlIbuPemohon),
		"kewarganegaraan-ibu-pemohon": person.KewarganegaraanIbuPemohon,
		"agama-ibu-pemohon":           person.AgamaIbuPemohon,
		"pekerjaan-ibu-pemohon":       person.PekerjaanIbuPemohon,
		"alamat-ibu-pemohon":          person.AlamatIbuPemohon,
		"tanggal-dibuat":              timeindo,
	}

	fileTempate := fmt.Sprintf("derektori/surat/template/%s", fileTempateMain)
	doc, err := docx.Open(fileTempate)
	if err != nil {
		return "", err
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		return "", err
	}

	NameSurat, path := GenerateSuratName(data)
	err = doc.WriteToFile(path)
	if err != nil {
		return "", err
	}
	return NameSurat, nil
}

func SuratNikahNSatuRiplace(data Surat, kodeSurat string, fileTempateMain string) (string, error) {
	var person InputSuratPengantarNikahNSatu

	if fileTempateMain == "" {
		return "", errors.New("File tempate Surat tidak di temukan!")
	}

	err := json.Unmarshal(data.Data, &person)
	if err != nil {
		return "", err
	}

	dateTimeNowPemohon, _ := helper.StringToDateSepesific(person.TanggalLahirPemohon)
	ttlPemohon, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirPemohon, dateTimeNowPemohon)

	dateTimeNowAyahPemohon, _ := helper.StringToDateSepesific(person.TanggalLahirAyahPemohon)
	ttlAyahPemohon, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirAyahPemohon, dateTimeNowAyahPemohon)

	dateTimeNowIbuPemohon, _ := helper.StringToDateSepesific(person.TanggalLahirIbuPemohon)
	ttlIbuPemohon, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirIbuPemohon, dateTimeNowIbuPemohon)

	date := time.Now()
	timeindo, _ := helper.IndonesiaFormat(date)

	fileTempate := fmt.Sprintf("./derektori/surat/template/%s", fileTempateMain)
	r, err := DocxReplace.ReadDocxFile(fileTempate)
	if err != nil {
		return "", err
	}
	docx1 := r.Editable()
	docx1.Replace("no-surat", helper.NomerSurat(data.NoSurat), -1)
	docx1.Replace("kode-surat", kodeSurat, -1)
	docx1.Replace("tahun", fmt.Sprint(date.Year()), -1)

	docx1.Replace("tanggal-dibuat", timeindo, -1)

	docx1.Replace("nama-pemohon", person.NamaPemohon, -1)
	docx1.Replace("nik-pemohon", fmt.Sprint(person.NikPemohon), -1)
	docx1.Replace("jenis-kelamin-pemohon", person.JenisKelaminPemohon, -1)
	docx1.Replace("ttl-pemohon", strings.TrimSpace(ttlPemohon), -1)
	docx1.Replace("kewarganegaraan-pemohon", person.KewarganegaraanPemohon, -1)
	docx1.Replace("agama-pemohon", person.AgamaPemohon, -1)
	docx1.Replace("pekerjaan-pemohon", person.PekerjaanPemohon, -1)
	docx1.Replace("alamat-pemohon", person.AlamatPemohon, -1)
	docx1.Replace("status-pernikahan-pemohon", person.StatusPernikahanPemohon, -1)

	docx1.Replace("nama-ayah-pemohon", person.NamaAyahPemohon, -1)
	docx1.Replace("nik-ayah-pemohon", fmt.Sprint(person.NikAyahPemohon), -1)
	docx1.Replace("ttl-ayah-pemohon", strings.TrimSpace(ttlAyahPemohon), -1)
	docx1.Replace("kewarganegaraan-ayah-pemohon", person.KewarganegaraanAyahPemohon, -1)
	docx1.Replace("agama-ayah-pemohon", person.AgamaAyahPemohon, -1)
	docx1.Replace("pekerjaan-ayah-pemohon", person.PekerjaanAyahPemohon, -1)
	docx1.Replace("alamat-ayah-pemohon", person.AlamatAyahPemohon, -1)

	docx1.Replace("nama-ibu-pemohon", person.NamaIbuPemohon, -1)
	docx1.Replace("nik-ibu-pemohon", fmt.Sprint(person.NikIbuPemohon), -1)
	docx1.Replace("ttl-ibu-pemohon", strings.TrimSpace(ttlIbuPemohon), -1)
	docx1.Replace("kewarganegaraan-ibu-pemohon", person.KewarganegaraanIbuPemohon, -1)
	docx1.Replace("agama-ibu-pemohon", person.AgamaIbuPemohon, -1)
	docx1.Replace("pekerjaan-ibu-pemohon", person.PekerjaanIbuPemohon, -1)
	docx1.Replace("alamat-ibu-pemohon", person.AlamatIbuPemohon, -1)

	// docx1.Replace("no_surat", helper.NomerSurat(data.NoSurat), -1)
	// docx1.Replace("kode_surat", kodeSurat, -1)
	// docx1.Replace("tahun_", fmt.Sprint(date.Year()), -1)

	// docx1.Replace("tanggal_dibuat_24", timeindo, -1)

	// docx1.Replace("nama_pemohon_1", person.NamaPemohon, -1)
	// docx1.Replace("nik_pemohon_2", fmt.Sprint(person.NikPemohon), -1)
	// docx1.Replace("jenis_kelamin_pemohon_3", person.JenisKelaminPemohon, -1)
	// docx1.Replace("ttl_pemohon_4", strings.TrimSpace(ttlPemohon), -1)
	// docx1.Replace("kewarganegaraan_pemohon_5", person.KewarganegaraanPemohon, -1)
	// docx1.Replace("agama_pemohon_6", person.AgamaPemohon, -1)
	// docx1.Replace("pekerjaan_pemohon_7", person.PekerjaanPemohon, -1)
	// docx1.Replace("alamat_pemohon_8", person.AlamatPemohon, -1)
	// docx1.Replace("status_pernikahan_pemohon_9", person.StatusPernikahanPemohon, -1)

	// docx1.Replace("nama_ayah_pemohon_10", person.NamaAyahPemohon, -1)
	// docx1.Replace("nik_ayah_pemohon_11", fmt.Sprint(person.NikAyahPemohon), -1)
	// docx1.Replace("ttl_ayah_pemohon_12", strings.TrimSpace(ttlAyahPemohon), -1)
	// docx1.Replace("kewarganegaraan_ayah_pemohon_13", person.KewarganegaraanAyahPemohon, -1)
	// docx1.Replace("agama_ayah_pemohon_14", person.AgamaAyahPemohon, -1)
	// docx1.Replace("pekerjaan_ayah_pemohon_15", person.PekerjaanAyahPemohon, -1)
	// docx1.Replace("alamat_ayah_pemohon_16", person.AlamatAyahPemohon, -1)

	// docx1.Replace("nama_ibu_pemohon_17", person.NamaIbuPemohon, -1)
	// docx1.Replace("nik_ibu_pemohon_18", fmt.Sprint(person.NikIbuPemohon), -1)
	// docx1.Replace("ttl_ibu_pemohon_19", strings.TrimSpace(ttlIbuPemohon), -1)
	// docx1.Replace("kewarganegaraan_ibu_pemohon_20", person.KewarganegaraanIbuPemohon, -1)
	// docx1.Replace("agama_ibu_pemohon_21", person.AgamaIbuPemohon, -1)
	// docx1.Replace("pekerjaan_ibu_pemohon_22", person.PekerjaanIbuPemohon, -1)
	// docx1.Replace("alamat_ibu_pemohon_23", person.AlamatIbuPemohon, -1)

	NameSurat, path := GenerateSuratName(data)
	err = docx1.WriteToFile(path)
	if err != nil {
		return "", err
	}

	r.Close()

	return NameSurat, nil
}

func SuratNikahNEmapat(data Surat, kodeSurat string, fileTempateMain string) (string, error) {
	var person InputSuratPengatarNikahNEmpat

	if fileTempateMain == "" {
		return "", errors.New("File tempate Surat tidak di temukan!")
	}

	err := json.Unmarshal(data.Data, &person)
	if err != nil {
		return "", err
	}

	dateTimeNowSuami, _ := helper.StringToDateSepesific(person.TanggalLahirCalonSuami)
	ttlSuami, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirCalonSuami, dateTimeNowSuami)

	dateTimeNowIstri, _ := helper.StringToDateSepesific(person.TanggalLahirCalonIstri)
	ttlAIstri, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirCalonIstri, dateTimeNowIstri)
	date := time.Now()
	timeindo, _ := helper.IndonesiaFormat(date)

	replaceMap := docx.PlaceholderMap{
		"id-surat":       "",
		"tahun":          date.Year(),
		"kode-surat":     kodeSurat,
		"no-surat":       helper.NomerSurat(data.NoSurat),
		"tanggal-dibuat": timeindo,

		"nama-calon-suami":  person.NamaCalonSuami,
		"binti-calon-suami": person.BintiCalonSuami,
		"nik-calon-suami":   person.NikCalonSuami,
		// "jenis-kelamin-calon-suami":   person.JenisKelaminCalonSuami,
		"ttl-calon-suami":             strings.TrimSpace(ttlSuami),
		"kewarganegaraan-calon-suami": person.KewarganegaraanCalonSuami,
		"agama-calon-suami":           person.AgamaCalonSuami,
		"pekerjaan-calon-suami":       person.PekerjaanCalonSuami,
		"alamat-calon-suami":          person.AlamatCalonSuami,

		"nama-calon-istri":  person.NamaCalonIstri,
		"binti-calon-istri": person.BintiCalonIstri,
		"nik-calon-istri":   person.NikCalonIstri,
		// "jenis-kelamin-calon-istri":   person.JenisKelaminCalonIstri,
		"ttl-calon-istri":             strings.TrimSpace(ttlAIstri),
		"kewarganegaraan-calon-istri": person.KewarganegaraanCalonIstri,
		"agama-calon-istri":           person.AgamaCalonIstri,
		"pekerjaan-calon-istri":       person.PekerjaanCalonIstri,
		"alamat-calon-istri":          person.AlamatCalonIstri,
	}

	fileTempate := fmt.Sprintf("derektori/surat/template/%s", fileTempateMain)
	doc, err := docx.Open(fileTempate)
	if err != nil {
		return "", err
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		return "", err
	}

	NameSurat, path := GenerateSuratName(data)
	err = doc.WriteToFile(path)
	if err != nil {
		return "", err
	}
	return NameSurat, nil
}

func SuratNikahNEmapatRiplace(data Surat, kodeSurat string, fileTempateMain string) (string, error) {
	var person InputSuratPengatarNikahNEmpat

	if fileTempateMain == "" {
		return "", errors.New("File tempate Surat tidak di temukan!")
	}

	err := json.Unmarshal(data.Data, &person)
	if err != nil {
		return "", err
	}

	dateTimeNowSuami, _ := helper.StringToDateSepesific(person.TanggalLahirCalonSuami)
	ttlSuami, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirCalonSuami, dateTimeNowSuami)

	dateTimeNowIstri, _ := helper.StringToDateSepesific(person.TanggalLahirCalonIstri)
	ttlAIstri, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirCalonIstri, dateTimeNowIstri)
	date := time.Now()
	timeindo, _ := helper.IndonesiaFormat(date)

	fileTempate := fmt.Sprintf("./derektori/surat/template/%s", fileTempateMain)
	r, err := DocxReplace.ReadDocxFile(fileTempate)
	if err != nil {
		fmt.Println("c4 :", err)
		return "", err
	}
	docx1 := r.Editable()
	docx1.Replace("no-surat", helper.NomerSurat(data.NoSurat), -1)
	docx1.Replace("kode-surat", kodeSurat, -1)
	docx1.Replace("tahun", timeindo, -1)

	docx1.Replace("tanggal-dibuat", timeindo, -1)

	docx1.Replace("nama-calon-suami", person.NamaCalonSuami, -1)
	docx1.Replace("binti-calon-suami", person.BintiCalonSuami, -1)
	docx1.Replace("nik-calon-suami", person.NikCalonSuami, -1)
	docx1.Replace("ttl-calon-suami", strings.TrimSpace(ttlSuami), -1)
	docx1.Replace("kewarganegaraan-calon-suami", person.KewarganegaraanCalonSuami, -1)
	docx1.Replace("agama-calon-suami", person.AgamaCalonSuami, -1)
	docx1.Replace("pekerjaan-calon-suami", person.PekerjaanCalonSuami, -1)
	docx1.Replace("alamat-calon-suami", person.AlamatCalonSuami, -1)

	docx1.Replace("nama-calon-istri", person.NamaCalonIstri, -1)
	docx1.Replace("binti-calon-istri", person.BintiCalonIstri, -1)
	docx1.Replace("nik-calon-istri", person.NikCalonIstri, -1)
	docx1.Replace("ttl-calon-istri", strings.TrimSpace(ttlAIstri), -1)
	docx1.Replace("kewarganegaraan-calon-istri", person.KewarganegaraanCalonIstri, -1)
	docx1.Replace("agama-calon-istri", person.AgamaCalonIstri, -1)
	docx1.Replace("pekerjaan-calon-istri", person.PekerjaanCalonIstri, -1)
	docx1.Replace("alamat-calon-istri", person.AlamatCalonIstri, -1)

	NameSurat, path := GenerateSuratName(data)
	err = docx1.WriteToFile(path)
	if err != nil {
		return "", err
	}

	r.Close()

	return NameSurat, nil
}

func SuratNikahNLima(data Surat, kodeSurat string, fileTempateMain string) (string, error) {
	var person InputSuratPengatarNikahNLima
	fmt.Println("c1 :", data)

	if fileTempateMain == "" {
		fmt.Println("c2 :", fileTempateMain)

		return "", errors.New("File tempate Surat tidak di temukan!")
	}

	err := json.Unmarshal(data.Data, &person)
	if err != nil {
		fmt.Println("c3 :", err)

		return "", err
	}

	dateTimeNowPemohon, _ := helper.StringToDateSepesific(person.TanggalLahirPemohon)
	ttlPemohon, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirPemohon, dateTimeNowPemohon)

	dateTimeNowPendamping, _ := helper.StringToDateSepesific(person.TanggalLahirPendaping)
	ttlPendamping, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirPendaping, dateTimeNowPendamping)

	dateTimeNowAyahPemohon, _ := helper.StringToDateSepesific(person.TanggalLahirAyahPemohon)
	ttlAyahPemohon, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirAyahPemohon, dateTimeNowAyahPemohon)

	dateTimeNowIbuPemohon, _ := helper.StringToDateSepesific(person.TanggalLahirIbuPemohon)
	ttlIbuPemohon, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirIbuPemohon, dateTimeNowIbuPemohon)

	date := time.Now()
	timeindo, _ := helper.IndonesiaFormat(date)

	replaceMap := docx.PlaceholderMap{
		"tanggal-dibuat": timeindo,
		"kode-surat":     kodeSurat,
		"no-surat":       helper.NomerSurat(data.NoSurat),
		// Ayah/wali/pengampu
		"nama-ayah-pemohon":            person.NamaAyahPemohon,
		"nama-ayah-ttd":                person.NamaAyahPemohon,
		"binti-ayah-pemohon":           person.BintiAyahPemohon,
		"nik-ayah-pemohon":             person.NikAyahPemohon,
		"jenis-kelamin-ayah-pemohon":   person.JenisKelaminAyahPemohon,
		"ttl-ayah-pemohon":             strings.TrimSpace(ttlAyahPemohon),
		"kewarganegaraan-ayah-pemohon": person.KewarganegaraanAyahPemohon,
		"agama-ayah-pemohon":           person.AgamaAyahPemohon,
		"pekerjaan-ayah-pemohon":       person.PekerjaanAyahPemohon,
		"alamat-ayah-pemohon":          person.AlamatAyahPemohon,
		"status-ayahbertandatangan":    person.StatusAyahPemohon,
		// ibu/wali/pengampu
		"nama-ibu-pemohon":            person.NamaPemohon,
		"nama-ibu-ttd":                person.NamaPemohon,
		"binti-ibu-pemohon":           person.BintiIbuPemohon,
		"nik-ibu-pemohon":             person.NikIbuPemohon,
		"jenis-kelamin-ibu-pemohon":   person.JenisKelaminIbuPemohon,
		"ttl-ibu-pemohon":             strings.TrimSpace(ttlIbuPemohon),
		"kewarganegaraan-ibu-pemohon": person.KewarganegaraanIbuPemohon,
		"agama-ibu-pemohon":           person.AgamaIbuPemohon,
		"pekerjaan-ibu-pemohon":       person.PekerjaanIbuPemohon,
		"alamat-ibu-pemohon":          person.AlamatIbuPemohon,
		"status-ibubertandatangan":    person.StatusIbuPemohon,
		// pemohon
		"nama-pemohon":            person.NamaPemohon,
		"binti-pemohon":           person.BintiPemohon,
		"nik-pemohon":             person.NikPemohon,
		"jenis-kelamin-pemohon":   person.JenisKelaminPemohon,
		"ttl-pemohon":             strings.TrimSpace(ttlPemohon),
		"kewarganegaraan-pemohon": person.KewarganegaraanPemohon,
		"agama-pemohon":           person.AgamaPemohon,
		"pekerjaan-pemohon":       person.PekerjaanPemohon,
		"alamat-pemohon":          person.AlamatPemohon,
		// pendamping
		"nama-pendamping":            person.NamaPendaping,
		"binti-pendamping":           person.BintiPendamping,
		"nik-pendamping":             person.NikPendaping,
		"jenis-kelamin-pendamping":   person.JenisKelaminPendaping,
		"ttl-pendamping":             strings.TrimSpace(ttlPendamping),
		"kewarganegaraan-pendamping": person.KewarganegaraanPendaping,
		"agama-pendamping":           person.AgamaPendaping,
		"pekerjaan-pendamping":       person.PekerjaanPendaping,
		"alamat-pendamping":          person.AlamatPendaping,
	}

	fileTempate := fmt.Sprintf("derektori/surat/template/%s", fileTempateMain)
	doc, err := docx.Open(fileTempate)
	if err != nil {
		fmt.Println("c4 :", err)

		return "", err
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		fmt.Println("c5 :", err)
		return "", err
	}

	NameSurat, path := GenerateSuratName(data)
	err = doc.WriteToFile(path)
	if err != nil {
		fmt.Println("c6 :", err)

		return "", err
	}

	fmt.Println("c7 :end")

	return NameSurat, nil
}

func SuratNikahNLimaRiplace(data Surat, kodeSurat string, fileTempateMain string) (string, error) {
	var person InputSuratPengatarNikahNLima
	fmt.Println("c1 :", data)

	if fileTempateMain == "" {
		fmt.Println("c2 :", fileTempateMain)

		return "", errors.New("File tempate Surat tidak di temukan!")
	}

	err := json.Unmarshal(data.Data, &person)
	if err != nil {
		fmt.Println("c3 :", err)

		return "", err
	}

	dateTimeNowPemohon, _ := helper.StringToDateSepesific(person.TanggalLahirPemohon)
	ttlPemohon, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirPemohon, dateTimeNowPemohon)
	// s2 := strings.TrimRight(ttlPemohon, " ")
	dateTimeNowPendamping, _ := helper.StringToDateSepesific(person.TanggalLahirPendaping)
	ttlPendamping, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirPendaping, dateTimeNowPendamping)

	dateTimeNowAyahPemohon, _ := helper.StringToDateSepesific(person.TanggalLahirAyahPemohon)
	ttlAyahPemohon, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirAyahPemohon, dateTimeNowAyahPemohon)

	dateTimeNowIbuPemohon, _ := helper.StringToDateSepesific(person.TanggalLahirIbuPemohon)
	ttlIbuPemohon, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahirIbuPemohon, dateTimeNowIbuPemohon)

	date := time.Now()
	timeindo, _ := helper.IndonesiaFormat(date)
	fileTempate := fmt.Sprintf("./derektori/surat/template/%s", fileTempateMain)
	r, err := DocxReplace.ReadDocxFile(fileTempate)
	if err != nil {
		fmt.Println("c4 :", err)
		return "", err
	}
	docx1 := r.Editable()
	docx1.Replace("kode-surat", kodeSurat, -1)
	docx1.Replace("no-surat", helper.NomerSurat(data.NoSurat), -1)
	docx1.Replace("tahun", timeindo, -1)

	docx1.Replace("tanggal-dibuat", timeindo, -1)

	docx1.Replace("nama-ayah-pemohon", person.NamaAyahPemohon, -1)
	docx1.Replace("nama-ayah-ttd", fmt.Sprintf(person.NamaAyahPemohon), -1)
	docx1.Replace("binti-ayah-pemohon", fmt.Sprintf(person.BintiAyahPemohon), -1)
	docx1.Replace("nik-ayah-pemohon", person.NikAyahPemohon, -1)
	docx1.Replace("jenis-kelamin-ayah-pemohon", person.JenisKelaminAyahPemohon, -1)
	docx1.Replace("ttl-ayah-pemohon", strings.TrimSpace(ttlAyahPemohon), -1)
	docx1.Replace("kewarganegaraan-ayah-pemohon", person.KewarganegaraanAyahPemohon, -1)
	docx1.Replace("agama-ayah-pemohon", person.AgamaAyahPemohon, -1)
	docx1.Replace("pekerjaan-ayah-pemohon", person.PekerjaanAyahPemohon, -1)
	docx1.Replace("alamat-ayah-pemohon", person.AlamatAyahPemohon, -1)
	docx1.Replace("status-ayah-bertandatangan", person.StatusAyahPemohon, -1)

	docx1.Replace("nama-ibu-pemohon", person.NamaPemohon, -1)
	docx1.Replace("nama-ibu-ttd", person.NamaPemohon, -1)
	docx1.Replace("binti-ibu-pemohon", person.BintiIbuPemohon, -1)
	docx1.Replace("nik-ibu-pemohon", person.NikIbuPemohon, -1)
	docx1.Replace("jenis-kelamin-ibu-pemohon", person.JenisKelaminIbuPemohon, -1)
	docx1.Replace("ttl-ibu-pemohon", strings.TrimSpace(ttlIbuPemohon), -1)
	docx1.Replace("kewarganegaraan-ibu-pemohon", person.KewarganegaraanIbuPemohon, -1)
	docx1.Replace("agama-ibu-pemohon", person.AgamaIbuPemohon, -1)
	docx1.Replace("pekerjaan-ibu-pemohon", person.PekerjaanIbuPemohon, -1)
	docx1.Replace("alamat-ibu-pemohon", person.AlamatIbuPemohon, -1)
	docx1.Replace("status-ibu-bertandatangan", person.StatusIbuPemohon, -1)

	docx1.Replace("nama-pemohon", person.NamaPemohon, -1)
	docx1.Replace("binti-pemohon", person.BintiPemohon, -1)
	docx1.Replace("nik-pemohon", person.NikPemohon, -1)
	docx1.Replace("jenis-kelamin-pemohon", person.JenisKelaminPemohon, -1)
	docx1.Replace("ttl-pemohon", strings.TrimSpace(ttlPemohon), -1)
	docx1.Replace("kewarganegaraan-pemohon", person.KewarganegaraanPemohon, -1)
	docx1.Replace("agama-pemohon", person.AgamaPemohon, -1)
	docx1.Replace("pekerjaan-pemohon", person.PekerjaanPemohon, -1)
	docx1.Replace("alamat-pemohon", person.AlamatPemohon, -1)

	docx1.Replace("nama-pendaping", person.NamaPendaping, -1)
	docx1.Replace("binti-pendaping", person.BintiPendamping, -1)
	docx1.Replace("nik-pendaping", person.NikPendaping, -1)
	docx1.Replace("jenis-kelamin-pendaping", person.JenisKelaminPendaping, -1)
	docx1.Replace("ttl-pendaping", strings.TrimSpace(ttlPendamping), -1)
	docx1.Replace("kewarganegaraan-pendaping", person.KewarganegaraanPendaping, -1)
	docx1.Replace("agama-pendaping", person.AgamaPendaping, -1)
	docx1.Replace("pekerjaan-pendaping", person.PekerjaanPendaping, -1)
	docx1.Replace("alamat-pendaping", person.AlamatPendaping, -1)

	NameSurat, path := GenerateSuratName(data)
	err = docx1.WriteToFile(path)
	if err != nil {
		return "", err
	}

	r.Close()

	fmt.Println("c7 :end")
	return NameSurat, nil
}

func SuratKepemilikanTanah(data Surat, kodeSurat string, fileTempateMain string) (string, error) {
	var person InputSuratKepemilikanTanah

	if fileTempateMain == "" {
		return "", errors.New("File tempate Surat tidak di temukan!")
	}

	err := json.Unmarshal(data.Data, &person)
	if err != nil {
		return "", err
	}

	// ttl, _ := helper.TempatTanggalLahirFormatIndonesia(person.TempatLahir, person.TanggalLahir)
	// tanggalBuat, _ := helper.IndonesiaFormat(person.TanggalMulai)
	// tanggalSelesai, _ := helper.IndonesiaFormat(person.TanggalSelesai)
	// date := time.Now()

	replaceMap := docx.PlaceholderMap{
		"id-surat":   "",
		"kode-surat": fmt.Sprintln(data.NoSurat, "/", data.KodeSurat, "/"),

		"nama-pemohon":      person.NamaPemilik,
		"nik-pemohon":       person.NIKPemilik,
		"ttl-pemohon":       strings.TrimSpace(person.TempatLahirPemilik),
		"umur-pemohon":      person.UmurPemilik,
		"agama-pemohon":     person.AgamaPemilik,
		"status-pernikahan": person.StatusPenikahanPemilik,
		"pekerjaan-pemohon": person.PekerjaanPemilik,
		"alamat-pemohon":    person.AlamatPemilik,

		"luas-tanah":                        person.LuasTanah,
		"alamat-tanah":                      person.AlamatTanah,
		"tanah-hasil":                       person.TanahHasil,
		"nomer-sertifikat-tanah":            person.NomerSertifikatTanah,
		"tanggal-sertifikat-tanah":          person.TanggalSertifikatTanah,
		"instansi-pembuat-sertifikat-tanah": person.InstansiPeresmi,
		"wilayah-intansi":                   person.Wilayah,

		"batas-tanah-utara":   person.BatasTanahUtara,
		"batas-tanah-seletan": person.BatasTanahSeletan,
		"batas-tanah-timur":   person.BatasTanahTimur,
		"batas-tanah-barat":   person.BatasTanahBarat,

		"nomer-sertifikat-tanah-instansi":   person.NomerSertifikatTanahInstansi,
		"tanggal-sertifikat-tanah-instansi": person.TanggalSertifikatTanahInstansi,

		"saksi-batas-tanah-utara":   person.SaksiBatasTanahUtara,
		"saksi-batas-tanah-selatan": person.SaksiBatasTanahSeletan,
		"saksi-batas-tanah-timur":   person.SaksiBatasTanahTimur,
		"saksi-batas-tanah-barat":   person.SaksiBatasTanahBarat,
	}

	fileTempate := fmt.Sprintf("derektori/surat/template/%s", fileTempateMain)
	doc, err := docx.Open(fileTempate)
	if err != nil {
		return "", err
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		return "", err
	}

	NameSurat, path := GenerateSuratName(data)
	err = doc.WriteToFile(path)
	if err != nil {
		return "", err
	}
	return NameSurat, nil
}
