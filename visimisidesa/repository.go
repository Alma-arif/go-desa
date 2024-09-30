package visimisidesa

import "gorm.io/gorm"

type Repository interface {
	Save(informasi VisiMisi) (VisiMisi, error)
	FindAllDeletedAt() ([]VisiMisi, error)
	FindAll() ([]VisiMisi, error)
	FindAllLimit() (VisiMisi, error)
	FindByID(id uint) (VisiMisi, error)
	FindByIDDeletedAt(id uint) (VisiMisi, error)
	Update(informasi VisiMisi) (VisiMisi, error)
	UpdateDeletedAt(id uint) (VisiMisi, error)
	DeletedSoft(id uint) (VisiMisi, error)
	Deleted(id uint) (VisiMisi, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(informasi VisiMisi) (VisiMisi, error) {
	err := r.db.Create(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindAllDeletedAt() ([]VisiMisi, error) {
	var informasi []VisiMisi
	err := r.db.Unscoped().Where("deleted_at > 0").Find(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindAll() ([]VisiMisi, error) {
	var informasi []VisiMisi
	err := r.db.Find(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindAllLimit() (VisiMisi, error) {
	var informasi VisiMisi
	err := r.db.Find(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) FindByID(id uint) (VisiMisi, error) {
	var informasi VisiMisi

	err := r.db.Where("id = ?", id).Find(&informasi).Error
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (r *repository) FindByIDDeletedAt(id uint) (VisiMisi, error) {
	var informasi VisiMisi

	err := r.db.Unscoped().Where("id = ?", id).Find(&informasi).Error
	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (r *repository) Update(informasi VisiMisi) (VisiMisi, error) {
	err := r.db.Save(&informasi).Error

	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) UpdateDeletedAt(id uint) (VisiMisi, error) {
	var informasi VisiMisi

	err := r.db.Unscoped().Model(&informasi).Where("id", id).Update("deleted_at", nil).Error

	if err != nil {
		return informasi, err
	}

	return informasi, nil
}

func (r *repository) DeletedSoft(id uint) (VisiMisi, error) {
	var informasi VisiMisi
	err := r.db.Where("id = ?", id).Delete(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

func (r *repository) Deleted(id uint) (VisiMisi, error) {
	var informasi VisiMisi
	err := r.db.Unscoped().Where("id = ?", id).Delete(&informasi).Error
	if err != nil {
		return informasi, err
	}
	return informasi, nil
}

// func (r *repository) FindAllVisiMisiWeb(perPage, page int) ([]VisiMisi, int64, error) {

// 	var informasi []VisiMisi
// 	var count int64

// 	err := r.db.Debug().Model(&VisiMisi{}).Count(&count).Error
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	offset := (page - 1) * perPage

// 	err = r.db.Debug().Model(&VisiMisi{}).Where("status = 1").Order("created_at desc").Limit(perPage).Offset(offset).Find(&informasi).Error
// 	if err != nil {
// 		return nil, 0, err

// 	}

// 	return informasi, count, nil
// }

// func (r *repository) FindAllVisiMisiLimitWeb(page int) ([]VisiMisi, error) {

// 	var informasi []VisiMisi

// 	err := r.db.Debug().Model(&VisiMisi{}).Where("status = 1").Order("created_at desc").Limit(page).Find(&informasi).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return informasi, nil
// }
