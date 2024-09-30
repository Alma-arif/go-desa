package profiledesa

import (
	"app-desa-kepuk/helper"
	"errors"
	"fmt"
	"html/template"
	"os"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateProfileDesa(input ProfileDesaInput, imageName string) (ProfileDesa, error)
	GetAllProfileDesa() ([]ProfileDesaView, error)
	GetAllProfileDesaDeleted() ([]ProfileDesaView, error)
	GetProfileDesaByID(id uint) (ProfileDesaView, error)
	UpdateProfileDesa(input ProfileDesaUpdate, image string) (ProfileDesa, error)
	DeletedProfileDesaSoft(id uint) (ProfileDesa, error)
	DeletedProfileDesa(id uint) (ProfileDesa, error)
	DeletedProfileDesaImage(id uint) (ProfileDesa, error)
	RestoreProfileDesa(id uint) (ProfileDesa, error)

	GetAllProfileDesaWeb() (ProfileDesaView, error)
}

type service struct {
	repository Repository
	Validate   *validator.Validate
}

func NewService(repository Repository, Validate *validator.Validate) *service {
	return &service{repository, Validate}
}

func (s *service) CreateProfileDesa(input ProfileDesaInput, imageName string) (ProfileDesa, error) {
	var profile ProfileDesa

	err := s.Validate.Struct(input)
	if err != nil {
		return profile, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	resultProfile, err := s.repository.FindAll()
	if err != nil {
		return profile, err
	}

	if len(resultProfile) >= 1 {
		return profile, errors.New("Data yang bisa ditambahkan sudah mencapai batas!")

	}

	// profile.Nama = input.Nama
	profile.ProfileDesa = input.ProfileDesa
	profile.ImageDesa = imageName

	newprofile, err := s.repository.Save(profile)
	if err != nil {
		return newprofile, err
	}

	return newprofile, nil
}

func (s *service) GetAllProfileDesa() ([]ProfileDesaView, error) {
	var profiles []ProfileDesaView

	resultProfile, err := s.repository.FindAll()
	if err != nil {
		return profiles, err
	}

	for i, profile := range resultProfile {
		var ProfileView ProfileDesaView
		ProfileView.ID = profile.ID
		ProfileView.Index = i + 1
		// ProfileView.Nama = profile.Nama
		ProfileView.ProfileDesa = template.HTML(profile.ProfileDesa)
		ProfileView.ImageDesa = profile.ImageDesa
		ProfileView.CreatedAt = profile.CreatedAt
		ProfileView.UpdatedAt = profile.UpdatedAt
		profiles = append(profiles, ProfileView)
	}

	return profiles, nil
}

func (s *service) GetAllProfileDesaDeleted() ([]ProfileDesaView, error) {
	var profileResult []ProfileDesaView

	profiles, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return profileResult, err
	}

	for i, profile := range profiles {
		profilesRow := ProfileDesaView{}
		profilesRow.ID = profile.ID
		profilesRow.Index = i + 1
		// profilesRow.Nama = profile.Nama
		profilesRow.ProfileDesa = template.HTML(profile.ProfileDesa)
		profilesRow.ImageDesa = profile.ImageDesa
		profilesRow.CreatedAt = profile.CreatedAt
		profilesRow.DeletedAt = profile.DeletedAt
		profileResult = append(profileResult, profilesRow)
	}

	return profileResult, nil
}

func (s *service) GetProfileDesaByID(id uint) (ProfileDesaView, error) {
	var resultProfileDesa ProfileDesaView

	profile, err := s.repository.FindByID(id)
	if err != nil {
		return resultProfileDesa, err
	}

	resultProfileDesa.ID = profile.ID
	// resultProfileDesa.Nama = profile.Nama
	resultProfileDesa.ProfileDesa = template.HTML(profile.ProfileDesa)
	resultProfileDesa.ImageDesa = profile.ImageDesa
	resultProfileDesa.CreatedAt = profile.CreatedAt

	return resultProfileDesa, nil
}

