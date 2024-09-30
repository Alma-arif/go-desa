package helper

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func IsAllowedFileType(file *multipart.FileHeader) (bool, string) {
	// Mendapatkan ekstensi file
	// allowedExtensions := []string{".jpg", ".jpeg", ".png", ".PNG", ".pdf", ".PDF", ".doc", ".docx", ".xls", "xlsx", ".pptx", ".ppt"}
	allowedExtensions := strings.Split(".PNG|.png|.jpeg|.jpg|.pdf|.docx|.doc|.xls|.xlsx|.ppt|.pptx", "|")

	// allowedExtensions := strings.Split(os.Getenv("APP_FILE_TYPE"), "|")
	ext := filepath.Ext(file.Filename)

	// Memeriksa apakah ekstensi file diizinkan
	for _, allowedExt := range allowedExtensions {
		if strings.EqualFold(allowedExt, ext) {
			return true, allowedExt
		}
	}

	return false, ""
}

func IsAllowedFileTypeImage(file *multipart.FileHeader) (bool, string) {

	allowedExtensions := strings.Split(os.Getenv("APP_FILE_TYPE_IMAGE_PROFILE"), "|")
	ext := filepath.Ext(file.Filename)

	for _, allowedExt := range allowedExtensions {
		if strings.EqualFold(allowedExt, ext) {
			return true, allowedExt
		}
	}

	return false, ""
}
