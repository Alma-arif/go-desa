package slideshow

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"errors"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateImageSlideShow(input ImageSlideShowInput, imageFile string) (ImageSlideShow, error)
	GetAllImageSlideShow() ([]ImageSlideShowView, error)
	GetAllImageSlideShowDeletedAt() ([]ImageSlideShowView, error)
	GetImageSlideShowByID(id uint) (ImageSlideShowView, error)
	GetImageSlideShowByIDDeleted(id uint) (ImageSlideShow, error)
	UpdateImageSlideShow(input ImageSlideShowUpdate, imageFile string) (ImageSlideShow, error)
	DeletedImageSlideShowSoft(id uint) (ImageSlideShow, error)
	DeletedImageSlideShow(id uint) (ImageSlideShow, error)
	RestoreImageSlideShow(id uint) (ImageSlideShow, error)
	GetAllImageSlideShowWeb() (ImageFormatter, error)
}

type service struct {
	repository     Repository
	Validate       *validator.Validate
	userRepository user.Repository
}

func NewService(repository Repository, Validate *validator.Validate, userRepository user.Repository) *service {
	return &service{repository, Validate, userRepository}
}

func (s *service) CreateImageSlideShow(input ImageSlideShowInput, imageFile string) (ImageSlideShow, error) {
	var image ImageSlideShow

	err := s.Validate.Struct(input)
	if err != nil {

		return image, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	image.Judul = input.Judul
	image.Keterangan = input.Keterangan
	image.SlideShowImages = imageFile
	image.Utama = input.Utama
	image.Link = input.Link

	_, err = s.repository.Save(image)

	if err != nil {
		return image, err
	}

	return image, nil
}

func (s *service) GetAllImageSlideShow() ([]ImageSlideShowView, error) {
	var image []ImageSlideShowView

	resultImage, err := s.repository.FindAll()
	if err != nil {
		return image, err
	}

	for i, rowImage := range resultImage {
		var imageView ImageSlideShowView
		imageView.ID = rowImage.ID
		imageView.Index = i + 1
		imageView.Judul = rowImage.Judul
		imageView.Keterangan = rowImage.Keterangan
		imageView.SlideShowImages = rowImage.SlideShowImages
		imageView.Link = rowImage.Link
		imageView.Utama = rowImage.Utama
		imageView.CreatedAt = rowImage.CreatedAt
		image = append(image, imageView)
	}

	return image, nil
}

func (s *service) GetAllImageSlideShowDeletedAt() ([]ImageSlideShowView, error) {
	var image []ImageSlideShowView

	resultImage, err := s.repository.FindAllDeletedAt()
	if err != nil {
		return image, err
	}

	for i, rowImage := range resultImage {
		var imageView ImageSlideShowView
		imageView.ID = rowImage.ID
		imageView.Index = i + 1
		imageView.Judul = rowImage.Judul
		imageView.Keterangan = rowImage.Keterangan
		imageView.SlideShowImages = rowImage.SlideShowImages
		imageView.Link = rowImage.Link
		imageView.Utama = rowImage.Utama
		imageView.CreatedAt = rowImage.CreatedAt
		image = append(image, imageView)
	}

	return image, nil
}

func (s *service) GetImageSlideShowByID(id uint) (ImageSlideShowView, error) {
	var image ImageSlideShowView

	rowImage, err := s.repository.FindByID(id)
	if err != nil {
		return image, err
	}

	image.ID = rowImage.ID
	// image.Index = i + 1
	image.Judul = rowImage.Judul
	image.Keterangan = rowImage.Keterangan
	image.SlideShowImages = rowImage.SlideShowImages
	image.Link = rowImage.Link
	image.Utama = rowImage.Utama
	image.CreatedAt = rowImage.CreatedAt

	if image.ID == 0 {
		return image, errors.New("No Image found on with that ID")
	}

	return image, nil
}

func (s *service) GetImageSlideShowByIDDeleted(id uint) (ImageSlideShow, error) {
	var image ImageSlideShow

	rowImage, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return image, err
	}

	return rowImage, nil
}

