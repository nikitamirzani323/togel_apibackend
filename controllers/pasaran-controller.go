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
	Pasaran_limitline3dd int    `json:"pasaran_limitline3dd"`
	Pasaran_limitline2d  int    `json:"pasaran_limitline2d"`
	Pasaran_limitline2dd int    `json:"pasaran_limitline2dd"`
	Pasaran_limitline2dt int    `json:"pasaran_limitline2dt"`
}
type pasaranconf432 struct {
	Idpasaran                   int     `json:"idpasaran"`
	Idpasarantogel              string  `json:"idpasarantogel"`
	Page                        string  `json:"page"`
	Pasaran_minbet_432d         int     `json:"pasaran_minbet_432d" validate:"required,numeric"`
	Pasaran_maxbet4d_432d       int     `json:"pasaran_maxbet4d_432d" validate:"required,numeric"`
	Pasaran_maxbet3d_432d       int     `json:"pasaran_maxbet3d_432d" validate:"required,numeric"`
	Pasaran_maxbet3dd_432d      int     `json:"pasaran_maxbet3dd_432d" validate:"required,numeric"`
	Pasaran_maxbet2d_432d       int     `json:"pasaran_maxbet2d_432d" validate:"required,numeric"`
	Pasaran_maxbet2dd_432d      int     `json:"pasaran_maxbet2dd_432d" validate:"required,numeric"`
	Pasaran_maxbet2dt_432d      int     `json:"pasaran_maxbet2dt_432d" validate:"required,numeric"`
	Pasaran_limitotal4d_432d    int     `json:"pasaran_limitotal4d_432d" validate:"required,numeric"`
	Pasaran_limitotal3d_432d    int     `json:"pasaran_limitotal3d_432d" validate:"required,numeric"`
	Pasaran_limitotal3dd_432d   int     `json:"pasaran_limitotal3dd_432d" validate:"required,numeric"`
	Pasaran_limitotal2d_432d    int     `json:"pasaran_limitotal2d_432d" validate:"required,numeric"`
	Pasaran_limitotal2dd_432d   int     `json:"pasaran_limitotal2dd_432d" validate:"required,numeric"`
	Pasaran_limitotal2dt_432d   int     `json:"pasaran_limitotal2dt_432d" validate:"required,numeric"`
	Pasaran_limitglobal4d_432d  int     `json:"pasaran_limitglobal4d_432d" validate:"required,numeric"`
	Pasaran_limitglobal3d_432d  int     `json:"pasaran_limitglobal3d_432d" validate:"required,numeric"`
	Pasaran_limitglobal3dd_432d int     `json:"pasaran_limitglobal3dd_432d" validate:"required,numeric"`
	Pasaran_limitglobal2d_432d  int     `json:"pasaran_limitglobal2d_432d" validate:"required,numeric"`
	Pasaran_limitglobal2dd_432d int     `json:"pasaran_limitglobal2dd_432d" validate:"required,numeric"`
	Pasaran_limitglobal2dt_432d int     `json:"pasaran_limitglobal2dt_432d" validate:"required,numeric"`
	Pasaran_win4d_432d          int     `json:"pasaran_win4d_432d" validate:"required,numeric"`
	Pasaran_win3d_432d          int     `json:"pasaran_win3d_432d" validate:"required,numeric"`
	Pasaran_win3dd_432d         int     `json:"pasaran_win3dd_432d" validate:"required,numeric"`
	Pasaran_win2d_432d          int     `json:"pasaran_win2d_432d" validate:"required,numeric"`
	Pasaran_win2dd_432d         int     `json:"pasaran_win2dd_432d" validate:"required,numeric"`
	Pasaran_win2dt_432d         int     `json:"pasaran_win2dt_432d" validate:"required,numeric"`
	Pasaran_win4dnodisc_432d    int     `json:"pasaran_Win4dnodisc_432d" validate:"required,numeric"`
	Pasaran_win3dnodisc_432d    int     `json:"pasaran_Win3dnodisc_432d" validate:"required,numeric"`
	Pasaran_win3ddnodisc_432d   int     `json:"pasaran_win3ddnodisc_432d" validate:"required,numeric"`
	Pasaran_win2dnodisc_432d    int     `json:"pasaran_win2dnodisc_432d" validate:"required,numeric"`
	Pasaran_win2ddnodisc_432d   int     `json:"pasaran_win2ddnodisc_432d" validate:"required,numeric"`
	Pasaran_win2dtnodisc_432d   int     `json:"pasaran_win2dtnodisc_432d" validate:"required,numeric"`
	Pasaran_win4dbb_kena_432d   int     `json:"pasaran_Win4dbb_kena_432d" validate:"required,numeric"`
	Pasaran_win3dbb_kena_432d   int     `json:"pasaran_Win3dbb_kena_432d" validate:"required,numeric"`
	Pasaran_win3ddbb_kena_432d  int     `json:"pasaran_win3ddbb_kena_432d" validate:"required,numeric"`
	Pasaran_win2dbb_kena_432d   int     `json:"pasaran_win2dbb_kena_432d" validate:"required,numeric"`
	Pasaran_win2ddbb_kena_432d  int     `json:"pasaran_win2ddbb_kena_432d" validate:"required,numeric"`
	Pasaran_win2dtbb_kena_432d  int     `json:"pasaran_win2dtbb_kena_432d" validate:"required,numeric"`
	Pasaran_win4dbb_432d        int     `json:"pasaran_win4dbb_432d" validate:"required,numeric"`
	Pasaran_win3dbb_432d        int     `json:"pasaran_win3dbb_432d" validate:"required,numeric"`
	Pasaran_win3ddbb_432d       int     `json:"pasaran_win3ddbb_432d" validate:"required,numeric"`
	Pasaran_win2dbb_432d        int     `json:"pasaran_win2dbb_432d" validate:"required,numeric"`
	Pasaran_win2ddbb_432d       int     `json:"pasaran_win2ddbb_432d" validate:"required,numeric"`
	Pasaran_win2dtbb_432d       int     `json:"pasaran_win2dtbb_432d" validate:"required,numeric"`
	Pasaran_disc4d_432d         float32 `json:"pasaran_disc4d_432d" validate:"required,numeric"`
	Pasaran_disc3d_432d         float32 `json:"pasaran_disc3d_432d" validate:"required,numeric"`
	Pasaran_disc3dd_432d        float32 `json:"pasaran_disc3dd_432d" validate:"required,numeric"`
	Pasaran_disc2d_432d         float32 `json:"pasaran_disc2d_432d" validate:"required,numeric"`
	Pasaran_disc2dd_432d        float32 `json:"pasaran_disc2dd_432d" validate:"required,numeric"`
	Pasaran_disc2dt_432d        float32 `json:"pasaran_disc2dt_432d" validate:"required,numeric"`
}
type pasaranconfcbebas struct {
	Idpasaran                  int     `json:"idpasaran"`
	Idpasarantogel             string  `json:"idpasarantogel"`
	Page                       string  `json:"page"`
	Pasaran_minbet_cbebas      int     `json:"pasaran_minbet_cbebas" validate:"required,numeric"`
	Pasaran_maxbet_cbebas      int     `json:"pasaran_maxbet_cbebas" validate:"required,numeric"`
	Pasaran_limitotal_cbebas   int     `json:"pasaran_limitotal_cbebas" validate:"required,numeric"`
	Pasaran_limitglobal_cbebas int     `json:"pasaran_limitglobal_cbebas" validate:"required,numeric"`
	Pasaran_win_cbebas         float32 `json:"pasaran_win_cbebas" validate:"required,numeric"`
	Pasaran_disc_cbebas        float32 `json:"pasaran_disc_cbebas" validate:"required,numeric"`
}
type pasaranconfcmacau struct {
	Idpasaran                  int     `json:"idpasaran"`
	Idpasarantogel             string  `json:"idpasarantogel"`
	Page                       string  `json:"page"`
	Pasaran_minbet_cmacau      int     `json:"pasaran_minbet_cmacau" validate:"required,numeric"`
	Pasaran_maxbet_cmacau      int     `json:"pasaran_maxbet_cmacau" validate:"required,numeric"`
	Pasaran_limitotal_cmacau   int     `json:"pasaran_limitotal_cmacau" validate:"required,numeric"`
	Pasaran_limitglobal_cmacau int     `json:"pasaran_limitglobal_cmacau" validate:"required,numeric"`
	Pasaran_win2_cmacau        float32 `json:"pasaran_win2_cmacau" validate:"required,numeric"`
	Pasaran_win3_cmacau        float32 `json:"pasaran_win3_cmacau" validate:"required,numeric"`
	Pasaran_win4_cmacau        float32 `json:"pasaran_win4_cmacau" validate:"required,numeric"`
	Pasaran_disc_cmacau        float32 `json:"pasaran_disc_cmacau" validate:"required,numeric"`
}
type pasaranconfcnaga struct {
	Idpasaran                 int     `json:"idpasaran"`
	Idpasarantogel            string  `json:"idpasarantogel"`
	Page                      string  `json:"page"`
	Pasaran_minbet_cnaga      int     `json:"pasaran_minbet_cnaga" validate:"required,numeric"`
	Pasaran_maxbet_cnaga      int     `json:"pasaran_maxbet_cnaga" validate:"required,numeric"`
	Pasaran_limittotal_cnaga  int     `json:"pasaran_limittotal_cnaga" validate:"required,numeric"`
	Pasaran_limitglobal_cnaga int     `json:"pasaran_limitglobal_cnaga" validate:"required,numeric"`
	Pasaran_win3_cnaga        float32 `json:"pasaran_win3_cnaga" validate:"required,numeric"`
	Pasaran_win4_cnaga        float32 `json:"pasaran_win4_cnaga" validate:"required,numeric"`
	Pasaran_disc_cnaga        float32 `json:"pasaran_disc_cnaga" validate:"required,numeric"`
}
type pasaranconfcjitu struct {
	Idpasaran                 int     `json:"idpasaran"`
	Idpasarantogel            string  `json:"idpasarantogel"`
	Page                      string  `json:"page"`
	Pasaran_minbet_cjitu      int     `json:"pasaran_minbet_cjitu" validate:"required,numeric"`
	Pasaran_maxbet_cjitu      int     `json:"pasaran_maxbet_cjitu" validate:"required,numeric"`
	Pasaran_limittotal_cjitu  int     `json:"pasaran_limittotal_cjitu" validate:"required,numeric"`
	Pasaran_limitglobal_cjitu int     `json:"pasaran_limitglobal_cjitu" validate:"required,numeric"`
	Pasaran_winas_cjitu       float32 `json:"pasaran_winas_cjitu" validate:"required,numeric"`
	Pasaran_winkop_cjitu      float32 `json:"pasaran_winkop_cjitu" validate:"required,numeric"`
	Pasaran_winkepala_cjitu   float32 `json:"pasaran_winkepala_cjitu" validate:"required,numeric"`
	Pasaran_winekor_cjitu     float32 `json:"pasaran_winekor_cjitu" validate:"required,numeric"`
	Pasaran_desc_cjitu        float32 `json:"pasaran_desc_cjitu" validate:"required,numeric"`
}
type pasaranconfc5050 struct {
	Idpasaran                    int     `json:"idpasaran"`
	Idpasarantogel               string  `json:"idpasarantogel"`
	Page                         string  `json:"page"`
	Pasaran_minbet_5050umum      int     `json:"pasaran_minbet_5050umum" validate:"required,numeric"`
	Pasaran_maxbet_5050umum      int     `json:"pasaran_maxbet_5050umum" validate:"required,numeric"`
	Pasaran_limittotal_5050umum  int     `json:"pasaran_limittotal_5050umum" validate:"required,numeric"`
	Pasaran_limitglobal_5050umum int     `json:"pasaran_limitglobal_5050umum" validate:"required,numeric"`
	Pasaran_keibesar_5050umum    float32 `json:"pasaran_keibesar_5050umum" validate:"required,numeric"`
	Pasaran_keikecil_5050umum    float32 `json:"pasaran_keikecil_5050umum" validate:"required,numeric"`
	Pasaran_keigenap_5050umum    float32 `json:"pasaran_keigenap_5050umum" validate:"required,numeric"`
	Pasaran_keiganjil_5050umum   float32 `json:"pasaran_keiganjil_5050umum" validate:"required,numeric"`
	Pasaran_keitengah_5050umum   float32 `json:"pasaran_keitengah_5050umum" validate:"required,numeric"`
	Pasaran_keitepi_5050umum     float32 `json:"pasaran_keitepi_5050umum" validate:"required,numeric"`
	Pasaran_discbesar_5050umum   float32 `json:"pasaran_discbesar_5050umum" validate:"numeric"`
	Pasaran_disckecil_5050umum   float32 `json:"pasaran_disckecil_5050umum" validate:"numeric"`
	Pasaran_discgenap_5050umum   float32 `json:"pasaran_discgenap_5050umum" validate:"numeric"`
	Pasaran_discganjil_5050umum  float32 `json:"pasaran_discganjil_5050umum" validate:"numeric"`
	Pasaran_disctengah_5050umum  float32 `json:"pasaran_disctengah_5050umum" validate:"numeric"`
	Pasaran_disctepi_5050umum    float32 `json:"pasaran_disctepi_5050umum" validate:"numeric"`
}
type pasaranconfc5050special struct {
	Idpasaran                            int     `json:"idpasaran"`
	Idpasarantogel                       string  `json:"idpasarantogel"`
	Page                                 string  `json:"page"`
	Pasaran_minbet_5050special           int     `json:"pasaran_minbet_5050special" validate:"required,numeric"`
	Pasaran_maxbet_5050special           int     `json:"pasaran_maxbet_5050special" validate:"required,numeric"`
	Pasaran_limittotal_5050special       int     `json:"pasaran_limittotal_5050special" validate:"required,numeric"`
	Pasaran_limitglobal_5050special      int     `json:"pasaran_limitglobal_5050special" validate:"required,numeric"`
	Pasaran_keiasganjil_5050special      float32 `json:"pasaran_keiasganjil_5050special" validate:"required,numeric"`
	Pasaran_keiasgenap_5050special       float32 `json:"pasaran_keiasgenap_5050special" validate:"required,numeric"`
	Pasaran_keiasbesar_5050special       float32 `json:"pasaran_keiasbesar_5050special" validate:"required,numeric"`
	Pasaran_keiaskecil_5050special       float32 `json:"pasaran_keiaskecil_5050special" validate:"required,numeric"`
	Pasaran_keikopganjil_5050special     float32 `json:"pasaran_keikopganjil_5050special" validate:"required,numeric"`
	Pasaran_keikopgenap_5050special      float32 `json:"pasaran_keikopgenap_5050special" validate:"required,numeric"`
	Pasaran_keikopbesar_5050special      float32 `json:"pasaran_keikopbesar_5050special" validate:"required,numeric"`
	Pasaran_keikopkecil_5050special      float32 `json:"pasaran_keikopkecil_5050special" validate:"required,numeric"`
	Pasaran_keikepalaganjil_5050special  float32 `json:"pasaran_keikepalaganjil_5050special" validate:"required,numeric"`
	Pasaran_keikepalagenap_5050special   float32 `json:"pasaran_keikepalagenap_5050special" validate:"required,numeric"`
	Pasaran_keikepalabesar_5050special   float32 `json:"pasaran_keikepalabesar_5050special" validate:"required,numeric"`
	Pasaran_keikepalakecil_5050special   float32 `json:"pasaran_keikepalakecil_5050special" validate:"required,numeric"`
	Pasaran_keiekorganjil_5050special    float32 `json:"pasaran_keiekorganjil_5050special" validate:"required,numeric"`
	Pasaran_keiekorgenap_5050special     float32 `json:"pasaran_keiekorgenap_5050special" validate:"required,numeric"`
	Pasaran_keiekorbesar_5050special     float32 `json:"pasaran_keiekorbesar_5050special" validate:"required,numeric"`
	Pasaran_keiekorkecil_5050special     float32 `json:"pasaran_keiekorkecil_5050special" validate:"required,numeric"`
	Pasaran_discasganjil_5050special     float32 `json:"pasaran_discasganjil_5050special" validate:"numeric"`
	Pasaran_discasgenap_5050special      float32 `json:"pasaran_discasgenap_5050special" validate:"numeric"`
	Pasaran_discasbesar_5050special      float32 `json:"pasaran_discasbesar_5050special" validate:"numeric"`
	Pasaran_discaskecil_5050special      float32 `json:"pasaran_discaskecil_5050special" validate:"numeric"`
	Pasaran_disckopganjil_5050special    float32 `json:"pasaran_disckopganjil_5050special" validate:"numeric"`
	Pasaran_disckopgenap_5050special     float32 `json:"pasaran_disckopgenap_5050special" validate:"numeric"`
	Pasaran_disckopbesar_5050special     float32 `json:"pasaran_disckopbesar_5050special" validate:"numeric"`
	Pasaran_disckopkecil_5050special     float32 `json:"pasaran_disckopkecil_5050special" validate:"numeric"`
	Pasaran_disckepalaganjil_5050special float32 `json:"pasaran_disckepalaganjil_5050special" validate:"numeric"`
	Pasaran_disckepalagenap_5050special  float32 `json:"pasaran_disckepalagenap_5050special" validate:"numeric"`
	Pasaran_disckepalabesar_5050special  float32 `json:"pasaran_disckepalabesar_5050special" validate:"numeric"`
	Pasaran_disckepalakecil_5050special  float32 `json:"pasaran_disckepalakecil_5050special" validate:"numeric"`
	Pasaran_discekorganjil_5050special   float32 `json:"pasaran_discekorganjil_5050special" validate:"numeric"`
	Pasaran_discekorgenap_5050special    float32 `json:"pasaran_discekorgenap_5050special" validate:"numeric"`
	Pasaran_discekorbesar_5050special    float32 `json:"pasaran_discekorbesar_5050special" validate:"numeric"`
	Pasaran_discekorkecil_5050special    float32 `json:"pasaran_discekorkecil_5050special" validate:"numeric"`
}
type pasaranconfc5050kombinasi struct {
	Idpasaran                                 int     `json:"idpasaran"`
	Idpasarantogel                            string  `json:"idpasarantogel"`
	Page                                      string  `json:"page"`
	Pasaran_minbet_5050kombinasi              int     `json:"pasaran_minbet_5050kombinasi" validate:"required,numeric"`
	Pasaran_maxbet_5050kombinasi              int     `json:"pasaran_maxbet_5050kombinasi" validate:"required,numeric"`
	Pasaran_limittotal_5050kombinasi          int     `json:"pasaran_limittotal_5050kombinasi" validate:"required,numeric"`
	Pasaran_limitglobal_5050kombinasi         int     `json:"pasaran_limitglobal_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeimono_5050kombinasi     float32 `json:"pasaran_belakangkeimono_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeistereo_5050kombinasi   float32 `json:"pasaran_belakangkeistereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeikembang_5050kombinasi  float32 `json:"pasaran_belakangkeikembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeikempis_5050kombinasi   float32 `json:"pasaran_belakangkeikempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeikembar_5050kombinasi   float32 `json:"pasaran_belakangkeikembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeimono_5050kombinasi       float32 `json:"pasaran_tengahkeimono_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeistereo_5050kombinasi     float32 `json:"pasaran_tengahkeistereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeikembang_5050kombinasi    float32 `json:"pasaran_tengahkeikembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeikempis_5050kombinasi     float32 `json:"pasaran_tengahkeikempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeikembar_5050kombinasi     float32 `json:"pasaran_tengahkeikembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeimono_5050kombinasi        float32 `json:"pasaran_depankeimono_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeistereo_5050kombinasi      float32 `json:"pasaran_depankeistereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeikembang_5050kombinasi     float32 `json:"pasaran_depankeikembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeikempis_5050kombinasi      float32 `json:"pasaran_depankeikempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeikembar_5050kombinasi      float32 `json:"pasaran_depankeikembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangdiscmono_5050kombinasi    float32 `json:"pasaran_belakangdiscmono_5050kombinasi" validate:"numeric"`
	Pasaran_belakangdiscstereo_5050kombinasi  float32 `json:"pasaran_belakangdiscstereo_5050kombinasi" validate:"numeric"`
	Pasaran_belakangdisckembang_5050kombinasi float32 `json:"pasaran_belakangdisckembang_5050kombinasi" validate:"numeric"`
	Pasaran_belakangdisckempis_5050kombinasi  float32 `json:"pasaran_belakangdisckempis_5050kombinasi" validate:"numeric"`
	Pasaran_belakangdisckembar_5050kombinasi  float32 `json:"pasaran_belakangdisckembar_5050kombinasi" validate:"numeric"`
	Pasaran_tengahdiscmono_5050kombinasi      float32 `json:"pasaran_tengahdiscmono_5050kombinasi" validate:"numeric"`
	Pasaran_tengahdiscstereo_5050kombinasi    float32 `json:"pasaran_tengahdiscstereo_5050kombinasi" validate:"numeric"`
	Pasaran_tengahdisckembang_5050kombinasi   float32 `json:"pasaran_tengahdisckembang_5050kombinasi" validate:"numeric"`
	Pasaran_tengahdisckempis_5050kombinasi    float32 `json:"pasaran_tengahdisckempis_5050kombinasi" validate:"numeric"`
	Pasaran_tengahdisckembar_5050kombinasi    float32 `json:"pasaran_tengahdisckembar_5050kombinasi" validate:"numeric"`
	Pasaran_depandiscmono_5050kombinasi       float32 `json:"pasaran_depandiscmono_5050kombinasi" validate:"numeric"`
	Pasaran_depandiscstereo_5050kombinasi     float32 `json:"pasaran_depandiscstereo_5050kombinasi" validate:"numeric"`
	Pasaran_depandisckembang_5050kombinasi    float32 `json:"pasaran_depandisckembang_5050kombinasi" validate:"numeric"`
	Pasaran_depandisckempis_5050kombinasi     float32 `json:"pasaran_depandisckempis_5050kombinasi" validate:"numeric"`
	Pasaran_depandisckembar_5050kombinasi     float32 `json:"pasaran_depandisckembar_5050kombinasi" validate:"numeric"`
}
type pasaranconfmakaukombinasi struct {
	Idpasaran                     int     `json:"idpasaran"`
	Idpasarantogel                string  `json:"idpasarantogel"`
	Page                          string  `json:"page"`
	Pasaran_minbet_kombinasi      int     `json:"pasaran_minbet_kombinasi" validate:"required,numeric"`
	Pasaran_maxbet_kombinasi      int     `json:"pasaran_maxbet_kombinasi" validate:"required,numeric"`
	Pasaran_limittotal_kombinasi  int     `json:"pasaran_limittotal_kombinasi" validate:"required,numeric"`
	Pasaran_limitglobal_kombinasi int     `json:"pasaran_limitglobal_kombinasi" validate:"required,numeric"`
	Pasaran_win_kombinasi         float32 `json:"pasaran_win_kombinasi" validate:"required,numeric"`
	Pasaran_disc_kombinasi        float32 `json:"pasaran_disc_kombinasi" validate:"required,numeric"`
}
type pasaranconfdasar struct {
	Idpasaran                 int     `json:"idpasaran"`
	Idpasarantogel            string  `json:"idpasarantogel"`
	Page                      string  `json:"page"`
	Pasaran_minbet_dasar      int     `json:"pasaran_minbet_dasar" validate:"required,numeric"`
	Pasaran_maxbet_dasar      int     `json:"pasaran_maxbet_dasar" validate:"required,numeric"`
	Pasaran_limittotal_dasar  int     `json:"pasaran_limittotal_dasar" validate:"required,numeric"`
	Pasaran_limitglobal_dasar int     `json:"pasaran_limitglobal_dasar" validate:"required,numeric"`
	Pasaran_keibesar_dasar    float32 `json:"pasaran_keibesar_dasar" validate:"numeric"`
	Pasaran_keikecil_dasar    float32 `json:"pasaran_keikecil_dasar" validate:"numeric"`
	Pasaran_keigenap_dasar    float32 `json:"pasaran_keigenap_dasar" validate:"numeric"`
	Pasaran_keiganjil_dasar   float32 `json:"pasaran_keiganjil_dasar" validate:"numeric"`
	Pasaran_discbesar_dasar   float32 `json:"pasaran_discbesar_dasar" validate:"numeric"`
	Pasaran_disckecil_dasar   float32 `json:"pasaran_disckecil_dasar" validate:"numeric"`
	Pasaran_discgenap_dasar   float32 `json:"pasaran_discgenap_dasar" validate:"numeric"`
	Pasaran_discganjil_dasar  float32 `json:"pasaran_discganjil_dasar" validate:"numeric"`
}
type pasaranconfshio struct {
	Idpasaran                int     `json:"idpasaran"`
	Idpasarantogel           string  `json:"idpasarantogel"`
	Page                     string  `json:"page"`
	Pasaran_minbet_shio      int     `json:"pasaran_minbet_shio" validate:"required,numeric"`
	Pasaran_maxbet_shio      int     `json:"pasaran_maxbet_shio" validate:"required,numeric"`
	Pasaran_limittotal_shio  int     `json:"pasaran_limittotal_shio" validate:"required,numeric"`
	Pasaran_limitglobal_shio int     `json:"pasaran_limitglobal_shio" validate:"required,numeric"`
	Pasaran_shioyear_shio    string  `json:"pasaran_shioyear_shio" validate:"required,numeric"`
	Pasaran_disc_shio        float32 `json:"pasaran_disc_shio" validate:"numeric"`
	Pasaran_win_shio         float32 `json:"pasaran_win_shio" validate:"required,numeric"`
}

type responseredis_pasarandetail struct {
	Idpasarantogel                    string  `json:"idpasarantogel"`
	Nmpasarantogel                    string  `json:"nmpasaran"`
	PasaranDiundi                     string  `json:"pasarandiundi"`
	PasaranURL                        string  `json:"pasaranurl"`
	Jamtutup                          string  `json:"jamtutup"`
	Jamjadwal                         string  `json:"jamjadwal"`
	Jamopen                           string  `json:"jamopen"`
	Limitline4d                       int     `json:"limitline_4d"`
	Limitline3d                       int     `json:"limitline_3d"`
	Limitline3dd                      int     `json:"limitline_3dd"`
	Limitline2d                       int     `json:"limitline_2d"`
	Limitline2dd                      int     `json:"limitline_2dd"`
	Limitline2dt                      int     `json:"limitline_2dt"`
	Bbfs                              int     `json:"bbfs"`
	Minbet_432d                       float32 `json:"minbet_432d"`
	Maxbet4d_432d                     float32 `json:"maxbet4d_432d"`
	Maxbet3d_432d                     float32 `json:"maxbet3d_432d"`
	Maxbet3dd_432d                    float32 `json:"maxbet3dd_432d"`
	Maxbet2d_432d                     float32 `json:"maxbet2d_432d"`
	Maxbet2dd_432d                    float32 `json:"maxbet2dd_432d"`
	Maxbet2dt_432d                    float32 `json:"maxbet2dt_432d"`
	Limitotal4d_432d                  float32 `json:"limitotal4d_432d"`
	Limitotal3d_432d                  float32 `json:"limitotal3d_432d"`
	Limitotal3dd_432d                 float32 `json:"limitotal3dd_432d"`
	Limitotal2d_432d                  float32 `json:"limitotal2d_432d"`
	Limitotal2dd_432d                 float32 `json:"limitotal2dd_432d"`
	Limitotal2dt_432d                 float32 `json:"limitotal2dt_432d"`
	Limitglobal4d_432d                float32 `json:"limitglobal4d_432d"`
	Limitglobal3d_432d                float32 `json:"limitglobal3d_432d"`
	Limitglobal3dd_432d               float32 `json:"limitglobal3dd_432d"`
	Limitglobal2d_432d                float32 `json:"limitglobal2d_432d"`
	Limitglobal2dd_432d               float32 `json:"limitglobal2dd_432d"`
	Limitglobal2dt_432d               float32 `json:"limitglobal2dt_432d"`
	Disc4d_432d                       float32 `json:"disc4d_432d"`
	Disc3d_432d                       float32 `json:"disc3d_432d"`
	Disc3dd_432d                      float32 `json:"disc3dd_432d"`
	Disc2d_432d                       float32 `json:"disc2d_432d"`
	Disc2dd_432d                      float32 `json:"disc2dd_432d"`
	Disc2dt_432d                      float32 `json:"disc2dt_432d"`
	Win4d_432d                        float32 `json:"win4d_432d"`
	Win3d_432d                        float32 `json:"win3d_432d"`
	Win3dd_432d                       float32 `json:"win3dd_432d"`
	Win2d_432d                        float32 `json:"win2d_432d"`
	Win2dd_432d                       float32 `json:"win2dd_432d"`
	Win2dt_432d                       float32 `json:"win2dt_432d"`
	Win4dnodisc_432d                  float32 `json:"win4dnodisc_432d"`
	Win3dnodisc_432d                  float32 `json:"win3dnodisc_432d"`
	Win3ddnodisc_432d                 float32 `json:"win3ddnodisc_432d"`
	Win2dnodisc_432d                  float32 `json:"win2dnodisc_432d"`
	Win2ddnodisc_432d                 float32 `json:"win2ddnodisc_432d"`
	Win2dtnodisc_432d                 float32 `json:"win2dtnodisc_432d"`
	Win4dbb_kena_432d                 float32 `json:"win4dbb_kena_432d"`
	Win3dbb_kena_432d                 float32 `json:"win3dbb_kena_432d"`
	Win3ddbb_kena_432d                float32 `json:"win3ddbb_kena_432d"`
	Win2dbb_kena_432d                 float32 `json:"win2dbb_kena_432d"`
	Win2ddbb_kena_432d                float32 `json:"win2ddbb_kena_432d"`
	Win2dtbb_kena_432d                float32 `json:"win2dtbb_kena_432d"`
	Win4dbb_432d                      float32 `json:"win4dbb_432d"`
	Win3dbb_432d                      float32 `json:"win3dbb_432d"`
	Win3ddbb_432d                     float32 `json:"win3ddbb_432d"`
	Win2dbb_432d                      float32 `json:"win2dbb_432d"`
	Win2ddbb_432d                     float32 `json:"win2ddbb_432d"`
	Win2dtbb_432d                     float32 `json:"win2dtbb_432d"`
	Minbet_cbebas                     float32 `json:"minbet_cbebas"`
	Maxbet_cbebas                     float32 `json:"maxbet_cbebas"`
	Win_cbebas                        float32 `json:"win_cbebas"`
	Disc_cbebas                       float32 `json:"disc_cbebas"`
	Limitglobal_cbebas                float32 `json:"limitglobal_cbebas"`
	Limittotal_cbebas                 float32 `json:"limittotal_cbebas"`
	Minbet_cmacau                     float32 `json:"minbet_cmacau"`
	Maxbet_cmacau                     float32 `json:"maxbet_cmacau"`
	Win2d_cmacau                      float32 `json:"win2d_cmacau"`
	Win3d_cmacau                      float32 `json:"win3d_cmacau"`
	Win4d_cmacau                      float32 `json:"win4d_cmacau"`
	Disc_cmacau                       float32 `json:"disc_cmacau"`
	Limitglobal_cmacau                float32 `json:"limitglobal_cmacau"`
	Limitotal_cmacau                  float32 `json:"limitotal_cmacau"`
	Minbet_cnaga                      float32 `json:"minbet_cnaga"`
	Maxbet_cnaga                      float32 `json:"maxbet_cnaga"`
	Win3_cnaga                        float32 `json:"win3_cnaga"`
	Win4_cnaga                        float32 `json:"win4_cnaga"`
	Disc_cnaga                        float32 `json:"disc_cnaga"`
	Limitglobal_cnaga                 float32 `json:"limitglobal_cnaga"`
	Limittotal_cnaga                  float32 `json:"limittotal_cnaga"`
	Minbet_cjitu                      float32 `json:"minbet_cjitu"`
	Maxbet_cjitu                      float32 `json:"maxbet_cjitu"`
	Winas_cjitu                       float32 `json:"winas_cjitu"`
	Winkop_cjitu                      float32 `json:"winkop_cjitu"`
	Winkepala_cjitu                   float32 `json:"winkepala_cjitu"`
	Winekor_cjitu                     float32 `json:"winekor_cjitu"`
	Desc_cjitu                        float32 `json:"desc_cjitu"`
	Limitglobal_cjitu                 float32 `json:"limitglobal_cjitu"`
	Limittotal_cjitu                  float32 `json:"limittotal_cjitu"`
	Minbet_5050umum                   float32 `json:"minbet_5050umum"`
	Maxbet_5050umum                   float32 `json:"maxbet_5050umum"`
	Keibesar_5050umum                 float32 `json:"keibesar_5050umum"`
	Keikecil_5050umum                 float32 `json:"keikecil_5050umum"`
	Keigenap_5050umum                 float32 `json:"keigenap_5050umum"`
	Keiganjil_5050umum                float32 `json:"keiganjil_5050umum"`
	Keitengah_5050umum                float32 `json:"keitengah_5050umum"`
	Keitepi_5050umum                  float32 `json:"keitepi_5050umum"`
	Discbesar_5050umum                float32 `json:"discbesar_5050umum"`
	Disckecil_5050umum                float32 `json:"disckecil_5050umum"`
	Discgenap_5050umum                float32 `json:"discgenap_5050umum"`
	Discganjil_5050umum               float32 `json:"discganjil_5050umum"`
	Disctengah_5050umum               float32 `json:"disctengah_5050umum"`
	Disctepi_5050umum                 float32 `json:"disctepi_5050umum"`
	Limitglobal_5050umum              float32 `json:"limitglobal_5050umum"`
	Limittotal_5050umum               float32 `json:"limittotal_5050umum"`
	Minbet_5050special                float32 `json:"minbet_5050special"`
	Maxbet_5050special                float32 `json:"maxbet_5050special"`
	Keiasganjil_5050special           float32 `json:"keiasganjil_5050special"`
	Keiasgenap_5050special            float32 `json:"keiasgenap_5050special"`
	Keiasbesar_5050special            float32 `json:"keiasbesar_5050special"`
	Keiaskecil_5050special            float32 `json:"keiaskecil_5050special"`
	Keikopganjil_5050special          float32 `json:"keikopganjil_5050special"`
	Keikopgenap_5050special           float32 `json:"keikopgenap_5050special"`
	Keikopbesar_5050special           float32 `json:"keikopbesar_5050special"`
	Keikopkecil_5050special           float32 `json:"keikopkecil_5050special"`
	Keikepalaganjil_5050special       float32 `json:"keikepalaganjil_5050special"`
	Keikepalagenap_5050special        float32 `json:"keikepalagenap_5050special"`
	Keikepalabesar_5050special        float32 `json:"keikepalabesar_5050special"`
	Keikepalakecil_5050special        float32 `json:"keikepalakecil_5050special"`
	Keiekorganjil_5050special         float32 `json:"keiekorganjil_5050special"`
	Keiekorgenap_5050special          float32 `json:"keiekorgenap_5050special"`
	Keiekorbesar_5050special          float32 `json:"keiekorbesar_5050special"`
	Keiekorkecil_5050special          float32 `json:"keiekorkecil_5050special"`
	Discasganjil_5050special          float32 `json:"discasganjil_5050special"`
	Discasgenap_5050special           float32 `json:"discasgenap_5050special"`
	Discasbesar_5050special           float32 `json:"discasbesar_5050special"`
	Discaskecil_5050special           float32 `json:"discaskecil_5050special"`
	Disckopganjil_5050special         float32 `json:"disckopganjil_5050special"`
	Disckopgenap_5050special          float32 `json:"disckopgenap_5050special"`
	Disckopbesar_5050special          float32 `json:"disckopbesar_5050special"`
	Disckopkecil_5050special          float32 `json:"disckopkecil_5050special"`
	Disckepalaganjil_5050special      float32 `json:"disckepalaganjil_5050special"`
	Disckepalagenap_5050special       float32 `json:"disckepalagenap_5050special"`
	Disckepalabesar_5050special       float32 `json:"disckepalabesar_5050special"`
	Disckepalakecil_5050special       float32 `json:"disckepalakecil_5050special"`
	Discekorganjil_5050special        float32 `json:"discekorganjil_5050special"`
	Discekorgenap_5050special         float32 `json:"discekorgenap_5050special"`
	Discekorbesar_5050special         float32 `json:"discekorbesar_5050special"`
	Discekorkecil_5050special         float32 `json:"discekorkecil_5050special"`
	Limitglobal_5050special           float32 `json:"limitglobal_5050special"`
	Limittotal_5050special            float32 `json:"limittotal_5050special"`
	Minbet_5050kombinasi              float32 `json:"minbet_5050kombinasi"`
	Maxbet_5050kombinasi              float32 `json:"maxbet_5050kombinasi"`
	Belakangkeimono_5050kombinasi     float32 `json:"belakangkeimono_5050kombinasi"`
	Belakangkeistereo_5050kombinasi   float32 `json:"belakangkeistereo_5050kombinasi"`
	Belakangkeikembang_5050kombinasi  float32 `json:"belakangkeikembang_5050kombinasi"`
	Belakangkeikempis_5050kombinasi   float32 `json:"belakangkeikempis_5050kombinasi"`
	Belakangkeikembar_5050kombinasi   float32 `json:"belakangkeikembar_5050kombinasi"`
	Tengahkeimono_5050kombinasi       float32 `json:"tengahkeimono_5050kombinasi"`
	Tengahkeistereo_5050kombinasi     float32 `json:"tengahkeistereo_5050kombinasi"`
	Tengahkeikembang_5050kombinasi    float32 `json:"tengahkeikembang_5050kombinasi"`
	Tengahkeikempis_5050kombinasi     float32 `json:"tengahkeikempis_5050kombinasi"`
	Tengahkeikembar_5050kombinasi     float32 `json:"tengahkeikembar_5050kombinasi"`
	Depankeimono_5050kombinasi        float32 `json:"depankeimono_5050kombinasi"`
	Depankeistereo_5050kombinasi      float32 `json:"depankeistereo_5050kombinasi"`
	Depankeikembang_5050kombinasi     float32 `json:"depankeikembang_5050kombinasi"`
	Depankeikempis_5050kombinasi      float32 `json:"depankeikempis_5050kombinasi"`
	Depankeikembar_5050kombinasi      float32 `json:"depankeikembar_5050kombinasi"`
	Belakangdiscmono_5050kombinasi    float32 `json:"belakangdiscmono_5050kombinasi"`
	Belakangdiscstereo_5050kombinasi  float32 `json:"belakangdiscstereo_5050kombinasi"`
	Belakangdisckembang_5050kombinasi float32 `json:"belakangdisckembang_5050kombinasi"`
	Belakangdisckempis_5050kombinasi  float32 `json:"belakangdisckempis_5050kombinasi"`
	Belakangdisckembar_5050kombinasi  float32 `json:"belakangdisckembar_5050kombinasi"`
	Tengahdiscmono_5050kombinasi      float32 `json:"tengahdiscmono_5050kombinasi"`
	Tengahdiscstereo_5050kombinasi    float32 `json:"tengahdiscstereo_5050kombinasi"`
	Tengahdisckembang_5050kombinasi   float32 `json:"tengahdisckembang_5050kombinasi"`
	Tengahdisckempis_5050kombinasi    float32 `json:"tengahdisckempis_5050kombinasi"`
	Tengahdisckembar_5050kombinasi    float32 `json:"tengahdisckembar_5050kombinasi"`
	Depandiscmono_5050kombinasi       float32 `json:"depandiscmono_5050kombinasi"`
	Depandiscstereo_5050kombinasi     float32 `json:"depandiscstereo_5050kombinasi"`
	Depandisckembang_5050kombinasi    float32 `json:"depandisckembang_5050kombinasi"`
	Depandisckempis_5050kombinasi     float32 `json:"depandisckempis_5050kombinasi"`
	Depandisckembar_5050kombinasi     float32 `json:"depandisckembar_5050kombinasi"`
	Limitglobal_5050kombinasi         float32 `json:"limitglobal_5050kombinasi"`
	Limittotal_5050kombinasi          float32 `json:"limittotal_5050kombinasi"`
	Minbet_kombinasi                  float32 `json:"minbet_kombinasi"`
	Maxbet_kombinasi                  float32 `json:"maxbet_kombinasi"`
	Win_kombinasi                     float32 `json:"win_kombinasi"`
	Disc_kombinasi                    float32 `json:"disc_kombinasi"`
	Limitglobal_kombinasi             float32 `json:"limitglobal_kombinasi"`
	Limittotal_kombinasi              float32 `json:"limittotal_kombinasi"`
	Minbet_dasar                      float32 `json:"minbet_dasar"`
	Maxbet_dasar                      float32 `json:"maxbet_dasar"`
	Keibesar_dasar                    float32 `json:"keibesar_dasar"`
	Keikecil_dasar                    float32 `json:"keikecil_dasar"`
	Keigenap_dasar                    float32 `json:"keigenap_dasar"`
	Keiganjil_dasar                   float32 `json:"keiganjil_dasar"`
	Discbesar_dasar                   float32 `json:"discbesar_dasar"`
	Disckecil_dasar                   float32 `json:"disckecil_dasar"`
	Discgenap_dasar                   float32 `json:"discgenap_dasar"`
	Discganjil_dasar                  float32 `json:"discganjil_dasar"`
	Limitglobal_dasar                 float32 `json:"limitglobal_dasar"`
	Limittotal_dasar                  float32 `json:"limittotal_dasar"`
	Minbet_shio                       float32 `json:"minbet_shio"`
	Maxbet_shio                       float32 `json:"maxbet_shio"`
	Win_shio                          float32 `json:"win_shio"`
	Disc_shio                         float32 `json:"disc_shio"`
	Shioyear_shio                     string  `json:"shioyear_shio"`
	Limitglobal_shio                  float32 `json:"limitglobal_shio"`
	Limittotal_shio                   float32 `json:"limittotal_shio"`
	Displaypasaran                    int     `json:"displaypasaran"`
	StatusPasaranActive               string  `json:"statuspasaranactive"`
	Create                            string  `json:"create"`
	Createdate                        string  `json:"createdate"`
	Update                            string  `json:"update"`
	Updatedate                        string  `json:"updatedate"`
}
type responseredis_pasaranonline struct {
	Idpasaranonline int    `json:"idpasaranonline`
	Haripasaran     string `json:"haripasaran"`
}

