package file

type CreateFileInput struct {
	ArsipID       uint   `form:"arsip_id"`
	NamaFile      string `form:"nama"`
	DeskripsiFile string `form:"deskripsi_file"`
	FileSize      float64
	FileLocation  string
	Status        int
}

type CreateFileInputtest struct {
	ArsipID       uint   `form:"arsip_id"`
	NamaFile      string `form:"nama"`
	DeskripsiFile string `form:"deskripsi_file"`
	FileLocation  string
}

type UpdateFileInput struct {
	ID            uint   `form:"file_id" validate:"required"`
	ArsipID       uint   `form:"arsip_id" `
	NamaFile      string `form:"nama"`
	DeskripsiFile string `form:"deskripsi_file"`
	FileSize      float64
	FileLocation  string
	Status        int
}

type UpdateFileArispInput struct {
	ArsipID uint `form:"id" validate:"required"`
	FileID  uint `form:"file" `
}
