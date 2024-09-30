package filesurat

import (
	"errors"
	"fmt"
	"os"
)

type Service interface {
	GetAllFileSurat() ([]FileSuratView, error)
	GetFileSuratByID(id uint) (FileSuratView, error)
	GetFileSuratByKodeFix(kode string) (FileSuratView, error)
	CreateFileSurat(input InputFileSurat) (FileSurat, error)
	UpdateFileSurat(input UpdateFileSurat) (FileSurat, error)
	UpdateFileSuratSecone(input UpdateFileSuratSecone) (FileSurat, error)
	DeletedFileSuartSoft(id uint) (FileSurat, error)
	DeleteFileSurat(id uint) (FileSurat, error)
	RestoreFileSurat(id uint) (FileSurat, error)
	GetAllFileSuratDeleted() ([]FileSuratView, error)
	DeleteFile(id uint) (FileSurat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllFileSurat() ([]FileSuratView, error) {
	var suratResult []FileSuratView

	resultSurat, err := s.repository.FindAll()
	if err != nil {
		return suratResult, err
	}

	for i, surat := range resultSurat {
		suratRow := FileSuratView{}

		suratRow.ID = surat.ID
		suratRow.Index = i + 1
		suratRow.Nama = surat.Nama
		suratRow.KodeSuratFix = surat.KodeSuratFix
		suratRow.KodeSurat = surat.KodeSurat
		suratRow.FileMain = surat.FileMain
		suratRow.NamaFileMain = surat.NamaFileMain
		suratRow.File = surat.File
		suratRow.NamaFile = surat.NamaFile
		suratRow.CreatedAt = surat.CreatedAt

		suratResult = append(suratResult, suratRow)
	}

	return suratResult, nil
}

func (s *service) GetFileSuratByID(id uint) (FileSuratView, error) {
	var ResultSurat FileSuratView

	surat, err := s.repository.FindByID(id)
	if err != nil {
		return ResultSurat, err
	}

	ResultSurat.ID = surat.ID
	ResultSurat.Nama = surat.Nama
	ResultSurat.KodeSuratFix = surat.KodeSuratFix
	ResultSurat.KodeSurat = surat.KodeSurat
	ResultSurat.FileMain = surat.FileMain
	ResultSurat.NamaFileMain = surat.NamaFileMain
	ResultSurat.File = surat.File
	ResultSurat.NamaFile = surat.NamaFile
	ResultSurat.CreatedAt = surat.CreatedAt
	if surat.ID == 0 {
		return ResultSurat, errors.New("ID File surat tidak di temukan")
	}

	return ResultSurat, nil
}

func (s *service) GetFileSuratByKodeFix(kode string) (FileSuratView, error) {
	var ResultSurat FileSuratView

	surat, err := s.repository.FindByIDKodeFix(kode)
	if err != nil {
		return ResultSurat, err
	}

	ResultSurat.ID = surat.ID
	ResultSurat.Nama = surat.Nama
	ResultSurat.KodeSuratFix = surat.KodeSuratFix
	ResultSurat.KodeSurat = surat.KodeSurat
	ResultSurat.FileMain = surat.FileMain
	ResultSurat.NamaFileMain = surat.NamaFileMain
	ResultSurat.File = surat.File
	ResultSurat.NamaFile = surat.NamaFile
	ResultSurat.CreatedAt = surat.CreatedAt
	if surat.ID == 0 {
		return ResultSurat, errors.New("ID File surat tidak di temukan")
	}

	return ResultSurat, nil
}

func (s *service) CreateFileSurat(input InputFileSurat) (FileSurat, error) {
	var surat FileSurat

	surat.KodeSuratFix = input.KodeSuratFix
	surat.KodeSurat = input.KodeSurat
	surat.Nama = input.Nama
	surat.KodeSurat = input.KodeSurat
	surat.FileMain = input.FileMain
	surat.NamaFileMain = input.NamaFileMain
	surat.File = input.File
	surat.NamaFile = input.NamaFile

	resultSurat, err := s.repository.Save(surat)
	if err != nil {

		return surat, err
	}

	return resultSurat, nil
}

func (s *service) UpdateFileSurat(input UpdateFileSurat) (FileSurat, error) {
	var surat FileSurat

	rowFileSurat, err := s.repository.FindByID(input.ID)
	if err != nil {
		return surat, err
	}
	if rowFileSurat.ID <= 0 {
		return surat, errors.New("Perubahan data gagal!")
	}

	surat.ID = rowFileSurat.ID
	surat.KodeSuratFix = rowFileSurat.KodeSuratFix
	surat.KodeSurat = input.KodeSurat
	surat.Nama = input.Nama
	surat.FileMain = rowFileSurat.FileMain
	surat.NamaFileMain = rowFileSurat.NamaFileMain
	surat.File = rowFileSurat.File
	surat.NamaFile = rowFileSurat.NamaFile
	surat.CreatedAt = rowFileSurat.CreatedAt

	result, err := s.repository.Update(surat)
	if err != nil {

		return surat, err
	}

	return result, nil
}

func (s *service) UpdateFileSuratSecone(input UpdateFileSuratSecone) (FileSurat, error) {
	var surat FileSurat
	fmt.Println("c1 :", input.File, " dan ", input.NamaFile)

	rowFileSurat, err := s.repository.FindByID(input.ID)
	if err != nil {
		fmt.Println("c2 :", err.Error())
		return surat, err
	}

	surat.ID = rowFileSurat.ID
	surat.KodeSuratFix = rowFileSurat.KodeSuratFix
	surat.KodeSurat = rowFileSurat.KodeSurat
	surat.Nama = rowFileSurat.Nama
	surat.KodeSurat = rowFileSurat.KodeSurat
	surat.FileMain = rowFileSurat.FileMain
	surat.NamaFileMain = rowFileSurat.NamaFileMain
	surat.CreatedAt = rowFileSurat.CreatedAt

	if rowFileSurat.File == "" {
		fmt.Println("c4 :", input.File, " dan ", input.NamaFile)
		surat.File = input.File
		surat.NamaFile = input.NamaFile
	} else {
		surat.File = input.File
		surat.NamaFile = input.NamaFile
		fmt.Println("c4a :", input.File, " dan ", input.NamaFile)
		path := fmt.Sprintf("derektori/surat/template/%s", rowFileSurat.File)
		err = os.Remove(path)
		if err != nil {
			fmt.Println("c5 :", err.Error())
			return surat, err
		}
	}

	result, err := s.repository.Update(surat)
	if err != nil {
		fmt.Println("c6 :", err.Error())
		return surat, err
	}

	fmt.Println("c7 :")
	return result, nil
}

func (s *service) DeletedFileSuartSoft(id uint) (FileSurat, error) {
	var surat FileSurat

	_, err := s.repository.FindByID(id)
	if err != nil {
		return surat, err
	}

	surat, err = s.repository.DeletedSoft(id)
	if err != nil {
		return surat, err
	}

	return surat, nil
}

func (s *service) DeleteFileSurat(id uint) (FileSurat, error) {
	var surat FileSurat

	rowSurat, err := s.repository.FindByID(id)
	if err != nil {
		return surat, err
	}

	pathRemoveFile := fmt.Sprintf("derektori/surat/template/%s", rowSurat.File)
	err = os.Remove(pathRemoveFile)
	if err != nil {
		return surat, err
	}

	pathRemoveFileMAin := fmt.Sprintf("derektori/surat/template/%s", rowSurat.FileMain)
	err = os.Remove(pathRemoveFileMAin)
	if err != nil {
		return surat, err
	}

	surat, err = s.repository.Deleted(id)
	if err != nil {
		return surat, err
	}

	return surat, nil
}

func (s *service) RestoreFileSurat(id uint) (FileSurat, error) {
	var surat FileSurat

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

func (s *service) GetAllFileSuratDeleted() ([]FileSuratView, error) {
	var suratResult []FileSuratView

	resultSurat, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return suratResult, err
	}

	for i, surat := range resultSurat {
		suratRow := FileSuratView{}

		suratRow.ID = surat.ID
		suratRow.Index = i + 1
		suratRow.Nama = surat.Nama
		suratRow.KodeSuratFix = surat.KodeSuratFix
		suratRow.KodeSurat = surat.KodeSurat
		suratRow.FileMain = surat.FileMain
		suratRow.NamaFileMain = surat.NamaFileMain
		suratRow.File = surat.File
		suratRow.NamaFile = surat.NamaFile
		suratRow.CreatedAt = surat.CreatedAt

		suratResult = append(suratResult, suratRow)
	}

	return suratResult, nil
}

func (s *service) DeleteFile(id uint) (FileSurat, error) {
	var surat FileSurat

	rowSurat, err := s.repository.FindByID(id)
	if err != nil {
		return surat, err
	}

	pathRemoveFile := fmt.Sprintf("derektori/surat/template/%s", rowSurat.File)
	err = os.Remove(pathRemoveFile)
	if err != nil {
		return surat, err
	}
	surat.ID = rowSurat.ID
	surat.KodeSuratFix = rowSurat.KodeSuratFix
	surat.KodeSurat = rowSurat.KodeSurat
	surat.Nama = rowSurat.Nama
	surat.KodeSurat = rowSurat.KodeSurat
	surat.FileMain = rowSurat.FileMain
	surat.NamaFileMain = rowSurat.NamaFileMain
	surat.File = ""
	surat.NamaFile = rowSurat.NamaFile

	resultSurat, err := s.repository.Update(surat)
	if err != nil {
		return surat, err
	}

	return resultSurat, nil
}
