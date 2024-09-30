package pegawai

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user Pegawai) (Pegawai, error)
	FindAll() ([]Pegawai, error)
	FindAllDeletedAt() ([]Pegawai, error)
	FindByID(id uint) (Pegawai, error)
	FindByIDDeletedAt(id uint) (Pegawai, error)
	Update(user Pegawai) (Pegawai, error)
	UpdateDeletedAt(id uint) (Pegawai, error)
	DeletedSoft(id uint) (Pegawai, error)
	Deleted(id uint) (Pegawai, error)

	// FindBySlideShowIDEnduserPrimary() (Pegawai, error)
	// FindBySlideShowIDEnduserNoPrimary() ([]Pegawai, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user Pegawai) (Pegawai, error) {
	err := r.db.Create(&user).Error
	// fmt.Println("c 1 :", err)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAll() ([]Pegawai, error) {
	var user []Pegawai
	err := r.db.Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindAllDeletedAt() ([]Pegawai, error) {
	var user []Pegawai
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(id uint) (Pegawai, error) {
	var user Pegawai

	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (Pegawai, error) {
	var user Pegawai

	err := r.db.Unscoped().Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user Pegawai) (Pegawai, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) UpdateDeletedAt(id uint) (Pegawai, error) {
	var user Pegawai

	err := r.db.Unscoped().Model(&user).Where("id", id).Update("deleted_at", nil).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) DeletedSoft(id uint) (Pegawai, error) {
	var user Pegawai
	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Deleted(id uint) (Pegawai, error) {
	var user Pegawai
	err := r.db.Unscoped().Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *repository) FindAllSlideShowWeb(page int) ([]Pegawai, error) {

	var user []Pegawai

	err := r.db.Debug().Model(&Pegawai{}).Order("created_at desc").Limit(page).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) FindBySlideShowIDEnduserPrimary() (Pegawai, error) {
	var user Pegawai
	// err := r.db.Debug().Model(&Pegawai{}).Order("created_at desc").Limit(page).Find(&user).Error
	err := r.db.Where("utama = 1").Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindBySlideShowIDEnduserNoPrimary() ([]Pegawai, error) {
	var user []Pegawai
	err := r.db.Model(&Pegawai{}).Where("utama = 0").Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
