package handler

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/pengumuman"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type pengumumanHandler struct {
	service      pengumuman.Service
	userService  user.Service
	sessionStore *session.Store
}

func NewPengumumanHandler(service pengumuman.Service, userService user.Service, sessionStore *session.Store) *pengumumanHandler {
	return &pengumumanHandler{service, userService, sessionStore}
}

func (h *pengumumanHandler) ShowPengumumanList(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-pengumuman")
		staussession = session.Get("msg-alert-new-pengumuman-status")
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

	informasi, _ := h.service.GetAllPengumuman()

	data := map[string]interface{}{
		"pengumuman": informasi,
	}

	return c.Render("admin/pengumuman/dasboard_pengumuman_list", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *pengumumanHandler) ShowPengumumanDetail(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-detail-pengumuman")
		staussession = session.Get("msg-alert-detail-pengumuman-status")
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
		return c.Redirect("/dasboard/admin/pengumuman/")
	}

	berita, err := h.service.GetPengumumanByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	// image, _ := h.imageService.GetBeritaImageByBeritaID(berita.ID)
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"pengumuman": berita,
	}

	return c.Render("admin/pengumuman/dasboard_pengumuman_detail", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *pengumumanHandler) NewPengumuman(c *fiber.Ctx) error {
	root := c.Get("Referer")
	var fileto *multipart.FileHeader

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	cookisUserId := c.Cookies("sessionLog")
	idUser, err := helper.GetSessionID(cookisUserId)
	if err != nil {
		return c.Redirect("/login")
	}

	input := new(pengumuman.PengumumanInput)
	if err := c.BodyParser(input); err != nil {

		helper.AlertMassage("msg-alert-new-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, _ := c.FormFile("image")
	// if err != nil {
	// 	helper.AlertMassage("msg-alert-new-pengumuman", err.Error(), "error", session)
	// 	return c.Redirect(root)
	// }
	var fileNameToEncripTo string
	if file == fileto {
		fileNameToEncripTo = ""

	} else {
		a, _ := helper.IsAllowedFileTypeImage(file)
		if a != true {
			helper.AlertMassage("msg-alert-new-pengumuman", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileName := helper.StringWithoutSpaces(file.Filename)
		fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			helper.AlertMassage("msg-alert-new-pengumuman", err.Error(), "error", session)
			return c.Redirect(root)
		}

		path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncrip)
		err = c.SaveFile(file, path)
		if err != nil {
			helper.AlertMassage("msg-alert-new-pengumuman", err.Error(), "error", session)
			return c.Redirect(root)
		}
		fileNameToEncripTo = fileNameToEncrip

	}

	path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncripTo)

	if _, err := h.service.CreatePengumuman(*input, fileNameToEncripTo, idUser); err != nil {

		os.Remove(path)
		// if err != nil {
		// 	helper.AlertMassage("msg-alert-new-pengumuman", err.Error(), "error", session)
		// return c.Redirect(root)
		// }
		helper.AlertMassage("msg-alert-new-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-pengumuman", "Pengumuman baru berhasil di buat!.", "success", session)
	return c.Redirect(root)
}

func (h *pengumumanHandler) UpdatePengumumanView(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-update-pengumuman")
		staussession = session.Get("msg-alert-update-pengumuman-status")
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
		return c.Redirect("/dasboard/admin/pengumuman")
	}

	pengumuman, err := h.service.GetPengumumanByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	// kategori, err := h.kategoriService.GetAllBeritaKategori()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"pengumuman": pengumuman,
	}
	return c.Render("admin/pengumuman/dasboard_pengumuman_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *pengumumanHandler) UpdatePengumuman(c *fiber.Ctx) error {
	var fileto *multipart.FileHeader

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-pengumuman", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	input := new(pengumuman.PengumumanUpdate)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, _ := c.FormFile("image")

	var fileNameToEncripTo string
	if file == fileto {
		fileNameToEncripTo = ""

	} else {

		a, _ := helper.IsAllowedFileTypeImage(file)
		if a != true {
			helper.AlertMassage("msg-alert-update-pengumuman", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileName := helper.StringWithoutSpaces(file.Filename)
		fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			helper.AlertMassage("msg-alert-update-pengumuman", err.Error(), "error", session)
			return c.Redirect(root)
		}

		path := fmt.Sprintf("derektori/images_berita/%s", fileNameToEncrip)
		err = c.SaveFile(file, path)
		if err != nil {
			helper.AlertMassage("msg-alert-update-pengumuman", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileNameToEncripTo = fileNameToEncrip
	}

	pengumumanRow, err := h.service.UpdatePengumuman(*input, fileNameToEncripTo)
	if err != nil {
		// os.Remove(path)
		fmt.Println("a10")
		helper.AlertMassage("msg-alert-update-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Berita %s berhasil di ubah!", pengumumanRow.Judul)
	helper.AlertMassage("msg-alert-update-berita", returnString, "success", session)

	return c.Redirect(root)
}

func (h *pengumumanHandler) ShowPengumumanListRecycle(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-restore-pengumuman")
		staussession = session.Get("msg-alert-restore-pengumuman-status")
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

	pengumuman, _ := h.service.GetAllPengumumanDeleted()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"pengumuman": pengumuman,
	}

	return c.Render("admin/pengumuman/dasboard_pengumuman_list_recycle", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *pengumumanHandler) RestorePengumuman(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.RestorePengumuman(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-pengumuman", "Pengumuman berhasil dipulihkan!", "success", session)

	return c.Redirect(root)
}

func (h *pengumumanHandler) DeletePengumumanSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	pengumuman, err := h.service.DeletedPengumumanSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Pengumuman %s berhasil di hapus. ", pengumuman.Judul)
	helper.AlertMassage("msg-alert-new-pengumuman", returnString, "success", session)

	return c.Redirect(root)
}

func (h *pengumumanHandler) DeletePengumuman(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	pengumuman, err := h.service.DeletedPengumuman(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-pengumuman", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Pengumuman %s berhasil di hapus.", pengumuman.Judul)
	helper.AlertMassage("msg-alert-restore-pengumuman", returnString, "success", session)

	return c.Redirect(root)
}
