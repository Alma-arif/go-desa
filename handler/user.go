package handler

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type userHandler struct {
	userService  user.Service
	sessionStore *session.Store
}

func NewUserHandler(userService user.Service, sessionStore *session.Store) *userHandler {
	return &userHandler{userService, sessionStore}
}

func (h *userHandler) ShowUserAll(c *fiber.Ctx) error {

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

	return c.Render("user/user/dasboard_user_list", fiber.Map{
		"header": userMain,
		"data":   users,
		"layout": "table",
		"alert":  template.HTML(alert),
	})
}
