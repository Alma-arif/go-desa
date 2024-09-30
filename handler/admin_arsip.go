package handler

import (
	"app-desa-kepuk/arsip"
	"app-desa-kepuk/arsipkategori"
	"app-desa-kepuk/file"
	"app-desa-kepuk/filedetail"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type arsipHandler struct {
	Service           arsip.Service
	KategoriArsip     arsipkategori.Service
	fileService       file.Service
	fileDetailService filedetail.Service
	userService       user.Service
	sessionStore      *session.Store
}

func NewArsipHandler(arsipeService arsip.Service, KategoriArsip arsipkategori.Service, fileService file.Service, fileDetailService filedetail.Service, userService user.Service, sessionStore *session.Store) *arsipHandler {
	return &arsipHandler{arsipeService, KategoriArsip, fileService, fileDetailService, userService, sessionStore}
}

func (h *arsipHandler) ShowArsipList(c *fiber.Ctx) error {

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

	arsip, _ := h.Service.GetAllArsip()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	kategori, _ := h.KategoriArsip.GetAllArsipKategori()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"arsip":    arsip,
		"kategori": kategori,
	}

	return c.Render("admin/arsip/dasboard_arsip_list", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *arsipHandler) ShowArsipDetail(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-detail-arsip")
		staussession = session.Get("msg-alert-detail-arsip-status")
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
		return c.Redirect("/dasboard/arsip/list")
	}

	arsip, _ := h.Service.GetArsipByID(uint(id))

	file, _ := h.fileDetailService.GetFileAllArsipID(arsip.ID)

	fileResult, _ := h.fileService.GetFileAllByArsipIDNull()

	data := map[string]interface{}{
		"arsip":      arsip,
		"file":       file,
		"resultFile": fileResult,
	}
	return c.Render("admin/arsip/dasboard_arsip_detail", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *arsipHandler) NewArsip(c *fiber.Ctx) error {

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputNew := new(arsip.CreateArsipInput)

	if err := c.BodyParser(inputNew); err != nil {
		helper.AlertMassage("msg-alert-new-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if _, err := h.Service.CreateArsip(*inputNew); err != nil {
		helper.AlertMassage("msg-alert-new-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-arsip", "Arsip baru berhasil ditambahkan!", "success", session)
	// end
	return c.Redirect(root)
}

func (h *arsipHandler) UpdateArsipView(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-update-arsip")
		staussession = session.Get("msg-alert-update-arsip-status")
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
		return c.Redirect("/dasboard/admin/arsip")
	}

	arsip, err := h.Service.GetArsipByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	kategori, _ := h.KategoriArsip.GetAllArsipKategori()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"arsip":    arsip,
		"kategori": kategori,
	}
	return c.Render("admin/arsip/dasboard_arsip_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *arsipHandler) UpdateArsip(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-arsip", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	input := new(arsip.UpdateArsipInput)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}
	arsipRow, err := h.Service.UpdateArsip(*input)

	if err != nil {
		helper.AlertMassage("msg-alert-update-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Arsip %s berhasil di ubah!", arsipRow.Nama)
	helper.AlertMassage("msg-alert-update-arsip", returnString, "success", session)

	return c.Redirect(root)
}

func (h *arsipHandler) ShowArsipListRecycle(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-restore-arsip")
		staussession = session.Get("msg-alert-restore-arsip-status")
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

	arsip, _ := h.Service.GetAllArsipDeleted()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"arsip": arsip,
	}

	return c.Render("admin/arsip/dasboard_arsip_list_recycle", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *arsipHandler) RestoreArsip(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.Service.RestoreArsip(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-arsip", "User berhasil di hapus!", "success", session)

	return c.Redirect(root)
}

func (h *arsipHandler) DeleteArsipSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	arsip, err := h.Service.DeletedSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Arsip %s berhasil di hapus. ", arsip.Nama)
	helper.AlertMassage("msg-alert-new-arsip", returnString, "success", session)

	return c.Redirect(root)
}

func (h *arsipHandler) DeleteArsip(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	arsip, err := h.Service.Deleted(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Arsip %s berhasil di hapus.", arsip.Nama)
	helper.AlertMassage("msg-alert-restore-arsip", returnString, "success", session)

	return c.Redirect(root)
}
