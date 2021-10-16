package routes

import (
	"bitbucket.org/isbtotogroup/apibackend_go/controllers"
	"bitbucket.org/isbtotogroup/apibackend_go/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())

	app.Post("/api/login", controllers.CheckLogin)

	app.Post("/api/home", middleware.JWTProtected(), controllers.Home)
	app.Post("/api/generatepassword", controllers.GenerateHashPassword)
	app.Post("/api/allpasaran", middleware.JWTProtected(), controllers.PasaranHome)
	app.Post("/api/editpasaran", middleware.JWTProtected(), controllers.PasaranDetail)
	app.Post("/api/savepasaran", middleware.JWTProtected(), controllers.PasaranSave)
	app.Post("/api/savepasaranonline", middleware.JWTProtected(), controllers.PasaranSaveOnline)
	app.Post("/api/deletepasaranonline", middleware.JWTProtected(), controllers.PasaranDeleteOnline)
	app.Post("/api/savepasaranlimitline", middleware.JWTProtected(), controllers.PasaranSaveLimit)
	app.Post("/api/savepasaranconf432", middleware.JWTProtected(), controllers.PasaranSaveConf432d)
	app.Post("/api/savepasaranconfcolokbebas", middleware.JWTProtected(), controllers.PasaranSaveConfColokBebas)
	app.Post("/api/savepasaranconfcolokmacau", middleware.JWTProtected(), controllers.PasaranSaveConfColokMacau)
	app.Post("/api/savepasaranconfcoloknaga", middleware.JWTProtected(), controllers.PasaranSaveConfColokNaga)
	app.Post("/api/savepasaranconfcolokjitu", middleware.JWTProtected(), controllers.PasaranSaveConfColokJitu)
	app.Post("/api/savepasaranconf5050umum", middleware.JWTProtected(), controllers.PasaranSaveConf5050Umum)
	app.Post("/api/savepasaranconf5050special", middleware.JWTProtected(), controllers.PasaranSaveConf5050Special)
	app.Post("/api/savepasaranconf5050kombinasi", middleware.JWTProtected(), controllers.PasaranSaveConf5050Kombinasi)
	app.Post("/api/savepasaranconfmacaukombinasi", middleware.JWTProtected(), controllers.PasaranSaveConfMacauKombinasi)
	app.Post("/api/savepasaranconfdasar", middleware.JWTProtected(), controllers.PasaranSaveConfDasar)
	app.Post("/api/savepasaranconfshio", middleware.JWTProtected(), controllers.PasaranSaveConfShio)

	app.Post("/api/allperiode", middleware.JWTProtected(), controllers.PeriodeHome)
	app.Post("/api/listpasaran", middleware.JWTProtected(), controllers.Periodelistpasaran)
	app.Post("/api/listprediksi", middleware.JWTProtected(), controllers.Periodelistprediksi)
	app.Post("/api/editperiode", middleware.JWTProtected(), controllers.PeriodeDetail)
	app.Post("/api/periodelistmemberbynomor", middleware.JWTProtected(), controllers.PeriodeListMemberByNomor)
	app.Post("/api/periodelistmember", middleware.JWTProtected(), controllers.PeriodeListMember)
	app.Post("/api/periodelistbet", middleware.JWTProtected(), controllers.PeriodeListBet)
	app.Post("/api/periodelistbettable", middleware.JWTProtected(), controllers.PeriodeListBetTable)
	app.Post("/api/periodebettable", middleware.JWTProtected(), controllers.PeriodeBetTable)
	app.Post("/api/saveperiode", middleware.JWTProtected(), controllers.PeriodeSave)
	app.Post("/api/savepasarannew", middleware.JWTProtected(), controllers.PeriodeSaveNew)
	app.Post("/api/saveperioderevisi", middleware.JWTProtected(), controllers.PeriodeSaveRevisi)
	app.Post("/api/cancelbet", middleware.JWTProtected(), controllers.PeriodeCancelBet)

	app.Post("/api/alladmin", middleware.JWTProtected(), controllers.AdminHome)
	app.Post("/api/editadmin", middleware.JWTProtected(), controllers.AdminDetail)
	app.Post("/api/saveadmin", middleware.JWTProtected(), controllers.AdminSave)
	app.Post("/api/saveadminiplist", middleware.JWTProtected(), controllers.AdminSaveIplist)
	app.Post("/api/deleteadminiplist", middleware.JWTProtected(), controllers.AdminDeleteIplist)

	app.Post("/api/alladminrule", middleware.JWTProtected(), controllers.AdminruleHome)
	app.Post("/api/editadminrule", middleware.JWTProtected(), controllers.AdminruleDetail)
	app.Post("/api/saveadminrule", middleware.JWTProtected(), controllers.SaveAdminruleDetail)
	app.Post("/api/saveadminruleconf", middleware.JWTProtected(), controllers.SaveAdminruleConf)

	app.Post("/api/dashboardwinlose", middleware.JWTProtected(), controllers.DashboardHome)
	app.Post("/api/reportwinlose", middleware.JWTProtected(), controllers.Reportwinlose)
	return app
}