func PasaranHome(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, client_company, _, _ := helpers.Parsing_Decry(temp_decp, "==")
	field_redis := "LISTPASARAN_AGENT_" + strings.ToLower(client_company)
	var obj entities.Model_pasaranHome
	var arraobj []entities.Model_pasaranHome
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		idcomppasaran, _ := jsonparser.GetInt(value, "idcomppasaran")
		nmpasarantogel, _ := jsonparser.GetString(value, "nmpasarantogel")
		tipepasaran, _ := jsonparser.GetString(value, "tipepasaran")
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
		obj.Tipepasaran = tipepasaran
		obj.PasaranDiundi = pasarandiundi
		obj.Jamtutup = jamtutup
		obj.Jamjadwal = jamjadwal
		obj.Jamopen = jamopen
		obj.Displaypasaran = int(displaypasaran)
		obj.StatusPasaran = statuspasaran
		obj.StatusPasaranActive = statuspasaranactive
		obj.StatusPasarancss = statuspasaran_css
		obj.StatusPasaranActivecss = statuspasaranactive_css
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_home(client_company)
		helpers.SetRedis(field_redis, result, time.Hour*24)
		log.Println("PASARAN MYSQL")
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
		log.Println("PASARAN CACHE")
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
	field_redis := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
	var obj responseredis_pasarandetail
	var arraobj []responseredis_pasarandetail
	var obj2 responseredis_pasaranonline
	var arraobj2 []responseredis_pasaranonline
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	pasaranonline_RD, _, _, _ := jsonparser.Get(jsonredis, "pasaranonline")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		idpasarantogel, _ := jsonparser.GetString(value, "idpasarantogel")
		nmpasarantogel, _ := jsonparser.GetString(value, "nmpasaran")
		pasarandiundi, _ := jsonparser.GetString(value, "pasarandiundi")
		pasaranurl, _ := jsonparser.GetString(value, "pasaranurl")
		jamtutup, _ := jsonparser.GetString(value, "jamtutup")
		jamjadwal, _ := jsonparser.GetString(value, "jamjadwal")
		jamopen, _ := jsonparser.GetString(value, "jamopen")
		limitline4d, _ := jsonparser.GetInt(value, "limitline_4d")
		limitline3d, _ := jsonparser.GetInt(value, "limitline_3d")
		limitline3dd, _ := jsonparser.GetInt(value, "limitline_3dd")
		limitline2d, _ := jsonparser.GetInt(value, "limitline_2d")
		limitline2dd, _ := jsonparser.GetInt(value, "limitline_2dd")
		limitline2dt, _ := jsonparser.GetInt(value, "limitline_2dt")
		bbfs, _ := jsonparser.GetInt(value, "bbfs")
		minbet_432d, _ := jsonparser.GetInt(value, "minbet_432d")
		maxbet4d_432d, _ := jsonparser.GetInt(value, "maxbet4d_432d")
		maxbet3d_432d, _ := jsonparser.GetInt(value, "maxbet3d_432d")
		maxbet3dd_432d, _ := jsonparser.GetInt(value, "maxbet3dd_432d")
		maxbet2d_432d, _ := jsonparser.GetInt(value, "maxbet2d_432d")
		maxbet2dd_432d, _ := jsonparser.GetInt(value, "maxbet2dd_432d")
		maxbet2dt_432d, _ := jsonparser.GetInt(value, "maxbet2dt_432d")
		limitotal4d_432d, _ := jsonparser.GetInt(value, "limitotal4d_432d")
		limitotal3d_432d, _ := jsonparser.GetInt(value, "limitotal3d_432d")
		limitotal3dd_432d, _ := jsonparser.GetInt(value, "limitotal3dd_432d")
		limitotal2d_432d, _ := jsonparser.GetInt(value, "limitotal2d_432d")
		limitotal2dd_432d, _ := jsonparser.GetInt(value, "limitotal2dd_432d")
		limitotal2dt_432d, _ := jsonparser.GetInt(value, "limitotal2dt_432d")
		limitglobal4d_432d, _ := jsonparser.GetInt(value, "limitglobal4d_432d")
		limitglobal3d_432d, _ := jsonparser.GetInt(value, "limitglobal3d_432d")
		limitglobal3dd_432d, _ := jsonparser.GetInt(value, "limitglobal3dd_432d")
		limitglobal2d_432d, _ := jsonparser.GetInt(value, "limitglobal2d_432d")
		limitglobal2dd_432d, _ := jsonparser.GetInt(value, "limitglobal2dd_432d")
		limitglobal2dt_432d, _ := jsonparser.GetInt(value, "limitglobal2dt_432d")
		disc4d_432d, _ := jsonparser.GetFloat(value, "disc4d_432d")
		disc3d_432d, _ := jsonparser.GetFloat(value, "disc3d_432d")
		disc3dd_432d, _ := jsonparser.GetFloat(value, "disc3dd_432d")
		disc2d_432d, _ := jsonparser.GetFloat(value, "disc2d_432d")
		disc2dd_432d, _ := jsonparser.GetFloat(value, "disc2dd_432d")
		disc2dt_432d, _ := jsonparser.GetFloat(value, "disc2dt_432d")
		win4d_432d, _ := jsonparser.GetFloat(value, "win4d_432d")
		win3d_432d, _ := jsonparser.GetFloat(value, "win3d_432d")
		win3dd_432d, _ := jsonparser.GetFloat(value, "win3dd_432d")
		win2d_432d, _ := jsonparser.GetFloat(value, "win2d_432d")
		win2dd_432d, _ := jsonparser.GetFloat(value, "win2dd_432d")
		win2dt_432d, _ := jsonparser.GetFloat(value, "win2dt_432d")
		win4dnodisc_432d, _ := jsonparser.GetFloat(value, "win4dnodisc_432d")
		win3dnodisc_432d, _ := jsonparser.GetFloat(value, "win3dnodisc_432d")
		win3ddnodisc_432d, _ := jsonparser.GetFloat(value, "win3ddnodisc_432d")
		win2dnodisc_432d, _ := jsonparser.GetFloat(value, "win2dnodisc_432d")
		win2ddnodisc_432d, _ := jsonparser.GetFloat(value, "win2ddnodisc_432d")
		win2dtnodisc_432d, _ := jsonparser.GetFloat(value, "win2dtnodisc_432d")
		win4dbb_kena_432d, _ := jsonparser.GetFloat(value, "win4dbb_kena_432d")
		win3dbb_kena_432d, _ := jsonparser.GetFloat(value, "win3dbb_kena_432d")
		win3ddbb_kena_432d, _ := jsonparser.GetFloat(value, "win3ddbb_kena_432d")
		win2dbb_kena_432d, _ := jsonparser.GetFloat(value, "win2dbb_kena_432d")
		win2ddbb_kena_432d, _ := jsonparser.GetFloat(value, "win2ddbb_kena_432d")
		win2dtbb_kena_432d, _ := jsonparser.GetFloat(value, "win2dtbb_kena_432d")
		win4dbb_432d, _ := jsonparser.GetFloat(value, "win4dbb_432d")
		win3dbb_432d, _ := jsonparser.GetFloat(value, "win3dbb_432d")
		win3ddbb_432d, _ := jsonparser.GetFloat(value, "win3ddbb_432d")
		win2dbb_432d, _ := jsonparser.GetFloat(value, "win2dbb_432d")
		win2ddbb_432d, _ := jsonparser.GetFloat(value, "win2ddbb_432d")
		win2dtbb_432d, _ := jsonparser.GetFloat(value, "win2dtbb_432d")
		minbet_cbebas, _ := jsonparser.GetInt(value, "minbet_cbebas")
		maxbet_cbebas, _ := jsonparser.GetInt(value, "maxbet_cbebas")
		win_cbebas, _ := jsonparser.GetFloat(value, "win_cbebas")
		disc_cbebas, _ := jsonparser.GetFloat(value, "disc_cbebas")
		limitglobal_cbebas, _ := jsonparser.GetInt(value, "limitglobal_cbebas")
		limittotal_cbebas, _ := jsonparser.GetInt(value, "limittotal_cbebas")
		minbet_cmacau, _ := jsonparser.GetInt(value, "minbet_cmacau")
		maxbet_cmacau, _ := jsonparser.GetInt(value, "maxbet_cmacau")
		win2d_cmacau, _ := jsonparser.GetFloat(value, "win2d_cmacau")
		win3d_cmacau, _ := jsonparser.GetFloat(value, "win3d_cmacau")
		win4d_cmacau, _ := jsonparser.GetFloat(value, "win4d_cmacau")
		disc_cmacau, _ := jsonparser.GetFloat(value, "disc_cmacau")
		limitglobal_cmacau, _ := jsonparser.GetInt(value, "limitglobal_cmacau")
		limitotal_cmacau, _ := jsonparser.GetInt(value, "limitotal_cmacau")
		minbet_cnaga, _ := jsonparser.GetInt(value, "minbet_cnaga")
		maxbet_cnaga, _ := jsonparser.GetInt(value, "maxbet_cnaga")
		win3_cnaga, _ := jsonparser.GetFloat(value, "win3_cnaga")
		win4_cnaga, _ := jsonparser.GetFloat(value, "win4_cnaga")
		disc_cnaga, _ := jsonparser.GetFloat(value, "disc_cnaga")
		limitglobal_cnaga, _ := jsonparser.GetInt(value, "limitglobal_cnaga")
		limittotal_cnaga, _ := jsonparser.GetInt(value, "limittotal_cnaga")
		minbet_cjitu, _ := jsonparser.GetInt(value, "minbet_cjitu")
		maxbet_cjitu, _ := jsonparser.GetInt(value, "maxbet_cjitu")
		winas_cjitu, _ := jsonparser.GetFloat(value, "winas_cjitu")
		winkop_cjitu, _ := jsonparser.GetFloat(value, "winkop_cjitu")
		winkepala_cjitu, _ := jsonparser.GetFloat(value, "winkepala_cjitu")
		winekor_cjitu, _ := jsonparser.GetFloat(value, "winekor_cjitu")
		desc_cjitu, _ := jsonparser.GetFloat(value, "desc_cjitu")
		limitglobal_cjitu, _ := jsonparser.GetInt(value, "limitglobal_cjitu")
		limittotal_cjitu, _ := jsonparser.GetInt(value, "limittotal_cjitu")
		minbet_5050umum, _ := jsonparser.GetInt(value, "minbet_5050umum")
		maxbet_5050umum, _ := jsonparser.GetInt(value, "maxbet_5050umum")
		keibesar_5050umum, _ := jsonparser.GetFloat(value, "keibesar_5050umum")
		keikecil_5050umum, _ := jsonparser.GetFloat(value, "keikecil_5050umum")
		keigenap_5050umum, _ := jsonparser.GetFloat(value, "keigenap_5050umum")
		keiganjil_5050umum, _ := jsonparser.GetFloat(value, "keiganjil_5050umum")
		keitengah_5050umum, _ := jsonparser.GetFloat(value, "keitengah_5050umum")
		keitepi_5050umum, _ := jsonparser.GetFloat(value, "keitepi_5050umum")
		discbesar_5050umum, _ := jsonparser.GetFloat(value, "discbesar_5050umum")
		disckecil_5050umum, _ := jsonparser.GetFloat(value, "disckecil_5050umum")
		discgenap_5050umum, _ := jsonparser.GetFloat(value, "discgenap_5050umum")
		discganjil_5050umum, _ := jsonparser.GetFloat(value, "discganjil_5050umum")
		disctengah_5050umum, _ := jsonparser.GetFloat(value, "disctengah_5050umum")
		disctepi_5050umum, _ := jsonparser.GetFloat(value, "disctepi_5050umum")
		limitglobal_5050umum, _ := jsonparser.GetInt(value, "limitglobal_5050umum")
		limittotal_5050umum, _ := jsonparser.GetInt(value, "limittotal_5050umum")
		minbet_5050special, _ := jsonparser.GetInt(value, "minbet_5050special")
		maxbet_5050special, _ := jsonparser.GetInt(value, "maxbet_5050special")
		keiasganjil_5050special, _ := jsonparser.GetFloat(value, "keiasganjil_5050special")
		keiasgenap_5050special, _ := jsonparser.GetFloat(value, "keiasgenap_5050special")
		keiasbesar_5050special, _ := jsonparser.GetFloat(value, "keiasbesar_5050special")
		keiaskecil_5050special, _ := jsonparser.GetFloat(value, "keiaskecil_5050special")
		keikopganjil_5050special, _ := jsonparser.GetFloat(value, "keikopganjil_5050special")
		keikopgenap_5050special, _ := jsonparser.GetFloat(value, "keikopgenap_5050special")
		keikopbesar_5050special, _ := jsonparser.GetFloat(value, "keikopbesar_5050special")
		keikopkecil_5050special, _ := jsonparser.GetFloat(value, "keikopkecil_5050special")
		keikepalaganjil_5050special, _ := jsonparser.GetFloat(value, "keikepalaganjil_5050special")
		keikepalagenap_5050special, _ := jsonparser.GetFloat(value, "keikepalagenap_5050special")
		keikepalabesar_5050special, _ := jsonparser.GetFloat(value, "keikepalabesar_5050special")
		keikepalakecil_5050special, _ := jsonparser.GetFloat(value, "keikepalakecil_5050special")
		keiekorganjil_5050special, _ := jsonparser.GetFloat(value, "keiekorganjil_5050special")
		keiekorgenap_5050special, _ := jsonparser.GetFloat(value, "keiekorgenap_5050special")
		keiekorbesar_5050special, _ := jsonparser.GetFloat(value, "keiekorbesar_5050special")
		keiekorkecil_5050special, _ := jsonparser.GetFloat(value, "keiekorkecil_5050special")
		discasganjil_5050special, _ := jsonparser.GetFloat(value, "discasganjil_5050special")
		discasgenap_5050special, _ := jsonparser.GetFloat(value, "discasgenap_5050special")
		discasbesar_5050special, _ := jsonparser.GetFloat(value, "discasbesar_5050special")
		discaskecil_5050special, _ := jsonparser.GetFloat(value, "discaskecil_5050special")
		disckopganjil_5050special, _ := jsonparser.GetFloat(value, "disckopganjil_5050special")
		disckopgenap_5050special, _ := jsonparser.GetFloat(value, "disckopgenap_5050special")
		disckopbesar_5050special, _ := jsonparser.GetFloat(value, "disckopbesar_5050special")
		disckopkecil_5050special, _ := jsonparser.GetFloat(value, "disckopkecil_5050special")
		disckepalaganjil_5050special, _ := jsonparser.GetFloat(value, "disckepalaganjil_5050special")
		disckepalagenap_5050special, _ := jsonparser.GetFloat(value, "disckepalagenap_5050special")
		disckepalabesar_5050special, _ := jsonparser.GetFloat(value, "disckepalabesar_5050special")
		disckepalakecil_5050special, _ := jsonparser.GetFloat(value, "disckepalakecil_5050special")
		discekorganjil_5050special, _ := jsonparser.GetFloat(value, "discekorganjil_5050special")
		discekorgenap_5050special, _ := jsonparser.GetFloat(value, "discekorgenap_5050special")
		discekorbesar_5050special, _ := jsonparser.GetFloat(value, "discekorbesar_5050special")
		discekorkecil_5050special, _ := jsonparser.GetFloat(value, "discekorkecil_5050special")
		limitglobal_5050special, _ := jsonparser.GetInt(value, "limitglobal_5050special")
		limittotal_5050special, _ := jsonparser.GetInt(value, "limittotal_5050special")
		minbet_5050kombinasi, _ := jsonparser.GetInt(value, "minbet_5050kombinasi")
		maxbet_5050kombinasi, _ := jsonparser.GetInt(value, "maxbet_5050kombinasi")
		belakangkeimono_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangkeimono_5050kombinasi")
		belakangkeistereo_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangkeistereo_5050kombinasi")
		belakangkeikembang_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangkeikembang_5050kombinasi")
		belakangkeikempis_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangkeikempis_5050kombinasi")
		belakangkeikembar_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangkeikembar_5050kombinasi")
		tengahkeimono_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahkeimono_5050kombinasi")
		tengahkeistereo_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahkeistereo_5050kombinasi")
		tengahkeikembang_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahkeikembang_5050kombinasi")
		tengahkeikempis_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahkeikempis_5050kombinasi")
		tengahkeikembar_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahkeikembar_5050kombinasi")
		depankeimono_5050kombinasi, _ := jsonparser.GetFloat(value, "depankeimono_5050kombinasi")
		depankeistereo_5050kombinasi, _ := jsonparser.GetFloat(value, "depankeistereo_5050kombinasi")
		depankeikembang_5050kombinasi, _ := jsonparser.GetFloat(value, "depankeikembang_5050kombinasi")
		depankeikempis_5050kombinasi, _ := jsonparser.GetFloat(value, "depankeikempis_5050kombinasi")
		depankeikembar_5050kombinasi, _ := jsonparser.GetFloat(value, "depankeikembar_5050kombinasi")
		belakangdiscmono_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangdiscmono_5050kombinasi")
		belakangdiscstereo_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangdiscstereo_5050kombinasi")
		belakangdisckembang_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangdisckembang_5050kombinasi")
		belakangdisckempis_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangdisckempis_5050kombinasi")
		belakangdisckembar_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangdisckembar_5050kombinasi")
		tengahdiscmono_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahdiscmono_5050kombinasi")
		tengahdiscstereo_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahdiscstereo_5050kombinasi")
		tengahdisckembang_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahdisckembang_5050kombinasi")
		tengahdisckempis_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahdisckempis_5050kombinasi")
		tengahdisckembar_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahdisckembar_5050kombinasi")
		depandiscmono_5050kombinasi, _ := jsonparser.GetFloat(value, "depandiscmono_5050kombinasi")
		depandiscstereo_5050kombinasi, _ := jsonparser.GetFloat(value, "depandiscstereo_5050kombinasi")
		depandisckembang_5050kombinasi, _ := jsonparser.GetFloat(value, "depandisckembang_5050kombinasi")
		depandisckempis_5050kombinasi, _ := jsonparser.GetFloat(value, "depandisckempis_5050kombinasi")
		depandisckembar_5050kombinasi, _ := jsonparser.GetFloat(value, "depandisckembar_5050kombinasi")
		limitglobal_5050kombinasi, _ := jsonparser.GetInt(value, "limitglobal_5050kombinasi")
		limittotal_5050kombinasi, _ := jsonparser.GetInt(value, "limittotal_5050kombinasi")
		minbet_kombinasi, _ := jsonparser.GetInt(value, "minbet_kombinasi")
		maxbet_kombinasi, _ := jsonparser.GetInt(value, "maxbet_kombinasi")
		win_kombinasi, _ := jsonparser.GetFloat(value, "win_kombinasi")
		disc_kombinasi, _ := jsonparser.GetFloat(value, "disc_kombinasi")
		limitglobal_kombinasi, _ := jsonparser.GetInt(value, "limitglobal_kombinasi")
		limittotal_kombinasi, _ := jsonparser.GetInt(value, "limittotal_kombinasi")
		minbet_dasar, _ := jsonparser.GetInt(value, "minbet_dasar")
		maxbet_dasar, _ := jsonparser.GetInt(value, "maxbet_dasar")
		keibesar_dasar, _ := jsonparser.GetFloat(value, "keibesar_dasar")
		keikecil_dasar, _ := jsonparser.GetFloat(value, "keikecil_dasar")
		keigenap_dasar, _ := jsonparser.GetFloat(value, "keigenap_dasar")
		keiganjil_dasar, _ := jsonparser.GetFloat(value, "keiganjil_dasar")
		discbesar_dasar, _ := jsonparser.GetFloat(value, "discbesar_dasar")
		disckecil_dasar, _ := jsonparser.GetFloat(value, "disckecil_dasar")
		discgenap_dasar, _ := jsonparser.GetFloat(value, "discgenap_dasar")
		discganjil_dasar, _ := jsonparser.GetFloat(value, "discganjil_dasar")
		limitglobal_dasar, _ := jsonparser.GetInt(value, "limitglobal_dasar")
		limittotal_dasar, _ := jsonparser.GetInt(value, "limittotal_dasar")
		minbet_shio, _ := jsonparser.GetInt(value, "minbet_shio")
		maxbet_shio, _ := jsonparser.GetInt(value, "maxbet_shio")
		win_shio, _ := jsonparser.GetFloat(value, "win_shio")
		disc_shio, _ := jsonparser.GetFloat(value, "disc_shio")
		shioyear_shio, _ := jsonparser.GetString(value, "shioyear_shio")
		limitglobal_shio, _ := jsonparser.GetInt(value, "limitglobal_shio")
		limittotal_shio, _ := jsonparser.GetInt(value, "limittotal_shio")
		displaypasaran, _ := jsonparser.GetInt(value, "displaypasaran")
		statuspasaranactive, _ := jsonparser.GetString(value, "statuspasaranactive")
		create, _ := jsonparser.GetString(value, "create")
		createdate, _ := jsonparser.GetString(value, "createdate")
		update, _ := jsonparser.GetString(value, "update")
		updatedate, _ := jsonparser.GetString(value, "updatedate")

		obj.Idpasarantogel = idpasarantogel
		obj.Nmpasarantogel = nmpasarantogel
		obj.PasaranDiundi = pasarandiundi
		obj.PasaranURL = pasaranurl
		obj.Jamtutup = jamtutup
		obj.Jamjadwal = jamjadwal
		obj.Jamopen = jamopen
		obj.Limitline4d = int(limitline4d)
		obj.Limitline3d = int(limitline3d)
		obj.Limitline3dd = int(limitline3dd)
		obj.Limitline2d = int(limitline2d)
		obj.Limitline2dd = int(limitline2dd)
		obj.Limitline2dt = int(limitline2dt)
		obj.Bbfs = int(bbfs)
		obj.Minbet_432d = float32(minbet_432d)
		obj.Maxbet4d_432d = float32(maxbet4d_432d)
		obj.Maxbet3d_432d = float32(maxbet3d_432d)
		obj.Maxbet3dd_432d = float32(maxbet3dd_432d)
		obj.Maxbet2d_432d = float32(maxbet2d_432d)
		obj.Maxbet2dd_432d = float32(maxbet2dd_432d)
		obj.Maxbet2dt_432d = float32(maxbet2dt_432d)
		obj.Limitotal4d_432d = float32(limitotal4d_432d)
		obj.Limitotal3d_432d = float32(limitotal3d_432d)
		obj.Limitotal3dd_432d = float32(limitotal3dd_432d)
		obj.Limitotal2d_432d = float32(limitotal2d_432d)
		obj.Limitotal2dd_432d = float32(limitotal2dd_432d)
		obj.Limitotal2dt_432d = float32(limitotal2dt_432d)
		obj.Limitglobal4d_432d = float32(limitglobal4d_432d)
		obj.Limitglobal3d_432d = float32(limitglobal3d_432d)
		obj.Limitglobal3dd_432d = float32(limitglobal3dd_432d)
		obj.Limitglobal2d_432d = float32(limitglobal2d_432d)
		obj.Limitglobal2dd_432d = float32(limitglobal2dd_432d)
		obj.Limitglobal2dt_432d = float32(limitglobal2dt_432d)
		obj.Disc4d_432d = float32(disc4d_432d)
		obj.Disc3d_432d = float32(disc3d_432d)
		obj.Disc3dd_432d = float32(disc3dd_432d)
		obj.Disc2d_432d = float32(disc2d_432d)
		obj.Disc2dd_432d = float32(disc2dd_432d)
		obj.Disc2dt_432d = float32(disc2dt_432d)
		obj.Win4d_432d = float32(win4d_432d)
		obj.Win3d_432d = float32(win3d_432d)
		obj.Win3dd_432d = float32(win3dd_432d)
		obj.Win2d_432d = float32(win2d_432d)
		obj.Win2dd_432d = float32(win2dd_432d)
		obj.Win2dt_432d = float32(win2dt_432d)
		obj.Win4dnodisc_432d = float32(win4dnodisc_432d)
		obj.Win3dnodisc_432d = float32(win3dnodisc_432d)
		obj.Win3ddnodisc_432d = float32(win3ddnodisc_432d)
		obj.Win2dnodisc_432d = float32(win2dnodisc_432d)
		obj.Win2ddnodisc_432d = float32(win2ddnodisc_432d)
		obj.Win2dtnodisc_432d = float32(win2dtnodisc_432d)
		obj.Win4dbb_kena_432d = float32(win4dbb_kena_432d)
		obj.Win3dbb_kena_432d = float32(win3dbb_kena_432d)
		obj.Win3ddbb_kena_432d = float32(win3ddbb_kena_432d)
		obj.Win2dbb_kena_432d = float32(win2dbb_kena_432d)
		obj.Win2ddbb_kena_432d = float32(win2ddbb_kena_432d)
		obj.Win2dtbb_kena_432d = float32(win2dtbb_kena_432d)
		obj.Win4dbb_432d = float32(win4dbb_432d)
		obj.Win3dbb_432d = float32(win3dbb_432d)
		obj.Win3ddbb_432d = float32(win3ddbb_432d)
		obj.Win2dbb_432d = float32(win2dbb_432d)
		obj.Win2ddbb_432d = float32(win2ddbb_432d)
		obj.Win2dtbb_432d = float32(win2dtbb_432d)
		obj.Minbet_cbebas = float32(minbet_cbebas)
		obj.Maxbet_cbebas = float32(maxbet_cbebas)
		obj.Win_cbebas = float32(win_cbebas)
		obj.Disc_cbebas = float32(disc_cbebas)
		obj.Limitglobal_cbebas = float32(limitglobal_cbebas)
		obj.Limittotal_cbebas = float32(limittotal_cbebas)
		obj.Minbet_cmacau = float32(minbet_cmacau)
		obj.Maxbet_cmacau = float32(maxbet_cmacau)
		obj.Win2d_cmacau = float32(win2d_cmacau)
		obj.Win3d_cmacau = float32(win3d_cmacau)
		obj.Win4d_cmacau = float32(win4d_cmacau)
		obj.Disc_cmacau = float32(disc_cmacau)
		obj.Limitglobal_cmacau = float32(limitglobal_cmacau)
		obj.Limitotal_cmacau = float32(limitotal_cmacau)
		obj.Minbet_cnaga = float32(minbet_cnaga)
		obj.Maxbet_cnaga = float32(maxbet_cnaga)
		obj.Win3_cnaga = float32(win3_cnaga)
		obj.Win4_cnaga = float32(win4_cnaga)
		obj.Disc_cnaga = float32(disc_cnaga)
		obj.Limitglobal_cnaga = float32(limitglobal_cnaga)
		obj.Limittotal_cnaga = float32(limittotal_cnaga)
		obj.Minbet_cjitu = float32(minbet_cjitu)
		obj.Maxbet_cjitu = float32(maxbet_cjitu)
		obj.Winas_cjitu = float32(winas_cjitu)
		obj.Winkop_cjitu = float32(winkop_cjitu)
		obj.Winkepala_cjitu = float32(winkepala_cjitu)
		obj.Winekor_cjitu = float32(winekor_cjitu)
		obj.Desc_cjitu = float32(desc_cjitu)
		obj.Limitglobal_cjitu = float32(limitglobal_cjitu)
		obj.Limittotal_cjitu = float32(limittotal_cjitu)
		obj.Minbet_5050umum = float32(minbet_5050umum)
		obj.Maxbet_5050umum = float32(maxbet_5050umum)
		obj.Keibesar_5050umum = float32(keibesar_5050umum)
		obj.Keikecil_5050umum = float32(keikecil_5050umum)
		obj.Keigenap_5050umum = float32(keigenap_5050umum)
		obj.Keiganjil_5050umum = float32(keiganjil_5050umum)
		obj.Keitengah_5050umum = float32(keitengah_5050umum)
		obj.Keitepi_5050umum = float32(keitepi_5050umum)
		obj.Discbesar_5050umum = float32(discbesar_5050umum)
		obj.Disckecil_5050umum = float32(disckecil_5050umum)
		obj.Discgenap_5050umum = float32(discgenap_5050umum)
		obj.Discganjil_5050umum = float32(discganjil_5050umum)
		obj.Disctengah_5050umum = float32(disctengah_5050umum)
		obj.Disctepi_5050umum = float32(disctepi_5050umum)
		obj.Limitglobal_5050umum = float32(limitglobal_5050umum)
		obj.Limittotal_5050umum = float32(limittotal_5050umum)
		obj.Minbet_5050special = float32(minbet_5050special)
		obj.Maxbet_5050special = float32(maxbet_5050special)
		obj.Keiasganjil_5050special = float32(keiasganjil_5050special)
		obj.Keiasgenap_5050special = float32(keiasgenap_5050special)
		obj.Keiasbesar_5050special = float32(keiasbesar_5050special)
		obj.Keiaskecil_5050special = float32(keiaskecil_5050special)
		obj.Keikopganjil_5050special = float32(keikopganjil_5050special)
		obj.Keikopgenap_5050special = float32(keikopgenap_5050special)
		obj.Keikopbesar_5050special = float32(keikopbesar_5050special)
		obj.Keikopkecil_5050special = float32(keikopkecil_5050special)
		obj.Keikepalaganjil_5050special = float32(keikepalaganjil_5050special)
		obj.Keikepalagenap_5050special = float32(keikepalagenap_5050special)
		obj.Keikepalabesar_5050special = float32(keikepalabesar_5050special)
		obj.Keikepalakecil_5050special = float32(keikepalakecil_5050special)
		obj.Keiekorganjil_5050special = float32(keiekorganjil_5050special)
		obj.Keiekorgenap_5050special = float32(keiekorgenap_5050special)
		obj.Keiekorbesar_5050special = float32(keiekorbesar_5050special)
		obj.Keiekorkecil_5050special = float32(keiekorkecil_5050special)
		obj.Discasganjil_5050special = float32(discasganjil_5050special)
		obj.Discasgenap_5050special = float32(discasgenap_5050special)
		obj.Discasbesar_5050special = float32(discasbesar_5050special)
		obj.Discaskecil_5050special = float32(discaskecil_5050special)
		obj.Disckopganjil_5050special = float32(disckopganjil_5050special)
		obj.Disckopgenap_5050special = float32(disckopgenap_5050special)
		obj.Disckopbesar_5050special = float32(disckopbesar_5050special)
		obj.Disckopkecil_5050special = float32(disckopkecil_5050special)
		obj.Disckepalaganjil_5050special = float32(disckepalaganjil_5050special)
		obj.Disckepalagenap_5050special = float32(disckepalagenap_5050special)
		obj.Disckepalabesar_5050special = float32(disckepalabesar_5050special)
		obj.Disckepalakecil_5050special = float32(disckepalakecil_5050special)
		obj.Discekorganjil_5050special = float32(discekorganjil_5050special)
		obj.Discekorgenap_5050special = float32(discekorgenap_5050special)
		obj.Discekorbesar_5050special = float32(discekorbesar_5050special)
		obj.Discekorkecil_5050special = float32(discekorkecil_5050special)
		obj.Limitglobal_5050special = float32(limitglobal_5050special)
		obj.Limittotal_5050special = float32(limittotal_5050special)
		obj.Minbet_5050kombinasi = float32(minbet_5050kombinasi)
		obj.Maxbet_5050kombinasi = float32(maxbet_5050kombinasi)
		obj.Belakangkeimono_5050kombinasi = float32(belakangkeimono_5050kombinasi)
		obj.Belakangkeistereo_5050kombinasi = float32(belakangkeistereo_5050kombinasi)
		obj.Belakangkeikembang_5050kombinasi = float32(belakangkeikembang_5050kombinasi)
		obj.Belakangkeikempis_5050kombinasi = float32(belakangkeikempis_5050kombinasi)
		obj.Belakangkeikembar_5050kombinasi = float32(belakangkeikembar_5050kombinasi)
		obj.Tengahkeimono_5050kombinasi = float32(tengahkeimono_5050kombinasi)
		obj.Tengahkeistereo_5050kombinasi = float32(tengahkeistereo_5050kombinasi)
		obj.Tengahkeikembang_5050kombinasi = float32(tengahkeikembang_5050kombinasi)
		obj.Tengahkeikempis_5050kombinasi = float32(tengahkeikempis_5050kombinasi)
		obj.Tengahkeikembar_5050kombinasi = float32(tengahkeikembar_5050kombinasi)
		obj.Depankeimono_5050kombinasi = float32(depankeimono_5050kombinasi)
		obj.Depankeistereo_5050kombinasi = float32(depankeistereo_5050kombinasi)
		obj.Depankeikembang_5050kombinasi = float32(depankeikembang_5050kombinasi)
		obj.Depankeikempis_5050kombinasi = float32(depankeikempis_5050kombinasi)
		obj.Depankeikembar_5050kombinasi = float32(depankeikembar_5050kombinasi)
		obj.Belakangdiscmono_5050kombinasi = float32(belakangdiscmono_5050kombinasi)
		obj.Belakangdiscstereo_5050kombinasi = float32(belakangdiscstereo_5050kombinasi)
		obj.Belakangdisckembang_5050kombinasi = float32(belakangdisckembang_5050kombinasi)
		obj.Belakangdisckempis_5050kombinasi = float32(belakangdisckempis_5050kombinasi)
		obj.Belakangdisckembar_5050kombinasi = float32(belakangdisckembar_5050kombinasi)
		obj.Tengahdiscmono_5050kombinasi = float32(tengahdiscmono_5050kombinasi)
		obj.Tengahdiscstereo_5050kombinasi = float32(tengahdiscstereo_5050kombinasi)
		obj.Tengahdisckembang_5050kombinasi = float32(tengahdisckembang_5050kombinasi)
		obj.Tengahdisckempis_5050kombinasi = float32(tengahdisckempis_5050kombinasi)
		obj.Tengahdisckembar_5050kombinasi = float32(tengahdisckembar_5050kombinasi)
		obj.Depandiscmono_5050kombinasi = float32(depandiscmono_5050kombinasi)
		obj.Depandiscstereo_5050kombinasi = float32(depandiscstereo_5050kombinasi)
		obj.Depandisckembang_5050kombinasi = float32(depandisckembang_5050kombinasi)
		obj.Depandisckempis_5050kombinasi = float32(depandisckempis_5050kombinasi)
		obj.Depandisckembar_5050kombinasi = float32(depandisckembar_5050kombinasi)
		obj.Limitglobal_5050kombinasi = float32(limitglobal_5050kombinasi)
		obj.Limittotal_5050kombinasi = float32(limittotal_5050kombinasi)
		obj.Minbet_kombinasi = float32(minbet_kombinasi)
		obj.Maxbet_kombinasi = float32(maxbet_kombinasi)
		obj.Win_kombinasi = float32(win_kombinasi)
		obj.Disc_kombinasi = float32(disc_kombinasi)
		obj.Limitglobal_kombinasi = float32(limitglobal_kombinasi)
		obj.Limittotal_kombinasi = float32(limittotal_kombinasi)
		obj.Minbet_dasar = float32(minbet_dasar)
		obj.Maxbet_dasar = float32(maxbet_dasar)
		obj.Keibesar_dasar = float32(keibesar_dasar)
		obj.Keikecil_dasar = float32(keikecil_dasar)
		obj.Keigenap_dasar = float32(keigenap_dasar)
		obj.Keiganjil_dasar = float32(keiganjil_dasar)
		obj.Discbesar_dasar = float32(discbesar_dasar)
		obj.Disckecil_dasar = float32(disckecil_dasar)
		obj.Discgenap_dasar = float32(discgenap_dasar)
		obj.Discganjil_dasar = float32(discganjil_dasar)
		obj.Limitglobal_dasar = float32(limitglobal_dasar)
		obj.Limittotal_dasar = float32(limittotal_dasar)
		obj.Minbet_shio = float32(minbet_shio)
		obj.Maxbet_shio = float32(maxbet_shio)
		obj.Win_shio = float32(win_shio)
		obj.Disc_shio = float32(disc_shio)
		obj.Shioyear_shio = shioyear_shio
		obj.Limitglobal_shio = float32(limitglobal_shio)
		obj.Limittotal_shio = float32(limittotal_shio)
		obj.Displaypasaran = int(displaypasaran)
		obj.StatusPasaranActive = statuspasaranactive
		obj.Create = create
		obj.Createdate = createdate
		obj.Update = update
		obj.Updatedate = updatedate
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(pasaranonline_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		Idpasaranonline, _ := jsonparser.GetInt(value, "Idpasaranonline")
		haripasaran, _ := jsonparser.GetString(value, "haripasaran")

		obj2.Idpasaranonline = int(Idpasaranonline)
		obj2.Haripasaran = haripasaran
		arraobj2 = append(arraobj2, obj2)
	})
	if !flag {
		result, err := models.Fetch_detail(client_company, client.Idpasaran)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(field_redis, result, time.Hour*24)
		log.Println("PASARAN DETAIL MYSQL")
		return c.JSON(result)
	} else {
		log.Println("PASARAN DETAIL CACHE")
		return c.JSON(fiber.Map{
			"status":        fiber.StatusOK,
			"message":       "Success",
			"record":        arraobj,
			"pasaranonline": arraobj2,
			"time":          time.Since(render_page).String(),
		})
	}

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
	field_redis := "LISTPASARAN_AGENT_" + strings.ToLower(client_company)
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
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
		//FRONTEND
		val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
		log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
		//AGEN
		val_agent := helpers.DeleteRedis(field_redis)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			//FRONTEND
			val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
			log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
			//AGEN
			val_agent := helpers.DeleteRedis(field_redis)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
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
	field_redis := "LISTPASARAN_AGENT_" + strings.ToLower(client_company)
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
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
		//FRONTEND
		val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
		log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
		//AGEN
		val_agent := helpers.DeleteRedis(field_redis)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			//FRONTEND
			val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
			log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
			//AGEN
			val_agent := helpers.DeleteRedis(field_redis)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
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
	field_redis := "LISTPASARAN_AGENT_" + strings.ToLower(client_company)
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
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
		//FRONTEND
		val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
		log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
		//AGEN
		val_agent := helpers.DeleteRedis(field_redis)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			//FRONTEND
			val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
			log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
			//AGEN
			val_agent := helpers.DeleteRedis(field_redis)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
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
	field_redis := "LISTPASARAN_AGENT_" + strings.ToLower(client_company)
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranLimitline(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_limitline4d,
			client.Pasaran_limitline3d,
			client.Pasaran_limitline3dd,
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
		//FRONTEND
		val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
		log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
		//AGEN
		val_agent := helpers.DeleteRedis(field_redis)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			result, err := models.Save_PasaranLimitline(
				client_username,
				client_company,
				client.Idpasaran,
				client.Pasaran_limitline4d,
				client.Pasaran_limitline3d,
				client.Pasaran_limitline3dd,
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
			//FRONTEND
			val_frontend_listpasaran := helpers.DeleteRedis("listpasaran_" + strings.ToLower(client_company))
			log.Printf("Redis Delete FRONTEND LISTPASARAN status: %d", val_frontend_listpasaran)
			//AGEN
			val_agent := helpers.DeleteRedis(field_redis)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN status: %d", val_agent)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConf432d(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconf432)
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
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConf432(
			client_username, client_company, client.Idpasaran,
			client.Pasaran_minbet_432d,
			client.Pasaran_maxbet4d_432d, client.Pasaran_maxbet3d_432d, client.Pasaran_maxbet3dd_432d, client.Pasaran_maxbet2d_432d, client.Pasaran_maxbet2dd_432d, client.Pasaran_maxbet2dt_432d,
			client.Pasaran_win4d_432d, client.Pasaran_win3d_432d, client.Pasaran_win3dd_432d, client.Pasaran_win2d_432d, client.Pasaran_win2dd_432d, client.Pasaran_win2dt_432d,
			client.Pasaran_win4dnodisc_432d, client.Pasaran_win3dnodisc_432d, client.Pasaran_win3ddnodisc_432d, client.Pasaran_win2dnodisc_432d, client.Pasaran_win2ddnodisc_432d, client.Pasaran_win2dtnodisc_432d,
			client.Pasaran_win4dbb_kena_432d, client.Pasaran_win3dbb_kena_432d, client.Pasaran_win3ddbb_kena_432d, client.Pasaran_win2dbb_kena_432d, client.Pasaran_win2ddbb_kena_432d, client.Pasaran_win2dtbb_kena_432d,
			client.Pasaran_win4dbb_432d, client.Pasaran_win3dbb_432d, client.Pasaran_win3ddbb_432d, client.Pasaran_win2dbb_432d, client.Pasaran_win2ddbb_432d, client.Pasaran_win2dtbb_432d,
			client.Pasaran_disc4d_432d, client.Pasaran_disc3d_432d, client.Pasaran_disc3dd_432d, client.Pasaran_disc2d_432d, client.Pasaran_disc2dd_432d, client.Pasaran_disc2dt_432d,
			client.Pasaran_limitglobal4d_432d, client.Pasaran_limitglobal3d_432d, client.Pasaran_limitglobal3dd_432d, client.Pasaran_limitglobal2d_432d, client.Pasaran_limitglobal2dd_432d, client.Pasaran_limitglobal2dt_432d,
			client.Pasaran_limitotal4d_432d, client.Pasaran_limitotal3d_432d, client.Pasaran_limitotal3dd_432d, client.Pasaran_limitotal2d_432d, client.Pasaran_limitotal2dd_432d, client.Pasaran_limitotal2dt_432d)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_4-3-2")
		log.Printf("Redis Delete Client - CONF 432 status: %d", val_frontend)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			result, err := models.Save_PasaranConf432(
				client_username, client_company, client.Idpasaran,
				client.Pasaran_minbet_432d,
				client.Pasaran_maxbet4d_432d, client.Pasaran_maxbet3d_432d, client.Pasaran_maxbet3dd_432d, client.Pasaran_maxbet2d_432d, client.Pasaran_maxbet2dd_432d, client.Pasaran_maxbet2dt_432d,
				client.Pasaran_win4d_432d, client.Pasaran_win3d_432d, client.Pasaran_win3dd_432d, client.Pasaran_win2d_432d, client.Pasaran_win2dd_432d, client.Pasaran_win2dt_432d,
				client.Pasaran_win4dnodisc_432d, client.Pasaran_win3dnodisc_432d, client.Pasaran_win3ddnodisc_432d, client.Pasaran_win2dnodisc_432d, client.Pasaran_win2ddnodisc_432d, client.Pasaran_win2dtnodisc_432d,
				client.Pasaran_win4dbb_kena_432d, client.Pasaran_win3dbb_kena_432d, client.Pasaran_win3ddbb_kena_432d, client.Pasaran_win2dbb_kena_432d, client.Pasaran_win2ddbb_kena_432d, client.Pasaran_win2dtbb_kena_432d,
				client.Pasaran_win4dbb_432d, client.Pasaran_win3dbb_432d, client.Pasaran_win3ddbb_432d, client.Pasaran_win2dbb_432d, client.Pasaran_win2ddbb_432d, client.Pasaran_win2dtbb_432d,
				client.Pasaran_disc4d_432d, client.Pasaran_disc3d_432d, client.Pasaran_disc3dd_432d, client.Pasaran_disc2d_432d, client.Pasaran_disc2dd_432d, client.Pasaran_disc2dt_432d,
				client.Pasaran_limitglobal4d_432d, client.Pasaran_limitglobal3d_432d, client.Pasaran_limitglobal3dd_432d, client.Pasaran_limitglobal2d_432d, client.Pasaran_limitglobal2dd_432d, client.Pasaran_limitglobal2dt_432d,
				client.Pasaran_limitotal4d_432d, client.Pasaran_limitotal3d_432d, client.Pasaran_limitotal3dd_432d, client.Pasaran_limitotal2d_432d, client.Pasaran_limitotal2dd_432d, client.Pasaran_limitotal2dt_432d)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"record":  nil,
				})
			}
			val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_4-3-2")
			log.Printf("Redis Delete Client - CONF 432 status: %d", val_frontend)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfColokBebas(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconfcbebas)
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
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
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
		val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_colok")
		log.Printf("Redis Delete Client - CONF COLOK status: %d", val_frontend)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_colok")
			log.Printf("Redis Delete Client - CONF COLOK status: %d", val_frontend)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfColokMacau(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconfcmacau)
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
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
	if typeadmin == "MASTER" {
		result, err := models.Save_PasaranConfColokMacau(
			client_username,
			client_company,
			client.Idpasaran,
			client.Pasaran_minbet_cmacau,
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
		val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_colok")
		log.Printf("Redis Delete Client - CONF COLOK status: %d", val_frontend)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_colok")
			log.Printf("Redis Delete Client - CONF COLOK status: %d", val_frontend)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfColokNaga(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	validate := validator.New()
	client := new(pasaranconfcnaga)
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
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
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
		val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_colok")
		log.Printf("Redis Delete Client - CONF COLOK status: %d", val_frontend)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_colok")
			log.Printf("Redis Delete Client - CONF COLOK status: %d", val_frontend)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfColokJitu(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	validate := validator.New()
	client := new(pasaranconfcjitu)
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
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
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
		val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_colok")
		log.Printf("Redis Delete Client - CONF COLOK status: %d", val_frontend)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_colok")
			log.Printf("Redis Delete Client - CONF COLOK status: %d", val_frontend)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConf5050Umum(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	validate := validator.New()
	client := new(pasaranconfc5050)
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
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
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
		val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_5050")
		log.Printf("Redis Delete Client - CONF 5050 status: %d", val_frontend)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_5050")
			log.Printf("Redis Delete Client - CONF 5050 status: %d", val_frontend)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConf5050Special(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	validate := validator.New()
	client := new(pasaranconfc5050special)
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
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
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
		val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_5050")
		log.Printf("Redis Delete Client - CONF 5050 status: %d", val_frontend)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_5050")
			log.Printf("Redis Delete Client - CONF 5050 status: %d", val_frontend)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConf5050Kombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	validate := validator.New()
	client := new(pasaranconfc5050kombinasi)
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
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
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
		val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_5050")
		log.Printf("Redis Delete Client - CONF 5050 status: %d", val_frontend)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_5050")
			log.Printf("Redis Delete Client - CONF 5050 status: %d", val_frontend)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfMacauKombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	validate := validator.New()
	client := new(pasaranconfmakaukombinasi)
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
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
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

		val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_macaukombinasi")
		log.Printf("Redis Delete Client - CONF MACAUKOMBINASI status: %d", val_frontend)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_macaukombinasi")
			log.Printf("Redis Delete Client - CONF MACAUKOMBINASI status: %d", val_frontend)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfDasar(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	validate := validator.New()
	client := new(pasaranconfdasar)
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
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
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

		val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_dasar")
		log.Printf("Redis Delete Client - CONF DASAR status: %d", val_frontend)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_dasar")
			log.Printf("Redis Delete Client - CONF DASAR status: %d", val_frontend)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
func PasaranSaveConfShio(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	validate := validator.New()
	client := new(pasaranconfshio)
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
	field_redis2 := "LISTPASARAN_AGENT_" + strings.ToLower(client_company) + "_" + strconv.Itoa(client.Idpasaran)
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

		val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_shio")
		log.Printf("Redis Delete Client - CONF SHIO status: %d", val_frontend)
		val_agent2 := helpers.DeleteRedis(field_redis2)
		log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
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
			val_frontend := helpers.DeleteRedis("config_" + strings.ToLower(client_company) + "_" + strings.ToLower(client.Idpasarantogel) + "_shio")
			log.Printf("Redis Delete Client - CONF SHIO status: %d", val_frontend)
			val_agent2 := helpers.DeleteRedis(field_redis2)
			log.Printf("Redis Delete Agent - PASARAN DETAIL status: %d", val_agent2)
			log_redis := "LISTLOG_AGENT_" + strings.ToLower(client_company)
			val_agent_redis := helpers.DeleteRedis(log_redis)
			log.Printf("Redis Delete LOG status: %d", val_agent_redis)
			return c.JSON(result)
		}
	}
}
