package handler

import (
	"app-desa-kepuk/arsip"
	"app-desa-kepuk/arsipkategori"
	"app-desa-kepuk/file"
	"app-desa-kepuk/filedetail"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type arsipUserHandler struct {
	Service           arsip.Service
	KategoriArsip     arsipkategori.Service
	fileService       file.Service
	fileDetailService filedetail.Service
	userService       user.Service
	sessionStore      *session.Store
}

func NewArsipUSerHandler(arsipeService arsip.Service, KategoriArsip arsipkategori.Service, fileService file.Service, fileDetailService filedetail.Service, userService user.Service, sessionStore *session.Store) *arsipUserHandler {
	return &arsipUserHandler{arsipeService, KategoriArsip, fileService, fileDetailService, userService, sessionStore}
}

func (h *arsipUserHandler) ShowArsipList(c *fiber.Ctx) error {

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

	arsip, err := h.Service.GetAllArsip()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	kategori, err := h.KategoriArsip.GetAllArsipKategori()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	data := map[string]interface{}{
		"arsip":    arsip,
		"kategori": kategori,
	}

	return c.Render("user/arsip/dasboard_arsip_list", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *arsipUserHandler) ShowArsipDetail(c *fiber.Ctx) error {
	var idUser uint

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
		return c.Redirect("/dasboard/arsip/list")
	}

	arsip, err := h.Service.GetArsipByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	file, err := h.fileService.GetFileByArsipID(arsip.ID)
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	data := map[string]interface{}{
		"arsip": arsip,
		"file":  file,
	}
	return c.Render("user/arsip/dasboard_arsip_detail", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
	})
}
