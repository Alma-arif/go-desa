package handler

import (
	"app-desa-kepuk/arsip"
	"app-desa-kepuk/file"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type dasboardHandler struct {
	userService  user.Service
	fileService  file.Service
	arsipService arsip.Service
	sessionStore *session.Store
}

func NewDasboardHandler(userService user.Service, fileService file.Service, arsipService arsip.Service, sessionStore *session.Store) *dasboardHandler {
	return &dasboardHandler{userService, fileService, arsipService, sessionStore}
}

func (h *dasboardHandler) ShowDasbordAdminView(c *fiber.Ctx) error {

	cookisUserId := c.Cookies("sessionLog")
	idUser, err := helper.GetSessionID(cookisUserId)

	if err != nil {
		return c.Redirect("/login")
	}

	userMain, err := h.userService.GetUserByID(idUser)
	if err != nil {
		return c.Redirect("/login")
	}

	user, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	file, err := h.fileService.GetFileAll()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	arsips, err := h.arsipService.GetAllArsip()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	data := map[string]interface{}{
		"file":  len(file),
		"arsip": len(arsips),
		"user":  len(user),
	}

	return c.Render("admin/dasboardAdminHome", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
	})
}

func (h *dasboardHandler) ShowDasbordUserView(c *fiber.Ctx) error {

	cookisUserId := c.Cookies("sessionLog")
	idUser, err := helper.GetSessionID(cookisUserId)
	if err != nil {
		return c.Redirect("/login")
	}

	userMain, err := h.userService.GetUserByID(idUser)
	if err != nil {
		return c.Redirect("/login")
	}

	file, err := h.fileService.GetFileAll()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	arsips, err := h.arsipService.GetAllArsip()
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	data := map[string]interface{}{
		"file":  len(file),
		"arsip": len(arsips),
	}

	return c.Render("user/dasboard_user_home", fiber.Map{
		"header": userMain,
		"data":   data,
		"layout": "table",
	})
}