func (s *service) UpdateImageSlideShow(input ImageSlideShowUpdate, imageFile string) (ImageSlideShow, error) {
	var image ImageSlideShow

	err := s.Validate.Struct(input)
	if err != nil {
		return image, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	rowimage, err := s.repository.FindByID(input.ID)
	if err != nil {
		return image, err
	}

	image.ID = rowimage.ID
	image.Judul = input.Judul
	image.Keterangan = input.Keterangan

	if imageFile == "" {
		imageFile = rowimage.SlideShowImages
	} else {
		pathRemoveFile := fmt.Sprintf("derektori/image/%s", rowimage.SlideShowImages)
		err = os.Remove(pathRemoveFile)
		if err != nil {
			return image, err
		}
	}

	image.SlideShowImages = imageFile
	image.Link = input.Link
	image.Utama = input.Utama
	image.CreatedAt = rowimage.CreatedAt

	imageUpdate, err := s.repository.Update(image)
	if err != nil {
		return image, err
	}
	return imageUpdate, nil
}

func (s *service) DeletedImageSlideShowSoft(id uint) (ImageSlideShow, error) {
	var image ImageSlideShow
	imageRow, err := s.repository.FindByID(id)
	if err != nil {
		return image, err
	}

	var newFile string
	if imageRow.SlideShowImages != "" {
		path := fmt.Sprintf("derektori/image/%s", imageRow.SlideShowImages)
		_, err = os.Stat(path)
		if err != nil {
			return image, err
		}

		fileName, err := helper.GetFileNameDekrip(imageRow.SlideShowImages)
		if err != nil {
			return image, err
		}

		newFile, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			return image, err
		}

		NewPath := fmt.Sprintf("derektori/image/%s", newFile)
		err = os.Rename(path, NewPath)
		if err != nil {
			return image, err
		}

	}

	image.ID = imageRow.ID
	image.Judul = imageRow.Judul
	image.Keterangan = imageRow.Keterangan
	image.SlideShowImages = newFile
	image.Link = imageRow.Link

	image.Utama = 0
	image.CreatedAt = imageRow.CreatedAt

	update, err := s.repository.Update(image)
	if err != nil {
		return update, err
	}

	deletedimage, err := s.repository.DeletedSoft(update.ID)
	if err != nil {
		return image, err
	}

	return deletedimage, nil
}

func (s *service) DeletedImageSlideShow(id uint) (ImageSlideShow, error) {
	var image ImageSlideShow

	image, err := s.repository.Deleted(id)
	if err != nil {
		return image, err
	}
	if image.SlideShowImages != "" {
		pathRemoveFile := fmt.Sprintf("derektori/image/%s", image.SlideShowImages)
		err = os.Remove(pathRemoveFile)
		if err != nil {
			return image, err
		}
	}
	return image, nil
}

func (s *service) RestoreImageSlideShow(id uint) (ImageSlideShow, error) {
	var image ImageSlideShow

	imageDeleted, err := s.repository.FindByIDDeletedAt(id)
	if err != nil {
		return image, err
	}

	updateImage, err := s.repository.UpdateDeletedAt(imageDeleted.ID)
	if err != nil {
		return image, err
	}

	return updateImage, nil
}

func (s *service) GetAllImageSlideShowWeb() (ImageFormatter, error) {
	var image ImageFormatter

	imagePromary, err := s.repository.FindBySlideShowIDEndImagePrimary()
	if err != nil {
		return image, err
	}

	imageNoPromary, err := s.repository.FindBySlideShowIDEndImageNoPrimary()
	if err != nil {
		return image, err
	}

	imageResult := ImageBeritaFormatter(imagePromary, imageNoPromary)

	return imageResult, nil
}
