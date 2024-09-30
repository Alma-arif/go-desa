package handler

import (
	"app-desa-kepuk/berita"
	"app-desa-kepuk/demografi"
	"app-desa-kepuk/pegawai"
	"app-desa-kepuk/pengumuman"
	"app-desa-kepuk/profiledesa"
	"app-desa-kepuk/sejarah"
	"app-desa-kepuk/slideshow"
	"app-desa-kepuk/visimisidesa"

	"github.com/gofiber/fiber/v2"
)

type WebHomeHandler struct {
	service            slideshow.Service
	pengumumanService  pengumuman.Service
	beritaService      berita.Service
	pegawaiService     pegawai.Service
	profileDesaService profiledesa.Service
	visiMisiService    visimisidesa.Service
	demografiService   demografi.Service
	sejarahService     sejarah.Service
}

func NewWebHomeHandler(service slideshow.Service, pengumumanService pengumuman.Service, beritaService berita.Service, pegawaiService pegawai.Service, profileDesaService profiledesa.Service, visiMisiService visimisidesa.Service, demografiService demografi.Service, sejarahService sejarah.Service) *WebHomeHandler {
	return &WebHomeHandler{service, pengumumanService, beritaService, pegawaiService, profileDesaService, visiMisiService, demografiService, sejarahService}
}

func (h *WebHomeHandler) ShowWebHome(c *fiber.Ctx) error {

	imageSlide, _ := h.service.GetAllImageSlideShowWeb()
	pengumumanLimit, _ := h.pengumumanService.GetAllPengumumanWebLimit(5)
	beritaLimit, _ := h.beritaService.GetAllBeritaWebLimit(5)
	pegawai, _ := h.pegawaiService.GetAllPegawai()

	data := map[string]interface{}{
		"SlideShow":     imageSlide,
		"pengumumanNew": pengumumanLimit,
		"beritaNew":     beritaLimit,
		"pegawai":       pegawai,
	}

	return c.Render("home/home", fiber.Map{
		"data": data,
	})
}

func (h *WebHomeHandler) ProfileDesaWeb(c *fiber.Ctx) error {

	profileDesa, _ := h.profileDesaService.GetAllProfileDesaWeb()
	data := map[string]interface{}{
		"profiledesa": profileDesa,
	}

	return c.Render("home/profile_desa", fiber.Map{
		"data": data,
	})
}

func (h *WebHomeHandler) VisiMisiWeb(c *fiber.Ctx) error {

	visiMisi, _ := h.visiMisiService.GetAllVisiMisiWeb()

	data := map[string]interface{}{
		"visimisi": visiMisi,
	}

	return c.Render("home/visi_misi", fiber.Map{
		"data": data,
	})
}

func (h *WebHomeHandler) DemografiWeb(c *fiber.Ctx) error {

	demografi, _ := h.demografiService.GetAllDemografiWeb()

	data := map[string]interface{}{
		"demografi": demografi,
	}

	return c.Render("home/demografi", fiber.Map{
		"data": data,
	})
}

func (h *WebHomeHandler) SejarahfiWeb(c *fiber.Ctx) error {

	sejarah, _ := h.sejarahService.GetAllSejarahWeb()

	data := map[string]interface{}{
		"sejarah": sejarah,
	}

	return c.Render("home/sejarah", fiber.Map{
		"data": data,
	})
}
