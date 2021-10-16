package controllers

import (
	"log"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"bitbucket.org/isbtotogroup/apibackend_go/models"
	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type periodedetail struct {
	Idtrxkeluaran int    `json:"idinvoice"`
	Permainan     string `json:"permainan"`
}
type periodedetailstatus struct {
	Idtrxkeluaran int    `json:"idinvoice"`
	Status        string `json:"status"`
}
type periodedetailmembernomor struct {
	Idtrxkeluaran int    `json:"idinvoice" validate:"required"`
	Permainan     string `json:"permainan" validate:"required"`
	Nomor         string `json:"nomor" validate:"required"`
}
type periodeSave struct {
	Sdata          string `json:"sData" validate:"required"`
	Page           string `json:"page"`
	Idtrxkeluaran  int    `json:"idinvoice" validate:"required"`
	Nomorkeluaran  string `json:"nomorkeluaran" validate:"required,min=4,max=4"`
	Idpasarantogel string `json:"idpasarancode" validate:"required"`
}
type periodeSaveNew struct {
	Sdata         string `json:"sData" validate:"required"`
	Page          string `json:"page"`
	Idcomppasaran int    `json:"pasaran_code" validate:"required"`
}
type periodesaverevisi struct {
	Sdata     string `json:"sData" validate:"required"`
	Page      string `json:"page"`
	Idinvoice int    `json:"idinvoice" validate:"required"`
	Msgrevisi string `json:"msgrevisi" validate:"required"`
}
type periodecancelbet struct {
	Sdata           string `json:"sData" validate:"required"`
	Page            string `json:"page"`
	Idinvoice       int    `json:"idinvoice" validate:"required"`
	Idinvoicedetail int    `json:"idinvoicedetail" validate:"required"`
}
type periodeSavePrediksi struct {
	Sdata     string `json:"sData" validate:"required"`
	Page      string `json:"page"`
	Idinvoice int    `json:"idinvoice" validate:"required"`
}
type periodePrediksi struct {
	Nomorkeluaran string `json:"nomorkeluaran" validate:"required,min=4,max=4"`
	Idcomppasaran int    `json:"pasaran_code" validate:"required"`
}
type responseredis_periodehome struct {
	Pasaran_no               int    `json:"pasaran_no"`
	Pasaran_invoice          int    `json:"pasaran_invoice"`
	Pasaran_idcompp          int    `json:"pasaran_idcompp"`
	Pasaran_periode          string `json:"pasaran_periode"`
	Pasaran_code             string `json:"pasaran_code"`
	Pasaran_name             string `json:"pasaran_name"`
	Pasaran_tanggal          string `json:"pasaran_tanggal"`
	Pasaran_keluaran         string `json:"pasaran_keluaran"`
	Pasaran_status           string `json:"pasaran_status"`
	Pasaran_status_css       string `json:"pasaran_status_css"`
	Pasaran_totalmember      int    `json:"pasaran_totalmember"`
	Pasaran_totalbet         int    `json:"pasaran_totalbet"`
	Pasaran_totaloutstanding int    `json:"pasaran_totaloutstanding"`
	Pasaran_winlose          int    `json:"pasaran_winlose"`
	Pasaran_revisi           int    `json:"pasaran_revisi"`
	Pasaran_msgrevisi        string `json:"pasaran_msgrevisi"`
}
type responseredis_periodelistpasaranonline struct {
	Pasarancomp_idcompp int    `json:"pasarancomp_idcompp"`
	Pasarancomp_nama    string `json:"pasarancomp_nama"`
}

func PeriodeHome(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	field_redis := "LISTPERIODE_AGENT_" + client_company
	var obj responseredis_periodehome
	var arraobj []responseredis_periodehome
	var obj_pasaranonline responseredis_periodelistpasaranonline
	var arraobj_pasaranonline []responseredis_periodelistpasaranonline
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	pasaranonline_RD, _, _, _ := jsonparser.Get(jsonredis, "pasaranonline")

	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasaran_no, _ := jsonparser.GetInt(value, "pasaran_no")
		pasaran_invoice, _ := jsonparser.GetInt(value, "pasaran_invoice")
		pasaran_idcompp, _ := jsonparser.GetInt(value, "pasaran_idcompp")
		pasaran_code, _ := jsonparser.GetString(value, "pasaran_code")
		pasaran_periode, _ := jsonparser.GetString(value, "pasaran_periode")
		pasaran_name, _ := jsonparser.GetString(value, "pasaran_name")
		pasaran_tanggal, _ := jsonparser.GetString(value, "pasaran_tanggal")
		pasaran_keluaran, _ := jsonparser.GetString(value, "pasaran_keluaran")
		pasaran_status, _ := jsonparser.GetString(value, "pasaran_status")
		pasaran_status_css, _ := jsonparser.GetString(value, "pasaran_status_css")
		pasaran_totalmember, _ := jsonparser.GetInt(value, "pasaran_totalmember")
		pasaran_totalbet, _ := jsonparser.GetInt(value, "pasaran_totalbet")
		pasaran_totaloutstanding, _ := jsonparser.GetInt(value, "pasaran_totaloutstanding")
		pasaran_winlose, _ := jsonparser.GetInt(value, "pasaran_winlose")
		pasaran_revisi, _ := jsonparser.GetInt(value, "pasaran_revisi")
		pasaran_msgrevisi, _ := jsonparser.GetString(value, "pasaran_msgrevisi")

		obj.Pasaran_no = int(pasaran_no)
		obj.Pasaran_invoice = int(pasaran_invoice)
		obj.Pasaran_idcompp = int(pasaran_idcompp)
		obj.Pasaran_periode = pasaran_periode
		obj.Pasaran_code = pasaran_code
		obj.Pasaran_name = pasaran_name
		obj.Pasaran_tanggal = pasaran_tanggal
		obj.Pasaran_keluaran = pasaran_keluaran
		obj.Pasaran_status = pasaran_status
		obj.Pasaran_status_css = pasaran_status_css
		obj.Pasaran_totalmember = int(pasaran_totalmember)
		obj.Pasaran_totalbet = int(pasaran_totalbet)
		obj.Pasaran_totaloutstanding = int(pasaran_totaloutstanding)
		obj.Pasaran_winlose = int(pasaran_winlose)
		obj.Pasaran_revisi = int(pasaran_revisi)
		obj.Pasaran_msgrevisi = pasaran_msgrevisi
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(pasaranonline_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasarancomp_idcompp, _ := jsonparser.GetInt(value, "pasarancomp_idcompp")
		pasarancomp_nama, _ := jsonparser.GetString(value, "pasarancomp_nama")

		obj_pasaranonline.Pasarancomp_idcompp = int(pasarancomp_idcompp)
		obj_pasaranonline.Pasarancomp_nama = pasarancomp_nama
		arraobj_pasaranonline = append(arraobj_pasaranonline, obj_pasaranonline)
	})
	if !flag {
		result, err := models.Fetch_periode(client_company)
		helpers.SetRedis(field_redis, result, 0)
		log.Println("MYSQL")
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
		log.Println("cache")
		return c.JSON(fiber.Map{
			"status":        fiber.StatusOK,
			"message":       "Success",
			"record":        arraobj,
			"pasaranonline": arraobj_pasaranonline,
			"time":          time.Since(render_page).String(),
		})
	}
}
func PeriodeDetail(c *fiber.Ctx) error {
	client := new(periodedetail)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	result, err := models.Fetch_periodedetail(client_company, client.Idtrxkeluaran)
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
func PeriodeListMember(c *fiber.Ctx) error {
	client := new(periodedetail)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	result, err := models.Fetch_membergroup(client_company, client.Idtrxkeluaran)
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
func PeriodeListBet(c *fiber.Ctx) error {
	client := new(periodedetail)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	result, err := models.Fetch_listbet(client_company, client.Permainan, client.Idtrxkeluaran)
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
func PeriodeListBetstatus(c *fiber.Ctx) error {
	client := new(periodedetailstatus)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	result, err := models.Fetch_listbetbystatus(client_company, client.Status, client.Idtrxkeluaran)
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
func PeriodeListBetTable(c *fiber.Ctx) error {
	client := new(periodedetail)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	result, err := models.Fetch_listbettable(client_company, client.Idtrxkeluaran)
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
func PeriodeBetTable(c *fiber.Ctx) error {
	client := new(periodedetail)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	result, err := models.Fetch_bettable(client_company, client.Permainan, client.Idtrxkeluaran)
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
func PeriodeListMemberByNomor(c *fiber.Ctx) error {
	client := new(periodedetailmembernomor)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	result, err := models.Fetch_membergroupbynomor(client_company, client.Permainan, client.Nomor, client.Idtrxkeluaran)
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
func PeriodeSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodeSave)
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
		result, err := models.Save_Periode(
			client_username,
			client_company,
			client.Idtrxkeluaran, client.Nomorkeluaran)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("listresult_" + client_company + "_" + client.Idpasarantogel)
		val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + client_company)
		val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + client_company)
		val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + client_company)
		val_frontend_listresult := helpers.DeleteRedis("listresult_" + client_company)
		log.Printf("Redis Delete status: %d", val)
		log.Printf("Redis Delete Agent status: %d", val_agent)
		log.Printf("Redis Delete Agent DASHBOARD status: %d", val_agent_dashboard)
		log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
		log.Printf("Redis Delete FRONTEND LISTRESULT status: %d", val_frontend_listresult)
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
			result, err := models.Save_Periode(
				client_username,
				client_company,
				client.Idtrxkeluaran, client.Nomorkeluaran)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("listresult_" + client_company + "_" + client.Idpasarantogel)
			val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + client_company)
			val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + client_company)
			val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + client_company)
			val_frontend_listresult := helpers.DeleteRedis("listresult_" + client_company)
			log.Printf("Redis Delete status: %d", val)
			log.Printf("Redis Delete Agent status: %d", val_agent)
			log.Printf("Redis Delete Agent DASHBOARD status: %d", val_agent_dashboard)
			log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
			log.Printf("Redis Delete FRONTEND LISTRESULT status: %d", val_frontend_listresult)
			return c.JSON(result)
		}
	}
}
func PeriodeSaveNew(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodeSaveNew)
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
	idpasarantogel := models.Get_CompanyPasaran(client_company, "idpasarantogel", client.Idcomppasaran)
	log.Println(idpasarantogel)
	if typeadmin == "MASTER" {
		result, err := models.Save_PeriodeNew(client_username, client_company, client.Idcomppasaran)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("listresult_" + client_company + "_" + idpasarantogel)
		val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + client_company)
		val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + client_company)
		log.Printf("Redis Delete status: %d", val)
		log.Printf("Redis Delete Agent status: %d", val_agent)
		log.Printf("Redis Delete Agent DASHBOARD status: %d", val_agent_dashboard)
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
			result, err := models.Save_PeriodeNew(client_username, client_company, client.Idcomppasaran)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("listresult_" + client_company + "_" + idpasarantogel)
			val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + client_company)
			val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + client_company)
			log.Printf("Redis Delete status: %d", val)
			log.Printf("Redis Delete Agent status: %d", val_agent)
			log.Printf("Redis Delete Agent DASHBOARD status: %d", val_agent_dashboard)
			return c.JSON(result)
		}
	}
}
func PeriodeSaveRevisi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodesaverevisi)
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
	idcomppasaran := models.Get_Trxkeluaran(client_company, "idpasarantogel", client.Idinvoice)
	idcomppasaran_int, _ := strconv.Atoi(idcomppasaran)
	idpasarantogel := models.Get_CompanyPasaran(client_company, "idpasarantogel", idcomppasaran_int)

	if typeadmin == "MASTER" {
		result, err := models.Save_PeriodeRevisi(client_username, client_company, client.Msgrevisi, client.Idinvoice)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("listresult_" + client_company + "_" + idpasarantogel)
		val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + client_company)
		val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + client_company)
		log.Printf("Redis Delete status: %d", val)
		log.Printf("Redis Delete Agent status: %d", val_agent)
		log.Printf("Redis Delete Agent DASHBOARD status: %d", val_agent_dashboard)
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
			result, err := models.Save_PeriodeRevisi(client_username, client_company, client.Msgrevisi, client.Idinvoice)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("listresult_" + client_company + "_" + idpasarantogel)
			val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + client_company)
			val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + client_company)
			log.Printf("Redis Delete status: %d", val)
			log.Printf("Redis Delete Agent status: %d", val_agent)
			log.Printf("Redis Delete Agent DASHBOARD status: %d", val_agent_dashboard)
			return c.JSON(result)
		}
	}
}
func PeriodeCancelBet(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodecancelbet)
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
	idcomppasaran := models.Get_Trxkeluaran(client_company, "idpasarantogel", client.Idinvoice)
	idcomppasaran_int, _ := strconv.Atoi(idcomppasaran)
	idpasarantogel := models.Get_CompanyPasaran(client_company, "idpasarantogel", idcomppasaran_int)

	if typeadmin == "MASTER" {
		result, err := models.Cancelbet_Periode(client_username, client_company, client.Idinvoice, client.Idinvoicedetail)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("listresult_" + client_company + "_" + idpasarantogel)
		val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + client_company)
		val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + client_company)
		log.Printf("Redis Delete status: %d", val)
		log.Printf("Redis Delete Agent status: %d", val_agent)
		log.Printf("Redis Delete Agent DASHBOARD status: %d", val_agent_dashboard)
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
			result, err := models.Cancelbet_Periode(client_username, client_company, client.Idinvoice, client.Idinvoicedetail)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("listresult_" + client_company + "_" + idpasarantogel)
			val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + client_company)
			val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + client_company)
			log.Printf("Redis Delete status: %d", val)
			log.Printf("Redis Delete Agent status: %d", val_agent)
			log.Printf("Redis Delete Agent DASHBOARD status: %d", val_agent_dashboard)
			return c.JSON(result)
		}
	}
}
func Periodelistpasaran(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	result, err := models.Fetch_listpasaran(client_company)
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
func Periodelistprediksi(c *fiber.Ctx) error {
	client := new(periodePrediksi)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	result, err := models.Fetch_listprediksi(client_company, client.Nomorkeluaran, client.Idcomppasaran)
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
