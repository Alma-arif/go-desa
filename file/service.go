package file

import (
	"app-desa-kepuk/helper"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	GetFileAll() ([]File, error)
	GetFileByID(id uint) (File, error)
	RestoreFile(id uint) (File, error)
	GetFileByArsipID(id uint) ([]File, error)
	GetFileAllByArsipIDNull() ([]File, error)
	CreateFile(input CreateFileInput, fileName string) (File, error)
	UpdateFile(input UpdateFileInput) (File, error)
	DeleteFileSoft(id uint) (File, error)
	DeleteFile(id uint) (File, error)
	Enkripsi(id uint) (File, error)
	Dekripsi(id uint) (File, error)
	UpdateFileArispID(id uint, idarsip uint) (File, error)

	EnkripsiRC(id uint) (File, error)
	DekripsiRC(id uint) (File, error)
}

type service struct {
	repository Repository
	Validate   *validator.Validate
}

func NewService(repository Repository, Validate *validator.Validate) *service {
	return &service{repository, Validate}
}

func (s *service) RestoreFile(id uint) (File, error) {
	var file File

	file, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return file, err
	}

	update, err := s.repository.UpdateDeletetAt(file.ID)
	if err != nil {
		return update, err
	}

	return update, nil
}

func (s *service) GetFileAll() ([]File, error) {
	file, err := s.repository.FindAll()
	if err != nil {
		return file, err
	}

	return file, nil
}

func (s *service) GetFileByID(id uint) (File, error) {

	file, err := s.repository.FindByID(id)
	if err != nil {
		return file, err
	}

	return file, nil
}

func (s *service) GetFileByArsipID(id uint) ([]File, error) {
	file, err := s.repository.FindFileByArsipID(id)
	if err != nil {
		return file, err
	}

	return file, nil
}

func (s *service) GetFileAllByArsipIDNull() ([]File, error) {
	file, err := s.repository.FindFileByArsipIDNull()
	if err != nil {
		return file, err
	}

	return file, nil
}

func (s *service) CreateFile(input CreateFileInput, fileName string) (File, error) {
	var file File

	err := s.Validate.Struct(input)
	if err != nil {
		return file, errors.New("isi form dengan benar")
	}
	if input.NamaFile == "" {
		file.NamaFile = fileName
	} else {
		file.NamaFile = input.NamaFile
	}

	file.Deskripsi = input.DeskripsiFile
	file.ArsipID = input.ArsipID
	file.FileSize = input.FileSize
	file.FileLocation = input.FileLocation
	file.Status = input.Status

	newFile, err := s.repository.Save(file)
	if err != nil {
		return newFile, err
	}

	return newFile, nil
}

func (s *service) UpdateFile(input UpdateFileInput) (File, error) {
	var file File

	err := s.Validate.Struct(input)
	if err != nil {
		return file, errors.New("isi form dengan benar")
	}

	fileRow, err := s.repository.FindByID(input.ID)
	if err != nil {
		return file, err
	}

	file.ID = fileRow.ID

	if input.NamaFile == fileRow.NamaFile {
		file.NamaFile = fileRow.NamaFile
	} else {
		path := fmt.Sprintf("derektori/file/%s", fileRow.FileLocation)
		_, err = os.Stat(path)
		if err != nil {
			return file, err
		}

		extension := filepath.Ext(fileRow.NamaFile)
		fileString := fmt.Sprintf("%s%s", input.NamaFile, extension)

		filenamaa, err := helper.GetFileNameEnkrip(fileString)
		if err != nil {
			return file, err
		}

		pathEnkripsi := fmt.Sprintf("derektori/file/%s", filenamaa)
		err = os.Rename(path, pathEnkripsi)
		if err != nil {
			return file, err
		}

		file.NamaFile = fileString

	}
	file.Deskripsi = input.DeskripsiFile
	file.ArsipID = input.ArsipID
	file.FileSize = fileRow.FileSize
	file.FileLocation = fileRow.FileLocation
	file.CreatedAt = fileRow.CreatedAt

	update, err := s.repository.Update(file)
	if err != nil {
		return update, err
	}

	return update, nil
}

func (s *service) DeleteFileSoft(id uint) (File, error) {
	var file File

	resultFile, err := s.repository.FindByID(id)
	if err != nil {
		return file, err
	}

	path := fmt.Sprintf("derektori/file/%s", resultFile.FileLocation)
	_, err = os.Stat(path)
	if err != nil {
		return file, err
	}

	fileName, err := helper.GetFileNameDekrip(resultFile.FileLocation)
	if err != nil {
		return file, err
	}

	newFile, err := helper.GetFileNameEnkrip(fileName)
	if err != nil {
		return file, err
	}

	NewPath := fmt.Sprintf("derektori/file/%s", newFile)
	err = os.Rename(path, NewPath)
	if err != nil {
		return file, err
	}

	file.ID = resultFile.ID
	file.ArsipID = resultFile.ArsipID
	file.FileLocation = newFile
	file.NamaFile = fileName
	file.Deskripsi = resultFile.Deskripsi
	file.FileSize = resultFile.FileSize
	file.Status = resultFile.Status
	file.CreatedAt = resultFile.CreatedAt

	update, err := s.repository.Update(file)
	if err != nil {
		return update, err
	}

	file, err = s.repository.DeletedSoft(id)
	if err != nil {
		return file, err
	}

	return file, nil
}

