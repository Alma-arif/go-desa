package surat

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(pesan Surat) (Surat, error)
	FindAll() ([]Surat, error)
	FindAllDeletedAt() ([]Surat, error)
	FindByID(id uint) (Surat, error)
	FindByKodeSurat(kodeSuart uint) (Surat, error)

	FindByIDDeletedAt(id uint) (Surat, error)
	Update(pesan Surat) (Surat, error)
	UpdateDeletedAt(id uint) (Surat, error)
	DeletedSoft(id uint) (Surat, error)
	Deleted(id uint) (Surat, error)

	// FindBySlideShowIDEndpesanPrimary() (Surat, error)
	// FindBySlideShowIDEndpesanNoPrimary() ([]Surat, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(pesan Surat) (Surat, error) {
	err := r.db.Create(&pesan).Error
	if err != nil {
		return pesan, err
	}

	return pesan, nil
}

func (r *repository) FindAll() ([]Surat, error) {
	var pesan []Surat
	err := r.db.Find(&pesan).Error
	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) FindAllDeletedAt() ([]Surat, error) {
	var pesan []Surat
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&pesan).Error
	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) FindByID(id uint) (Surat, error) {
	var pesan Surat

	err := r.db.Where("id = ?", id).Find(&pesan).Error
	if err != nil {
		return pesan, err
	}

	return pesan, nil
}

func (r *repository) FindByKodeSurat(kodeSuart uint) (Surat, error) {
	var pesan Surat

	err := r.db.Model(&Surat{}).Where("kode_surat = ?", kodeSuart).Order("no_surat DESC").Limit(1).First(&pesan).Error
	if err != nil {
		return pesan, err
	}

	return pesan, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (Surat, error) {
	var pesan Surat

	err := r.db.Unscoped().Where("id = ?", id).Find(&pesan).Error
	if err != nil {
		return pesan, err
	}

	return pesan, nil
}

func (r *repository) Update(pesan Surat) (Surat, error) {
	err := r.db.Save(&pesan).Error

	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) UpdateDeletedAt(id uint) (Surat, error) {
	var pesan Surat

	err := r.db.Unscoped().Model(&pesan).Where("id", id).Update("deleted_at", nil).Error
	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) DeletedSoft(id uint) (Surat, error) {
	var pesan Surat
	err := r.db.Where("id = ?", id).Delete(&pesan).Error
	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) Deleted(id uint) (Surat, error) {
	var pesan Surat
	err := r.db.Unscoped().Where("id = ?", id).Delete(&pesan).Error
	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) FindAllSlideShowWeb(page int) ([]Surat, error) {

	var pesan []Surat

	err := r.db.Debug().Model(&Surat{}).Order("created_at desc").Limit(page).Find(&pesan).Error
	if err != nil {
		return nil, err
	}

	return pesan, nil
}

func (r *repository) FindBySlideShowIDEndpesanPrimary() (Surat, error) {
	var pesan Surat
	// err := r.db.Debug().Model(&Surat{}).Order("created_at desc").Limit(page).Find(&pesan).Error
	err := r.db.Where("utama = 1").Find(&pesan).Error
	if err != nil {
		return pesan, err
	}

	return pesan, nil
}

func (r *repository) FindBySlideShowIDEndpesanNoPrimary() ([]Surat, error) {
	var pesan []Surat
	err := r.db.Model(&Surat{}).Where("utama = 0").Find(&pesan).Error
	if err != nil {
		return pesan, err
	}

	return pesan, nil
}
