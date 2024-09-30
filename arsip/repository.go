package arsip

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]ArsipDesa, error)
	FindAllDeletedAt() ([]ArsipDesa, error)
	FindByID(id uint) (ArsipDesa, error)
	FindByIDDeletedAt(id uint) (ArsipDesa, error)
	Save(file ArsipDesa) (ArsipDesa, error)
	Update(file ArsipDesa) (ArsipDesa, error)
	UpdateDeletedAt(id uint) (ArsipDesa, error)
	DeletedSoft(id uint) (ArsipDesa, error)
	Deleted(id uint) (ArsipDesa, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]ArsipDesa, error) {
	var arsipDesa []ArsipDesa

	err := r.db.Find(&arsipDesa).Error

	if err != nil {
		return arsipDesa, err
	}

	return arsipDesa, nil
}

func (r *repository) FindAllDeletedAt() ([]ArsipDesa, error) {
	var arsipDesa []ArsipDesa

	err := r.db.Unscoped().Where("deleted_at > 0").Find(&arsipDesa).Error

	if err != nil {
		return arsipDesa, err
	}

	return arsipDesa, nil
}

func (r *repository) FindByID(id uint) (ArsipDesa, error) {
	var arsipDesa ArsipDesa

	err := r.db.Find(&arsipDesa).Where("id = ?", id).Error

	if err != nil {
		return arsipDesa, err
	}

	return arsipDesa, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (ArsipDesa, error) {
	var arsipDesa ArsipDesa

	err := r.db.Unscoped().Where("id = ?", id).Find(&arsipDesa).Error

	if err != nil {
		return arsipDesa, err
	}

	return arsipDesa, nil
}

func (r *repository) Save(arsip ArsipDesa) (ArsipDesa, error) {
	err := r.db.Create(&arsip).Error
	if err != nil {
		return arsip, err
	}
	return arsip, nil
}

func (r *repository) Update(arsip ArsipDesa) (ArsipDesa, error) {
	err := r.db.Save(&arsip).Error

	if err != nil {
		return arsip, err
	}
	return arsip, nil
}
func (r *repository) UpdateDeletedAt(id uint) (ArsipDesa, error) {
	var arsip ArsipDesa
	err := r.db.Unscoped().Model(&arsip).Where("id", id).Update("deleted_at", nil).Error

	if err != nil {
		return arsip, err
	}
	return arsip, nil
}

func (r *repository) DeletedSoft(id uint) (ArsipDesa, error) {
	var arsip ArsipDesa

	err := r.db.Where("id = ?", id).Delete(&arsip).Error
	if err != nil {
		return arsip, err
	}

	return arsip, nil
}

func (r *repository) Deleted(id uint) (ArsipDesa, error) {
	var arsip ArsipDesa

	err := r.db.Unscoped().Where("id = ?", id).Delete(&arsip).Error
	if err != nil {
		return arsip, err
	}

	return arsip, nil
}
