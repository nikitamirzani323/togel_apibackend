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

type pasarandetail struct {
	Idpasaran int `json:"idpasaran"`
}
type pasaransave struct {
	Idpasaran         int    `json:"idpasaran"`
	Page              string `json:"page"`
	Pasaran_diundi    string `json:"pasaran_diundi"`
	Pasaran_url       string `json:"pasaran_url"`
	Pasaran_jamtutup  string `json:"pasaran_jamtutup"`
	Pasaran_jamjadwal string `json:"pasaran_jamjadwal"`
	Pasaran_jamopen   string `json:"pasaran_jamopen"`
	Pasaran_display   int    `json:"pasaran_display"`
	Pasaran_status    string `json:"pasaran_status"`
}
type Pasaransaveonline struct {
	Idpasaran   int    `json:"idpasaran" validate:"required"`
	Page        string `json:"page"`
	Haripasaran string `json:"pasaran_hariraya" validate:"required,alpha"`
}
type Pasarandeleteonline struct {
	Idpasaran      int    `json:"idpasaran" validate:"required"`
	Idpasaraonline int    `json:"idpasaraonline" validate:"required,numeric"`
	Page           string `json:"page"`
}
type pasaransavelimit struct {
	Idpasaran            int    `json:"idpasaran"`
	Page                 string `json:"page"`
	Pasaran_limitline4d  int    `json:"pasaran_limitline4d"`
	Pasaran_limitline3d  int    `json:"pasaran_limitline3d"`
	Pasaran_limitline2d  int    `json:"pasaran_limitline2d"`
	Pasaran_limitline2dd int    `json:"pasaran_limitline2dd"`
	Pasaran_limitline2dt int    `json:"pasaran_limitline2dt"`
}
type pasaranconf432 struct {
	Idpasaran                   int     `json:"idpasaran"`
	Idpasarantogel              string  `json:"idpasarantogel"`
	Page                        string  `json:"page"`
	Pasaran_minbet_432d         int     `json:"pasaran_minbet_432d"`
	Pasaran_maxbet4d_432d       int     `json:"pasaran_maxbet4d_432d"`
	Pasaran_maxbet3d_432d       int     `json:"pasaran_maxbet3d_432d"`
	Pasaran_maxbet2d_432d       int     `json:"pasaran_maxbet2d_432d"`
	Pasaran_maxbet2dd_432d      int     `json:"pasaran_maxbet2dd_432d"`
	Pasaran_maxbet2dt_432d      int     `json:"pasaran_maxbet2dt_432d"`
	Pasaran_limitotal4d_432d    int     `json:"pasaran_limitotal4d_432d"`
	Pasaran_limitotal3d_432d    int     `json:"pasaran_limitotal3d_432d"`
	Pasaran_limitotal2d_432d    int     `json:"pasaran_limitotal2d_432d"`
	Pasaran_limitotal2dd_432d   int     `json:"pasaran_limitotal2dd_432d"`
	Pasaran_limitotal2dt_432d   int     `json:"pasaran_limitotal2dt_432d"`
	Pasaran_limitglobal4d_432d  int     `json:"pasaran_limitglobal4d_432d"`
	Pasaran_limitglobal3d_432d  int     `json:"pasaran_limitglobal3d_432d"`
	Pasaran_limitglobal2d_432d  int     `json:"pasaran_limitglobal2d_432d"`
	Pasaran_limitglobal2dd_432d int     `json:"pasaran_limitglobal2dd_432d"`
	Pasaran_limitglobal2dt_432d int     `json:"pasaran_limitglobal2dt_432d"`
	Pasaran_win4d_432d          int     `json:"pasaran_win4d_432d"`
	Pasaran_win3d_432d          int     `json:"pasaran_win3d_432d"`
	Pasaran_win2d_432d          int     `json:"pasaran_win2d_432d"`
	Pasaran_win2dd_432d         int     `json:"pasaran_win2dd_432d"`
	Pasaran_win2dt_432d         int     `json:"pasaran_win2dt_432d"`
	Pasaran_disc4d_432d         float32 `json:"pasaran_disc4d_432d"`
	Pasaran_disc3d_432d         float32 `json:"pasaran_disc3d_432d"`
	Pasaran_disc2d_432d         float32 `json:"pasaran_disc2d_432d"`
	Pasaran_disc2dd_432d        float32 `json:"pasaran_disc2dd_432d"`
	Pasaran_disc2dt_432d        float32 `json:"pasaran_disc2dt_432d"`
}
type pasaranconfcbebas struct {
	Idpasaran                  int     `json:"idpasaran"`
	Idpasarantogel             string  `json:"idpasarantogel"`
	Page                       string  `json:"page"`
	Pasaran_minbet_cbebas      int     `json:"pasaran_minbet_cbebas"`
	Pasaran_maxbet_cbebas      int     `json:"pasaran_maxbet_cbebas"`
	Pasaran_limitotal_cbebas   int     `json:"pasaran_limitotal_cbebas"`
	Pasaran_limitglobal_cbebas int     `json:"pasaran_limitglobal_cbebas"`
	Pasaran_win_cbebas         float32 `json:"pasaran_win_cbebas"`
	Pasaran_disc_cbebas        float32 `json:"pasaran_disc_cbebas"`
}
type pasaranconfcmacau struct {
	Idpasaran                  int     `json:"idpasaran"`
	Idpasarantogel             string  `json:"idpasarantogel"`
	Page                       string  `json:"page"`
	Pasaran_minbet_cmacau      int     `json:"pasaran_minbet_cmacau"`
	Pasaran_maxbet_cmacau      int     `json:"pasaran_maxbet_cmacau"`
	Pasaran_limitotal_cmacau   int     `json:"pasaran_limitotal_cmacau"`
	Pasaran_limitglobal_cmacau int     `json:"pasaran_limitglobal_cmacau"`
	Pasaran_win2_cmacau        float32 `json:"pasaran_win2_cmacau"`
	Pasaran_win3_cmacau        float32 `json:"pasaran_win3_cmacau"`
	Pasaran_win4_cmacau        float32 `json:"pasaran_win4_cmacau"`
	Pasaran_disc_cmacau        float32 `json:"pasaran_disc_cmacau"`
}
type pasaranconfcnaga struct {
	Idpasaran                 int     `json:"idpasaran"`
	Idpasarantogel            string  `json:"idpasarantogel"`
	Page                      string  `json:"page"`
	Pasaran_minbet_cnaga      int     `json:"pasaran_minbet_cnaga"`
	Pasaran_maxbet_cnaga      int     `json:"pasaran_maxbet_cnaga"`
	Pasaran_limittotal_cnaga  int     `json:"pasaran_limittotal_cnaga"`
	Pasaran_limitglobal_cnaga int     `json:"pasaran_limitglobal_cnaga"`
	Pasaran_win3_cnaga        float32 `json:"pasaran_win3_cnaga"`
	Pasaran_win4_cnaga        float32 `json:"pasaran_win4_cnaga"`
	Pasaran_disc_cnaga        float32 `json:"pasaran_disc_cnaga"`
}
type pasaranconfcjitu struct {
	Idpasaran                 int     `json:"idpasaran"`
	Idpasarantogel            string  `json:"idpasarantogel"`
	Page                      string  `json:"page"`
	Pasaran_minbet_cjitu      int     `json:"pasaran_minbet_cjitu"`
	Pasaran_maxbet_cjitu      int     `json:"pasaran_maxbet_cjitu"`
	Pasaran_limittotal_cjitu  int     `json:"pasaran_limittotal_cjitu"`
	Pasaran_limitglobal_cjitu int     `json:"pasaran_limitglobal_cjitu"`
	Pasaran_winas_cjitu       float32 `json:"pasaran_winas_cjitu"`
	Pasaran_winkop_cjitu      float32 `json:"pasaran_winkop_cjitu"`
	Pasaran_winkepala_cjitu   float32 `json:"pasaran_winkepala_cjitu"`
	Pasaran_winekor_cjitu     float32 `json:"pasaran_winekor_cjitu"`
	Pasaran_desc_cjitu        float32 `json:"pasaran_desc_cjitu"`
}
type pasaranconfc5050 struct {
	Idpasaran                    int     `json:"idpasaran"`
	Idpasarantogel               string  `json:"idpasarantogel"`
	Page                         string  `json:"page"`
	Pasaran_minbet_5050umum      int     `json:"pasaran_minbet_5050umum"`
	Pasaran_maxbet_5050umum      int     `json:"pasaran_maxbet_5050umum"`
	Pasaran_limittotal_5050umum  int     `json:"pasaran_limittotal_5050umum"`
	Pasaran_limitglobal_5050umum int     `json:"pasaran_limitglobal_5050umum"`
	Pasaran_keibesar_5050umum    float32 `json:"pasaran_keibesar_5050umum"`
	Pasaran_keikecil_5050umum    float32 `json:"pasaran_keikecil_5050umum"`
	Pasaran_keigenap_5050umum    float32 `json:"pasaran_keigenap_5050umum"`
	Pasaran_keiganjil_5050umum   float32 `json:"pasaran_keiganjil_5050umum"`
	Pasaran_keitengah_5050umum   float32 `json:"pasaran_keitengah_5050umum"`
	Pasaran_keitepi_5050umum     float32 `json:"pasaran_keitepi_5050umum"`
	Pasaran_discbesar_5050umum   float32 `json:"pasaran_discbesar_5050umum"`
	Pasaran_disckecil_5050umum   float32 `json:"pasaran_disckecil_5050umum"`
	Pasaran_discgenap_5050umum   float32 `json:"pasaran_discgenap_5050umum"`
	Pasaran_discganjil_5050umum  float32 `json:"pasaran_discganjil_5050umum"`
	Pasaran_disctengah_5050umum  float32 `json:"pasaran_disctengah_5050umum"`
	Pasaran_disctepi_5050umum    float32 `json:"pasaran_disctepi_5050umum"`
}
type pasaranconfc5050special struct {
	Idpasaran                            int     `json:"idpasaran"`
	Idpasarantogel                       string  `json:"idpasarantogel"`
	Page                                 string  `json:"page"`
	Pasaran_minbet_5050special           int     `json:"pasaran_minbet_5050special"`
	Pasaran_maxbet_5050special           int     `json:"pasaran_maxbet_5050special"`
	Pasaran_limittotal_5050special       int     `json:"pasaran_limittotal_5050special"`
	Pasaran_limitglobal_5050special      int     `json:"pasaran_limitglobal_5050special"`
	Pasaran_keiasganjil_5050special      float32 `json:"pasaran_keiasganjil_5050special"`
	Pasaran_keiasgenap_5050special       float32 `json:"pasaran_keiasgenap_5050special"`
	Pasaran_keiasbesar_5050special       float32 `json:"pasaran_keiasbesar_5050special"`
	Pasaran_keiaskecil_5050special       float32 `json:"pasaran_keiaskecil_5050special"`
	Pasaran_keikopganjil_5050special     float32 `json:"pasaran_keikopganjil_5050special"`
	Pasaran_keikopgenap_5050special      float32 `json:"pasaran_keikopgenap_5050special"`
	Pasaran_keikopbesar_5050special      float32 `json:"pasaran_keikopbesar_5050special"`
	Pasaran_keikopkecil_5050special      float32 `json:"pasaran_keikopkecil_5050special"`
	Pasaran_keikepalaganjil_5050special  float32 `json:"pasaran_keikepalaganjil_5050special"`
	Pasaran_keikepalagenap_5050special   float32 `json:"pasaran_keikepalagenap_5050special"`
	Pasaran_keikepalabesar_5050special   float32 `json:"pasaran_keikepalabesar_5050special"`
	Pasaran_keikepalakecil_5050special   float32 `json:"pasaran_keikepalakecil_5050special"`
	Pasaran_keiekorganjil_5050special    float32 `json:"pasaran_keiekorganjil_5050special"`
	Pasaran_keiekorgenap_5050special     float32 `json:"pasaran_keiekorgenap_5050special"`
	Pasaran_keiekorbesar_5050special     float32 `json:"pasaran_keiekorbesar_5050special"`
	Pasaran_keiekorkecil_5050special     float32 `json:"pasaran_keiekorkecil_5050special"`
	Pasaran_discasganjil_5050special     float32 `json:"pasaran_discasganjil_5050special"`
	Pasaran_discasgenap_5050special      float32 `json:"pasaran_discasgenap_5050special"`
	Pasaran_discasbesar_5050special      float32 `json:"pasaran_discasbesar_5050special"`
	Pasaran_discaskecil_5050special      float32 `json:"pasaran_discaskecil_5050special"`
	Pasaran_disckopganjil_5050special    float32 `json:"pasaran_disckopganjil_5050special"`
	Pasaran_disckopgenap_5050special     float32 `json:"pasaran_disckopgenap_5050special"`
	Pasaran_disckopbesar_5050special     float32 `json:"pasaran_disckopbesar_5050special"`
	Pasaran_disckopkecil_5050special     float32 `json:"pasaran_disckopkecil_5050special"`
	Pasaran_disckepalaganjil_5050special float32 `json:"pasaran_disckepalaganjil_5050special"`
	Pasaran_disckepalagenap_5050special  float32 `json:"pasaran_disckepalagenap_5050special"`
	Pasaran_disckepalabesar_5050special  float32 `json:"pasaran_disckepalabesar_5050special"`
	Pasaran_disckepalakecil_5050special  float32 `json:"pasaran_disckepalakecil_5050special"`
	Pasaran_discekorganjil_5050special   float32 `json:"pasaran_discekorganjil_5050special"`
	Pasaran_discekorgenap_5050special    float32 `json:"pasaran_discekorgenap_5050special"`
	Pasaran_discekorbesar_5050special    float32 `json:"pasaran_discekorbesar_5050special"`
	Pasaran_discekorkecil_5050special    float32 `json:"pasaran_discekorkecil_5050special"`
}
type pasaranconfc5050kombinasi struct {
	Idpasaran                                 int     `json:"idpasaran"`
	Idpasarantogel                            string  `json:"idpasarantogel"`
	Page                                      string  `json:"page"`
	Pasaran_minbet_5050kombinasi              int     `json:"pasaran_minbet_5050kombinasi"`
	Pasaran_maxbet_5050kombinasi              int     `json:"pasaran_maxbet_5050kombinasi"`
	Pasaran_limittotal_5050kombinasi          int     `json:"pasaran_limittotal_5050kombinasi"`
	Pasaran_limitglobal_5050kombinasi         int     `json:"pasaran_limitglobal_5050kombinasi"`
	Pasaran_belakangkeimono_5050kombinasi     float32 `json:"pasaran_belakangkeimono_5050kombinasi"`
	Pasaran_belakangkeistereo_5050kombinasi   float32 `json:"pasaran_belakangkeistereo_5050kombinasi"`
	Pasaran_belakangkeikembang_5050kombinasi  float32 `json:"pasaran_belakangkeikembang_5050kombinasi"`
	Pasaran_belakangkeikempis_5050kombinasi   float32 `json:"pasaran_belakangkeikempis_5050kombinasi"`
	Pasaran_belakangkeikembar_5050kombinasi   float32 `json:"pasaran_belakangkeikembar_5050kombinasi"`
	Pasaran_tengahkeimono_5050kombinasi       float32 `json:"pasaran_tengahkeimono_5050kombinasi"`
	Pasaran_tengahkeistereo_5050kombinasi     float32 `json:"pasaran_tengahkeistereo_5050kombinasi"`
	Pasaran_tengahkeikembang_5050kombinasi    float32 `json:"pasaran_tengahkeikembang_5050kombinasi"`
	Pasaran_tengahkeikempis_5050kombinasi     float32 `json:"pasaran_tengahkeikempis_5050kombinasi"`
	Pasaran_tengahkeikembar_5050kombinasi     float32 `json:"pasaran_tengahkeikembar_5050kombinasi"`
	Pasaran_depankeimono_5050kombinasi        float32 `json:"pasaran_depankeimono_5050kombinasi"`
	Pasaran_depankeistereo_5050kombinasi      float32 `json:"pasaran_depankeistereo_5050kombinasi"`
	Pasaran_depankeikembang_5050kombinasi     float32 `json:"pasaran_depankeikembang_5050kombinasi"`
	Pasaran_depankeikempis_5050kombinasi      float32 `json:"pasaran_depankeikempis_5050kombinasi"`
	Pasaran_depankeikembar_5050kombinasi      float32 `json:"pasaran_depankeikembar_5050kombinasi"`
	Pasaran_belakangdiscmono_5050kombinasi    float32 `json:"pasaran_belakangdiscmono_5050kombinasi"`
	Pasaran_belakangdiscstereo_5050kombinasi  float32 `json:"pasaran_belakangdiscstereo_5050kombinasi"`
	Pasaran_belakangdisckembang_5050kombinasi float32 `json:"pasaran_belakangdisckembang_5050kombinasi"`
	Pasaran_belakangdisckempis_5050kombinasi  float32 `json:"pasaran_belakangdisckempis_5050kombinasi"`
	Pasaran_belakangdisckembar_5050kombinasi  float32 `json:"pasaran_belakangdisckembar_5050kombinasi"`
	Pasaran_tengahdiscmono_5050kombinasi      float32 `json:"pasaran_tengahdiscmono_5050kombinasi"`
	Pasaran_tengahdiscstereo_5050kombinasi    float32 `json:"pasaran_tengahdiscstereo_5050kombinasi"`
	Pasaran_tengahdisckembang_5050kombinasi   float32 `json:"pasaran_tengahdisckembang_5050kombinasi"`
	Pasaran_tengahdisckempis_5050kombinasi    float32 `json:"pasaran_tengahdisckempis_5050kombinasi"`
	Pasaran_tengahdisckembar_5050kombinasi    float32 `json:"pasaran_tengahdisckembar_5050kombinasi"`
	Pasaran_depandiscmono_5050kombinasi       float32 `json:"pasaran_depandiscmono_5050kombinasi"`
	Pasaran_depandiscstereo_5050kombinasi     float32 `json:"pasaran_depandiscstereo_5050kombinasi"`
	Pasaran_depandisckembang_5050kombinasi    float32 `json:"pasaran_depandisckembang_5050kombinasi"`
	Pasaran_depandisckempis_5050kombinasi     float32 `json:"pasaran_depandisckempis_5050kombinasi"`
	Pasaran_depandisckembar_5050kombinasi     float32 `json:"pasaran_depandisckembar_5050kombinasi"`
}
type pasaranconfmakaukombinasi struct {
	Idpasaran                     int     `json:"idpasaran"`
	Idpasarantogel                string  `json:"idpasarantogel"`
	Page                          string  `json:"page"`
	Pasaran_minbet_kombinasi      int     `json:"pasaran_minbet_kombinasi"`
	Pasaran_maxbet_kombinasi      int     `json:"pasaran_maxbet_kombinasi"`
	Pasaran_limittotal_kombinasi  int     `json:"pasaran_limittotal_kombinasi"`
	Pasaran_limitglobal_kombinasi int     `json:"pasaran_limitglobal_kombinasi"`
	Pasaran_win_kombinasi         float32 `json:"pasaran_win_kombinasi"`
	Pasaran_disc_kombinasi        float32 `json:"pasaran_disc_kombinasi"`
}
type pasaranconfdasar struct {
	Idpasaran                 int     `json:"idpasaran"`
	Idpasarantogel            string  `json:"idpasarantogel"`
	Page                      string  `json:"page"`
	Pasaran_minbet_dasar      int     `json:"pasaran_minbet_dasar"`
	Pasaran_maxbet_dasar      int     `json:"pasaran_maxbet_dasar"`
	Pasaran_limittotal_dasar  int     `json:"pasaran_limittotal_dasar"`
	Pasaran_limitglobal_dasar int     `json:"pasaran_limitglobal_dasar"`
	Pasaran_keibesar_dasar    float32 `json:"pasaran_keibesar_dasar"`
	Pasaran_keikecil_dasar    float32 `json:"pasaran_keikecil_dasar"`
	Pasaran_keigenap_dasar    float32 `json:"pasaran_keigenap_dasar"`
	Pasaran_keiganjil_dasar   float32 `json:"pasaran_keiganjil_dasar"`
	Pasaran_discbesar_dasar   float32 `json:"pasaran_discbesar_dasar"`
	Pasaran_disckecil_dasar   float32 `json:"pasaran_disckecil_dasar"`
	Pasaran_discgenap_dasar   float32 `json:"pasaran_discgenap_dasar"`
	Pasaran_discganjil_dasar  float32 `json:"pasaran_discganjil_dasar"`
}
type pasaranconfshio struct {
	Idpasaran                int     `json:"idpasaran"`
	Idpasarantogel           string  `json:"idpasarantogel"`
	Page                     string  `json:"page"`
	Pasaran_minbet_shio      int     `json:"pasaran_minbet_shio"`
	Pasaran_maxbet_shio      int     `json:"pasaran_maxbet_shio"`
	Pasaran_limittotal_shio  int     `json:"pasaran_limittotal_shio"`
	Pasaran_limitglobal_shio int     `json:"pasaran_limitglobal_shio"`
	Pasaran_shioyear_shio    string  `json:"pasaran_shioyear_shio"`
	Pasaran_disc_shio        float32 `json:"pasaran_disc_shio"`
	Pasaran_win_shio         float32 `json:"pasaran_win_shio"`
}
type responseredis_pasaranhome struct {
	Idcomppasaran           int    `json:"idcomppasaran"`
	Nmpasarantogel          string `json:"nmpasarantogel"`
	Pasarandiundi           string `json:"pasarandiundi"`
	Jamtutup                string `json:"jamtutup"`
	Jamjadwal               string `json:"jamjadwal"`
	Jamopen                 string `json:"jamopen"`
	Displaypasaran          int    `json:"displaypasaran"`
	Statuspasaran           string `json:"statuspasaran"`
	Statuspasaranactive     string `json:"statuspasaranactive"`
	Statuspasaran_css       string `json:"statuspasaran_css"`
	Statuspasaranactive_css string `json:"statuspasaranactive_css"`
}

