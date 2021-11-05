package controllers

import (
	"log"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"bitbucket.org/isbtotogroup/apibackend_go/models"
	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type adminruledetail struct {
	Idrule int `json:"idrule" validate:"required"`
}
type adminrulesavedetail struct {
	Idrule int    `json:"idrule" validate:"required"`
	Page   string `json:"page"`
	Sdata  string `json:"sData" validate:"required"`
	Nama   string `json:"nama" validate:"required"`
}
type adminrulesaveconf struct {
	Idrule int    `json:"idrule" validate:"required"`
	Page   string `json:"page"`
	Sdata  string `json:"sData" validate:"required"`
	Rule   string `json:"rule" validate:"required"`
}
type responseredis_adminrulehome struct {
	Adminrule_no   int    `json:"adminrule_no"`
	Adminrule_id   int    `json:"adminrule_id"`
	Adminrule_nama string `json:"adminrule_nama"`
}

func AdminruleHome(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	field_redis := "LISTADMINRULE_AGENT_" + strings.ToLower(client_company)
	var obj responseredis_adminrulehome
	var arraobj []responseredis_adminrulehome
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		adminrule_no, _ := jsonparser.GetInt(value, "adminrule_no")
		adminrule_id, _ := jsonparser.GetInt(value, "adminrule_id")
		adminrule_nama, _ := jsonparser.GetString(value, "adminrule_nama")

		obj.Adminrule_no = int(adminrule_no)
		obj.Adminrule_id = int(adminrule_id)
		obj.Adminrule_nama = adminrule_nama
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_adminruleHome(client_company)
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
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func AdminruleDetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(adminruledetail)
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
	result, err := models.Fetch_adminruleDetail(client_company, client.Idrule)
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
func SaveAdminruleDetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(adminrulesavedetail)
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
		result, err := models.Save_Adminrule(client_username, client_company, client.Sdata, client.Nama, client.Idrule)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		field_redis := "LISTADMINRULE_AGENT_" + strings.ToLower(client_company)
		val_agent := helpers.DeleteRedis(field_redis)
		log.Printf("Redis Delete AGEN - ADMINRULE status: %d", val_agent)
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
			result, err := models.Save_Adminrule(client_username, client_company, client.Sdata, client.Nama, client.Idrule)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			field_redis := "LISTADMINRULE_AGENT_" + strings.ToLower(client_company)
			val_agent := helpers.DeleteRedis(field_redis)
			log.Printf("Redis Delete AGEN - ADMINRULE status: %d", val_agent)
			return c.JSON(result)
		}
	}
}
func SaveAdminruleConf(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(adminrulesaveconf)
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
		result, err := models.Save_Adminruleconf(client_username, client_company, client.Sdata, client.Rule, client.Idrule)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		field_redis := "LISTADMINRULE_AGENT_" + strings.ToLower(client_company)
		val_agent := helpers.DeleteRedis(field_redis)
		log.Printf("Redis Delete AGEN - ADMINRULE status: %d", val_agent)
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
			result, err := models.Save_Adminruleconf(client_username, client_company, client.Sdata, client.Rule, client.Idrule)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			field_redis := "LISTADMINRULE_AGENT_" + strings.ToLower(client_company)
			val_agent := helpers.DeleteRedis(field_redis)
			log.Printf("Redis Delete AGEN - ADMINRULE status: %d", val_agent)
			return c.JSON(result)
		}
	}
}
