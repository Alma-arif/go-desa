package handler

import (
	"app-desa-kepuk/helper"
	"app-desa-kepuk/user"
	"fmt"
	"html/template"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type sessionHandler struct {
	userService  user.Service
	sessionStore *session.Store
}

func NewSessionHandler(userService user.Service, sessionStore *session.Store) *sessionHandler {
	return &sessionHandler{userService, sessionStore}
}

func (h *sessionHandler) LoginView(c *fiber.Ctx) error {
	var sessionResult, staussession interface{}
	session, err := h.sessionStore.Get(c)
	defer session.Save()

	if err != nil {
		sessionResult = err.Error()
		session.Destroy()
	} else {
		sessionResult = session.Get("msg-alert-login-user")
		staussession = session.Get("msg-alert-login-user-status")
	}

	alert := helper.AlertString(sessionResult, staussession)
	return c.Render("session/login", fiber.Map{
		"alert": template.HTML(alert),
	})
}

func (h *sessionHandler) Login(c *fiber.Ctx) error {
	var input user.LoginInput
	root := c.Get("Referer")
	session, err := h.sessionStore.Get(c)
	if err != nil {
		// fmt.Println()
		helper.AlertMassage("msg-alert-login-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	if err := c.BodyParser(&input); err != nil {
		helper.AlertMassage("msg-alert-login-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	rowUser, err := h.userService.Login(input)
	if err != nil {
		helper.AlertMassage("msg-alert-login-user", err.Error(), "error", session)
		return c.Redirect(root)
	}

	helper.FormDeleteError("msg-alert-login-user", session)

	idString := fmt.Sprintln(rowUser.ID)

	cookie := new(fiber.Cookie)
	cookie.Name = "sessionLog"
	cookie.Value = idString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	// cookie.HTTPOnly = true
	// cookie.Secure = true
	c.Cookie(cookie)
	if rowUser.Role == "user" {
		return c.Redirect("/dasboard/user/")
	}

	return c.Redirect("/dasboard/admin/")
}

func (h *sessionHandler) Destroy(c *fiber.Ctx) error {
	// root := c.Get("Referer")

	cookisUserId := c.Cookies("sessionLog")
	idUser, err := helper.GetSessionID(cookisUserId)
	if err != nil {
		return c.Redirect("/login")
	}

	if idUser == 0 {
		return c.Redirect("/login")

	}

	c.ClearCookie("sessionLog")
	c.ClearCookie()

	return c.Redirect("/login")
}
