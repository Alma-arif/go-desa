package main

import (
	"app-desa-kepuk/arsip"
	"app-desa-kepuk/arsipkategori"
	"app-desa-kepuk/berita"
	beritaimage "app-desa-kepuk/beritaImage"
	"app-desa-kepuk/beritakategori"
	"app-desa-kepuk/config"
	"app-desa-kepuk/demografi"
	"app-desa-kepuk/file"
	"app-desa-kepuk/filedetail"
	"app-desa-kepuk/filesurat"
	"app-desa-kepuk/handler"
	"app-desa-kepuk/helper"
	"app-desa-kepuk/pegawai"
	"app-desa-kepuk/pengumuman"
	profiledesa "app-desa-kepuk/profiledesa"
	"app-desa-kepuk/sejarah"
	"app-desa-kepuk/slideshow"
	"app-desa-kepuk/surat"
	"app-desa-kepuk/user"
	"app-desa-kepuk/visimisidesa"
	"os"

	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	// env config
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal(err.Error())
	}
	helper.GenerateHMACMD5("sandipostsidesa")
	// fmt.Println("-----------------------------------------------------")
	// helper.GenerateHMACMD5("passwordqwertysidesa")
	// fmt.Println("-----------------------------------------------------")
	// helper.GenerateHMACMD5("089343passordsidesa")
	// fmt.Println("-----------------------------------------------------")
	// helper.GenerateHMACMD5("sidesasandipassword")
	// fmt.Println("-----------------------------------------------------")
	// helper.GenerateHMACMD5("Password4532sandidesa")

	validate := validator.New() // validator
	store := session.New()      // session
	db := config.InitDB()       // database

	db.AutoMigrate(&user.User{}, &file.File{}, &arsip.ArsipDesa{}, &arsipkategori.KategoriArsip{})
	db.AutoMigrate(&berita.Berita{}, &beritakategori.KategoriBerita{}, &beritaimage.ImageBerita{})
	db.AutoMigrate(&profiledesa.ProfileDesa{}, &visimisidesa.VisiMisi{}, &slideshow.ImageSlideShow{})
	db.AutoMigrate(&pengumuman.Pengumuman{}, &pegawai.Pegawai{}, &demografi.Demografi{}, &sejarah.Sejarah{})
	db.AutoMigrate(&surat.Surat{}, &filesurat.FileSurat{})

	// repository
	userRepository := user.NewRepository(db)                     // user repository
	arsipkategoriRepository := arsipkategori.NewRepository(db)   // arsip kategori repository
	arsipRepository := arsip.NewRepository(db)                   // arsip repository
	fileRepository := file.NewRepository(db)                     // dokumen repository
	beritaRepository := berita.NewRepository(db)                 // berita repository
	beritaKategoriRepository := beritakategori.NewRepository(db) // berita kategori repository
	beritaimageRepository := beritaimage.NewRepository(db)       // berita image repository
	profileDesaReposiotry := profiledesa.NewRepository(db)       // profile desa repository
	imageSlideShowReposiotry := slideshow.NewRepository(db)      // image slide show repository
	pengumumanRepository := pengumuman.NewRepository(db)         // pengumuman repository
	pegawaiRepository := pegawai.NewRepository(db)               //pegawai repository
	visiMisiRepository := visimisidesa.NewRepository(db)         //visiMisi repository
	demografiRepository := demografi.NewRepository(db)           //demografi repository
	sejarahRepository := sejarah.NewRepository(db)               //sejarah repository
	suratRepository := surat.NewRepository(db)                   //surat repository
	fileSuratRepository := filesurat.NewRepository(db)           // file surat repository

	// service
	userService := user.NewService(userRepository, validate)                                                                        // user service
	arsipkategoriService := arsipkategori.NewService(arsipkategoriRepository, validate)                                             // arsip kategori service
	arsipService := arsip.NewService(arsipRepository, arsipkategoriRepository, fileRepository, validate)                            // arsip service
	fileService := file.NewService(fileRepository, validate)                                                                        // dokumen service
	fileDetailService := filedetail.NewService(fileRepository, arsipkategoriRepository, arsipRepository)                            // dokumen detail service
	beritaService := berita.NewService(beritaRepository, validate, userRepository, beritaKategoriRepository, beritaimageRepository) // berita service
	beritaKategoriService := beritakategori.NewService(beritaKategoriRepository, validate)                                          // berita kategori service
	beritaImageService := beritaimage.NewService(beritaimageRepository, validate)                                                   // berita image service
	profileDesaService := profiledesa.NewService(profileDesaReposiotry, validate)                                                   // profile desa service
	imageSlidShowService := slideshow.NewService(imageSlideShowReposiotry, validate, userRepository)                                // image Slide Show service
	pengumumanService := pengumuman.NewService(pengumumanRepository, validate, userRepository)                                      // pengumuman service
	pegawaiService := pegawai.NewService(pegawaiRepository, validate, userRepository)                                               //pegawai service
	visiMisiService := visimisidesa.NewService(visiMisiRepository, validate, userRepository)                                        //visiMisi service
	demografiService := demografi.NewService(demografiRepository, validate, userRepository)                                         //demografi service
	sejarahService := sejarah.NewService(sejarahRepository, validate, userRepository)                                               //sejarah service)
	suratService := surat.NewService(suratRepository, fileSuratRepository)                                                          // surat service
	fileSuratService := filesurat.NewService(fileSuratRepository)                                                                   // file surat service
	// surat := surat.InputSuratKeteranganUsaha{}
	// surat.KodeSurat = "SKU"
	// surat.NoSurat = "100"
	// surat.Nama = "alma arif"
	// surat.JenisKelamin = "laki-laki"
	// surat.Agama = "islam"
	// surat.Status = "belum menikah"
	// surat.TempatTanggalLahir = "Jepara, 12 januari 1992"
	// surat.Pekerjaan = "tukang kayu"
	// surat.Keperluan = "permohonan surat usaha"
	// surat.JenisUsaha = "mable kayu"
	// surat.Keterangan = "kososng"
	// surat.TanggalMulai = "12 Januari 2024"
	// surat.TanggalSelesai = "30 maret 2024"

	// suratService.CreateSuratUsaha(surat, "file kososng")

	// inputlogin := user.LoginInput{}
	// inputlogin.Email = "kisah"
	// inputlogin.Password = "123"

	// userRow, err := userService.Login(inputlogin)
	// if err != nil {
	// 	fmt.Println("Error : ", err.Error())
	// } else {
	// 	fmt.Println("Login Berhsil ")
	// 	fmt.Println("Nama : ", userRow.Nama)
	// 	fmt.Println("Email : ", userRow.Email)
	// }

	// handler //
	dasboardHandle := handler.NewDasboardHandler(userService, fileService, arsipService, store)
	sessionHandler := handler.NewSessionHandler(userService, store)
	adminHandler := handler.NewAdminHandler(userService, store)
	// userHandler := handler.NewUserHandler(userService, store)

	arsipkategoriHandler := handler.NewArsipKategoriHandler(arsipkategoriService, userService, store)
	arsipHandler := handler.NewArsipHandler(arsipService, arsipkategoriService, fileService, fileDetailService, userService, store)
	fileHandler := handler.NewFileHandler(fileService, arsipService, userService, fileDetailService, store)
	beritaHandler := handler.NewBeritaHandler(beritaService, beritaImageService, beritaKategoriService, userService, store)
	beritaKategoriHandler := handler.NewBeritaKategoriHandler(beritaKategoriService, userService, store)
	beritaImageHandler := handler.NewBeritaImageHandler(beritaImageService, beritaService, beritaKategoriService, userService, store)
	profileDesaHandler := handler.NewProfileDesaHandler(profileDesaService, userService, store)
	imageSlideShowHandler := handler.NewImageSlideShowHandler(imageSlidShowService, userService, store)
	pengumumanHanlder := handler.NewPengumumanHandler(pengumumanService, userService, store)
	pegawaiHandler := handler.NewPegawaiHandler(pegawaiService, userService, store)
	visiMisiHandler := handler.NewVisiMisiHandler(visiMisiService, userService, store)
	demografiHandler := handler.NewDemografiHandler(demografiService, userService, store)
	sejarahHandler := handler.NewSejarahHandler(sejarahService, userService, store)
	suratHandler := handler.NewsuratHandler(suratService, fileSuratService, userService, store, validate)
	fileSuratHandler := handler.NewFileSuratHandler(fileSuratService, userService, store, validate)
	// inputFileTest := file.CreateFileInputtest{}
	// inputFileTest.NamaFile = "SuratPengantarNikah"
	// inputFileTest.DeskripsiFile = "Surat Pengatar nikah N!"
	// inputFileTest.ArsipID = 0
	// inputFileTest.FileLocation = "Derektori/file/suratpengatarnikahNsatu.pdf"

	// user
	// arsipKategoriHandlerUser := handler.NewArsipKategoriUserHandler(arsipkategoriService, userService, store)
	arsipHandlerUser := handler.NewArsipUSerHandler(arsipService, arsipkategoriService, fileService, fileDetailService, userService, store)
	fileHandlerUser := handler.NewFileUserHandler(fileService, arsipService, userService, fileDetailService, store)

	// web
	homeWebHandler := handler.NewWebHomeHandler(imageSlidShowService, pengumumanService, beritaService, pegawaiService, profileDesaService, visiMisiService, demografiService, sejarahService)
	beritaWebHandler := handler.NewWebBeritaHandler(beritaService, beritaImageService, beritaKategoriService)
	pengumumanHandler := handler.NewWebPengumumanHandler(pengumumanService, beritaService)

	// Testing

	// User Testing
	// user.NewTestingUserService(userService)

	// root
	engine := html.New("./web/template/", ".html")

	app := fiber.New(fiber.Config{
		Views:     engine,
		BodyLimit: 50 * 1024 * 1024,
	})

	// file static css end js
	app.Static("/tem", "./web/assets/AdminLTE/")
	app.Static("/themes", "./web/assets/web/")

	// file arsip public access
	app.Static("/arsip", "./derektori/file/")

	// image profile user
	app.Static("/image", "./derektori/image/")
	app.Static("/image/profile", "./derektori/images_profile/")
	app.Static("/image/berita", "./derektori/images_berita/")

	// web
	app.Get("/", homeWebHandler.ShowWebHome)

	app.Get("/berita", beritaWebHandler.ShowBeritaAll)
	app.Get("/berita/:judul", beritaWebHandler.ShowBeritaDetail)

	app.Get("/pengumuman", pengumumanHandler.ShowPengumumanAll)
	app.Get("/pengumuman/:judul", beritaWebHandler.ShowBeritaDetail)

	app.Get("/visi-misi", homeWebHandler.VisiMisiWeb)
	app.Get("/about-us", homeWebHandler.ProfileDesaWeb)
	app.Get("/demografi", homeWebHandler.DemografiWeb)
	app.Get("/sejarah-desa", homeWebHandler.SejarahfiWeb)

	// login
	app.Get("/login", LoginMiddleware(userService), sessionHandler.LoginView)
	app.Post("/session", LoginMiddleware(userService), sessionHandler.Login)
	app.Get("/logout", sessionHandler.Destroy)

	// dasboard admin
	dasboard := app.Group("/dasboard")
	// admin
	admin := dasboard.Group("/admin")
	admin.Get("/", authRoleAdminMiddleware(userService), dasboardHandle.ShowDasbordAdminView)
	admin.Get("/profile", authRoleAdminMiddleware(userService), adminHandler.ShowUserProfile)
	// admin.Get("/profile/image/:id", authRoleAdminMiddleware(userService), adminHandler.UploadImageProfileView)
	// admin.Post("/profile/image/", authRoleAdminMiddleware(userService), adminHandler.UploadImageProfile)

	// user admin
	adminUser := admin.Group("/user")
	adminUser.Get("", authRoleAdminMiddleware(userService), adminHandler.ShowAdminAll)
	adminUser.Get("/detail/:id", authRoleAdminMiddleware(userService), adminHandler.ShowAdminDetail)
	adminUser.Get("/update/:id", authRoleAdminMiddleware(userService), adminHandler.UpdateUserView)
	adminUser.Post("/update", authRoleAdminMiddleware(userService), adminHandler.UpdateUser)
	adminUser.Get("/new-password/:id", authRoleAdminMiddleware(userService), adminHandler.UpdatePasswordView)
	adminUser.Post("/new-password", authMiddleware(userService), adminHandler.ResetPassword)
	adminUser.Post("/new", authRoleAdminMiddleware(userService), adminHandler.NewUser)
	adminUser.Get("/profile/image/:id", authRoleAdminMiddleware(userService), adminHandler.UploadImageProfileView)
	adminUser.Post("/profile/image/", authRoleAdminMiddleware(userService), adminHandler.UploadImageProfile)
	adminUser.Get("/delete/:id", authRoleAdminMiddleware(userService), adminHandler.DeleteUserSoft)
	adminUser.Get("/recycle", authRoleAdminMiddleware(userService), adminHandler.ShowAdminAllRecycle)
	adminUser.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), adminHandler.RestoreUser)
	adminUser.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), adminHandler.DeleteUserRecycle)

	// arsip admin
	adminArsip := admin.Group("/arsip")
	adminArsip.Get("", authRoleAdminMiddleware(userService), arsipHandler.ShowArsipList)
	adminArsip.Get("/detail/:id", authRoleAdminMiddleware(userService), arsipHandler.ShowArsipDetail)
	adminArsip.Post("/new", authRoleAdminMiddleware(userService), arsipHandler.NewArsip)
	adminArsip.Get("/update/:id", authRoleAdminMiddleware(userService), arsipHandler.UpdateArsipView)
	adminArsip.Post("/update", authRoleAdminMiddleware(userService), arsipHandler.UpdateArsip)
	adminArsip.Get("/delete/:id", authRoleAdminMiddleware(userService), arsipHandler.DeleteArsipSoft)
	adminArsip.Get("/recycle", authRoleAdminMiddleware(userService), arsipHandler.ShowArsipListRecycle)
	adminArsip.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), arsipHandler.DeleteArsip)
	adminArsip.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), arsipHandler.RestoreArsip)

	// kategori arsip admin
	adminKategoriArsip := adminArsip.Group("/kategori")
	adminKategoriArsip.Get("", authRoleAdminMiddleware(userService), arsipkategoriHandler.ShowArsipKategoriList)
	adminKategoriArsip.Post("/new", authRoleAdminMiddleware(userService), arsipkategoriHandler.NewArsipKategori)
	adminKategoriArsip.Get("/update/:id", authRoleAdminMiddleware(userService), arsipkategoriHandler.UpdateArsipKategoriView)
	adminKategoriArsip.Post("/update", authRoleAdminMiddleware(userService), arsipkategoriHandler.UpdateArsipKategori)
	adminKategoriArsip.Get("/delete-dokumen/:id", authRoleAdminMiddleware(userService), fileHandler.UpdateFileArsip)
	adminKategoriArsip.Get("/delete/:id", authRoleAdminMiddleware(userService), arsipkategoriHandler.DeletedArsipKategoriSoft)
	adminKategoriArsip.Get("/recycle", authRoleAdminMiddleware(userService), arsipkategoriHandler.ShowKategoriArsipListRecycle)
	adminKategoriArsip.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), arsipkategoriHandler.DeletedArsipKategoriRecycle)
	adminKategoriArsip.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), arsipkategoriHandler.RestoreArsipKategori)

	// dokumen file admin
	adminDokumen := admin.Group("/dokumen")
	adminDokumen.Get("", authRoleAdminMiddleware(userService), fileHandler.ShowFileAll)
	adminDokumen.Get("/detail/:id", authRoleAdminMiddleware(userService), fileHandler.ShowFileDetail)
	adminDokumen.Post("/new", authRoleAdminMiddleware(userService), fileHandler.NewFile)
	adminDokumen.Get("/update/:id", authRoleAdminMiddleware(userService), fileHandler.UpdateFileView)
	adminDokumen.Post("/update", authRoleAdminMiddleware(userService), fileHandler.UpdateFile)
	adminDokumen.Get("/delete/:id", authRoleAdminMiddleware(userService), fileHandler.DeletetFileSoft)
	adminDokumen.Get("/recycle", authRoleAdminMiddleware(userService), fileHandler.ShowFileAllRecycle)
	adminDokumen.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), fileHandler.DeletetFileRecycle)
	adminDokumen.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), fileHandler.RestoreFile)
	adminDokumen.Get("/Key/:id", authRoleAdminMiddleware(userService), fileHandler.EnkripsiEndDekripsiFile)
	adminDokumen.Get("/download/:id", authRoleAdminMiddleware(userService), fileHandler.DownloadFile)
	adminDokumen.Post("/update/arsip", authRoleAdminMiddleware(userService), fileHandler.FileArsipIDUpdate)
	adminDokumen.Get("/arsip/delete-dokumen/:id", authRoleAdminMiddleware(userService), fileHandler.FileArsipIDUpdateDelete)

	// admin berita
	adminBerita := admin.Group("/berita")
	adminBerita.Get("", authRoleAdminMiddleware(userService), beritaHandler.ShowBeritaList)
	adminBerita.Get("/detail/:id", authRoleAdminMiddleware(userService), beritaHandler.ShowBeritaDetail)
	adminBerita.Post("/new", authRoleAdminMiddleware(userService), beritaHandler.NewBerita)
	adminBerita.Get("/update/:id", authRoleAdminMiddleware(userService), beritaHandler.UpdateBeritaView)
	adminBerita.Post("/update", authRoleAdminMiddleware(userService), beritaHandler.UpdateBerita)
	adminBerita.Get("/delete/:id", authRoleAdminMiddleware(userService), beritaHandler.DeleteBeritaSoft)
	adminBerita.Get("/recycle", authRoleAdminMiddleware(userService), beritaHandler.ShowBeritaListRecycle)
	adminBerita.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), beritaHandler.DeleteBerita)
	adminBerita.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), beritaHandler.RestoreBerita)

	// admin berita image
	adminBeritaImage := adminBerita.Group("/image")
	adminBeritaImage.Post("/new", authRoleAdminMiddleware(userService), beritaImageHandler.NewBeritaImage)
	adminBeritaImage.Get("/delete/:id", authRoleAdminMiddleware(userService), beritaImageHandler.DeleteBeritaImage)

	// admin berita kategori
	adminBeritaKategori := adminBerita.Group("/kategori")
	adminBeritaKategori.Get("", authRoleAdminMiddleware(userService), beritaKategoriHandler.ShowBeritaKategoriList)
	adminBeritaKategori.Post("/new", authRoleAdminMiddleware(userService), beritaKategoriHandler.NewBeritaKategori)
	adminBeritaKategori.Get("/update/:id", authRoleAdminMiddleware(userService), beritaKategoriHandler.UpdateBeritaKategoriView)
	adminBeritaKategori.Post("/update", authRoleAdminMiddleware(userService), beritaKategoriHandler.UpdateBeritaKategori)
	adminBeritaKategori.Get("/delete/:id", authRoleAdminMiddleware(userService), beritaKategoriHandler.DeletedBeritaKategoriSoft)
	adminBeritaKategori.Get("/recycle", authRoleAdminMiddleware(userService), beritaKategoriHandler.ShowKategoriBeritaListRecycle)
	adminBeritaKategori.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), beritaKategoriHandler.DeletedBeritaKategoriRecycle)
	adminBeritaKategori.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), beritaKategoriHandler.RestoreBeritaKategori)

	// penumuman
	adminPengumuman := admin.Group("/pengumuman")
	adminPengumuman.Get("", authRoleAdminMiddleware(userService), pengumumanHanlder.ShowPengumumanList)
	adminPengumuman.Get("/detail/:id", authRoleAdminMiddleware(userService), pengumumanHanlder.ShowPengumumanDetail)
	adminPengumuman.Post("/new", authRoleAdminMiddleware(userService), pengumumanHanlder.NewPengumuman)
	adminPengumuman.Get("/update/:id", authRoleAdminMiddleware(userService), pengumumanHanlder.UpdatePengumumanView)
	adminPengumuman.Post("/update", authRoleAdminMiddleware(userService), pengumumanHanlder.UpdatePengumuman)
	adminPengumuman.Get("/delete/:id", authRoleAdminMiddleware(userService), pengumumanHanlder.DeletePengumumanSoft)
	adminPengumuman.Get("/recycle", authRoleAdminMiddleware(userService), pengumumanHanlder.ShowPengumumanListRecycle)
	adminPengumuman.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), pengumumanHanlder.DeletePengumuman)
	adminPengumuman.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), pengumumanHanlder.RestorePengumuman)

	// visi & misi
	adminVisiMisi := admin.Group("/visi-misi")
	adminVisiMisi.Get("", authRoleAdminMiddleware(userService), visiMisiHandler.ShowVisiMisiList)
	adminVisiMisi.Get("/detail/:id", authRoleAdminMiddleware(userService), visiMisiHandler.ShowVisiMisiDetail)
	adminVisiMisi.Post("/new", authRoleAdminMiddleware(userService), visiMisiHandler.NewVisiMisi)
	adminVisiMisi.Get("/update/:id", authRoleAdminMiddleware(userService), visiMisiHandler.UpdateVisiMisiView)
	adminVisiMisi.Post("/update", authRoleAdminMiddleware(userService), visiMisiHandler.UpdateVisiMisi)
	adminVisiMisi.Get("/delete/:id", authRoleAdminMiddleware(userService), visiMisiHandler.DeletevisiMisiSoft)
	adminVisiMisi.Get("/image/delete/:id", authRoleAdminMiddleware(userService), visiMisiHandler.DeleteVisiMisiImage)
	adminVisiMisi.Get("/recycle", authRoleAdminMiddleware(userService), visiMisiHandler.ShowVisiMisiListRecycle)
	adminVisiMisi.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), visiMisiHandler.DeleteVisiMisi)
	adminVisiMisi.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), visiMisiHandler.RestorevisiMisi)

	// Demografi
	adminDemografi := admin.Group("demografi")
	adminDemografi.Get("", authRoleAdminMiddleware(userService), demografiHandler.ShowDemografiList)
	adminDemografi.Get("/detail/:id", authRoleAdminMiddleware(userService), demografiHandler.ShowDemografiDetail)
	adminDemografi.Post("/new", authRoleAdminMiddleware(userService), demografiHandler.NewDemografi)
	adminDemografi.Get("/update/:id", authRoleAdminMiddleware(userService), demografiHandler.UpdateDemografiView)
	adminDemografi.Post("/update", authRoleAdminMiddleware(userService), demografiHandler.UpdateDemografi)
	adminDemografi.Get("/delete/:id", authRoleAdminMiddleware(userService), demografiHandler.DeleteDemografiSoft)
	adminDemografi.Get("/image/delete/:id", authRoleAdminMiddleware(userService), demografiHandler.DeleteDemografiImage)
	adminDemografi.Get("/recycle", authRoleAdminMiddleware(userService), demografiHandler.ShowDemografiListRecycle)
	adminDemografi.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), demografiHandler.DeleteDemografi)
	adminDemografi.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), demografiHandler.RestoreDemografi)

	// Sejarah
	adminSejarah := admin.Group("sejarah")
	adminSejarah.Get("", authRoleAdminMiddleware(userService), sejarahHandler.ShowSejarahList)
	adminSejarah.Get("/detail/:id", authRoleAdminMiddleware(userService), sejarahHandler.ShowSejarahDetail)
	adminSejarah.Post("/new", authRoleAdminMiddleware(userService), sejarahHandler.NewSejarah)
	adminSejarah.Get("/update/:id", authRoleAdminMiddleware(userService), sejarahHandler.UpdateSejarahView)
	adminSejarah.Post("/update", authRoleAdminMiddleware(userService), sejarahHandler.UpdateSejarah)
	adminSejarah.Get("/delete/:id", authRoleAdminMiddleware(userService), sejarahHandler.DeleteSejarahSoft)
	adminSejarah.Get("/image/delete/:id", authRoleAdminMiddleware(userService), sejarahHandler.DeleteSejarahImage)
	adminSejarah.Get("/recycle", authRoleAdminMiddleware(userService), sejarahHandler.ShowSejarahListRecycle)
	adminSejarah.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), sejarahHandler.DeleteSejarah)
	adminSejarah.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), sejarahHandler.RestoreSejarah)

	// profile Desa
	adminProfileDesa := admin.Group("/profile-desa")
	adminProfileDesa.Get("", authRoleAdminMiddleware(userService), profileDesaHandler.ShowProfileDesaList)
	adminProfileDesa.Get("/detail/:id", authRoleAdminMiddleware(userService), profileDesaHandler.ShowProfileDesaDetail)
	adminProfileDesa.Post("/new", authRoleAdminMiddleware(userService), profileDesaHandler.NewProfileDesa)
	adminProfileDesa.Get("/update/:id", authRoleAdminMiddleware(userService), profileDesaHandler.UpdateProfileDesaView)
	adminProfileDesa.Post("/update", authRoleAdminMiddleware(userService), profileDesaHandler.UpdateProfileDesa)
	adminProfileDesa.Get("/delete/:id", authRoleAdminMiddleware(userService), profileDesaHandler.DeleteProfileDesaSoft)
	adminProfileDesa.Get("/image/delete/:id", authRoleAdminMiddleware(userService), profileDesaHandler.DeleteProfileDesaImage)
	adminProfileDesa.Get("/recycle", authRoleAdminMiddleware(userService), profileDesaHandler.ShowProfileDesaListRecycle)
	adminProfileDesa.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), profileDesaHandler.DeleteProfileDesa)
	adminProfileDesa.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), profileDesaHandler.RestoreProfileDesa)

	// image show
	adminImageShow := admin.Group("/image-show")
	adminImageShow.Get("", authRoleAdminMiddleware(userService), imageSlideShowHandler.ShowimageSlideShowList)
	adminImageShow.Get("/detail/:id", authRoleAdminMiddleware(userService), imageSlideShowHandler.ShowImageSlideShowDetail)
	adminImageShow.Post("/new", authRoleAdminMiddleware(userService), imageSlideShowHandler.NewImageSlideShow)
	adminImageShow.Get("/update/:id", authRoleAdminMiddleware(userService), imageSlideShowHandler.UpdateImageSlideShowView)
	adminImageShow.Post("/update", authRoleAdminMiddleware(userService), imageSlideShowHandler.UpdateImageSlideShow)
	adminImageShow.Get("/delete/:id", authRoleAdminMiddleware(userService), imageSlideShowHandler.DeleteImageSlideShowSoft)
	adminImageShow.Get("/recycle", authRoleAdminMiddleware(userService), imageSlideShowHandler.ShowImageSlideShowListRecycle)
	adminImageShow.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), imageSlideShowHandler.DeleteImageSlideShow)
	adminImageShow.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), imageSlideShowHandler.RestoreImageSlideShow)

	// pegawai
	adminPegawai := admin.Group("/pegawai")
	adminPegawai.Get("", authRoleAdminMiddleware(userService), pegawaiHandler.ShowPegawaiList)
	adminPegawai.Get("/detail/:id", authRoleAdminMiddleware(userService), pegawaiHandler.ShowPegawaiDetail)
	adminPegawai.Post("/new", authRoleAdminMiddleware(userService), pegawaiHandler.NewPegawai)
	adminPegawai.Get("/update/:id", authRoleAdminMiddleware(userService), pegawaiHandler.UpdatePegawaiView)
	adminPegawai.Post("/update", authRoleAdminMiddleware(userService), pegawaiHandler.UpdatePegawai)
	adminPegawai.Get("/delete/:id", authRoleAdminMiddleware(userService), pegawaiHandler.DeletePegawaiSoft)
	adminUser.Get("/profile/image/:id", authRoleAdminMiddleware(userService), pegawaiHandler.UploadImageProfilePegawaiView)
	adminUser.Post("/profile/image/", authRoleAdminMiddleware(userService), pegawaiHandler.UploadImageProfilePegawai)
	adminPegawai.Get("/recycle", authRoleAdminMiddleware(userService), pegawaiHandler.ShowPegawaiListRecycle)
	adminPegawai.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), pegawaiHandler.DeletePegawai)
	adminPegawai.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), pegawaiHandler.RestorePegawai)

	//surat
	adminSurat := admin.Group("/surat")
	adminSurat.Get("/", authRoleAdminMiddleware(userService), suratHandler.ShowSuratList)
	adminSurat.Get("/delete/:unik/:id", authRoleAdminMiddleware(userService), suratHandler.DeleteSuratSoft)
	adminSurat.Get("/recycle", authRoleAdminMiddleware(userService), suratHandler.ShowSuratListRecycle)
	adminSurat.Get("/recycle/delete/:unik/:id", authRoleAdminMiddleware(userService), suratHandler.DeleteSurat)
	adminSurat.Get("/recycle/restore/:unik/:id", authRoleAdminMiddleware(userService), suratHandler.RestoreSurat)
	adminSurat.Get("/download/:unik/:id", authRoleAdminMiddleware(userService), suratHandler.DownloadFileSurat)

	// surat usaha
	adminSurat.Get("/surat-usaha/new", authRoleAdminMiddleware(userService), suratHandler.InputSuartUsahaView)
	adminSurat.Post("/surat-usaha/new", authRoleAdminMiddleware(userService), suratHandler.InputSuartUsaha)
	adminSurat.Get("/surat-usaha/update/:id", authRoleAdminMiddleware(userService), suratHandler.UpdateSuartUsahaView)
	adminSurat.Post("/surat-usaha/update/", authRoleAdminMiddleware(userService), suratHandler.UpdateSuratUsaha)

	// surat keteranagn kematian
	adminSurat.Get("/surat-keterangan-kematian/new", authRoleAdminMiddleware(userService), suratHandler.InputSuartKematianView)
	adminSurat.Post("/surat-keterangan-kematian/new", authRoleAdminMiddleware(userService), suratHandler.InputSuartKematian)
	adminSurat.Get("/surat-keterangan-kematian/update/:id", authRoleAdminMiddleware(userService), suratHandler.UpdateSuartKematianView)
	adminSurat.Post("/surat-keterangan-kematian/update/", authRoleAdminMiddleware(userService), suratHandler.UpdateSuratKematian)

	// surat pengatar nikah N1
	adminSurat.Get("/surat-pengatar-nikah/new", authRoleAdminMiddleware(userService), suratHandler.InputSuratKeteranganNikahNSatuView)
	adminSurat.Post("/surat-pengatar-nikah/new", authRoleAdminMiddleware(userService), suratHandler.InputSuratKeteranganNikahNSatu)
	adminSurat.Get("/surat-pengatar-nikah/update/:id", authRoleAdminMiddleware(userService), suratHandler.UpdateSuratKeteranganNikahNSatuView)
	adminSurat.Post("/surat-pengatar-nikah/update", authRoleAdminMiddleware(userService), suratHandler.UpdateSuratKeteranganNikahNSatu)

	// surat persetujuan pengatin N4
	adminSurat.Get("/surat-persetujuan-pengatin/new", authRoleAdminMiddleware(userService), suratHandler.InputSuratKeteranganNikahNEmpatView)
	adminSurat.Post("/surat-persetujuan-pengatin/new", authRoleAdminMiddleware(userService), suratHandler.InputSuratKeteranganNikahNEmpat)
	adminSurat.Get("/surat-persetujuan-pengatin/update/:id", authRoleAdminMiddleware(userService), suratHandler.UpdateSuratKeteranganNikahNEmpatView)
	adminSurat.Post("/surat-persetujuan-pengatin/update", authRoleAdminMiddleware(userService), suratHandler.UpdateSuratKeteranganNikahNEmpat)

	// surat izin orang tua N5
	adminSurat.Get("/surat-izin-orang-tua/new", authRoleAdminMiddleware(userService), suratHandler.InputSuratKeteranganNikahNLimaView)
	adminSurat.Post("/surat-izin-orang-tua/new", authRoleAdminMiddleware(userService), suratHandler.InputSuratKeteranganNikahNLima)
	adminSurat.Get("/surat-izin-orang-tua/update/:id", authRoleAdminMiddleware(userService), suratHandler.UpdateSuratKeteranganNikahNLimaView)
	adminSurat.Post("/surat-izin-orang-tua/update", authRoleAdminMiddleware(userService), suratHandler.UpdateSuratKeteranganNikahNLima)

	// adminSurat.Get("/surat-pernyataan-kepemilikan-tanah/new", authRoleAdminMiddleware(userService), suratHandler.InputSuratKeteranganNikahNLimaView)
	// adminSurat.Post("/surat-pernyataan-kepemilikan-tanah/new", authRoleAdminMiddleware(userService), suratHandler.InputSuratKeteranganNikahNLima)

	//surat setting
	adminFileSurat := adminSurat.Group("/setting")
	adminFileSurat.Get("", authRoleAdminMiddleware(userService), fileSuratHandler.ShowFileSuratAll)
	adminFileSurat.Post("/new", authRoleAdminMiddleware(userService), fileSuratHandler.NewFileSurat)
	adminFileSurat.Get("/update/:id", authRoleAdminMiddleware(userService), fileSuratHandler.UpdateFileSuratView)
	adminFileSurat.Post("/update", authRoleAdminMiddleware(userService), fileSuratHandler.UpdateFileSurat)
	adminFileSurat.Get("/update-file/:id", authRoleAdminMiddleware(userService), fileSuratHandler.UpdateFileSuratSeconView)
	adminFileSurat.Post("/update-file", authRoleAdminMiddleware(userService), fileSuratHandler.UpdateFileSuratSeconeFile)
	adminFileSurat.Get("/delete/:id", authRoleAdminMiddleware(userService), fileSuratHandler.DeleteFileSuratSoft)
	adminFileSurat.Get("/recycle", authRoleAdminMiddleware(userService), fileSuratHandler.ShowFileSuratListRecycle)
	adminFileSurat.Get("/recycle/delete/:id", authRoleAdminMiddleware(userService), fileSuratHandler.DeleteFileSurat)
	adminFileSurat.Get("/recycle/restore/:id", authRoleAdminMiddleware(userService), fileSuratHandler.RestoreFileSurat)
	adminFileSurat.Get("/file-main/download/:id", authRoleAdminMiddleware(userService), fileSuratHandler.DownloadFileSuratMain)
	adminFileSurat.Get("/file/download/:id", authRoleAdminMiddleware(userService), fileSuratHandler.DownloadFileSurat)
	adminFileSurat.Get("/delete-file/:id", authRoleAdminMiddleware(userService), fileSuratHandler.RemoveFileSurat)

	// user
	user := dasboard.Group("/user")
	user.Get("/", authRoleUsereMiddleware(userService), dasboardHandle.ShowDasbordUserView)
	user.Get("/profile", authRoleUsereMiddleware(userService), adminHandler.ShowUserProfile)

	// user user
	userUser := user.Group("/user")
	userUser.Get("", authRoleUsereMiddleware(userService), adminHandler.ShowAdminAll)
	userUser.Get("/detail/:id", authRoleUsereMiddleware(userService), adminHandler.ShowAdminDetail)

	// arsip user
	userArsip := user.Group("/arsip")
	userArsip.Get("", authRoleUsereMiddleware(userService), arsipHandlerUser.ShowArsipList)
	userArsip.Get("/detail/:id", authRoleUsereMiddleware(userService), arsipHandlerUser.ShowArsipDetail)

	// dokumen file
	userDokumen := user.Group("/dokumen")
	userDokumen.Get("", authRoleUsereMiddleware(userService), authRoleUsereMiddleware(userService), fileHandlerUser.ShowFileAll)
	userDokumen.Get("/detail/:id", authRoleUsereMiddleware(userService), authRoleUsereMiddleware(userService), fileHandlerUser.ShowFileDetail)
	userDokumen.Get("/download/:id", authRoleUsereMiddleware(userService), fileHandlerUser.DownloadFile)

	app.Listen(os.Getenv("APP_PORT"))

}

