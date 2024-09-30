package arsip

import (
	"app-desa-kepuk/arsipkategori"
	"app-desa-kepuk/file"
	"app-desa-kepuk/helper"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	GetAllArsip() ([]ArsipList, error)
	GetAllArsipDeleted() ([]ArsipList, error)
	GetArsipByID(id uint) (ArsipList, error)
	CreateArsip(input CreateArsipInput) (ArsipDesa, error)
	UpdateArsip(input UpdateArsipInput) (ArsipDesa, error)
	DeletedSoft(id uint) (ArsipDesa, error)
	Deleted(id uint) (ArsipDesa, error)
	RestoreArsip(id uint) (ArsipDesa, error)
}

type service struct {
	repository         Repository
	kategoriRepository arsipkategori.Repository
	fileRepository     file.Repository
	Validate           *validator.Validate
}

func NewService(repository Repository, kategoriRepository arsipkategori.Repository, fileRepository file.Repository, Validate *validator.Validate) *service {
	return &service{repository, kategoriRepository, fileRepository, Validate}
}

func (s *service) GetAllArsip() ([]ArsipList, error) {
	var arsipResult []ArsipList

	arsips, err := s.repository.FindAll()
	if err != nil {
		return arsipResult, err
	}

	for i, arsip := range arsips {
		arsipRow := ArsipList{}

		arsipRow.ID = arsip.ID
		arsipRow.Index = i + 1
		arsipRow.Nama = arsip.Nama
		arsipRow.Deskripsi = arsip.Deskripsi
		arsipRow.CreatedAt = arsip.CreatedAt

		kategori, err := s.kategoriRepository.FindByID(arsip.ID)
		if err != nil {
			return arsipResult, err
		}

		arsipRow.KategoriID = kategori.ID
		arsipRow.Kategori = kategori.Nama

		arsipResult = append(arsipResult, arsipRow)
	}

	return arsipResult, nil
}

func (s *service) GetAllArsipDeleted() ([]ArsipList, error) {
	var arsipResult []ArsipList

	arsips, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return arsipResult, err
	}

	for i, arsip := range arsips {
		arsipRow := ArsipList{}

		arsipRow.ID = arsip.ID
		arsipRow.Index = i + 1
		arsipRow.Nama = arsip.Nama
		arsipRow.Deskripsi = arsip.Deskripsi
		arsipRow.CreatedAt = arsip.CreatedAt
		arsiptimeDeletedAt := fmt.Sprint(arsip.DeletedAt)
		timeArsipDelete, err := helper.StringToDateTimeIndoFormat(arsiptimeDeletedAt)
		if err != nil {
			return arsipResult, err
		}
		arsipRow.DeletedAt = timeArsipDelete

		kategori, err := s.kategoriRepository.FindByID(arsip.ID)
		if err != nil {
			return arsipResult, err
		}

		arsipRow.KategoriID = kategori.ID
		arsipRow.Kategori = kategori.Nama

		arsipResult = append(arsipResult, arsipRow)
	}

	return arsipResult, nil
}

func (s *service) GetArsipByID(id uint) (ArsipList, error) {
	var ResiltArsip ArsipList

	arsip, err := s.repository.FindByID(id)
	if err != nil {
		return ResiltArsip, err
	}

	ResiltArsip.ID = arsip.ID
	ResiltArsip.Nama = arsip.Nama
	ResiltArsip.Deskripsi = arsip.Deskripsi
	ResiltArsip.CreatedAt = arsip.CreatedAt
	kategori, err := s.kategoriRepository.FindByID(arsip.ID)
	if err != nil {
		return ResiltArsip, err
	}

	ResiltArsip.KategoriID = kategori.ID
	ResiltArsip.Kategori = kategori.Nama

	if arsip.ID == 0 {
		return ResiltArsip, errors.New("No arisp found on with that ID")
	}

	return ResiltArsip, nil
}

func (s *service) CreateArsip(input CreateArsipInput) (ArsipDesa, error) {
	var arsip ArsipDesa

	err := s.Validate.Struct(input)
	if err != nil {
		return arsip, errors.New("isi form dengan benar")
	}
	arsip.Nama = input.Nama
	arsip.KategoriID = input.KategoriID
	arsip.Deskripsi = input.Deskripsi

	newArsip, err := s.repository.Save(arsip)
	if err != nil {
		return newArsip, err
	}

	return newArsip, nil
}

func (s *service) UpdateArsip(input UpdateArsipInput) (ArsipDesa, error) {
	var arsip ArsipDesa

	err := s.Validate.Struct(input)
	if err != nil {
		return arsip, errors.New("isi form dengan benar")
	}

	arsipRow, err := s.repository.FindByID(uint(input.ID))
	if err != nil {
		return arsip, err
	}

	arsip.ID = arsipRow.ID
	arsip.Nama = input.Nama
	arsip.KategoriID = input.KategoriID
	arsip.Deskripsi = input.Deskripsi
	arsip.CreatedAt = arsipRow.CreatedAt

	arsip, err = s.repository.Update(arsip)
	if err != nil {
		return arsip, err
	}

	return arsip, nil
}

func (s *service) DeletedSoft(id uint) (ArsipDesa, error) {
	var arsip ArsipDesa

	file, err := s.fileRepository.FindFileByArsipID(id)
	if err != nil {
		return arsip, err
	}

	if len(file) >= 1 {
		return arsip, errors.New("Arsip tidak dapat di hapus, terdapat File Pada Arsip!")
	}

	arsip, err = s.repository.DeletedSoft(id)
	if err != nil {
		return arsip, err
	}

	return arsip, nil
}

func (s *service) Deleted(id uint) (ArsipDesa, error) {
	var arsip ArsipDesa

	file, err := s.fileRepository.FindFileByArsipID(id)
	if err != nil {
		return arsip, err
	}

	if len(file) >= 1 {
		return arsip, errors.New("Arsip tidak dapat di hapus, terdapat File Pada Arsip!")
	}

	arsip, err = s.repository.Deleted(id)
	if err != nil {
		return arsip, err
	}

	return arsip, nil
}

func (s *service) RestoreArsip(id uint) (ArsipDesa, error) {
	var arsip ArsipDesa

	arsipDeleted, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return arsip, err
	}

	updatedArsip, err := s.repository.UpdateDeletedAt(arsipDeleted.ID)
	if err != nil {
		return arsip, err
	}

	return updatedArsip, nil
}
