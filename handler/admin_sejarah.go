package handler

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/sejarah"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type sejarahHandler struct {
	service      sejarah.Service
	userService  user.Service
	sessionStore *session.Store
}

func NewSejarahHandler(service sejarah.Service, userService user.Service, sessionStore *session.Store) *sejarahHandler {
	return &sejarahHandler{service, userService, sessionStore}
}

func (h *sejarahHandler) ShowSejarahList(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-sejarah")
		staussession = session.Get("msg-alert-new-sejarah-status")
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

	informasi, _ := h.service.GetAllSejarahWeb()

	data := map[string]interface{}{
		"sejarah": informasi,
	}

	return c.Render("admin/sejarah/dasboard_sejarah_view", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *sejarahHandler) ShowSejarahDetail(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-detail-sejarah")
		staussession = session.Get("msg-alert-detail-sejarah-status")
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
		return c.Redirect("/dasboard/admin/sejarah/")
	}

	Sejarah, err := h.service.GetSejarahByID(uint(id))
	if err != nil {
		return c.Redirect("/dasboard/admin/sejarah/")
	}

	data := map[string]interface{}{
		"sejarah": Sejarah,
	}

	return c.Render("admin/sejarah/dasboard_sejarah_detail", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *sejarahHandler) NewSejarah(c *fiber.Ctx) error {
	root := c.Get("Referer")
	var fileto *multipart.FileHeader

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(sejarah.SejarahInput)
	if err := c.BodyParser(input); err != nil {

		helper.AlertMassage("msg-alert-new-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, _ := c.FormFile("image")

	var fileNameToEncripTo string
	if file == fileto {
		fileNameToEncripTo = ""

	} else {
		a, _ := helper.IsAllowedFileTypeImage(file)
		if a != true {
			helper.AlertMassage("msg-alert-new-sejarah", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileName := helper.StringWithoutSpaces(file.Filename)
		fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			helper.AlertMassage("msg-alert-new-sejarah", err.Error(), "error", session)
			return c.Redirect(root)
		}

		path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncrip)
		err = c.SaveFile(file, path)
		if err != nil {
			helper.AlertMassage("msg-alert-new-sejarah", err.Error(), "error", session)
			return c.Redirect(root)
		}
		fileNameToEncripTo = fileNameToEncrip

	}
	path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncripTo)

	if _, err := h.service.CreateSejarah(*input, fileNameToEncripTo); err != nil {
		os.Remove(path)
		helper.AlertMassage("msg-alert-new-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-sejarah", "Data baru berhasil di buat!.", "success", session)
	return c.Redirect(root)
}

func (h *sejarahHandler) UpdateSejarahView(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-update-sejarah")
		staussession = session.Get("msg-alert-update-sejarah-status")
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
		return c.Redirect("/dasboard/admin/sejarah")
	}

	Sejarah, err := h.service.GetSejarahByID(uint(id))
	if err != nil {
		return c.Redirect("/dasboard/admin/sejarah")
	}

	data := map[string]interface{}{
		"sejarah": Sejarah,
	}
	return c.Render("admin/sejarah/dasboard_sejarah_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *sejarahHandler) UpdateSejarah(c *fiber.Ctx) error {
	var fileto *multipart.FileHeader

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-sejarah", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	input := new(sejarah.SejarahUpdate)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, _ := c.FormFile("image")

	var fileNameToEncripTo string
	if file == fileto {
		fileNameToEncripTo = ""

	} else {
		a, _ := helper.IsAllowedFileTypeImage(file)
		if a != true {
			helper.AlertMassage("msg-alert-update-sejarah", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileName := helper.StringWithoutSpaces(file.Filename)
		fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			helper.AlertMassage("msg-alert-update-sejarah", err.Error(), "error", session)
			return c.Redirect(root)
		}

		path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncrip)
		err = c.SaveFile(file, path)
		if err != nil {
			helper.AlertMassage("msg-alert-update-sejarah", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileNameToEncripTo = fileNameToEncrip
	}

	_, err = h.service.UpdateSejarah(*input, fileNameToEncripTo)
	if err != nil {
		helper.AlertMassage("msg-alert-update-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-update-sejarah", "Data berhasil di ubah!", "success", session)
	return c.Redirect(root)
}

func (h *sejarahHandler) ShowSejarahListRecycle(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-restore-sejarah")
		staussession = session.Get("msg-alert-restore-sejarah-status")
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

	Sejarah, _ := h.service.GetAllSejarahDeleted()

	data := map[string]interface{}{
		"sejarah": Sejarah,
	}

	return c.Render("admin/sejarah/dasboard_sejarah_list_recycle", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *sejarahHandler) RestoreSejarah(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.RestoreSejarah(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-sejarah", "Data berhasil dipulihkan!", "success", session)

	return c.Redirect(root)
}

func (h *sejarahHandler) DeleteSejarahSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedSejarahSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-sejarah", "Data berhasil di hapus!", "success", session)

	return c.Redirect(root)
}

func (h *sejarahHandler) DeleteSejarah(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedSejarah(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-sejarah", "Data berhasil di hapus.", "success", session)

	return c.Redirect(root)
}

func (h *sejarahHandler) DeleteSejarahImage(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedSejarahImage(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-sejarah", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-sejarah", "Data berhasil di hapus.", "success", session)

	return c.Redirect(root)
}
