package profiledesa

import (
	"gorm.io/gorm"
)

type Repository interface {
	// profile desa
	Save(desa ProfileDesa) (ProfileDesa, error)
	FindAll() ([]ProfileDesa, error)
	FindAllLimit() (ProfileDesa, error)
	FindAllDeletedAt() ([]ProfileDesa, error)
	FindByID(id uint) (ProfileDesa, error)
	FindByIDDeletedAt(id uint) (ProfileDesa, error)
	Update(desa ProfileDesa) (ProfileDesa, error)
	UpdateDeletedAt(id uint) (ProfileDesa, error)
	DeletedSoft(id uint) (ProfileDesa, error)
	Deleted(id uint) (ProfileDesa, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(desa ProfileDesa) (ProfileDesa, error) {
	err := r.db.Create(&desa).Error
	if err != nil {
		return desa, err
	}
	return desa, nil
}

func (r *repository) FindAll() ([]ProfileDesa, error) {
	var desa []ProfileDesa
	err := r.db.Find(&desa).Error
	if err != nil {
		return desa, err
	}
	return desa, nil
}

func (r *repository) FindAllLimit() (ProfileDesa, error) {
	var desa ProfileDesa
	err := r.db.Find(&desa).Error
	if err != nil {
		return desa, err
	}
	return desa, nil
}

func (r *repository) FindAllDeletedAt() ([]ProfileDesa, error) {
	var desa []ProfileDesa
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&desa).Error
	if err != nil {
		return desa, err
	}
	return desa, nil
}

func (r *repository) FindByID(id uint) (ProfileDesa, error) {
	var desa ProfileDesa

	err := r.db.Where("id = ?", id).Find(&desa).Error
	if err != nil {
		return desa, err
	}

	return desa, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (ProfileDesa, error) {
	var desa ProfileDesa

	err := r.db.Unscoped().Where("id = ?", id).Find(&desa).Error
	if err != nil {
		return desa, err
	}

	return desa, nil
}

func (r *repository) Update(desa ProfileDesa) (ProfileDesa, error) {
	err := r.db.Save(&desa).Error

	if err != nil {
		return desa, err
	}
	return desa, nil
}

func (r *repository) UpdateDeletedAt(id uint) (ProfileDesa, error) {
	var desa ProfileDesa

	err := r.db.Unscoped().Model(&desa).Where("id", id).Update("deleted_at", nil).Error
	if err != nil {
		return desa, err
	}
	return desa, nil
}

func (r *repository) DeletedSoft(id uint) (ProfileDesa, error) {
	var desa ProfileDesa
	err := r.db.Where("id = ?", id).Delete(&desa).Error
	if err != nil {
		return desa, err
	}
	return desa, nil
}

func (r *repository) Deleted(id uint) (ProfileDesa, error) {
	var desa ProfileDesa
	err := r.db.Unscoped().Where("id = ?", id).Delete(&desa).Error
	if err != nil {
		return desa, err
	}
	return desa, nil

}
