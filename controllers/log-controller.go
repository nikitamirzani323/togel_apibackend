package controllers

import (
	"log"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/entities"
	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"bitbucket.org/isbtotogroup/apibackend_go/models"
	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const REDIS_LOGHOME = "LISTLOG_AGENT"

func LogHome(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")

	var obj entities.Model_log
	var arraobj []entities.Model_log
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(REDIS_LOGHOME + "_" + strings.ToLower(client_company))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		log_id, _ := jsonparser.GetInt(value, "log_id")
		log_datetime, _ := jsonparser.GetString(value, "log_datetime")
		log_username, _ := jsonparser.GetString(value, "log_username")
		log_page, _ := jsonparser.GetString(value, "log_page")
		log_tipe, _ := jsonparser.GetString(value, "log_tipe")
		log_note, _ := jsonparser.GetString(value, "log_note")

		obj.Log_id = int(log_id)
		obj.Log_datetime = log_datetime
		obj.Log_username = log_username
		obj.Log_page = log_page
		obj.Log_tipe = log_tipe
		obj.Log_note = log_note
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_loghome(client_company)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(REDIS_LOGHOME+"_"+strings.ToLower(client_company), result, time.Hour*24)
		log.Println("LOG MYSQL")
		return c.JSON(result)
	} else {
		log.Println("LOG CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
