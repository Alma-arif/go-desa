package berita

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(berita Berita) (Berita, error)
	FindAllDeletedAt() ([]Berita, error)
	FindAll() ([]Berita, error)
	FindByID(id uint) (Berita, error)
	FindByIDDeletedAt(id uint) (Berita, error)
	Update(berita Berita) (Berita, error)
	UpdateDeletedAt(id uint) (Berita, error)
	DeletedSoft(id uint) (Berita, error)
	Deleted(id uint) (Berita, error)

	FindAllBeritaWeb(perPage, page int) ([]Berita, int64, error)
	FindAllBeritaLimitWeb(page int) ([]Berita, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(berita Berita) (Berita, error) {
	err := r.db.Create(&berita).Error
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) FindAllDeletedAt() ([]Berita, error) {
	var berita []Berita
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&berita).Error
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) FindAll() ([]Berita, error) {
	var berita []Berita
	err := r.db.Find(&berita).Error
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) FindByID(id uint) (Berita, error) {
	var berita Berita

	err := r.db.Where("id = ?", id).Find(&berita).Error
	if err != nil {
		return berita, err
	}

	return berita, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (Berita, error) {
	var berita Berita

	err := r.db.Unscoped().Where("id = ?", id).Find(&berita).Error
	if err != nil {
		return berita, err
	}

	return berita, nil
}

func (r *repository) Update(berita Berita) (Berita, error) {
	err := r.db.Save(&berita).Error

	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) UpdateDeletedAt(id uint) (Berita, error) {
	var berita Berita

	err := r.db.Unscoped().Model(&berita).Where("id", id).Update("deleted_at", nil).Error

	if err != nil {
		return berita, err
	}

	return berita, nil
}

func (r *repository) DeletedSoft(id uint) (Berita, error) {
	var berita Berita
	err := r.db.Where("id = ?", id).Delete(&berita).Error
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) Deleted(id uint) (Berita, error) {
	var berita Berita
	err := r.db.Unscoped().Where("id = ?", id).Delete(&berita).Error
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) FindAllBeritaWeb(perPage, page int) ([]Berita, int64, error) {

	var berita []Berita
	var count int64

	err := r.db.Model(&Berita{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage

	err = r.db.Model(&Berita{}).Where("status = 1").Order("created_at desc").Limit(perPage).Offset(offset).Find(&berita).Error
	if err != nil {
		return nil, 0, err

	}

	return berita, count, nil
}

func (r *repository) FindAllBeritaLimitWeb(page int) ([]Berita, error) {

	var berita []Berita
	// var count int64

	err := r.db.Model(&Berita{}).Where("status = 1").Order("created_at desc").Limit(page).Find(&berita).Error
	if err != nil {
		return nil, err
	}

	return berita, nil
}
