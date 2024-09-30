package handler

import (
	"app-desa-kepuk/berita"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/pengumuman"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type webPengumumabHandler struct {
	service       pengumuman.Service
	beritaService berita.Service
}

func NewWebPengumumanHandler(service pengumuman.Service, beritaService berita.Service) *webPengumumabHandler {
	return &webPengumumabHandler{service, beritaService}
}

func (h *webPengumumabHandler) ShowPengumumanAll(c *fiber.Ctx) error {

	q := c.Queries()
	page, _ := strconv.Atoi(q["page"])
	if page <= 0 {
		page = 1
	}

	perPage := 9

	pengumuman, totalRows, err := h.service.GetAllPengumumanWeb(perPage, page)
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	pagination, _ := helper.GetPaginationLinks(helper.PaginationParams{
		Path:        "pengumuman",
		TotalRows:   int32(totalRows),
		PerPage:     int32(perPage),
		CurrentPage: int32(page),
	})

	pengumumanLimit, _ := h.service.GetAllPengumumanWebLimit(5)

	beritaLimit, _ := h.beritaService.GetAllBeritaWebLimit(5)

	data := map[string]interface{}{
		"main":          pengumuman,
		"pengumumanNew": pengumumanLimit,
		"beritaNew":     beritaLimit,
	}

	return c.Render("home/pengumuman", fiber.Map{
		"data":       data,
		"pagination": pagination,
	})
}

func (h *webPengumumabHandler) ShowPengumumanDetail(c *fiber.Ctx) error {
	parameter := c.Params("judul")

	judulID := helper.StrinParameterJudulID(parameter)
	berita, err := h.service.GetPengumumanByID(judulID)
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	pengumumanLimit, _ := h.service.GetAllPengumumanWebLimit(5)

	beritaLimit, _ := h.beritaService.GetAllBeritaWebLimit(5)

	data := map[string]interface{}{
		"main":          berita,
		"beritaNew":     beritaLimit,
		"pengumumanNew": pengumumanLimit,
	}

	return c.Render("home/pengumuman_detail", fiber.Map{
		"data": data,
	})
}
