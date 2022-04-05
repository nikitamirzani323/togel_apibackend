package controllers

import (
	"log"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/entities"
	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"bitbucket.org/isbtotogroup/apibackend_go/models"
	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type responseredis_adminhome struct {
	Admin_no           int    `json:"admin_no"`
	Admin_username     string `json:"admin_username"`
	Admin_nama         string `json:"admin_nama"`
	Admin_tipe         string `json:"admin_tipe"`
	Admin_rule         string `json:"admin_rule"`
	Admin_joindate     string `json:"admin_joindate"`
	Admin_timezone     string `json:"admin_timezone"`
	Admin_lastlogin    string `json:"admin_lastlogin"`
	Admin_lastipaddres string `json:"admin_lastipaddres"`
	Admin_status       string `json:"admin_status"`
}
type responseredis_adminhome_listruleadmin struct {
	Adminrule_idruleadmin int    `json:"adminrule_idruleadmin"`
	Adminrule_name        string `json:"adminrule_name"`
}

const Fieldadmin_home_redis = "LISTADMIN_AGENT_"

func AdminHome(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	field_redis := Fieldadmin_home_redis + strings.ToLower(client_company)
	var obj entities.Model_admin
	var arraobj []entities.Model_admin
	var obj_listruleadmin responseredis_adminhome_listruleadmin
	var arraobj_listruleadmin []responseredis_adminhome_listruleadmin
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listruleadmin_RD, _, _, _ := jsonparser.Get(jsonredis, "listruleadmin")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		admin_no, _ := jsonparser.GetInt(value, "admin_no")
		admin_username, _ := jsonparser.GetString(value, "admin_username")
		admin_nama, _ := jsonparser.GetString(value, "admin_nama")
		admin_tipe, _ := jsonparser.GetString(value, "admin_tipe")
		admin_rule, _ := jsonparser.GetString(value, "admin_rule")
		admin_joindate, _ := jsonparser.GetString(value, "admin_joindate")
		admin_timezone, _ := jsonparser.GetString(value, "admin_timezone")
		admin_lastlogin, _ := jsonparser.GetString(value, "admin_lastlogin")
		admin_lastipaddres, _ := jsonparser.GetString(value, "admin_lastipaddres")
		admin_status, _ := jsonparser.GetString(value, "admin_status")

		obj.No = int(admin_no)
		obj.Username = admin_username
		obj.Nama = admin_nama
		obj.Tipeadmin = admin_tipe
		obj.Rule = admin_rule
		obj.Joindate = admin_joindate
		obj.Timezone = admin_timezone
		obj.Lastlogin = admin_lastlogin
		obj.LastIpaddress = admin_lastipaddres
		obj.Status = admin_status
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listruleadmin_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		adminrule_idruleadmin, _ := jsonparser.GetInt(value, "adminrule_idruleadmin")
		adminrule_name, _ := jsonparser.GetString(value, "adminrule_name")

		obj_listruleadmin.Adminrule_idruleadmin = int(adminrule_idruleadmin)
		obj_listruleadmin.Adminrule_name = adminrule_name
		arraobj_listruleadmin = append(arraobj_listruleadmin, obj_listruleadmin)
	})
	if !flag {
		result, err := models.Fetch_adminHome(client_company)
		helpers.SetRedis(field_redis, result, 30*time.Minute)
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
			"listruleadmin": arraobj_listruleadmin,
			"time":          time.Since(render_page).String(),
		})
	}
}
func AdminDetail(c *fiber.Ctx) error {
	type admindetail struct {
		Username string `json:"username" validate:"required"`
	}
	var errors []*helpers.ErrorResponse
	client := new(admindetail)
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
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	result, err := models.Fetch_adminDetail(client_company, client.Username)
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
func AdminSave(c *fiber.Ctx) error {
	type adminsave struct {
		Sdata       string `json:"sdata" validate:"required"`
		Page        string `json:"page"`
		Idruleadmin int    `json:"idruleadmin" `
		Username    string `json:"username" validate:"required,alphanum,max=20"`
		Password    string `json:"password" `
		Name        string `json:"nama" validate:"required,alphanum,max=70"`
		Status      string `json:"status" validate:"required,alpha"`
	}
	var errors []*helpers.ErrorResponse
	client := new(adminsave)
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
		result, err := models.Save_Admin(
			client_username,
			client_company,
			client.Sdata,
			client.Username,
			client.Password,
			client.Name, client.Status, client.Idruleadmin)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		_deleteredis_admin(client_company)
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
			result, err := models.Save_Admin(
				client_username,
				client_company,
				client.Sdata,
				client.Username,
				client.Password,
				client.Name, client.Status, client.Idruleadmin)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			_deleteredis_admin(client_company)
			return c.JSON(result)
		}
	}
}
func AdminSaveIplist(c *fiber.Ctx) error {
	type adminsaveiplist struct {
		Sdata     string `json:"sdata" validate:"required"`
		Page      string `json:"page"`
		Username  string `json:"username" validate:"required,alphanum,max=20"`
		Ipaddress string `json:"ipaddress" validate:"required,max=20"`
	}
	var errors []*helpers.ErrorResponse
	client := new(adminsaveiplist)
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
		result, err := models.Save_AdminIplist(
			client_username,
			client_company,
			client.Sdata,
			client.Username,
			client.Ipaddress)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		_deleteredis_admin(client_company)
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
			result, err := models.Save_AdminIplist(
				client_username,
				client_company,
				client.Sdata,
				client.Username,
				client.Ipaddress)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			_deleteredis_admin(client_company)
			return c.JSON(result)
		}
	}
}
func AdminDeleteIplist(c *fiber.Ctx) error {
	type admindeleteiplist struct {
		Idcompiplist int    `json:"idcompiplist" validate:"required"`
		Username     string `json:"username" validate:"required,alphanum,max=20"`
		Page         string `json:"page"`
	}
	var errors []*helpers.ErrorResponse
	client := new(admindeleteiplist)
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
	_, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Delete_AdminIplist(client.Username, client.Idcompiplist)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		_deleteredis_admin(client_company)
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
			result, err := models.Delete_AdminIplist(client.Username, client.Idcompiplist)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			_deleteredis_admin(client_company)
			return c.JSON(result)
		}
	}
}
func _deleteredis_admin(company string) {
	field_redis := Fieldadmin_home_redis + strings.ToLower(company)
	val_adminagent := helpers.DeleteRedis(field_redis)
	log.Printf("Redis Delete AGEN - ADMIN status: %d", val_adminagent)
	log_redis := "LISTLOG_AGENT_" + strings.ToLower(company)
	val_agent_redis := helpers.DeleteRedis(log_redis)
	log.Printf("Redis Delete AGEN - LOG status: %d", val_agent_redis)
}
