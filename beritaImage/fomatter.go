package beritaimage

import (
	"time"
)

type ImageFormatter struct {
	ID               uint
	Index            int
	NamaImagePrimary string
	IdBerita         uint
	ImageUtama       int
	CreatedAt        time.Time
	ImageNoPrimary   []ImageAllNoPrimary
}

type ImageAllNoPrimary struct {
	ID                 uint
	Index              int
	NamaImageNoPrimary string
	IdBerita           uint
	CreatedAt          time.Time
}

func ImageBeritaFormatter(imagePrimary ImageBerita, imageNoPrimary []ImageBerita) ImageFormatter {

	var imageFormatter ImageFormatter

	imageFormatter.ID = imagePrimary.ID
	imageFormatter.Index = 1
	imageFormatter.NamaImagePrimary = imagePrimary.NamaImageFile
	imageFormatter.ImageUtama = imagePrimary.ImageUtama
	imageFormatter.CreatedAt = imagePrimary.CreatedAt

	images := []ImageAllNoPrimary{}
	for i, image := range imageNoPrimary {
		var imageAllNoPrimary ImageAllNoPrimary

		imageAllNoPrimary.ID = image.ID
		imageAllNoPrimary.Index = i + 2
		imageAllNoPrimary.NamaImageNoPrimary = image.NamaImageFile
		imageAllNoPrimary.IdBerita = image.IdBerita
		imageAllNoPrimary.CreatedAt = image.CreatedAt

		images = append(images, imageAllNoPrimary)
	}

	imageFormatter.ImageNoPrimary = images

	return imageFormatter
}