func (s *service) UpdateProfileDesa(input ProfileDesaUpdate, image string) (ProfileDesa, error) {
	var profile ProfileDesa

	err := s.Validate.Struct(input)
	if err != nil {
		return profile, errors.New("isi form dengan benar")
	}

	profileRow, err := s.repository.FindByID(uint(input.ID))
	if err != nil {
		return profile, err
	}

	profile.ID = profileRow.ID
	profile.ProfileDesa = input.ProfileDesa

	if image != "" && profileRow.ImageDesa != "" {
		pathRemoveFile := fmt.Sprintf("derektori/images_berita/%s", profileRow.ImageDesa)
		err = os.Remove(pathRemoveFile)
		if err != nil {
			return profile, err
		}

		image = image

	} else if image != "" {
		image = image

	} else {
		image = profileRow.ImageDesa
	}

	profile.ImageDesa = image
	profile.CreatedAt = profileRow.CreatedAt

	profile, err = s.repository.Update(profile)
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (s *service) DeletedProfileDesaSoft(id uint) (ProfileDesa, error) {
	var profile ProfileDesa

	profileRow, err := s.repository.FindByID(id)
	if err != nil {
		return profile, err
	}

	var newFileTo string
	if profileRow.ImageDesa != "" {
		path := fmt.Sprintf("derektori/images_berita/%s", profileRow.ImageDesa)
		_, err = os.Stat(path)
		if err != nil {
			return profile, err
		}

		fileName, err := helper.GetFileNameDekrip(profileRow.ImageDesa)
		if err != nil {
			return profile, err
		}

		newFile, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			return profile, err
		}

		NewPath := fmt.Sprintf("derektori/images_berita/%s", newFile)
		err = os.Rename(path, NewPath)
		if err != nil {
			return profile, err
		}

		newFileTo = newFile
	}

	profile.ID = profileRow.ID
	// profile.Nama = profileRow.Nama
	profile.ProfileDesa = profileRow.ProfileDesa
	profile.ImageDesa = newFileTo
	profile.CreatedAt = profileRow.CreatedAt

	update, err := s.repository.Update(profile)
	if err != nil {
		return update, err
	}

	profile, err = s.repository.DeletedSoft(profileRow.ID)
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (s *service) DeletedProfileDesa(id uint) (ProfileDesa, error) {
	var profile ProfileDesa

	profile, err := s.repository.Deleted(id)
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (s *service) DeletedProfileDesaImage(id uint) (ProfileDesa, error) {
	var informasi ProfileDesa

	informasiRow, err := s.repository.FindByID(id)
	if err != nil {
		return informasi, err
	}
	if informasiRow.ImageDesa != "" {
		pathRemoveFile := fmt.Sprintf("derektori/images_berita/%s", informasiRow.ImageDesa)
		err = os.Remove(pathRemoveFile)
		if err != nil {
			return informasi, err
		}
	}

	informasi.ID = informasiRow.ID
	informasi.ProfileDesa = informasiRow.ProfileDesa
	informasi.ImageDesa = ""
	informasi.CreatedAt = informasiRow.CreatedAt

	_, err = s.repository.Update(informasi)
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (s *service) RestoreProfileDesa(id uint) (ProfileDesa, error) {
	var profile ProfileDesa

	profileDeleted, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return profile, err
	}

	updatedProfile, err := s.repository.UpdateDeletedAt(profileDeleted.ID)
	if err != nil {
		return profile, err
	}

	return updatedProfile, nil
}

func (s *service) GetAllProfileDesaWeb() (ProfileDesaView, error) {
	var profiles ProfileDesaView

	resultProfile, err := s.repository.FindAllLimit()
	if err != nil {
		return profiles, err
	}

	profiles.ID = resultProfile.ID
	profiles.ProfileDesa = template.HTML(resultProfile.ProfileDesa)
	profiles.ImageDesa = resultProfile.ImageDesa
	profiles.CreatedAt = resultProfile.CreatedAt
	profiles.UpdatedAt = resultProfile.UpdatedAt

	return profiles, nil
}
