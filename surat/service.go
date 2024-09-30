package surat

import (
	"app-desa-kepuk/filesurat"
	"app-desa-kepuk/helper"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Service interface {
	GetAllSurat() ([]SuratView, error)
	GetAllSuratDeleted() ([]SuratView, error)
	GetSuratByID(id uint) (SuratView, error)
	DeletedSuratSoft(id uint) (Surat, error)
	DeleteSurat(id uint) (Surat, error)
	RestoreSurat(id uint) (Surat, error)
	GetSuratByKodeSurat(kodeSurat uint) (SuratView, error)

	CreateSuratUsaha(input InputSuratKeteranganUsaha) (Surat, error)
	UpdateSuratUsaha(input UpdateSuratKeteranganUsaha) (Surat, error)

	CreateSuratKeteranganKematian(input InputSuratKeteranganMeninggal) (Surat, error)
	UpdateSurateteranganKematian(input UpdateSuratKeteranganMeninggal) (Surat, error)

	CreateSuratKeteranganNikahNSatu(input InputSuratPengantarNikahNSatu) (Surat, error)
	UpdateSuratKeteranganNikahNSatu(input UpdateSuratPengantarNikahNSatu) (Surat, error)

	CreateSuratKeteranganNikahNEmpat(input InputSuratPengatarNikahNEmpat) (Surat, error)
	UpdateSuratKeteranganNikahNEmpat(input UpdateSuratPengatarNikahNEmpat) (Surat, error)

	CreateSuratKeteranganNikahNLima(input InputSuratPengatarNikahNLima) (Surat, error)
	UpdateSuratKeteranganNikahNLima(input UpdateSuratPengatarNikahNLima) (Surat, error)

	CreateSuratKepemilikanTanah(input InputSuratKepemilikanTanah) (Surat, error)
}

type service struct {
	repository Repository
	fileSurat  filesurat.Repository
}

func NewService(repository Repository, fileSurat filesurat.Repository) *service {
	return &service{repository, fileSurat}
}

func (s *service) GetAllSurat() ([]SuratView, error) {
	var suratResult []SuratView

	resultSurat, err := s.repository.FindAll()
	if err != nil {
		return suratResult, err
	}

	for i, surat := range resultSurat {
		suratRow := SuratView{}
		timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

		suratRow.ID = surat.ID
		suratRow.Unix = timestamp
		suratRow.Index = i + 1
		suratRow.Nama = surat.Nama
		suratRow.KodeSurat = surat.KodeSurat
		suratRow.NoSurat = surat.NoSurat
		suratRow.NoSuratString = helper.NomerSurat(surat.NoSurat)
		data, _ := s.fileSurat.FindByID(surat.ID)
		suratRow.KeteranagnKodeSurat = data.Nama
		suratRow.KodeSuratString = data.KodeSurat
		suratRow.Data = surat.Data
		suratRow.Path = surat.Path
		suratRow.Perihal = surat.Perihal
		suratRow.Keterangan = surat.Keterangan
		suratRow.FileLocation = surat.FileLocation
		suratRow.FilePDF = surat.FilePDF
		suratRow.CreatedAt = surat.CreatedAt

		suratResult = append(suratResult, suratRow)
	}

	return suratResult, nil
}

func (s *service) GetAllSuratDeleted() ([]SuratView, error) {
	var suratResult []SuratView

	resultSurat, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return suratResult, err
	}

	for i, surat := range resultSurat {
		suratRow := SuratView{}
		timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

		suratRow.ID = surat.ID
		suratRow.Unix = timestamp
		suratRow.Index = i + 1
		suratRow.Nama = surat.Nama
		suratRow.KodeSurat = surat.KodeSurat
		suratRow.NoSurat = surat.NoSurat
		suratRow.NoSuratString = helper.NomerSurat(surat.NoSurat)
		data, _ := s.fileSurat.FindByID(surat.ID)
		suratRow.KeteranagnKodeSurat = data.Nama
		suratRow.KodeSuratString = data.KodeSurat
		suratRow.Path = surat.Path
		suratRow.Perihal = surat.Perihal
		suratRow.Keterangan = surat.Keterangan
		suratRow.FileLocation = surat.FileLocation
		suratRow.FilePDF = surat.FilePDF
		suratRow.CreatedAt = surat.CreatedAt

		suratResult = append(suratResult, suratRow)
	}

	return suratResult, nil
}

