package berita

import (
	beritaimage "app-desa-kepuk/beritaImage"
	"app-desa-kepuk/beritakategori"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"errors"
	"fmt"
	"html/template"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateBerita(input BeritaInput, userID uint) (Berita, error)
	GetAllBerita() ([]BeritaView, error)
	GetAllBeritaDeleted() ([]BeritaView, error)
	GetBeritaByID(id uint) (BeritaView, error)
	GetBeritaByIDDeleted(id uint) (Berita, error)
	UpdateBerita(input BeritaUpdate) (Berita, error)
	DeletedBeritaSoft(id uint) (Berita, error)
	DeletedBerita(id uint) (Berita, error)
	RestoreBerita(id uint) (Berita, error)

	GetAllBeritaWeb(perPage, page int) ([]BeritaViewWeb, int64, error)
	GetBeritaWebByID(id uint) (BeritaViewWeb, error)
	GetAllBeritaWebLimit(page int) ([]BeritaViewWeb, error)
}

type service struct {
	repository         Repository
	Validate           *validator.Validate
	userRepository     user.Repository
	kategoriRepository beritakategori.Repository
	imagesRepository   beritaimage.Repository
}

func NewService(repository Repository, Validate *validator.Validate, userRepository user.Repository, kategoriRepository beritakategori.Repository, imagesRepository beritaimage.Repository) *service {
	return &service{repository, Validate, userRepository, kategoriRepository, imagesRepository}
}

func (s *service) CreateBerita(input BeritaInput, userID uint) (Berita, error) {
	var berita Berita

	err := s.Validate.Struct(input)
	if err != nil {
		return berita, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	userRow, err := s.userRepository.FindByID(userID)
	if err != nil {
		return berita, err
	}
	berita.Header = strings.Replace(input.Judul, " ", "-", -1)
	berita.Judul = input.Judul
	berita.Berita = input.Berita
	berita.IDUser = userRow.ID
	berita.Status = input.Status
	berita.IdKategori = input.IdKategori
	newBerita, err := s.repository.Save(berita)
	if err != nil {
		return newBerita, err
	}

	return newBerita, nil
}

func (s *service) GetAllBerita() ([]BeritaView, error) {
	var berita []BeritaView

	resultBerita, err := s.repository.FindAll()
	if err != nil {
		return berita, err
	}
	//belum mengambil data user

	for i, rowBerita := range resultBerita {
		var beritaView BeritaView
		beritaView.ID = rowBerita.ID
		beritaView.Index = i + 1
		beritaView.Judul = rowBerita.Judul
		beritaView.Berita = template.HTML(rowBerita.Berita)
		beritaView.IDUser = rowBerita.IDUser

		userRow, err := s.userRepository.FindByID(rowBerita.IDUser)
		if err != nil {
			return berita, err
		}
		beritaView.UserName = userRow.Nama

		image, _ := s.imagesRepository.FindByBeritaIDEndImagePrimary(rowBerita.ID)
		// if err != nil {
		// 	return berita, err
		// }
		beritaView.Image = image.NamaImageFile

		kategori, _ := s.kategoriRepository.FindByID(rowBerita.IdKategori)
		// if err != nil {
		// 	return berita, err
		// }
		beritaView.IdKategori = kategori.ID
		beritaView.Kategori = kategori.Nama

		beritaView.Status = rowBerita.Status
		beritaView.CreatedAt = rowBerita.CreatedAt
		berita = append(berita, beritaView)
	}

	return berita, nil
}

func (s *service) GetAllBeritaDeleted() ([]BeritaView, error) {
	var berita []BeritaView

	resultBerita, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return berita, err
	}

	for i, rowBerita := range resultBerita {
		var beritaView BeritaView
		beritaView.ID = rowBerita.ID
		beritaView.Index = i + 1
		beritaView.Judul = rowBerita.Judul
		beritaView.Berita = template.HTML(rowBerita.Berita)
		beritaView.IDUser = rowBerita.IDUser

		userRow, err := s.userRepository.FindByID(rowBerita.IDUser)
		if err != nil {
			return berita, err
		}
		beritaView.UserName = userRow.Nama

		image, err := s.imagesRepository.FindByBeritaIDEndImagePrimary(rowBerita.ID)
		if err != nil {
			return berita, err
		}
		beritaView.Image = image.NamaImageFile

		kategori, err := s.kategoriRepository.FindByID(rowBerita.IdKategori)
		if err != nil {
			return berita, err
		}
		beritaView.IdKategori = kategori.ID
		beritaView.Kategori = kategori.Nama

		beritaView.Status = rowBerita.Status
		beritaView.CreatedAt = rowBerita.CreatedAt
		beritaDeleteAt := fmt.Sprint(rowBerita.DeletedAt)
		timeBeritaDeletedAt, err := helper.StringToDateTimeIndoFormat(beritaDeleteAt)
		if err != nil {
			return berita, err
		}
		beritaView.DeletedAt = timeBeritaDeletedAt

		berita = append(berita, beritaView)
	}

	return berita, nil
}

