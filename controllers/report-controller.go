package controllers

import (
	"log"
	"math"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"bitbucket.org/isbtotogroup/apibackend_go/models"
	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type winlose struct {
	Client_start string `json:"client_start" validate:"required"`
	Client_end   string `json:"client_end" validate:"required"`
}
type responseredis_winlose struct {
	Report_client_username string  `json:"report_client_username"`
	Report_client_turnover float64 `json:"report_client_turnover"`
	Report_client_winlose  int     `json:"report_client_winlose"`
	Report_agent_winlose   float64 `json:"report_agent_winlose"`
}

func Reportwinlose(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(winlose)
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
	field_redis := "REPORTWINLOSE_AGENT_" + strings.ToLower(client_company) + "_" + client.Client_start + "_" + client.Client_end
	render_page := time.Now()
	var obj responseredis_winlose
	var arraobj []responseredis_winlose
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	subtotalturnover_RD, _ := jsonparser.GetInt(jsonredis, "subtotalturnover")
	Subtotalwinlose_RD, _ := jsonparser.GetInt(jsonredis, "Subtotalwinlose")
	Subtotalwinlosecompany_RD, _ := jsonparser.GetInt(jsonredis, "Subtotalwinlosecompany")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		report_client_username, _ := jsonparser.GetString(value, "report_client_username")
		report_client_winlose, _ := jsonparser.GetFloat(value, "report_client_turnover")
		report_client_turnover, _ := jsonparser.GetInt(value, "report_client_winlose")
		report_agent_winlose, _ := jsonparser.GetFloat(value, "report_agent_winlose")

		obj.Report_client_username = report_client_username
		obj.Report_client_turnover = float64(report_client_winlose)
		obj.Report_client_winlose = int(report_client_turnover)
		obj.Report_agent_winlose = math.Abs(float64(report_agent_winlose))
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_winlose(client_company, client.Client_start, client.Client_end)

		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(field_redis, result, 10*time.Minute)
		log.Println("REPORTWINLOSE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("REPORTWINLOSE CACHE")
		return c.JSON(fiber.Map{
			"status":                 fiber.StatusOK,
			"message":                "Success",
			"record":                 arraobj,
			"subtotalturnover":       int(subtotalturnover_RD),
			"Subtotalwinlose":        int(Subtotalwinlose_RD),
			"Subtotalwinlosecompany": int(Subtotalwinlosecompany_RD),
			"time":                   time.Since(render_page).String(),
		})
	}
}
