package filedetail

import (
	"app-desa-kepuk/arsip"
	"app-desa-kepuk/arsipkategori"
	"app-desa-kepuk/file"
	"app-desa-kepuk/helper"
	"fmt"
	"strconv"
)

type Service interface {
	GetFileDetailAll() ([]FileDetail, error)
	GetFileAllArsipID(id uint) ([]FileDetail, error)
	GetFileDetailAllDeleted() ([]FileDetail, error)
	GetFileDetailByID(id uint) (FileDetail, error)
	GetFileDetailByIDDeleted(id uint) (FileDetail, error)
}

type service struct {
	repository         file.Repository
	kategoriRepository arsipkategori.Repository
	arsipRepository    arsip.Repository
}

func NewService(repository file.Repository, kategoriRepository arsipkategori.Repository, arsipRepository arsip.Repository) *service {
	return &service{repository, kategoriRepository, arsipRepository}
}

func (s *service) GetFileDetailAll() ([]FileDetail, error) {
	var fileDetail []FileDetail

	files, err := s.repository.FindAll()
	if err != nil {
		return fileDetail, err
	}

	for i, file := range files {
		fileRow := FileDetail{}
		fileRow.Index = i + 1
		fileRow.ID = file.ID
		fileRow.FileName = file.NamaFile
		fileRow.Deskripsi = file.Deskripsi

		angkaDiformat := fmt.Sprintf("%.3f", file.FileSize)
		angka, _ := strconv.ParseFloat(angkaDiformat, 64)
		fileRow.FileSize = angka

		fileRow.FileStatus = file.Status
		fileRow.FileLocation = file.FileLocation
		timeFileCreate, err := helper.DateTmeToDateTimeIndoFormat(file.CreatedAt)
		if err != nil {
			return fileDetail, err
		}
		fileRow.CreatedAt = timeFileCreate
		arsip, err := s.arsipRepository.FindByID(file.ArsipID)
		if err != nil {
			return fileDetail, err
		}

		fileRow.ArsipID = arsip.ID
		fileRow.ArsipName = arsip.Nama

		fileDetail = append(fileDetail, fileRow)
	}

	return fileDetail, nil
}

func (s *service) GetFileAllArsipID(id uint) ([]FileDetail, error) {
	var fileDetail []FileDetail

	files, err := s.repository.FindFileByArsipID(id)
	if err != nil {
		return fileDetail, err
	}

	for i, file := range files {
		fileRow := FileDetail{}
		fileRow.Index = i + 1
		fileRow.ID = file.ID
		fileRow.FileName = file.NamaFile
		fileRow.Deskripsi = file.Deskripsi

		angkaDiformat := fmt.Sprintf("%.3f", file.FileSize)
		angka, _ := strconv.ParseFloat(angkaDiformat, 64)
		fileRow.FileSize = angka

		fileRow.FileStatus = file.Status
		fileRow.FileLocation = file.FileLocation
		timeFileCreate, err := helper.DateTmeToDateTimeIndoFormat(file.CreatedAt)
		if err != nil {
			return fileDetail, err
		}
		fileRow.CreatedAt = timeFileCreate
		arsip, err := s.arsipRepository.FindByID(file.ArsipID)
		if err != nil {
			return fileDetail, err
		}

		fileRow.ArsipID = arsip.ID
		fileRow.ArsipName = arsip.Nama

		fileDetail = append(fileDetail, fileRow)
	}

	return fileDetail, nil
}

func (s *service) GetFileDetailAllDeleted() ([]FileDetail, error) {
	var fileDetail []FileDetail

	files, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return fileDetail, err
	}

	for i, file := range files {
		fileRow := FileDetail{}
		fileRow.Index = i + 1
		fileRow.ID = file.ID
		fileRow.FileName = file.NamaFile
		fileRow.Deskripsi = file.Deskripsi

		angkaDiformat := fmt.Sprintf("%.3f", file.FileSize)
		angka, _ := strconv.ParseFloat(angkaDiformat, 64)
		fileRow.FileSize = angka

		fileRow.FileStatus = file.Status
		fileRow.FileLocation = file.FileLocation
		timeFileCreate, err := helper.DateTmeToDateTimeIndoFormat(file.CreatedAt)
		if err != nil {
			return fileDetail, err
		}
		fileRow.CreatedAt = timeFileCreate

		filertimeDeletedAt := fmt.Sprint(file.DeletedAt)
		timeFilerDelete, err := helper.StringToDateTimeIndoFormat(filertimeDeletedAt)
		if err != nil {
			return fileDetail, err
		}
		fileRow.DeletedAt = timeFilerDelete

		arsip, err := s.arsipRepository.FindByID(file.ArsipID)
		if err != nil {
			return fileDetail, err

		}
		fileRow.ArsipID = arsip.ID
		fileRow.ArsipName = arsip.Nama

		fileDetail = append(fileDetail, fileRow)
	}

	return fileDetail, nil
}

func (s *service) GetFileDetailByID(id uint) (FileDetail, error) {
	var fileRow FileDetail

	file, err := s.repository.FindByID(id)
	if err != nil {
		return fileRow, err
	}

	fileRow.ID = file.ID
	fileRow.FileName = file.NamaFile
	fileRow.Deskripsi = file.Deskripsi

	angkaDiformat := fmt.Sprintf("%.3f", file.FileSize)
	angka, _ := strconv.ParseFloat(angkaDiformat, 64)
	fileRow.FileSize = angka

	fileRow.FileLocation = file.FileLocation
	fileRow.FileStatus = file.Status
	timeFileCreate, err := helper.DateTmeToDateTimeIndoFormat(file.CreatedAt)
	if err != nil {
		return fileRow, err
	}
	fileRow.CreatedAt = timeFileCreate
	arsip, err := s.arsipRepository.FindByID(file.ArsipID)
	if err != nil {
		return fileRow, err
	}

	fileRow.ArsipName = arsip.Nama

	return fileRow, nil
}

func (s *service) GetFileDetailByIDDeleted(id uint) (FileDetail, error) {
	var fileRow FileDetail

	file, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return fileRow, err
	}

	fileRow.ID = file.ID
	fileRow.FileName = file.NamaFile
	fileRow.Deskripsi = file.Deskripsi

	angkaDiformat := fmt.Sprintf("%.3f", file.FileSize)
	angka, _ := strconv.ParseFloat(angkaDiformat, 64)
	fileRow.FileSize = angka

	fileRow.FileLocation = file.FileLocation
	fileRow.FileStatus = file.Status
	arsip, err := s.arsipRepository.FindByID(file.ArsipID)
	if err != nil {
		return fileRow, err
	}
	fileRow.ArsipID = arsip.ID
	fileRow.ArsipName = arsip.Nama

	return fileRow, nil
}
