package pengumuman

import "gorm.io/gorm"

type Repository interface {
	// dasboard
	Save(informasi Pengumuman) (Pengumuman, error)
	FindAllDeletedAt() ([]Pengumuman, error)
	FindAll() ([]Pengumuman, error)
	FindByID(id uint) (Pengumuman, error)
	FindByIDDeletedAt(id uint) (Pengumuman, error)
	Update(informasi Pengumuman) (Pengumuman, error)
	UpdateDeletedAt(id uint) (Pengumuman, error)
	DeletedSoft(id uint) (Pengumuman, error)
	Deleted(id uint) (Pengumuman, error)
	// web
	FindAllPengumumanWeb(perPage, page int) ([]Pengumuman, int64, error)
	FindAllPengumumanLimitWeb(page int) ([]Pengumuman, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(informasi Pengumuman) (Pengumuman, error) {
	err := r.db.Create(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindAllDeletedAt() ([]Pengumuman, error) {
	var informasi []Pengumuman
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindAll() ([]Pengumuman, error) {
	var informasi []Pengumuman
	err := r.db.Find(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindByID(id uint) (Pengumuman, error) {
	var informasi Pengumuman

	err := r.db.Where("id = ?", id).Find(&informasi).Error
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (Pengumuman, error) {
	var informasi Pengumuman

	err := r.db.Unscoped().Where("id = ?", id).Find(&informasi).Error
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (r *repository) Update(informasi Pengumuman) (Pengumuman, error) {
	err := r.db.Save(&informasi).Error

	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) UpdateDeletedAt(id uint) (Pengumuman, error) {
	var informasi Pengumuman

	err := r.db.Unscoped().Model(&informasi).Where("id", id).Update("deleted_at", nil).Error

	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (r *repository) DeletedSoft(id uint) (Pengumuman, error) {
	var informasi Pengumuman
	err := r.db.Where("id = ?", id).Delete(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) Deleted(id uint) (Pengumuman, error) {
	var informasi Pengumuman
	err := r.db.Unscoped().Where("id = ?", id).Delete(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindAllPengumumanWeb(perPage, page int) ([]Pengumuman, int64, error) {

	var informasi []Pengumuman
	var count int64

	err := r.db.Debug().Model(&Pengumuman{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage

	err = r.db.Debug().Model(&Pengumuman{}).Where("status = 1").Order("created_at desc").Limit(perPage).Offset(offset).Find(&informasi).Error
	if err != nil {
		return nil, 0, err

	}

	return informasi, count, nil
}

func (r *repository) FindAllPengumumanLimitWeb(page int) ([]Pengumuman, error) {

	var informasi []Pengumuman

	err := r.db.Debug().Model(&Pengumuman{}).Where("status = 1").Order("created_at desc").Limit(page).Find(&informasi).Error
	if err != nil {
		return nil, err
	}

	return informasi, nil
}
