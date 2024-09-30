package handler

import (
	"app-desa-kepuk/filesurat"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"errors"
	"fmt"
	"html/template"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type fileSuratHandler struct {
	service      filesurat.Service //1.file
	userService  user.Service      //3.user
	sessionStore *session.Store
	Validate     *validator.Validate
}

func NewFileSuratHandler(service filesurat.Service, userService user.Service, sessionStore *session.Store, Validate *validator.Validate) *fileSuratHandler {
	return &fileSuratHandler{service, userService, sessionStore, Validate}
}

func (h *fileSuratHandler) ShowFileSuratAll(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}

	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-new-file-surat")
		staussession = session.Get("msg-alert-new-file-surat-status")
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

	fileSurat, _ := h.service.GetAllFileSurat()

	data := map[string]interface{}{
		"file": fileSurat,
	}

	return c.Render("admin/filesurat/dasboard_file_list", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *fileSuratHandler) NewFileSurat(c *fiber.Ctx) error {
	root := c.Get("Referer")
	var result filesurat.InputFileSurat

	session, err := h.sessionStore.Get(c)
	if err != nil {
		fmt.Println("a1")
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(filesurat.InputFileSurat)
	if err := c.BodyParser(input); err != nil {
		fmt.Println("a2")
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	err = h.Validate.Struct(input)
	if err != nil {
		fmt.Println("a3")
		helper.AlertMassage("msg-alert-update-file-surat", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	// File main
	fileMain, err := c.FormFile("file-main")
	if err != nil {
		fmt.Println("a4")
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}
	if fileMain == nil {
		fmt.Println("a5")
		helper.AlertMassage("msg-alert-new-file-surat", errors.New("File gagal di upload!").Error(), "error", session)
		return c.Redirect(root)
	}

	fileSizeMain := float64(fileMain.Size) / 1024 / 1024
	if fileSizeMain > float64(15) {
		fmt.Println("a6")
		helper.AlertMassage("msg-alert-new-file-surat", "Ukuran File Melebih yang sistem tentukan maksimal 15MB", "error", session)
		return c.Redirect(root)
	}

	fileTypeMain, _ := helper.IsAllowedFileType(fileMain)
	if fileTypeMain != true {
		fmt.Println("a7")
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileNameMain := helper.StringWithoutSpaces(fileMain.Filename)
	fileNameMainToEnkripsi, err := helper.GetFileNameEnkrip(fileNameMain)
	if err != nil {
		fmt.Println("a8")
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	path := fmt.Sprintf("derektori/surat/template/%s", fileNameMainToEnkripsi)
	err = c.SaveFile(fileMain, path)
	if err != nil {
		fmt.Println("a9")
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileSecon, _ := c.FormFile("file")
	// File main End

	// Secon main
	var fileSuratNameSecon string
	var namafilesecone string
	if fileSecon != nil {
		fmt.Println("a11")
		fileSizeSecon := float64(fileSecon.Size) / 1024 / 1024
		if fileSizeSecon > float64(15) {
			fmt.Println("a12")
			helper.AlertMassage("msg-alert-new-file-surat", "Ukuran File Melebih yang sistem tentukan maksimal 15MB", "error", session)
			return c.Redirect(root)
		}

		fileTypeSecon, _ := helper.IsAllowedFileType(fileSecon)
		if fileTypeSecon != true {
			fmt.Println("a13")
			helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileNameSecon := helper.StringWithoutSpaces(fileSecon.Filename)
		fileNameSeconToEnkripsi, err := helper.GetFileNameEnkrip(fileNameSecon)
		if err != nil {
			fmt.Println("a14")
			helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
			return c.Redirect(root)
		}

		path := fmt.Sprintf("derektori/surat/template/%s", fileNameSeconToEnkripsi)
		err = c.SaveFile(fileSecon, path)
		if err != nil {
			fmt.Println("a15")
			helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
			return c.Redirect(root)
		}

		fileSuratNameSecon = fileNameSeconToEnkripsi
		namafilesecone = fileNameSecon
	}
	// Secon main End

	result.KodeSurat = input.KodeSurat
	result.Nama = input.Nama
	result.FileMain = fileNameMainToEnkripsi
	result.NamaFileMain = fileNameMain
	result.File = fileSuratNameSecon
	result.NamaFile = namafilesecone

	_, err = h.service.CreateFileSurat(result)
	if err != nil {
		fmt.Println("a16")
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-file-surat", "Data berhasil di upload!", "success", session)
	return c.Redirect(root)
}

func (h *fileSuratHandler) UpdateFileSuratView(c *fiber.Ctx) error {
	root := c.Get("Referer")

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)

	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-update-file-surat")
		staussession = session.Get("msg-alert-update-file-surat-status")
	}

	alert := helper.AlertString(sessionResult, staussession)

	var idUser uint
	cookisUserId := c.Cookies("sessionLog")
	idUser, err = helper.GetSessionID(cookisUserId)
	if err != nil {
		return c.Redirect("/login")
	}

	userMain, err := h.userService.GetUserByID(idUser)
	if err != nil {
		return c.Redirect("/login")
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Redirect(root)
	}

	file, err := h.service.GetFileSuratByID(uint(id))
	if err != nil {
		return c.Redirect(root)
	}

	data := map[string]interface{}{
		"file": file,
	}

	return c.Render("admin/filesurat/dasboard_file_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *fileSuratHandler) UpdateFileSurat(c *fiber.Ctx) error {
	var inputFile filesurat.UpdateFileSurat
	// var f string
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(filesurat.UpdateFileSurat)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	err = h.Validate.Struct(input)
	if err != nil {
		helper.AlertMassage("msg-alert-update-file-surat", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	inputFile.ID = input.ID
	inputFile.KodeSurat = input.KodeSurat
	inputFile.Nama = input.Nama

	_, err = h.service.UpdateFileSurat(inputFile)
	if err != nil {
		helper.AlertMassage("msg-alert-update-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-update-file-surat", "Update berhasil!", "success", session)
	return c.Redirect(root)
}

func (h *fileSuratHandler) UpdateFileSuratSeconView(c *fiber.Ctx) error {
	root := c.Get("Referer")
	fmt.Println("a1")
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)

	if err != nil {
		sessionResult = err.Error()
		fmt.Println("a2 : ", err.Error())
	} else {
		sessionResult = session.Get("msg-alert-new-file-surat-file")
		staussession = session.Get("msg-alert-new-file-surat-file-status")
	}

	alert := helper.AlertString(sessionResult, staussession)

	var idUser uint
	cookisUserId := c.Cookies("sessionLog")
	idUser, err = helper.GetSessionID(cookisUserId)
	if err != nil {
		fmt.Println("a3 : ", err.Error())

		return c.Redirect("/login")
	}

	userMain, err := h.userService.GetUserByID(idUser)
	if err != nil {
		fmt.Println("a4 : ", err.Error())

		return c.Redirect("/login")
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println("a5 : ", err.Error())

		return c.Redirect(root)
	}

	file, err := h.service.GetFileSuratByID(uint(id))
	if err != nil {
		fmt.Println("a6 : ", err.Error())

		return c.Redirect(root)
	}

	data := map[string]interface{}{
		"file": file,
	}
	fmt.Println("a7 : ")

	return c.Render("admin/filesurat/dasboard_file_update_file", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *fileSuratHandler) UpdateFileSuratSeconeFile(c *fiber.Ctx) error {
	var inputFile filesurat.UpdateFileSuratSecone

	root := c.Get("Referer")
	fmt.Println("b1 : ")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		fmt.Println("b2 : ", err.Error())

		helper.AlertMassage("msg-alert-update-file-surat-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(filesurat.UpdateFileSuratSecone)
	if err := c.BodyParser(input); err != nil {
		fmt.Println("b3 : ", err.Error())

		helper.AlertMassage("msg-alert-update-file-surat-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	err = h.Validate.Struct(input)
	if err != nil {
		fmt.Println("b4 : ", err.Error())

		helper.AlertMassage("msg-alert-update-file-surat", errors.New("Pastikan Form tidak Kosong, dan Terisi dengan benar!").Error(), "error", session)
		return c.Redirect(root)
	}

	result, err := h.service.GetFileSuratByID(input.ID)
	if err != nil {
		fmt.Println("b5 : ", err.Error())

		helper.AlertMassage("msg-alert-update-file-surat-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if int(result.ID) <= 0 {
		fmt.Println("b6 : ", err.Error())

		helper.AlertMassage("msg-alert-update-file-surat-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputFile.ID = input.ID

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("b7 : ", err.Error())

		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if file == nil {
		fmt.Println("b8 : ", err.Error())
		helper.AlertMassage("msg-alert-new-file-surat", errors.New("File gagal di upload!").Error(), "error", session)
		return c.Redirect(root)
	}

	fileSizeMain := float64(file.Size) / 1024 / 1024
	if fileSizeMain > float64(15) {
		fmt.Println("b9 : ", err.Error())

		helper.AlertMassage("msg-alert-new-file-surat", "Ukuran File Melebih yang sistem tentukan maksimal 15MB", "error", session)
		return c.Redirect(root)
	}

	fileTypeMain, _ := helper.IsAllowedFileType(file)
	if fileTypeMain != true {
		fmt.Println("b10 : ", err.Error())
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileName := helper.StringWithoutSpaces(file.Filename)
	fileNameToEnkripsi, err := helper.GetFileNameEnkrip(fileName)
	if err != nil {
		fmt.Println("b11 : ", err.Error())
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	path := fmt.Sprintf("derektori/surat/template/%s", fileNameToEnkripsi)
	err = c.SaveFile(file, path)
	if err != nil {
		fmt.Println("b12 : ", err.Error())
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputFile.File = fileNameToEnkripsi
	inputFile.NamaFile = fileName

	data, err := h.service.UpdateFileSuratSecone(inputFile)
	if err != nil {
		fmt.Println("b14 : ", err.Error())
		pathRemoveFile := fmt.Sprintf("derektori/surat/template/%s", data.File)
		err = os.Remove(pathRemoveFile)
		if err != nil {
			fmt.Println("b13 : ", err.Error())
			helper.AlertMassage("msg-alert-update-file-surat-file", err.Error(), "error", session)
			return c.Redirect(root)
		}

		helper.AlertMassage("msg-alert-update-file-surat-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fmt.Println("b15 : ")
	helper.AlertMassage("msg-alert-update-file-surat-file", "Update berhasil!", "success", session)
	return c.Redirect(root)
}

func (h *fileSuratHandler) ShowFileSuratListRecycle(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-restore-file-surat")
		staussession = session.Get("msg-alert-restore-file-surat-status")
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

	file, _ := h.service.GetAllFileSuratDeleted()

	data := map[string]interface{}{
		"file": file,
	}

	return c.Render("admin/filesurat/dasboard_file_list_recycle", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *fileSuratHandler) RestoreFileSurat(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.RestoreFileSurat(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-file-surat", "Data berhasil dipulihkan!", "success", session)

	return c.Redirect(root)
}

func (h *fileSuratHandler) DeleteFileSuratSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeletedFileSuartSoft(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-file-surat", "Data berhasil di hapus!", "success", session)

	return c.Redirect(root)
}

func (h *fileSuratHandler) DeleteFileSurat(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeleteFileSurat(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-file-surat", "Data berhasil di hapus.", "success", session)

	return c.Redirect(root)
}

func (h *fileSuratHandler) DownloadFileSuratMain(c *fiber.Ctx) error {

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, err := h.service.GetFileSuratByID(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileLocation := fmt.Sprintf("./derektori/surat/template/%s", file.FileMain)
	return c.Download(fileLocation, file.NamaFileMain)
}

func (h *fileSuratHandler) DownloadFileSurat(c *fiber.Ctx) error {

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, err := h.service.GetFileSuratByID(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileLocation := fmt.Sprintf("./derektori/surat/template/%s", file.File)
	return c.Download(fileLocation, file.NamaFile)
}

func (h *fileSuratHandler) RemoveFileSurat(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.service.DeleteFile(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-file-surat", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-file-surat", "Data berhasil di hapus.", "success", session)
	return c.Redirect(root)
}
