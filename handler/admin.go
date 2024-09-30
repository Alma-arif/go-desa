package handler

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type adminHandler struct {
	userService  user.Service
	sessionStore *session.Store
}

func NewAdminHandler(userService user.Service, sessionStore *session.Store) *adminHandler {
	return &adminHandler{userService, sessionStore}
}

func (h *adminHandler) ShowUserProfile(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	// ambil sessio pesan alert dan error form
	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-update-password")
		staussession = session.Get("msg-alert-update-password-status")
	}

	alert := helper.AlertString(sessionResult, staussession)
	// end

	// ambil cookies id user yang ter enkrisi
	var idUser uint
	cookisUserId := c.Cookies("sessionLog")
	idUser, err = helper.GetSessionID(cookisUserId)
	if err != nil {
		return c.Redirect("/login")
	}
	// end

	// ambil data user dari id cookie untuk di tampilkan sebagai data user yang login
	userMain, err := h.userService.GetUserByID(idUser)
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}
	// end

	return c.Render("profile/dasboard_user_profile", fiber.Map{
		"header": userMain,
		"data":   userMain,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *adminHandler) ShowAdminAll(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-new-user")
		staussession = session.Get("msg-alert-new-user-status")
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

	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	return c.Render("admin/user/dasboard_user_list", fiber.Map{
		"header": userMain,
		"data":   users,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *adminHandler) ShowAdminDetail(c *fiber.Ctx) error {
	// ambil cookies id user yang ter enkrisi
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

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	return c.Render("admin/user/dasboard_user_detail", fiber.Map{
		"header": userMain,
		"data":   user,
		"layout": "table",
	})
}

func (h *adminHandler) NewUser(c *fiber.Ctx) error {

	root := c.Get("Referer")

	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	inputNew := new(user.RegisterUserInput)

	if err := c.BodyParser(inputNew); err != nil {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	userRow, err := h.userService.RegisterUser(*inputNew)
	if err != nil {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-user", "Resgistrasi Pengguna baru berhasil.", "success", session)
	pathRoot := fmt.Sprintf("/dasboard/admin/user/profile/image/%d", userRow.ID)
	return c.Redirect(pathRoot)
}

func (h *adminHandler) UploadImageProfileView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {
		sessionResult = session.Get("msg-alert-new-user")
		staussession = session.Get("msg-alert-new-user-status")
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
		return c.Redirect("/login")
	}

	userRow, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		return c.Redirect("/dasboard/admin/user")
	}
	return c.Render("admin/user/dasboard_user_image_profile", fiber.Map{
		"header": userMain,
		"data":   userRow,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *adminHandler) UploadImageProfile(c *fiber.Ctx) error {
	root := c.Get("Referer")

	session, _ := h.sessionStore.Get(c)
	id := c.FormValue("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	file, err := c.FormFile("image_avatar")
	if err != nil {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	a, _ := helper.IsAllowedFileTypeImage(file)
	if a != true {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	fileName := helper.StringWithoutSpaces(file.Filename)
	fileNameToEncrip, err := helper.GetFileNameEnkrip(fileName)
	if err != nil {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	path := fmt.Sprintf("derektori/images_profile/%s", fileNameToEncrip)
	userRow, err := h.userService.GetUserByID(uint(idInt))
	if err != nil {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	// update image profile
	if userRow.ProfileFile != "" {
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

		pathRemoveFile := fmt.Sprintf("derektori/images_profile/%s", userRow.ProfileFile)
		err = os.Remove(pathRemoveFile)
		if err != nil {
			helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
			return c.Redirect(root)
		}

		helper.AlertMassage("msg-alert-new-user", "Image profile berhasil di perbarui!", "success", session)
		return c.Redirect("/dasboard/admin/user")
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

func (h *adminHandler) UpdateUserView(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()

	} else {
		sessionResult = session.Get("msg-alert-update-user")
		staussession = session.Get("msg-alert-update-user-status")
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
		return c.Render("error400", fiber.Map{})
	}

	userRow, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}
	return c.Render("admin/user/dasboard_user_update", fiber.Map{
		"header": userMain,
		"data":   userRow,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *adminHandler) UpdateUser(c *fiber.Ctx) error {
	root := c.Get("Referer")
	input := new(user.UpdateUserInput)
	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-update-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if err = c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if _, err = h.userService.UpdateUser(*input); err != nil {
		helper.AlertMassage("msg-alert-update-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-update-user", "Update User berhasil.", "success", session)
	return c.Redirect(root)
}

func (h *adminHandler) UpdatePasswordView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()

	} else {
		sessionResult = session.Get("msg-alert-update-password")
		staussession = session.Get("msg-alert-update-password-status")
	}

	// end from error
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
		return c.Render("error400", fiber.Map{})
	}

	userRow, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	return c.Render("admin/user/dasboard_user_password_reset", fiber.Map{
		"header": userMain,
		"data":   userRow,
		"layout": "form",
		"alert":  template.HTML(alert),
	})
}

func (h *adminHandler) ResetPassword(c *fiber.Ctx) error {
	root := c.Get("Referer")
	input := new(user.UpdatePasswordInput)

	session, err := h.sessionStore.Get(c)

	if err != nil {
		helper.AlertMassage("msg-alert-update-password", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if err = c.BodyParser(input); err != nil {
		helper.AlertMassage("msg-alert-update-password", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if _, err = h.userService.UpdatePassword(*input); err != nil {
		helper.AlertMassage("msg-alert-update-password", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-update-password", "Password User berhasil di perbarui.", "success", session)
	return c.Redirect(root)
}

func (h *adminHandler) DeleteUserSoft(c *fiber.Ctx) error {
	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {

		fmt.Println("2")
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.userService.DeleteUserSoft(uint(id))
	if err != nil {
		fmt.Println("2")
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-user", "User berhasil di hapus!", "success", session)

	return c.Redirect(root)
}

func (h *adminHandler) DeleteUserRecycle(c *fiber.Ctx) error {
	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {

		fmt.Println("2")
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.userService.DeleteUser(uint(id))
	if err != nil {
		fmt.Println("2")
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-user", "User berhasil di hapus!", "success", session)

	return c.Redirect(root)
}

func (h *adminHandler) ShowAdminAllRecycle(c *fiber.Ctx) error {

	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
	} else {

		sessionResult = session.Get("msg-alert-new-user")
		staussession = session.Get("msg-alert-new-user-status")
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

	users, err := h.userService.GetAllUsersDeleted()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	return c.Render("admin/user/dasboard_user_list_recycle", fiber.Map{
		"header": userMain,
		"data":   users,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}

func (h *adminHandler) RestoreUser(c *fiber.Ctx) error {

	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)

	if err != nil {

		fmt.Println("2")
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	_, err = h.userService.RestoreUser(uint(id))
	if err != nil {
		fmt.Println("2")
		helper.AlertMassage("msg-alert-new-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.AlertMassage("msg-alert-new-user", "User berhasil di hapus!", "success", session)

	return c.Redirect(root)
}
