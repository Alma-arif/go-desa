package sejarah

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"errors"
	"fmt"
	"html/template"
	"os"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateSejarah(input SejarahInput, image string) (Sejarah, error)
	GetAllSejarah() ([]SejarahView, error)
	GetAllSejarahDeleted() ([]SejarahView, error)
	GetSejarahByID(id uint) (SejarahView, error)
	GetSejarahByIDDeleted(id uint) (Sejarah, error)
	UpdateSejarah(input SejarahUpdate, image string) (Sejarah, error)
	DeletedSejarahSoft(id uint) (Sejarah, error)
	DeletedSejarah(id uint) (Sejarah, error)
	DeletedSejarahImage(id uint) (Sejarah, error)
	RestoreSejarah(id uint) (Sejarah, error)

	GetAllSejarahWeb() (SejarahView, error)
}

type service struct {
	repository     Repository
	Validate       *validator.Validate
	userRepository user.Repository
}

func NewService(repository Repository, Validate *validator.Validate, userRepository user.Repository) *service {
	return &service{repository, Validate, userRepository}
}

func (s *service) CreateSejarah(input SejarahInput, image string) (Sejarah, error) {
	var informasi Sejarah

	err := s.Validate.Struct(input)
	if err != nil {
		return informasi, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	result, err := s.repository.FindAll()
	if err != nil {
		return informasi, err
	}

	if len(result) >= 1 {
		return informasi, errors.New("Data yang bisa ditambahkan sudah mencapai batas!")
	}

	informasi.Sejarah = input.Sejarah
	informasi.Image = image

	_, err = s.repository.Save(informasi)
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (s *service) GetAllSejarah() ([]SejarahView, error) {
	var informasi []SejarahView

	resultInformasi, err := s.repository.FindAll()
	if err != nil {
		return informasi, err
	}

	for i, rowInformasi := range resultInformasi {
		var informasiView SejarahView
		informasiView.ID = rowInformasi.ID
		informasiView.Index = i + 1
		informasiView.Sejarah = template.HTML(rowInformasi.Sejarah)
		informasiView.Image = rowInformasi.Image
		informasiView.CreatedAt = rowInformasi.CreatedAt
		informasi = append(informasi, informasiView)
	}

	return informasi, nil
}

func (s *service) GetAllSejarahDeleted() ([]SejarahView, error) {
	var informasi []SejarahView

	resultInformasi, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return informasi, err
	}

	for i, rowInformasi := range resultInformasi {
		var informasiView SejarahView
		informasiView.ID = rowInformasi.ID
		informasiView.Index = i + 1
		informasiView.Sejarah = template.HTML(rowInformasi.Sejarah)
		informasiView.Image = rowInformasi.Image
		informasiView.CreatedAt = rowInformasi.CreatedAt

		informasi = append(informasi, informasiView)
	}

	return informasi, nil
}

func (s *service) GetSejarahByID(id uint) (SejarahView, error) {
	var informasi SejarahView

	rowInformasi, err := s.repository.FindByID(id)
	if err != nil {
		return informasi, err
	}

	informasi.ID = rowInformasi.ID
	informasi.Sejarah = template.HTML(rowInformasi.Sejarah)
	informasi.Image = rowInformasi.Image
	informasi.CreatedAt = rowInformasi.CreatedAt

	if informasi.ID == 0 {
		return informasi, errors.New("No Pengumuman found on with that ID")
	}

	return informasi, nil
}

func (s *service) GetSejarahByIDDeleted(id uint) (Sejarah, error) {
	var informasi Sejarah

	rowInformasi, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return informasi, err
	}

	return rowInformasi, nil
}

func (s *service) UpdateSejarah(input SejarahUpdate, image string) (Sejarah, error) {
	var informasi Sejarah

	err := s.Validate.Struct(input)
	if err != nil {
		return informasi, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	rowInformasi, err := s.repository.FindByID(input.ID)
	if err != nil {
		return informasi, err
	}

	informasi.ID = rowInformasi.ID
	informasi.Sejarah = input.Sejarah

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

	informasi.CreatedAt = rowInformasi.CreatedAt

	informasiUpdate, err := s.repository.Update(informasi)
	if err != nil {
		return informasi, err
	}
	return informasiUpdate, nil
}

func (s *service) DeletedSejarahSoft(id uint) (Sejarah, error) {
	var informasi Sejarah
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
	informasi.Sejarah = informasiRow.Sejarah
	informasi.Image = newFileTo
	informasi.CreatedAt = informasiRow.CreatedAt

	update, err := s.repository.Update(informasi)
	if err != nil {
		return update, err
	}

	deletedimage, err := s.repository.DeletedSoft(informasi.ID)
	if err != nil {
		return informasi, err
	}

	return deletedimage, nil
}

func (s *service) DeletedSejarah(id uint) (Sejarah, error) {
	var informasi Sejarah

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

func (s *service) DeletedSejarahImage(id uint) (Sejarah, error) {
	var informasi Sejarah

	informasiRow, err := s.repository.FindByID(id)
	if err != nil {
		return informasi, err
	}
	if informasiRow.Image != "" {
		pathRemoveFile := fmt.Sprintf("derektori/images_berita/%s", informasiRow.Image)
		err = os.Remove(pathRemoveFile)
		if err != nil {
			return informasi, err
		}
	}

	informasi.ID = informasiRow.ID
	informasi.Sejarah = informasiRow.Sejarah
	informasi.Image = ""
	informasi.CreatedAt = informasiRow.CreatedAt

	_, err = s.repository.Update(informasi)
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (s *service) RestoreSejarah(id uint) (Sejarah, error) {
	var informasi Sejarah

	result, err := s.repository.FindAll()
	if err != nil {
		return informasi, err
	}

	if len(result) >= 1 {
		return informasi, errors.New("Data tidak bisa dipulihkan, Hapus data yang masih aktif!")
	}

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

func (s *service) GetAllSejarahWeb() (SejarahView, error) {
	var informasi SejarahView

	resultInformasi, err := s.repository.FindAllLimit()
	if err != nil {
		return informasi, err
	}

	informasi.ID = resultInformasi.ID
	informasi.Sejarah = template.HTML(resultInformasi.Sejarah)
	informasi.Image = resultInformasi.Image
	informasi.CreatedAt = resultInformasi.CreatedAt

	return informasi, nil
}
