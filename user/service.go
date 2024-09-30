package user

import (
	"app-desa-kepuk/helper"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	GetAllUsers() ([]UserView, error)
	GetAllUsersDeleted() ([]UserView, error)
	GetUserByID(id uint) (User, error)
	GetUserByIDDeleted(id uint) (User, error)
	CreateImageProfile(file string, id uint) (User, error)
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	UpdateUser(input UpdateUserInput) (User, error)
	UpdatePassword(input UpdatePasswordInput) (User, error)
	DeleteUserSoft(id uint) (User, error)
	DeleteUser(id uint) (User, error)
	RestoreUser(id uint) (User, error)
}
type service struct {
	repository Repository
	Validate   *validator.Validate
}

func NewService(repository Repository, Validate *validator.Validate) *service {
	return &service{repository, Validate}
}

func (s *service) GetAllUsers() ([]UserView, error) {
	var userResult []UserView
	users, err := s.repository.FindAll()
	if err != nil {
		return userResult, err
	}

	for i, user := range users {
		var userRow UserView
		userRow.ID = user.ID
		userRow.Index = i + 1
		userRow.Nama = user.Nama
		userRow.Email = user.Email
		userRow.NoHp = user.NoHp
		date, err := helper.DateToDateIndoFormat(user.TanggalLahir)
		if err != nil {
			return userResult, err
		}
		userRow.TanggalLahir = date
		userRow.Role = user.Role

		if user.ProfileFile == "" {
			userRow.ProfileFile = "image-user-no-poto.png"
		} else {
			userRow.ProfileFile = user.ProfileFile
		}

		timeUserCreate, err := helper.DateTmeToDateTimeIndoFormat(user.CreatedAt)
		if err != nil {
			return userResult, err
		}
		userRow.CreatedAt = timeUserCreate
		timeUserUpdate, _ := helper.DateTmeToDateTimeIndoFormat(user.UpdatedAt)
		userRow.UpdatedAt = timeUserUpdate
		userResult = append(userResult, userRow)

	}

	return userResult, nil
}

func (s *service) GetAllUsersDeleted() ([]UserView, error) {
	var userResult []UserView
	users, err := s.repository.FindAllDeleted()
	if err != nil {
		return userResult, err
	}

	for i, user := range users {
		var userRow UserView
		userRow.ID = user.ID
		userRow.Index = i + 1
		userRow.Nama = user.Nama
		userRow.Email = user.Email
		userRow.NoHp = user.NoHp

		date, _ := helper.DateToDateIndoFormat(user.TanggalLahir)
		userRow.TanggalLahir = date
		userRow.Role = user.Role
		if user.ProfileFile == "" {
			userRow.ProfileFile = "image-user-no-poto.png"
		} else {
			userRow.ProfileFile = user.ProfileFile
		}
		timeUserCreate, err := helper.DateTmeToDateTimeIndoFormat(user.CreatedAt)
		if err != nil {
			return userResult, err
		}
		userRow.CreatedAt = timeUserCreate

		usertimeDeletedAt := fmt.Sprint(user.DeletedAt)
		timeUserDelete, err := helper.StringToDateTimeIndoFormat(usertimeDeletedAt)
		if err != nil {
			return userResult, err
		}
		userRow.DeletedAt = timeUserDelete

		userResult = append(userResult, userRow)

	}

	return userResult, nil
}

func (s *service) GetUserByID(id uint) (User, error) {
	var user User
	user, err := s.repository.FindByID(id)

	if err != nil {
		return user, err
	}

	date, err := helper.DateToDateIndoFormat(user.TanggalLahir)
	if err != nil {
		return user, err
	}
	user.TanggalLahir = date
	timeUserCreate, err := helper.DateTmeToDateTimeIndoFormat(user.CreatedAt)
	if err != nil {
		return user, err
	}
	user.CreatedAt = timeUserCreate

	if user.ProfileFile == "" {
		user.ProfileFile = "image-user-no-poto.png"
	}

	if user.ID == 0 {
		return user, errors.New("No user found no with that ID")
	}

	return user, nil
}

