package slideshow

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Save(image ImageSlideShow) (ImageSlideShow, error)
	FindAll() ([]ImageSlideShow, error)
	FindAllDeletedAt() ([]ImageSlideShow, error)
	FindByID(id uint) (ImageSlideShow, error)
	FindByIDDeletedAt(id uint) (ImageSlideShow, error)
	Update(image ImageSlideShow) (ImageSlideShow, error)
	UpdateDeletedAt(id uint) (ImageSlideShow, error)
	DeletedSoft(id uint) (ImageSlideShow, error)
	Deleted(id uint) (ImageSlideShow, error)

	FindBySlideShowIDEndImagePrimary() (ImageSlideShow, error)
	FindBySlideShowIDEndImageNoPrimary() ([]ImageSlideShow, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(image ImageSlideShow) (ImageSlideShow, error) {
	err := r.db.Create(&image).Error
	if err != nil {
		return image, err
	}

	fmt.Println("3 :", image)

	return image, nil
}

func (r *repository) FindAll() ([]ImageSlideShow, error) {
	var image []ImageSlideShow
	err := r.db.Find(&image).Error
	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) FindAllDeletedAt() ([]ImageSlideShow, error) {
	var image []ImageSlideShow
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&image).Error
	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) FindByID(id uint) (ImageSlideShow, error) {
	var image ImageSlideShow

	err := r.db.Where("id = ?", id).Find(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (ImageSlideShow, error) {
	var image ImageSlideShow

	err := r.db.Unscoped().Where("id = ?", id).Find(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (r *repository) Update(image ImageSlideShow) (ImageSlideShow, error) {
	err := r.db.Save(&image).Error

	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) UpdateDeletedAt(id uint) (ImageSlideShow, error) {
	var image ImageSlideShow

	err := r.db.Unscoped().Model(&image).Where("id", id).Update("deleted_at", nil).Error
	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) DeletedSoft(id uint) (ImageSlideShow, error) {
	var image ImageSlideShow
	err := r.db.Where("id = ?", id).Delete(&image).Error
	if err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) Deleted(id uint) (ImageSlideShow, error) {
	var image ImageSlideShow
	err := r.db.Unscoped().Where("id = ?", id).Delete(&image).Error
	if err != nil {
		return image, err
	}
	return image, nil

}

func (r *repository) FindAllSlideShowWeb(page int) ([]ImageSlideShow, error) {

	var image []ImageSlideShow

	err := r.db.Debug().Model(&ImageSlideShow{}).Order("created_at desc").Limit(page).Find(&image).Error
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (r *repository) FindBySlideShowIDEndImagePrimary() (ImageSlideShow, error) {
	var image ImageSlideShow
	// err := r.db.Debug().Model(&ImageSlideShow{}).Order("created_at desc").Limit(page).Find(&image).Error
	err := r.db.Where("utama = 1").Find(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (r *repository) FindBySlideShowIDEndImageNoPrimary() ([]ImageSlideShow, error) {
	var image []ImageSlideShow
	err := r.db.Model(&ImageSlideShow{}).Where("utama = 0").Find(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}
