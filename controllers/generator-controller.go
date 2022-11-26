package controllers

import (
	"log"
	"strconv"
	"strings"

	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"bitbucket.org/isbtotogroup/apibackend_go/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Generatorsave struct {
	Page        string `json:"page"`
	Invoice     string `json:"invoice" `
	Totalmember int    `json:"totalmember" `
	Totalrow    int    `json:"totalrow" `
}

func GeneratorSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(Generatorsave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")
	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)
	if typeadmin == "MASTER" {
		result, err := models.Save_Generator(client_username, client_company, client.Invoice, client.Totalmember, client.Totalrow)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		return c.JSON(result)
	} else {
		if !flag_page {
			c.Status(fiber.StatusForbidden)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusForbidden,
				"message": "Mohon maaf Anda tidak bisa akses halaman ini",
				"record":  nil,
			})
		} else {
			result, err := models.Save_Generator(client_username, client_company, client.Invoice, client.Totalmember, client.Totalrow)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			return c.JSON(result)
		}
	}
}
func _deleteredis_generator(company string, idtrxkeluaran int) {
	log.Println("REDIS DELETE")
	log.Println("COMPANY :", company)
	log.Println("INVOICE :", idtrxkeluaran)

	//AGEN
	field_home_redis := Fieldperiode_home_redis + strings.ToLower(company)
	val_homeredis := helpers.DeleteRedis(field_home_redis)
	log.Printf("Redis Delete AGEN - PERIODE HOME : %d", val_homeredis)

	field_homedetail_redis := Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran)
	val_homedetailredis := helpers.DeleteRedis(field_homedetail_redis)
	log.Printf("%s\n", field_homedetail_redis)
	log.Printf("Redis Delete AGEN - PERIODE DETAIL : %d", val_homedetailredis)

	field_homedetail_listmember_redis := Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTMEMBER"
	val_homedetaillistmember_redis := helpers.DeleteRedis(field_homedetail_listmember_redis)
	log.Printf("%s\n", field_homedetail_listmember_redis)
	log.Printf("Redis Delete AGEN - PERIODE DETAIL LISTMEMBER : %d", val_homedetaillistmember_redis)

	field_homedetail_listbettable_redis := Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTBETTABLE"
	val_homedetaillistbettable_redis := helpers.DeleteRedis(field_homedetail_listbettable_redis)
	log.Printf("Redis Delete AGEN - PERIODE DETAIL LISTBETTABLE : %d", val_homedetaillistbettable_redis)

	log_redis := "LISTLOG_AGENT_" + strings.ToLower(company)
	val_agent_redis := helpers.DeleteRedis(log_redis)
	log.Printf("Redis Delete AGEN - LOG status: %d", val_agent_redis)

	val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + strings.ToLower(company))
	log.Printf("Redis Delete AGENT DASHBOARD status: %d", val_agent_dashboard)

	val_agent_dashboard_pasaranhome := helpers.DeleteRedis("LISTDASHBOARDPASARAN_AGENT_" + strings.ToLower(company))
	log.Printf("Redis Delete AGENT DASHBOARD PASARAN status: %d", val_agent_dashboard_pasaranhome)

	val_agent4d := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_4D")
	val_agent3d := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_3D")
	val_agent2d := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_2D")
	val_agent2dd := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_2DD")
	val_agent2dt := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_2DT")
	val_agentcolokbebas := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_COLOK_BEBAS")
	val_agentcolokmacau := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_COLOK_MACAU")
	val_agentcoloknaga := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_COLOK_NAGA")
	val_agentcolokjitu := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_COLOK_JITU")
	val_agent5050umum := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_50_50_UMUM")
	val_agent5050special := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_50_50_SPECIAL")
	val_agent5050kombinasi := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_50_50_KOMBINASI")
	val_agentmacaukombinasi := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_MACAU_KOMBINASI")
	val_agentdasar := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_DASAR")
	val_agentshio := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_SHIO")
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 4D: %d", val_agent4d)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 3D: %d", val_agent3d)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 2D: %d", val_agent2d)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 2DD: %d", val_agent2dd)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 2DT: %d", val_agent2dt)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET COLOK BEBAS: %d", val_agentcolokbebas)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET COLOK MACAU: %d", val_agentcolokmacau)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET COLOK NAGA: %d", val_agentcoloknaga)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET COLOK JITU: %d", val_agentcolokjitu)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 5050UMUM: %d", val_agent5050umum)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 5050SPECIAL: %d", val_agent5050special)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 5050KOMBINASI: %d", val_agent5050kombinasi)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET MACAU KOMBINASI: %d", val_agentmacaukombinasi)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET DASAR: %d", val_agentdasar)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET SHIO: %d", val_agentshio)
	val_agentall := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTBET_all")
	val_agentwinner := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTBET_winner")
	val_agentcancel := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTBET_cancel")
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS ALL: %d", val_agentall)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS WINNER: %d", val_agentwinner)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS CANCEL: %d", val_agentcancel)

}
