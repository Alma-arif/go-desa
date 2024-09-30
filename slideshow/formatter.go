package slideshow

import "time"

type ImageFormatter struct {
	ID              uint
	Judul           string
	Keterangan      string
	SlideShowImages string
	Link            string
	Utama           int
	CreatedAt       time.Time
	ImageNoPrimary  []ImageAllNoPrimary
}

type ImageAllNoPrimary struct {
	ID              uint
	Index           int
	Judul           string
	Keterangan      string
	SlideShowImages string
	Link            string
	Utama           int
	CreatedAt       time.Time
}

func ImageBeritaFormatter(imagePrimary ImageSlideShow, imageNoPrimary []ImageSlideShow) ImageFormatter {

	var imageFormatter ImageFormatter

	imageFormatter.ID = imagePrimary.ID
	imageFormatter.Judul = imagePrimary.Judul
	imageFormatter.Keterangan = imagePrimary.Keterangan
	imageFormatter.SlideShowImages = imagePrimary.SlideShowImages
	imageFormatter.Link = imagePrimary.Link
	imageFormatter.Utama = imagePrimary.Utama
	imageFormatter.CreatedAt = imagePrimary.CreatedAt

	images := []ImageAllNoPrimary{}
	for i, image := range imageNoPrimary {
		var imageAllNoPrimary ImageAllNoPrimary

		imageAllNoPrimary.ID = image.ID
		imageAllNoPrimary.Index = i + 1
		imageAllNoPrimary.Judul = image.Judul
		imageAllNoPrimary.Keterangan = image.Keterangan
		imageAllNoPrimary.SlideShowImages = image.SlideShowImages
		imageAllNoPrimary.Link = image.Link
		imageAllNoPrimary.Utama = image.Utama
		imageAllNoPrimary.CreatedAt = image.CreatedAt

		images = append(images, imageAllNoPrimary)
	}

	imageFormatter.ImageNoPrimary = images

	return imageFormatter
}
