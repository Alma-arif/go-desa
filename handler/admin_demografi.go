package handler

import (
	"app-desa-kepuk/demografi"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type demografiHandler struct {
	service      demografi.Service
	userService  user.Service
	sessionStore *session.Store
}

func NewDemografiHandler(service demografi.Service, userService user.Service, sessionStore *session.Store) *demografiHandler {
	return &demografiHandler{service, userService, sessionStore}
}

func (h *demografiHandler) ShowDemografiList(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-demografi")
		staussession = session.Get("msg-alert-new-demografi-status")
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

	informasi, _ := h.service.GetAllDemografiWeb()

	data := map[string]interface{}{
		"demografi": informasi,
	}

	return c.Render("admin/demografi/dasboard_demografi_view", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *demografiHandler) ShowDemografiDetail(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-detail-demografi")
		staussession = session.Get("msg-alert-detail-demografi-status")
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
		return c.Redirect("/dasboard/admin/demografi/")
	}

	Demografi, err := h.service.GetDemografiByID(uint(id))
	if err != nil {
		return c.Redirect("/dasboard/admin/demografi/")
	}

	data := map[string]interface{}{
		"demografi": Demografi,
	}

	return c.Render("admin/demografi/dasboard_demografi_detail", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *demografiHandler) NewDemografi(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(demografi.DemografiInput)
	if err := c.BodyParser(input); err != nil {

		helper.AlertMassage("msg-alert-new-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, err := c.FormFile("image")
	if err != nil {
		helper.AlertMassage("msg-alert-new-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	a, _ := helper.IsAllowedFileTypeImage(file)
	if a != true {
		helper.AlertMassage("msg-alert-new-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileName := helper.StringWithoutSpaces(file.Filename)
	fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
	if err != nil {
		helper.AlertMassage("msg-alert-new-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncrip)
	err = c.SaveFile(file, path)
	if err != nil {
		helper.AlertMassage("msg-alert-new-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if _, err := h.service.CreateDemografi(*input, fileNameToEncrip); err != nil {
		os.Remove(path)

		helper.AlertMassage("msg-alert-new-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-demografi", "Pengumuman baru berhasil di buat!.", "success", session)
	return c.Redirect(root)
}

func (h *demografiHandler) UpdateDemografiView(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-update-demografi")
		staussession = session.Get("msg-alert-update-demografi-status")
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
		return c.Redirect("/dasboard/admin/demografi")
	}

	Demografi, err := h.service.GetDemografiByID(uint(id))
	if err != nil {
		return c.Redirect("/dasboard/admin/demografi")
	}

	data := map[string]interface{}{
		"demografi": Demografi,
	}
	return c.Render("admin/demografi/dasboard_demografi_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *demografiHandler) UpdateDemografi(c *fiber.Ctx) error {
	var fileto *multipart.FileHeader

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-demografi", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	input := new(demografi.DemografiUpdate)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, _ := c.FormFile("image")

	var fileNameToEncripTo string
	if file == fileto {
		fileNameToEncripTo = ""

	} else {
		a, _ := helper.IsAllowedFileTypeImage(file)
		if a != true {
			helper.AlertMassage("msg-alert-update-demografi", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileName := helper.StringWithoutSpaces(file.Filename)
		fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			helper.AlertMassage("msg-alert-update-demografi", err.Error(), "error", session)
			return c.Redirect(root)
		}

		path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncrip)
		err = c.SaveFile(file, path)
		if err != nil {
			helper.AlertMassage("msg-alert-update-demografi", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileNameToEncripTo = fileNameToEncrip
	}

	_, err = h.service.UpdateDemografi(*input, fileNameToEncripTo)
	if err != nil {
		helper.AlertMassage("msg-alert-update-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-update-demografi", "Data berhasil di ubah!", "success", session)
	return c.Redirect(root)
}

func (h *demografiHandler) ShowDemografiListRecycle(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-restore-demografi")
		staussession = session.Get("msg-alert-restore-demografi-status")
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

	Demografi, _ := h.service.GetAllDemografiDeleted()

	data := map[string]interface{}{
		"demografi": Demografi,
	}

	return c.Render("admin/demografi/dasboard_demografi_list_recycle", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *demografiHandler) RestoreDemografi(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.RestoreDemografi(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-demografi", "Data berhasil dipulihkan!", "success", session)

	return c.Redirect(root)
}

func (h *demografiHandler) DeleteDemografiSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedDemografiSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-demografi", "Data berhasil di hapus!", "success", session)

	return c.Redirect(root)
}

func (h *demografiHandler) DeleteDemografi(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedDemografi(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-demografi", "Data berhasil di hapus.", "success", session)

	return c.Redirect(root)
}

func (h *demografiHandler) DeleteDemografiImage(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedDemografiImage(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-demografi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-demografi", "Data berhasil di hapus.", "success", session)

	return c.Redirect(root)
}