func (s *service) GetSuratByID(id uint) (SuratView, error) {
	var ResultSurat SuratView

	surat, err := s.repository.FindByID(id)
	if err != nil {
		return ResultSurat, err
	}

	if surat.ID <= 0 {
		return ResultSurat, errors.New("Data tidak ditemukan!")
	}

	ResultSurat.ID = surat.ID
	ResultSurat.Nama = surat.Nama
	ResultSurat.KodeSurat = surat.KodeSurat
	ResultSurat.NoSurat = surat.NoSurat
	ResultSurat.NoSuratString = helper.NomerSurat(surat.NoSurat)

	data, _ := s.fileSurat.FindByID(surat.ID)
	ResultSurat.KeteranagnKodeSurat = data.Nama
	ResultSurat.KodeSuratString = data.KodeSurat
	ResultSurat.Data = surat.Data
	ResultSurat.Path = surat.Path
	ResultSurat.Perihal = surat.Perihal
	ResultSurat.Keterangan = surat.Keterangan
	ResultSurat.FileLocation = surat.FileLocation
	ResultSurat.FilePDF = surat.FilePDF
	ResultSurat.CreatedAt = surat.CreatedAt

	return ResultSurat, nil
}

func (s *service) DeletedSuratSoft(id uint) (Surat, error) {
	var surat Surat
	fmt.Println("b 1 :", id)
	rowSurat, err := s.repository.FindByID(id)
	if err != nil {
		fmt.Println("b 2 :", err)

		return surat, err
	}

	if rowSurat.ID <= 0 {
		fmt.Println("b 3 :", rowSurat.ID)

		return surat, errors.New("Data tidak di temukan!")
	}

	path := fmt.Sprintf("./derektori/surat/file_surat/%s", rowSurat.FileLocation)
	fmt.Println("b 4 a:", rowSurat.FileLocation)

	_, err = os.Stat(path)
	if err != nil {
		fmt.Println("b 4 :", err)

		return surat, err
	}

	namaSurat, NewPath := GenerateSuratName(rowSurat)
	fmt.Println("b 5 a:", namaSurat, " / ", NewPath)

	err = os.Rename(path, NewPath)
	if err != nil {
		fmt.Println("b 5 :", err)

		return surat, err
	}

	surat.ID = rowSurat.ID
	surat.KodeSurat = rowSurat.KodeSurat
	surat.NoSurat = rowSurat.NoSurat
	surat.Nama = rowSurat.Nama
	surat.Keterangan = rowSurat.Keterangan
	surat.Perihal = rowSurat.Perihal
	surat.FileLocation = namaSurat
	surat.FilePDF = rowSurat.FilePDF
	surat.Data = rowSurat.Data
	surat.Path = rowSurat.Path
	surat.CreatedAt = rowSurat.CreatedAt
	surat.UpdatedAt = rowSurat.UpdatedAt

	suratUpdate, err := s.repository.Update(surat)
	fmt.Println("b 6 :", err)

	if err != nil {
		fmt.Println("b 6 :", err)

		return surat, err
	}

	surat, err = s.repository.DeletedSoft(suratUpdate.ID)
	if err != nil {
		fmt.Println("b 7 :", err)

		return surat, err
	}

	fmt.Println("b 8 : end")

	return surat, nil
}