func LoginMiddleware(userSession user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		cookisUserId := c.Cookies("sessionLog")
		idUser, err := helper.GetSessionID(cookisUserId)
		if err != nil {
			return c.Next()
		}

		user, err := userSession.GetUserByID(idUser)
		if err != nil {
			return c.Next()
		}

		if user.Role == "admin" {
			return c.Redirect("/dasboard/admin/")
		} else if user.Role == "user" {
			return c.Redirect("/dasboard/user/")
		}

		return c.Next()
	}
}

func authMiddleware(userSession user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookisUserId := c.Cookies("sessionLog")
		idUser, err := helper.GetSessionID(cookisUserId)
		if err != nil {
			return c.Redirect("/login")
		}

		_, err = userSession.GetUserByID(idUser)
		if err != nil {
			return c.Redirect("/login")
		}

		return c.Next()
	}
}

func authRoleAdminMiddleware(userSession user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		cookisUserId := c.Cookies("sessionLog")
		idUser, err := helper.GetSessionID(cookisUserId)
		if err != nil {
			return c.Redirect("/login")
		}

		user, err := userSession.GetUserByID(idUser)
		if err != nil {
			return c.Redirect("/login")
		}

		if user.Role == "user" {
			return c.Redirect("/dasboard/user/")
		}

		return c.Next()
	}
}

func authRoleUsereMiddleware(userSession user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		cookisUserId := c.Cookies("sessionLog")
		idUser, err := helper.GetSessionID(cookisUserId)
		if err != nil {
			return c.Redirect("/login")
		}

		user, err := userSession.GetUserByID(idUser)
		if err != nil {
			return c.Redirect("/login")
		}

		if user.Role == "admin" {
			return c.Redirect("/dasboard/admin/")
		}

		return c.Next()
	}
}
