package demografi

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
	CreateDemografi(input DemografiInput, image string) (Demografi, error)
	GetAllDemografi() ([]DemografiView, error)
	GetAllDemografiDeleted() ([]DemografiView, error)
	GetDemografiByID(id uint) (DemografiView, error)
	GetDemografiByIDDeleted(id uint) (Demografi, error)
	UpdateDemografi(input DemografiUpdate, image string) (Demografi, error)
	DeletedDemografiSoft(id uint) (Demografi, error)
	DeletedDemografi(id uint) (Demografi, error)
	DeletedDemografiImage(id uint) (Demografi, error)
	RestoreDemografi(id uint) (Demografi, error)

	GetAllDemografiWeb() (DemografiView, error)
}

type service struct {
	repository     Repository
	Validate       *validator.Validate
	userRepository user.Repository
}

func NewService(repository Repository, Validate *validator.Validate, userRepository user.Repository) *service {
	return &service{repository, Validate, userRepository}
}

func (s *service) CreateDemografi(input DemografiInput, image string) (Demografi, error) {
	var informasi Demografi

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

	informasi.Demografi = input.Demografi
	informasi.Image = image

	_, err = s.repository.Save(informasi)
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (s *service) GetAllDemografi() ([]DemografiView, error) {
	var informasi []DemografiView

	resultInformasi, err := s.repository.FindAll()
	if err != nil {
		return informasi, err
	}

	for i, rowInformasi := range resultInformasi {
		var informasiView DemografiView
		informasiView.ID = rowInformasi.ID
		informasiView.Index = i + 1
		informasiView.Demografi = template.HTML(rowInformasi.Demografi)
		informasiView.Image = rowInformasi.Image
		informasiView.CreatedAt = rowInformasi.CreatedAt
		informasi = append(informasi, informasiView)
	}

	return informasi, nil
}

func (s *service) GetAllDemografiDeleted() ([]DemografiView, error) {
	var informasi []DemografiView

	resultInformasi, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return informasi, err
	}

	for i, rowInformasi := range resultInformasi {
		var informasiView DemografiView
		informasiView.ID = rowInformasi.ID
		informasiView.Index = i + 1
		informasiView.Demografi = template.HTML(rowInformasi.Demografi)
		informasiView.Image = rowInformasi.Image
		informasiView.CreatedAt = rowInformasi.CreatedAt

		informasi = append(informasi, informasiView)
	}

	return informasi, nil
}

func (s *service) GetDemografiByID(id uint) (DemografiView, error) {
	var informasi DemografiView

	rowInformasi, err := s.repository.FindByID(id)
	if err != nil {
		return informasi, err
	}

	informasi.ID = rowInformasi.ID
	informasi.Demografi = template.HTML(rowInformasi.Demografi)
	informasi.Image = rowInformasi.Image
	informasi.CreatedAt = rowInformasi.CreatedAt

	if informasi.ID == 0 {
		return informasi, errors.New("No Pengumuman found on with that ID")
	}

	return informasi, nil
}

func (s *service) GetDemografiByIDDeleted(id uint) (Demografi, error) {
	var informasi Demografi

	rowInformasi, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return informasi, err
	}

	return rowInformasi, nil
}

func (s *service) UpdateDemografi(input DemografiUpdate, image string) (Demografi, error) {
	var informasi Demografi

	err := s.Validate.Struct(input)
	if err != nil {
		return informasi, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	rowInformasi, err := s.repository.FindByID(input.ID)
	if err != nil {
		return informasi, err
	}

	informasi.ID = rowInformasi.ID
	informasi.Demografi = input.Demografi

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

func (s *service) DeletedDemografiSoft(id uint) (Demografi, error) {
	var informasi Demografi
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
	informasi.Demografi = informasiRow.Demografi
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

func (s *service) DeletedDemografi(id uint) (Demografi, error) {
	var informasi Demografi

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

func (s *service) DeletedDemografiImage(id uint) (Demografi, error) {
	var informasi Demografi

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
	informasi.Demografi = informasiRow.Demografi
	informasi.Image = ""
	informasi.CreatedAt = informasiRow.CreatedAt

	_, err = s.repository.Update(informasi)
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (s *service) RestoreDemografi(id uint) (Demografi, error) {
	var informasi Demografi

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

func (s *service) GetAllDemografiWeb() (DemografiView, error) {
	var informasi DemografiView

	resultInformasi, err := s.repository.FindAllLimit()
	if err != nil {
		return informasi, err
	}

	informasi.ID = resultInformasi.ID
	informasi.Demografi = template.HTML(resultInformasi.Demografi)
	informasi.Image = resultInformasi.Image
	informasi.CreatedAt = resultInformasi.CreatedAt

	return informasi, nil
}
