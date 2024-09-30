package handler

import (
	"app-desa-kepuk/arsip"
	"app-desa-kepuk/file"
	"app-desa-kepuk/filedetail"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"errors"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type fileHandler struct {
	fileService       file.Service       //1.file
	arsipService      arsip.Service      //2.arsip
	userService       user.Service       //3.user
	fileDetailService filedetail.Service //4.fileDetail
	sessionStore      *session.Store
}

func NewFileHandler(fileService file.Service, arsipService arsip.Service, userService user.Service, fileDetailService filedetail.Service, sessionStore *session.Store) *fileHandler {
	return &fileHandler{fileService, arsipService, userService, fileDetailService, sessionStore}
}

func (h *fileHandler) ShowFileAll(c *fiber.Ctx) error {
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
		"file":  file,
		"arsip": arsips,
	}

	return c.Render("admin/file/dasboard_file_list", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *fileHandler) ShowFileDetail(c *fiber.Ctx) error {
	var idUser uint
	var sessionResult, staussession interface{}

	cookisUserId := c.Cookies("sessionLog")

	idUser, err := helper.GetSessionID(cookisUserId)
	if err != nil {
		return c.Redirect("/login")
	}

	session, err := h.sessionStore.Get(c)
	defer session.Save()

	// ambil sessio pesan alert dan error form
	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-detail-file")
		staussession = session.Get("msg-alert-detail-file-status")
	}

	alert := helper.AlertString(sessionResult, staussession)

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

	return c.Render("admin/file/dasboard_file_detail", fiber.Map{
		"header": userMain,
		"data":   rowFile,
		"layout": "table",
		"error":  false,
		"alert":  template.HTML(alert),
	})
}

func (h *fileHandler) NewFile(c *fiber.Ctx) error {
	root := c.Get("Referer")
	var result file.CreateFileInput

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(file.CreateFileInput)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, err := c.FormFile("file-dokumen")

	da := float64(file.Size) / 1024 / 1024
	if da > float64(15) {
		helper.AlertMassage("msg-alert-new-file", "Ukuran File Melebih yang sistem tentukan maksimal 15MB", "error", session)
		return c.Redirect(root)
	}

	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if file == nil {
		helper.AlertMassage("msg-alert-new-file", errors.New("File gagal di upload!").Error(), "error", session)
		return c.Redirect(root)
	}

	a, _ := helper.IsAllowedFileType(file)
	if a != true {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}
	var fileNamestringdokumen string
	var filestringtow string
	if input.NamaFile == "" {
		fileNamestringdokumen = file.Filename
		filestringtow = file.Filename
	} else {

		extension := filepath.Ext(file.Filename)
		fileString := fmt.Sprintf("%s%s", input.NamaFile, extension)

		filestringtow = fileString
		fileNamestringdokumen = input.NamaFile
	}

	fileName := helper.StringWithoutSpaces(filestringtow)
	fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	path := fmt.Sprintf("derektori/file/%s", fileNameToEncrip)
	err = c.SaveFile(file, path)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	result.NamaFile = fileNamestringdokumen
	result.Status = 0
	result.FileSize = float64(file.Size) / 1024 / 1024
	result.FileLocation = fmt.Sprintf("%s", fileNameToEncrip)
	result.ArsipID = input.ArsipID
	result.DeskripsiFile = input.DeskripsiFile

	_, err = h.fileService.CreateFile(result, fileName)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-file", "File berhasil di upload!", "success", session)
	return c.Redirect(root)
}

func (h *fileHandler) NewFiletest(c *fiber.Ctx) error {
	root := c.Get("Referer")
	var result file.CreateFileInput

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(file.CreateFileInput)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, err := c.FormFile("file-dokumen")

	da := float64(file.Size) / 1024 / 1024
	if da > float64(15) {
		helper.AlertMassage("msg-alert-new-file", "Ukuran File Melebih yang sistem tentukan maksimal 15MB", "error", session)
		return c.Redirect(root)
	}

	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if file == nil {
		helper.AlertMassage("msg-alert-new-file", errors.New("File gagal di upload!").Error(), "error", session)
		return c.Redirect(root)
	}

	//Cek type file pdf
	a, _ := helper.IsAllowedFileType(file)
	if a != true {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}
	var fileNamestringdokumen string
	var filestringtow string
	if input.NamaFile == "" {
		fileNamestringdokumen = file.Filename
		filestringtow = file.Filename
	} else {

		extension := filepath.Ext(file.Filename)
		fileString := fmt.Sprintf("%s%s", input.NamaFile, extension)

		filestringtow = fileString
		fileNamestringdokumen = input.NamaFile
	}

	fileName := helper.StringWithoutSpaces(filestringtow)
	fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	path := fmt.Sprintf("derektori/file/%s", fileNameToEncrip)
	err = c.SaveFile(file, path)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	result.NamaFile = fileNamestringdokumen
	result.Status = 0
	result.FileSize = float64(file.Size) / 1024 / 1024
	result.FileLocation = fmt.Sprintf("%s", fileNameToEncrip)
	result.ArsipID = input.ArsipID
	result.DeskripsiFile = input.DeskripsiFile

	_, err = h.fileService.CreateFile(result, fileName)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-file", "File berhasil di upload!", "success", session)
	return c.Redirect(root)
}

