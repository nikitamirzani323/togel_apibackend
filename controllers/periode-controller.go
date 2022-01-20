package controllers

import (
	"log"
	"strconv"
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

type periodedetail struct {
	Idtrxkeluaran int    `json:"idinvoice"`
	Permainan     string `json:"permainan"`
}
type periodedetailstatus struct {
	Idtrxkeluaran int    `json:"idinvoice"`
	Status        string `json:"status"`
}
type periodedetailusername struct {
	Idtrxkeluaran int    `json:"idinvoice"`
	Username      string `json:"username"`
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

type periodeSavePrediksi struct {
	Sdata     string `json:"sData" validate:"required"`
	Page      string `json:"page"`
	Idinvoice int    `json:"idinvoice" validate:"required"`
}
type periodePrediksi struct {
	Nomorkeluaran string `json:"nomorkeluaran" validate:"required,min=4,max=4"`
	Idcomppasaran int    `json:"Idcomppasaran" validate:"required"`
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
	Pasaran_totalcancelbet   int    `json:"pasaran_totalcancelbet"`
	Pasaran_winlose          int    `json:"pasaran_winlose"`
	Pasaran_revisi           int    `json:"pasaran_revisi"`
	Pasaran_msgrevisi        string `json:"pasaran_msgrevisi"`
}
type responseredis_periodedetail struct {
	Idinvoice          string `json:"idinvoice"`
	TanggalPeriode     string `json:"periode_tanggalkeluaran"`
	TanggalNext        string `json:"periode_tanggalnext"`
	PeriodeKeluaran    string `json:"periode_keluaranperiode"`
	Keluaran           string `json:"periode_keluaran"`
	Statusrevisi       string `json:"periode_statusrevisi"`
	StatusOnlineOffice string `json:"periode_statusonline"`
	Create             string `json:"periode_create"`
	CreateDate         string `json:"periode_createdate"`
	Update             string `json:"periode_update"`
	UpdateDate         string `json:"periode_updatedate"`
}
type responseredis_periodelistmember struct {
	Member         string `json:"member"`
	Totalbet       int    `json:"totalbet"`
	Totalbayar     int    `json:"totalbayar"`
	Totalcancelbet int    `json:"totalcancelbet"`
	Totalwin       int    `json:"totalwin"`
}
type responseredis_periodelistbettable struct {
	Permainan string `json:"permainan"`
}
type responseredis_periodelistbet struct {
	Bet_id           int     `json:"bet_id"`
	Bet_datetime     string  `json:"bet_datetime"`
	Bet_ipaddress    string  `json:"bet_ipaddress"`
	Bet_device       string  `json:"bet_device"`
	Bet_timezone     string  `json:"bet_timezone"`
	Bet_username     string  `json:"bet_username"`
	Bet_typegame     string  `json:"bet_typegame"`
	Bet_nomortogel   string  `json:"bet_nomortogel"`
	Bet_posisitogel  string  `json:"bet_posisitogel"`
	Bet_bet          int     `json:"bet_bet"`
	Bet_diskon       int     `json:"bet_diskon"`
	Bet_diskonpercen int     `json:"bet_diskonpercen"`
	Bet_kei          int     `json:"bet_kei"`
	Bet_keipercen    int     `json:"bet_keipercen"`
	Bet_win          float32 `json:"bet_win"`
	Bet_totalwin     int     `json:"bet_totalwin"`
	Bet_bayar        int     `json:"bet_bayar"`
	Bet_status       string  `json:"bet_status"`
	Bet_statuscss    string  `json:"bet_statuscss"`
	Bet_create       string  `json:"bet_create"`
	Bet_createDate   string  `json:"bet_createdate"`
	Bet_update       string  `json:"bet_update"`
	Bet_updateDate   string  `json:"bet_updatedate"`
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
	field_redis := "LISTPERIODE_AGENT_" + strings.ToLower(client_company)
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
		pasaran_totalcancelbet, _ := jsonparser.GetInt(value, "pasaran_totalcancelbet")
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
		obj.Pasaran_totalcancelbet = int(pasaran_totalcancelbet)
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
		helpers.SetRedis(field_redis, result, time.Hour*24)
		log.Println("PERIODE MYSQL")
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
		log.Println("PERIODE CACHE")
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

	field_redis := "LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran)
	render_page := time.Now()
	var obj responseredis_periodedetail
	var arraobj []responseredis_periodedetail
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		idinvoice, _ := jsonparser.GetString(value, "idinvoice")
		periode_tanggalkeluaran, _ := jsonparser.GetString(value, "periode_tanggalkeluaran")
		periode_tanggalnext, _ := jsonparser.GetString(value, "periode_tanggalnext")
		periode_keluaranperiode, _ := jsonparser.GetString(value, "periode_keluaranperiode")
		periode_keluaran, _ := jsonparser.GetString(value, "periode_keluaran")
		periode_statusrevisi, _ := jsonparser.GetString(value, "periode_statusrevisi")
		periode_statusonline, _ := jsonparser.GetString(value, "periode_statusonline")
		periode_create, _ := jsonparser.GetString(value, "periode_create")
		periode_createdate, _ := jsonparser.GetString(value, "periode_createdate")
		periode_update, _ := jsonparser.GetString(value, "periode_update")
		periode_updatedate, _ := jsonparser.GetString(value, "periode_updatedate")

		obj.Idinvoice = idinvoice
		obj.TanggalPeriode = periode_tanggalkeluaran
		obj.TanggalNext = periode_tanggalnext
		obj.PeriodeKeluaran = periode_keluaranperiode
		obj.Keluaran = periode_keluaran
		obj.Statusrevisi = periode_statusrevisi
		obj.StatusOnlineOffice = periode_statusonline
		obj.Create = periode_create
		obj.CreateDate = periode_createdate
		obj.Update = periode_update
		obj.UpdateDate = periode_updatedate
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_periodedetail(client_company, client.Idtrxkeluaran)
		helpers.SetRedis(field_redis, result, time.Minute*1)
		log.Println("PERIODE DETAIL MYSQL")
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
		log.Println("PERIODE DETAIL CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
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
	field_redis := "LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTMEMBER"
	render_page := time.Now()
	var obj responseredis_periodelistmember
	var arraobj []responseredis_periodelistmember
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		member, _ := jsonparser.GetString(value, "member")
		totalbet, _ := jsonparser.GetInt(value, "totalbet")
		totalbayar, _ := jsonparser.GetInt(value, "totalbayar")
		totalcancelbet, _ := jsonparser.GetInt(value, "totalcancelbet")
		totalwin, _ := jsonparser.GetInt(value, "totalwin")

		obj.Member = member
		obj.Totalbet = int(totalbet)
		obj.Totalbayar = int(totalbayar)
		obj.Totalcancelbet = int(totalcancelbet)
		obj.Totalwin = int(totalwin)
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_membergroup(client_company, client.Idtrxkeluaran)
		helpers.SetRedis(field_redis, result, time.Minute*30)
		log.Println("PERIODE DETAIL LIST MEMBER MYSQL")
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
		log.Println("PERIODE DETAIL LIST MEMBER  CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}

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
	field_redis := "LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_" + client.Permainan
	render_page := time.Now()
	var obj responseredis_periodelistbet
	var arraobj []responseredis_periodelistbet
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	totalbet_RD, _ := jsonparser.GetInt(jsonredis, "totalbet")
	subtotal_RD, _ := jsonparser.GetInt(jsonredis, "subtotal")
	subtotalwin_RD, _ := jsonparser.GetInt(jsonredis, "subtotalwin")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		bet_id, _ := jsonparser.GetInt(value, "bet_id")
		bet_datetime, _ := jsonparser.GetString(value, "bet_datetime")
		bet_ipaddress, _ := jsonparser.GetString(value, "bet_ipaddress")
		bet_device, _ := jsonparser.GetString(value, "bet_device")
		bet_timezone, _ := jsonparser.GetString(value, "bet_timezone")
		bet_username, _ := jsonparser.GetString(value, "bet_username")
		bet_typegame, _ := jsonparser.GetString(value, "bet_typegame")
		bet_nomortogel, _ := jsonparser.GetString(value, "bet_nomortogel")
		bet_posisitogel, _ := jsonparser.GetString(value, "bet_posisitogel")
		bet_bet, _ := jsonparser.GetInt(value, "bet_bet")
		bet_diskon, _ := jsonparser.GetInt(value, "bet_diskon")
		bet_diskonpercen, _ := jsonparser.GetInt(value, "bet_diskonpercen")
		bet_kei, _ := jsonparser.GetInt(value, "bet_kei")
		bet_keipercen, _ := jsonparser.GetInt(value, "bet_keipercen")
		bet_win, _ := jsonparser.GetFloat(value, "bet_win")
		bet_totalwin, _ := jsonparser.GetInt(value, "bet_totalwin")
		bet_bayar, _ := jsonparser.GetInt(value, "bet_bayar")
		bet_status, _ := jsonparser.GetString(value, "bet_status")
		bet_statuscss, _ := jsonparser.GetString(value, "bet_statuscss")
		bet_create, _ := jsonparser.GetString(value, "bet_create")
		bet_createdate, _ := jsonparser.GetString(value, "bet_createdate")
		bet_update, _ := jsonparser.GetString(value, "bet_update")
		bet_updatedate, _ := jsonparser.GetString(value, "bet_updatedate")

		obj.Bet_id = int(bet_id)
		obj.Bet_datetime = bet_datetime
		obj.Bet_ipaddress = bet_ipaddress
		obj.Bet_device = bet_device
		obj.Bet_timezone = bet_timezone
		obj.Bet_username = bet_username
		obj.Bet_typegame = bet_typegame
		obj.Bet_nomortogel = bet_nomortogel
		obj.Bet_posisitogel = bet_posisitogel
		obj.Bet_bet = int(bet_bet)
		obj.Bet_diskon = int(bet_diskon)
		obj.Bet_diskonpercen = int(bet_diskonpercen)
		obj.Bet_kei = int(bet_kei)
		obj.Bet_keipercen = int(bet_keipercen)
		obj.Bet_win = float32(bet_win)
		obj.Bet_totalwin = int(bet_totalwin)
		obj.Bet_bayar = int(bet_bayar)
		obj.Bet_status = bet_status
		obj.Bet_statuscss = bet_statuscss
		obj.Bet_create = bet_create
		obj.Bet_createDate = bet_createdate
		obj.Bet_update = bet_update
		obj.Bet_updateDate = bet_updatedate
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_listbet(client_company, client.Permainan, client.Idtrxkeluaran)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(field_redis, result, time.Minute*30)
		log.Println("PERIODE DETAIL LIST BET MYSQL " + client.Permainan)
		return c.JSON(result)
	} else {
		log.Println("PERIODE DETAIL LIST BET CACHE " + client.Permainan)
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"totalbet":    int(totalbet_RD),
			"subtotal":    int(subtotal_RD),
			"subtotalwin": int(subtotalwin_RD),
			"time":        time.Since(render_page).String(),
		})
	}
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
	field_redis := "LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTBET_" + client.Status
	render_page := time.Now()
	var obj responseredis_periodelistbet
	var arraobj []responseredis_periodelistbet
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	totalbet_RD, _ := jsonparser.GetInt(jsonredis, "totalbet")
	subtotal_RD, _ := jsonparser.GetInt(jsonredis, "subtotal")
	subtotalwin_RD, _ := jsonparser.GetInt(jsonredis, "subtotalwin")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		bet_id, _ := jsonparser.GetInt(value, "bet_id")
		bet_datetime, _ := jsonparser.GetString(value, "bet_datetime")
		bet_ipaddress, _ := jsonparser.GetString(value, "bet_ipaddress")
		bet_device, _ := jsonparser.GetString(value, "bet_device")
		bet_timezone, _ := jsonparser.GetString(value, "bet_timezone")
		bet_username, _ := jsonparser.GetString(value, "bet_username")
		bet_typegame, _ := jsonparser.GetString(value, "bet_typegame")
		bet_nomortogel, _ := jsonparser.GetString(value, "bet_nomortogel")
		bet_bet, _ := jsonparser.GetInt(value, "bet_bet")
		bet_diskon, _ := jsonparser.GetInt(value, "bet_diskon")
		bet_diskonpercen, _ := jsonparser.GetInt(value, "bet_diskonpercen")
		bet_kei, _ := jsonparser.GetInt(value, "bet_kei")
		bet_keipercen, _ := jsonparser.GetInt(value, "bet_keipercen")
		bet_win, _ := jsonparser.GetFloat(value, "bet_win")
		bet_totalwin, _ := jsonparser.GetInt(value, "bet_totalwin")
		bet_bayar, _ := jsonparser.GetInt(value, "bet_bayar")
		bet_status, _ := jsonparser.GetString(value, "bet_status")
		bet_statuscss, _ := jsonparser.GetString(value, "bet_statuscss")
		bet_create, _ := jsonparser.GetString(value, "bet_create")
		bet_createdate, _ := jsonparser.GetString(value, "bet_createdate")
		bet_update, _ := jsonparser.GetString(value, "bet_update")
		bet_updatedate, _ := jsonparser.GetString(value, "bet_updatedate")

		obj.Bet_id = int(bet_id)
		obj.Bet_datetime = bet_datetime
		obj.Bet_ipaddress = bet_ipaddress
		obj.Bet_device = bet_device
		obj.Bet_timezone = bet_timezone
		obj.Bet_username = bet_username
		obj.Bet_typegame = bet_typegame
		obj.Bet_nomortogel = bet_nomortogel
		obj.Bet_bet = int(bet_bet)
		obj.Bet_diskon = int(bet_diskon)
		obj.Bet_diskonpercen = int(bet_diskonpercen)
		obj.Bet_kei = int(bet_kei)
		obj.Bet_keipercen = int(bet_keipercen)
		obj.Bet_win = float32(bet_win)
		obj.Bet_totalwin = int(bet_totalwin)
		obj.Bet_bayar = int(bet_bayar)
		obj.Bet_status = bet_status
		obj.Bet_statuscss = bet_statuscss
		obj.Bet_create = bet_create
		obj.Bet_createDate = bet_createdate
		obj.Bet_update = bet_update
		obj.Bet_updateDate = bet_updatedate
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_listbetbystatus(client_company, client.Status, client.Idtrxkeluaran)
		helpers.SetRedis(field_redis, result, time.Minute*30)
		log.Println("PERIODE DETAIL LIST BET STATUS MYSQL " + client.Status)
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
		log.Println("PERIODE DETAIL LIST BET STATUS CACHE " + client.Status)
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"totalbet":    int(totalbet_RD),
			"subtotal":    int(subtotal_RD),
			"subtotalwin": int(subtotalwin_RD),
			"time":        time.Since(render_page).String(),
		})
	}
}
func PeriodeListBetusername(c *fiber.Ctx) error {
	client := new(periodedetailusername)
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
	result, err := models.Fetch_listbetbyusername(client_company, client.Username, client.Idtrxkeluaran)
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
	field_redis := "LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTBETTABLE"
	render_page := time.Now()
	var obj responseredis_periodelistbettable
	var arraobj []responseredis_periodelistbettable
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		permainan, _ := jsonparser.GetString(value, "permainan")

		obj.Permainan = permainan
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_listbettable(client_company, client.Idtrxkeluaran)
		helpers.SetRedis(field_redis, result, time.Minute*30)
		log.Println("PERIODE DETAIL LIST BET TABLE MYSQL")
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
		log.Println("PERIODE DETAIL LIST BET TABLE  CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
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
		//FRONTEND
		val_frontend := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel))
		val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
		val_frontend_listresult := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company))
		log.Printf("Redis Delete FRONTEND status: %d", val_frontend)
		log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
		log.Printf("Redis Delete FRONTEND LISTRESULT status: %d", val_frontend_listresult)
		//AGEN
		val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company))
		val_agent2 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran))
		val_agent3 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTMEMBER")
		val_agent4 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTBETTABLE")
		val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + strings.ToLower(client_company))
		log.Printf("Redis Delete AGENT PERIODE : %d", val_agent)
		log.Printf("Redis Delete AGENT PERIODE DETAIL: %d", val_agent2)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LISTMEMBER: %d", val_agent3)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LISTBETTABLE: %d", val_agent4)
		log.Printf("Redis Delete AGENT DASHBOARD status: %d", val_agent_dashboard)
		val_agent4d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_4D")
		val_agent3d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_3D")
		val_agent2d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_2D")
		val_agent2dd := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_2DD")
		val_agent2dt := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_2DT")
		val_agentcolokbebas := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_COLOK_BEBAS")
		val_agentcolokmacau := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_COLOK_MACAU")
		val_agentcoloknaga := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_COLOK_NAGA")
		val_agentcolokjitu := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_COLOK_JITU")
		val_agent5050umum := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_50_50_UMUM")
		val_agent5050special := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_50_50_SPECIAL")
		val_agent5050kombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_50_50_KOMBINASI")
		val_agentmacaukombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_MACAU_KOMBINASI")
		val_agentdasar := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_DASAR")
		val_agentshio := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_SHIO")
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
		val_agentall := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTBET_all")
		val_agentwinner := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTBET_winner")
		val_agentcancel := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTBET_cancel")
		log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS ALL: %d", val_agentall)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS WINNER: %d", val_agentwinner)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS CANCEL: %d", val_agentcancel)
		log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
		val_agent_redis := helpers.DeleteRedis(log_redis)
		log.Printf("Redis Delete LOG status: %d", val_agent_redis)
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
			//FRONTEND
			val_frontend := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel))
			val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
			val_frontend_listresult := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company))
			log.Printf("Redis Delete FRONTEND status: %d", val_frontend)
			log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
			log.Printf("Redis Delete FRONTEND LISTRESULT status: %d", val_frontend_listresult)
			//AGEN
			val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company))
			val_agent2 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran))
			val_agent3 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTMEMBER")
			val_agent4 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTBETTABLE")
			val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + strings.ToLower(client_company))
			log.Printf("Redis Delete AGENT PERIODE : %d", val_agent)
			log.Printf("Redis Delete AGENT PERIODE DETAIL: %d", val_agent2)
			log.Printf("Redis Delete AGENT PERIODE DETAIL LISTMEMBER: %d", val_agent3)
			log.Printf("Redis Delete AGENT PERIODE DETAIL LISTBETTABLE: %d", val_agent4)
			log.Printf("Redis Delete AGENT DASHBOARD status: %d", val_agent_dashboard)
			val_agent4d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_4D")
			val_agent3d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_3D")
			val_agent2d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_2D")
			val_agent2dd := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_2DD")
			val_agent2dt := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_2DT")
			val_agentcolokbebas := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_COLOK_BEBAS")
			val_agentcolokmacau := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_COLOK_MACAU")
			val_agentcoloknaga := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_COLOK_NAGA")
			val_agentcolokjitu := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_COLOK_JITU")
			val_agent5050umum := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_50_50_UMUM")
			val_agent5050special := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_50_50_SPECIAL")
			val_agent5050kombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_50_50_KOMBINASI")
			val_agentmacaukombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_MACAU_KOMBINASI")
			val_agentdasar := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_DASAR")
			val_agentshio := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTPERMAINAN_SHIO")
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
			val_agentall := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTBET_all")
			val_agentwinner := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTBET_winner")
			val_agentcancel := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idtrxkeluaran) + "_LISTBET_cancel")
			log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS ALL: %d", val_agentall)
			log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS WINNER: %d", val_agentwinner)
			log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS CANCEL: %d", val_agentcancel)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
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
		//FRONTEND
		val_frontend := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company) + "_" + strings.ToLower(idpasarantogel))
		val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
		val_frontend_listresult := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company))
		log.Printf("Redis Delete FRONTEND status: %d", val_frontend)
		log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
		log.Printf("Redis Delete FRONTEND LISTRESULT status: %d", val_frontend_listresult)
		//AGEN
		val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company))
		val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + strings.ToLower(client_company))
		log.Printf("Redis Delete AGENT status: %d", val_agent)
		log.Printf("Redis Delete AGENT DASHBOARD status: %d", val_agent_dashboard)
		log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
		val_agent_redis := helpers.DeleteRedis(log_redis)
		log.Printf("Redis Delete LOG status: %d", val_agent_redis)
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
			//FRONTEND
			val_frontend := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company) + "_" + strings.ToLower(idpasarantogel))
			val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
			val_frontend_listresult := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company))
			log.Printf("Redis Delete FRONTEND status: %d", val_frontend)
			log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
			log.Printf("Redis Delete FRONTEND LISTRESULT status: %d", val_frontend_listresult)
			//AGEN
			val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company))
			val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + strings.ToLower(client_company))
			log.Printf("Redis Delete AGENT status: %d", val_agent)
			log.Printf("Redis Delete AGENT DASHBOARD status: %d", val_agent_dashboard)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
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

		//FRONTEND
		val_frontend := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company) + "_" + strings.ToLower(idpasarantogel))
		val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
		val_frontend_listresult := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company))
		log.Printf("Redis Delete FRONTEND status: %d", val_frontend)
		log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
		log.Printf("Redis Delete FRONTEND LISTRESULT status: %d", val_frontend_listresult)
		//AGEN
		val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company))
		val_agent2 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice))
		val_agent3 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTMEMBER")
		val_agent4 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBETTABLE")
		val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + strings.ToLower(client_company))
		log.Printf("Redis Delete AGENT PERIODE : %d", val_agent)
		log.Printf("Redis Delete AGENT PERIODE DETAIL: %d", val_agent2)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LISTMEMBER: %d", val_agent3)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LISTBETTABLE: %d", val_agent4)
		log.Printf("Redis Delete AGENT DASHBOARD status: %d", val_agent_dashboard)
		val_agent4d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_4D")
		val_agent3d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_3D")
		val_agent2d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_2D")
		val_agent2dd := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_2DD")
		val_agent2dt := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_2DT")
		val_agentcolokbebas := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_BEBAS")
		val_agentcolokmacau := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_MACAU")
		val_agentcoloknaga := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_NAGA")
		val_agentcolokjitu := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_JITU")
		val_agent5050umum := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_50_50_UMUM")
		val_agent5050special := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_50_50_SPECIAL")
		val_agent5050kombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_50_50_KOMBINASI")
		val_agentmacaukombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_MACAU_KOMBINASI")
		val_agentdasar := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_DASAR")
		val_agentshio := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_SHIO")
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
		val_agentall := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBET_all")
		val_agentwinner := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBET_winner")
		val_agentcancel := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBET_cancel")
		log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS ALL: %d", val_agentall)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS WINNER: %d", val_agentwinner)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS CANCEL: %d", val_agentcancel)
		log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
		val_agent_redis := helpers.DeleteRedis(log_redis)
		log.Printf("Redis Delete LOG status: %d", val_agent_redis)
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
			//FRONTEND
			val_frontend := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company) + "_" + strings.ToLower(idpasarantogel))
			val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
			val_frontend_listresult := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company))
			log.Printf("Redis Delete FRONTEND status: %d", val_frontend)
			log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
			log.Printf("Redis Delete FRONTEND LISTRESULT status: %d", val_frontend_listresult)
			//AGEN
			val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company))
			val_agent2 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice))
			val_agent3 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTMEMBER")
			val_agent4 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBETTABLE")
			val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + strings.ToLower(client_company))
			log.Printf("Redis Delete AGENT PERIODE : %d", val_agent)
			log.Printf("Redis Delete AGENT PERIODE DETAIL: %d", val_agent2)
			log.Printf("Redis Delete AGENT PERIODE DETAIL LISTMEMBER: %d", val_agent3)
			log.Printf("Redis Delete AGENT PERIODE DETAIL LISTBETTABLE: %d", val_agent4)
			log.Printf("Redis Delete AGENT DASHBOARD status: %d", val_agent_dashboard)
			val_agent4d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_4D")
			val_agent3d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_3D")
			val_agent2d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_2D")
			val_agent2dd := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_2DD")
			val_agent2dt := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_2DT")
			val_agentcolokbebas := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_BEBAS")
			val_agentcolokmacau := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_MACAU")
			val_agentcoloknaga := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_NAGA")
			val_agentcolokjitu := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_JITU")
			val_agent5050umum := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_50_50_UMUM")
			val_agent5050special := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_50_50_SPECIAL")
			val_agent5050kombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_50_50_KOMBINASI")
			val_agentmacaukombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_MACAU_KOMBINASI")
			val_agentdasar := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_DASAR")
			val_agentshio := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_SHIO")
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
			val_agentall := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBET_all")
			val_agentwinner := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBET_winner")
			val_agentcancel := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBET_cancel")
			log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS ALL: %d", val_agentall)
			log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS WINNER: %d", val_agentwinner)
			log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS CANCEL: %d", val_agentcancel)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
