package beritakategori

import (
	"app-desa-kepuk/helper"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateBeritaKategori(input KategoriBeritaInput) (KategoriBerita, error)

	GetAllBeritaKategori() ([]KategoriBeritaView, error)
	GetAllBeritaKategoriDeleted() ([]KategoriBeritaView, error)

	GetBeritaKategoriByID(id uint) (KategoriBerita, error)
	GetBeritaKategoriByIDDeleted(id uint) (KategoriBerita, error)

	UpdateBeritaKategori(input KategoriBeritaUpdate) (KategoriBerita, error)

	DeletedBeritaKategoriSoft(id uint) (KategoriBerita, error)
	DeletedBeritaKategori(id uint) (KategoriBerita, error)

	RestoreBeritaKategori(id uint) (KategoriBerita, error)
}

type service struct {
	repository Repository
	Validate   *validator.Validate
}

func NewService(repository Repository, Validate *validator.Validate) *service {
	return &service{repository, Validate}
}

func (s *service) CreateBeritaKategori(input KategoriBeritaInput) (KategoriBerita, error) {
	var kategori KategoriBerita

	err := s.Validate.Struct(input)
	if err != nil {
		return kategori, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	kategori.Nama = input.Nama
	kategori.Deskripsi = input.Deskripsi
	// kategori.Status = input.Status

	newKategori, err := s.repository.Save(kategori)
	if err != nil {
		return newKategori, err
	}

	return newKategori, nil
}

func (s *service) GetAllBeritaKategori() ([]KategoriBeritaView, error) {
	var kategoris []KategoriBeritaView

	resultKategori, err := s.repository.FindAll()
	if err != nil {
		return kategoris, err
	}

	for i, kategori := range resultKategori {
		var kategoriView KategoriBeritaView
		kategoriView.ID = kategori.ID
		kategoriView.Index = i + 1
		kategoriView.Nama = kategori.Nama
		kategoriView.Deskripsi = kategori.Deskripsi
		// kategoriView.Status = kategori.Status
		kategoriView.CreatedAt = kategori.CreatedAt
		kategoriView.UpdatedAt = kategori.UpdatedAt
		kategoris = append(kategoris, kategoriView)
	}

	return kategoris, nil
}

func (s *service) GetAllBeritaKategoriDeleted() ([]KategoriBeritaView, error) {
	var kategoris []KategoriBeritaView

	resultKategori, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return kategoris, err
	}

	for i, kategori := range resultKategori {
		var kategoriView KategoriBeritaView
		kategoriView.ID = kategori.ID
		kategoriView.Index = i + 1
		kategoriView.Nama = kategori.Nama
		kategoriView.Deskripsi = kategori.Deskripsi
		// kategoriView.Status = kategori.Status
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

func (s *service) GetBeritaKategoriByID(id uint) (KategoriBerita, error) {
	var kategori KategoriBerita

	kategoriRow, err := s.repository.FindByID(id)
	if err != nil {
		return kategori, err
	}

	return kategoriRow, nil
}

func (s *service) GetBeritaKategoriByIDDeleted(id uint) (KategoriBerita, error) {
	var kategori KategoriBerita

	kategoriRow, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return kategori, err
	}

	return kategoriRow, nil
}

func (s *service) UpdateBeritaKategori(input KategoriBeritaUpdate) (KategoriBerita, error) {
	var kategori KategoriBerita

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
	// kategori.Status = input.Status
	kategoriUpdate, err := s.repository.Update(kategori)
	if err != nil {
		return kategori, err
	}
	return kategoriUpdate, nil
}

func (s *service) DeletedBeritaKategoriSoft(id uint) (KategoriBerita, error) {
	var kategori KategoriBerita

	kategori, err := s.repository.DeletedSoft(id)

	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (s *service) DeletedBeritaKategori(id uint) (KategoriBerita, error) {
	var kategori KategoriBerita

	kategori, err := s.repository.Deleted(id)

	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (s *service) RestoreBeritaKategori(id uint) (KategoriBerita, error) {
	var kategori KategoriBerita

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