func (s *service) DeleteFile(id uint) (File, error) {
	var file File

	resultFile, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return file, err
	}

	path := fmt.Sprintf("derektori/file/%s", resultFile.FileLocation)

	err = os.Remove(path)
	if err != nil {
		return file, err
	}

	file, err = s.repository.Deleted(resultFile.ID)
	if err != nil {
		return file, err
	}

	return file, nil
}

func (s *service) Enkripsi(id uint) (File, error) {
	var file File

	start := time.Now()

	resultFile, err := s.repository.FindByID(id)
	if err != nil {
		return file, err
	}

	path := fmt.Sprintf("derektori/file/%s", resultFile.FileLocation)
	// filetow, err := os.Stat(path)
	// if err != nil {
	// 	return file, err
	// }
	// fmt.Println("file size sebelum di enkripsi (mb):", float64(filetow.Size())/1024/1024)

	fileData, err := os.ReadFile(path)
	if err != nil {
		return file, err
	}

	enkripsiData := helper.Rc4Data([]byte(os.Getenv("APP_FILE_SECRET_KEY")), fileData)
	// if err != nil {
	// 	return file, err
	// }

	filenama, err := helper.GetFileNameDekrip(resultFile.FileLocation)
	if err != nil {
		return file, err
	}

	filenamaa, err := helper.GetFileNameEnkrip(filenama)
	if err != nil {
		return file, err
	}

	err = os.WriteFile(path, enkripsiData, 0644)
	if err != nil {
		return file, err
	}

	pathEnkripsi := fmt.Sprintf("derektori/file/%s", filenamaa)
	err = os.Rename(path, pathEnkripsi)
	if err != nil {
		return file, err
	}

	fileDataEnkrisi, err := os.Open(pathEnkripsi)
	defer fileDataEnkrisi.Close()
	if err != nil {
		return file, err
	}

	fileInfo, err := fileDataEnkrisi.Stat()
	if err != nil {
		return file, err
	}

	file.ID = resultFile.ID
	file.ArsipID = resultFile.ArsipID
	file.NamaFile = resultFile.NamaFile
	file.Status = 1 // 1 status enkrispi
	file.FileSize = float64(fileInfo.Size()) / 1024 / 1024
	file.FileLocation = filenamaa
	file.Deskripsi = resultFile.Deskripsi
	file.CreatedAt = resultFile.CreatedAt

	_, err = s.repository.Update(file)
	if err != nil {
		return file, err
	}

	duration := time.Since(start)

	fmt.Println("file size sesudah di enkripsi (mb):", float64(fileInfo.Size())/1024/1024)
	fmt.Println("time dekrissi in detik:", duration.Seconds())

	return file, nil
}

func (s *service) EnkripsiRC(id uint) (File, error) {
	var file File

	resultFile, err := s.repository.FindByID(id)
	if err != nil {
		return file, err
	}

	path := fmt.Sprintf("../derektori/file/%s", resultFile.FileLocation)

	fileData, err := os.ReadFile(path)
	if err != nil {
		return file, err
	}

	enkripsiData, _ := helper.Rc4Encrypt(fileData, []byte("SIDESAdatadesaKepukBangsriJeparaFILEterEnkripsiRc4"))

	filenama, err := helper.GetFileNameDekrip(resultFile.FileLocation)
	if err != nil {

		return file, err
	}

	filenamaa, err := helper.GetFileNameEnkrip(filenama)
	if err != nil {

		return file, err
	}

	err = os.WriteFile(path, enkripsiData, 0644)
	if err != nil {

		return file, err
	}

	pathEnkripsi := fmt.Sprintf("../derektori/file/%s", filenamaa)
	err = os.Rename(path, pathEnkripsi)
	if err != nil {

		return file, err
	}

	fileDataEnkrisi, err := os.Open(pathEnkripsi)
	defer fileDataEnkrisi.Close()
	if err != nil {
		return file, err
	}

	fileInfo, err := fileDataEnkrisi.Stat()
	if err != nil {

		return file, err
	}

	file.ID = resultFile.ID
	file.ArsipID = resultFile.ArsipID
	file.NamaFile = resultFile.NamaFile
	file.Status = 1 // 1 status enkrispi
	file.FileSize = float64(fileInfo.Size()) / 1024 / 1024
	file.FileLocation = filenamaa
	file.Deskripsi = resultFile.Deskripsi
	file.CreatedAt = resultFile.CreatedAt

	_, err = s.repository.Update(file)
	if err != nil {

		return file, err
	}

	return file, nil
}

