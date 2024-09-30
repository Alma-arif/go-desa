package pengumuman

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"errors"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	// dasboard
	CreatePengumuman(input PengumumanInput, image string, userID uint) (Pengumuman, error)
	GetAllPengumuman() ([]PengumumanView, error)
	GetAllPengumumanDeleted() ([]PengumumanView, error)
	GetPengumumanByID(id uint) (PengumumanView, error)
	GetPengumumanByIDDeleted(id uint) (Pengumuman, error)
	UpdatePengumuman(input PengumumanUpdate, image string) (Pengumuman, error)
	DeletedPengumumanSoft(id uint) (Pengumuman, error)
	DeletedPengumuman(id uint) (Pengumuman, error)
	RestorePengumuman(id uint) (Pengumuman, error)

	// web
	GetAllPengumumanWeb(perPage, page int) ([]PengumumanViewWeb, int64, error)
	GetAllPengumumanWebLimit(page int) ([]PengumumanViewWeb, error)
}

type service struct {
	repository     Repository
	Validate       *validator.Validate
	userRepository user.Repository
}

func NewService(repository Repository, Validate *validator.Validate, userRepository user.Repository) *service {
	return &service{repository, Validate, userRepository}
}

func (s *service) CreatePengumuman(input PengumumanInput, image string, userID uint) (Pengumuman, error) {
	var informasi Pengumuman

	err := s.Validate.Struct(input)
	if err != nil {
		return informasi, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	userRow, err := s.userRepository.FindByID(userID)
	if err != nil {
		return informasi, err
	}
	informasi.Header = strings.Replace(input.Judul, " ", "-", -1)
	informasi.Judul = input.Judul
	informasi.Pengumuman = input.Pengumuman
	informasi.IDUser = userRow.ID
	informasi.Status = input.Status
	informasi.Image = image
	_, err = s.repository.Save(informasi)
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (s *service) GetAllPengumuman() ([]PengumumanView, error) {
	var informasi []PengumumanView

	resultInformasi, err := s.repository.FindAll()
	if err != nil {
		return informasi, err
	}

	for i, rowInformasi := range resultInformasi {
		var informasiView PengumumanView
		informasiView.ID = rowInformasi.ID
		informasiView.Index = i + 1
		informasiView.Judul = rowInformasi.Judul
		informasiView.Pengumuman = template.HTML(rowInformasi.Pengumuman)
		informasiView.IDUser = rowInformasi.IDUser

		userRow, err := s.userRepository.FindByID(rowInformasi.IDUser)
		if err != nil {
			return informasi, err
		}

		informasiView.Username = userRow.Nama
		informasiView.Image = rowInformasi.Image
		informasiView.Status = rowInformasi.Status
		informasiView.CreatedAt = rowInformasi.CreatedAt
		informasi = append(informasi, informasiView)
	}

	return informasi, nil
}

func (s *service) GetAllPengumumanDeleted() ([]PengumumanView, error) {
	var informasi []PengumumanView

	resultInformasi, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return informasi, err
	}

	for i, rowInformasi := range resultInformasi {
		var informasiView PengumumanView
		informasiView.ID = rowInformasi.ID
		informasiView.Index = i + 1
		informasiView.Judul = rowInformasi.Judul
		informasiView.Pengumuman = template.HTML(rowInformasi.Pengumuman)
		informasiView.IDUser = rowInformasi.IDUser

		userRow, err := s.userRepository.FindByID(rowInformasi.IDUser)
		if err != nil {
			return informasi, err
		}
		informasiView.Username = userRow.Nama

		informasiView.Image = rowInformasi.Image
		informasiView.Status = rowInformasi.Status
		informasiView.CreatedAt = rowInformasi.CreatedAt
		informasiView.DeletedAt = rowInformasi.DeletedAt

		informasi = append(informasi, informasiView)
	}

	return informasi, nil
}

func (s *service) GetPengumumanByID(id uint) (PengumumanView, error) {
	var informasi PengumumanView

	rowInformasi, err := s.repository.FindByID(id)
	if err != nil {
		return informasi, err
	}

	informasi.ID = rowInformasi.ID
	informasi.Judul = rowInformasi.Judul
	informasi.Pengumuman = template.HTML(rowInformasi.Pengumuman)
	informasi.IDUser = rowInformasi.IDUser

	userRow, err := s.userRepository.FindByID(rowInformasi.IDUser)
	if err != nil {
		return informasi, err
	}
	informasi.Username = userRow.Nama

	informasi.Image = rowInformasi.Image
	informasi.Status = rowInformasi.Status
	informasi.CreatedAt = rowInformasi.CreatedAt
	informasi.DeletedAt = rowInformasi.DeletedAt

	if informasi.ID == 0 {
		return informasi, errors.New("No Pengumuman found on with that ID")
	}

	return informasi, nil
}

func (s *service) GetPengumumanByIDDeleted(id uint) (Pengumuman, error) {
	var informasi Pengumuman

	rowInformasi, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return informasi, err
	}

	return rowInformasi, nil
}

