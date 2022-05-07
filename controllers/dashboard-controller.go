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

type responseredis_dashboardchart struct {
	Pasaran_name   string      `json:"pasaran_name"`
	Pasaran_detial interface{} `json:"pasaran_detial"`
}
type responseredis_dashboardchartchild struct {
	Pasaranwinlose int `json:"Pasaranwinlose"`
}

const Fielddashboard_redis = "DASHBOARD_CHART_AGEN_WINLOSE"

func DashboardWinlose(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_dashboardwinlose)
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

	var obj entities.Model_dashboardwinlose
	var arraobj []entities.Model_dashboardwinlose
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fielddashboard_redis + "_" + strings.ToLower(client_company) + "_" + client.Year)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		winlose, _ := jsonparser.GetInt(value, "winlose")

		obj.Winlose = int(winlose)
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_dashboardwinlose(client_company, client.Year)
		helpers.SetRedis(Fielddashboard_redis+"_"+strings.ToLower(client_company)+"_"+client.Year, result, time.Hour*24)
		log.Println("DASHBOARD WINLOSE MYSQL")
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
		log.Println("DASHBOARD WINLOSE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func DashboardHome(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")

	field_redis := "DASHBOARD_CHART_AGENT_" + strings.ToLower(client_company)
	var obj responseredis_dashboardchart
	var arraobj []responseredis_dashboardchart
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasaran_name, _ := jsonparser.GetString(value, "pasaran_name")
		child_RD, _, _, _ := jsonparser.Get(value, "pasaran_detial")
		var obj_child responseredis_dashboardchartchild
		var arraobj_child []responseredis_dashboardchartchild
		jsonparser.ArrayEach(child_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			Pasaranwinlose, _ := jsonparser.GetInt(value, "Pasaranwinlose")
			obj_child.Pasaranwinlose = int(Pasaranwinlose)
			arraobj_child = append(arraobj_child, obj_child)
		})

		obj.Pasaran_name = pasaran_name
		obj.Pasaran_detial = arraobj_child
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_dashboard(client_company)
		helpers.SetRedis(field_redis, result, time.Hour*24)
		log.Println("DASHBOARD MYSQL")
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
		log.Println("DASHBOARD CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