func (s *service) GetUserByIDDeleted(id uint) (User, error) {
	var user User
	user, err := s.repository.FindByIDDeleted(id)

	if err != nil {
		return user, err
	}

	// tanggal, _ := time.Parse(time.RFC3339, user.TanggalLahir)
	// user.TanggalLahir = tanggal.Format("2006-01-02")

	date, _ := helper.DateToDateIndoFormat(user.TanggalLahir)
	user.TanggalLahir = date

	if user.ID == 0 {
		return user, errors.New("No user found no with that ID")
	}

	return user, nil
}

func (s *service) CreateImageProfile(file string, id uint) (User, error) {
	var user User

	userRow, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}
	user.ID = userRow.ID
	user.Nama = userRow.Nama
	user.Email = userRow.Email
	user.Password = userRow.Password
	user.ProfileFile = file
	user.TanggalLahir = userRow.TanggalLahir
	user.NoHp = userRow.NoHp
	user.Role = userRow.Role

	updateUser, err := s.repository.Update(user)
	if err != nil {
		return user, err
	}

	return updateUser, nil
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	var user User

	err := s.Validate.Struct(input)
	if err != nil {
		return user, errors.New("isi form dengan benar!")
	}

	if input.Password != input.PasswordRetype {
		return user, errors.New("Password yang anda masukan salah")
	}

	userEmailAfalabel, _ := s.repository.FindByEmail(input.Email)

	if userEmailAfalabel.ID != 0 {
		return user, errors.New("Email sudah pernah digunakan!")
	}

	user.Nama = input.Nama
	user.Email = input.Email
	user.NoHp = input.NoHp
	date, err := helper.StringToDate(input.TanggalLahir)
	if err != nil {
		return user, errors.New("Tangal Lahir tidak sesuai.")
	}
	user.TanggalLahir = date
	passwordHash := helper.GenerateHMAC(input.Password)
	user.Password = passwordHash
	user.Role = "user"

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	var user User
	start := time.Now()

	err := s.Validate.Struct(input)
	if err != nil {
		return user, errors.New("isi form dengan benar")
	}

	email := input.Email
	password := input.Password

	user, err = s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Pengguna dengan email tersebut tidak di temaukan")
	}

	passwordHash := helper.GenerateHMACTest(password)
	if user.Password != passwordHash {

		return user, errors.New("Password tidak sesuai")
	}

	duration := time.Since(start)

	fmt.Println("Waktu Proses login dan Hasing (detik) : ", duration.Seconds())
	return user, nil
}

func (s *service) UpdateUser(input UpdateUserInput) (User, error) {
	var user User

	err := s.Validate.Struct(input)
	if err != nil {
		return user, errors.New("isi form dengan benar")
	}

	userRow, err := s.repository.FindByID(input.ID)
	if err != nil {
		return user, err
	}

	user.ID = userRow.ID
	user.Nama = input.Nama
	user.Email = input.Email
	user.Password = userRow.Password
	date, _ := helper.StringToDate(input.TanggalLahir)
	user.TanggalLahir = date
	user.NoHp = strings.Replace(input.NoHp, "_", "", -1)
	user.Role = input.Role
	user.CreatedAt = userRow.CreatedAt
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) UpdatePassword(input UpdatePasswordInput) (User, error) {
	var user User

	err := s.Validate.Struct(input)
	if err != nil {
		return user, errors.New("isi form dengan benar")
	}

	userRow, err := s.repository.FindByID(input.ID)
	if err != nil {
		return user, err
	}

	if input.Password != input.PasswordRetype {
		return user, errors.New("Password salah")
	}

	user.Nama = userRow.Nama
	user.Email = userRow.Email
	user.Password = helper.GenerateHMAC(input.Password)
	user.TanggalLahir = userRow.TanggalLahir
	user.NoHp = userRow.NoHp
	user.Role = userRow.Role

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) DeleteUserSoft(id uint) (User, error) {
	var user User

	user, err := s.repository.DeleteSoft(id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) DeleteUser(id uint) (User, error) {
	var user User

	user, err := s.repository.Delete(id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) RestoreUser(id uint) (User, error) {
	var user User

	userDeleted, err := s.repository.FindByIDDeleted(id)
	if err != nil {
		return user, err
	}

	updatedUser, err := s.repository.UpdateDeletedAt(userDeleted.ID)
	if err != nil {
		return user, err
	}

	return updatedUser, nil
}
