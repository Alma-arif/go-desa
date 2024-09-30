package handler

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/profiledesa"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type profileDesaHandler struct {
	service      profiledesa.Service
	userService  user.Service
	sessionStore *session.Store
}

func NewProfileDesaHandler(service profiledesa.Service, userService user.Service, sessionStore *session.Store) *profileDesaHandler {
	return &profileDesaHandler{service, userService, sessionStore}
}

func (h *profileDesaHandler) ShowProfileDesaList(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-new-profile-desa")
		staussession = session.Get("msg-alert-new-profile-desa-status")
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

	ProfileDesa, _ := h.service.GetAllProfileDesaWeb()

	data := map[string]interface{}{
		"profiledesa": ProfileDesa,
	}

	return c.Render("admin/profiledesa/dasboard_profile_desa_view", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *profileDesaHandler) ShowProfileDesaDetail(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-detail-profile-desa")
		staussession = session.Get("msg-alert-detail-profile-desa-status")
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
		return c.Redirect("/dasboard/admin/profile-desa/")
	}

	profileDesa, err := h.service.GetProfileDesaByID(uint(id))
	if err != nil {
		return c.Redirect("/dasboard/admin/profile-desa/")
	}

	data := map[string]interface{}{
		"profiledesa": profileDesa,
	}

	return c.Render("admin/profiledesa/dasboard_profile_desa_detail", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *profileDesaHandler) NewProfileDesa(c *fiber.Ctx) error {
	root := c.Get("Referer")
	var fileto *multipart.FileHeader

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(profiledesa.ProfileDesaInput)
	if err := c.BodyParser(input); err != nil {

		helper.AlertMassage("msg-alert-new-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, _ := c.FormFile("image")

	var fileNameToEncripTo string
	if file == fileto {
		fileNameToEncripTo = ""

	} else {
		a, _ := helper.IsAllowedFileTypeImage(file)
		if a != true {
			helper.AlertMassage("msg-alert-new-profile-desa", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileName := helper.StringWithoutSpaces(file.Filename)
		fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			helper.AlertMassage("msg-alert-new-profile-desa", err.Error(), "error", session)
			return c.Redirect(root)
		}

		path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncrip)
		err = c.SaveFile(file, path)
		if err != nil {
			helper.AlertMassage("msg-alert-new-profile-desa", err.Error(), "error", session)
			return c.Redirect(root)
		}
		fileNameToEncripTo = fileNameToEncrip

	}
	path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncripTo)

	if _, err := h.service.CreateProfileDesa(*input, fileNameToEncripTo); err != nil {
		os.Remove(path)

		helper.AlertMassage("msg-alert-new-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-profile-desa", "Data baru berhasil di buat!.", "success", session)
	return c.Redirect(root)
}

func (h *profileDesaHandler) UpdateProfileDesaView(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-update-profile-desa")
		staussession = session.Get("msg-alert-update-profile-desa-status")
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
		return c.Redirect("/dasboard/admin/profile-desa")
	}

	profileDesa, err := h.service.GetProfileDesaByID(uint(id))
	if err != nil {
		return c.Redirect("/dasboard/admin/profile-desa")
	}

	data := map[string]interface{}{
		"profiledesa": profileDesa,
	}
	return c.Render("admin/profiledesa/dasboard_profile_desa_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *profileDesaHandler) UpdateProfileDesa(c *fiber.Ctx) error {
	var fileto *multipart.FileHeader

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-profile-desa", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	input := new(profiledesa.ProfileDesaUpdate)

	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, _ := c.FormFile("image")

	var fileNameToEncripTo string

	if file == fileto {
		fileNameToEncripTo = ""

	} else {
		a, _ := helper.IsAllowedFileTypeImage(file)
		if a != true {
			helper.AlertMassage("msg-alert-update-profile-desa", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileName := helper.StringWithoutSpaces(file.Filename)
		fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			helper.AlertMassage("msg-alert-update-profile-desa", err.Error(), "error", session)
			return c.Redirect(root)
		}

		path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncrip)
		err = c.SaveFile(file, path)
		if err != nil {
			helper.AlertMassage("msg-alert-update-profile-desa", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileNameToEncripTo = fileNameToEncrip
	}

	_, err = h.service.UpdateProfileDesa(*input, fileNameToEncripTo)
	if err != nil {
		helper.AlertMassage("msg-alert-update-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-update-profile-desa", "Data berhasil di ubah!", "success", session)
	return c.Redirect(root)
}

func (h *profileDesaHandler) ShowProfileDesaListRecycle(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-restore-profile-desa")
		staussession = session.Get("msg-alert-restore-profile-desa-status")
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

	profileDesa, _ := h.service.GetAllProfileDesaDeleted()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"profiledesa": profileDesa,
	}

	return c.Render("admin/profiledesa/dasboard_profile_desa_list_recycle", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *profileDesaHandler) RestoreProfileDesa(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.RestoreProfileDesa(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-profile-desa", "Data berhasil dipulihkan!", "success", session)

	return c.Redirect(root)
}

func (h *profileDesaHandler) DeleteProfileDesaSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedProfileDesaSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-profile-desa", "Data berhasil di hapus!", "success", session)

	return c.Redirect(root)
}

func (h *profileDesaHandler) DeleteProfileDesa(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedProfileDesa(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-profile-desa", "Data berhasil di hapus.", "success", session)

	return c.Redirect(root)
}

func (h *profileDesaHandler) DeleteProfileDesaImage(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedProfileDesaImage(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-profile-desa", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-profile-desa", "Data berhasil di hapus.", "success", session)

	return c.Redirect(root)
}
