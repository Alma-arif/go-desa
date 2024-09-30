package arsipkategori

import (
	"app-desa-kepuk/helper"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateArsipKategori(input KategoriArsipInput) (KategoriArsip, error)

	GetAllArsipKategori() ([]KategoriArsipView, error)
	GetAllArsipKategoriDeleted() ([]KategoriArsipView, error)

	GetArsipKategoriByID(id uint) (KategoriArsip, error)
	GetArsipKategoriByIDDeleted(id uint) (KategoriArsip, error)

	UpdateArsipKategori(input KategoriArsipUpdate) (KategoriArsip, error)

	DeletedArsipKategoriSoft(id uint) (KategoriArsip, error)
	DeletedArsipKategori(id uint) (KategoriArsip, error)

	RestoreArsipKategori(id uint) (KategoriArsip, error)
}

type service struct {
	repository Repository
	Validate   *validator.Validate
}

func NewService(repository Repository, Validate *validator.Validate) *service {
	return &service{repository, Validate}
}

func (s *service) CreateArsipKategori(input KategoriArsipInput) (KategoriArsip, error) {
	var kategori KategoriArsip

	err := s.Validate.Struct(input)
	if err != nil {
		return kategori, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	kategori.Nama = input.Nama
	kategori.Deskripsi = input.Deskripsi

	newKategori, err := s.repository.Save(kategori)
	if err != nil {
		return newKategori, err
	}

	return newKategori, nil
}

func (s *service) GetAllArsipKategori() ([]KategoriArsipView, error) {
	var kategoris []KategoriArsipView

	resultKategori, err := s.repository.FindAll()
	if err != nil {
		return kategoris, err
	}

	for i, kategori := range resultKategori {
		var kategoriView KategoriArsipView
		kategoriView.ID = kategori.ID
		kategoriView.Index = i + 1
		kategoriView.Nama = kategori.Nama
		kategoriView.Deskripsi = kategori.Deskripsi
		kategoriView.CreatedAt = kategori.CreatedAt
		kategoriView.UpdatedAt = kategori.UpdatedAt
		kategoris = append(kategoris, kategoriView)
	}

	return kategoris, nil
}

func (s *service) GetAllArsipKategoriDeleted() ([]KategoriArsipView, error) {
	var kategoris []KategoriArsipView

	resultKategori, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return kategoris, err
	}

	for i, kategori := range resultKategori {
		var kategoriView KategoriArsipView
		kategoriView.ID = kategori.ID
		kategoriView.Index = i + 1
		kategoriView.Nama = kategori.Nama
		kategoriView.Deskripsi = kategori.Deskripsi
		kategoriView.CreatedAt = kategori.CreatedAt
		kategoriView.UpdatedAt = kategori.UpdatedAt
		kategoritimeDeletedAt := fmt.Sprint(kategori.DeletedAt)
		timeKategoripDelete, err := helper.StringToDateTimeIndoFormat(kategoritimeDeletedAt)
		if err != nil {
			return kategoris, err
		}
		kategoriView.DeletedAt = timeKategoripDelete
		kategoris = append(kategoris, kategoriView)
	}

	return kategoris, nil
}

func (s *service) GetArsipKategoriByID(id uint) (KategoriArsip, error) {
	var kategori KategoriArsip

	kategoriRow, err := s.repository.FindByID(id)
	if err != nil {
		return kategori, err
	}

	return kategoriRow, nil
}

func (s *service) GetArsipKategoriByIDDeleted(id uint) (KategoriArsip, error) {
	var kategori KategoriArsip

	kategoriRow, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return kategori, err
	}

	return kategoriRow, nil
}

func (s *service) UpdateArsipKategori(input KategoriArsipUpdate) (KategoriArsip, error) {
	var kategori KategoriArsip

	err := s.Validate.Struct(input)
	if err != nil {
		return kategori, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	KategoriRow, err := s.repository.FindByID(input.ID)
	if err != nil {
		return kategori, err
	}

	kategori.ID = KategoriRow.ID
	kategori.Nama = input.Nama
	kategori.Deskripsi = input.Deskripsi
	kategoriUpdate, err := s.repository.Update(kategori)
	if err != nil {
		return kategori, err
	}
	return kategoriUpdate, nil
}

func (s *service) DeletedArsipKategoriSoft(id uint) (KategoriArsip, error) {
	var kategori KategoriArsip

	kategori, err := s.repository.DeletedSoft(id)

	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (s *service) DeletedArsipKategori(id uint) (KategoriArsip, error) {
	var kategori KategoriArsip

	kategori, err := s.repository.Deleted(id)

	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (s *service) RestoreArsipKategori(id uint) (KategoriArsip, error) {
	var kategori KategoriArsip

	kategoriDeleted, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return kategori, err
	}

	updatedKategori, err := s.repository.UpdateDeletedAt(kategoriDeleted.ID)
	if err != nil {
		return kategori, err
	}

	return updatedKategori, nil
}
