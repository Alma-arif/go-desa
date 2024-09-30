package handler

import (
	"app-desa-kepuk/berita"
	beritaimage "app-desa-kepuk/beritaImage"
	"app-desa-kepuk/beritakategori"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type beritaImageHandler struct {
	Service         beritaimage.Service
	beritaService   berita.Service
	kategoriService beritakategori.Service
	userService     user.Service
	sessionStore    *session.Store
}

func NewBeritaImageHandler(Service beritaimage.Service, beritaService berita.Service, kategoriService beritakategori.Service, userService user.Service, sessionStore *session.Store) *beritaImageHandler {
	return &beritaImageHandler{Service, beritaService, kategoriService, userService, sessionStore}
}

func (h *beritaImageHandler) NewBeritaImage(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, err := c.FormFile("image_berita")
	if err != nil {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	a, _ := helper.IsAllowedFileTypeImage(file)
	if a != true {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileName := helper.StringWithoutSpaces(file.Filename)
	fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
	if err != nil {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	da := float64(file.Size) / 1024 / 1024
	if da > float64(5) {
		helper.AlertMassage("msg-alert-new-berita", "Ukuran gambar Melebih yang sistem tentukan maksimal 4MB", "error", session)
		return c.Redirect(root)
	}

	path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncrip)

	err = c.SaveFile(file, path)
	if err != nil {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(beritaimage.ImageBeritaInput)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if _, err := h.Service.CreateBeritaImage(*input, fileNameToEncrip); err != nil {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-berita", "Berita baru berhasil di buat!.", "success", session)
	return c.Redirect(root)
}

func (h *beritaImageHandler) DeleteBeritaImage(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-detail-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-detail-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.Service.DeleteImageBerita(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-detail-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-detail-berita", "Image berhasil di hapus.", "success", session)

	return c.Redirect(root)
}
