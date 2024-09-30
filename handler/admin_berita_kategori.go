package handler

import (
	"app-desa-kepuk/beritakategori"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type beritaKategoriHandler struct {
	service      beritakategori.Service
	userService  user.Service
	sessionStore *session.Store
}

func NewBeritaKategoriHandler(service beritakategori.Service, userService user.Service, sessionStore *session.Store) *beritaKategoriHandler {
	return &beritaKategoriHandler{service, userService, sessionStore}
}

func (h *beritaKategoriHandler) ShowBeritaKategoriList(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-new-berita-kategori")
		staussession = session.Get("msg-alert-new-berita-kategori-status")
	}

	alert := helper.AlertString(sessionResult, staussession)

	cookisUserId := c.Cookies("sessionLog")
	idUser, err := helper.GetSessionID(cookisUserId)
	if err != nil {
		return c.Redirect("/login")
	}

	userMain, err := h.userService.GetUserByID(idUser)
	if err != nil {
		return c.Redirect("/login")
	}

	kategoriBerita, _ := h.service.GetAllBeritaKategori()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	return c.Render("admin/beritakategori/dasboard_berita_kategori_list", fiber.Map{
		"header": userMain,
		"data":   kategoriBerita,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *beritaKategoriHandler) NewBeritaKategori(c *fiber.Ctx) error {

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-new-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(beritakategori.KategoriBeritaInput)

	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-new-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if _, err := h.service.CreateBeritaKategori(*input); err != nil {
		helper.AlertMassage("msg-alert-new-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-berita-kategori", "Kategori baru berhasil di buat!.", "success", session)
	return c.Redirect(root)
}

func (h *beritaKategoriHandler) UpdateBeritaKategoriView(c *fiber.Ctx) error {
	//form error
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()

	} else {

		sessionResult = session.Get("msg-alert-update-berita-kategori")
		staussession = session.Get("msg-alert-update-berita-kategori-status")
	}

	alert := helper.AlertString(sessionResult, staussession)

	cookisUserId := c.Cookies("sessionLog")

	idUser, err := helper.GetSessionID(cookisUserId)
	if err != nil {
		return c.Redirect("/login")
	}

	userMain, err := h.userService.GetUserByID(idUser)
	if err != nil {
		return c.Redirect("/login")
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Redirect("/dasboard/admin/arsip/kategori")
	}

	kategori, err := h.service.GetBeritaKategoriByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	return c.Render("admin/beritakategori/dasboard_berita_kategori_update", fiber.Map{
		"header": userMain,
		"data":   kategori,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *beritaKategoriHandler) UpdateBeritaKategori(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-berita-kategori", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	input := new(beritakategori.KategoriBeritaUpdate)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	kategoriRow, err := h.service.UpdateBeritaKategori(*input)
	if err != nil {
		helper.AlertMassage("msg-alert-update-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Kategori %s berhasil di Ubah!", kategoriRow.Nama)
	helper.AlertMassage("msg-alert-update-berita-kategori", returnString, "success", session)

	return c.Redirect(root)
}

func (h *beritaKategoriHandler) DeletedBeritaKategoriSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-new-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedBeritaKategoriSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-berita-kategori", "Kategori berhasil di hapus!", "success", session)
	return c.Redirect(root)
}

func (h *beritaKategoriHandler) DeletedBeritaKategoriRecycle(c *fiber.Ctx) error {
	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedBeritaKategori(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-berita-kategori", "Kategori berhasil di hapus!", "success", session)
	return c.Redirect(root)
}

func (h *beritaKategoriHandler) ShowKategoriBeritaListRecycle(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-restore-berita-kategori")
		staussession = session.Get("msg-alert-restore-berita-kategori-status")
	}

	alert := helper.AlertString(sessionResult, staussession)

	cookisUserId := c.Cookies("sessionLog")
	idUser, err := helper.GetSessionID(cookisUserId)
	if err != nil {
		return c.Redirect("/login")
	}

	userMain, err := h.userService.GetUserByID(idUser)
	if err != nil {
		return c.Redirect("/login")
	}

	kategoriBerita, _ := h.service.GetAllBeritaKategoriDeleted()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	return c.Render("admin/beritakategori/dasboard_berita_kategori_list_recycle", fiber.Map{
		"header": userMain,
		"data":   kategoriBerita,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *beritaKategoriHandler) RestoreBeritaKategori(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.RestoreBeritaKategori(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-berita-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-berita-kategori", "katgori berhasil di hapus!", "success", session)

	return c.Redirect(root)
}
