package handler

import (
	"app-desa-kepuk/filesurat"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/surat"
	"app-desa-kepuk/user"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type suratHandler struct {
	service      surat.Service
	fileSurat    filesurat.Service
	userService  user.Service
	sessionStore *session.Store
	Validate     *validator.Validate
}

func NewsuratHandler(service surat.Service, fileSurat filesurat.Service, userService user.Service, sessionStore *session.Store, Validate *validator.Validate) *suratHandler {
	return &suratHandler{service, fileSurat, userService, sessionStore, Validate}
}

func (h *suratHandler) ShowSuratList(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}

	session, err := h.sessionStore.Get(c)
	defer session.Save()

	// ambil sessio pesan alert dan error form
	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-new-surat")
		staussession = session.Get("msg-alert-new-surat-status")
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

	fileSurat, _ := h.service.GetAllSurat()

	data := map[string]interface{}{
		"file": fileSurat,
	}

	return c.Render("admin/surat/dasboard_surat_list", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *suratHandler) ShowSuratListRecycle(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}

	session, err := h.sessionStore.Get(c)
	defer session.Save()

	// ambil sessio pesan alert dan error form
	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-new-surat")
		staussession = session.Get("msg-alert-new-surat-status")
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

	fileSurat, _ := h.service.GetAllSuratDeleted()

	data := map[string]interface{}{
		"file": fileSurat,
	}

	return c.Render("admin/surat/dasboard_surat_recycle", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *suratHandler) RestoreSurat(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.RestoreSurat(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-surat", "Data berhasil dipulihkan!", "success", session)

	return c.Redirect(root)
}

func (h *suratHandler) DeleteSuratSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedSuratSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-surat", "Data berhasil di hapus!", "success", session)

	return c.Redirect(root)
}

func (h *suratHandler) DeleteSurat(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeleteSurat(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-surat", "Data berhasil di hapus.", "success", session)

	return c.Redirect(root)
}

func (h *suratHandler) DownloadFileSurat(c *fiber.Ctx) error {
	root := c.Get("Referer")
	fmt.Println("Don1 : ")
	session, err := h.sessionStore.Get(c)
	if err != nil {
		fmt.Println("Don2 : ", err)

		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println("Don3 : ", err)
		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, err := h.service.GetSuratByID(uint(id))
	fmt.Println("Don4 : ", file)

	if err != nil {
		fmt.Println("Don5 : ", err)

		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fmt.Println("Don6 : end")
	fileLocation := fmt.Sprintf("./derektori/surat/file_surat/%s", file.FileLocation)
	return c.Download(fileLocation, file.FileLocation)
}

// surat Usaha
func (h *suratHandler) InputSuartUsahaView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-surat")
		staussession = session.Get("msg-alert-new-surat-status")
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

	FileSurat, _ := h.fileSurat.GetFileSuratByKodeFix(os.Getenv("Code_Surat_Usaha"))
	fmt.Println("data form :", FileSurat)
	rowSurat, _ := h.service.GetSuratByKodeSurat(FileSurat.ID)

	no := helper.NomerSurat(rowSurat.NoSurat + 1)

	data := map[string]interface{}{
		"file":     FileSurat,
		"no_surat": no,
	}

	return c.Render("admin/surat/suratusaha/dasboard_surat_usaha_satu_input", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *suratHandler) InputSuartUsaha(c *fiber.Ctx) error {

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	fmt.Println("a1")
	if err != nil {
		fmt.Println("a2")

		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputNew := new(surat.InputSuratKeteranganUsaha)
	if err := c.BodyParser(inputNew); err != nil {
		fmt.Println("a3")
		fmt.Println("error : ", err)

		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	err = h.Validate.Struct(inputNew)
	if err != nil {
		fmt.Println("a4 : ", err)

		helper.AlertMassage("msg-alert-update-file-surat", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.CreateSuratUsaha(*inputNew)
	if err != nil {
		fmt.Println("a5 :", err)

		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}
	fmt.Println("a6 end")

	helper.AlertMassage("msg-alert-new-surat", "Surat baru berhasil dibuat.", "success", session)
	// pathRoot := fmt.Sprintf("/dasboard/admin/user/profile/image/%d", surat.ID)
	return c.Redirect(root)
}

func (h *suratHandler) UpdateSuartUsahaView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	var datasurat surat.InputSuratKeteranganUsaha

	session, err := h.sessionStore.Get(c)
	defer session.Save()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}
	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-Update-surat-usaha")
		staussession = session.Get("msg-alert-Update-surat-usaha-status")
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
	surat, err := h.service.GetSuratByID(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	err = json.Unmarshal(surat.Data, &datasurat)
	if err != nil {
		fmt.Println("data surat err :", err)
	}

	// fmt.Println("data surat", datasurat)
	// fmt.Println("data surat 1", surat)
	// fmt.Println("data surat 2", surat.Data)

	data := map[string]interface{}{
		"nosurat":   helper.NomerSurat(surat.NoSurat),
		"surat":     datasurat,
		"Suartfull": surat,
	}

	return c.Render("admin/surat/suratusaha/dasboard_surat_usaha_satu_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *suratHandler) UpdateSuratUsaha(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)
	fmt.Println("a1")
	if err != nil {
		fmt.Println("a2 :", err)

		helper.AlertMassage("msg-alert-update-surat-usaha", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputNew := new(surat.UpdateSuratKeteranganUsaha)
	if err := c.BodyParser(inputNew); err != nil {
		fmt.Println("a3 :", err)
		helper.AlertMassage("msg-alert-update-surat-usaha", err.Error(), "error", session)
		return c.Redirect(root)
	}

	err = h.Validate.Struct(inputNew)
	if err != nil {
		fmt.Println("a4 : ", err)
		helper.AlertMassage("msg-alert-update-surat-usaha", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.UpdateSuratUsaha(*inputNew)
	if err != nil {
		fmt.Println("a5 :", err)
		helper.AlertMassage("msg-alert-update-surat-usaha", err.Error(), "error", session)
		return c.Redirect(root)
	}
	fmt.Println("a6 end")

	helper.AlertMassage("msg-alert-update-surat-usaha", "Surat baru berhasil Update.", "success", session)
	return c.Redirect(root)
}

// surat Usaha end

// surat Kematian
func (h *suratHandler) InputSuartKematianView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-surat-kematian")
		staussession = session.Get("msg-alert-new-surat-kematian-status")
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

	FileSurat, _ := h.fileSurat.GetFileSuratByKodeFix(os.Getenv("Code_Surat_Kematian"))
	rowSurat, _ := h.service.GetSuratByKodeSurat(FileSurat.ID)
	no := helper.NomerSurat(rowSurat.NoSurat + 1)
	fmt.Println(FileSurat)
	data := map[string]interface{}{
		"file":     FileSurat,
		"no_surat": no,
	}

	return c.Render("admin/surat/suratkematian/dasboard_surat_kematian_satu_input", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *suratHandler) InputSuartKematian(c *fiber.Ctx) error {

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-new-surat-kematian", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputNew := new(surat.InputSuratKeteranganMeninggal)

	if err := c.BodyParser(inputNew); err != nil {
		helper.AlertMassage("msg-alert-new-surat-kematian", err.Error(), "error", session)
		return c.Redirect(root)
	}

	err = h.Validate.Struct(inputNew)
	if err != nil {
		helper.AlertMassage("msg-alert-update-file-surat-kematian", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.CreateSuratKeteranganKematian(*inputNew)
	if err != nil {
		helper.AlertMassage("msg-alert-new-surat-kematian", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-surat-kematian", "Surat baru berhasil dibuat.", "success", session)
	// pathRoot := fmt.Sprintf("/dasboard/admin/user/profile/image/%d", surat.ID)
	return c.Redirect(root)
}

func (h *suratHandler) UpdateSuartKematianView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	var datasurat surat.InputSuratKeteranganMeninggal

	session, err := h.sessionStore.Get(c)
	defer session.Save()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}
	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-Update-surat-kematian")
		staussession = session.Get("msg-alert-Update-surat-kematian-status")
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
	surat, err := h.service.GetSuratByID(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	err = json.Unmarshal(surat.Data, &datasurat)
	if err != nil {
		fmt.Println("data surat err :", err)
	}
	// fmt.Println("data surat", datasurat)
	// fmt.Println("data surat 1", surat)
	// fmt.Println("data surat 2", surat.Data)

	data := map[string]interface{}{
		"nosurat":   helper.NomerSurat(surat.NoSurat),
		"surat":     datasurat,
		"Suartfull": surat,
	}

	return c.Render("admin/surat/suratkematian/dasboard_surat_kematian_satu_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *suratHandler) UpdateSuratKematian(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)
	fmt.Println("a1")
	if err != nil {
		fmt.Println("a2 :", err)

		helper.AlertMassage("msg-alert-update-surat-kematian", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputNew := new(surat.UpdateSuratKeteranganMeninggal)
	if err := c.BodyParser(inputNew); err != nil {
		fmt.Println("a3 :", err)
		helper.AlertMassage("msg-alert-update-surat-kematian", err.Error(), "error", session)
		return c.Redirect(root)
	}

	err = h.Validate.Struct(inputNew)
	if err != nil {
		fmt.Println("a4 : ", err)
		helper.AlertMassage("msg-alert-update-surat-kematian", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.UpdateSurateteranganKematian(*inputNew)
	if err != nil {
		fmt.Println("a5 :", err)
		helper.AlertMassage("msg-alert-update-surat-kematian", err.Error(), "error", session)
		return c.Redirect(root)
	}
	fmt.Println("a6 end")

	helper.AlertMassage("msg-alert-update-surat-kematian", "Surat baru berhasil Update.", "success", session)
	return c.Redirect(root)
}

// surat Kematian end

// surat nikah N1

func (h *suratHandler) InputSuratKeteranganNikahNSatuView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-surat-nikah-satu")
		staussession = session.Get("msg-alert-new-surat-nikah-satu-status")
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

	FileSurat, _ := h.fileSurat.GetFileSuratByKodeFix(os.Getenv("Code_Surat_Nikah_N_Satu"))
	rowSurat, _ := h.service.GetSuratByKodeSurat(FileSurat.ID)
	no := helper.NomerSurat(rowSurat.NoSurat + 1)
	fmt.Println("data a :", FileSurat)
	data := map[string]interface{}{
		"file":     FileSurat,
		"no_surat": no,
	}

	return c.Render("admin/surat/suratnikah/dasboard_surat_nikah_satu_input", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *suratHandler) InputSuratKeteranganNikahNSatu(c *fiber.Ctx) error {

	root := c.Get("Referer")
	fmt.Println("data b1 :")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-surat-nikah-satu", err.Error(), "error", session)
		fmt.Println("data b2 :", err)

		return c.Redirect(root)
	}

	inputNew := new(surat.InputSuratPengantarNikahNSatu)

	if err := c.BodyParser(inputNew); err != nil {
		helper.AlertMassage("msg-alert-new-surat-nikah-satu", err.Error(), "error", session)
		fmt.Println("data b3 :", err)

		return c.Redirect(root)
	}

	err = h.Validate.Struct(inputNew)
	if err != nil {
		fmt.Println("data b4 :", err)

		helper.AlertMassage("msg-alert-update-file-surat-nikah-satu", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	// inputNew.JenisKelaminAyahPemohon = "Laki-laki"
	// inputNew.JenisKelaminIbuPemohon = "Perempuan"

	_, err = h.service.CreateSuratKeteranganNikahNSatu(*inputNew)
	if err != nil {
		fmt.Println("data b5 :", err)

		helper.AlertMassage("msg-alert-new-surat-nikah-satu", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-surat-nikah-satu", "Surat baru berhasil dibuat.", "success", session)
	// pathRoot := fmt.Sprintf("/dasboard/admin/user/profile/image/%d", surat.ID)
	return c.Redirect(root)
}

func (h *suratHandler) UpdateSuratKeteranganNikahNSatuView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	var datasurat surat.InputSuratPengantarNikahNSatu

	session, err := h.sessionStore.Get(c)
	defer session.Save()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}
	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-Update-surat-nikah-n-satu")
		staussession = session.Get("msg-alert-Update-surat-nikah-n-satu")
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
	surat, err := h.service.GetSuratByID(uint(id))
	if err != nil {
		return c.Redirect("/login")
	}

	err = json.Unmarshal(surat.Data, &datasurat)
	if err != nil {
		fmt.Println("data surat err :", err)
	}
	// fmt.Println("data surat", datasurat)
	// fmt.Println("data surat 1", surat)
	// fmt.Println("data surat 2", surat.Data)

	data := map[string]interface{}{
		"nosurat":   helper.NomerSurat(surat.NoSurat),
		"surat":     datasurat,
		"Suartfull": surat,
	}

	return c.Render("admin/surat/suratnikah/dasboard_surat_nikah_satu_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *suratHandler) UpdateSuratKeteranganNikahNSatu(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)
	fmt.Println("a1")
	if err != nil {
		fmt.Println("a2 :", err)

		helper.AlertMassage("msg-alert-Update-surat-nikah-n-satu", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputNew := new(surat.UpdateSuratPengantarNikahNSatu)
	if err := c.BodyParser(inputNew); err != nil {
		fmt.Println("a3 :", err)
		helper.AlertMassage("msg-alert-Update-surat-nikah-n-satu", err.Error(), "error", session)
		return c.Redirect(root)
	}

	err = h.Validate.Struct(inputNew)
	if err != nil {
		fmt.Println("a4 : ", err)
		helper.AlertMassage("msg-alert-Update-surat-nikah-n-satu", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.UpdateSuratKeteranganNikahNSatu(*inputNew)
	if err != nil {
		fmt.Println("a5 :", err)
		helper.AlertMassage("msg-alert-Update-surat-nikah-n-satu", err.Error(), "error", session)
		return c.Redirect(root)
	}
	fmt.Println("a6 end")

	helper.AlertMassage("msg-alert-Update-surat-nikah-n-satu", "Surat baru berhasil Update.", "success", session)
	return c.Redirect(root)
}

// surat nikah N1 end

// surat persetujuan pengatin N4

func (h *suratHandler) InputSuratKeteranganNikahNEmpatView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-surat-nikah-empat")
		staussession = session.Get("msg-alert-new-surat-nikah-empat-status")
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

	FileSurat, _ := h.fileSurat.GetFileSuratByKodeFix(os.Getenv("Code_Surat_Nikah_N_Empat"))
	rowSurat, _ := h.service.GetSuratByKodeSurat(FileSurat.ID)
	no := helper.NomerSurat(rowSurat.NoSurat + 1)

	fmt.Println("data a : ", FileSurat)
	data := map[string]interface{}{
		"file":     FileSurat,
		"no_surat": no,
	}

	return c.Render("admin/surat/suratnikah/dasboard_surat_nikah_empat_input", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *suratHandler) InputSuratKeteranganNikahNEmpat(c *fiber.Ctx) error {

	root := c.Get("Referer")
	fmt.Println("a1")
	session, err := h.sessionStore.Get(c)
	if err != nil {
		fmt.Println("a1 : ", err)
		helper.AlertMassage("msg-alert-new-surat-nikah-empat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputNew := new(surat.InputSuratPengatarNikahNEmpat)
	if err := c.BodyParser(inputNew); err != nil {
		fmt.Println("a2 : ", err)

		helper.AlertMassage("msg-alert-new-surat-nikah-empat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	err = h.Validate.Struct(inputNew)
	if err != nil {
		fmt.Println("a3 : ", err)

		helper.AlertMassage("msg-alert-update-file-surat-nikah-empat", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.CreateSuratKeteranganNikahNEmpat(*inputNew)
	if err != nil {
		fmt.Println("a4 : ", err)

		helper.AlertMassage("msg-alert-new-surat-nikah-empat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fmt.Println("a5 : end")
	helper.AlertMassage("msg-alert-new-surat-nikah-empat", "Surat baru berhasil dibuat.", "success", session)
	return c.Redirect(root)
}

func (h *suratHandler) UpdateSuratKeteranganNikahNEmpatView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	var datasurat surat.InputSuratPengatarNikahNEmpat

	session, err := h.sessionStore.Get(c)
	defer session.Save()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}
	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-Update-surat-nikah-n-empat")
		staussession = session.Get("msg-alert-Update-surat-nikah-n-empat-status")
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
	surat, err := h.service.GetSuratByID(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	err = json.Unmarshal(surat.Data, &datasurat)
	if err != nil {
		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		fmt.Println("data surat err :", err)
	}
	// fmt.Println("data surat", datasurat)
	// fmt.Println("data surat 1", surat)
	// fmt.Println("data surat 2", surat.Data)

	data := map[string]interface{}{
		"nosurat":   helper.NomerSurat(surat.NoSurat),
		"surat":     datasurat,
		"Suartfull": surat,
	}

	return c.Render("admin/surat/suratnikah/dasboard_surat_nikah_empat_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *suratHandler) UpdateSuratKeteranganNikahNEmpat(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)
	fmt.Println("a1")
	if err != nil {
		fmt.Println("a2 :", err)

		helper.AlertMassage("msg-alert-Update-surat-nikah-n-empat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputNew := new(surat.UpdateSuratPengatarNikahNEmpat)
	if err := c.BodyParser(inputNew); err != nil {
		fmt.Println("a3 :", err)
		helper.AlertMassage("msg-alert-Update-surat-nikah-n-empat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	err = h.Validate.Struct(inputNew)
	if err != nil {
		fmt.Println("a4 : ", err)
		helper.AlertMassage("msg-alert-Update-surat-nikah-n-empat", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.UpdateSuratKeteranganNikahNEmpat(*inputNew)
	if err != nil {
		fmt.Println("a5 :", err)
		helper.AlertMassage("msg-alert-Update-surat-nikah-n-empat", err.Error(), "error", session)
		return c.Redirect(root)
	}
	fmt.Println("a6 end")

	helper.AlertMassage("msg-alert-Update-surat-nikah-n-empat", "Surat baru berhasil Update.", "success", session)
	return c.Redirect(root)
}

// surat persetujuan pengatin N4

// surat izin orang tua N5

func (h *suratHandler) InputSuratKeteranganNikahNLimaView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-surat-nikah-lima")
		staussession = session.Get("msg-alert-new-surat-nikah-lima-status")
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

	FileSurat, _ := h.fileSurat.GetFileSuratByKodeFix(os.Getenv("Code_Surat_Nikah_N_Lima"))
	rowSurat, _ := h.service.GetSuratByKodeSurat(FileSurat.ID)
	no := helper.NomerSurat(rowSurat.NoSurat + 1)
	fmt.Println("data form :", FileSurat)

	data := map[string]interface{}{
		"file":     FileSurat,
		"no_surat": no,
	}

	return c.Render("admin/surat/suratnikah/dasboard_surat_nikah_lima_input", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *suratHandler) InputSuratKeteranganNikahNLima(c *fiber.Ctx) error {
	root := c.Get("Referer")

	fmt.Println("a1 :")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		fmt.Println("a2 :", err)

		helper.AlertMassage("msg-alert-new-surat-nikah-lima", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputNew := new(surat.InputSuratPengatarNikahNLima)
	if err := c.BodyParser(inputNew); err != nil {
		fmt.Println("a3 :", err)

		helper.AlertMassage("msg-alert-new-surat-nikah-lima", err.Error(), "error", session)
		return c.Redirect(root)
	}

	err = h.Validate.Struct(inputNew)
	if err != nil {
		fmt.Println("a4 :", err)

		helper.AlertMassage("msg-alert-update-file-surat-nikah-lima", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.CreateSuratKeteranganNikahNLima(*inputNew)
	if err != nil {
		fmt.Println("a5 :", err)

		helper.AlertMassage("msg-alert-new-surat-nikah-lima", err.Error(), "error", session)
		return c.Redirect(root)
	}
	fmt.Println("a6 : end")

	helper.AlertMassage("msg-alert-new-surat-nikah-lima", "Surat baru berhasil dibuat.", "success", session)
	return c.Redirect(root)
}

func (h *suratHandler) UpdateSuratKeteranganNikahNLimaView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	var datasurat surat.InputSuratPengatarNikahNEmpat

	session, err := h.sessionStore.Get(c)
	defer session.Save()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}
	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-Update-surat-nikah-n-lima")
		staussession = session.Get("msg-alert-Update-surat-nikah-n-lima-status")
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
	surat, err := h.service.GetSuratByID(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	err = json.Unmarshal(surat.Data, &datasurat)
	if err != nil {
		helper.AlertMassage("msg-alert-new-surat", err.Error(), "error", session)
		fmt.Println("data surat err :", err)
	}
	// fmt.Println("data surat", datasurat)
	// fmt.Println("data surat 1", surat)
	// fmt.Println("data surat 2", surat.Data)

	data := map[string]interface{}{
		"nosurat":   helper.NomerSurat(surat.NoSurat),
		"surat":     datasurat,
		"Suartfull": surat,
	}

	return c.Render("admin/surat/suratnikah/dasboard_surat_nikah_lima_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *suratHandler) UpdateSuratKeteranganNikahNLima(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)
	fmt.Println("a1")
	if err != nil {
		fmt.Println("a2 :", err)

		helper.AlertMassage("msg-alert-Update-surat-nikah-n-lima", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputNew := new(surat.UpdateSuratPengatarNikahNLima)
	if err := c.BodyParser(inputNew); err != nil {
		fmt.Println("a3 :", err)
		helper.AlertMassage("msg-alert-Update-surat-nikah-n-lima", err.Error(), "error", session)
		return c.Redirect(root)
	}

	err = h.Validate.Struct(inputNew)
	if err != nil {
		fmt.Println("a4 : ", err)
		helper.AlertMassage("msg-alert-Update-surat-nikah-n-lima", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.UpdateSuratKeteranganNikahNLima(*inputNew)
	if err != nil {
		fmt.Println("a5 :", err)
		helper.AlertMassage("msg-alert-Update-surat-nikah-n-lima", err.Error(), "error", session)
		return c.Redirect(root)
	}
	fmt.Println("a6 end")

	helper.AlertMassage("msg-alert-Update-surat-nikah-n-lima", "Surat baru berhasil Update.", "success", session)
	return c.Redirect(root)
}

// surat izin orang tua N5 end
