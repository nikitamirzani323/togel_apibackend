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

	var obj entities.Model_dashboardwinlose_parent
	var arraobj []entities.Model_dashboardwinlose_parent
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fielddashboard_redis + "_" + strings.ToLower(client_company) + "_" + client.Year)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		dashboardwinlose_nmagen, _ := jsonparser.GetString(value, "dashboardwinlose_nmagen")
		child_RD, _, _, _ := jsonparser.Get(value, "dashboardwinlose_detail")

		var obj_child entities.Model_dashboardwinlose_child
		var arraobj_child []entities.Model_dashboardwinlose_child
		jsonparser.ArrayEach(child_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			dashboardwinlose_winlose, _ := jsonparser.GetInt(value, "dashboardwinlose_winlose")
			obj_child.Dashboardwinlose_winlose = int(dashboardwinlose_winlose)
			arraobj_child = append(arraobj_child, obj_child)
		})

		obj.Dashboardwinlose_nmagen = dashboardwinlose_nmagen
		obj.Dashboardwinlose_detail = arraobj_child
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_dashboardwinlose(client_company, client.Year)
		helpers.SetRedis(Fielddashboard_redis+"_"+strings.ToLower(client_company)+"_"+client.Year, result, time.Minute*30)
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
func DashboardWinlosepasaran(c *fiber.Ctx) error {
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

	var obj entities.Model_dashboardagenpasaranwinlose_parent
	var arraobj []entities.Model_dashboardagenpasaranwinlose_parent
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fielddashboard_redis + "_pasaran_" + strings.ToLower(client_company) + "_" + client.Year)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		dashboardagenpasaran_nmpasaran, _ := jsonparser.GetString(value, "dashboardagenpasaran_nmpasaran")
		child_RD, _, _, _ := jsonparser.Get(value, "dashboardagenpasaran_detail")
		var obj_child entities.Model_dashboardagenpasaranwinlose_child
		var arraobj_child []entities.Model_dashboardagenpasaranwinlose_child
		jsonparser.ArrayEach(child_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			dashboardagenpasaran_winlose, _ := jsonparser.GetInt(value, "dashboardagenpasaran_winlose")
			obj_child.Dashboardagenpasaran_winlose = int(dashboardagenpasaran_winlose)
			arraobj_child = append(arraobj_child, obj_child)
		})

		obj.Dashboardagenpasaran_nmpasaran = dashboardagenpasaran_nmpasaran
		obj.Dashboardagenpasaran_detail = arraobj_child
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_dashboard(client_company, client.Year)
		helpers.SetRedis(Fielddashboard_redis+"_pasaran_"+strings.ToLower(client_company)+"_"+client.Year, result, time.Minute*30)
		log.Println("DASHBOARD WINLOSE PASARAN MYSQL")
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
		log.Println("DASHBOARD WINLOSE PASARAN CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