func (s *service) DeleteSurat(id uint) (Surat, error) {
	var surat Surat

	fmt.Println("b1 :", id)
	rowSurat, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		fmt.Println("b2 :", err)

		return surat, err
	}

	if rowSurat.ID <= 0 {
		fmt.Println("b3 :", rowSurat.ID)
		return surat, errors.New("Data tidak ditemukan!")
	}

	path := fmt.Sprintf("./derektori/surat/file_surat/%s", rowSurat.FileLocation)
	_, err = os.Stat(path)
	if err != nil {
		fmt.Println("b4 :", err)
		return surat, err
	}

	err = os.Remove(path)
	if err != nil {
		fmt.Println("b5 :", err)

		return surat, err
	}

	surat, err = s.repository.Deleted(id)
	if err != nil {
		fmt.Println("b6 :", err)

		return surat, err
	}

	fmt.Println("b7 : end")

	return surat, nil
}

func (s *service) RestoreSurat(id uint) (Surat, error) {
	var surat Surat

	surat, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return surat, err
	}

	updatedSuart, err := s.repository.UpdateDeletedAt(surat.ID)
	if err != nil {
		return surat, err
	}

	return updatedSuart, nil
}

func (s *service) GetSuratByKodeSurat(kodeSurat uint) (SuratView, error) {
	var suratResult SuratView

	resultSurat, err := s.repository.FindByKodeSurat(kodeSurat)
	if err != nil {
		return suratResult, err
	}

	suratResult.ID = resultSurat.ID
	suratResult.Nama = resultSurat.Nama
	suratResult.KodeSurat = resultSurat.KodeSurat
	suratResult.NoSurat = resultSurat.NoSurat
	suratResult.NoSuratString = helper.NomerSurat(resultSurat.NoSurat)

	data, _ := s.fileSurat.FindByID(resultSurat.ID)
	suratResult.KeteranagnKodeSurat = resultSurat.Nama
	suratResult.KodeSuratString = data.KodeSurat
	suratResult.Path = resultSurat.Path
	suratResult.Perihal = resultSurat.Perihal
	suratResult.Keterangan = resultSurat.Keterangan
	suratResult.FileLocation = resultSurat.FileLocation
	suratResult.FilePDF = resultSurat.FilePDF
	suratResult.CreatedAt = resultSurat.CreatedAt

	return suratResult, nil
}

// surat usaha
func (s *service) CreateSuratUsaha(input InputSuratKeteranganUsaha) (Surat, error) {
	var surat Surat

	fmt.Println("b1")

	// ubah data ke json format
	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("b2 :", err)
		return surat, err
	}

	// end

	kodeSurat, err := s.fileSurat.FindByID(input.KodeSurat)
	if err != nil {
		fmt.Println("b3 :", err)
		return surat, err
	}

	// nomer surat
	if input.NoSurat <= 0 {
		fmt.Println("b4")

		resultNoSurat, err := s.repository.FindByKodeSurat(kodeSurat.ID)
		if err != nil {
			fmt.Println("b5")

			return surat, err
		}

		surat.NoSurat = resultNoSurat.NoSurat + 1
	} else {
		fmt.Println("b6")
		surat.NoSurat = int(input.NoSurat)
	}
	// nomer surat End

	surat.KodeSurat = input.KodeSurat
	surat.Nama = input.Nama
	surat.Perihal = "SURAT KETERANGAN USAHA"
	surat.Keterangan = "Permohonan surat keterangan usaha"
	surat.Data = jsonData
	surat.Path = "surat-usaha"

	// pembuatan nama file surat
	var fileNameSurat string
	if kodeSurat.File == "" {
		fmt.Println("b7 :", kodeSurat.FileMain)

		fileNameSurat = kodeSurat.FileMain
	} else {
		fmt.Println("b8 :", kodeSurat.File)
		fileNameSurat = kodeSurat.File
	}

	fileName, err := SuratUsaha(surat, kodeSurat.KodeSurat, fileNameSurat)
	fmt.Println("b9a :", fileName)

	if err != nil {
		fmt.Println("b9 :", err)

		return surat, err
	}
	// end

	surat.FileLocation = fileName

	surat, err = s.repository.Save(surat)

	if err != nil {
		fmt.Println("b10 :", err)

		return surat, err
	}

	fmt.Println("b11 End")

	return surat, nil
}

