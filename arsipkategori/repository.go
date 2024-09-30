package arsipkategori

import "gorm.io/gorm"

type Repository interface {
	Save(kategori KategoriArsip) (KategoriArsip, error)
	FindAll() ([]KategoriArsip, error)
	FindAllDeletedAt() ([]KategoriArsip, error)
	FindByID(id uint) (KategoriArsip, error)
	FindByIDDeletedAt(id uint) (KategoriArsip, error)
	Update(kategori KategoriArsip) (KategoriArsip, error)
	UpdateDeletedAt(id uint) (KategoriArsip, error)
	DeletedSoft(id uint) (KategoriArsip, error)
	Deleted(id uint) (KategoriArsip, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(kategori KategoriArsip) (KategoriArsip, error) {
	err := r.db.Create(&kategori).Error
	if err != nil {
		return kategori, err
	}
	return kategori, nil
}

func (r *repository) FindAll() ([]KategoriArsip, error) {
	var kategori []KategoriArsip
	err := r.db.Find(&kategori).Error
	if err != nil {
		return kategori, err
	}
	return kategori, nil
}

func (r *repository) FindAllDeletedAt() ([]KategoriArsip, error) {
	var kategori []KategoriArsip
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&kategori).Error
	if err != nil {
		return kategori, err
	}
	return kategori, nil
}

func (r *repository) FindByID(id uint) (KategoriArsip, error) {
	var kategori KategoriArsip

	err := r.db.Where("id = ?", id).Find(&kategori).Error
	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (KategoriArsip, error) {
	var kategori KategoriArsip

	err := r.db.Unscoped().Where("id = ?", id).Find(&kategori).Error
	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (r *repository) Update(kategori KategoriArsip) (KategoriArsip, error) {
	err := r.db.Save(&kategori).Error

	if err != nil {
		return kategori, err
	}
	return kategori, nil
}

func (r *repository) UpdateDeletedAt(id uint) (KategoriArsip, error) {
	var arsipKategori KategoriArsip

	err := r.db.Unscoped().Model(&arsipKategori).Where("id", id).Update("deleted_at", nil).Error

	if err != nil {
		return arsipKategori, err
	}

	return arsipKategori, nil
}

func (r *repository) DeletedSoft(id uint) (KategoriArsip, error) {
	var kategori KategoriArsip
	err := r.db.Where("id = ?", id).Delete(&kategori).Error
	if err != nil {
		return kategori, err
	}

	return kategori, nil

}

func (r *repository) Deleted(id uint) (KategoriArsip, error) {
	var kategori KategoriArsip
	err := r.db.Unscoped().Where("id = ?", id).Delete(&kategori).Error
	if err != nil {
		return kategori, err
	}

	return kategori, nil

}