func PasaranHome(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	field_redis := "LISTPASARAN_AGENT_" + client_company
	var obj responseredis_pasaranhome
	var arraobj []responseredis_pasaranhome
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		idcomppasaran, _ := jsonparser.GetInt(value, "idcomppasaran")
		nmpasarantogel, _ := jsonparser.GetString(value, "nmpasarantogel")
		pasarandiundi, _ := jsonparser.GetString(value, "pasarandiundi")
		jamtutup, _ := jsonparser.GetString(value, "jamtutup")
		jamjadwal, _ := jsonparser.GetString(value, "jamjadwal")
		jamopen, _ := jsonparser.GetString(value, "jamopen")
		displaypasaran, _ := jsonparser.GetInt(value, "displaypasaran")
		statuspasaran, _ := jsonparser.GetString(value, "statuspasaran")
		statuspasaranactive, _ := jsonparser.GetString(value, "statuspasaranactive")
		statuspasaran_css, _ := jsonparser.GetString(value, "statuspasaran_css")
		statuspasaranactive_css, _ := jsonparser.GetString(value, "statuspasaranactive_css")

		obj.Idcomppasaran = int(idcomppasaran)
		obj.Nmpasarantogel = nmpasarantogel
		obj.Pasarandiundi = pasarandiundi
		obj.Jamtutup = jamtutup
		obj.Jamjadwal = jamjadwal
		obj.Jamopen = jamopen
		obj.Displaypasaran = int(displaypasaran)
		obj.Statuspasaran = statuspasaran
		obj.Statuspasaranactive = statuspasaranactive
		obj.Statuspasaran_css = statuspasaran_css
		obj.Statuspasaranactive_css = statuspasaranactive_css
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_home(client_company)
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
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PasaranDetail(c *fiber.Ctx) error {
	client := new(pasarandetail)
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
	result, err := models.Fetch_detail(client_company, client.Idpasaran)
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
func PasaranSave(c *fiber.Ctx) error {
	client := new(pasaransave)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")
	field_redis := "LISTPASARAN_AGENT_" + client_company
	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Save_Pasaran(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_diundi,
			client.Pasaran_url,
			client.Pasaran_jamtutup,
			client.Pasaran_jamjadwal,
			client.Pasaran_jamopen,
			client.Pasaran_status,
			client.Pasaran_display)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val_agent := helpers.DeleteRedis(field_redis)
		log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
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
			result, err := models.Save_Pasaran(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_diundi,
				client.Pasaran_url,
				client.Pasaran_jamtutup,
				client.Pasaran_jamjadwal,
				client.Pasaran_jamopen,
				client.Pasaran_status,
				client.Pasaran_display)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val_agent := helpers.DeleteRedis(field_redis)
			log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
			return c.JSON(result)
		}
	}
}