func (s *service) UpdateSuratUsaha(input UpdateSuratKeteranganUsaha) (Surat, error) {
	var surat Surat

	kodeSurat, err := s.fileSurat.FindByID(input.KodeSurat)
	if err != nil {
		fmt.Println("b3 :", err)

		return surat, err
	}

	rowSurat, err := s.repository.FindByID(input.ID)
	if err != nil {
		return surat, err
	}

	if rowSurat.ID <= 0 {
		return surat, errors.New("Update data gagal.")
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error marshalling JSON 2 :", err)
		return surat, err
	}

	surat.ID = input.ID
	surat.KodeSurat = input.KodeSurat
	surat.NoSurat = input.NoSurat
	surat.Nama = input.Nama
	surat.Perihal = rowSurat.Perihal
	surat.Keterangan = rowSurat.Keterangan
	surat.Data = jsonData
	surat.CreatedAt = rowSurat.CreatedAt
	surat.Path = rowSurat.Path

	// pembuatan nama file surat
	var fileNameSurat string
	if kodeSurat.File == "" {
		fmt.Println("b7 :", kodeSurat.FileMain)
		fileNameSurat = kodeSurat.FileMain
	} else {
		fmt.Println("b8 :", kodeSurat.File)
		fileNameSurat = kodeSurat.File
	}

	fileName, err := SuratUsaha(surat, kodeSurat.KodeSurat, fileNameSurat)
	fmt.Println("b9a :", fileName)
	if err != nil {
		fmt.Println("b9 :", err)
		return surat, err
	}
	// end

	surat.FileLocation = fileName

	path := fmt.Sprint("./derektori/surat/file_surat/", rowSurat.FileLocation)
	err = os.Remove(path)
	if err != nil {
		fmt.Println("b10 :", err)

		return surat, err
	}

	dataSurat, err := s.repository.Update(surat)
	if err != nil {
		fmt.Println("Error marshalling JSON 3 :", err)
		return surat, err
	}

	return dataSurat, nil
}

// surat usaha end

// surat kematian
func (s *service) CreateSuratKeteranganKematian(input InputSuratKeteranganMeninggal) (Surat, error) {
	var surat Surat

	// fmt.Println("b1")

	// ubah data ke json format
	jsonData, err := json.Marshal(input)
	if err != nil {
		// fmt.Println("b2 :", err)

		return surat, err
	}

	// end

	kodeSurat, err := s.fileSurat.FindByID(input.KodeSurat)
	if err != nil {
		// fmt.Println("b3 :", err)

		return surat, err
	}

	// nomer surat
	if input.NoSurat <= 0 {
		// fmt.Println("b4")

		resultNoSurat, err := s.repository.FindByKodeSurat(kodeSurat.ID)
		if err != nil {
			// fmt.Println("b5")

			return surat, err
		}
		surat.NoSurat = resultNoSurat.NoSurat + 1
	} else {
		// fmt.Println("b6")
		surat.NoSurat = int(input.NoSurat)
	}
	// nomer surat End

	surat.KodeSurat = input.KodeSurat
	surat.Nama = input.Nama
	surat.Perihal = "SURAT KETERANGAN Kematian"
	surat.Keterangan = "Permohonan surat keterangan Kematian"
	surat.Data = jsonData
	surat.Path = "surat-keterangan-kematian"

	// pembuatan nama file surat
	var fileNameSurat string
	if kodeSurat.File == "" {
		// fmt.Println("b7 :", kodeSurat.FileMain)

		fileNameSurat = kodeSurat.FileMain
	} else {
		// fmt.Println("b8 :", kodeSurat.File)
		fileNameSurat = kodeSurat.File
	}

	fileName, err := SuratKeteranagnKematian(surat, kodeSurat.KodeSurat, fileNameSurat)
	if err != nil {
		// fmt.Println("b9 :", err)

		return surat, err
	}
	// end

	surat.FileLocation = fileName

	surat, err = s.repository.Save(surat)

	if err != nil {
		fmt.Println("b10 :", err)

		return surat, err
	}

	fmt.Println("b11 End")

	return surat, nil
}

