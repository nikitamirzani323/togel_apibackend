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
	api := app.Group("/api/", middleware.JWTProtected())

	api.Post("home", controllers.Home)
	api.Post("generatepassword", controllers.GenerateHashPassword)
	api.Post("allpasaran", controllers.PasaranHome)
	api.Post("editpasaran", controllers.PasaranDetail)
	api.Post("savepasaran", controllers.PasaranSave)
	api.Post("savepasaranonline", controllers.PasaranSaveOnline)
	api.Post("deletepasaranonline", controllers.PasaranDeleteOnline)
	api.Post("savepasaranlimitline", controllers.PasaranSaveLimit)
	api.Post("savepasaranconf432", controllers.PasaranSaveConf432d)
	api.Post("savepasaranconfcolokbebas", controllers.PasaranSaveConfColokBebas)
	api.Post("savepasaranconfcolokmacau", controllers.PasaranSaveConfColokMacau)
	api.Post("savepasaranconfcoloknaga", controllers.PasaranSaveConfColokNaga)
	api.Post("savepasaranconfcolokjitu", controllers.PasaranSaveConfColokJitu)
	api.Post("savepasaranconf5050umum", controllers.PasaranSaveConf5050Umum)
	api.Post("savepasaranconf5050special", controllers.PasaranSaveConf5050Special)
	api.Post("savepasaranconf5050kombinasi", controllers.PasaranSaveConf5050Kombinasi)
	api.Post("savepasaranconfmacaukombinasi", controllers.PasaranSaveConfMacauKombinasi)
	api.Post("savepasaranconfdasar", controllers.PasaranSaveConfDasar)
	api.Post("savepasaranconfshio", controllers.PasaranSaveConfShio)

	api.Post("allperiode", controllers.PeriodeHome)
	api.Post("listpasaran", controllers.Periodelistpasaran)
	api.Post("listprediksi", controllers.Periodelistprediksi)
	api.Post("editperiode", controllers.PeriodeDetail)
	api.Post("periodelistmemberbynomor", controllers.PeriodeListMemberByNomor)
	api.Post("periodelistmember", controllers.PeriodeListMember)
	api.Post("periodelistbet", controllers.PeriodeListBet)
	api.Post("periodelistbetstatus", controllers.PeriodeListBetstatus)
	api.Post("periodelistbetusername", controllers.PeriodeListBetusername)
	api.Post("periodelistbettable", controllers.PeriodeListBetTable)
	api.Post("periodebettable", controllers.PeriodeBetTable)
	api.Post("saveperiode", controllers.PeriodeSave)
	api.Post("savepasarannew", controllers.PeriodeSaveNew)
	api.Post("saveperioderevisi", controllers.PeriodeSaveRevisi)
	api.Post("cancelbet", controllers.PeriodeCancelBet)

	api.Post("alladmin", controllers.AdminHome)
	api.Post("editadmin", controllers.AdminDetail)
	api.Post("saveadmin", controllers.AdminSave)
	api.Post("saveadminiplist", controllers.AdminSaveIplist)
	api.Post("deleteadminiplist", controllers.AdminDeleteIplist)

	api.Post("alladminrule", controllers.AdminruleHome)
	api.Post("editadminrule", controllers.AdminruleDetail)
	api.Post("saveadminrule", controllers.SaveAdminruleDetail)
	api.Post("saveadminruleconf", controllers.SaveAdminruleConf)

	api.Post("dashboardwinlose", controllers.DashboardWinlose)
	api.Post("dashboardpasaranwinlose", controllers.DashboardWinlosepasaran)
	api.Post("reportwinlose", controllers.Reportwinlose)

	api.Post("log", controllers.LogHome)
	return app
}