func (s *service) GetBeritaByID(id uint) (BeritaView, error) {
	var berita BeritaView

	RowBerita, err := s.repository.FindByID(id)
	if err != nil {
		return berita, err
	}

	berita.ID = RowBerita.ID
	berita.Judul = RowBerita.Judul
	berita.Berita = template.HTML(RowBerita.Berita)
	user, err := s.userRepository.FindByID(RowBerita.IDUser)
	if err != nil {
		return berita, err
	}
	berita.IDUser = RowBerita.IDUser
	berita.UserName = user.Nama
	kategori, _ := s.kategoriRepository.FindByID(RowBerita.IdKategori)
	berita.IdKategori = RowBerita.IdKategori
	berita.Kategori = kategori.Nama
	berita.Status = RowBerita.Status
	berita.CreatedAt = RowBerita.CreatedAt

	if berita.ID == 0 {
		return berita, errors.New("No arisp found on with that ID")
	}

	return berita, nil
}

func (s *service) GetBeritaByIDDeleted(id uint) (Berita, error) {
	var berita Berita

	beritaRow, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return berita, err
	}

	return beritaRow, nil
}

func (s *service) UpdateBerita(input BeritaUpdate) (Berita, error) {
	var berita Berita

	err := s.Validate.Struct(input)
	if err != nil {
		return berita, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	beritaRow, err := s.repository.FindByID(input.ID)
	if err != nil {
		return berita, err
	}

	berita.ID = beritaRow.ID
	berita.Judul = input.Judul
	berita.Berita = input.Berita
	berita.IDUser = beritaRow.IDUser
	berita.Status = input.Status
	berita.IdKategori = input.IdKategori
	berita.CreatedAt = beritaRow.CreatedAt
	beritaUpdate, err := s.repository.Update(berita)
	if err != nil {
		return berita, err
	}
	return beritaUpdate, nil
}

func (s *service) DeletedBeritaSoft(id uint) (Berita, error) {
	var berita Berita

	berita, err := s.repository.DeletedSoft(id)

	if err != nil {
		return berita, err
	}

	return berita, nil
}

func (s *service) DeletedBerita(id uint) (Berita, error) {
	var berita Berita

	berita, err := s.repository.Deleted(id)

	if err != nil {
		return berita, err
	}

	return berita, nil
}

func (s *service) RestoreBerita(id uint) (Berita, error) {
	var berita Berita

	beritaDeleted, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return berita, err
	}

	updatedberita, err := s.repository.UpdateDeletedAt(beritaDeleted.ID)
	if err != nil {
		return berita, err
	}

	return updatedberita, nil
}