func (s *service) UpdateSurateteranganKematian(input UpdateSuratKeteranganMeninggal) (Surat, error) {
	var surat Surat

	kodeSurat, err := s.fileSurat.FindByID(input.KodeSurat)
	if err != nil {
		fmt.Println("b3 :", err)

		return surat, err
	}

	rowSurat, err := s.repository.FindByID(input.ID)
	if err != nil {
		return surat, err
	}

	if rowSurat.ID <= 0 {
		return surat, errors.New("Update data gagal.")
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error marshalling JSON 2 :", err)
		return surat, err
	}

	surat.ID = input.ID
	surat.KodeSurat = input.KodeSurat
	surat.NoSurat = input.NoSurat
	surat.Nama = input.Nama
	surat.Perihal = rowSurat.Perihal
	surat.Keterangan = rowSurat.Keterangan
	surat.Data = jsonData
	surat.CreatedAt = rowSurat.CreatedAt
	surat.Path = rowSurat.Path

	// pembuatan nama file surat
	var fileNameSurat string
	if kodeSurat.File == "" {
		fmt.Println("b7 :", kodeSurat.FileMain)

		fileNameSurat = kodeSurat.FileMain
	} else {
		fmt.Println("b8 :", kodeSurat.File)
		fileNameSurat = kodeSurat.File
	}

	fileName, err := SuratKeteranagnKematian(surat, kodeSurat.KodeSurat, fileNameSurat)
	if err != nil {
		fmt.Println("b9 :", err)

		return surat, err
	}
	// end

	surat.FileLocation = fileName

	path := fmt.Sprint("./derektori/surat/file_surat/", rowSurat.FileLocation)
	err = os.Remove(path)
	if err != nil {
		return surat, err
	}

	dataSurat, err := s.repository.Update(surat)
	if err != nil {
		fmt.Println("Error marshalling JSON 3 :", err)
		return surat, err
	}

	return dataSurat, nil
}

// surat kematian end

// surat nikah N1
func (s *service) CreateSuratKeteranganNikahNSatu(input InputSuratPengantarNikahNSatu) (Surat, error) {
	var surat Surat

	// ubah data ke json format
	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("b2 :", err)
		return surat, err
	}
	// end

	kodeSurat, err := s.fileSurat.FindByID(input.KodeSurat)
	if err != nil {
		fmt.Println("b3 :", err)
		return surat, err
	}

	// nomer surat
	if input.NoSurat <= 0 {
		fmt.Println("b4")
		resultNoSurat, err := s.repository.FindByKodeSurat(kodeSurat.ID)
		if err != nil {
			fmt.Println("b5")
			return surat, err
		}
		surat.NoSurat = resultNoSurat.NoSurat + 1
	} else {
		fmt.Println("b6")
		surat.NoSurat = int(input.NoSurat)
	}
	// nomer surat End

	surat.KodeSurat = input.KodeSurat
	surat.Nama = input.NamaPemohon
	surat.Perihal = "SURAT PENGATAR NIKAH N1"
	surat.Keterangan = "Permohonan surat Pengatar Nikah N1"
	surat.Data = jsonData
	surat.Path = "surat-pengatar-nikah"

	// pembuatan nama file surat
	var fileNameSurat string
	if kodeSurat.File == "" {
		fmt.Println("b7 :", kodeSurat.FileMain)

		fileNameSurat = kodeSurat.FileMain
	} else {
		fmt.Println("b8 :", kodeSurat.File)
		fileNameSurat = kodeSurat.File
	}

	fileName, err := SuratNikahNSatu(surat, kodeSurat.KodeSurat, fileNameSurat)
	if err != nil {
		fmt.Println("b9 :", err)
		return surat, err
	}
	// end

	surat.FileLocation = fileName

	surat, err = s.repository.Save(surat)
	if err != nil {
		fmt.Println("b10 :", err)
		return surat, err
	}

	fmt.Println("b11 End")
	return surat, nil
}