func PeriodeCancelBet(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_periodecancelbet)
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
		//FRONTEND
		val_frontend := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company) + "_" + strings.ToLower(idpasarantogel))
		val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
		val_frontend_listresult := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company))
		log.Printf("Redis Delete FRONTEND status: %d", val_frontend)
		log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
		log.Printf("Redis Delete FRONTEND LISTRESULT status: %d", val_frontend_listresult)
		//AGEN
		val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company))
		val_agent2 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice))
		val_agent3 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTMEMBER")
		val_agent4 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBETTABLE")
		val_agent5 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_" + client.Permainan)
		val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + strings.ToLower(client_company))
		log.Printf("Redis Delete AGENT PERIODE : %d", val_agent)
		log.Printf("Redis Delete AGENT PERIODE DETAIL: %d", val_agent2)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LISTMEMBER: %d", val_agent3)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LISTBETTABLE: %d", val_agent4)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LISTPERMAINAN: %d", val_agent5)
		log.Printf("Redis Delete AGENT DASHBOARD status: %d", val_agent_dashboard)
		val_agent4d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_4D")
		val_agent3d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_3D")
		val_agent2d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_2D")
		val_agent2dd := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_2DD")
		val_agent2dt := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_2DT")
		val_agentcolokbebas := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_BEBAS")
		val_agentcolokmacau := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_MACAU")
		val_agentcoloknaga := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_NAGA")
		val_agentcolokjitu := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_JITU")
		val_agent5050umum := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_50_50_UMUM")
		val_agent5050special := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_50_50_SPECIAL")
		val_agent5050kombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_50_50_KOMBINASI")
		val_agentmacaukombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_MACAU_KOMBINASI")
		val_agentdasar := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_DASAR")
		val_agentshio := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_SHIO")
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
		val_agentall := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBET_all")
		val_agentwinner := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBET_winner")
		val_agentcancel := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBET_cancel")
		log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS ALL: %d", val_agentall)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS WINNER: %d", val_agentwinner)
		log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS CANCEL: %d", val_agentcancel)
		log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
		val_agent_redis := helpers.DeleteRedis(log_redis)
		log.Printf("Redis Delete LOG status: %d", val_agent_redis)
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
			//FRONTEND
			val_frontend := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company) + "_" + strings.ToLower(idpasarantogel))
			val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
			val_frontend_listresult := helpers.DeleteRedis("listresult_" + strings.ToLower(client_company))
			log.Printf("Redis Delete FRONTEND status: %d", val_frontend)
			log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
			log.Printf("Redis Delete FRONTEND LISTRESULT status: %d", val_frontend_listresult)
			//AGEN
			val_agent := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company))
			val_agent2 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice))
			val_agent3 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTMEMBER")
			val_agent4 := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBETTABLE")
			val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + strings.ToLower(client_company))
			log.Printf("Redis Delete AGENT PERIODE : %d", val_agent)
			log.Printf("Redis Delete AGENT PERIODE DETAIL: %d", val_agent2)
			log.Printf("Redis Delete AGENT PERIODE DETAIL LISTMEMBER: %d", val_agent3)
			log.Printf("Redis Delete AGENT PERIODE DETAIL LISTBETTABLE: %d", val_agent4)
			log.Printf("Redis Delete AGENT DASHBOARD status: %d", val_agent_dashboard)
			val_agent4d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_4D")
			val_agent3d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_3D")
			val_agent2d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_2D")
			val_agent2dd := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_2DD")
			val_agent2dt := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_2DT")
			val_agentcolokbebas := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_BEBAS")
			val_agentcolokmacau := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_MACAU")
			val_agentcoloknaga := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_NAGA")
			val_agentcolokjitu := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_COLOK_JITU")
			val_agent5050umum := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_50_50_UMUM")
			val_agent5050special := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_50_50_SPECIAL")
			val_agent5050kombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_50_50_KOMBINASI")
			val_agentmacaukombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_MACAU_KOMBINASI")
			val_agentdasar := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_DASAR")
			val_agentshio := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTPERMAINAN_SHIO")
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
			val_agentall := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBET_all")
			val_agentwinner := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBET_winner")
			val_agentcancel := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client_company) + "_INVOICE_" + strconv.Itoa(client.Idinvoice) + "_LISTBET_cancel")
			log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS ALL: %d", val_agentall)
			log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS WINNER: %d", val_agentwinner)
			log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS CANCEL: %d", val_agentcancel)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
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
