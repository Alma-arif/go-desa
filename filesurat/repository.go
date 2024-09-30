package filesurat

import "gorm.io/gorm"

type Repository interface {
	Save(pesan FileSurat) (FileSurat, error)
	FindAll() ([]FileSurat, error)
	FindByID(id uint) (FileSurat, error)
	FindByIDKodeFix(kode string) (FileSurat, error)
	FindByIDEndKodeSuart(id uint, kodeSuart int) (FileSurat, error)
	Update(pesan FileSurat) (FileSurat, error)
	UpdateDeletedAt(id uint) (FileSurat, error)
	DeletedSoft(id uint) (FileSurat, error)
	Deleted(id uint) (FileSurat, error)
	FindAllDeletedAt() ([]FileSurat, error)
	FindByIDDeletedAt(id uint) (FileSurat, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(pesan FileSurat) (FileSurat, error) {
	err := r.db.Create(&pesan).Error
	if err != nil {
		return pesan, err
	}

	return pesan, nil
}

func (r *repository) FindAll() ([]FileSurat, error) {
	var pesan []FileSurat
	err := r.db.Find(&pesan).Error
	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) FindByID(id uint) (FileSurat, error) {
	var pesan FileSurat

	err := r.db.Where("id = ?", id).Find(&pesan).Error
	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) FindByIDKodeFix(kode string) (FileSurat, error) {
	var pesan FileSurat

	err := r.db.Where("kode_surat_fix = ?", kode).Find(&pesan).Error
	if err != nil {
		return pesan, err
	}

	return pesan, nil
}

func (r *repository) FindByIDEndKodeSuart(id uint, kodeSuart int) (FileSurat, error) {
	var pesan FileSurat

	err := r.db.Model(&FileSurat{}).Where("kode_surat = ?", kodeSuart).Order("id DESC").Limit(1).First(&pesan).Error
	if err != nil {
		return pesan, err
	}

	return pesan, nil
}

func (r *repository) Update(pesan FileSurat) (FileSurat, error) {
	err := r.db.Save(&pesan).Error

	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) UpdateDeletedAt(id uint) (FileSurat, error) {
	var pesan FileSurat

	err := r.db.Unscoped().Model(&pesan).Where("id", id).Update("deleted_at", nil).Error
	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) DeletedSoft(id uint) (FileSurat, error) {
	var pesan FileSurat
	err := r.db.Where("id = ?", id).Delete(&pesan).Error
	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) Deleted(id uint) (FileSurat, error) {
	var pesan FileSurat
	err := r.db.Unscoped().Where("id = ?", id).Delete(&pesan).Error
	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) FindAllDeletedAt() ([]FileSurat, error) {
	var pesan []FileSurat
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&pesan).Error
	if err != nil {
		return pesan, err
	}
	return pesan, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (FileSurat, error) {
	var pesan FileSurat

	err := r.db.Unscoped().Where("id = ?", id).Find(&pesan).Error
	if err != nil {
		return pesan, err
	}

	return pesan, nil
}