func (h *fileHandler) UpdateFileView(c *fiber.Ctx) error {
	root := c.Get("Referer")

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	// defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-update-file")
		staussession = session.Get("msg-alert-update-file-status")
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

	file, err := h.fileDetailService.GetFileDetailByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{
			"header": userMain,
		})
	}

	if file.FileStatus == 1 {
		helper.AlertMassage("msg-alert-new-file", "File masih terkunci!", "error", session)
		return c.Redirect("/dasboard/admin/dokumen")
	}

	arsip, err := h.arsipService.GetAllArsip()
	if err != nil {
		return c.Render("error400", fiber.Map{
			"header": userMain,
		})

	}
	data := map[string]interface{}{
		"file":  file,
		"arsip": arsip,
	}

	return c.Render("admin/file/dasboard_file_update", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *fileHandler) UpdateFile(c *fiber.Ctx) error {
	var inputFile file.UpdateFileInput
	// var f string
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-update-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(file.UpdateFileInput)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputFile.ID = input.ID
	inputFile.ArsipID = input.ArsipID
	inputFile.NamaFile = input.NamaFile
	inputFile.DeskripsiFile = input.DeskripsiFile

	_, err = h.fileService.UpdateFile(inputFile)
	if err != nil {
		helper.AlertMassage("msg-alert-update-file", err.Error(), "error", session)

		return c.Redirect(root)
	}
	helper.AlertMassage("msg-alert-update-file", "Update Dokumen berhasil!", "success", session)

	return c.Redirect(root)
}

func (h *fileHandler) UpdateFileArsip(c *fiber.Ctx) error {
	var inputFile file.UpdateFileInput
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-detail-arsip", err.Error(), "error", session)
		return c.Redirect("/login")
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Redirect(root)
	}

	file, err := h.fileService.GetFileByID(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-detail-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputFile.ID = file.ID
	inputFile.ArsipID = 0
	inputFile.DeskripsiFile = file.Deskripsi
	inputFile.FileSize = file.FileSize
	inputFile.NamaFile = file.NamaFile
	inputFile.FileLocation = file.FileLocation

	_, err = h.fileService.UpdateFile(inputFile)
	if err != nil {
		helper.AlertMassage("msg-alert-detail-arsip", err.Error(), "error", session)
		return c.Redirect(root)
	}
	fileSucces := fmt.Sprintf("File %s berhasil di hapus dari arsip!", file.NamaFile)
	helper.AlertMassage("msg-alert-detail-arsip", fileSucces, "success", session)
	return c.Redirect(root)
}

func (h *fileHandler) DeletetFileSoft(c *fiber.Ctx) error {
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

	fileRow, err := h.fileService.GetFileByID(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if fileRow.Status == 1 {
		helper.AlertMassage("msg-alert-new-file", "File tidak bisa Di hapus, Buka kunci file!", "error", session)
		return c.Redirect(root)
	}

	_, err = h.fileService.DeleteFileSoft(fileRow.ID)
	if err != nil {
		helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileSucces := fmt.Sprintf("File %s berhasil terhapus!", fileRow.NamaFile)
	helper.AlertMassage("msg-alert-new-file", fileSucces, "success", session)

	return c.Redirect(root)
}

func (h *fileHandler) DeletetFileRecycle(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileRow, err := h.fileDetailService.GetFileDetailByIDDeleted(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.fileService.DeleteFile(fileRow.ID)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	newFile, err := helper.GetFileNameDekrip(fileRow.FileLocation)
	if err != nil {
		helper.AlertMassage("msg-alert-restore-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileSucces := fmt.Sprintf("File %s berhasil terhapus!", newFile)
	helper.AlertMassage("msg-alert-restore-file", fileSucces, "success", session)

	return c.Redirect(root)
}

func (h *fileHandler) ShowFileAllRecycle(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}

	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-restore-file")
		staussession = session.Get("msg-alert-restore-file-status")
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

	file, err := h.fileDetailService.GetFileDetailAllDeleted()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	data := map[string]interface{}{
		"file": file,
	}

	return c.Render("admin/file/dasboard_file_list_recycle", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *fileHandler) RestoreFile(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-restore-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-restore-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.fileService.RestoreFile(uint(id))
	if err != nil {
		helper.AlertMassage("msg-alert-restore-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-restore-file", "File berhasil di Kembalikan!", "success", session)

	return c.Redirect(root)
}

func (h *fileHandler) EnkripsiEndDekripsiFile(c *fiber.Ctx) error {
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

	if file.Status == 1 {

		_, err = h.fileService.Dekripsi(file.ID)
		if err != nil {
			helper.AlertMassage("msg-alert-new-file", err.Error(), "error", session)
			return c.Redirect(root)
		}

		helper.AlertMassage("msg-alert-new-file", "Kunci File berhasil terbuka!", "success", session)
		return c.Redirect(root)
	}

	_, err = h.fileService.Enkripsi(file.ID)
	if err != nil {
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-file", "Kunci File berhasil di kunci!", "success", session)
	return c.Redirect(root)
}

func (h *fileHandler) DownloadFile(c *fiber.Ctx) error {

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

func (h *fileHandler) FileArsipIDUpdate(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-detail-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	input := new(file.UpdateFileArispInput)
	if err := c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-detail-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	arsip, err := h.arsipService.GetArsipByID(input.ArsipID)
	if err != nil {
		helper.AlertMassage("msg-alert-detail-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.fileService.UpdateFileArispID(input.FileID, arsip.ID)
	if err != nil {
		helper.AlertMassage("msg-alert-detail-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-detail-file", "File berhasil di tambahkan ke arsip!", "success", session)
	return c.Redirect(root)
}

func (h *fileHandler) FileArsipIDUpdateDelete(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)
	if err != nil {
		helper.AlertMassage("msg-alert-detail-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	_, err = h.fileService.UpdateFileArispID(uint(id), 0)
	if err != nil {
		helper.AlertMassage("msg-alert-detail-file", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-detail-file", "File berhasil di hapus dari arsip!", "success", session)
	return c.Redirect(root)
}
