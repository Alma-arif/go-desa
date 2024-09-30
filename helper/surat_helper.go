package helper

// func SuratHelper(data surat.Surat) (string, error) {
// 	return "", nil

// }

// func surateName(data surat.Surat) string {
// 	return ""

// }

// func suratUsaha(data surat.Surat) (string, error) {
// 	var person surat.InputSuratKeteranganUsaha

// 	err := json.Unmarshal(data.Data, &person)
// 	if err != nil {
// 		return "", err
// 	}

// 	replaceMap := docx.PlaceholderMap{
// 		"nama":                 person.Nama,
// 		"jenis-kelamin":        person.JenisKelamin,
// 		"agama":                person.Agama,
// 		"status":               person.Status,
// 		"nik":                  person.NIK,
// 		"tempat-tanggal-lahir": person.TempatTanggalLahir,
// 		"pekerjaan":            person.Pekerjaan,
// 		"keperluan":            person.Keperluan,
// 		"jenis-usaha":          person.JenisUsaha,
// 		"keterangan-lain":      person.Keterangan,
// 		"tanggal-dibuat":       person.TanggalMulai,
// 		"tanggal-berakhir":     person.TanggalSelesai,
// 	}

// 	doc, err := docx.Open("suratsekalinikah.docx")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// replace the keys with values from replaceMap
// 	err = doc.ReplaceAll(replaceMap)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// write out a new file
// 	err = doc.WriteToFile("replacedaa.docx")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return "", nil
// }

// func suratNikah(data surat.Surat) (string, error) {
// 	return "", nil
// }
