package pegawai

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"errors"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreatePegawai(input PegawaiInput) (Pegawai, error)
	CreateImageProfilePegawai(image string, id uint) (Pegawai, error)
	GetAllPegawai() ([]PegawaiView, error)
	GetAllPegawaiDeletedAt() ([]PegawaiView, error)
	GetPegawaiByID(id uint) (PegawaiView, error)
	GetPegawaiByIDDeleted(id uint) (Pegawai, error)
	UpdatePegawai(input PegawaiUpdate) (Pegawai, error)
	DeletedPegawaiSoft(id uint) (Pegawai, error)
	DeletedPegawai(id uint) (Pegawai, error)
	RestorePegawai(id uint) (Pegawai, error)
}

type service struct {
	repository     Repository
	Validate       *validator.Validate
	userRepository user.Repository
}

func NewService(repository Repository, Validate *validator.Validate, userRepository user.Repository) *service {
	return &service{repository, Validate, userRepository}
}

func (s *service) CreatePegawai(input PegawaiInput) (Pegawai, error) {
	var user Pegawai

	err := s.Validate.Struct(input)

	if err != nil {
		return user, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	user.Nama = input.Nama
	user.Jabatan = input.Jabatan
	user.NoHP = input.NoHP
	user.Alamat = input.Alamat
	user.Image = ""
	date, _ := helper.StringToDate(input.TanggalLahir)
	user.TanggalLahir = date

	_, err = s.repository.Save(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) CreateImageProfilePegawai(image string, id uint) (Pegawai, error) {
	var user Pegawai

	userRow, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}
	user.ID = userRow.ID
	user.Nama = userRow.Nama
	user.Jabatan = userRow.Jabatan
	user.NoHP = userRow.NoHP
	user.Alamat = userRow.Alamat
	user.Image = image
	user.TanggalLahir = userRow.TanggalLahir

	updateUser, err := s.repository.Update(user)
	if err != nil {
		return user, err
	}

	return updateUser, nil
}

func (s *service) GetAllPegawai() ([]PegawaiView, error) {
	var user []PegawaiView

	resultuser, err := s.repository.FindAll()
	if err != nil {
		return user, err
	}

	for i, rowuser := range resultuser {
		var userView PegawaiView
		userView.ID = rowuser.ID
		userView.Index = i + 1
		userView.Nama = rowuser.Nama
		userView.Jabatan = rowuser.Jabatan
		userView.NoHP = rowuser.NoHP
		userView.Alamat = rowuser.Alamat
		userView.Image = rowuser.Image
		userView.TanggalLahir = rowuser.TanggalLahir
		userView.CreatedAt = rowuser.CreatedAt

		user = append(user, userView)
	}

	return user, nil
}

func (s *service) GetAllPegawaiDeletedAt() ([]PegawaiView, error) {
	var user []PegawaiView

	resultuser, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return user, err
	}

	for i, rowuser := range resultuser {
		var userView PegawaiView
		userView.ID = rowuser.ID
		userView.Index = i + 1
		userView.Nama = rowuser.Nama
		userView.Jabatan = rowuser.Jabatan
		userView.NoHP = rowuser.NoHP
		userView.Alamat = rowuser.Alamat
		userView.Image = rowuser.Image
		userView.TanggalLahir = rowuser.TanggalLahir
		userView.CreatedAt = rowuser.CreatedAt
		user = append(user, userView)
	}

	return user, nil
}

func (s *service) GetPegawaiByID(id uint) (PegawaiView, error) {
	var user PegawaiView

	rowuser, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	user.ID = rowuser.ID
	user.Nama = rowuser.Nama
	user.Jabatan = rowuser.Jabatan
	user.NoHP = rowuser.NoHP
	user.Alamat = rowuser.Alamat
	user.Image = rowuser.Image
	user.TanggalLahir = rowuser.TanggalLahir
	user.CreatedAt = rowuser.CreatedAt

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *service) GetPegawaiByIDDeleted(id uint) (Pegawai, error) {
	var user Pegawai

	rowuser, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return user, err
	}

	return rowuser, nil
}

func (s *service) UpdatePegawai(input PegawaiUpdate) (Pegawai, error) {
	var user Pegawai

	err := s.Validate.Struct(input)
	if err != nil {
		return user, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	rowuser, err := s.repository.FindByID(input.ID)
	if err != nil {
		return user, err
	}

	user.ID = rowuser.ID
	user.Nama = input.Nama
	user.Jabatan = input.Jabatan
	user.NoHP = input.NoHP
	user.Alamat = input.Alamat
	date, err := helper.StringToDate(input.TanggalLahir)
	if err != nil {
		return user, errors.New("Tanggal lahir tidak sesuai.")
	}
	user.TanggalLahir = date
	user.Image = rowuser.Image
	user.CreatedAt = rowuser.CreatedAt

	userUpdate, err := s.repository.Update(user)
	if err != nil {
		return user, err
	}
	return userUpdate, nil
}

func (s *service) DeletedPegawaiSoft(id uint) (Pegawai, error) {
	var user Pegawai
	userRow, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}
	data := ""
	// var newFile string
	if userRow.Image != "" {
		path := fmt.Sprintf("derektori/image/%s", userRow.Image)
		_, err = os.Stat(path)
		if err != nil {
			return user, err
		}

		fileName, err := helper.GetFileNameDekrip(userRow.Image)
		if err != nil {
			return user, err
		}

		newFile, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			return user, err
		}
		data = newFile
		NewPath := fmt.Sprintf("derektori/image/%s", newFile)
		err = os.Rename(path, NewPath)
		if err != nil {
			return user, err
		}

	}

	user.ID = userRow.ID
	user.Nama = userRow.Nama
	user.Jabatan = userRow.Jabatan
	user.NoHP = userRow.NoHP
	user.Alamat = userRow.Alamat
	user.TanggalLahir = userRow.TanggalLahir
	user.Image = data
	user.CreatedAt = userRow.CreatedAt

	update, err := s.repository.Update(user)
	if err != nil {
		return update, err
	}

	deleteduser, err := s.repository.DeletedSoft(update.ID)
	if err != nil {
		return user, err
	}

	return deleteduser, nil
}

func (s *service) DeletedPegawai(id uint) (Pegawai, error) {
	var user Pegawai

	user, err := s.repository.Deleted(id)
	if err != nil {
		return user, err
	}
	if user.Image != "" {
		pathRemoveFile := fmt.Sprintf("derektori/image/%s", user.Image)
		err = os.Remove(pathRemoveFile)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func (s *service) RestorePegawai(id uint) (Pegawai, error) {
	var user Pegawai

	userDeleted, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return user, err
	}

	updateuser, err := s.repository.UpdateDeletedAt(userDeleted.ID)
	if err != nil {
		return user, err
	}

	return updateuser, nil
}
