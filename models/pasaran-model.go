package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/config"
	"bitbucket.org/isbtotogroup/apibackend_go/db"
	"bitbucket.org/isbtotogroup/apibackend_go/entities"
	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

type pasaranEdit_Online struct {
	Idpasaranonline int    `json:"idpasaranonline`
	Haripasaran     string `json:"haripasaran"`
}

func Fetch_home(company string) (helpers.Response, error) {
	var obj entities.Model_pasaranHome
	var arraobj []entities.Model_pasaranHome
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()

	sql_select := `SELECT 
		A.idcomppasaran, A.pasarandiundi, A.jamtutup, A.jamjadwal, A.jamopen, 
		A.displaypasaran, A.statuspasaran, A.statuspasaranactive, 
		B.nmpasarantogel, B.tipepasaran   
		FROM ` + config.DB_tbl_mst_company_game_pasaran + ` as A 
		JOIN ` + config.DB_tbl_mst_pasaran_togel + ` as B ON B.idpasarantogel = A.idpasarantogel 
		WHERE A.idcompany = ? 
		ORDER BY A.displaypasaran ASC 
	`

	row, err := con.QueryContext(ctx, sql_select, company)
	defer row.Close()

	if err != nil {
		return res, err
	}

	for row.Next() {
		var (
			idcomppasaran, displaypasaran                               int
			pasarandiundi, jamtutup, jamjadwal, jamopen, nmpasarantogel string
			statuspasaran, statuspasaranactive, tipepasaran             string
		)

		err = row.Scan(
			&idcomppasaran,
			&pasarandiundi,
			&jamtutup, &jamjadwal, &jamopen,
			&displaypasaran,
			&statuspasaran, &statuspasaranactive,
			&nmpasarantogel, &tipepasaran)

		if err != nil {
			return res, err
		}
		statuscss := config.STATUS_RUNNING
		statusactivecss := config.STATUS_COMPLETE
		if statuspasaran == "OFFLINE" {
			statuscss = config.STATUS_CANCEL
		}
		if statuspasaranactive == "Y" {
			statuspasaranactive = "ACTIVE"
		} else {
			statuspasaranactive = "DEACTIVE"
			statusactivecss = config.STATUS_CANCEL
		}

		obj.Idcomppasaran = idcomppasaran
		obj.Nmpasarantogel = nmpasarantogel
		obj.Tipepasaran = tipepasaran
		obj.PasaranDiundi = pasarandiundi
		obj.Jamtutup = jamtutup
		obj.Jamjadwal = jamjadwal
		obj.Jamopen = jamopen
		obj.Displaypasaran = displaypasaran
		obj.StatusPasaran = statuspasaran
		obj.StatusPasaranActive = statuspasaranactive
		obj.StatusPasarancss = statuscss
		obj.StatusPasaranActivecss = statusactivecss
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()

	return res, nil
}

func Fetch_detail(company string, idcomppasaran int) (helpers.ResponsePasaran, error) {
	var obj entities.Model_pasaranEdit
	var arraobj []entities.Model_pasaranEdit
	var res helpers.ResponsePasaran
	msg := "Success"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()

	sql_select := `SELECT 
		A.idpasarantogel, A.pasarandiundi, A.pasaranurl, A.jamtutup, A.jamjadwal, A.jamopen, 
		A.statuspasaranactive, A.displaypasaran, B.nmpasarantogel, 
		A.limitline_4d, A.limitline_3d, A.limitline_3dd, A.limitline_2d, A.limitline_2dd, A.limitline_2dt, A.bbfs, 
		A.1_minbet as minbet_432d, A.1_maxbet4d as maxbet4d_432d, A.1_maxbet3d as maxbet3d_432d, A.1_maxbet3dd as maxbet3dd_432d, 
		A.1_maxbet2d as maxbet2d_432d, A.1_maxbet2dd as maxbet2dd_432d, A.1_maxbet2dt as maxbet2dt_432d, 
		A.1_maxbet4d_fullbb as maxbet4d_fullbb_432d, A.1_maxbet3d_fullbb as maxbet3d_fullbb_432d, A.1_maxbet3dd_fullbb as maxbet3dd_fullbb_432d,
		A.1_maxbet2d_fullbb as maxbet2d_fullbb_432d, A.1_maxbet2dd_fullbb as maxbet2dd_fullbb_432d, A.1_maxbet2dt_fullbb as maxbet2dt_fullbb_432d,
		A.1_maxbuy4d as maxbuy4d_432d, A.1_maxbuy3d as maxbuy3d_432d, A.1_maxbuy3dd as maxbuy3dd_432d,
		A.1_maxbuy2d as maxbuy2d_432d, A.1_maxbuy2dd as maxbuy2dd_432d, A.1_maxbuy2dt as maxbuy2dt_432d,
		A.1_limittotal4d as limitotal4d_432d, A.1_limittotal3d as limitotal3d_432d, A.1_limittotal3dd as limitotal3dd_432d, 
		A.1_limittotal2d as limitotal2d_432d, A.1_limittotal2dd as limitotal2dd_432d, A.1_limittotal2dt as limitotal2dt_432d, 
		A.1_limitbuang4d as limitglobal4d_432d, A.1_limitbuang3d as limitglobal3d_432d, A.1_limitbuang3dd as limitglobal3dd_432d, 
		A.1_limitbuang2d as limitglobal2d_432d, A.1_limitbuang2dd as limitglobal2dd_432d, A.1_limitbuang2dt as limitglobal2dt_432d, 
		A.1_limitbuang4d_fullbb as limitbuang4d_fullbb_432d, A.1_limitbuang3d_fullbb as limitbuang3d_fullbb_432d, A.1_limitbuang3dd_fullbb as limitbuang3dd_fullbb_432d, 
		A.1_limitbuang2d_fullbb as limitbuang2d_fullbb_432d, A.1_limitbuang2dd_fullbb as limitbuang2dd_fullbb_432d, A.1_limitbuang2dt_fullbb as limitbuang2dt_fullbb_432d, 
		A.1_limittotal4d_fullbb as limittotal4d_fullbb_432d, A.1_limittotal3d_fullbb as limittotal3d_fullbb_432d, A.1_limittotal3dd_fullbb as limittotal3dd_fullbb_432d, 
		A.1_limittotal2d_fullbb as limittotal2d_fullbb_432d, A.1_limittotal2dd_fullbb as limittotal2dd_fullbb_432d, A.1_limittotal2dt_fullbb as limittotal2dt_fullbb_432d, 
		A.1_disc4d as disc4d_432d, A.1_disc3d as disc3d_432d, A.1_disc3dd as disc3dd_432d, 
		A.1_disc2d as disc2d_432d, A.1_disc2dd as disc2dd_432d, A.1_disc2dt as disc2dt_432d, 
		A.1_win4d as win4d_432d, A.1_win3d as win3d_432d, A.1_win3dd as win3dd_432d, 
		A.1_win2d as win2d_432d, A.1_win2dd as win2dd_432d, A.1_win2dt as win2dt_432d, 
		A.1_win4dnodisc as win4dnodisc_432d, A.1_win3dnodisc as win3dnodisc_432d, A.1_win3ddnodisc as win3ddnodisc_432d, 
		A.1_win2dnodisc as win2dnodisc_432d, A.1_win2ddnodisc as win2ddnodisc_432d, A.1_win2dtnodisc as win2dtnodisc_432d,
		A.1_win4dbb_kena as win4dbb_kena_432d, A.1_win3dbb_kena as win3dbb_kena_432d, A.1_win3ddbb_kena as win3dbb_kena_432d, 
		A.1_win2dbb_kena as win2dbb_kena_432d, A.1_win2ddbb_kena as win2ddbb_kena_432d, A.1_win2dtbb_kena as win2dtbb_kena_432d,
		A.1_win4dbb as win4dbb_432d, A.1_win3dbb as win3dbb_432d, A.1_win3ddbb as win3dbb_432d, 
		A.1_win2dbb as win2dbb_432d, A.1_win2ddbb as win2ddbb_432d, A.1_win2dtbb as win2dtbb_432d,
		A.2_minbet as minbet_cbebas, A.2_maxbet as maxbet_cbebas, A.2_maxbuy as maxbuy_cbebas,
		A.2_win as win_cbebas, A.2_disc as disc_cbebas, 
		A.2_limitbuang as limitglobal_cbebas, A.2_limitotal as limittotal_cbebas, 
		A.3_minbet as minbet_cmacau, A.3_maxbet as maxbet_cmacau, A.3_maxbuy as maxbuy_cmacau,
		A.3_win2digit as win2d_cmacau, A.3_win3digit as win3d_cmacau, A.3_win4digit as win4d_cmacau, 
		A.3_disc as disc_cmacau, A.3_limitbuang as limitglobal_cmacau, A.3_limittotal as limitotal_cmacau, 
		A.4_minbet as minbet_cnaga, A.4_maxbet as maxbet_cnaga, A.4_maxbuy as maxbuy_cnaga,
		A.4_win3digit as win3_cnaga, A.4_win4digit as win4_cnaga, 
		A.4_disc as disc_cnaga, A.4_limitbuang as limitglobal_cnaga, A.4_limittotal as limittotal_cnaga, 
		A.5_minbet as minbet_cjitu, A.5_maxbet as maxbet_cjitu, A.5_maxbuy as maxbuy_cjitu, 
		A.5_winas as winas_cjitu, A.5_winkop as winkop_cjitu, A.5_winkepala as winkepala_cjitu, A.5_winekor as winekor_cjitu, 
		A.5_desic as desc_cjitu, A.5_limitbuang as limitglobal_cjitu, A.5_limitotal as limittotal_cjitu, 
		A.6_minbet as minbet_5050umum, A.6_maxbet as maxbet_5050umum, A.6_maxbuy as maxbuy_5050umum,
		A.6_keibesar as keibesar_5050umum, A.6_keikecil as keikecil_5050umum, A.6_keigenap as keigenap_5050umum, 
		A.6_keiganjil as keiganjil_5050umum, A.6_keitengah as keitengah_5050umum, A.6_keitepi as keitepi_5050umum, 
		A.6_discbesar as discbesar_5050umum, A.6_disckecil as disckecil_5050umum, A.6_discgenap as discgenap_5050umum, 
		A.6_discganjil as discganjil_5050umum, A.6_disctengah as disctengah_5050umum, A.6_disctepi as disctepi_5050umum, 
		A.6_limitbuang as limitglobal_5050umum, A.6_limittotal as limittotal_5050umum, 
		A.7_minbet as minbet_5050special, A.7_maxbet as maxbet_5050special, A.7_maxbuy as maxbuy_5050special,
		A.7_keiasganjil as keiasganjil_5050special, A.7_keiasgenap as keiasgenap_5050special, A.7_keiasbesar as keiasbesar_5050special, 
		A.7_keiaskecil as keiaskecil_5050special, A.7_keikopganjil as keikopganjil_5050special, A.7_keikopgenap as keikopgenap_5050special, 
		A.7_keikopbesar as keikopbesar_5050special, A.7_keikopkecil as keikopkecil_5050special, A.7_keikepalaganjil as keikepalaganjil_5050special, 
		A.7_keikepalagenap as keikepalagenap_5050special, A.7_keikepalabesar as keikepalabesar_5050special, A.7_keikepalakecil as keikepalakecil_5050special, 
		A.7_keiekorganjil as keiekorganjil_5050special, A.7_keiekorgenap as keiekorgenap_5050special, A.7_keiekorbesar as keiekorbesar_5050special, 
		A.7_keiekorkecil as keiekorkecil_5050special, 
		A.7_discasganjil as discasganjil_5050special, A.7_discasgenap as discasgenap_5050special, A.7_discasbesar as discasbesar_5050special, 
		A.7_discaskecil as discaskecil_5050special, A.7_disckopganjil as disckopganjil_5050special, A.7_disckopgenap as disckopgenap_5050special, 
		A.7_disckopbesar as disckopbesar_5050special, A.7_disckopkecil as disckopkecil_5050special, A.7_disckepalaganjil as disckepalaganjil_5050special, 
		A.7_disckepalagenap as disckepalagenap_5050special, A.7_disckepalabesar as disckepalabesar_5050special, A.7_disckepalakecil as disckepalakecil_5050special, 
		A.7_discekorganjil as discekorganjil_5050special, A.7_discekorgenap as discekorgenap_5050special, A.7_discekorbesar as discekorbesar_5050special, 
		A.7_discekorkecil as discekorkecil_5050special, A.7_limitbuang as limitglobal_5050special, A.7_limittotal as limittotal_5050special, 
		A.8_minbet as minbet_5050kombinasi, A.8_maxbet as maxbet_5050kombinasi, A.8_maxbuy as maxbuy_5050kombinasi,
		A.8_belakangkeimono as belakangkeimono_5050kombinasi, A.8_belakangkeistereo as belakangkeistereo_5050kombinasi, A.8_belakangkeikembang as belakangkeikembang_5050kombinasi, A.8_belakangkeikempis as belakangkeikempis_5050kombinasi, A.8_belakangkeikembar as belakangkeikembar_5050kombinasi, 
		A.8_tengahkeimono as tengahkeimono_5050kombinasi, A.8_tengahkeistereo as tengahkeistereo_5050kombinasi, A.8_tengahkeikembang as tengahkeikembang_5050kombinasi, A.8_tengahkeikempis as tengahkeikempis_5050kombinasi, A.8_tengahkeikembar as tengahkeikembar_5050kombinasi, 
		A.8_depankeimono as depankeimono_5050kombinasi, A.8_depankeistereo as depankeistereo_5050kombinasi, A.8_depankeikembang as depankeikembang_5050kombinasi, A.8_depankeikempis as depankeikempis_5050kombinasi, A.8_depankeikembar as depankeikembar_5050kombinasi, 
		A.8_belakangdiscmono as belakangdiscmono_5050kombinasi, A.8_belakangdiscstereo as belakangdiscstereo_5050kombinasi, A.8_belakangdisckembang as belakangdisckembang_5050kombinasi, A.8_belakangdisckempis as belakangdisckempis_5050kombinasi, A.8_belakangdisckembar as belakangdisckembar_5050kombinasi, 
		A.8_tengahdiscmono as tengahdiscmono_5050kombinasi, A.8_tengahdiscstereo as tengahdiscstereo_5050kombinasi, A.8_tengahdisckembang as tengahdisckembang_5050kombinasi, A.8_tengahdisckempis as tengahdisckempis_5050kombinasi, A.8_tengahdisckembar as tengahdisckembar_5050kombinasi, 
		A.8_depandiscmono as depandiscmono_5050kombinasi, A.8_depandiscstereo as depandiscstereo_5050kombinasi, A.8_depandisckembang as depandisckembang_5050kombinasi, A.8_depandisckempis as depandisckempis_5050kombinasi, A.8_depandisckembar as depandisckembar_5050kombinasi, 
		A.8_limitbuang as limitglobal_5050kombinasi, A.8_limittotal as limittotal_5050kombinasi, 
		A.9_minbet as minbet_kombinasi, A.9_maxbet as maxbet_kombinasi, A.9_maxbuy as maxbuy_kombinasi,
		A.9_win as win_kombinasi, A.9_discount as disc_kombinasi, A.9_limitbuang as limitglobal_kombinasi, A.9_limittotal as limittotal_kombinasi, 
		A.10_minbet as minbet_dasar, A.10_maxbet as maxbet_dasar, A.10_maxbuy as maxbuy_dasar,
		A.10_keibesar as keibesar_dasar, A.10_keikecil as keikecil_dasar, A.10_keigenap as keigenap_dasar, A.10_keiganjil as keiganjil_dasar, 
		A.10_discbesar as discbesar_dasar, A.10_disckecil as disckecil_dasar, A.10_discigenap as discgenap_dasar, A.10_discganjil as discganjil_dasar, 
		A.10_limitbuang as limitglobal_dasar, A.10_limittotal as limittotal_dasar, 
		A.11_minbet as minbet_shio, A.11_maxbet as maxbet_shio, A.11_maxbuy as maxbuy_shio,
		A.11_win as win_shio, A.11_disc as disc_shio, A.11_limitbuang as limitglobal_shio, A.11_limittotal as limittotal_shio, 
		A.11_shiotahunini as shioyear_shio, 
		A.createcomppas, A.createdatecomppas, 
		A.updatecomppas, A.updatedatecompas 
		FROM ` + config.DB_tbl_mst_company_game_pasaran + ` as A 
		JOIN ` + config.DB_tbl_mst_pasaran_togel + ` as B ON B.idpasarantogel = A.idpasarantogel
		WHERE A.idcompany = ? 
		AND A.idcomppasaran = ? 
	`
	var (
		idpasarantogel_db, pasarandiundi_db, pasaranurl_db, jamtutup_db, jamjadwal_db, jamopen_db, statuspasaranactive_db, nmpasarantogel_db                                                                                                                                                                     string
		createcomppas_db, createdatecomppas_db, updatecomppas_db, updatedatecompas_db                                                                                                                                                                                                                            string
		displaypasaran_db, limitline_4d_db, limitline_3d_db, limitline_3dd_db, limitline_2d_db, limitline_2dd_db, limitline_2dt_db, bbfs_db                                                                                                                                                                      int
		minbet_432d_db, maxbet4d_432d_db, maxbet3d_432d_db, maxbet3dd_432d_db, maxbet2d_432d_db, maxbet2dd_432d_db, maxbet2dt_432d_db                                                                                                                                                                            float32
		maxbet4d_fullbb_432d_db, maxbet3d_fullbb_432d_db, maxbet3dd_fullbb_432d_db, maxbet2d_fullbb_432d_db, maxbet2dd_fullbb_432d_db, maxbet2dt_fullbb_432d_db                                                                                                                                                  float32
		maxbuy4d_432d_db, maxbuy3d_432d_db, maxbuy3dd_432d_db, maxbuy2d_432d_db, maxbuy2dd_432d_db, maxbuy2dt_432d_db                                                                                                                                                                                            float32
		limitotal4d_432d_db, limitotal3d_432d_db, limitotal3dd_432d_db, limitotal2d_432d_db, limitotal2dd_432d_db, limitotal2dt_432d_db                                                                                                                                                                          float32
		limitglobal4d_432d_db, limitglobal3d_432d_db, limitglobal3dd_432d_db, limitglobal2d_432d_db, limitglobal2dd_432d_db, limitglobal2dt_432d_db                                                                                                                                                              float32
		limitbuang4d_fullbb_432d_db, limitbuang3d_fullbb_432d_db, limitbuang3dd_fullbb_432d_db, limitbuang2d_fullbb_432d_db, limitbuang2dd_fullbb_432d_db, limitbuang2dt_fullbb_432d_db                                                                                                                          float32
		limittotal4d_fullbb_432d_db, limittotal3d_fullbb_432d_db, limittotal3dd_fullbb_432d_db, limittotal2d_fullbb_432d_db, limittotal2dd_fullbb_432d_db, limittotal2dt_fullbb_432d_db                                                                                                                          float32
		disc4d_432d_db, disc3d_432d_db, disc3dd_432d_db, disc2d_432d_db, disc2dd_432d_db, disc2dt_432d_db                                                                                                                                                                                                        float32
		win4d_432d_db, win3d_432d_db, win3dd_432d_db, win2d_432d_db, win2dd_432d_db, win2dt_432d_db                                                                                                                                                                                                              float32
		win4dnodisc_432d_db, win3dnodisc_432d_db, win3ddnodisc_432d_db, win2dnodisc_432d_db, win2ddnodisc_432d_db, win2dtnodisc_432d_db                                                                                                                                                                          float32
		win4dbb_kena_432d_db, win3dbb_kena_432d_db, win3ddbb_kena_432d_db, win2dbb_kena_432d_db, win2ddbb_kena_432d_db, win2dtbb_kena_432d_db                                                                                                                                                                    float32
		win4dbb_432d_db, win3dbb_432d_db, win3ddbb_432d_db, win2dbb_432d_db, win2ddbb_432d_db, win2dtbb_432d_db                                                                                                                                                                                                  float32
		minbet_cbebas_db, maxbet_cbebas_db, maxbuy_cbebas_db, win_cbebas_db, disc_cbebas_db, limitglobal_cbebas_db, limittotal_cbebas_db                                                                                                                                                                         float32
		minbet_cmacau_db, maxbet_cmacau_db, maxbuy_cmacau_db, win2d_cmacau_db, win3d_cmacau_db, win4d_cmacau_db, disc_cmacau_db, limitglobal_cmacau_db, limitotal_cmacau_db                                                                                                                                      float32
		minbet_cnaga_db, maxbet_cnaga_db, maxbuy_cnaga_db, win3_cnaga_db, win4_cnaga_db, disc_cnaga_db, limitglobal_cnaga_db, limittotal_cnaga_db                                                                                                                                                                float32
		minbet_cjitu_db, maxbet_cjitu_db, maxbuy_cjitu_db, winas_cjitu_db, winkop_cjitu_db, winkepala_cjitu_db, winekor_cjitu_db, desc_cjitu_db, limitglobal_cjitu_db, limittotal_cjitu_db                                                                                                                       float32
		minbet_5050umum_db, maxbet_5050umum_db, maxbuy_5050umum_db, keibesar_5050umum_db, keikecil_5050umum_db, keigenap_5050umum_db, keiganjil_5050umum_db, keitengah_5050umum_db, keitepi_5050umum_db                                                                                                          float32
		discbesar_5050umum_db, disckecil_5050umum_db, discgenap_5050umum_db, discganjil_5050umum_db, disctengah_5050umum_db, disctepi_5050umum_db, limitglobal_5050umum_db, limittotal_5050umum_db                                                                                                               float32
		minbet_5050special_db, maxbet_5050special_db, maxbuy_5050special_db, keiasganjil_5050special_db, keiasgenap_5050special_db, keiasbesar_5050special_db, keiaskecil_5050special_db, keikopganjil_5050special_db, keikopgenap_5050special_db                                                                float32
		keikopbesar_5050special_db, keikopkecil_5050special_db, keikepalaganjil_5050special_db, keikepalagenap_5050special_db, keikepalabesar_5050special_db, keikepalakecil_5050special_db, keiekorganjil_5050special_db, keiekorgenap_5050special_db, keiekorbesar_5050special_db, keiekorkecil_5050special_db float32
		discasganjil_5050special_db, discasgenap_5050special_db, discasbesar_5050special_db, discaskecil_5050special_db, disckopganjil_5050special_db, disckopgenap_5050special_db, disckopbesar_5050special_db, disckopkecil_5050special_db, disckepalaganjil_5050special_db, disckepalagenap_5050special_db    float32
		disckepalabesar_5050special_db, disckepalakecil_5050special_db, discekorganjil_5050special_db, discekorgenap_5050special_db, discekorbesar_5050special_db, discekorkecil_5050special_db                                                                                                                  float32
		limitglobal_5050special_db, limittotal_5050special_db                                                                                                                                                                                                                                                    float32
		minbet_5050kombinasi_db, maxbet_5050kombinasi_db, maxbuy_5050kombinasi_db                                                                                                                                                                                                                                float32
		belakangkeimono_5050kombinasi_db, belakangkeistereo_5050kombinasi_db, belakangkeikembang_5050kombinasi_db, belakangkeikempis_5050kombinasi_db, belakangkeikembar_5050kombinasi_db                                                                                                                        float32
		tengahkeimono_5050kombinasi_db, tengahkeistereo_5050kombinasi_db, tengahkeikembang_5050kombinasi_db, tengahkeikempis_5050kombinasi_db, tengahkeikembar_5050kombinasi_db                                                                                                                                  float32
		depankeimono_5050kombinasi_db, depankeistereo_5050kombinasi_db, depankeikembang_5050kombinasi_db, depankeikempis_5050kombinasi_db, depankeikembar_5050kombinasi_db                                                                                                                                       float32
		belakangdiscmono_5050kombinasi_db, belakangdiscstereo_5050kombinasi_db, belakangdisckembang_5050kombinasi_db, belakangdisckempis_5050kombinasi_db, belakangdisckembar_5050kombinasi_db                                                                                                                   float32
		tengahdiscmono_5050kombinasi_db, tengahdiscstereo_5050kombinasi_db, tengahdisckembang_5050kombinasi_db, tengahdisckempis_5050kombinasi_db, tengahdisckembar_5050kombinasi_db                                                                                                                             float32
		depandiscmono_5050kombinasi_db, depandiscstereo_5050kombinasi_db, depandisckembang_5050kombinasi_db, depandisckempis_5050kombinasi_db, depandisckembar_5050kombinasi_db                                                                                                                                  float32
		limitglobal_5050kombinasi_db, limittotal_5050kombinasi_db                                                                                                                                                                                                                                                float32
		minbet_kombinasi_db, maxbet_kombinasi_db, maxbuy_kombinasi_db, win_kombinasi_db, disc_kombinasi_db, limitglobal_kombinasi_db, limittotal_kombinasi_db                                                                                                                                                    float32
		minbet_dasar_db, maxbet_dasar_db, maxbuy_dasar_db, keibesar_dasar_db, keikecil_dasar_db, keigenap_dasar_db, keiganjil_dasar_db, discbesar_dasar_db, disckecil_dasar_db, discgenap_dasar_db, discganjil_dasar_db, limitglobal_dasar_db, limittotal_dasar_db                                               float32
		minbet_shio_db, maxbet_shio_db, maxbuy_shio_db, win_shio_db, disc_shio_db, limitglobal_shio_db, limittotal_shio_db                                                                                                                                                                                       float32
		shioyear_shio_db                                                                                                                                                                                                                                                                                         string
	)
	err := con.QueryRowContext(ctx, sql_select, company, idcomppasaran).Scan(
		&idpasarantogel_db, &pasarandiundi_db, &pasaranurl_db, &jamtutup_db,
		&jamjadwal_db, &jamopen_db, &statuspasaranactive_db, &displaypasaran_db, &nmpasarantogel_db,
		&limitline_4d_db, &limitline_3d_db, &limitline_3dd_db, &limitline_2d_db, &limitline_2dd_db, &limitline_2dt_db, &bbfs_db,
		&minbet_432d_db, &maxbet4d_432d_db, &maxbet3d_432d_db, &maxbet3dd_432d_db, &maxbet2d_432d_db, &maxbet2dd_432d_db, &maxbet2dt_432d_db,
		&maxbet4d_fullbb_432d_db, &maxbet3d_fullbb_432d_db, &maxbet3dd_fullbb_432d_db, &maxbet2d_fullbb_432d_db, &maxbet2dd_fullbb_432d_db, &maxbet2dt_fullbb_432d_db,
		&maxbuy4d_432d_db, &maxbuy3d_432d_db, &maxbuy3dd_432d_db, &maxbuy2d_432d_db, &maxbuy2dd_432d_db, &maxbuy2dt_432d_db, &maxbet2dt_fullbb_432d_db,
		&limitotal4d_432d_db, &limitotal3d_432d_db, &limitotal3dd_432d_db, &limitotal2d_432d_db, &limitotal2dd_432d_db, &limitotal2dt_432d_db,
		&limitglobal4d_432d_db, &limitglobal3d_432d_db, &limitglobal3dd_432d_db, &limitglobal2d_432d_db, &limitglobal2dd_432d_db, &limitglobal2dt_432d_db,
		&limitbuang4d_fullbb_432d_db, &limitbuang3d_fullbb_432d_db, &limitbuang3dd_fullbb_432d_db, &limitbuang2d_fullbb_432d_db, &limitbuang2dd_fullbb_432d_db, &limitbuang2dt_fullbb_432d_db,
		&limittotal4d_fullbb_432d_db, &limittotal3d_fullbb_432d_db, &limittotal3dd_fullbb_432d_db, &limittotal2d_fullbb_432d_db, &limittotal2dd_fullbb_432d_db, &limittotal2dt_fullbb_432d_db,
		&disc4d_432d_db, &disc3d_432d_db, &disc3dd_432d_db, &disc2d_432d_db, &disc2dd_432d_db, &disc2dt_432d_db,
		&win4d_432d_db, &win3d_432d_db, &win3dd_432d_db, &win2d_432d_db, &win2dd_432d_db, &win2dt_432d_db,
		&win4dnodisc_432d_db, &win3dnodisc_432d_db, &win3ddnodisc_432d_db, &win2dnodisc_432d_db, &win2ddnodisc_432d_db, &win2dtnodisc_432d_db,
		&win4dbb_kena_432d_db, &win3dbb_kena_432d_db, &win3ddbb_kena_432d_db, &win2dbb_kena_432d_db, &win2ddbb_kena_432d_db, &win2dtbb_kena_432d_db,
		&win4dbb_432d_db, &win3dbb_432d_db, &win3ddbb_432d_db, &win2dbb_432d_db, &win2ddbb_432d_db, &win2dtbb_432d_db,
		&minbet_cbebas_db, &maxbet_cbebas_db, &maxbuy_cbebas_db, &win_cbebas_db, &disc_cbebas_db, &limitglobal_cbebas_db, &limittotal_cbebas_db,
		&minbet_cmacau_db, &maxbet_cmacau_db, &maxbuy_cmacau_db, &win2d_cmacau_db, &win3d_cmacau_db, &win4d_cmacau_db, &disc_cmacau_db, &limitglobal_cmacau_db, &limitotal_cmacau_db,
		&minbet_cnaga_db, &maxbet_cnaga_db, &maxbuy_cbebas_db, &win3_cnaga_db, &win4_cnaga_db, &disc_cnaga_db, &limitglobal_cnaga_db, &limittotal_cnaga_db,
		&minbet_cjitu_db, &maxbet_cjitu_db, &maxbuy_cjitu_db,
		&winas_cjitu_db, &winkop_cjitu_db, &winkepala_cjitu_db, &winekor_cjitu_db,
		&desc_cjitu_db, &limitglobal_cjitu_db, &limittotal_cjitu_db,
		&minbet_5050umum_db, &maxbet_5050umum_db, &maxbuy_5050umum_db,
		&keibesar_5050umum_db, &keikecil_5050umum_db, &keigenap_5050umum_db, &keiganjil_5050umum_db, &keitengah_5050umum_db, &keitepi_5050umum_db,
		&discbesar_5050umum_db, &disckecil_5050umum_db, &discgenap_5050umum_db, &discganjil_5050umum_db, &disctengah_5050umum_db, &disctepi_5050umum_db, &limitglobal_5050umum_db, &limittotal_5050umum_db,
		&minbet_5050special_db, &maxbet_5050special_db, &maxbuy_5050special_db, &keiasganjil_5050special_db, &keiasgenap_5050special_db, &keiasbesar_5050special_db, &keiaskecil_5050special_db, &keikopganjil_5050special_db, &keikopgenap_5050special_db,
		&keikopbesar_5050special_db, &keikopkecil_5050special_db, &keikepalaganjil_5050special_db, &keikepalagenap_5050special_db, &keikepalabesar_5050special_db, &keikepalakecil_5050special_db, &keiekorganjil_5050special_db, &keiekorgenap_5050special_db, &keiekorbesar_5050special_db, &keiekorkecil_5050special_db,
		&discasganjil_5050special_db, &discasgenap_5050special_db, &discasbesar_5050special_db, &discaskecil_5050special_db, &disckopganjil_5050special_db, &disckopgenap_5050special_db, &disckopbesar_5050special_db, &disckopkecil_5050special_db, &disckepalaganjil_5050special_db, &disckepalagenap_5050special_db,
		&disckepalabesar_5050special_db, &disckepalakecil_5050special_db, &discekorganjil_5050special_db, &discekorgenap_5050special_db, &discekorbesar_5050special_db, &discekorkecil_5050special_db, &limitglobal_5050special_db, &limittotal_5050special_db,
		&minbet_5050kombinasi_db, &maxbet_5050kombinasi_db, &maxbuy_5050kombinasi_db,
		&belakangkeimono_5050kombinasi_db, &belakangkeistereo_5050kombinasi_db, &belakangkeikembang_5050kombinasi_db, &belakangkeikempis_5050kombinasi_db, &belakangkeikembar_5050kombinasi_db,
		&tengahkeimono_5050kombinasi_db, &tengahkeistereo_5050kombinasi_db, &tengahkeikembang_5050kombinasi_db, &tengahkeikempis_5050kombinasi_db, &tengahkeikembar_5050kombinasi_db,
		&depankeimono_5050kombinasi_db, &depankeistereo_5050kombinasi_db, &depankeikembang_5050kombinasi_db, &depankeikempis_5050kombinasi_db, &depankeikembar_5050kombinasi_db,
		&belakangdiscmono_5050kombinasi_db, &belakangdiscstereo_5050kombinasi_db, &belakangdisckembang_5050kombinasi_db, &belakangdisckempis_5050kombinasi_db, &belakangdisckembar_5050kombinasi_db,
		&tengahdiscmono_5050kombinasi_db, &tengahdiscstereo_5050kombinasi_db, &tengahdisckembang_5050kombinasi_db, &tengahdisckempis_5050kombinasi_db, &tengahdisckembar_5050kombinasi_db,
		&depandiscmono_5050kombinasi_db, &depandiscstereo_5050kombinasi_db, &depandisckembang_5050kombinasi_db, &depandisckempis_5050kombinasi_db, &depandisckembar_5050kombinasi_db,
		&limitglobal_5050kombinasi_db, &limittotal_5050kombinasi_db,
		&minbet_kombinasi_db, &maxbet_kombinasi_db, &maxbuy_kombinasi_db, &win_kombinasi_db, &disc_kombinasi_db, &limitglobal_kombinasi_db, &limittotal_kombinasi_db,
		&minbet_dasar_db, &maxbet_dasar_db, &maxbuy_dasar_db, &keibesar_dasar_db, &keikecil_dasar_db, &keigenap_dasar_db, &keiganjil_dasar_db, &discbesar_dasar_db, &disckecil_dasar_db, &discgenap_dasar_db, &discganjil_dasar_db, &limitglobal_dasar_db, &limittotal_dasar_db,
		&minbet_shio_db, &maxbet_shio_db, &maxbuy_shio_db, &win_shio_db, &disc_shio_db, &limitglobal_shio_db, &limittotal_shio_db, &shioyear_shio_db,
		&createcomppas_db,
		&createdatecomppas_db,
		&updatecomppas_db,
		&updatedatecompas_db)

	helpers.ErrorCheck(err)
	obj.Idpasarantogel = idpasarantogel_db
	obj.Nmpasarantogel = nmpasarantogel_db
	obj.PasaranDiundi = pasarandiundi_db
	obj.PasaranURL = pasaranurl_db
	obj.Jamtutup = jamtutup_db
	obj.Jamjadwal = jamjadwal_db
	obj.Jamopen = jamopen_db
	obj.Limitline4d = limitline_4d_db
	obj.Limitline3d = limitline_3d_db
	obj.Limitline3dd = limitline_3dd_db
	obj.Limitline2d = limitline_2d_db
	obj.Limitline2dd = limitline_2dd_db
	obj.Limitline2dt = limitline_2dt_db
	obj.Bbfs = bbfs_db
	obj.Minbet_432d = minbet_432d_db
	obj.Maxbet4d_432d = maxbet4d_432d_db
	obj.Maxbet3d_432d = maxbet3d_432d_db
	obj.Maxbet3dd_432d = maxbet3dd_432d_db
	obj.Maxbet2d_432d = maxbet2d_432d_db
	obj.Maxbet2dd_432d = maxbet2dd_432d_db
	obj.Maxbet2dt_432d = maxbet2dt_432d_db
	obj.Maxbet4d_fullbb_432d = maxbet4d_fullbb_432d_db
	obj.Maxbet3d_fullbb_432d = maxbet3d_fullbb_432d_db
	obj.Maxbet3dd_fullbb_432d = maxbet3dd_fullbb_432d_db
	obj.Maxbet2d_fullbb_432d = maxbet2d_fullbb_432d_db
	obj.Maxbet2dd_fullbb_432d = maxbet2dd_fullbb_432d_db
	obj.Maxbet2dt_fullbb_432d = maxbet2dt_fullbb_432d_db
	obj.Maxbuy4d_432d = maxbuy4d_432d_db
	obj.Maxbuy3d_432d = maxbuy3d_432d_db
	obj.Maxbuy3dd_432d = maxbuy3dd_432d_db
	obj.Maxbuy2d_432d = maxbuy2d_432d_db
	obj.Maxbuy2dd_432d = maxbuy2dd_432d_db
	obj.Maxbuy2dt_432d = maxbuy2dt_432d_db
	obj.Limitotal4d_432d = limitotal4d_432d_db
	obj.Limitotal3d_432d = limitotal3d_432d_db
	obj.Limitotal3dd_432d = limitotal3dd_432d_db
	obj.Limitotal2d_432d = limitotal2d_432d_db
	obj.Limitotal2dd_432d = limitotal2dd_432d_db
	obj.Limitotal2dt_432d = limitotal2dt_432d_db
	obj.Limitglobal4d_432d = limitglobal4d_432d_db
	obj.Limitglobal3d_432d = limitglobal3d_432d_db
	obj.Limitglobal3dd_432d = limitglobal3dd_432d_db
	obj.Limitglobal2d_432d = limitglobal2d_432d_db
	obj.Limitglobal2dd_432d = limitglobal2dd_432d_db
	obj.Limitglobal2dt_432d = limitglobal2dt_432d_db
	obj.Limitotal4d_fullbb_432d = limitbuang4d_fullbb_432d_db
	obj.Limitotal3d_fullbb_432d = limitbuang3d_fullbb_432d_db
	obj.Limitotal3dd_fullbb_432d = limitbuang3dd_fullbb_432d_db
	obj.Limitotal2d_fullbb_432d = limitbuang2d_fullbb_432d_db
	obj.Limitotal2dd_fullbb_432d = limitbuang2dd_fullbb_432d_db
	obj.Limitotal2dt_fullbb_432d = limitbuang2dt_fullbb_432d_db
	obj.Limitglobal4d_fullbb_432d = limitbuang4d_fullbb_432d_db
	obj.Limitglobal3d_fullbb_432d = limitbuang3d_fullbb_432d_db
	obj.Limitglobal3dd_fullbb_432d = limitbuang3dd_fullbb_432d_db
	obj.Limitglobal2d_fullbb_432d = limitbuang2d_fullbb_432d_db
	obj.Limitglobal2dd_fullbb_432d = limitbuang2dd_fullbb_432d_db
	obj.Limitglobal2dt_fullbb_432d = limitbuang2dt_fullbb_432d_db
	obj.Disc4d_432d = disc4d_432d_db
	obj.Disc3d_432d = disc3d_432d_db
	obj.Disc3dd_432d = disc3dd_432d_db
	obj.Disc2d_432d = disc2d_432d_db
	obj.Disc2dd_432d = disc2dd_432d_db
	obj.Disc2dt_432d = disc2dt_432d_db
	obj.Win4d_432d = win4d_432d_db
	obj.Win3d_432d = win3d_432d_db
	obj.Win3dd_432d = win3dd_432d_db
	obj.Win2d_432d = win2d_432d_db
	obj.Win2dd_432d = win2dd_432d_db
	obj.Win2dt_432d = win2dt_432d_db
	obj.Win4dnodisc_432d = win4dnodisc_432d_db
	obj.Win3dnodisc_432d = win3dnodisc_432d_db
	obj.Win3ddnodisc_432d = win3ddnodisc_432d_db
	obj.Win2dnodisc_432d = win2dnodisc_432d_db
	obj.Win2ddnodisc_432d = win2ddnodisc_432d_db
	obj.Win2dtnodisc_432d = win2dtnodisc_432d_db
	obj.Win4dbb_kena_432d = win4dbb_kena_432d_db
	obj.Win3dbb_kena_432d = win3dbb_kena_432d_db
	obj.Win3ddbb_kena_432d = win3ddbb_kena_432d_db
	obj.Win2dbb_kena_432d = win2dbb_kena_432d_db
	obj.Win2ddbb_kena_432d = win2ddbb_kena_432d_db
	obj.Win2dtbb_kena_432d = win2dtbb_kena_432d_db
	obj.Win4dbb_432d = win4dbb_432d_db
	obj.Win3dbb_432d = win3dbb_432d_db
	obj.Win3ddbb_432d = win3ddbb_432d_db
	obj.Win2dbb_432d = win2dbb_432d_db
	obj.Win2ddbb_432d = win2ddbb_432d_db
	obj.Win2dtbb_432d = win2dtbb_432d_db
	obj.Minbet_cbebas = minbet_cbebas_db
	obj.Maxbet_cbebas = maxbet_cbebas_db
	obj.Maxbuy_cbebas = maxbuy_cbebas_db
	obj.Win_cbebas = win_cbebas_db
	obj.Disc_cbebas = disc_cbebas_db
	obj.Limitglobal_cbebas = limitglobal_cbebas_db
	obj.Limittotal_cbebas = limittotal_cbebas_db
	obj.Minbet_cmacau = minbet_cmacau_db
	obj.Maxbet_cmacau = maxbet_cmacau_db
	obj.Maxbuy_cmacau = maxbuy_cmacau_db
	obj.Win2d_cmacau = win2d_cmacau_db
	obj.Win3d_cmacau = win3d_cmacau_db
	obj.Win4d_cmacau = win4d_cmacau_db
	obj.Disc_cmacau = disc_cmacau_db
	obj.Limitglobal_cmacau = limitglobal_cmacau_db
	obj.Limitotal_cmacau = limitotal_cmacau_db
	obj.Minbet_cnaga = minbet_cnaga_db
	obj.Maxbet_cnaga = maxbet_cnaga_db
	obj.Maxbuy_cnaga = maxbuy_cnaga_db
	obj.Win3_cnaga = win3_cnaga_db
	obj.Win4_cnaga = win4_cnaga_db
	obj.Disc_cnaga = disc_cnaga_db
	obj.Limitglobal_cnaga = limitglobal_cnaga_db
	obj.Limittotal_cnaga = limittotal_cnaga_db
	obj.Minbet_cjitu = minbet_cjitu_db
	obj.Maxbet_cjitu = maxbet_cjitu_db
	obj.Maxbuy_cjitu = maxbuy_cjitu_db
	obj.Winas_cjitu = winas_cjitu_db
	obj.Winkop_cjitu = winkop_cjitu_db
	obj.Winkepala_cjitu = winkepala_cjitu_db
	obj.Winekor_cjitu = winekor_cjitu_db
	obj.Desc_cjitu = desc_cjitu_db
	obj.Limitglobal_cjitu = limitglobal_cjitu_db
	obj.Limittotal_cjitu = limittotal_cjitu_db
	obj.Minbet_5050umum = minbet_5050umum_db
	obj.Maxbet_5050umum = maxbet_5050umum_db
	obj.Maxbuy_5050umum = maxbuy_5050umum_db
	obj.Keibesar_5050umum = keibesar_5050umum_db
	obj.Keikecil_5050umum = keikecil_5050umum_db
	obj.Keigenap_5050umum = keigenap_5050umum_db
	obj.Keiganjil_5050umum = keiganjil_5050umum_db
	obj.Keitengah_5050umum = keitengah_5050umum_db
	obj.Keitepi_5050umum = keitengah_5050umum_db
	obj.Discbesar_5050umum = discbesar_5050umum_db
	obj.Disckecil_5050umum = disckecil_5050umum_db
	obj.Discgenap_5050umum = discgenap_5050umum_db
	obj.Discganjil_5050umum = discganjil_5050umum_db
	obj.Disctengah_5050umum = disctengah_5050umum_db
	obj.Disctepi_5050umum = disctepi_5050umum_db
	obj.Limitglobal_5050umum = limitglobal_5050umum_db
	obj.Limittotal_5050umum = limittotal_5050umum_db
	obj.Minbet_5050special = minbet_5050special_db
	obj.Maxbet_5050special = maxbet_5050special_db
	obj.Maxbuy_5050special = maxbuy_5050special_db
	obj.Keiasganjil_5050special = keiasganjil_5050special_db
	obj.Keiasgenap_5050special = keiasgenap_5050special_db
	obj.Keiasbesar_5050special = keiasbesar_5050special_db
	obj.Keiaskecil_5050special = keiaskecil_5050special_db
	obj.Keikopganjil_5050special = keikopganjil_5050special_db
	obj.Keikopgenap_5050special = keikopgenap_5050special_db
	obj.Keikopbesar_5050special = keikopbesar_5050special_db
	obj.Keikopkecil_5050special = keikopkecil_5050special_db
	obj.Keikepalaganjil_5050special = keikepalaganjil_5050special_db
	obj.Keikepalagenap_5050special = keikepalagenap_5050special_db
	obj.Keikepalabesar_5050special = keikepalabesar_5050special_db
	obj.Keikepalakecil_5050special = keikepalakecil_5050special_db
	obj.Keiekorganjil_5050special = keiekorganjil_5050special_db
	obj.Keiekorgenap_5050special = keiekorgenap_5050special_db
	obj.Keiekorbesar_5050special = keiekorbesar_5050special_db
	obj.Keiekorkecil_5050special = keiekorkecil_5050special_db
	obj.Discasganjil_5050special = discasganjil_5050special_db
	obj.Discasgenap_5050special = discasgenap_5050special_db
	obj.Discasbesar_5050special = discasbesar_5050special_db
	obj.Discaskecil_5050special = discaskecil_5050special_db
	obj.Disckopganjil_5050special = disckopganjil_5050special_db
	obj.Disckopgenap_5050special = disckopgenap_5050special_db
	obj.Disckopbesar_5050special = disckopbesar_5050special_db
	obj.Disckopkecil_5050special = disckopkecil_5050special_db
	obj.Disckepalaganjil_5050special = disckepalaganjil_5050special_db
	obj.Disckepalagenap_5050special = disckepalagenap_5050special_db
	obj.Disckepalabesar_5050special = disckepalabesar_5050special_db
	obj.Disckepalakecil_5050special = disckepalakecil_5050special_db
	obj.Discekorganjil_5050special = discekorganjil_5050special_db
	obj.Discekorgenap_5050special = discekorgenap_5050special_db
	obj.Discekorbesar_5050special = discekorbesar_5050special_db
	obj.Discekorkecil_5050special = discekorkecil_5050special_db
	obj.Limitglobal_5050special = limitglobal_5050special_db
	obj.Limittotal_5050special = limittotal_5050special_db
	obj.Minbet_5050kombinasi = minbet_5050kombinasi_db
	obj.Maxbet_5050kombinasi = maxbet_5050kombinasi_db
	obj.Maxbuy_5050kombinasi = maxbuy_5050kombinasi_db
	obj.Belakangkeimono_5050kombinasi = belakangkeimono_5050kombinasi_db
	obj.Belakangkeistereo_5050kombinasi = belakangkeistereo_5050kombinasi_db
	obj.Belakangkeikembang_5050kombinasi = belakangkeikembang_5050kombinasi_db
	obj.Belakangkeikempis_5050kombinasi = belakangkeikempis_5050kombinasi_db
	obj.Belakangkeikembar_5050kombinasi = belakangkeikembang_5050kombinasi_db
	obj.Tengahkeimono_5050kombinasi = tengahkeimono_5050kombinasi_db
	obj.Tengahkeistereo_5050kombinasi = tengahkeistereo_5050kombinasi_db
	obj.Tengahkeikembang_5050kombinasi = tengahkeikembang_5050kombinasi_db
	obj.Tengahkeikempis_5050kombinasi = tengahkeikempis_5050kombinasi_db
	obj.Tengahkeikembar_5050kombinasi = tengahkeikembar_5050kombinasi_db
	obj.Depankeimono_5050kombinasi = depankeimono_5050kombinasi_db
	obj.Depankeistereo_5050kombinasi = depankeistereo_5050kombinasi_db
	obj.Depankeikembang_5050kombinasi = depankeikembang_5050kombinasi_db
	obj.Depankeikempis_5050kombinasi = depankeikempis_5050kombinasi_db
	obj.Depankeikembar_5050kombinasi = depankeikembar_5050kombinasi_db
	obj.Belakangdiscmono_5050kombinasi = belakangdiscmono_5050kombinasi_db
	obj.Belakangdiscstereo_5050kombinasi = belakangdiscstereo_5050kombinasi_db
	obj.Belakangdisckembang_5050kombinasi = belakangdisckembang_5050kombinasi_db
	obj.Belakangdisckempis_5050kombinasi = belakangdisckempis_5050kombinasi_db
	obj.Belakangdisckembar_5050kombinasi = belakangdisckembang_5050kombinasi_db
	obj.Tengahdiscmono_5050kombinasi = tengahdiscmono_5050kombinasi_db
	obj.Tengahdiscstereo_5050kombinasi = tengahdiscstereo_5050kombinasi_db
	obj.Tengahdisckembang_5050kombinasi = tengahdisckembang_5050kombinasi_db
	obj.Tengahdisckempis_5050kombinasi = tengahdisckempis_5050kombinasi_db
	obj.Tengahdisckembar_5050kombinasi = tengahdisckembar_5050kombinasi_db
	obj.Depandiscmono_5050kombinasi = depandiscstereo_5050kombinasi_db
	obj.Depandiscstereo_5050kombinasi = depandiscstereo_5050kombinasi_db
	obj.Depandisckembang_5050kombinasi = depandisckembang_5050kombinasi_db
	obj.Depandisckempis_5050kombinasi = depandisckempis_5050kombinasi_db
	obj.Depandisckembar_5050kombinasi = depandisckembang_5050kombinasi_db
	obj.Limitglobal_5050kombinasi = limitglobal_5050kombinasi_db
	obj.Limittotal_5050kombinasi = limittotal_5050kombinasi_db
	obj.Minbet_kombinasi = minbet_kombinasi_db
	obj.Maxbet_kombinasi = maxbet_kombinasi_db
	obj.Maxbuy_kombinasi = maxbuy_kombinasi_db
	obj.Win_kombinasi = win_kombinasi_db
	obj.Disc_kombinasi = disc_kombinasi_db
	obj.Limitglobal_kombinasi = limitglobal_kombinasi_db
	obj.Limittotal_kombinasi = limittotal_kombinasi_db
	obj.Minbet_dasar = minbet_dasar_db
	obj.Maxbet_dasar = maxbet_dasar_db
	obj.Maxbuy_dasar = maxbuy_dasar_db
	obj.Keibesar_dasar = keibesar_dasar_db
	obj.Keikecil_dasar = keikecil_dasar_db
	obj.Keigenap_dasar = keigenap_dasar_db
	obj.Keiganjil_dasar = keiganjil_dasar_db
	obj.Discbesar_dasar = discbesar_dasar_db
	obj.Disckecil_dasar = disckecil_dasar_db
	obj.Discgenap_dasar = discgenap_dasar_db
	obj.Discganjil_dasar = discganjil_dasar_db
	obj.Limitglobal_dasar = limitglobal_dasar_db
	obj.Limittotal_dasar = limittotal_dasar_db
	obj.Minbet_shio = minbet_shio_db
	obj.Maxbet_shio = maxbet_shio_db
	obj.Maxbuy_shio = maxbuy_shio_db
	obj.Win_shio = win_shio_db
	obj.Disc_shio = disc_shio_db
	obj.Shioyear_shio = shioyear_shio_db
	obj.Limitglobal_shio = limitglobal_shio_db
	obj.Limittotal_shio = limittotal_shio_db
	obj.StatusPasaranActive = statuspasaranactive_db
	obj.Displaypasaran = displaypasaran_db
	obj.Create = createcomppas_db
	obj.Createdate = createdatecomppas_db
	obj.Update = updatecomppas_db
	obj.Updatedate = updatedatecompas_db
	arraobj = append(arraobj, obj)

	var objpasaranonline pasaranEdit_Online
	var arraobjpasaranonline []pasaranEdit_Online
	sqlpasaranonline := `SELECT
		idcomppasaranoff , haripasaran
		FROM ` + config.DB_tbl_mst_company_game_pasaran_offline + ` 
		WHERE idcompany = ?
		AND idcomppasaran = ?
	`
	row, err := con.QueryContext(ctx, sqlpasaranonline, company, idcomppasaran)
	defer row.Close()

	if err != nil {
		helpers.ErrorCheck(err)
	}

	for row.Next() {
		var idcomppasaranoff_db int
		var haripasaran_db string

		err = row.Scan(&idcomppasaranoff_db, &haripasaran_db)

		if err != nil {
			helpers.ErrorCheck(err)
		}

		objpasaranonline.Idpasaranonline = idcomppasaranoff_db
		objpasaranonline.Haripasaran = haripasaran_db
		arraobjpasaranonline = append(arraobjpasaranonline, objpasaranonline)
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Pasaraonline = arraobjpasaranonline
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_Pasaran(agent string, company string, idcomppasaran int, pasarandiundi, pasaranurl, jamtutup, jamjadwal, jamopen, statuspasaranactive string, displaypasaran int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	msg := "Failed"

	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		stmt, e := con.Prepare(`
					UPDATE 
					` + config.DB_tbl_mst_company_game_pasaran + `  
					SET pasarandiundi=? , pasaranurl=?, 
					jamtutup=?, jamjadwal=?, jamopen=?, statuspasaranactive=?, displaypasaran=?, 
					updatecomppas=?, updatedatecompas=? 
					WHERE idcomppasaran=? AND idcompany=? 
				`)
		helpers.ErrorCheck(e)
		rec, e := stmt.ExecContext(ctx,
			pasarandiundi,
			pasaranurl,
			jamtutup,
			jamjadwal,
			jamopen,
			statuspasaranactive,
			displaypasaran,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)
		helpers.ErrorCheck(e)

		update, err_update := rec.RowsAffected()
		helpers.ErrorCheck(err_update)
		if update > 0 {
			msg = "Success"

			idpasarantogel, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
			nmpasarantogel := Pasaranmaster_id(idpasarantogel, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "PASARAN DI UNDI - " + pasarandiundi + "<br />"
			noteafter += "PASARAN URL - " + pasaranurl + "<br />"
			noteafter += "PASARAN JAM TUTUP - " + jamtutup + "<br />"
			noteafter += "PASARAN JAM JADWAL - " + jamjadwal + "<br />"
			noteafter += "PASARAN JAM OPEN - " + jamopen + "<br />"
			noteafter += "PASARAN STATUS - " + statuspasaranactive + "<br />"
			noteafter += "PASARAN DISPLAY - " + strconv.Itoa(displaypasaran)
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN", "", noteafter)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_PasaranOnline(agent, company string, idcomppasaran int, haripasaran string) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	flag := false
	msg := "Failed"

	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		flag = Get_OnlinePasaran(company, idcomppasaran, haripasaran, "hari")
		if !flag {
			field_col := "tbl_mst_company_game_pasaran_offline_" + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_col)
			sql_insert := `
			INSERT INTO 
			` + config.DB_tbl_mst_company_game_pasaran_offline + `(
				idcomppasaranoff , idcomppasaran, idcompany, haripasaran,
				createcomppasaranoff, createdatecomppasaranoff)
			VALUES(?,?,?,?,?,?)
		`
			stmt, e := con.PrepareContext(ctx, sql_insert)
			helpers.ErrorCheck(e)
			var idrecord string = tglnow.Format("YYYY") + strconv.Itoa(idrecord_counter)
			rec, e := stmt.ExecContext(ctx,
				idrecord,
				idcomppasaran,
				company,
				haripasaran,
				agent,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"))
			helpers.ErrorCheck(e)

			insert, err_insert := rec.RowsAffected()
			helpers.ErrorCheck(err_insert)
			fmt.Printf("The last inserted row id: %d\n", insert)
			defer stmt.Close()
			if insert > 0 {
				msg = "Success"
				idpasarantogel, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
				nmpasarantogel := Pasaranmaster_id(idpasarantogel, "nmpasarantogel")

				noteafter := ""
				noteafter += "PASARAN - " + nmpasarantogel + "<br />"
				noteafter += "HARI PASARAN - " + haripasaran
				Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - NEW ONLINE", "", noteafter)
			}
		} else {
			msg = "Duplicate Entry"
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Delete_PasaranOnline(company string, idcomppasaran, idcomppasaranoff int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		stmt, e := con.PrepareContext(ctx, `
					DELETE FROM 
					`+config.DB_tbl_mst_company_game_pasaran_offline+`  
					WHERE idcomppasaranoff = ?
					AND idcomppasaran = ? 
					AND idcompany = ? 
				`)
		helpers.ErrorCheck(e)
		rec, e := stmt.ExecContext(ctx,
			idcomppasaranoff,
			idcomppasaran,
			company)
		helpers.ErrorCheck(e)

		delete, err_delete := rec.RowsAffected()
		helpers.ErrorCheck(err_delete)
		fmt.Printf("The last delete row id: %d\n", delete)
		defer stmt.Close()
		if delete > 0 {
			msg = "Success"
			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			haripasaran := _pasaran_online(idcomppasaran, company, "haripasaran")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "HARI PASARAN - " + haripasaran
			Insert_log(company, "", "PASARAN", "UPDATE PASARAN - DELETE ONLINE", "", noteafter)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Save_PasaranLimitline(agent string, company string, idcomppasaran int, limitline4d, limitline3d, limitline3dd, limitline2d, limitline2dd, limitline2dt int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		stmt, e := con.PrepareContext(ctx, `
					UPDATE 
					`+config.DB_tbl_mst_company_game_pasaran+`  
					SET limitline_4d=? , limitline_3d=?, limitline_3dd=?, 
					limitline_2d=?, limitline_2dd=?, limitline_2dt=?,   
					updatecomppas=?, updatedatecompas=? 
					WHERE idcomppasaran=? AND idcompany=? 
				`)
		helpers.ErrorCheck(e)
		rec, e := stmt.ExecContext(ctx, limitline4d, limitline3d, limitline3dd, limitline2d, limitline2dd, limitline2dt,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)
		helpers.ErrorCheck(e)

		update, err_update := rec.RowsAffected()
		helpers.ErrorCheck(err_update)
		if update > 0 {
			msg = "Success"

			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "LIMITLINE 4D - " + strconv.Itoa(limitline4d) + "<br />"
			noteafter += "LIMITLINE 3D - " + strconv.Itoa(limitline3d) + "<br />"
			noteafter += "LIMITLINE 3DD - " + strconv.Itoa(limitline3dd) + "<br />"
			noteafter += "LIMITLINE 2D - " + strconv.Itoa(limitline2d) + "<br />"
			noteafter += "LIMITLINE 2DD - " + strconv.Itoa(limitline2dd) + "<br />"
			noteafter += "LIMITLINE 2DT - " + strconv.Itoa(limitline2dt)
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - LIMITLINE", "", noteafter)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Save_PasaranConf432(
	agent string,
	company string,
	idcomppasaran int,
	minbet, maxbet4d, maxbet3d, maxbet3dd, maxbet2d, maxbet2dd, maxbet2dt int,
	maxbet4d_fulbb, maxbet3d_fulbb, maxbet3dd_fulbb, maxbet2d_fulbb, maxbet2dd_fulbb, maxbet2dt_fulbb int,
	maxbuy4d, maxbuy3d, maxbuy3dd, maxbuy2d, maxbuy2dd, maxbuy2dt int,
	win4d, win3d, win3dd, win2d, win2dd, win2dt int,
	win4dnodisc, win3dnodisc, win3ddnodisc, win2dnodisc, win2ddnodisc, win2dtnodisc int,
	win4dbb_kena, win3dbb_kena, win3ddbb_kena, win2dbb_kena, win2ddbb_kena, win2dtbb_kena int,
	win4dbb, win3dbb, win3ddbb, win2dbb, win2ddbb, win2dtbb int,
	disc4d, disc3d, disc3dd, disc2d, disc2dd, disc2dt float32,
	limitglobal4d, limitglobal3d, limitglobal3dd, limitglobal2d, limitglobal2dd, limitglobal2dt int,
	limittotal4d, limittotal3d, limittotal3dd, limittotal2d, limittotal2dd, limittotal2dt int,
	limitglobal4d_fullbb, limitglobal3d_fullbb, limitglobal3dd_fullbb, limitglobal2d_fullbb, limitglobal2dd_fullbb, limitglobal2dt_fullbb int,
	limittotal4d_fullbb, limittotal3d_fullbb, limittotal3dd_fullbb, limittotal2d_fullbb, limittotal2dd_fullbb, limittotal2dt_fullbb int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	render_page := time.Now()

	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		sql_update := `
			UPDATE 
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 1_minbet=? , 1_maxbet4d=?, 1_maxbet3d=?, 1_maxbet3dd=?, 1_maxbet2d=?, 1_maxbet2dd=?, 1_maxbet2dt=?, 
			1_maxbet4d_fullbb=?, 1_maxbet3d_fullbb=?, 1_maxbet3dd_fullbb=?, 1_maxbet2d_fullbb=?, 1_maxbet2dd_fullbb=?, 1_maxbet2dt_fullbb=?, 
			1_maxbuy4d=?, 1_maxbuy3d=?, 1_maxbuy3dd=?, 1_maxbuy2d=?, 1_maxbuy2dd=?, 1_maxbuy2dt=?, 
			1_win4d=?, 1_win3d=?, 1_win3dd=?, 1_win2d=?, 1_win2dd=?, 1_win2dt=?, 
			1_win4dnodisc=?, 1_win3dnodisc=?, 1_win3ddnodisc=?, 1_win2dnodisc=?, 1_win2ddnodisc=?, 1_win2dtnodisc=?, 
			1_win4dbb_kena=?, 1_win3dbb_kena=?, 1_win3ddbb_kena=?, 1_win2dbb_kena=?, 1_win2ddbb_kena=?, 1_win2dtbb_kena=?, 
			1_win4dbb=?, 1_win3dbb=?, 1_win3ddbb=?, 1_win2dbb=?, 1_win2ddbb=?, 1_win2dtbb=?, 
			1_disc4d=?, 1_disc3d=?, 1_disc3dd=?, 1_disc2d=?, 1_disc2dd=?, 1_disc2dt=?, 
			1_limitbuang4d=?, 1_limitbuang3d=?, 1_limitbuang3dd=?, 1_limitbuang2d=?, 1_limitbuang2dd=?, 1_limitbuang2dt=?,  
			1_limittotal4d=?, 1_limittotal3d=?, 1_limittotal3dd=?, 1_limittotal2d=?, 1_limittotal2dd=?, 1_limittotal2dt=?,  
			1_limitbuang4d_fullbb=?, 1_limitbuang3d_fullbb=?, 1_limitbuang3dd_fullbb=?, 1_limitbuang2d_fullbb=?, 1_limitbuang2dd_fullbb=?, 1_limitbuang2dt_fullbb=?,  
			1_limittotal4d_fullbb=?, 1_limittotal3d_fullbb=?, 1_limittotal3dd_fullbb=?, 1_limittotal2d_fullbb=?, 1_limittotal2dd_fullbb=?, 1_limittotal2dt_fullbb=?,   
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println("432D")
		log.Printf("DISC 4D : %s", fmt.Sprintf("%.3f", disc4d))
		log.Printf("DISC 3D : %s", fmt.Sprintf("%.3f", disc3d))
		log.Printf("DISC 3DD : %s", fmt.Sprintf("%.3f", disc3dd))
		log.Printf("DISC 2D : %s", fmt.Sprintf("%.3f", disc2d))
		log.Printf("DISC 2DD : %s", fmt.Sprintf("%.3f", disc2dd))
		log.Printf("DISC 2DT : %s", fmt.Sprintf("%.3f", disc2dt))
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_company_game_pasaran, "UPDATE",
			minbet, maxbet4d, maxbet3d, maxbet3dd, maxbet2d, maxbet2dd, maxbet2dt,
			maxbet4d_fulbb, maxbet3d_fulbb, maxbet3dd_fulbb, maxbet2d_fulbb, maxbet2dd_fulbb, maxbet2dt_fulbb,
			maxbuy4d, maxbuy3d, maxbuy3dd, maxbuy2d, maxbuy2dd, maxbuy2dt,
			win4d, win3d, win3dd, win2d, win2dd, win2dt,
			win4dnodisc, win3dnodisc, win3ddnodisc, win2dnodisc, win2ddnodisc, win2dtnodisc,
			win4dbb_kena, win3dbb_kena, win3ddbb_kena, win2dbb_kena, win2ddbb_kena, win2dtbb_kena,
			win4dbb, win3dbb, win3ddbb, win2dbb, win2ddbb, win2dtbb,
			fmt.Sprintf("%.3f", disc4d), fmt.Sprintf("%.3f", disc3d), fmt.Sprintf("%.3f", disc3dd), fmt.Sprintf("%.3f", disc2d), fmt.Sprintf("%.3f", disc2dd), fmt.Sprintf("%.3f", disc2dt),
			limitglobal4d, limitglobal3d, limitglobal3dd, limitglobal2d, limitglobal2dd, limitglobal2dt,
			limittotal4d, limittotal3d, limittotal3dd, limittotal2d, limittotal2dd, limittotal2dt,
			limitglobal4d_fullbb, limitglobal3d_fullbb, limitglobal3dd_fullbb, limitglobal2d_fullbb, limitglobal2dd_fullbb, limitglobal2dt_fullbb,
			limittotal4d_fullbb, limittotal3d_fullbb, limittotal3dd_fullbb, limittotal2d_fullbb, limittotal2dd_fullbb, limittotal2dt_fullbb,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)

		if flag_update {
			msg = "Success"
			log.Println(msg_update)

			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "MINIMAL BET - " + strconv.Itoa(minbet) + "<br />"
			noteafter += "MAXIMAL BET4D - " + strconv.Itoa(maxbet4d) + "<br />"
			noteafter += "MAXIMAL BET3D - " + strconv.Itoa(maxbet3d) + "<br />"
			noteafter += "MAXIMAL BET3DD - " + strconv.Itoa(maxbet3dd) + "<br />"
			noteafter += "MAXIMAL BET2D - " + strconv.Itoa(maxbet2d) + "<br />"
			noteafter += "MAXIMAL BET2DD - " + strconv.Itoa(maxbet2dd) + "<br />"
			noteafter += "MAXIMAL BET2DT - " + strconv.Itoa(maxbet2dt) + "<br />"
			noteafter += "MAXIMAL BET4D FULLBB - " + strconv.Itoa(maxbet4d_fulbb) + "<br />"
			noteafter += "MAXIMAL BET3D FULLBB - " + strconv.Itoa(maxbet3d_fulbb) + "<br />"
			noteafter += "MAXIMAL BET3DD FULLBB - " + strconv.Itoa(maxbet3dd_fulbb) + "<br />"
			noteafter += "MAXIMAL BET2D FULLBB - " + strconv.Itoa(maxbet2d_fulbb) + "<br />"
			noteafter += "MAXIMAL BET2DD FULLBB - " + strconv.Itoa(maxbet2dd_fulbb) + "<br />"
			noteafter += "MAXIMAL BET2DT FULLBB - " + strconv.Itoa(maxbet2dt_fulbb) + "<br />"
			noteafter += "MAXIMAL BUY4D - " + strconv.Itoa(maxbuy4d) + "<br />"
			noteafter += "MAXIMAL BUY3D - " + strconv.Itoa(maxbuy3d) + "<br />"
			noteafter += "MAXIMAL BUY3DD - " + strconv.Itoa(maxbuy3dd) + "<br />"
			noteafter += "MAXIMAL BUY2D - " + strconv.Itoa(maxbuy2d) + "<br />"
			noteafter += "MAXIMAL BUY2DD - " + strconv.Itoa(maxbuy2dd) + "<br />"
			noteafter += "MAXIMAL BUY2DT - " + strconv.Itoa(maxbuy2dt) + "<br />"
			noteafter += "WIN4D - " + strconv.Itoa(win4d) + "<br />"
			noteafter += "WIN3D - " + strconv.Itoa(win3d) + "<br />"
			noteafter += "WIN3DD - " + strconv.Itoa(win3dd) + "<br />"
			noteafter += "WIN2D - " + strconv.Itoa(win2d) + "<br />"
			noteafter += "WIN2DD - " + strconv.Itoa(win2dd) + "<br />"
			noteafter += "WIN2DT - " + strconv.Itoa(win2dt) + "<br />"
			noteafter += "WIN4D NO DISKON - " + strconv.Itoa(win4dnodisc) + "<br />"
			noteafter += "WIN3D NO DISKON - " + strconv.Itoa(win3dnodisc) + "<br />"
			noteafter += "WIN3DD NO DISKON - " + strconv.Itoa(win3ddnodisc) + "<br />"
			noteafter += "WIN2D NO DISKON - " + strconv.Itoa(win2dnodisc) + "<br />"
			noteafter += "WIN2DD NO DISKON - " + strconv.Itoa(win2ddnodisc) + "<br />"
			noteafter += "WIN2DT NO DISKON - " + strconv.Itoa(win2dtnodisc) + "<br />"
			noteafter += "WIN4D BB KENA - " + strconv.Itoa(win4dbb_kena) + "<br />"
			noteafter += "WIN3D BB KENA - " + strconv.Itoa(win3dbb_kena) + "<br />"
			noteafter += "WIN3DD BB KENA - " + strconv.Itoa(win3ddbb_kena) + "<br />"
			noteafter += "WIN2D BB KENA - " + strconv.Itoa(win2dbb_kena) + "<br />"
			noteafter += "WIN2DD BB KENA - " + strconv.Itoa(win2dbb_kena) + "<br />"
			noteafter += "WIN2DT BB KENA - " + strconv.Itoa(win2dtbb_kena) + "<br />"
			noteafter += "WIN4D BB  - " + strconv.Itoa(win4dbb) + "<br />"
			noteafter += "WIN3D BB  - " + strconv.Itoa(win3dbb) + "<br />"
			noteafter += "WIN3DD BB  - " + strconv.Itoa(win3ddbb) + "<br />"
			noteafter += "WIN2D BB  - " + strconv.Itoa(win2dbb) + "<br />"
			noteafter += "WIN2DD BB  - " + strconv.Itoa(win2ddbb) + "<br />"
			noteafter += "WIN2DT BB  - " + strconv.Itoa(win2dtbb) + "<br />"
			noteafter += "DISC4D - " + fmt.Sprintf("%.3f", disc4d) + "<br />"
			noteafter += "DISC3D - " + fmt.Sprintf("%.3f", disc3d) + "<br />"
			noteafter += "DISC3DD - " + fmt.Sprintf("%.3f", disc3dd) + "<br />"
			noteafter += "DISC2D - " + fmt.Sprintf("%.3f", disc2d) + "<br />"
			noteafter += "DISC2DD - " + fmt.Sprintf("%.3f", disc2dd) + "<br />"
			noteafter += "DISC2DT - " + fmt.Sprintf("%.3f", disc2dt) + "<br />"
			noteafter += "LIMIT GLOBAL 4D - " + strconv.Itoa(limitglobal4d) + "<br />"
			noteafter += "LIMIT GLOBAL 3D - " + strconv.Itoa(limitglobal3d) + "<br />"
			noteafter += "LIMIT GLOBAL 3DD - " + strconv.Itoa(limitglobal3dd) + "<br />"
			noteafter += "LIMIT GLOBAL 2D - " + strconv.Itoa(limitglobal2d) + "<br />"
			noteafter += "LIMIT GLOBAL 2DD - " + strconv.Itoa(limitglobal2dd) + "<br />"
			noteafter += "LIMIT GLOBAL 2DT - " + strconv.Itoa(limitglobal2dt) + "<br />"
			noteafter += "LIMIT TOTAL 4D - " + strconv.Itoa(limittotal4d) + "<br />"
			noteafter += "LIMIT TOTAL 3D - " + strconv.Itoa(limittotal3d) + "<br />"
			noteafter += "LIMIT TOTAL 3DD - " + strconv.Itoa(limittotal3dd) + "<br />"
			noteafter += "LIMIT TOTAL 2D - " + strconv.Itoa(limittotal2d) + "<br />"
			noteafter += "LIMIT TOTAL 2DD - " + strconv.Itoa(limittotal2dd) + "<br />"
			noteafter += "LIMIT GLOBAL 4D FULLBB - " + strconv.Itoa(limitglobal4d_fullbb) + "<br />"
			noteafter += "LIMIT GLOBAL 3D FULLBB - " + strconv.Itoa(limitglobal3d_fullbb) + "<br />"
			noteafter += "LIMIT GLOBAL 3DD FULLBB - " + strconv.Itoa(limitglobal3dd_fullbb) + "<br />"
			noteafter += "LIMIT GLOBAL 2D FULLBB - " + strconv.Itoa(limitglobal2d_fullbb) + "<br />"
			noteafter += "LIMIT GLOBAL 2DD FULLBB - " + strconv.Itoa(limitglobal2dd_fullbb) + "<br />"
			noteafter += "LIMIT GLOBAL 2DT FULLBB - " + strconv.Itoa(limitglobal2dt_fullbb) + "<br />"
			noteafter += "LIMIT TOTAL 4D FULLBB - " + strconv.Itoa(limittotal4d_fullbb) + "<br />"
			noteafter += "LIMIT TOTAL 3D FULLBB - " + strconv.Itoa(limittotal3d_fullbb) + "<br />"
			noteafter += "LIMIT TOTAL 3DD FULLBB - " + strconv.Itoa(limittotal3dd_fullbb) + "<br />"
			noteafter += "LIMIT TOTAL 2D FULLBB - " + strconv.Itoa(limittotal2d_fullbb) + "<br />"
			noteafter += "LIMIT TOTAL 2DD FULLBB - " + strconv.Itoa(limittotal2dd_fullbb) + "<br />"
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - 4-3-2", "", noteafter)
		} else {
			log.Println(msg_update)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Save_PasaranConfColokBebas(
	agent string,
	company string,
	idcomppasaran int,
	minbet, maxbet, maxbuy int,
	win, disc float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	render_page := time.Now()

	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		sql_update := `
			UPDATE 
			` + config.DB_tbl_mst_company_game_pasaran + `  
			SET 2_minbet=? , 2_maxbet=?, 2_maxbuy=?, 
			2_win=?, 2_disc=?, 
			2_limitbuang=?, 2_limitotal=?, 
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println("COLOK BEBAS")
		log.Printf("WIN : %s", fmt.Sprintf("%.3f", win))
		log.Printf("DISC : %s", fmt.Sprintf("%.3f", disc))
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_company_game_pasaran, "UPDATE",
			minbet, maxbet, maxbuy,
			fmt.Sprintf("%.3f", win), fmt.Sprintf("%.3f", disc),
			limitglobal, limittotal,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)

		if flag_update {
			msg = "Succes"
			log.Println(msg_update)

			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "MINIMAL BET - " + strconv.Itoa(minbet) + "<br />"
			noteafter += "MAXIMAL BET - " + strconv.Itoa(maxbet) + "<br />"
			noteafter += "MAXIMAL BUY - " + strconv.Itoa(maxbuy) + "<br />"
			noteafter += "WIN - " + fmt.Sprintf("%.3f", win) + "<br />"
			noteafter += "DISC - " + fmt.Sprintf("%.3f", disc) + "<br />"
			noteafter += "LIMIT GLOBAL  - " + strconv.Itoa(limitglobal) + "<br />"
			noteafter += "LIMIT TOTAL - " + strconv.Itoa(limittotal)
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - COLOK BEBAS", "", noteafter)
		} else {
			log.Println(msg_update)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Save_PasaranConfColokMacau(
	agent string,
	company string,
	idcomppasaran int,
	minbet, maxbet, maxbuy int,
	win2, win3, win4, disc float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	render_page := time.Now()

	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		sql_update := `
			UPDATE 
			` + config.DB_tbl_mst_company_game_pasaran + `  
			SET 3_minbet=? , 3_maxbet=?, 3_maxbuy=?,
			3_win2digit=?, 3_win3digit=?, 3_win4digit=?, 
			3_disc=?, 3_limitbuang=?, 3_limittotal=?,  
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println("COLOK MACAU")
		log.Printf("WIN 2 : %s", fmt.Sprintf("%.3f", win2))
		log.Printf("WIN 3 : %s", fmt.Sprintf("%.3f", win3))
		log.Printf("WIN 4 : %s", fmt.Sprintf("%.3f", win4))
		log.Printf("DISC : %s", fmt.Sprintf("%.3f", disc))
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_company_game_pasaran, "UPDATE",
			minbet, maxbet, maxbuy,
			fmt.Sprintf("%.3f", win2), fmt.Sprintf("%.3f", win3), fmt.Sprintf("%.3f", win4), fmt.Sprintf("%.3f", disc),
			limitglobal, limittotal,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)

		if flag_update {
			msg = "Succes"
			log.Println(msg_update)

			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "MINIMAL BET - " + strconv.Itoa(minbet) + "<br />"
			noteafter += "MAXIMAL BET - " + strconv.Itoa(maxbet) + "<br />"
			noteafter += "MAXIMAL BUY - " + strconv.Itoa(maxbuy) + "<br />"
			noteafter += "DISC - " + fmt.Sprintf("%.2f", disc) + "<br />"
			noteafter += "WIN2 - " + fmt.Sprintf("%.2f", win2) + "<br />"
			noteafter += "WIN3 - " + fmt.Sprintf("%.2f", win3) + "<br />"
			noteafter += "WIN4 - " + fmt.Sprintf("%.2f", win4) + "<br />"
			noteafter += "LIMIT GLOBAL  - " + strconv.Itoa(limitglobal) + "<br />"
			noteafter += "LIMIT TOTAL - " + strconv.Itoa(limittotal)
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - COLOK MACAU", "", noteafter)
		} else {
			log.Println(msg_update)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Save_PasaranConfColokNaga(
	agent string,
	company string,
	idcomppasaran int,
	minbet, maxbet, maxbuy int,
	win3, win4, disc float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	render_page := time.Now()

	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		sql_update := `
			UPDATE 
			` + config.DB_tbl_mst_company_game_pasaran + `  
			SET 4_minbet=? , 4_maxbet=?, 4_maxbuy=?,  
			4_win3digit=?, 4_win4digit=?,  
			4_disc=?, 4_limitbuang=?, 4_limittotal=?,  
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println("COLOK NAGA")
		log.Printf("WIN 3 : %s", fmt.Sprintf("%.3f", win3))
		log.Printf("WIN 4 : %s", fmt.Sprintf("%.3f", win4))
		log.Printf("DISC : %s", fmt.Sprintf("%.3f", disc))
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_company_game_pasaran, "UPDATE",
			minbet, maxbet, maxbuy,
			fmt.Sprintf("%.3f", win3), fmt.Sprintf("%.3f", win4), fmt.Sprintf("%.3f", disc),
			limitglobal, limittotal,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)

		if flag_update {
			msg = "Succes"
			log.Println(msg_update)

			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "MINIMAL BET - " + strconv.Itoa(minbet) + "<br />"
			noteafter += "MAXIMAL BET - " + strconv.Itoa(maxbet) + "<br />"
			noteafter += "MAXIMAL BUY - " + strconv.Itoa(maxbuy) + "<br />"
			noteafter += "DISC - " + fmt.Sprintf("%.2f", disc) + "<br />"
			noteafter += "WIN3 - " + fmt.Sprintf("%.2f", win3) + "<br />"
			noteafter += "WIN4 - " + fmt.Sprintf("%.2f", win4) + "<br />"
			noteafter += "LIMIT GLOBAL  - " + strconv.Itoa(limitglobal) + "<br />"
			noteafter += "LIMIT TOTAL - " + strconv.Itoa(limittotal)
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - COLOK NAGA", "", noteafter)
		} else {
			log.Println(msg_update)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Save_PasaranConfColokJitu(
	agent string,
	company string,
	idcomppasaran int,
	minbet, maxbet, maxbuy int,
	winas, winkop, winkepala, winekor, disc float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	render_page := time.Now()

	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		sql_update := `
			UPDATE 
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 5_minbet=? , 5_maxbet=?, 5_maxbuy=?, 
			5_winas=?, 5_winkop=?, 5_winkepala=?, 5_winekor=?,
			5_desic=?, 5_limitbuang=?, 5_limitotal=?,  
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println("COLOK JITU")
		log.Printf("WIN AS : %s", fmt.Sprintf("%.3f", winas))
		log.Printf("WIN KOP : %s", fmt.Sprintf("%.3f", winkop))
		log.Printf("WIN KEPALA : %s", fmt.Sprintf("%.3f", winkepala))
		log.Printf("WIN EKOR : %s", fmt.Sprintf("%.3f", winekor))
		log.Printf("DISC : %s", fmt.Sprintf("%.3f", disc))
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_company_game_pasaran, "UPDATE",
			minbet, maxbet, maxbuy,
			fmt.Sprintf("%.3f", winas), fmt.Sprintf("%.3f", winkop), fmt.Sprintf("%.3f", winkepala), fmt.Sprintf("%.3f", winekor),
			disc,
			limitglobal, limittotal,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)

		if flag_update {
			msg = "Succes"
			log.Println(msg_update)

			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "MINIMAL BET - " + strconv.Itoa(minbet) + "<br />"
			noteafter += "MAXIMAL BET - " + strconv.Itoa(maxbet) + "<br />"
			noteafter += "MAXIMAL BUY - " + strconv.Itoa(maxbuy) + "<br />"
			noteafter += "DISC - " + fmt.Sprintf("%.2f", disc) + "<br />"
			noteafter += "WIN_AS - " + fmt.Sprintf("%.2f", winas) + "<br />"
			noteafter += "WIN_KOP - " + fmt.Sprintf("%.2f", winkop) + "<br />"
			noteafter += "WIN_KEPALA - " + fmt.Sprintf("%.2f", winkepala) + "<br />"
			noteafter += "WIN_EKOR - " + fmt.Sprintf("%.2f", winekor) + "<br />"
			noteafter += "LIMIT GLOBAL  - " + strconv.Itoa(limitglobal) + "<br />"
			noteafter += "LIMIT TOTAL - " + strconv.Itoa(limittotal)
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - COLOK JITU", "", noteafter)
		} else {
			log.Println(msg_update)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Save_PasaranConf5050Umum(
	agent string,
	company string,
	idcomppasaran int,
	minbet, maxbet, maxbuy int,
	keibesar, keikecil, keigenap, keiganjil, keitengah, keitepi float32,
	discbesar, disckecil, discgenap, discganjil, disctengah, disctepi float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	render_page := time.Now()

	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		sql_update := `
			UPDATE 
			` + config.DB_tbl_mst_company_game_pasaran + `  
			SET 6_minbet=? , 6_maxbet=?, 6_maxbuy=?,
			6_keibesar=?, 6_keikecil=?, 6_keigenap=?, 6_keiganjil=?, 6_keitengah=?, 6_keitepi=?, 
			6_discbesar=?, 6_disckecil=?, 6_discgenap=?, 6_discganjil=?, 6_disctengah=?, 6_disctepi=?,  
			6_limitbuang=?, 6_limittotal=?,  
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println("5050 UMUM")
		log.Printf("KEI BESAR : %s", fmt.Sprintf("%.3f", keibesar))
		log.Printf("KEI KECIL : %s", fmt.Sprintf("%.3f", keikecil))
		log.Printf("KEI GENAP : %s", fmt.Sprintf("%.3f", keigenap))
		log.Printf("KEI GANJIL : %s", fmt.Sprintf("%.3f", keiganjil))
		log.Printf("KEI TENGAH : %s", fmt.Sprintf("%.3f", keitengah))
		log.Printf("KEI TEPI : %s", fmt.Sprintf("%.3f", keitepi))
		log.Printf("DISC BESAR : %s", fmt.Sprintf("%.3f", discbesar))
		log.Printf("DISC KECIL : %s", fmt.Sprintf("%.3f", disckecil))
		log.Printf("DISC GENAP : %s", fmt.Sprintf("%.3f", discgenap))
		log.Printf("DISC GANJIL : %s", fmt.Sprintf("%.3f", discganjil))
		log.Printf("DISC TENGAH : %s", fmt.Sprintf("%.3f", disctengah))
		log.Printf("DISC TEPI : %s", fmt.Sprintf("%.3f", disctepi))
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_company_game_pasaran, "UPDATE",
			minbet, maxbet, maxbuy,
			fmt.Sprintf("%.3f", keibesar), fmt.Sprintf("%.3f", keikecil), fmt.Sprintf("%.3f", keigenap), fmt.Sprintf("%.3f", keiganjil),
			fmt.Sprintf("%.3f", keitengah), fmt.Sprintf("%.3f", keitepi),
			fmt.Sprintf("%.3f", discbesar), fmt.Sprintf("%.3f", disckecil), fmt.Sprintf("%.3f", discgenap), fmt.Sprintf("%.3f", discganjil),
			fmt.Sprintf("%.3f", disctengah), fmt.Sprintf("%.3f", disctepi),
			limitglobal, limittotal,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)

		if flag_update {
			msg = "Success"
			log.Println(msg_update)

			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "\n"
			noteafter += "MINIMAL BET - " + strconv.Itoa(minbet) + "\n"
			noteafter += "MAXIMAL BET - " + strconv.Itoa(maxbet) + "\n"
			noteafter += "MAXIMAL BUY - " + strconv.Itoa(maxbuy) + "\n"
			noteafter += "DISC BESAR - " + fmt.Sprintf("%.3f", discbesar) + "\n"
			noteafter += "DISC KECIL - " + fmt.Sprintf("%.3f", disckecil) + "\n"
			noteafter += "DISC GENAP - " + fmt.Sprintf("%.3f", discgenap) + "\n"
			noteafter += "DISC GANJIL - " + fmt.Sprintf("%.3f", discganjil) + "\n"
			noteafter += "DISC TENGAH - " + fmt.Sprintf("%.3f", disctengah) + "\n"
			noteafter += "DISC TEPI - " + fmt.Sprintf("%.3f", disctepi) + "\n"
			noteafter += "KEI BESAR - " + fmt.Sprintf("%.3f", keibesar) + "\n"
			noteafter += "KEI KECIL - " + fmt.Sprintf("%.3f", keikecil) + "\n"
			noteafter += "KEI GENAP - " + fmt.Sprintf("%.3f", keigenap) + "\n"
			noteafter += "KEI GANJIL - " + fmt.Sprintf("%.3f", keiganjil) + "\n"
			noteafter += "KEI TENGAH - " + fmt.Sprintf("%.3f", keitengah) + "\n"
			noteafter += "KEI TEPI - " + fmt.Sprintf("%f", keitepi) + "\n"
			noteafter += "LIMIT GLOBAL  - " + strconv.Itoa(limitglobal) + "\n"
			noteafter += "LIMIT TOTAL - " + strconv.Itoa(limittotal)
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - 5050 UMUM", "", noteafter)
		} else {
			log.Println(msg_update)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Save_PasaranConf5050Special(
	agent string,
	company string,
	idcomppasaran int,
	minbet, maxbet, maxbuy int,
	keiasganjil, keiasgenap, keiasbesar, keiaskecil float32,
	keikopganjil, keikopgenap, keikopbesar, keikopkecil float32,
	keikepalaganjil, keikepalagenap, keikepalabesar, keikepalakecil float32,
	keiekorganjil, keiekorgenap, keiekorbesar, keiekorkecil float32,
	discasganjil, discasgenap, discasbesar, discaskecil float32,
	disckopganjil, disckopgenap, disckopbesar, disckopkecil float32,
	disckepalaganjil, disckepalagenap, disckepalabesar, disckepalakecil float32,
	discekorganjil, discekorgenap, discekorbesar, discekorkecil float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	render_page := time.Now()

	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		sql_update := `
			UPDATE 
			` + config.DB_tbl_mst_company_game_pasaran + ` 
			SET 7_minbet=? , 7_maxbet=?, 7_maxbuy=?,
			7_keiasganjil=?, 7_keiasgenap=?, 7_keiasbesar=?, 7_keiaskecil=?, 
			7_keikopganjil=?, 7_keikopgenap=?, 7_keikopbesar=?, 7_keikopkecil=?, 
			7_keikepalaganjil=?, 7_keikepalagenap=?, 7_keikepalabesar=?, 7_keikepalakecil=?,  
			7_keiekorganjil=?, 7_keiekorgenap=?, 7_keiekorbesar=?, 7_keiekorkecil=?, 
			7_discasganjil=?, 7_discasgenap=?, 7_discasbesar=?, 7_discaskecil=?, 
			7_disckopganjil=?, 7_disckopgenap=?, 7_disckopbesar=?, 7_disckopkecil=?, 
			7_disckepalaganjil=?, 7_disckepalagenap=?, 7_disckepalabesar=?, 7_disckepalakecil=?, 
			7_discekorganjil=?, 7_discekorgenap=?, 7_discekorbesar=?, 7_discekorkecil=?, 
			7_limitbuang=?, 7_limittotal=?,  
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println("5050 SPECIAL")
		log.Printf("KEI AS GANJIL : %s", fmt.Sprintf("%.3f", keiasganjil))
		log.Printf("KEI AS GENAP : %s", fmt.Sprintf("%.3f", keiasgenap))
		log.Printf("KEI AS BESAR : %s", fmt.Sprintf("%.3f", keiasbesar))
		log.Printf("KEI AS KECIL : %s", fmt.Sprintf("%.3f", keiaskecil))
		log.Printf("KEI KOP GANJIL : %s", fmt.Sprintf("%.3f", keikopganjil))
		log.Printf("KEI KOP GENAP : %s", fmt.Sprintf("%.3f", keikopgenap))
		log.Printf("KEI KOP BESAR : %s", fmt.Sprintf("%.3f", keikopbesar))
		log.Printf("KEI KOP KECIL : %s", fmt.Sprintf("%.3f", keikopkecil))
		log.Printf("KEI KEPALA GANJIL : %s", fmt.Sprintf("%.3f", keikepalaganjil))
		log.Printf("KEI KEPALA GENAP : %s", fmt.Sprintf("%.3f", keikepalagenap))
		log.Printf("KEI KEPALA BESAR : %s", fmt.Sprintf("%.3f", keikepalabesar))
		log.Printf("KEI KEPALA KECIL : %s", fmt.Sprintf("%.3f", keikepalakecil))
		log.Printf("KEI EKOR GANJIL : %s", fmt.Sprintf("%.3f", keiekorganjil))
		log.Printf("KEI EKOR GENAP : %s", fmt.Sprintf("%.3f", keiekorgenap))
		log.Printf("KEI EKOR BESAR : %s", fmt.Sprintf("%.3f", keiekorbesar))
		log.Printf("KEI EKOR KECIL : %s", fmt.Sprintf("%.3f", keiekorkecil))

		log.Printf("DISC AS GANJIL : %s", fmt.Sprintf("%.3f", discasganjil))
		log.Printf("DISC AS GENAP : %s", fmt.Sprintf("%.3f", discasgenap))
		log.Printf("DISC AS BESAR : %s", fmt.Sprintf("%.3f", discasbesar))
		log.Printf("DISC AS KECIL : %s", fmt.Sprintf("%.3f", discaskecil))
		log.Printf("DISC KOP GANJIL : %s", fmt.Sprintf("%.3f", disckopganjil))
		log.Printf("DISC KOP GENAP : %s", fmt.Sprintf("%.3f", disckopgenap))
		log.Printf("DISC KOP BESAR : %s", fmt.Sprintf("%.3f", disckopbesar))
		log.Printf("DISC KOP KECIL : %s", fmt.Sprintf("%.3f", disckopkecil))
		log.Printf("DISC KEPALA GANJIL : %s", fmt.Sprintf("%.3f", disckepalaganjil))
		log.Printf("DISC KEPALA GENAP : %s", fmt.Sprintf("%.3f", disckepalagenap))
		log.Printf("DISC KEPALA BESAR : %s", fmt.Sprintf("%.3f", disckepalabesar))
		log.Printf("DISC KEPALA KECIL : %s", fmt.Sprintf("%.3f", disckepalakecil))
		log.Printf("DISC EKOR GANJIL : %s", fmt.Sprintf("%.3f", discekorganjil))
		log.Printf("DISC EKOR GENAP : %s", fmt.Sprintf("%.3f", discekorgenap))
		log.Printf("DISC EKOR BESAR : %s", fmt.Sprintf("%.3f", discekorbesar))
		log.Printf("DISC EKOR KECIL : %s", fmt.Sprintf("%.3f", discekorkecil))
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_company_game_pasaran, "UPDATE",
			minbet, maxbet, maxbuy,
			fmt.Sprintf("%.3f", keiasganjil), fmt.Sprintf("%.3f", keiasgenap), fmt.Sprintf("%.3f", keiasbesar), fmt.Sprintf("%.3f", keiaskecil),
			fmt.Sprintf("%.3f", keikopganjil), fmt.Sprintf("%.3f", keikopgenap), fmt.Sprintf("%.3f", keikopbesar), fmt.Sprintf("%.3f", keikopkecil),
			fmt.Sprintf("%.3f", keikepalaganjil), fmt.Sprintf("%.3f", keikepalagenap), fmt.Sprintf("%.3f", keikepalabesar), fmt.Sprintf("%.3f", keikepalakecil),
			fmt.Sprintf("%.3f", keiekorganjil), fmt.Sprintf("%.3f", keiekorgenap), fmt.Sprintf("%.3f", keiekorbesar), fmt.Sprintf("%.3f", keiekorkecil),
			discasganjil, discasgenap, discasbesar, discaskecil,
			disckopganjil, disckopgenap, disckopbesar, disckopkecil,
			disckepalaganjil, disckepalagenap, disckepalabesar, disckepalakecil,
			discekorganjil, discekorgenap, discekorbesar, discekorkecil,
			limitglobal,
			limittotal,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)

		if flag_update {
			msg = "Success"
			log.Println(msg_update)

			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "MINIMAL BET - " + strconv.Itoa(minbet) + "<br />"
			noteafter += "MAXIMAL BET - " + strconv.Itoa(maxbet) + "<br />"
			noteafter += "MAXIMAL BUY - " + strconv.Itoa(maxbuy) + "<br />"
			noteafter += "DISC AS GANJIL - " + fmt.Sprintf("%.3f", discasganjil) + "<br />"
			noteafter += "DISC AS GENAP - " + fmt.Sprintf("%.3f", discasgenap) + "<br />"
			noteafter += "DISC AS BESAR - " + fmt.Sprintf("%.3f", discasbesar) + "<br />"
			noteafter += "DISC AS KECIL - " + fmt.Sprintf("%.3f", discaskecil) + "<br />"
			noteafter += "DISC KOP GANJIL - " + fmt.Sprintf("%.3f", disckopganjil) + "<br />"
			noteafter += "DISC KOP GENAP - " + fmt.Sprintf("%.3f", disckopgenap) + "<br />"
			noteafter += "DISC KOP BESAR - " + fmt.Sprintf("%.3f", disckopbesar) + "<br />"
			noteafter += "DISC KOP KECIL - " + fmt.Sprintf("%.3f", disckopkecil) + "<br />"
			noteafter += "DISC KEPALA GANJIL - " + fmt.Sprintf("%.3f", disckepalaganjil) + "<br />"
			noteafter += "DISC KEPALA GENAP - " + fmt.Sprintf("%.3f", disckepalagenap) + "<br />"
			noteafter += "DISC KEPALA BESAR - " + fmt.Sprintf("%.3f", disckepalabesar) + "<br />"
			noteafter += "DISC KEPALA KECIL - " + fmt.Sprintf("%.3f", disckepalakecil) + "<br />"
			noteafter += "DISC EKOR GANJIL - " + fmt.Sprintf("%.3f", discekorganjil) + "<br />"
			noteafter += "DISC EKOR GENAP - " + fmt.Sprintf("%.3f", discekorgenap) + "<br />"
			noteafter += "DISC EKOR BESAR - " + fmt.Sprintf("%.3f", discekorbesar) + "<br />"
			noteafter += "DISC EKOR KECIL - " + fmt.Sprintf("%.3f", discekorkecil) + "<br />"
			noteafter += "KEI AS GANJIL - " + fmt.Sprintf("%.3f", keiasganjil) + "<br />"
			noteafter += "KEI AS GENAP - " + fmt.Sprintf("%.3f", keiasgenap) + "<br />"
			noteafter += "KEI AS BESAR - " + fmt.Sprintf("%.3f", keiasbesar) + "<br />"
			noteafter += "KEI AS KECIL - " + fmt.Sprintf("%.3f", keiaskecil) + "<br />"
			noteafter += "KEI KOP GANJIL - " + fmt.Sprintf("%.3f", keikopganjil) + "<br />"
			noteafter += "KEI KOP GENAP - " + fmt.Sprintf("%.3f", keikopgenap) + "<br />"
			noteafter += "KEI KOP BESAR - " + fmt.Sprintf("%.3f", keikopbesar) + "<br />"
			noteafter += "KEI KOP KECIL - " + fmt.Sprintf("%.3f", keikopkecil) + "<br />"
			noteafter += "KEI KEPALA GANJIL - " + fmt.Sprintf("%.3f", keikepalaganjil) + "<br />"
			noteafter += "KEI KEPALA GENAP - " + fmt.Sprintf("%.3f", keikepalagenap) + "<br />"
			noteafter += "KEI KEPALA BESAR - " + fmt.Sprintf("%.3f", keikepalabesar) + "<br />"
			noteafter += "KEI KEPALA KECIL - " + fmt.Sprintf("%.3f", keikepalakecil) + "<br />"
			noteafter += "KEI EKOR GANJIL - " + fmt.Sprintf("%.3f", keiekorganjil) + "<br />"
			noteafter += "KEI EKOR GENAP - " + fmt.Sprintf("%.3f", keiekorgenap) + "<br />"
			noteafter += "KEI EKOR BESAR - " + fmt.Sprintf("%.3f", keiekorbesar) + "<br />"
			noteafter += "KEI EKOR KECIL - " + fmt.Sprintf("%.3f", keiekorkecil) + "<br />"
			noteafter += "LIMIT GLOBAL  - " + strconv.Itoa(limitglobal) + "<br />"
			noteafter += "LIMIT TOTAL - " + strconv.Itoa(limittotal)
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - 5050 SPECIAL", "", noteafter)
		} else {
			log.Println(msg_update)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Save_PasaranConf5050Kombinasi(
	agent string,
	company string,
	idcomppasaran int,
	minbet, maxbet, maxbuy int,
	belakangkeimono, belakangkeistereo, belakangkeikembang, belakangkeikempis, belakangkeikembar float32,
	tengahkeimono, tengahkeistereo, tengahkeikembang, tengahkeikempis, tengahkeikembar float32,
	depankeimono, depankeistereo, depankeikembang, depankeikempis, depankeikembar float32,
	belakangdiscmono, belakangdiscstereo, belakangdisckembang, belakangdisckempis, belakangdisckembar float32,
	tengahdiscmono, tengahdiscstereo, tengahdisckembang, tengahdisckempis, tengahdisckembar float32,
	depandiscmono, depandiscstereo, depandisckembang, depandisckempis, depandisckembar float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	render_page := time.Now()

	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		sql_update := `
			UPDATE 
			` + config.DB_tbl_mst_company_game_pasaran + `  
			SET 8_minbet=? , 8_maxbet=?, 8_maxbuy=?,
			8_belakangkeimono=?, 8_belakangkeistereo=?, 8_belakangkeikembang=?, 8_belakangkeikempis=?, 8_belakangkeikembar=?, 
			8_tengahkeimono=?, 8_tengahkeistereo=?, 8_tengahkeikembang=?, 8_tengahkeikempis=?, 8_tengahkeikembar=?, 
			8_depankeimono=?, 8_depankeistereo=?, 8_depankeikembang=?, 8_depankeikempis=?, 8_depankeikembar=?, 
			8_belakangdiscmono=?, 8_belakangdiscstereo=?, 8_belakangdisckembang=?, 8_belakangdisckempis=?, 8_belakangdisckembar=?, 
			8_tengahdiscmono=?, 8_tengahdiscstereo=?, 8_tengahdisckembang=?, 8_tengahdisckempis=?, 8_tengahdisckembar=?, 
			8_depandiscmono=?, 8_depandiscstereo=?, 8_depandisckembang=?, 8_depandisckempis=?, 8_depandisckembar=?, 
			8_limitbuang=?, 8_limittotal=?,  
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println("5050 KOMBINASI")
		log.Printf("BELAKANG KEI MONO : %s", fmt.Sprintf("%.3f", belakangkeimono))
		log.Printf("BELAKANG KEI STEREO : %s", fmt.Sprintf("%.3f", belakangkeistereo))
		log.Printf("BELAKANG KEI KEMBANG : %s", fmt.Sprintf("%.3f", belakangkeikembang))
		log.Printf("BELAKANG KEI KEMPIS : %s", fmt.Sprintf("%.3f", belakangkeikempis))
		log.Printf("BELAKANG KEI KEMBAR : %s", fmt.Sprintf("%.3f", belakangkeikembar))
		log.Printf("TENGAH KEI MONO : %s", fmt.Sprintf("%.3f", tengahkeimono))
		log.Printf("TENGAH KEI STEREO : %s", fmt.Sprintf("%.3f", tengahkeistereo))
		log.Printf("TENGAH KEI KEMBANG : %s", fmt.Sprintf("%.3f", tengahkeikembang))
		log.Printf("TENGAH KEI KEMPIS : %s", fmt.Sprintf("%.3f", tengahkeikempis))
		log.Printf("TENGAH KEI KEMBAR : %s", fmt.Sprintf("%.3f", tengahkeikembar))
		log.Printf("DEPAN KEI MONO : %s", fmt.Sprintf("%.3f", depankeimono))
		log.Printf("DEPAN KEI STEREO : %s", fmt.Sprintf("%.3f", depankeistereo))
		log.Printf("DEPAN KEI KEMBANG : %s", fmt.Sprintf("%.3f", depankeikembang))
		log.Printf("DEPAN KEI KEMPIS : %s", fmt.Sprintf("%.3f", depankeikempis))
		log.Printf("DEPAN KEI KEMBAR : %s", fmt.Sprintf("%.3f", depankeikembar))

		log.Printf("BELAKANG DISC MONO : %s", fmt.Sprintf("%.3f", belakangdiscmono))
		log.Printf("BELAKANG DISC STEREO : %s", fmt.Sprintf("%.3f", belakangdiscstereo))
		log.Printf("BELAKANG DISC KEMBANG : %s", fmt.Sprintf("%.3f", belakangdisckembang))
		log.Printf("BELAKANG DISC KEMPIS : %s", fmt.Sprintf("%.3f", belakangdisckempis))
		log.Printf("BELAKANG DISC KEMBAR : %s", fmt.Sprintf("%.3f", belakangdisckembar))
		log.Printf("TENGAH DISC MONO : %s", fmt.Sprintf("%.3f", tengahdiscmono))
		log.Printf("TENGAH DISC STEREO : %s", fmt.Sprintf("%.3f", tengahdiscstereo))
		log.Printf("TENGAH DISC KEMBANG : %s", fmt.Sprintf("%.3f", tengahdisckembang))
		log.Printf("TENGAH DISC KEMPIS : %s", fmt.Sprintf("%.3f", tengahdisckempis))
		log.Printf("TENGAH DISC KEMBAR : %s", fmt.Sprintf("%.3f", tengahdisckembar))
		log.Printf("DEPAN DISC MONO : %s", fmt.Sprintf("%.3f", depandiscmono))
		log.Printf("DEPAN DISC STEREO : %s", fmt.Sprintf("%.3f", depandiscstereo))
		log.Printf("DEPAN DISC KEMBANG : %s", fmt.Sprintf("%.3f", depandisckembang))
		log.Printf("DEPAN DISC KEMPIS : %s", fmt.Sprintf("%.3f", depandisckempis))
		log.Printf("DEPAN DISC KEMBAR : %s", fmt.Sprintf("%.3f", depandisckembar))

		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_company_game_pasaran, "UPDATE",
			minbet, maxbet, maxbuy,
			fmt.Sprintf("%.3f", belakangkeimono), fmt.Sprintf("%.3f", belakangkeistereo), fmt.Sprintf("%.3f", belakangkeikembang),
			fmt.Sprintf("%.3f", belakangkeikempis), fmt.Sprintf("%.3f", belakangkeikembar),
			fmt.Sprintf("%.3f", tengahkeimono), fmt.Sprintf("%.3f", tengahkeistereo), fmt.Sprintf("%.3f", tengahkeikembang),
			fmt.Sprintf("%.3f", tengahkeikempis), fmt.Sprintf("%.3f", tengahkeikembar),
			fmt.Sprintf("%.3f", depankeimono), fmt.Sprintf("%.3f", depankeistereo), fmt.Sprintf("%.3f", depankeikembang),
			fmt.Sprintf("%.3f", depankeikempis), fmt.Sprintf("%.3f", depankeikembar),
			fmt.Sprintf("%.3f", belakangdiscmono), fmt.Sprintf("%.3f", belakangdiscstereo), fmt.Sprintf("%.3f", belakangdisckembang),
			fmt.Sprintf("%.3f", belakangdisckempis), fmt.Sprintf("%.3f", belakangdisckembar),
			fmt.Sprintf("%.3f", tengahdiscmono), fmt.Sprintf("%.3f", tengahdiscstereo), fmt.Sprintf("%.3f", tengahdisckembang),
			fmt.Sprintf("%.3f", tengahdisckempis), fmt.Sprintf("%.3f", tengahdisckembar),
			fmt.Sprintf("%.3f", depandiscmono), fmt.Sprintf("%.3f", depandiscstereo), fmt.Sprintf("%.3f", depandisckembang),
			fmt.Sprintf("%.3f", depandisckempis), fmt.Sprintf("%.3f", depandisckembar),
			limitglobal,
			limittotal,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)

		if flag_update {
			msg = "Success"
			log.Println(msg_update)

			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "MINIMAL BET - " + strconv.Itoa(minbet) + "<br />"
			noteafter += "MAXIMAL BET - " + strconv.Itoa(maxbet) + "<br />"
			noteafter += "MAXIMAL BUY - " + strconv.Itoa(maxbuy) + "<br />"
			noteafter += "DEPAN KEI MONO - " + fmt.Sprintf("%.3f", depankeimono) + "<br />"
			noteafter += "DEPAN KEI STEREO - " + fmt.Sprintf("%.3f", depankeistereo) + "<br />"
			noteafter += "DEPAN KEI KEMBANG - " + fmt.Sprintf("%.3f", depankeikembang) + "<br />"
			noteafter += "DEPAN KEI KEMPIS - " + fmt.Sprintf("%.3f", depankeikempis) + "<br />"
			noteafter += "DEPAN KEI KEMBAR - " + fmt.Sprintf("%.3f", depankeikembar) + "<br />"
			noteafter += "TENGAH KEI MONO - " + fmt.Sprintf("%.3f", tengahkeimono) + "<br />"
			noteafter += "TENGAH KEI STEREO - " + fmt.Sprintf("%.3f", tengahkeistereo) + "<br />"
			noteafter += "TENGAH KEI KEMBANG - " + fmt.Sprintf("%.3f", tengahkeikembang) + "<br />"
			noteafter += "TENGAH KEI KEMPIS - " + fmt.Sprintf("%.3f", tengahkeikempis) + "<br />"
			noteafter += "TENGAH KEI KEMBAR - " + fmt.Sprintf("%.3f", tengahkeikembar) + "<br />"
			noteafter += "BELAKANG KEI MONO - " + fmt.Sprintf("%.3f", belakangkeimono) + "<br />"
			noteafter += "BELAKANG KEI STEREO - " + fmt.Sprintf("%.3f", belakangkeistereo) + "<br />"
			noteafter += "BELAKANG KEI KEMBANG - " + fmt.Sprintf("%.3f", belakangkeikembang) + "<br />"
			noteafter += "BELAKANG KEI KEMPIS - " + fmt.Sprintf("%.3f", belakangkeikempis) + "<br />"
			noteafter += "BELAKANG KEI KEMBAR - " + fmt.Sprintf("%.3f", belakangkeikembar) + "<br />"
			noteafter += "DEPAN DISC MONO - " + fmt.Sprintf("%.3f", depandiscmono) + "<br />"
			noteafter += "DEPAN DISC STEREO - " + fmt.Sprintf("%.3f", depandiscstereo) + "<br />"
			noteafter += "DEPAN DISC KEMBANG - " + fmt.Sprintf("%.3f", depandisckembang) + "<br />"
			noteafter += "DEPAN DISC KEMPIS - " + fmt.Sprintf("%.3f", depandisckempis) + "<br />"
			noteafter += "DEPAN DISC KEMBAR - " + fmt.Sprintf("%.3f", depandisckembar) + "<br />"
			noteafter += "TENGAH DISC MONO - " + fmt.Sprintf("%.3f", tengahdiscmono) + "<br />"
			noteafter += "TENGAH DISC STEREO - " + fmt.Sprintf("%.3f", tengahdiscstereo) + "<br />"
			noteafter += "TENGAH DISC KEMBANG - " + fmt.Sprintf("%.3f", tengahdisckembang) + "<br />"
			noteafter += "TENGAH DISC KEMPIS - " + fmt.Sprintf("%.3f", tengahdisckempis) + "<br />"
			noteafter += "TENGAH DISC KEMBAR - " + fmt.Sprintf("%.3f", tengahdisckembar) + "<br />"
			noteafter += "BELAKANG DISC MONO - " + fmt.Sprintf("%.3f", belakangdiscmono) + "<br />"
			noteafter += "BELAKANG DISC STEREO - " + fmt.Sprintf("%.3f", belakangdiscstereo) + "<br />"
			noteafter += "BELAKANG DISC KEMBANG - " + fmt.Sprintf("%.3f", belakangdisckembang) + "<br />"
			noteafter += "BELAKANG DISC KEMPIS - " + fmt.Sprintf("%.3f", belakangdisckempis) + "<br />"
			noteafter += "BELAKANG DISC KEMBAR - " + fmt.Sprintf("%.3f", belakangdisckembar) + "<br />"
			noteafter += "LIMIT GLOBAL  - " + strconv.Itoa(limitglobal) + "<br />"
			noteafter += "LIMIT TOTAL - " + strconv.Itoa(limittotal)
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - 5050 KOMBINASI", "", noteafter)
		} else {
			log.Println(msg_update)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Save_PasaranConfMacauKombinasi(
	agent string,
	company string,
	idcomppasaran int,
	minbet, maxbet, maxbuy int,
	win, disc float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	render_page := time.Now()

	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		sql_update := `
			UPDATE 
			` + config.DB_tbl_mst_company_game_pasaran + `  
			SET 9_minbet=? , 9_maxbet=?, 9_maxbuy=?,
			9_win=?, 9_discount=?, 
			9_limitbuang=?, 9_limittotal=?, 
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println("MACAU KOMBINASI")
		log.Printf("WIN : %s", fmt.Sprintf("%.3f", win))
		log.Printf("DISC : %s", fmt.Sprintf("%.3f", disc))
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_company_game_pasaran, "UPDATE",
			minbet, maxbet, maxbuy,
			fmt.Sprintf("%.3f", win), fmt.Sprintf("%.3f", disc),
			limitglobal, limittotal,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)

		if flag_update {
			msg = "Success"
			log.Println(msg_update)

			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "MINIMAL BET - " + strconv.Itoa(minbet) + "<br />"
			noteafter += "MAXIMAL BET - " + strconv.Itoa(maxbet) + "<br />"
			noteafter += "MAXIMAL BUY - " + strconv.Itoa(maxbuy) + "<br />"
			noteafter += "WIN - " + fmt.Sprintf("%.3f", win) + "<br />"
			noteafter += "DISC - " + fmt.Sprintf("%.3f", disc) + "<br />"
			noteafter += "LIMIT GLOBAL  - " + strconv.Itoa(limitglobal) + "<br />"
			noteafter += "LIMIT TOTAL - " + strconv.Itoa(limittotal)
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - MACAU", "", noteafter)
		} else {
			log.Println(msg_update)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Save_PasaranConfDasar(
	agent string,
	company string,
	idcomppasaran int,
	minbet, maxbet, maxbuy int,
	keibesar, keikecil, keigenap, keiganjil float32,
	discbesar, disckecil, discigenap, discganjil float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	render_page := time.Now()

	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		sql_update := `
			UPDATE 
			` + config.DB_tbl_mst_company_game_pasaran + `  
			SET 10_minbet=? , 10_maxbet=?, 10_maxbuy=?,
			10_keibesar=?, 10_keikecil=?, 10_keigenap=?, 10_keiganjil=?, 
			10_discbesar=?, 10_disckecil=?, 10_discigenap=?, 10_discganjil=?, 
			10_limitbuang=?, 10_limittotal=?, 
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println("DASAR")
		log.Printf("KEI BESAR : %s", fmt.Sprintf("%.3f", keibesar))
		log.Printf("KEI KECIL : %s", fmt.Sprintf("%.3f", keikecil))
		log.Printf("KEI GENAP : %s", fmt.Sprintf("%.3f", keigenap))
		log.Printf("KEI GANJIL : %s", fmt.Sprintf("%.3f", keiganjil))
		log.Printf("DISC BESAR : %s", fmt.Sprintf("%.3f", discbesar))
		log.Printf("DISC KECIL : %s", fmt.Sprintf("%.3f", disckecil))
		log.Printf("DISC GENAP : %s", fmt.Sprintf("%.3f", discigenap))
		log.Printf("DISC GANJIL : %s", fmt.Sprintf("%.3f", discganjil))
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_company_game_pasaran, "UPDATE",
			minbet, maxbet, maxbuy,
			fmt.Sprintf("%.3f", keibesar), fmt.Sprintf("%.3f", keikecil), fmt.Sprintf("%.3f", keigenap), fmt.Sprintf("%.3f", keiganjil),
			fmt.Sprintf("%.3f", discbesar), fmt.Sprintf("%.3f", disckecil), fmt.Sprintf("%.3f", discigenap), fmt.Sprintf("%.3f", discganjil),
			limitglobal,
			limittotal,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)

		if flag_update {
			msg = "Success"
			log.Println(msg_update)

			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "MINIMAL BET - " + strconv.Itoa(minbet) + "<br />"
			noteafter += "MAXIMAL BET - " + strconv.Itoa(maxbet) + "<br />"
			noteafter += "MAXIMAL BUY - " + strconv.Itoa(maxbuy) + "<br />"
			noteafter += "DISC BESAR - " + fmt.Sprintf("%.3f", discbesar) + "<br />"
			noteafter += "DISC KECIL - " + fmt.Sprintf("%.3f", disckecil) + "<br />"
			noteafter += "DISC GENAP - " + fmt.Sprintf("%.3f", discigenap) + "<br />"
			noteafter += "DISC GANJIL - " + fmt.Sprintf("%.3f", discganjil) + "<br />"
			noteafter += "LIMIT GLOBAL  - " + strconv.Itoa(limitglobal) + "<br />"
			noteafter += "LIMIT TOTAL - " + strconv.Itoa(limittotal)
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - DASAR", "", noteafter)
		} else {
			log.Println(msg_update)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Save_PasaranConfShio(
	agent string,
	company string,
	idcomppasaran int,
	shiotahunini string,
	minbet, maxbet, maxbuy int,
	win, disc float32,
	limitglobal, limittotal int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	render_page := time.Now()

	msg := "Failed"
	pasarancode, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
	tipepasaran := Pasaranmaster_id(pasarancode, "tipepasaran")
	if tipepasaran != "WAJIB" {
		sql_update := `
			UPDATE 
			` + config.DB_tbl_mst_company_game_pasaran + `  
			SET 11_shiotahunini=? , 11_minbet=?, 11_maxbet=?, 11_maxbuy=?, 
			11_win=?, 11_disc=?, 
			11_limitbuang=?, 11_limittotal=?, 
			updatecomppas=?, updatedatecompas=? 
			WHERE idcomppasaran=? AND idcompany=? 
		`
		log.Println("SHIO")
		log.Printf("WIN : %s", fmt.Sprintf("%.3f", win))
		log.Printf("DISC : %s", fmt.Sprintf("%.3f", disc))
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_company_game_pasaran, "UPDATE",
			shiotahunini,
			minbet, maxbet, maxbuy,
			fmt.Sprintf("%.3f", win), fmt.Sprintf("%.3f", disc),
			limitglobal, limittotal,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idcomppasaran,
			company)
		log.Printf("%d-%d", minbet, maxbet)
		if flag_update {
			msg = "Success"
			log.Println(msg_update)

			nmpasarantogel := Pasaranmaster_id(pasarancode, "nmpasarantogel")
			noteafter := ""
			noteafter += "PASARAN - " + nmpasarantogel + "<br />"
			noteafter += "MINIMAL BET - " + strconv.Itoa(minbet) + "<br />"
			noteafter += "MAXIMAL BET - " + strconv.Itoa(maxbet) + "<br />"
			noteafter += "MAXIMAL BUY - " + strconv.Itoa(maxbuy) + "<br />"
			noteafter += "WIN - " + fmt.Sprintf("%.3f", win) + "<br />"
			noteafter += "DISC - " + fmt.Sprintf("%.3f", disc) + "<br />"
			noteafter += "LIMIT GLOBAL  - " + strconv.Itoa(limitglobal) + "<br />"
			noteafter += "LIMIT TOTAL - " + strconv.Itoa(limittotal)
			Insert_log(company, agent, "PASARAN", "UPDATE PASARAN - SHIO", "", noteafter)
		} else {
			log.Println(msg_update)
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func _pasaran_online(idcomppasaran int, company, tipecolumn string) string {
	con := db.CreateCon()
	ctx := context.Background()
	var result string = ""
	sql_pasaran := `SELECT 
		haripasaran 
		FROM ` + config.DB_tbl_mst_company_game_pasaran_offline + `  
		WHERE idcomppasaran  = ? 
		AND idcompany = ? 
	`
	var (
		haripasaran_db string
	)
	rows := con.QueryRowContext(ctx, sql_pasaran, idcomppasaran, company)
	switch err := rows.Scan(&haripasaran_db); err {
	case sql.ErrNoRows:
		result = ""
	case nil:
		switch tipecolumn {
		case "haripasaran":
			result = haripasaran_db
		}
	default:
		helpers.ErrorCheck(err)
	}
	return result
}
