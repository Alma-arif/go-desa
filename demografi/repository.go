package demografi

import "gorm.io/gorm"

type Repository interface {
	Save(informasi Demografi) (Demografi, error)
	FindAllDeletedAt() ([]Demografi, error)
	FindAll() ([]Demografi, error)
	FindAllLimit() (Demografi, error)
	FindByID(id uint) (Demografi, error)
	FindByIDDeletedAt(id uint) (Demografi, error)
	Update(informasi Demografi) (Demografi, error)
	UpdateDeletedAt(id uint) (Demografi, error)
	DeletedSoft(id uint) (Demografi, error)
	Deleted(id uint) (Demografi, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(informasi Demografi) (Demografi, error) {
	err := r.db.Create(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindAllDeletedAt() ([]Demografi, error) {
	var informasi []Demografi
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindAll() ([]Demografi, error) {
	var informasi []Demografi
	err := r.db.Find(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindAllLimit() (Demografi, error) {
	var informasi Demografi
	err := r.db.Find(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindByID(id uint) (Demografi, error) {
	var informasi Demografi

	err := r.db.Where("id = ?", id).Find(&informasi).Error
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (Demografi, error) {
	var informasi Demografi

	err := r.db.Unscoped().Where("id = ?", id).Find(&informasi).Error
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (r *repository) Update(informasi Demografi) (Demografi, error) {
	err := r.db.Save(&informasi).Error

	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) UpdateDeletedAt(id uint) (Demografi, error) {
	var informasi Demografi

	err := r.db.Unscoped().Model(&informasi).Where("id", id).Update("deleted_at", nil).Error

	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (r *repository) DeletedSoft(id uint) (Demografi, error) {
	var informasi Demografi
	err := r.db.Where("id = ?", id).Delete(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) Deleted(id uint) (Demografi, error) {
	var informasi Demografi
	err := r.db.Unscoped().Where("id = ?", id).Delete(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}
