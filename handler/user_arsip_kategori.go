package handler

import (
	"app-desa-kepuk/arsipkategori"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type arsipKategoriUserHandler struct {
	service      arsipkategori.Service
	userService  user.Service
	sessionStore *session.Store
}

func NewArsipKategoriUserHandler(service arsipkategori.Service, userService user.Service, sessionStore *session.Store) *arsipKategoriUserHandler {
	return &arsipKategoriUserHandler{service, userService, sessionStore}
}

func (h *arsipKategoriUserHandler) ShowArsipKategoriList(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-new-arsip")
		staussession = session.Get("msg-alert-new-arsip-status")
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

	kategoriArsip, err := h.service.GetAllArsipKategori()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	return c.Render("user/arsipkategori/dasboard_arsip_kategori_list", fiber.Map{
		"header": userMain,
		"data":   kategoriArsip,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}