func (s *service) UpdateSuratKeteranganNikahNSatu(input UpdateSuratPengantarNikahNSatu) (Surat, error) {
	var surat Surat

	kodeSurat, err := s.fileSurat.FindByID(input.KodeSurat)
	if err != nil {
		fmt.Println("b3 :", err)
		return surat, err
	}

	rowSurat, err := s.repository.FindByID(input.ID)
	if err != nil {
		return surat, err
	}

	if rowSurat.ID <= 0 {
		return surat, errors.New("Update data gagal.")
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error marshalling JSON 2 :", err)
		return surat, err
	}

	surat.ID = input.ID
	surat.KodeSurat = input.KodeSurat
	surat.NoSurat = input.NoSurat
	surat.Nama = input.NamaPemohon
	surat.Perihal = rowSurat.Perihal
	surat.Keterangan = rowSurat.Keterangan
	surat.Data = jsonData
	surat.CreatedAt = rowSurat.CreatedAt
	surat.Path = rowSurat.Path

	// pembuatan nama file surat
	var fileNameSurat string
	if kodeSurat.File == "" {
		fmt.Println("b7 :", kodeSurat.FileMain)

		fileNameSurat = kodeSurat.FileMain
	} else {
		fmt.Println("b8 :", kodeSurat.File)
		fileNameSurat = kodeSurat.File
	}

	fileName, err := SuratNikahNSatu(surat, kodeSurat.KodeSurat, fileNameSurat)
	if err != nil {
		fmt.Println("b9 :", err)

		return surat, err
	}
	// end

	surat.FileLocation = fileName

	path := fmt.Sprint("./derektori/surat/file_surat/", rowSurat.FileLocation)
	err = os.Remove(path)
	if err != nil {
		return surat, err
	}

	dataSurat, err := s.repository.Update(surat)
	if err != nil {
		fmt.Println("Error marshalling JSON 3 :", err)
		return surat, err
	}

	return dataSurat, nil
}

// surat nikah N1 end

// surat persetujuan pengatin N4
func (s *service) CreateSuratKeteranganNikahNEmpat(input InputSuratPengatarNikahNEmpat) (Surat, error) {
	var surat Surat

	// ubah data ke json format
	fmt.Println("b1 :", input)
	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("b2 :", err)

		return surat, err
	}
	// end

	kodeSurat, err := s.fileSurat.FindByID(input.KodeSurat)
	if err != nil {
		fmt.Println("b3 :", err)

		return surat, err
	}

	// nomer surat
	if input.NoSurat <= 0 {
		resultNoSurat, err := s.repository.FindByKodeSurat(kodeSurat.ID)
		if err != nil {
			fmt.Println("b4 :", err)

			return surat, err
		}
		surat.NoSurat = resultNoSurat.NoSurat + 1
	} else {
		surat.NoSurat = int(input.NoSurat)
	}
	// nomer surat End

	surat.KodeSurat = input.KodeSurat
	surat.Nama = input.NamaCalonSuami
	surat.Perihal = "SURAT PERSETUJUAN PENGANTIN"
	surat.Keterangan = "Permohonan surat PERSETUJUAN PENGATIN"
	surat.Data = jsonData
	surat.Path = "surat-persetujuan-pengatin"

	// pembuatan nama file surat
	var fileNameSurat string
	if kodeSurat.File == "" {
		fileNameSurat = kodeSurat.FileMain
	} else {
		fileNameSurat = kodeSurat.File
	}

	fileName, err := SuratNikahNEmapat(surat, kodeSurat.KodeSurat, fileNameSurat)
	if err != nil {
		fmt.Println("b5 :", err)

		return surat, err
	}
	// end

	surat.FileLocation = fileName

	surat, err = s.repository.Save(surat)
	if err != nil {
		fmt.Println("b6 :", err)

		return surat, err
	}

	return surat, nil
}

