package sejarah

import "gorm.io/gorm"

type Repository interface {
	Save(informasi Sejarah) (Sejarah, error)
	FindAllDeletedAt() ([]Sejarah, error)
	FindAll() ([]Sejarah, error)
	FindAllLimit() (Sejarah, error)
	FindByID(id uint) (Sejarah, error)
	FindByIDDeletedAt(id uint) (Sejarah, error)
	Update(informasi Sejarah) (Sejarah, error)
	UpdateDeletedAt(id uint) (Sejarah, error)
	DeletedSoft(id uint) (Sejarah, error)
	Deleted(id uint) (Sejarah, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(informasi Sejarah) (Sejarah, error) {
	err := r.db.Create(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindAllDeletedAt() ([]Sejarah, error) {
	var informasi []Sejarah
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindAll() ([]Sejarah, error) {
	var informasi []Sejarah
	err := r.db.Find(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindAllLimit() (Sejarah, error) {
	var informasi Sejarah
	err := r.db.Find(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindByID(id uint) (Sejarah, error) {
	var informasi Sejarah

	err := r.db.Where("id = ?", id).Find(&informasi).Error
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (Sejarah, error) {
	var informasi Sejarah

	err := r.db.Unscoped().Where("id = ?", id).Find(&informasi).Error
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (r *repository) Update(informasi Sejarah) (Sejarah, error) {
	err := r.db.Save(&informasi).Error

	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) UpdateDeletedAt(id uint) (Sejarah, error) {
	var informasi Sejarah

	err := r.db.Unscoped().Model(&informasi).Where("id", id).Update("deleted_at", nil).Error

	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (r *repository) DeletedSoft(id uint) (Sejarah, error) {
	var informasi Sejarah
	err := r.db.Where("id = ?", id).Delete(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) Deleted(id uint) (Sejarah, error) {
	var informasi Sejarah
	err := r.db.Unscoped().Where("id = ?", id).Delete(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}
