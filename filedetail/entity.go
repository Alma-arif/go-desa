package filedetail

import "time"

type FileDetail struct {
	ID           uint
	Index        int
	ArsipID      uint
	ArsipName    string
	FileName     string
	Deskripsi    string
	FileSize     float64
	FileLocation string
	FileStatus   int
	KategoryName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
