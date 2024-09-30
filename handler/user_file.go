package handler

import (
	"app-desa-kepuk/arsip"
	"app-desa-kepuk/file"
	"app-desa-kepuk/filedetail"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type fileUserHandler struct {
	fileService       file.Service       //1.file
	arsipService      arsip.Service      //2.arsip
	userService       user.Service       //3.user
	fileDetailService filedetail.Service //4.fileDetail
	sessionStore      *session.Store
}

func NewFileUserHandler(fileService file.Service, arsipService arsip.Service, userService user.Service, fileDetailService filedetail.Service, sessionStore *session.Store) *fileUserHandler {
	return &fileUserHandler{fileService, arsipService, userService, fileDetailService, sessionStore}
}

func (h *fileUserHandler) ShowFileAll(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}

	session, err := h.sessionStore.Get(c)
	defer session.Save()

	// ambil sessio pesan alert dan error form
	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-new-file")
		staussession = session.Get("msg-alert-new-file-status")
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

	file, err := h.fileDetailService.GetFileDetailAll()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	arsips, err := h.arsipService.GetAllArsip()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	data := map[string]interface{}{
		"file":   file,
		"arsip":  arsips,
		"layout": "table",
	}

	return c.Render("user/file/dasboard_file_list", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",

		"alert": template.HTML(alert),
	})
}

func (h *fileUserHandler) ShowFileDetail(c *fiber.Ctx) error {
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
		return c.Render("error400", fiber.Map{})
	}

	rowFile, err := h.fileDetailService.GetFileDetailByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	return c.Render("user/file/dasboard_file_detail", fiber.Map{
		"header": userMain,
		"data":   rowFile,
		"layout": "table",
	})
}

func (h *fileUserHandler) DownloadFile(c *fiber.Ctx) error {

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, err := h.fileService.GetFileByID(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileNameSTR, err := helper.GetFileNameDekrip(file.FileLocation)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileLocation := fmt.Sprintf("./derektori/file/%s", file.FileLocation)
	return c.Download(fileLocation, fileNameSTR)
}
