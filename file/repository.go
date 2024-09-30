package file

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]File, error)
	FindAllDeletedAt() ([]File, error)
	FindByID(id uint) (File, error)
	FindByIDDeletedAt(id uint) (File, error)
	FindFileByArsipID(arsipID uint) ([]File, error)
	FindFileByArsipIDNull() ([]File, error)
	Save(file File) (File, error)
	Update(file File) (File, error)
	UpdateDeletetAt(id uint) (File, error)
	DeletedSoft(id uint) (File, error)
	Deleted(id uint) (File, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]File, error) {
	var file []File

	err := r.db.Find(&file).Error

	if err != nil {
		return file, err
	}

	return file, nil

}

func (r *repository) FindAllDeletedAt() ([]File, error) {
	var file []File

	err := r.db.Unscoped().Where("deleted_at > 0").Find(&file).Error

	if err != nil {
		return file, err
	}

	return file, nil

}

func (r *repository) FindByID(id uint) (File, error) {
	var file File

	err := r.db.Where("id = ?", id).Find(&file).Error

	if err != nil {
		return file, err
	}

	return file, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (File, error) {
	var file File
	err := r.db.Unscoped().Where("id = ?", id).Find(&file).Debug().Error
	if err != nil {
		return file, err
	}

	return file, nil
}

func (r *repository) FindFileByArsipID(arsipID uint) ([]File, error) {
	var file []File

	err := r.db.Where("arsip_id = ?", arsipID).Find(&file).Error
	if err != nil {
		return file, err
	}

	return file, nil
}

func (r *repository) FindFileByArsipIDNull() ([]File, error) {
	var file []File

	err := r.db.Where("arsip_id = 0").Order("created_at desc").Find(&file).Error
	if err != nil {
		return file, err
	}

	return file, nil
}

func (r *repository) Save(file File) (File, error) {

	err := r.db.Create(&file).Error
	if err != nil {
		return file, err
	}

	return file, nil
}

func (r *repository) Update(file File) (File, error) {

	err := r.db.Save(&file).Error

	if err != nil {
		return file, err
	}

	return file, nil
}

func (r *repository) UpdateDeletetAt(id uint) (File, error) {
	var file File
	err := r.db.Unscoped().Model(&file).Where("id", id).Update("deleted_at", nil).Error

	if err != nil {
		return file, err
	}

	return file, nil
}

func (r *repository) DeletedSoft(id uint) (File, error) {
	var file File

	err := r.db.Where("id = ?", id).Delete(&file).Error
	if err != nil {
		return file, err
	}

	return file, nil
}

func (r *repository) Deleted(id uint) (File, error) {
	var file File

	err := r.db.Unscoped().Where("id = ?", id).Delete(&file).Error
	if err != nil {
		return file, err
	}

	return file, nil
}
