package handler

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"app-desa-kepuk/visimisidesa"
	"fmt"
	"html/template"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type visiMisiHandler struct {
	service      visimisidesa.Service
	userService  user.Service
	sessionStore *session.Store
}

func NewVisiMisiHandler(service visimisidesa.Service, userService user.Service, sessionStore *session.Store) *visiMisiHandler {
	return &visiMisiHandler{service, userService, sessionStore}
}

func (h *visiMisiHandler) ShowVisiMisiList(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-visi-misi")
		staussession = session.Get("msg-alert-new-visi-misi-status")
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

	informasi, _ := h.service.GetAllVisiMisiWeb()

	data := map[string]interface{}{
		"visimisi": informasi,
	}

	return c.Render("admin/visimisi/dasboard_visi_misi_view", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *visiMisiHandler) ShowVisiMisiDetail(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-detail-visi-misi")
		staussession = session.Get("msg-alert-detail-visi-misi-status")
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
		return c.Redirect("/dasboard/admin/visi-misi/")
	}

	visiMisi, err := h.service.GetVisiMisiByID(uint(id))
	if err != nil {
		return c.Redirect("/dasboard/admin/visi-misi/")
	}

	data := map[string]interface{}{
		"visimisi": visiMisi,
	}

	return c.Render("admin/visimisi/dasboard_visi_misi_detail", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *visiMisiHandler) NewVisiMisi(c *fiber.Ctx) error {
	root := c.Get("Referer")
	var fileto *multipart.FileHeader

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(visimisidesa.VisiMisiInput)
	if err := c.BodyParser(input); err != nil {

		helper.AlertMassage("msg-alert-new-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}
	file, _ := c.FormFile("image")

	var fileNameToEncripTo string
	if file == fileto {
		fileNameToEncripTo = ""

	} else {
		a, _ := helper.IsAllowedFileTypeImage(file)
		if a != true {
			helper.AlertMassage("msg-alert-new-visi-misi", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileName := helper.StringWithoutSpaces(file.Filename)
		fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			helper.AlertMassage("msg-alert-new-visi-misi", err.Error(), "error", session)
			return c.Redirect(root)
		}

		path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncrip)
		err = c.SaveFile(file, path)
		if err != nil {
			helper.AlertMassage("msg-alert-new-visi-misi", err.Error(), "error", session)
			return c.Redirect(root)
		}
		fileNameToEncripTo = fileNameToEncrip

	}

	path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncripTo)

	if _, err := h.service.CreateVisiMisi(*input, fileNameToEncripTo); err != nil {
		os.Remove(path)
		helper.AlertMassage("msg-alert-new-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-visi-misi", "Data baru berhasil di buat!.", "success", session)
	return c.Redirect(root)
}

func (h *visiMisiHandler) UpdateVisiMisiView(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-update-visi-misi")
		staussession = session.Get("msg-alert-update-visi-misi-status")
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
		return c.Redirect("/dasboard/admin/visi-misi")
	}

	visiMisi, err := h.service.GetVisiMisiByID(uint(id))
	if err != nil {
		return c.Redirect("/dasboard/admin/visi-misi")
	}

	data := map[string]interface{}{
		"visimisi": visiMisi,
	}
	return c.Render("admin/visimisi/dasboard_visi_misi_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *visiMisiHandler) UpdateVisiMisi(c *fiber.Ctx) error {
	var fileto *multipart.FileHeader

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-visi-misi", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	input := new(visimisidesa.VisiMisiUpdate)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, _ := c.FormFile("image")

	var fileNameToEncripTo string
	if file == fileto {
		fileNameToEncripTo = ""

	} else {
		a, _ := helper.IsAllowedFileTypeImage(file)
		if a != true {
			helper.AlertMassage("msg-alert-update-visi-misi", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileName := helper.StringWithoutSpaces(file.Filename)
		fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			helper.AlertMassage("msg-alert-update-visi-misi", err.Error(), "error", session)
			return c.Redirect(root)
		}

		path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncrip)
		err = c.SaveFile(file, path)
		if err != nil {
			helper.AlertMassage("msg-alert-update-visi-misi", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileNameToEncripTo = fileNameToEncrip
	}

	_, err = h.service.UpdateVisiMisi(*input, fileNameToEncripTo)
	if err != nil {
		helper.AlertMassage("msg-alert-update-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-update-visi-misi", "Data berhasil di ubah!", "success", session)
	return c.Redirect(root)
}

func (h *visiMisiHandler) ShowVisiMisiListRecycle(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-restore-visi-misi")
		staussession = session.Get("msg-alert-restore-visi-misi-status")
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

	visiMisi, _ := h.service.GetAllVisiMisiDeleted()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"visimisi": visiMisi,
	}

	return c.Render("admin/visimisi/dasboard_visi_misi_list_recycle", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *visiMisiHandler) RestorevisiMisi(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.RestoreVisiMisi(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-visi-misi", "Data berhasil dipulihkan!", "success", session)

	return c.Redirect(root)
}

func (h *visiMisiHandler) DeletevisiMisiSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedVisiMisiSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-visi-misi", "Data berhasil di hapus!", "success", session)

	return c.Redirect(root)
}

func (h *visiMisiHandler) DeleteVisiMisi(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedVisiMisi(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-visi-misi", "Data berhasil di hapus.", "success", session)

	return c.Redirect(root)
}

func (h *visiMisiHandler) DeleteVisiMisiImage(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedVisiMisiImage(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-visi-misi", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-visi-misi", "Data berhasil di hapus.", "success", session)

	return c.Redirect(root)
}
