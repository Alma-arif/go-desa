package handler

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/slideshow"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type imageSlideShowHandler struct {
	service      slideshow.Service
	userService  user.Service
	sessionStore *session.Store
}

func NewImageSlideShowHandler(service slideshow.Service, userService user.Service, sessionStore *session.Store) *imageSlideShowHandler {
	return &imageSlideShowHandler{service, userService, sessionStore}
}

func (h *imageSlideShowHandler) ShowimageSlideShowList(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-imageSlideShow")
		staussession = session.Get("msg-alert-new-imageSlideShow-status")
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

	informasi, _ := h.service.GetAllImageSlideShow()

	data := map[string]interface{}{
		"imageShow": informasi,
	}

	return c.Render("admin/imageshow/dasboard_image_show_list", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *imageSlideShowHandler) ShowImageSlideShowDetail(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-detail-imageSlideShow")
		staussession = session.Get("msg-alert-detail-imageSlideShow-status")
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
		return c.Redirect("/dasboard/image-show/")
	}

	image, err := h.service.GetImageSlideShowByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	// image, _ := h.imageService.GetBeritaImageByBeritaID(berita.ID)
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"imageShow": image,
	}

	return c.Render("admin/imageshow/dasboard_image_show_detail", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *imageSlideShowHandler) NewImageSlideShow(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	// cookisUserId := c.Cookies("sessionLog")
	// idUser, err := helper.GetSessionID(cookisUserId)
	// if err != nil {
	// 	return c.Redirect("/login")
	// }

	input := new(slideshow.ImageSlideShowInput)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-new-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}
	fmt.Println("1 :", input)

	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println("a1")
		helper.AlertMassage("msg-alert-new-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	a, _ := helper.IsAllowedFileTypeImage(file)
	if a != true {
		fmt.Println("a2")

		helper.AlertMassage("msg-alert-new-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileName := helper.StringWithoutSpaces(file.Filename)
	fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
	if err != nil {
		fmt.Println("a3")

		helper.AlertMassage("msg-alert-new-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	path := fmt.Sprintf("derektori/image/%s", fileNameToEncrip)

	if err = c.SaveFile(file, path); err != nil {
		fmt.Println("a4")

		helper.AlertMassage("msg-alert-new-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if _, err := h.service.CreateImageSlideShow(*input, fileNameToEncrip); err != nil {
		fmt.Println("a5")

		helper.AlertMassage("msg-alert-new-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-imageSlideShow", "Image baru berhasil di buat!.", "success", session)
	return c.Redirect(root)
}

func (h *imageSlideShowHandler) UpdateImageSlideShowView(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-update-imageSlideShow")
		staussession = session.Get("msg-alert-update-imageSlideShow-status")
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
		return c.Redirect("/dasboard/admin/image-show")
	}

	image, err := h.service.GetImageSlideShowByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	// kategori, err := h.kategoriService.GetAllBeritaKategori()
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	data := map[string]interface{}{
		"imageShow": image,
	}
	return c.Render("admin/imageshow/dasboard_image_show_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *imageSlideShowHandler) UpdateImageSlideShow(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-imageSlideShow", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	input := new(slideshow.ImageSlideShowUpdate)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	// pengumuman, err := h.service.GetPengumumanByID(input.ID)
	// if err != nil {
	// 	return c.Render("error400", fiber.Map{})
	// }

	file, err := c.FormFile("image")
	if err != nil {
		helper.AlertMassage("msg-alert-update-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	var fileNameToEncrip string
	var path string
	if file.Filename != "" {
		a, _ := helper.IsAllowedFileTypeImage(file)
		if a != true {
			helper.AlertMassage("msg-alert-update-imageSlideShow", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileName := helper.StringWithoutSpaces(file.Filename)
		fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
		if err != nil {
			helper.AlertMassage("msg-alert-update-imageSlideShow", err.Error(), "error", session)
			return c.Redirect(root)
		}

		path := fmt.Sprintf("derektori/image/%s", fileNameToEncrip)

		if err = c.SaveFile(file, path); err != nil {
			helper.AlertMassage("msg-alert-update-imageSlideShow", err.Error(), "error", session)
			return c.Redirect(root)
		}

	}

	imageRow, err := h.service.UpdateImageSlideShow(*input, fileNameToEncrip)
	if err != nil {
		os.Remove(path)
		helper.AlertMassage("msg-alert-update-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Image %s berhasil di ubah!", imageRow.Judul)
	helper.AlertMassage("msg-alert-update-imageSlideShow", returnString, "success", session)

	return c.Redirect(root)
}

func (h *imageSlideShowHandler) ShowImageSlideShowListRecycle(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-restore-imageSlideShow")
		staussession = session.Get("msg-alert-restore-imageSlideShow-status")
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

	image, _ := h.service.GetAllImageSlideShowDeletedAt()

	data := map[string]interface{}{
		"imageShow": image,
	}

	return c.Render("admin/imageshow/dasboard_image_show_list_recycle", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *imageSlideShowHandler) RestoreImageSlideShow(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.RestoreImageSlideShow(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-imageSlideShow", "Image berhasil di hapus!", "success", session)
	return c.Redirect(root)
}

func (h *imageSlideShowHandler) DeleteImageSlideShowSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	image, err := h.service.DeletedImageSlideShowSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Image %s berhasil di hapus. ", image.Judul)
	helper.AlertMassage("msg-alert-new-imageSlideShow", returnString, "success", session)

	return c.Redirect(root)
}

func (h *imageSlideShowHandler) DeleteImageSlideShow(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	image, err := h.service.DeletedImageSlideShow(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-imageSlideShow", err.Error(), "error", session)
		return c.Redirect(root)
	}

	returnString := fmt.Sprintf("Image %s berhasil di hapus.", image.Judul)
	helper.AlertMassage("msg-alert-restore-imageSlideShow", returnString, "success", session)

	return c.Redirect(root)
}
