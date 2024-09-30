package handler

import (
	"app-desa-kepuk/berita"
	beritaimage "app-desa-kepuk/beritaImage"
	"app-desa-kepuk/beritakategori"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type beritaHandler struct {
	Service         berita.Service
	imageService    beritaimage.Service
	kategoriService beritakategori.Service
	userService     user.Service
	sessionStore    *session.Store
}

func NewBeritaHandler(Service berita.Service, imageService beritaimage.Service, kategoriService beritakategori.Service, userService user.Service, sessionStore *session.Store) *beritaHandler {
	return &beritaHandler{Service, imageService, kategoriService, userService, sessionStore}
}

func (h *beritaHandler) ShowBeritaList(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-berita")
		staussession = session.Get("msg-alert-new-berita-status")
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

	berita, _ := h.Service.GetAllBerita()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	kategori, _ := h.kategoriService.GetAllBeritaKategori()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"berita":   berita,
		"kategori": kategori,
	}

	return c.Render("admin/berita/dasboard_berita_list", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *beritaHandler) ShowBeritaDetail(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-detail-berita")
		staussession = session.Get("msg-alert-detail-berita-status")
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
		return c.Redirect("/dasboard/berita/")
	}

	berita, err := h.Service.GetBeritaByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	image, _ := h.imageService.GetBeritaImageByBeritaID(berita.ID)

	data := map[string]interface{}{
		"berita": berita,
		"image":  image,
	}

	return c.Render("admin/berita/dasboard_berita_detail", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

// menambah berita baru
func (h *beritaHandler) NewBerita(c *fiber.Ctx) error {

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	cookisUserId := c.Cookies("sessionLog")
	idUser, err := helper.GetSessionID(cookisUserId)
	if err != nil {
		return c.Redirect("/login")
	}

	input := new(berita.BeritaInput)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}
	if _, err := h.Service.CreateBerita(*input, idUser); err != nil {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-berita", "Berita baru berhasil di buat!.", "success", session)
	return c.Redirect(root)
}

// merubah berita
func (h *beritaHandler) UpdateBeritaView(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-update-berita")
		staussession = session.Get("msg-alert-update-berita-status")
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
		return c.Redirect("/dasboard/admin/berita")
	}

	berita, err := h.Service.GetBeritaByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	kategori, _ := h.kategoriService.GetAllBeritaKategori()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"berita":   berita,
		"kategori": kategori,
	}
	return c.Render("admin/berita/dasboard_berita_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *beritaHandler) UpdateBerita(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-berita", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	input := new(berita.BeritaUpdate)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	beritaRow, err := h.Service.UpdateBerita(*input)
	if err != nil {
		helper.AlertMassage("msg-alert-update-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Berita %s berhasil di ubah!", beritaRow.Judul)
	helper.AlertMassage("msg-alert-update-berita", returnString, "success", session)

	return c.Redirect(root)
}

func (h *beritaHandler) ShowBeritaListRecycle(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-restore-berita")
		staussession = session.Get("msg-alert-restore-berita-status")
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

	berita, _ := h.Service.GetAllBeritaDeleted()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"berita": berita,
	}

	return c.Render("admin/berita/dasboard_berita_list_recycle", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *beritaHandler) RestoreBerita(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.Service.RestoreBerita(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-berita", "Berita berhasil di hapus!", "success", session)

	return c.Redirect(root)
}

func (h *beritaHandler) DeleteBeritaSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	berita, err := h.Service.DeletedBeritaSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Berita %s berhasil di hapus. ", berita.Judul)
	helper.AlertMassage("msg-alert-new-berita", returnString, "success", session)

	return c.Redirect(root)
}

func (h *beritaHandler) DeleteBerita(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	berita, err := h.Service.DeletedBerita(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-berita", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Berita %s berhasil di hapus.", berita.Judul)
	helper.AlertMassage("msg-alert-restore-berita", returnString, "success", session)

	return c.Redirect(root)
}