func (s *service) GetAllBeritaWeb(perPage, page int) ([]BeritaViewWeb, int64, error) {
	var berita []BeritaViewWeb

	resultBerita, count, err := s.repository.FindAllBeritaWeb(perPage, page)
	if err != nil {
		return berita, 0, err
	}

	for i, rowBerita := range resultBerita {
		var beritaView BeritaViewWeb
		beritaView.ID = rowBerita.ID
		beritaView.Index = i + 1
		beritaView.Header = rowBerita.Header
		beritaView.Judul = rowBerita.Judul
		beritaView.Berita = template.HTML(helper.ShortTextWords(26, helper.StripHtmlRegex(rowBerita.Berita)))
		// helper.StripHtmlRegex(rowBerita.Berita)
		// template.HTML(rowBerita.Berita)
		// helper.ShortTextWords(26, helper.StripHtmlRegex(rowBerita.Berita))
		beritaView.IDUser = rowBerita.IDUser

		userRow, err := s.userRepository.FindByID(rowBerita.IDUser)
		if err != nil {
			return berita, 0, err
		}

		beritaView.UserName = userRow.Nama

		rowImagePrimary, _ := s.imagesRepository.FindByBeritaIDEndImagePrimary(rowBerita.ID)
		beritaView.Image = rowImagePrimary.NamaImageFile

		kategori, _ := s.kategoriRepository.FindByID(rowBerita.IdKategori)
		beritaView.IdKategori = kategori.ID
		beritaView.Kategori = kategori.Nama

		beritaView.Status = rowBerita.Status

		tanggal, err := helper.IndonesiaFormat(rowBerita.CreatedAt)
		if err != nil {
			return berita, 0, err
		}
		beritaView.CreatedAt = tanggal

		berita = append(berita, beritaView)
	}

	return berita, count, nil
}

func (s *service) GetBeritaWebByID(id uint) (BeritaViewWeb, error) {
	var berita BeritaViewWeb

	RowBerita, err := s.repository.FindByID(id)
	if err != nil {
		return berita, err
	}

	berita.ID = RowBerita.ID
	berita.Judul = RowBerita.Judul
	berita.Berita = template.HTML(RowBerita.Berita)
	user, err := s.userRepository.FindByID(RowBerita.IDUser)
	if err != nil {
		return berita, err
	}
	berita.IDUser = RowBerita.IDUser
	berita.UserName = user.Nama
	kategori, _ := s.kategoriRepository.FindByID(RowBerita.IdKategori)
	berita.IdKategori = RowBerita.IdKategori
	berita.Kategori = kategori.Nama

	rowImagePrimary, _ := s.imagesRepository.FindByBeritaIDEndImagePrimary(RowBerita.ID)
	resultImagePrimary, _ := s.imagesRepository.FindByBeritaIDEndImageNoPrimary(RowBerita.ID)
	images := beritaimage.ImageBeritaFormatter(rowImagePrimary, resultImagePrimary)
	berita.ImageBerita = images

	berita.Status = RowBerita.Status

	tanggal, err := helper.IndonesiaFormat(RowBerita.CreatedAt)
	if err != nil {
		return berita, err
	}

	berita.CreatedAt = tanggal

	if berita.ID == 0 {
		return berita, errors.New("No berita found on with that ID")
	}

	return berita, nil
}

func (s *service) GetAllBeritaWebLimit(page int) ([]BeritaViewWeb, error) {
	var berita []BeritaViewWeb

	resultBerita, err := s.repository.FindAllBeritaLimitWeb(page)
	if err != nil {
		return berita, err
	}

	for _, rowBerita := range resultBerita {
		var beritaView BeritaViewWeb
		beritaView.ID = rowBerita.ID
		beritaView.Header = rowBerita.Header
		beritaView.Judul = helper.LimitCharacters(rowBerita.Judul, 35)

		rowImagePrimary, _ := s.imagesRepository.FindByBeritaIDEndImagePrimary(rowBerita.ID)
		beritaView.Image = rowImagePrimary.NamaImageFile

		tanggal, err := helper.IndonesiaFormat(rowBerita.CreatedAt)
		if err != nil {
			return berita, err
		}
		beritaView.Berita = template.HTML(helper.ShortTextWords(26, helper.StripHtmlRegex(rowBerita.Berita)))
		beritaView.CreatedAt = tanggal

		berita = append(berita, beritaView)
	}

	return berita, nil
}
