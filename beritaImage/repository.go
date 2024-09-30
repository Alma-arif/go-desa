package beritaimage

import "gorm.io/gorm"

type Repository interface {
	Save(image ImageBerita) (ImageBerita, error)
	FindAllDeletedAt() ([]ImageBerita, error)
	FindAll() ([]ImageBerita, error)
	FindByID(id uint) (ImageBerita, error)
	FindByBeritaID(id uint) ([]ImageBerita, error)
	FindByBeritaIDEndImagePrimary(id uint) (ImageBerita, error)
	FindByBeritaIDEndImageNoPrimary(id uint) ([]ImageBerita, error)
	FindByIDDeletedAt(id uint) (ImageBerita, error)
	Update(image ImageBerita) (ImageBerita, error)
	UpdateDeletedAt(id uint) (ImageBerita, error)
	DeletedSoft(id uint) (ImageBerita, error)
	Deleted(id uint) (ImageBerita, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(image ImageBerita) (ImageBerita, error) {
	err := r.db.Create(&image).Error
	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) FindAllDeletedAt() ([]ImageBerita, error) {
	var image []ImageBerita
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&image).Error
	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) FindAll() ([]ImageBerita, error) {
	var image []ImageBerita
	err := r.db.Find(&image).Error
	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) FindByID(id uint) (ImageBerita, error) {
	var image ImageBerita

	err := r.db.Where("id = ?", id).Find(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (r *repository) FindByBeritaID(id uint) ([]ImageBerita, error) {
	var image []ImageBerita

	err := r.db.Where("id_berita = ?", id).Find(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (r *repository) FindByBeritaIDEndImagePrimary(id uint) (ImageBerita, error) {
	var image ImageBerita
	err := r.db.Where("id_berita = ?", id).Where("image_utama = 1").Find(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (r *repository) FindByBeritaIDEndImageNoPrimary(id uint) ([]ImageBerita, error) {
	var image []ImageBerita
	err := r.db.Where("id_berita = ?", id).Where("image_utama = 0").Find(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (ImageBerita, error) {
	var image ImageBerita

	err := r.db.Unscoped().Where("id = ?", id).Find(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (r *repository) Update(image ImageBerita) (ImageBerita, error) {
	err := r.db.Save(&image).Error

	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) UpdateDeletedAt(id uint) (ImageBerita, error) {
	var image ImageBerita

	err := r.db.Unscoped().Model(&image).Where("id", id).Update("deleted_at", nil).Error
	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) DeletedSoft(id uint) (ImageBerita, error) {
	var image ImageBerita
	err := r.db.Where("id = ?", id).Delete(&image).Error
	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) Deleted(id uint) (ImageBerita, error) {
	var image ImageBerita
	err := r.db.Unscoped().Where("id = ?", id).Delete(&image).Error
	if err != nil {
		return image, err
	}
	return image, nil

}
