package handler

import (
	"app-desa-kepuk/arsipkategori"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type arsipKategoriHandler struct {
	service      arsipkategori.Service
	userService  user.Service
	sessionStore *session.Store
}

func NewArsipKategoriHandler(service arsipkategori.Service, userService user.Service, sessionStore *session.Store) *arsipKategoriHandler {
	return &arsipKategoriHandler{service, userService, sessionStore}
}

func (h *arsipKategoriHandler) ShowArsipKategoriList(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-new-arsip-kategori")
		staussession = session.Get("msg-alert-new-arsip-kategori-status")
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

	kategoriArsip, _ := h.service.GetAllArsipKategori()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	return c.Render("admin/arsipkategori/dasboard_arsip_kategori_list", fiber.Map{
		"header": userMain,
		"data":   kategoriArsip,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *arsipKategoriHandler) NewArsipKategori(c *fiber.Ctx) error {

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-new-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(arsipkategori.KategoriArsipInput)

	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-new-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if _, err := h.service.CreateArsipKategori(*input); err != nil {
		helper.AlertMassage("msg-alert-new-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-arsip-kategori", "Kategori baru berhasil di buat!.", "success", session)
	return c.Redirect(root)
}

func (h *arsipKategoriHandler) UpdateArsipKategoriView(c *fiber.Ctx) error {
	//form error
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()

	} else {

		sessionResult = session.Get("msg-alert-update-arsip-kategori")
		staussession = session.Get("msg-alert-update-arsip-kategori-status")
	}

	// end from error
	alert := helper.AlertString(sessionResult, staussession)

	// cookie user id
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

	kategori, err := h.service.GetArsipKategoriByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	return c.Render("admin/arsipkategori/dasboard_arsip_kategori_update", fiber.Map{
		"header": userMain,
		"data":   kategori,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *arsipKategoriHandler) UpdateArsipKategori(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-arsip-kategori", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	input := new(arsipkategori.KategoriArsipUpdate)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	kategoriRow, err := h.service.UpdateArsipKategori(*input)
	if err != nil {
		helper.AlertMassage("msg-alert-update-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Kategori %s berhasil di Ubah!", kategoriRow.Nama)
	helper.AlertMassage("msg-alert-update-arsip-kategori", returnString, "success", session)

	return c.Redirect(root)
}

func (h *arsipKategoriHandler) DeletedArsipKategoriSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {

		helper.AlertMassage("msg-alert-new-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedArsipKategoriSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-arsip-kategori", "Kategori berhasil di hapus!", "success", session)

	return c.Redirect(root)
}

func (h *arsipKategoriHandler) DeletedArsipKategoriRecycle(c *fiber.Ctx) error {
	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {

		helper.AlertMassage("msg-alert-restore-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedArsipKategori(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-arsip-kategori", "Kategori berhasil di hapus!", "success", session)

	return c.Redirect(root)
}

func (h *arsipKategoriHandler) ShowKategoriArsipListRecycle(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-restore-arsip-kategori")
		staussession = session.Get("msg-alert-restore-arsip-kategori-status")
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

	kategoriArsip, _ := h.service.GetAllArsipKategoriDeleted()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	return c.Render("admin/arsipkategori/dasboard_arsip_kategori_list_recycle", fiber.Map{
		"header": userMain,
		"data":   kategoriArsip,
		"layout": "table",

		"alert": template.HTML(alert),
	})
}

func (h *arsipKategoriHandler) RestoreArsipKategori(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.RestoreArsipKategori(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-arsip-kategori", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-arsip-kategori", "User berhasil di hapus!", "success", session)

	return c.Redirect(root)
}