func (s *service) UpdateSuratKeteranganNikahNEmpat(input UpdateSuratPengatarNikahNEmpat) (Surat, error) {
	var surat Surat

	kodeSurat, err := s.fileSurat.FindByID(input.KodeSurat)
	if err != nil {
		fmt.Println("b3 :", err)
		return surat, err
	}

	rowSurat, err := s.repository.FindByID(input.ID)
	if err != nil {
		return surat, err
	}

	if rowSurat.ID <= 0 {
		return surat, errors.New("Update data gagal.")
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error marshalling JSON 2 :", err)
		return surat, err
	}

	surat.ID = input.ID
	surat.KodeSurat = input.KodeSurat
	surat.NoSurat = input.NoSurat
	surat.Nama = input.NamaCalonSuami
	surat.Perihal = rowSurat.Perihal
	surat.Keterangan = rowSurat.Keterangan
	surat.Data = jsonData
	surat.CreatedAt = rowSurat.CreatedAt
	surat.Path = rowSurat.Path

	// pembuatan nama file surat
	var fileNameSurat string
	if kodeSurat.File == "" {
		fmt.Println("b7 :", kodeSurat.FileMain)
		fileNameSurat = kodeSurat.FileMain
	} else {
		fmt.Println("b8 :", kodeSurat.File)
		fileNameSurat = kodeSurat.File
	}

	fileName, err := SuratNikahNEmapat(surat, kodeSurat.KodeSurat, fileNameSurat)
	if err != nil {
		fmt.Println("b9 :", err)

		return surat, err
	}
	// end

	surat.FileLocation = fileName

	path := fmt.Sprint("./derektori/surat/file_surat/", rowSurat.FileLocation)
	err = os.Remove(path)
	if err != nil {
		return surat, err
	}

	dataSurat, err := s.repository.Update(surat)
	if err != nil {
		fmt.Println("Error marshalling JSON 3 :", err)
		return surat, err
	}

	return dataSurat, nil
}

// surat persetujuan pengatin N4 end

// surat izin orang tua N5
func (s *service) CreateSuratKeteranganNikahNLima(input InputSuratPengatarNikahNLima) (Surat, error) {
	var surat Surat
	fmt.Println("b1 :", input)

	// ubah data ke json format
	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("b2 :", err)

		return surat, err
	}
	// end

	kodeSurat, err := s.fileSurat.FindByID(input.KodeSurat)
	if err != nil {
		fmt.Println("b3 :", err)

		return surat, err
	}

	// nomer surat
	if input.NoSurat <= 0 {
		resultNoSurat, err := s.repository.FindByKodeSurat(kodeSurat.ID)
		if err != nil {
			fmt.Println("b4 :", err)

			return surat, err
		}
		surat.NoSurat = resultNoSurat.NoSurat + 1
	} else {
		surat.NoSurat = int(input.NoSurat)
	}
	// nomer surat End

	surat.KodeSurat = input.KodeSurat
	surat.Nama = input.NamaPemohon
	surat.Perihal = "SURAT IZIN ORANG TUA"
	surat.Keterangan = "Permohonan surat Izin Orang Tua"
	surat.Data = jsonData
	surat.Path = "surat-izin-orang-tua"

	// pembuatan nama file surat
	var fileNameSurat string
	if kodeSurat.File == "" {
		fileNameSurat = kodeSurat.FileMain
	} else {
		fileNameSurat = kodeSurat.File
	}

	fileName, err := SuratNikahNLima(surat, kodeSurat.KodeSurat, fileNameSurat)
	if err != nil {
		fmt.Println("b5 :", err)

		return surat, err
	}
	// end

	surat.FileLocation = fileName

	surat, err = s.repository.Save(surat)
	if err != nil {
		fmt.Println("b6 :", err)

		return surat, err
	}

	fmt.Println("b7 : end")

	return surat, nil
}

