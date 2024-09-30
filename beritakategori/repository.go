package beritakategori

import "gorm.io/gorm"

type Repository interface {
	Save(kategori KategoriBerita) (KategoriBerita, error)
	FindAll() ([]KategoriBerita, error)
	FindAllDeletedAt() ([]KategoriBerita, error)
	FindByID(id uint) (KategoriBerita, error)
	FindByIDDeletedAt(id uint) (KategoriBerita, error)
	Update(kategori KategoriBerita) (KategoriBerita, error)
	UpdateDeletedAt(id uint) (KategoriBerita, error)
	DeletedSoft(id uint) (KategoriBerita, error)
	Deleted(id uint) (KategoriBerita, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(kategori KategoriBerita) (KategoriBerita, error) {
	err := r.db.Create(&kategori).Error
	if err != nil {
		return kategori, err
	}
	return kategori, nil
}

func (r *repository) FindAllDeletedAt() ([]KategoriBerita, error) {
	var kategori []KategoriBerita
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&kategori).Error
	if err != nil {
		return kategori, err
	}
	return kategori, nil
}

func (r *repository) FindAll() ([]KategoriBerita, error) {
	var kategori []KategoriBerita
	err := r.db.Find(&kategori).Error
	if err != nil {
		return kategori, err
	}
	return kategori, nil
}

func (r *repository) FindByID(id uint) (KategoriBerita, error) {
	var kategori KategoriBerita

	err := r.db.Where("id = ?", id).Find(&kategori).Error
	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (KategoriBerita, error) {
	var kategori KategoriBerita

	err := r.db.Unscoped().Where("id = ?", id).Find(&kategori).Error
	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (r *repository) Update(kategori KategoriBerita) (KategoriBerita, error) {
	err := r.db.Save(&kategori).Error

	if err != nil {
		return kategori, err
	}
	return kategori, nil
}

func (r *repository) UpdateDeletedAt(id uint) (KategoriBerita, error) {
	var kategori KategoriBerita

	err := r.db.Unscoped().Model(&kategori).Where("id", id).Update("deleted_at", nil).Error

	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (r *repository) DeletedSoft(id uint) (KategoriBerita, error) {
	var kategori KategoriBerita
	err := r.db.Where("id = ?", id).Delete(&kategori).Error
	if err != nil {
		return kategori, err
	}

	return kategori, nil

}

func (r *repository) Deleted(id uint) (KategoriBerita, error) {
	var kategori KategoriBerita
	err := r.db.Unscoped().Where("id = ?", id).Delete(&kategori).Error
	if err != nil {
		return kategori, err
	}

	return kategori, nil

}