func (s *service) UpdatePengumuman(input PengumumanUpdate, image string) (Pengumuman, error) {
	var informasi Pengumuman

	err := s.Validate.Struct(input)
	if err != nil {
		return informasi, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	rowInformasi, err := s.repository.FindByID(input.ID)
	if err != nil {
		return informasi, err
	}

	informasi.ID = rowInformasi.ID
	informasi.Judul = input.Judul
	informasi.Pengumuman = input.Pengumuman
	informasi.IDUser = rowInformasi.IDUser

	if image != "" && rowInformasi.Image != "" {
		pathRemoveFile := fmt.Sprintf("derektori/images_berita/%s", rowInformasi.Image)
		err = os.Remove(pathRemoveFile)
		if err != nil {
			return informasi, err
		}

		image = image

	} else if image != "" {
		image = image

	} else {
		image = rowInformasi.Image
	}

	informasi.Image = image
	informasi.Status = input.Status
	informasi.CreatedAt = rowInformasi.CreatedAt

	informasiUpdate, err := s.repository.Update(informasi)
	if err != nil {
		return informasi, err
	}
	return informasiUpdate, nil
}

func (s *service) DeletedPengumumanSoft(id uint) (Pengumuman, error) {
	var informasi Pengumuman
	informasiRow, err := s.repository.FindByID(id)
	if err != nil {
		return informasi, err
	}

	var newFileTo string
	if informasiRow.Image != "" {
		path := fmt.Sprintf("derektori/images_berita/%s", informasiRow.Image)
		_, err = os.Stat(path)
		if err != nil {
			return informasi, err
		}

		fileName, err := helper.GetFileNameDekrip(informasiRow.Image)
		if err != nil {
			return informasi, err
		}

		newFile, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			return informasi, err
		}

		NewPath := fmt.Sprintf("derektori/images_berita/%s", newFile)
		err = os.Rename(path, NewPath)
		if err != nil {
			return informasi, err
		}

		newFileTo = newFile
	}

	informasi.ID = informasiRow.ID
	informasi.Judul = informasiRow.Judul
	informasi.Pengumuman = informasiRow.Pengumuman
	informasi.IDUser = informasiRow.IDUser
	informasi.Image = newFileTo
	informasi.Status = 0
	informasi.CreatedAt = informasiRow.CreatedAt

	update, err := s.repository.Update(informasi)
	if err != nil {
		return update, err
	}

	deletedimage, err := s.repository.DeletedSoft(update.ID)
	if err != nil {
		return informasi, err
	}

	return deletedimage, nil
}

func (s *service) DeletedPengumuman(id uint) (Pengumuman, error) {
	var informasi Pengumuman

	informasi, err := s.repository.Deleted(id)
	if err != nil {
		return informasi, err
	}
	if informasi.Image != "" {
		pathRemoveFile := fmt.Sprintf("derektori/images_berita/%s", informasi.Image)
		err = os.Remove(pathRemoveFile)
		if err != nil {
			return informasi, err
		}
	}
	return informasi, nil
}

func (s *service) RestorePengumuman(id uint) (Pengumuman, error) {
	var informasi Pengumuman

	informasiDeleted, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return informasi, err
	}

	updateInformasi, err := s.repository.UpdateDeletedAt(informasiDeleted.ID)
	if err != nil {
		return informasi, err
	}

	return updateInformasi, nil
}

func (s *service) GetAllPengumumanWeb(perPage, page int) ([]PengumumanViewWeb, int64, error) {
	var informasi []PengumumanViewWeb

	resultInformasi, count, err := s.repository.FindAllPengumumanWeb(perPage, page)
	if err != nil {
		return informasi, 0, err
	}

	for i, rowInformasi := range resultInformasi {
		var informasiView PengumumanViewWeb
		informasiView.ID = rowInformasi.ID
		informasiView.Index = i + 1
		informasiView.Judul = rowInformasi.Judul
		informasiView.Pengumuman = template.HTML(rowInformasi.Pengumuman)
		informasiView.IDUser = rowInformasi.IDUser

		userRow, err := s.userRepository.FindByID(rowInformasi.IDUser)
		if err != nil {
			return informasi, 0, err
		}
		informasiView.Username = userRow.Nama

		informasiView.Image = rowInformasi.Image
		informasiView.Status = rowInformasi.Status
		tanggal, err := helper.IndonesiaFormat(rowInformasi.CreatedAt)
		if err != nil {
			return informasi, 0, err
		}

		informasiView.CreatedAt = tanggal
		informasi = append(informasi, informasiView)
	}

	return informasi, count, nil
}

func (s *service) GetAllPengumumanWebLimit(page int) ([]PengumumanViewWeb, error) {
	var informasi []PengumumanViewWeb

	resultInformasi, err := s.repository.FindAllPengumumanLimitWeb(page)
	if err != nil {
		return informasi, err
	}

	for i, rowInformasi := range resultInformasi {
		var informasiView PengumumanViewWeb
		informasiView.ID = rowInformasi.ID
		informasiView.Index = i + 1
		informasiView.Judul = rowInformasi.Judul
		informasiView.Pengumuman = template.HTML(rowInformasi.Pengumuman)
		informasiView.IDUser = rowInformasi.IDUser

		userRow, err := s.userRepository.FindByID(rowInformasi.IDUser)
		if err != nil {
			return informasi, err
		}
		informasiView.Username = userRow.Nama

		informasiView.Image = rowInformasi.Image
		informasiView.Status = rowInformasi.Status
		tanggal, err := helper.IndonesiaFormat(rowInformasi.CreatedAt)
		if err != nil {
			return informasi, err
		}

		informasiView.CreatedAt = tanggal
		informasi = append(informasi, informasiView)
	}

	return informasi, nil
}