func (s *service) UpdateSuratKeteranganNikahNLima(input UpdateSuratPengatarNikahNLima) (Surat, error) {
	var surat Surat

	kodeSurat, err := s.fileSurat.FindByID(input.KodeSurat)
	if err != nil {
		fmt.Println("b3 :", err)
		return surat, err
	}

	rowSurat, err := s.repository.FindByID(input.ID)
	if err != nil {
		return surat, err
	}

	if rowSurat.ID <= 0 {
		return surat, errors.New("Update data gagal.")
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error marshalling JSON 2 :", err)
		return surat, err
	}

	surat.ID = input.ID
	surat.KodeSurat = input.KodeSurat
	surat.NoSurat = input.NoSurat
	surat.Nama = input.NamaPemohon
	surat.Perihal = rowSurat.Perihal
	surat.Keterangan = rowSurat.Keterangan
	surat.Data = jsonData
	surat.CreatedAt = rowSurat.CreatedAt
	surat.Path = rowSurat.Path

	// pembuatan nama file surat
	var fileNameSurat string
	if kodeSurat.File == "" {
		fmt.Println("b7 :", kodeSurat.FileMain)
		fileNameSurat = kodeSurat.FileMain
	} else {
		fmt.Println("b8 :", kodeSurat.File)
		fileNameSurat = kodeSurat.File
	}

	fileName, err := SuratNikahNLima(surat, kodeSurat.KodeSurat, fileNameSurat)
	if err != nil {
		fmt.Println("b9 :", err)

		return surat, err
	}
	// end

	surat.FileLocation = fileName

	path := fmt.Sprint("./derektori/surat/file_surat/", rowSurat.FileLocation)
	err = os.Remove(path)
	if err != nil {
		return surat, err
	}

	dataSurat, err := s.repository.Update(surat)
	if err != nil {
		fmt.Println("Error marshalling JSON 3 :", err)
		return surat, err
	}

	return dataSurat, nil
}

// surat izin orang tua N5 end

func (s *service) CreateSuratKepemilikanTanah(input InputSuratKepemilikanTanah) (Surat, error) {
	var surat Surat

	// ubah data ke json format
	jsonData, err := json.Marshal(input)
	if err != nil {
		return surat, err
	}
	// end

	kodeSurat, err := s.fileSurat.FindByID(input.KodeSurat)
	if err != nil {
		return surat, err
	}

	// nomer surat
	if input.NoSurat <= 0 {
		resultNoSurat, err := s.repository.FindByKodeSurat(kodeSurat.ID)
		if err != nil {
			return surat, err
		}
		surat.NoSurat = resultNoSurat.NoSurat + 1
	} else {
		surat.NoSurat = int(input.NoSurat)
	}
	// nomer surat End

	surat.KodeSurat = input.KodeSurat
	surat.Nama = input.NamaPemilik
	surat.Perihal = "SURAT PERNYATAAN KEPEMILIKAN TANAH"
	surat.Keterangan = "Permohonan surat Pernyataan Kepemilikan Tanah"
	surat.Data = jsonData
	surat.Path = "surat-pernyataan-kepemilikan-tanah"

	// pembuatan nama file surat
	var fileNameSurat string
	if kodeSurat.File == "" {
		fileNameSurat = kodeSurat.FileMain
	} else {
		fileNameSurat = kodeSurat.File
	}

	fileName, err := SuratKepemilikanTanah(surat, kodeSurat.KodeSurat, fileNameSurat)
	if err != nil {
		return surat, err
	}
	// end

	surat.FileLocation = fileName

	surat, err = s.repository.Save(surat)
	if err != nil {
		return surat, err
	}

	return surat, nil
}
