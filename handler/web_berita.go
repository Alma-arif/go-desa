package handler

import (
	"app-desa-kepuk/berita"
	beritaimage "app-desa-kepuk/beritaImage"
	"app-desa-kepuk/beritakategori"
	"app-desa-kepuk/helper"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type webBeritaHandler struct {
	Service         berita.Service
	imageService    beritaimage.Service
	kategoriService beritakategori.Service
}

func NewWebBeritaHandler(Service berita.Service, imageService beritaimage.Service, kategoriService beritakategori.Service) *webBeritaHandler {
	return &webBeritaHandler{Service, imageService, kategoriService}
}

func (h *webBeritaHandler) ShowBeritaAll(c *fiber.Ctx) error {

	q := c.Queries()
	page, _ := strconv.Atoi(q["page"])
	if page <= 0 {
		page = 1
	}

	perPage := 9

	beritas, totalRows, err := h.Service.GetAllBeritaWeb(perPage, page)
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	// fmt.Println(beritas)

	pagination, _ := helper.GetPaginationLinks(helper.PaginationParams{
		Path:        "berita",
		TotalRows:   int32(totalRows),
		PerPage:     int32(perPage),
		CurrentPage: int32(page),
	})

	beritaLimit, err := h.Service.GetAllBeritaWebLimit(5)
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}
	data := map[string]interface{}{
		"main":      beritas,
		"beritaNew": beritaLimit,
	}

	return c.Render("home/berita", fiber.Map{
		"data":       data,
		"pagination": pagination,
	})
}

func (h *webBeritaHandler) ShowBeritaDetail(c *fiber.Ctx) error {
	parameter := c.Params("judul")

	judulID := helper.StrinParameterJudulID(parameter)
	berita, err := h.Service.GetBeritaWebByID(judulID)
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}

	beritaLimit, err := h.Service.GetAllBeritaWebLimit(5)
	if err != nil {
		return c.Render("error400", fiber.Map{})
	}
	data := map[string]interface{}{
		"main":      berita,
		"beritaNew": beritaLimit,
	}

	return c.Render("home/berita_detail", fiber.Map{
		"data": data,
	})
}
