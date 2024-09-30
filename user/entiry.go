package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	Nama         string
	Email        string
	Password     string
	NoHp         string
	TanggalLahir time.Time
	Role         string
	ProfileFile  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type UserView struct {
	ID           uint
	Index        int
	Nama         string
	Email        string
	Password     string
	NoHp         string
	TanggalLahir time.Time
	Role         string
	ProfileFile  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

// func (user *User) BeforeCretateUser(tx *gorm.DB) (err error) {
// 	user.ID = uuid.New
// 	return
// }