func (s *service) Dekripsi(id uint) (File, error) {
	var file File

	start := time.Now()

	resultFile, err := s.repository.FindByID(id)
	if err != nil {
		return file, err
	}

	path := fmt.Sprintf("derektori/file/%s", resultFile.FileLocation)
	fileResult, err := os.ReadFile(path)
	if err != nil {
		return file, err
	}

	filetow, err := os.Stat(path)
	// defer filetow
	if err != nil {
		return file, err
	}

	fmt.Println("file size sebelum di dekripsi (mb):", float64(filetow.Size())/1024/1024)

	dekripsiData := helper.Rc4Data([]byte(os.Getenv("APP_FILE_SECRET_KEY")), fileResult)
	// if err != nil {
	// 	return file, err
	// }

	fileModif, err := os.OpenFile(path, os.O_RDWR, 0644)
	defer fileModif.Close()
	if err != nil {
		return file, err
	}

	_, err = fileModif.WriteAt(dekripsiData, 0)
	if err != nil {
		return file, err
	}

	err = fileModif.Sync()
	if err != nil {
		return file, err
	}

	fileInfo, err := fileModif.Stat()
	if err != nil {
		return file, err
	}

	file.ID = resultFile.ID
	file.ArsipID = resultFile.ArsipID
	file.NamaFile = resultFile.NamaFile
	file.Status = 0 // 0 status dekropsi
	file.FileSize = float64(fileInfo.Size()) / 1024 / 1024
	file.FileLocation = resultFile.FileLocation
	file.Deskripsi = resultFile.Deskripsi
	file.CreatedAt = resultFile.CreatedAt

	_, err = s.repository.Update(file)
	if err != nil {
		// os.Remove(pathDekripsi)
		return file, err
	}

	// os.Remove(path)

	duration := time.Since(start)
	fmt.Println("file size sesudah di enkripsi (mb):", float64(fileInfo.Size())/1024/1024)
	fmt.Println("time dekrissi in detik:", duration.Seconds())
	fmt.Println("time dekrissi in menit:", duration.Minutes())
	fmt.Println("time dekrissi in jam:", duration.Hours())

	return file, nil
}

func (s *service) DekripsiRC(id uint) (File, error) {
	var file File

	resultFile, err := s.repository.FindByID(id)
	if err != nil {
		return file, err
	}

	path := fmt.Sprintf("../derektori/file/%s", resultFile.FileLocation)
	fileResult, err := os.ReadFile(path)
	if err != nil {
		return file, err
	}

	// filetow, err := os.Stat(path)
	// // defer filetow
	// if err != nil {
	// 	return file, err
	// }

	// fmt.Println("file size sebelum di dekripsi (mb):", float64(filetow.Size())/1024/1024)

	dekripsiData, _ := helper.Rc4Decrypt(fileResult, []byte("SIDESAdatadesaKepukBangsriJeparaFILEterEnkripsiRc4"))
	// if err != nil {
	// 	return file, err
	// }

	fileModif, err := os.OpenFile(path, os.O_RDWR, 0644)
	defer fileModif.Close()
	if err != nil {
		return file, err
	}

	_, err = fileModif.WriteAt(dekripsiData, 0)
	if err != nil {
		return file, err
	}

	err = fileModif.Sync()
	if err != nil {
		return file, err
	}

	fileInfo, err := fileModif.Stat()
	if err != nil {
		return file, err
	}

	file.ID = resultFile.ID
	file.ArsipID = resultFile.ArsipID
	file.NamaFile = resultFile.NamaFile
	file.Status = 0 // 0 status dekropsi
	file.FileSize = float64(fileInfo.Size()) / 1024 / 1024
	file.FileLocation = resultFile.FileLocation
	file.Deskripsi = resultFile.Deskripsi
	file.CreatedAt = resultFile.CreatedAt

	_, err = s.repository.Update(file)
	if err != nil {
		// os.Remove(pathDekripsi)
		return file, err
	}

	return file, nil
}

func (s *service) UpdateFileArispID(id uint, idarsip uint) (File, error) {
	var file File

	resultFile, err := s.repository.FindByID(id)
	if err != nil {
		return file, err
	}

	file.ID = resultFile.ID
	file.ArsipID = idarsip
	file.NamaFile = resultFile.NamaFile
	file.Status = resultFile.Status
	file.FileSize = resultFile.FileSize
	file.FileLocation = resultFile.FileLocation
	file.Deskripsi = resultFile.Deskripsi
	file.CreatedAt = resultFile.CreatedAt

	_, err = s.repository.Update(file)
	if err != nil {
		return file, err
	}

	return file, nil
}
