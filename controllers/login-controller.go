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

type Login struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Ipaddress string `json:"ipaddress" validate:"required"`
	Timezone  string `json:"timezone" validate:"required"`
}
type Gpassword struct {
	Password string `json:"password"`
}
type home struct {
	Page string `json:"page"`
}

func CheckLogin(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(Login)
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

	result, idcomp, typeadmin, ruleadmin, err := models.Login_Model(client.Username, client.Password, client.Ipaddress, client.Timezone)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	if !result {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": "Username or Password Not Found",
			})

	} else {
		dataclient := client.Username + "==" + idcomp + "==" + typeadmin + "==" + strconv.Itoa(ruleadmin)
		dataclient_encr, keymap := helpers.Encryption(dataclient)
		dataclient_encr_final := dataclient_encr + "|" + strconv.Itoa(keymap)
		t, err := helpers.GenerateNewAccessToken(dataclient_encr_final)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		_deleteredis_login(idcomp)
		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
			"token":  t,
		})

	}
}
func Home(c *fiber.Ctx) error {
	client := new(home)
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
	_, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  nil,
		})
	} else {
		if !flag {
			c.Status(fiber.StatusForbidden)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusForbidden,
				"message": "Anda tidak bisa akses halaman ini",
				"record":  nil,
			})
		} else {
			c.Status(fiber.StatusOK)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusOK,
				"message": "ADMIN",
				"record":  nil,
			})
		}
	}
}
func GenerateHashPassword(c *fiber.Ctx) error {
	client := new(Gpassword)

	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	hash := helpers.HashPasswordMD5(client.Password)

	return c.JSON(hash)
}
func _deleteredis_login(idcomp string) {
	log_admin := "LISTADMIN_AGENT_" + strings.ToLower(idcomp)
	val_agent_admin := helpers.DeleteRedis(log_admin)
	log.Printf("Redis Delete ADMIN status: %d", val_agent_admin)
	log_redis := "LISTLOG_AGENT_" + strings.ToLower(idcomp)
	val_agent_redis := helpers.DeleteRedis(log_redis)
	log.Printf("Redis Delete LOG status: %d", val_agent_redis)
}
