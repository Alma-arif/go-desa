package handler

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/pegawai"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type pegawaiHandler struct {
	service      pegawai.Service
	userService  user.Service
	sessionStore *session.Store
}

func NewPegawaiHandler(service pegawai.Service, userService user.Service, sessionStore *session.Store) *pegawaiHandler {
	return &pegawaiHandler{service, userService, sessionStore}
}

func (h *pegawaiHandler) ShowPegawaiList(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-pegawai")
		staussession = session.Get("msg-alert-new-pegawai-status")
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

	pegawai, _ := h.service.GetAllPegawai()

	data := map[string]interface{}{
		"pegawai": pegawai,
	}

	return c.Render("admin/pegawai/dasboard_pegawai_list", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *pegawaiHandler) ShowPegawaiDetail(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-detail-pegawai")
		staussession = session.Get("msg-alert-detail-pegawai-status")
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
		return c.Redirect("/dasboard/admin/pegawai/")
	}

	pegawai, err := h.service.GetPegawaiByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	// image, _ := h.imageService.GetBeritaImageByBeritaID(berita.ID)
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"pegawai": pegawai,
	}

	return c.Render("admin/pegawai/dasboard_pegawai_detail", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *pegawaiHandler) NewPegawai(c *fiber.Ctx) error {

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(pegawai.PegawaiInput)
	// fmt.Println("a 7 :", input)

	if err := c.BodyParser(input); err != nil {
		// fmt.Println("a 6 :", err)

		helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	pegawaiRow, err := h.service.CreatePegawai(*input)
	if err != nil {
		helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-pegawai", "Resgistrasi pegawai baru berhasil.", "success", session)
	pathRoot := fmt.Sprintf("/dasboard/admin/pegawai/profile/image/%d", pegawaiRow.ID)
	return c.Redirect(pathRoot)
}

func (h *pegawaiHandler) UploadImageProfilePegawaiView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-pegawai")
		staussession = session.Get("msg-alert-new-pegawai-status")
	}

	alert := helper.AlertString(sessionResult, staussession)

	cookisUserId := c.Cookies("sessionLog")
	idUser, err := helper.GetSessionID(cookisUserId)
	if err != nil {
		return c.Redirect("/login")
	}

	userMain, err := h.service.GetPegawaiByID(idUser)
	if err != nil {
		return c.Redirect("/login")
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Redirect("/login")
	}

	userRow, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		return c.Redirect("/dasboard/admin/pegawai")
	}
	return c.Render("admin/user/dasboard_pegawai_image_profile", fiber.Map{
		"header": userMain,
		"data":   userRow,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *pegawaiHandler) UploadImageProfilePegawai(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, _ := h.sessionStore.Get(c)
	id := c.FormValue("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, err := c.FormFile("image_avatar")
	if err != nil {
		helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	a, _ := helper.IsAllowedFileTypeImage(file)
	if a != true {
		helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileName := helper.StringWithoutSpaces(file.Filename)
	fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
	if err != nil {
		helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	path := fmt.Sprintf("derektori/image/%s", fileNameToEncrip)
	userRow, err := h.userService.GetUserByID(uint(idInt))
	if err != nil {
		helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	// update image profile
	if userRow.ProfileFile != "" {
		err = c.SaveFile(file, path)
		if err != nil {
			helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
			return c.Redirect(root)
		}

		_, err = h.service.CreateImageProfilePegawai(fileNameToEncrip, userRow.ID)
		if err != nil {
			helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
			return c.Redirect(root)
		}

		pathRemoveFile := fmt.Sprintf("derektori/image/%s", userRow.ProfileFile)
		err = os.Remove(pathRemoveFile)
		if err != nil {
			helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
			return c.Redirect(root)
		}

		helper.AlertMassage("msg-alert-new-pegawai", "Image profile berhasil di perbarui!", "success", session)
		return c.Redirect("/dasboard/admin/pegawai")
	}
	// end update image profile

	err = c.SaveFile(file, path)
	if err != nil {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.userService.CreateImageProfile(fileNameToEncrip, userRow.ID)
	if err != nil {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-user", "Image profile berhasil di upload !", "success", session)
	return c.Redirect("/dasboard/admin/user")
}

func (h *pegawaiHandler) UpdatePegawaiView(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-update-pegawai")
		staussession = session.Get("msg-alert-update-pegawai-status")
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
		return c.Redirect("/dasboard/admin/pegawai")
	}

	pegawai, err := h.service.GetPegawaiByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	data := map[string]interface{}{
		"pegawai": pegawai,
	}
	return c.Render("admin/pegawai/dasboard_pegawai_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *pegawaiHandler) UpdatePegawai(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-pegawai", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	input := new(pegawai.PegawaiUpdate)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	pegawaiRow, err := h.service.UpdatePegawai(*input)
	if err != nil {
		helper.AlertMassage("msg-alert-update-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Pegawai %s berhasil di ubah!", pegawaiRow.Nama)
	helper.AlertMassage("msg-alert-update-pegawai", returnString, "success", session)

	return c.Redirect(root)
}

func (h *pegawaiHandler) ShowPegawaiListRecycle(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-restore-pegawai")
		staussession = session.Get("msg-alert-restore-pegawai-status")
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

	pegawai, _ := h.service.GetAllPegawaiDeletedAt()

	data := map[string]interface{}{
		"pegawai": pegawai,
	}

	return c.Render("admin/pegawai/dasboard_pegawai_list_recycle", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *pegawaiHandler) RestorePegawai(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.RestorePegawai(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-pegawai", "Pegawai berhasil di aktifkan kembali!", "success", session)

	return c.Redirect(root)
}

func (h *pegawaiHandler) DeletePegawaiSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	pegawai, err := h.service.DeletedPegawaiSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Pegawai %s berhasil di hapus. ", pegawai.Nama)
	helper.AlertMassage("msg-alert-new-pegawai", returnString, "success", session)

	return c.Redirect(root)
}

func (h *pegawaiHandler) DeletePegawai(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	pegawai, err := h.service.DeletedPegawai(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-pegawai", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Pegawai %s berhasil di hapus.", pegawai.Nama)
	helper.AlertMassage("msg-alert-restore-pegawai", returnString, "success", session)

	return c.Redirect(root)
}