func PasaranSaveOnline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(Pasaransaveonline)
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
	field_redis := "LISTPASARAN_AGENT_" + client_company
	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranOnline(
			client_username,
			client_company,
			client.Idpasaran,
			client.Haripasaran)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val_agent := helpers.DeleteRedis(field_redis)
		log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
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
			result, err := models.Save_PasaranOnline(
				client_username,
				client_company,
				client.Idpasaran,
				client.Haripasaran)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val_agent := helpers.DeleteRedis(field_redis)
			log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
			return c.JSON(result)
		}
	}
}
func PasaranDeleteOnline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(Pasarandeleteonline)
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
	field_redis := "LISTPASARAN_AGENT_" + client_company
	if typeadmin == "MASTER" {
		result, err := models.Delete_PasaranOnline(
			client_company,
			client.Idpasaran,
			client.Idpasaraonline)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val_agent := helpers.DeleteRedis(field_redis)
		log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
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
			result, err := models.Delete_PasaranOnline(
				client_company,
				client.Idpasaran,
				client.Idpasaraonline)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val_agent := helpers.DeleteRedis(field_redis)
			log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
			return c.JSON(result)
		}
	}
}
func PasaranSaveLimit(c *fiber.Ctx) error {
	client := new(pasaransavelimit)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)
	field_redis := "LISTPASARAN_AGENT_" + client_company
	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranLimitline(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_limitline4d,
			client.Pasaran_limitline3d,
			client.Pasaran_limitline2d,
			client.Pasaran_limitline2dd,
			client.Pasaran_limitline2dt)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val_agent := helpers.DeleteRedis(field_redis)
		log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
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
			result, err := models.Save_PasaranLimitline(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_limitline4d,
				client.Pasaran_limitline3d,
				client.Pasaran_limitline2d,
				client.Pasaran_limitline2dd,
				client.Pasaran_limitline2dt)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val_agent := helpers.DeleteRedis(field_redis)
			log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConf432d(c *fiber.Ctx) error {
	client := new(pasaranconf432)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConf432(
			client_username, client_company, client.Idpasaran,
			client.Pasaran_minbet_432d,
			client.Pasaran_maxbet4d_432d, client.Pasaran_maxbet3d_432d, client.Pasaran_maxbet2d_432d, client.Pasaran_maxbet2dd_432d, client.Pasaran_maxbet2dt_432d,
			client.Pasaran_win4d_432d, client.Pasaran_win3d_432d, client.Pasaran_win2d_432d, client.Pasaran_win2dd_432d, client.Pasaran_win2dt_432d,
			client.Pasaran_disc4d_432d, client.Pasaran_disc3d_432d, client.Pasaran_disc2d_432d, client.Pasaran_disc2dd_432d, client.Pasaran_disc2dt_432d,
			client.Pasaran_limitglobal4d_432d, client.Pasaran_limitglobal3d_432d, client.Pasaran_limitglobal2d_432d, client.Pasaran_limitglobal2dd_432d, client.Pasaran_limitglobal2dt_432d,
			client.Pasaran_limitotal4d_432d, client.Pasaran_limitotal3d_432d, client.Pasaran_limitotal2d_432d, client.Pasaran_limitotal2dd_432d, client.Pasaran_limitotal2dt_432d)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_4-3-2")
		log.Printf("Redis Delete Client - CONF 432 status: %d", val)
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
			result, err := models.Save_PasaranConf432(
				client_username, client_company, client.Idpasaran,
				client.Pasaran_minbet_432d,
				client.Pasaran_maxbet4d_432d, client.Pasaran_maxbet3d_432d, client.Pasaran_maxbet2d_432d, client.Pasaran_maxbet2dd_432d, client.Pasaran_maxbet2dt_432d,
				client.Pasaran_win4d_432d, client.Pasaran_win3d_432d, client.Pasaran_win2d_432d, client.Pasaran_win2dd_432d, client.Pasaran_win2dt_432d,
				client.Pasaran_disc4d_432d, client.Pasaran_disc3d_432d, client.Pasaran_disc2d_432d, client.Pasaran_disc2dd_432d, client.Pasaran_disc2dt_432d,
				client.Pasaran_limitglobal4d_432d, client.Pasaran_limitglobal3d_432d, client.Pasaran_limitglobal2d_432d, client.Pasaran_limitglobal2dd_432d, client.Pasaran_limitglobal2dt_432d,
				client.Pasaran_limitotal4d_432d, client.Pasaran_limitotal3d_432d, client.Pasaran_limitotal2d_432d, client.Pasaran_limitotal2dd_432d, client.Pasaran_limitotal2dt_432d)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_4-3-2")
			log.Printf("Redis Delete Client - CONF 432 status: %d", val)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfColokBebas(c *fiber.Ctx) error {
	client := new(pasaranconfcbebas)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConfColokBebas(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_minbet_cbebas,
			client.Pasaran_maxbet_cbebas,
			client.Pasaran_win_cbebas,
			client.Pasaran_disc_cbebas,
			client.Pasaran_limitglobal_cbebas,
			client.Pasaran_limitotal_cbebas)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_colok")
		log.Printf("Redis Delete Client - CONF COLOK status: %d", val)
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
			result, err := models.Save_PasaranConfColokBebas(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_minbet_cbebas,
				client.Pasaran_maxbet_cbebas,
				client.Pasaran_win_cbebas,
				client.Pasaran_disc_cbebas,
				client.Pasaran_limitglobal_cbebas,
				client.Pasaran_limitotal_cbebas)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_colok")
			log.Printf("Redis Delete Client - CONF COLOK status: %d", val)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfColokMacau(c *fiber.Ctx) error {
	client := new(pasaranconfcmacau)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConfColokMacau(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_maxbet_cmacau,
			client.Pasaran_maxbet_cmacau,
			client.Pasaran_win2_cmacau,
			client.Pasaran_win3_cmacau,
			client.Pasaran_win4_cmacau,
			client.Pasaran_disc_cmacau,
			client.Pasaran_limitglobal_cmacau,
			client.Pasaran_limitotal_cmacau)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_colok")
		log.Printf("Redis Delete Client - CONF COLOK status: %d", val)
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
			result, err := models.Save_PasaranConfColokMacau(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_maxbet_cmacau,
				client.Pasaran_maxbet_cmacau,
				client.Pasaran_win2_cmacau,
				client.Pasaran_win3_cmacau,
				client.Pasaran_win4_cmacau,
				client.Pasaran_disc_cmacau,
				client.Pasaran_limitglobal_cmacau,
				client.Pasaran_limitotal_cmacau)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_colok")
			log.Printf("Redis Delete Client - CONF COLOK status: %d", val)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfColokNaga(c *fiber.Ctx) error {
	client := new(pasaranconfcnaga)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConfColokNaga(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_minbet_cnaga,
			client.Pasaran_maxbet_cnaga,
			client.Pasaran_win3_cnaga,
			client.Pasaran_win4_cnaga,
			client.Pasaran_disc_cnaga,
			client.Pasaran_limitglobal_cnaga,
			client.Pasaran_limittotal_cnaga)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_colok")
		log.Printf("Redis Delete Client - CONF COLOK status: %d", val)
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
			result, err := models.Save_PasaranConfColokNaga(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_minbet_cnaga,
				client.Pasaran_maxbet_cnaga,
				client.Pasaran_win3_cnaga,
				client.Pasaran_win4_cnaga,
				client.Pasaran_disc_cnaga,
				client.Pasaran_limitglobal_cnaga,
				client.Pasaran_limittotal_cnaga)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_colok")
			log.Printf("Redis Delete Client - CONF COLOK status: %d", val)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfColokJitu(c *fiber.Ctx) error {
	client := new(pasaranconfcjitu)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConfColokJitu(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_minbet_cjitu,
			client.Pasaran_maxbet_cjitu,
			client.Pasaran_winas_cjitu,
			client.Pasaran_winkop_cjitu,
			client.Pasaran_winkepala_cjitu,
			client.Pasaran_winekor_cjitu,
			client.Pasaran_desc_cjitu,
			client.Pasaran_limitglobal_cjitu,
			client.Pasaran_limittotal_cjitu)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_colok")
		log.Printf("Redis Delete Client - CONF COLOK status: %d", val)
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
			result, err := models.Save_PasaranConfColokJitu(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_minbet_cjitu,
				client.Pasaran_maxbet_cjitu,
				client.Pasaran_winas_cjitu,
				client.Pasaran_winkop_cjitu,
				client.Pasaran_winkepala_cjitu,
				client.Pasaran_winekor_cjitu,
				client.Pasaran_desc_cjitu,
				client.Pasaran_limitglobal_cjitu,
				client.Pasaran_limittotal_cjitu)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_colok")
			log.Printf("Redis Delete Client - CONF COLOK status: %d", val)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConf5050Umum(c *fiber.Ctx) error {
	client := new(pasaranconfc5050)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConf5050Umum(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_minbet_5050umum,
			client.Pasaran_maxbet_5050umum,
			client.Pasaran_keibesar_5050umum,
			client.Pasaran_keikecil_5050umum,
			client.Pasaran_keigenap_5050umum,
			client.Pasaran_keiganjil_5050umum,
			client.Pasaran_keitengah_5050umum,
			client.Pasaran_keitepi_5050umum,
			client.Pasaran_discbesar_5050umum,
			client.Pasaran_disckecil_5050umum,
			client.Pasaran_discgenap_5050umum,
			client.Pasaran_discganjil_5050umum,
			client.Pasaran_disctengah_5050umum,
			client.Pasaran_disctepi_5050umum,
			client.Pasaran_limitglobal_5050umum,
			client.Pasaran_limittotal_5050umum)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_5050")
		log.Printf("Redis Delete Client - CONF 5050 status: %d", val)
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
			result, err := models.Save_PasaranConf5050Umum(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_minbet_5050umum,
				client.Pasaran_maxbet_5050umum,
				client.Pasaran_keibesar_5050umum,
				client.Pasaran_keikecil_5050umum,
				client.Pasaran_keigenap_5050umum,
				client.Pasaran_keiganjil_5050umum,
				client.Pasaran_keitengah_5050umum,
				client.Pasaran_keitepi_5050umum,
				client.Pasaran_discbesar_5050umum,
				client.Pasaran_disckecil_5050umum,
				client.Pasaran_discgenap_5050umum,
				client.Pasaran_discganjil_5050umum,
				client.Pasaran_disctengah_5050umum,
				client.Pasaran_disctepi_5050umum,
				client.Pasaran_limitglobal_5050umum,
				client.Pasaran_limittotal_5050umum)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_5050")
			log.Printf("Redis Delete Client - CONF 5050 status: %d", val)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConf5050Special(c *fiber.Ctx) error {
	client := new(pasaranconfc5050special)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConf5050Special(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_minbet_5050special,
			client.Pasaran_maxbet_5050special,
			client.Pasaran_keiasganjil_5050special,
			client.Pasaran_keiasgenap_5050special,
			client.Pasaran_keiasbesar_5050special,
			client.Pasaran_keiaskecil_5050special,
			client.Pasaran_keikopganjil_5050special,
			client.Pasaran_keikopgenap_5050special,
			client.Pasaran_keikopbesar_5050special,
			client.Pasaran_keikopkecil_5050special,
			client.Pasaran_keikepalaganjil_5050special,
			client.Pasaran_keikepalagenap_5050special,
			client.Pasaran_keikepalabesar_5050special,
			client.Pasaran_keikepalakecil_5050special,
			client.Pasaran_keiekorganjil_5050special,
			client.Pasaran_keiekorgenap_5050special,
			client.Pasaran_keiekorbesar_5050special,
			client.Pasaran_keiekorkecil_5050special,
			client.Pasaran_discasganjil_5050special,
			client.Pasaran_discasgenap_5050special,
			client.Pasaran_discasbesar_5050special,
			client.Pasaran_discaskecil_5050special,
			client.Pasaran_disckopganjil_5050special,
			client.Pasaran_disckopgenap_5050special,
			client.Pasaran_disckopbesar_5050special,
			client.Pasaran_disckopkecil_5050special,
			client.Pasaran_disckepalaganjil_5050special,
			client.Pasaran_disckepalagenap_5050special,
			client.Pasaran_disckepalabesar_5050special,
			client.Pasaran_disckepalakecil_5050special,
			client.Pasaran_discekorganjil_5050special,
			client.Pasaran_discekorgenap_5050special,
			client.Pasaran_discekorbesar_5050special,
			client.Pasaran_discekorkecil_5050special,
			client.Pasaran_limitglobal_5050special,
			client.Pasaran_limittotal_5050special)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_5050")
		log.Printf("Redis Delete Client - CONF 5050 status: %d", val)
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
			result, err := models.Save_PasaranConf5050Special(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_minbet_5050special,
				client.Pasaran_maxbet_5050special,
				client.Pasaran_keiasganjil_5050special,
				client.Pasaran_keiasgenap_5050special,
				client.Pasaran_keiasbesar_5050special,
				client.Pasaran_keiaskecil_5050special,
				client.Pasaran_keikopganjil_5050special,
				client.Pasaran_keikopgenap_5050special,
				client.Pasaran_keikopbesar_5050special,
				client.Pasaran_keikopkecil_5050special,
				client.Pasaran_keikepalaganjil_5050special,
				client.Pasaran_keikepalagenap_5050special,
				client.Pasaran_keikepalabesar_5050special,
				client.Pasaran_keikepalakecil_5050special,
				client.Pasaran_keiekorganjil_5050special,
				client.Pasaran_keiekorgenap_5050special,
				client.Pasaran_keiekorbesar_5050special,
				client.Pasaran_keiekorkecil_5050special,
				client.Pasaran_discasganjil_5050special,
				client.Pasaran_discasgenap_5050special,
				client.Pasaran_discasbesar_5050special,
				client.Pasaran_discaskecil_5050special,
				client.Pasaran_disckopganjil_5050special,
				client.Pasaran_disckopgenap_5050special,
				client.Pasaran_disckopbesar_5050special,
				client.Pasaran_disckopkecil_5050special,
				client.Pasaran_disckepalaganjil_5050special,
				client.Pasaran_disckepalagenap_5050special,
				client.Pasaran_disckepalabesar_5050special,
				client.Pasaran_disckepalakecil_5050special,
				client.Pasaran_discekorganjil_5050special,
				client.Pasaran_discekorgenap_5050special,
				client.Pasaran_discekorbesar_5050special,
				client.Pasaran_discekorkecil_5050special,
				client.Pasaran_limitglobal_5050special,
				client.Pasaran_limittotal_5050special)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_5050")
			log.Printf("Redis Delete Client - CONF 5050 status: %d", val)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConf5050Kombinasi(c *fiber.Ctx) error {
	client := new(pasaranconfc5050kombinasi)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConf5050Kombinasi(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_minbet_5050kombinasi,
			client.Pasaran_maxbet_5050kombinasi,
			client.Pasaran_belakangkeimono_5050kombinasi,
			client.Pasaran_belakangkeistereo_5050kombinasi,
			client.Pasaran_belakangkeikembang_5050kombinasi,
			client.Pasaran_belakangkeikempis_5050kombinasi,
			client.Pasaran_belakangkeikembar_5050kombinasi,
			client.Pasaran_tengahkeimono_5050kombinasi,
			client.Pasaran_tengahkeistereo_5050kombinasi,
			client.Pasaran_tengahkeikembang_5050kombinasi,
			client.Pasaran_tengahkeikempis_5050kombinasi,
			client.Pasaran_tengahkeikembar_5050kombinasi,
			client.Pasaran_depankeimono_5050kombinasi,
			client.Pasaran_depankeistereo_5050kombinasi,
			client.Pasaran_depankeikembang_5050kombinasi,
			client.Pasaran_depankeikempis_5050kombinasi,
			client.Pasaran_depankeikembar_5050kombinasi,
			client.Pasaran_belakangdiscmono_5050kombinasi,
			client.Pasaran_belakangdiscstereo_5050kombinasi,
			client.Pasaran_belakangdisckembang_5050kombinasi,
			client.Pasaran_belakangdisckempis_5050kombinasi,
			client.Pasaran_belakangdisckembar_5050kombinasi,
			client.Pasaran_tengahdiscmono_5050kombinasi,
			client.Pasaran_tengahdiscstereo_5050kombinasi,
			client.Pasaran_tengahdisckembang_5050kombinasi,
			client.Pasaran_tengahdisckempis_5050kombinasi,
			client.Pasaran_tengahdisckembar_5050kombinasi,
			client.Pasaran_depandiscmono_5050kombinasi,
			client.Pasaran_depandiscstereo_5050kombinasi,
			client.Pasaran_depandisckembang_5050kombinasi,
			client.Pasaran_depandisckempis_5050kombinasi,
			client.Pasaran_depandisckembar_5050kombinasi,
			client.Pasaran_limitglobal_5050kombinasi,
			client.Pasaran_limittotal_5050kombinasi)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_5050")
		log.Printf("Redis Delete Client - CONF 5050 status: %d", val)
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
			result, err := models.Save_PasaranConf5050Kombinasi(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_minbet_5050kombinasi,
				client.Pasaran_maxbet_5050kombinasi,
				client.Pasaran_belakangkeimono_5050kombinasi,
				client.Pasaran_belakangkeistereo_5050kombinasi,
				client.Pasaran_belakangkeikembang_5050kombinasi,
				client.Pasaran_belakangkeikempis_5050kombinasi,
				client.Pasaran_belakangkeikembar_5050kombinasi,
				client.Pasaran_tengahkeimono_5050kombinasi,
				client.Pasaran_tengahkeistereo_5050kombinasi,
				client.Pasaran_tengahkeikembang_5050kombinasi,
				client.Pasaran_tengahkeikempis_5050kombinasi,
				client.Pasaran_tengahkeikembar_5050kombinasi,
				client.Pasaran_depankeimono_5050kombinasi,
				client.Pasaran_depankeistereo_5050kombinasi,
				client.Pasaran_depankeikembang_5050kombinasi,
				client.Pasaran_depankeikempis_5050kombinasi,
				client.Pasaran_depankeikembar_5050kombinasi,
				client.Pasaran_belakangdiscmono_5050kombinasi,
				client.Pasaran_belakangdiscstereo_5050kombinasi,
				client.Pasaran_belakangdisckembang_5050kombinasi,
				client.Pasaran_belakangdisckempis_5050kombinasi,
				client.Pasaran_belakangdisckembar_5050kombinasi,
				client.Pasaran_tengahdiscmono_5050kombinasi,
				client.Pasaran_tengahdiscstereo_5050kombinasi,
				client.Pasaran_tengahdisckembang_5050kombinasi,
				client.Pasaran_tengahdisckempis_5050kombinasi,
				client.Pasaran_tengahdisckembar_5050kombinasi,
				client.Pasaran_depandiscmono_5050kombinasi,
				client.Pasaran_depandiscstereo_5050kombinasi,
				client.Pasaran_depandisckembang_5050kombinasi,
				client.Pasaran_depandisckempis_5050kombinasi,
				client.Pasaran_depandisckembar_5050kombinasi,
				client.Pasaran_limitglobal_5050kombinasi,
				client.Pasaran_limittotal_5050kombinasi)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_5050")
			log.Printf("Redis Delete Client - CONF 5050 status: %d", val)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfMacauKombinasi(c *fiber.Ctx) error {
	client := new(pasaranconfmakaukombinasi)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConfMacauKombinasi(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_minbet_kombinasi,
			client.Pasaran_maxbet_kombinasi,
			client.Pasaran_win_kombinasi,
			client.Pasaran_disc_kombinasi,
			client.Pasaran_limitglobal_kombinasi,
			client.Pasaran_limittotal_kombinasi)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_macaukombinasi")
		log.Printf("Redis Delete Client - CONF MACAUKOMBINASI status: %d", val)
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
			result, err := models.Save_PasaranConfMacauKombinasi(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_minbet_kombinasi,
				client.Pasaran_maxbet_kombinasi,
				client.Pasaran_win_kombinasi,
				client.Pasaran_disc_kombinasi,
				client.Pasaran_limitglobal_kombinasi,
				client.Pasaran_limittotal_kombinasi)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_macaukombinasi")
			log.Printf("Redis Delete Client - CONF MACAUKOMBINASI status: %d", val)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfDasar(c *fiber.Ctx) error {
	client := new(pasaranconfdasar)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConfDasar(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_minbet_dasar,
			client.Pasaran_maxbet_dasar,
			client.Pasaran_keibesar_dasar,
			client.Pasaran_keikecil_dasar,
			client.Pasaran_keigenap_dasar,
			client.Pasaran_keiganjil_dasar,
			client.Pasaran_discbesar_dasar,
			client.Pasaran_disckecil_dasar,
			client.Pasaran_discgenap_dasar,
			client.Pasaran_discganjil_dasar,
			client.Pasaran_limitglobal_dasar,
			client.Pasaran_limittotal_dasar)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_dasar")
		log.Printf("Redis Delete Client - CONF DASAR status: %d", val)
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
			result, err := models.Save_PasaranConfDasar(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_minbet_dasar,
				client.Pasaran_maxbet_dasar,
				client.Pasaran_keibesar_dasar,
				client.Pasaran_keikecil_dasar,
				client.Pasaran_keigenap_dasar,
				client.Pasaran_keiganjil_dasar,
				client.Pasaran_discbesar_dasar,
				client.Pasaran_disckecil_dasar,
				client.Pasaran_discgenap_dasar,
				client.Pasaran_discganjil_dasar,
				client.Pasaran_limitglobal_dasar,
				client.Pasaran_limittotal_dasar)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_dasar")
			log.Printf("Redis Delete Client - CONF DASAR status: %d", val)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfShio(c *fiber.Ctx) error {
	client := new(pasaranconfshio)
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
	client_username, client_company, typeadmin, idruleadmin := helpers.Parsing_Decry(temp_decp, "==")

	ruleadmin := models.Get_AdminRule(client_company, "ruleadmin", idruleadmin)
	flag_page := models.Get_listitemsearch(ruleadmin, ",", client.Page)

	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConfShio(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_shioyear_shio,
			client.Pasaran_minbet_shio,
			client.Pasaran_maxbet_shio,
			client.Pasaran_win_shio,
			client.Pasaran_disc_shio,
			client.Pasaran_limitglobal_shio,
			client.Pasaran_limittotal_shio)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_shio")
		log.Printf("Redis Delete Client - CONF SHIO status: %d", val)
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
			result, err := models.Save_PasaranConfShio(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_shioyear_shio,
				client.Pasaran_minbet_shio,
				client.Pasaran_maxbet_shio,
				client.Pasaran_win_shio,
				client.Pasaran_disc_shio,
				client.Pasaran_limitglobal_shio,
				client.Pasaran_limittotal_shio)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val := helpers.DeleteRedis("config_" + client_company + "_" + client.Idpasarantogel + "_shio")
			log.Printf("Redis Delete Client - CONF SHIO status: %d", val)
			return c.JSON(result)
		}
	}
}
