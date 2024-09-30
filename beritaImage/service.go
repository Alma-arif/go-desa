package beritaimage

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateBeritaImage(input ImageBeritaInput, imageName string) (ImageBerita, error)
	GetAllBeritaImage() ([]ImageBeritaView, error)
	GetBeritaImageByBeritaID(id uint) (ImageFormatter, error)
	DeleteImageBerita(id uint) (ImageBerita, error)
}

type service struct {
	repository Repository
	Validate   *validator.Validate
}

func NewService(repository Repository, Validate *validator.Validate) *service {
	return &service{repository, Validate}
}

func (s *service) CreateBeritaImage(input ImageBeritaInput, imageName string) (ImageBerita, error) {
	var image ImageBerita

	err := s.Validate.Struct(input)
	if err != nil {
		return image, errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!")
	}

	image.NamaImageFile = imageName
	image.IdBerita = input.IdBerita
	image.ImageUtama = input.ImageUtama
	// image.Status = input.Status

	if input.ImageUtama == 1 {
		rowImagePrimary, err := s.repository.FindByBeritaIDEndImagePrimary(input.IdBerita)
		if err != nil {
			return image, err
		}

		if rowImagePrimary.ID >= uint(1) {
			var imageberita ImageBerita

			imageberita.ID = rowImagePrimary.ID
			imageberita.NamaImageFile = rowImagePrimary.NamaImageFile
			imageberita.IdBerita = rowImagePrimary.IdBerita
			imageberita.ImageUtama = 0
			imageberita.CreatedAt = rowImagePrimary.CreatedAt
			_, err := s.repository.Update(imageberita)
			if err != nil {
				return image, err
			}
		}

	}

	newImage, err := s.repository.Save(image)
	if err != nil {
		return newImage, err
	}

	return newImage, nil
}

func (s *service) GetAllBeritaImage() ([]ImageBeritaView, error) {
	var images []ImageBeritaView

	resultImages, err := s.repository.FindAll()
	if err != nil {
		return images, err
	}

	for i, image := range resultImages {
		var ImageView ImageBeritaView
		ImageView.ID = image.ID
		ImageView.Index = i + 1
		ImageView.NamaImageFile = image.NamaImageFile
		ImageView.IdBerita = image.IdBerita
		ImageView.ImageUtama = image.ImageUtama
		ImageView.CreatedAt = image.CreatedAt
		ImageView.UpdatedAt = image.UpdatedAt
		images = append(images, ImageView)
	}

	return images, nil
}

func (s *service) GetBeritaImageByBeritaID(id uint) (ImageFormatter, error) {
	var images ImageFormatter

	rowImagePrimary, err := s.repository.FindByBeritaIDEndImagePrimary(id)
	if err != nil {
		return images, err
	}

	resultImagePrimary, err := s.repository.FindByBeritaIDEndImageNoPrimary(id)
	if err != nil {
		return images, err
	}

	images = ImageBeritaFormatter(rowImagePrimary, resultImagePrimary)

	return images, nil
}

func (s *service) DeleteImageBerita(id uint) (ImageBerita, error) {
	var imageBerita ImageBerita

	imageBerita, err := s.repository.Deleted(id)
	if err != nil {
		return imageBerita, err
	}

	path := fmt.Sprintf("derektori/images_berita/%s", imageBerita.NamaImageFile)

	err = os.Remove(path)
	if err != nil {
		return imageBerita, err
	}

	return imageBerita, nil
}
