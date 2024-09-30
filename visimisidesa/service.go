package visimisidesa

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
	CreateVisiMisi(input VisiMisiInput, image string) (VisiMisi, error)
	GetAllVisiMisi() ([]VisiMisiView, error)
	GetAllVisiMisiDeleted() ([]VisiMisiView, error)
	GetVisiMisiByID(id uint) (VisiMisiView, error)
	GetVisiMisiByIDDeleted(id uint) (VisiMisi, error)
	UpdateVisiMisi(input VisiMisiUpdate, image string) (VisiMisi, error)
	DeletedVisiMisiSoft(id uint) (VisiMisi, error)
	DeletedVisiMisi(id uint) (VisiMisi, error)
	DeletedVisiMisiImage(id uint) (VisiMisi, error)
	RestoreVisiMisi(id uint) (VisiMisi, error)

	GetAllVisiMisiWeb() (VisiMisiView, error)
}

type service struct {
	repository     Repository
	Validate       *validator.Validate
	userRepository user.Repository
}

func NewService(repository Repository, Validate *validator.Validate, userRepository user.Repository) *service {
	return &service{repository, Validate, userRepository}
}

func (s *service) CreateVisiMisi(input VisiMisiInput, image string) (VisiMisi, error) {
	var informasi VisiMisi

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

	informasi.VisiMisi = input.VisiMisi
	informasi.Image = image

	_, err = s.repository.Save(informasi)
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (s *service) GetAllVisiMisi() ([]VisiMisiView, error) {
	var informasi []VisiMisiView

	resultInformasi, err := s.repository.FindAll()
	if err != nil {
		return informasi, err
	}

	for i, rowInformasi := range resultInformasi {
		var informasiView VisiMisiView
		informasiView.ID = rowInformasi.ID
		informasiView.Index = i + 1
		informasiView.VisiMisi = template.HTML(rowInformasi.VisiMisi)
		informasiView.Image = rowInformasi.Image
		informasiView.CreatedAt = rowInformasi.CreatedAt
		informasi = append(informasi, informasiView)
	}

	return informasi, nil
}

func (s *service) GetAllVisiMisiDeleted() ([]VisiMisiView, error) {
	var informasi []VisiMisiView

	resultInformasi, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return informasi, err
	}

	for i, rowInformasi := range resultInformasi {
		var informasiView VisiMisiView
		informasiView.ID = rowInformasi.ID
		informasiView.Index = i + 1
		informasiView.VisiMisi = template.HTML(rowInformasi.VisiMisi)
		informasiView.Image = rowInformasi.Image
		informasiView.CreatedAt = rowInformasi.CreatedAt

		informasi = append(informasi, informasiView)
	}

	return informasi, nil
}

func (s *service) GetVisiMisiByID(id uint) (VisiMisiView, error) {
	var informasi VisiMisiView

	rowInformasi, err := s.repository.FindByID(id)
	if err != nil {
		return informasi, err
	}

	informasi.ID = rowInformasi.ID
	informasi.VisiMisi = template.HTML(rowInformasi.VisiMisi)
	informasi.Image = rowInformasi.Image
	informasi.CreatedAt = rowInformasi.CreatedAt

	if informasi.ID == 0 {
		return informasi, errors.New("No Pengumuman found on with that ID")
	}

	return informasi, nil
}

func (s *service) GetVisiMisiByIDDeleted(id uint) (VisiMisi, error) {
	var informasi VisiMisi

	rowInformasi, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return informasi, err
	}

	return rowInformasi, nil
}

func (s *service) UpdateVisiMisi(input VisiMisiUpdate, image string) (VisiMisi, error) {
	var informasi VisiMisi

	err := s.Validate.Struct(input)
	if err != nil {
		return informasi, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	rowInformasi, err := s.repository.FindByID(input.ID)
	if err != nil {
		return informasi, err
	}

	informasi.ID = rowInformasi.ID
	informasi.VisiMisi = input.VisiMisi

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

func (s *service) DeletedVisiMisiSoft(id uint) (VisiMisi, error) {
	var informasi VisiMisi
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
	informasi.VisiMisi = informasiRow.VisiMisi
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

func (s *service) DeletedVisiMisi(id uint) (VisiMisi, error) {
	var informasi VisiMisi

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

func (s *service) DeletedVisiMisiImage(id uint) (VisiMisi, error) {
	var informasi VisiMisi

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
	informasi.VisiMisi = informasiRow.VisiMisi
	informasi.Image = ""
	informasi.CreatedAt = informasiRow.CreatedAt

	_, err = s.repository.Update(informasi)
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (s *service) RestoreVisiMisi(id uint) (VisiMisi, error) {
	var informasi VisiMisi

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

func (s *service) GetAllVisiMisiWeb() (VisiMisiView, error) {
	var informasi VisiMisiView

	resultInformasi, err := s.repository.FindAllLimit()
	if err != nil {
		return informasi, err
	}

	informasi.ID = resultInformasi.ID
	informasi.VisiMisi = template.HTML(resultInformasi.VisiMisi)
	informasi.Image = resultInformasi.Image
	informasi.CreatedAt = resultInformasi.CreatedAt

	return informasi, nil
}
