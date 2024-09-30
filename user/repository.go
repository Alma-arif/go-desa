package user

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindAll() ([]User, error)
	FindAllDeleted() ([]User, error)
	FindByID(id uint) (User, error)
	FindByIDDeleted(id uint) (User, error)
	FindByEmail(email string) (User, error)
	Update(user User) (User, error)
	UpdateDeletedAt(id uint) (User, error)
	DeleteSoft(id uint) (User, error)
	Delete(id uint) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) FindAllDeleted() ([]User, error) {
	var users []User
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) FindByID(id uint) (User, error) {
	var user User

	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	if user.ID <= 0 {
		return user, errors.New("data tidak di temukan")
	}

	return user, nil
}

func (r *repository) FindByIDDeleted(id uint) (User, error) {
	var user User

	err := r.db.Unscoped().Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	if user.ID <= 0 {
		return user, errors.New("data tidak di temukan")
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	if user.ID <= 0 {
		return user, errors.New("data tidak di temukan")
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
func (r *repository) UpdateDeletedAt(id uint) (User, error) {
	var resultUSer User

	err := r.db.Unscoped().Model(&resultUSer).Where("id", id).Update("deleted_at", nil).Error

	// err := r.db.Save(&user).Error

	if err != nil {
		return resultUSer, err
	}

	return resultUSer, nil
}

func (r *repository) DeleteSoft(id uint) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Delete(id uint) (User, error) {
	var user User
	err := r.db.Unscoped().Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
