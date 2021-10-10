package controllers

import (
	"log"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"bitbucket.org/isbtotogroup/apibackend_go/models"
	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type admindetail struct {
	Username string `json:"username" validate:"required"`
}
type adminsave struct {
	Sdata       string `json:"sdata" validate:"required"`
	Page        string `json:"page"`
	Idruleadmin int    `json:"idruleadmin" `
	Username    string `json:"username" validate:"required,alphanum,max=20"`
	Password    string `json:"password" `
	Name        string `json:"nama" validate:"required,alphanum,max=70"`
	Status      string `json:"status" validate:"required,alpha"`
}
type adminsaveiplist struct {
	Sdata     string `json:"sdata" validate:"required"`
	Page      string `json:"page"`
	Username  string `json:"username" validate:"required,alphanum,max=20"`
	Ipaddress string `json:"ipaddress" validate:"required,max=20"`
}
type admindeleteiplist struct {
	Idcompiplist int    `json:"idcompiplist" validate:"required"`
	Username     string `json:"username" validate:"required,alphanum,max=20"`
	Page         string `json:"page"`
}
type responseredis_adminhome struct {
	Admin_no           int    `json:"admin_no"`
	Admin_username     string `json:"admin_username"`
	Admin_nama         string `json:"admin_nama"`
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

func AdminHome(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	field_redis := "LISTADMINM_AGENT_" + client_company
	var obj responseredis_adminhome
	var arraobj []responseredis_adminhome
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
		admin_rule, _ := jsonparser.GetString(value, "admin_rule")
		admin_joindate, _ := jsonparser.GetString(value, "admin_joindate")
		admin_timezone, _ := jsonparser.GetString(value, "admin_timezone")
		admin_lastlogin, _ := jsonparser.GetString(value, "admin_lastlogin")
		admin_lastipaddres, _ := jsonparser.GetString(value, "admin_lastipaddres")
		admin_status, _ := jsonparser.GetString(value, "admin_status")

		obj.Admin_no = int(admin_no)
		obj.Admin_username = admin_username
		obj.Admin_nama = admin_nama
		obj.Admin_rule = admin_rule
		obj.Admin_joindate = admin_joindate
		obj.Admin_timezone = admin_timezone
		obj.Admin_lastlogin = admin_lastlogin
		obj.Admin_lastipaddres = admin_lastipaddres
		obj.Admin_status = admin_status
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
		helpers.SetRedis(field_redis, result, 5*time.Minute)
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
			return c.JSON(result)
		}
	}
}
func AdminSaveIplist(c *fiber.Ctx) error {
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
			return c.JSON(result)
		}
	}
}
func AdminDeleteIplist(c *fiber.Ctx) error {
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
			return c.JSON(result)
		}
	}
}
