package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/config"
	"bitbucket.org/isbtotogroup/apibackend_go/db"
	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
	amqp "github.com/rabbitmq/amqp091-go"
)

var mutex sync.RWMutex

type periodeHome struct {
	No                int     `json:"pasaran_no"`
	Idtrxkeluaran     int     `json:"pasaran_invoice"`
	Idcomppasaran     int     `json:"pasaran_idcompp"`
	Pasarancode       string  `json:"pasaran_code"`
	Keluaranperiode   string  `json:"pasaran_periode"`
	Nmpasaran         string  `json:"pasaran_name"`
	Tanggalperiode    string  `json:"pasaran_tanggal"`
	Keluarantogel     string  `json:"pasaran_keluaran"`
	Status            string  `json:"pasaran_status"`
	Status_css        string  `json:"pasaran_status_css"`
	Total_Member      float32 `json:"pasaran_totalmember"`
	Total_bet         float32 `json:"pasaran_totalbet"`
	Total_outstanding float32 `json:"pasaran_totaloutstanding"`
	Total_cancelbet   float32 `json:"pasaran_totalcancelbet"`
	Winlose           float32 `json:"pasaran_winlose"`
	Revisi            int     `json:"pasaran_revisi"`
	Msgrevisi         string  `json:"pasaran_msgrevisi"`
}
type periodeHomePasaran struct {
	Idcomppasaran int    `json:"pasarancomp_idcompp"`
	Pasarannama   string `json:"pasarancomp_nama"`
}
type periodeEdit struct {
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
type periodeBet struct {
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
	Bet_diskonpercen float32 `json:"bet_diskonpercen"`
	Bet_kei          int     `json:"bet_kei"`
	Bet_keipercen    float32 `json:"bet_keipercen"`
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
type listbetTable struct {
	Permainan string `json:"permainan"`
}

type ListBet struct {
	Nomortogel  string `json:"bet_keluaran"`
	Totalmember int    `json:"bet_totalmember"`
	Totalbet    int    `json:"bet_totalbet"`
}
type listMember struct {
	Member         string `json:"member"`
	Totalbet       int    `json:"totalbet"`
	Totalbayar     int    `json:"totalbayar"`
	Totalcancelbet int    `json:"totalcancelbet"`
	Totalwin       int    `json:"totalwin"`
}
type listMemberByNomor struct {
	Member      string  `json:"member"`
	Permainan   string  `json:"member_permainan"`
	Nomor       string  `json:"member_nomor"`
	Posisitogel string  `json:"member_posisitogel"`
	Bet         int     `json:"member_bet"`
	Disc        int     `json:"member_disc"`
	Discpercen  int     `json:"member_discpercen"`
	Kei         int     `json:"member_kei"`
	Keipercen   int     `json:"member_keipercen"`
	Bayar       int     `json:"member_bayar"`
	Win         float32 `json:"member_win"`
	Winhasil    int     `json:"member_winhasil"`
}
type listPasaran struct {
	Pasaran_idcomp int    `json:"pasaran_idcomp"`
	Pasaran_code   string `json:"pasaran_code"`
	Pasaran_name   string `json:"pasaran_name"`
}
type listPrediksi struct {
	Prediksi_invoice      string  `json:"prediksi_invoice"`
	Prediksi_code         string  `json:"prediksi_code"`
	Prediksi_tanggal      string  `json:"prediksi_tanggal"`
	Prediksi_username     string  `json:"prediksi_username"`
	Prediksi_permainan    string  `json:"prediksi_permainan"`
	Prediksi_nomor        string  `json:"prediksi_nomor"`
	Prediksi_posisitogel  string  `json:"prediksi_posisitogel"`
	Prediksi_bet          int     `json:"prediksi_bet"`
	Prediksi_diskon       int     `json:"prediksi_diskon"`
	Prediksi_diskonpercen float32 `json:"prediksi_diskonpercen"`
	Prediksi_kei          int     `json:"prediksi_kei"`
	Prediksi_keipercen    float32 `json:"prediksi_keipercen"`
	Prediksi_bayar        int     `json:"prediksi_bayar"`
	Prediksi_win          float32 `json:"prediksi_win"`
	Prediksi_totalwin     int     `json:"prediksi_totalwin"`
	Prediksi_status       string  `json:"prediksi_status"`
	Prediksi_statuscss    string  `json:"prediksi_statuscss"`
}

func Fetch_periode(company string) (helpers.ResponsePasaran, error) {
	var obj periodeHome
	var arraobj []periodeHome
	var res helpers.ResponsePasaran
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	flag := true
	tbl_trx_keluarantogel, _, _ := Get_mappingdatabase(company)

	sql_periode := `SELECT 
			A.idtrxkeluaran, A.idcomppasaran, A.keluaranperiode, A.datekeluaran, A.keluarantogel, 
			A.total_member, A.total_bet, A.total_outstanding, A.winlose, A.total_cancel, 
			C.nmpasarantogel, B.idpasarantogel, A.revisi, A.noterevisi    
			FROM ` + tbl_trx_keluarantogel + ` as A 
			JOIN ` + config.DB_tbl_mst_company_game_pasaran + ` as B ON B.idcomppasaran = A.idcomppasaran  
			JOIN ` + config.DB_tbl_mst_pasaran_togel + ` as C ON C.idpasarantogel  = B.idpasarantogel  
			WHERE A.idcompany = ? 
			AND A.keluarantogel = "" 
			ORDER BY A.datekeluaran DESC  
		`
	row, err := con.QueryContext(ctx, sql_periode, company)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			idtrxkeluaran_db, idcomppasaran_db, revisi_db                                                              int
			datekeluaran_db, keluarantogel_db, nmpasarantogel_db, idpasarantogel_db, keluaranperiode_db, noterevisi_db string
			total_member_db, total_bet_db, total_outstanding_db, winlose_db, total_cancel_db                           float32
		)

		err = row.Scan(
			&idtrxkeluaran_db, &idcomppasaran_db, &keluaranperiode_db,
			&datekeluaran_db, &keluarantogel_db, &total_member_db,
			&total_bet_db, &total_outstanding_db, &winlose_db, &total_cancel_db,
			&nmpasarantogel_db, &idpasarantogel_db, &revisi_db, &noterevisi_db)

		helpers.ErrorCheck(err)
		status := "DONE"
		status_css := config.STATUS_COMPLETE
		if keluarantogel_db == "" {
			status = "RUNNING"
			status_css = config.STATUS_RUNNING
		}
		totalwinlose := total_outstanding_db - total_cancel_db - winlose_db

		obj.No = no
		obj.Idtrxkeluaran = idtrxkeluaran_db
		obj.Idcomppasaran = idcomppasaran_db
		obj.Pasarancode = idpasarantogel_db
		obj.Nmpasaran = nmpasarantogel_db
		obj.Keluaranperiode = keluaranperiode_db + "-" + idpasarantogel_db
		obj.Tanggalperiode = datekeluaran_db
		obj.Keluarantogel = keluarantogel_db
		obj.Total_Member = total_member_db
		obj.Total_bet = total_bet_db
		obj.Total_outstanding = total_outstanding_db
		obj.Total_cancelbet = total_cancel_db
		obj.Winlose = totalwinlose
		obj.Revisi = revisi_db
		obj.Msgrevisi = noterevisi_db
		obj.Status = status
		obj.Status_css = status_css
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	var obj_pasar periodeHomePasaran
	var arraobj_pasar []periodeHomePasaran
	sql_pasaran := `SELECT 
			A.idcomppasaran, B.nmpasarantogel  
			FROM ` + config.DB_tbl_mst_company_game_pasaran + ` as A 
			JOIN ` + config.DB_tbl_mst_pasaran_togel + ` as B ON B.idpasarantogel = A.idpasarantogel  
			WHERE A.idcompany = ? 
			AND A.statuspasaranactive = 'Y'
			ORDER BY B.nmpasarantogel DESC 
		`
	row_pasaran, err_pasaran := con.QueryContext(ctx, sql_pasaran, company)
	helpers.ErrorCheck(err_pasaran)
	for row_pasaran.Next() {
		no += 1
		var (
			idcomppasaran_db  int
			nmpasarantogel_db string
		)

		err = row_pasaran.Scan(&idcomppasaran_db, &nmpasarantogel_db)
		helpers.ErrorCheck(err)
		flag = Get_OnlinePasaran(company, idcomppasaran_db, "", "total_pasaran")

		if flag {
			obj_pasar.Idcomppasaran = idcomppasaran_db
			obj_pasar.Pasarannama = nmpasarantogel_db
			arraobj_pasar = append(arraobj_pasar, obj_pasar)
			msg = "Success"
		}
	}
	defer row_pasaran.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Pasaraonline = arraobj_pasar
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Fetch_periodedetail(company string, idtrxkeluaran int) (helpers.Response, error) {
	var obj periodeEdit
	var arraobj []periodeEdit
	var res helpers.Response
	tglnow, _ := goment.New()
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	flag := false
	tbl_trx_keluarantogel, _, _ := Get_mappingdatabase(company)

	sql_select := `SELECT 
		A.idtrxkeluaran, A.idcomppasaran, A.keluaranperiode, A.datekeluaran, A.keluarantogel, 
		A.createkeluaran, A.createdatekeluaran, A.updatekeluaran, COALESCE(A.updatedatekeluaran,''), 
		B.idpasarantogel, B.jamtutup, B.jamjadwal, B.jamopen, A.revisi      
		FROM ` + tbl_trx_keluarantogel + ` as A 
		JOIN ` + config.DB_tbl_mst_company_game_pasaran + ` as B ON B.idcomppasaran = A.idcomppasaran 
		WHERE A.idcompany = ? 
		AND A.idtrxkeluaran = ? 
	`
	var (
		idtrxkeluaran_db, idcomppasaran_db, revisi_db                                                         int
		keluaranperiode_db, datekeluaran_db, keluarantogel_db                                                 string
		createkeluaran_db, createdatekeluaran_db, updatekeluaran_db, updatedatekeluaran_db, idpasarantogel_db string
		jamtutup_db, jamjadwal_db, jamopen_db                                                                 string
	)

	row := con.QueryRowContext(ctx, sql_select, company, idtrxkeluaran)
	switch e := row.Scan(&idtrxkeluaran_db, &idcomppasaran_db, &keluaranperiode_db, &datekeluaran_db, &keluarantogel_db,
		&createkeluaran_db, &createdatekeluaran_db, &updatekeluaran_db, &updatedatekeluaran_db, &idpasarantogel_db,
		&jamtutup_db, &jamjadwal_db, &jamopen_db, &revisi_db); e {
	case sql.ErrNoRows:
		log.Println("Data not found in DB")
		flag = false
	case nil:
		flag = true
	default:
		helpers.ErrorCheck(e)
		flag = false
	}
	if flag {
		tglopen, _ := goment.New(datekeluaran_db)
		tglskrg := tglnow.Format("YYYY-MM-DD HH:mm:ss")
		tglskrgend := tglnow.Format("YYYY-MM-DD") + " 23:59:59"
		jamtutup := tglnow.Format("YYYY-MM-DD") + " " + jamtutup_db
		jamtutup2 := tglopen.Format("YYYY-MM-DD") + " " + jamtutup_db
		jamopen := tglnow.Format("YYYY-MM-DD") + " " + jamopen_db
		jamopen2 := tglopen.Format("YYYY-MM-DD") + " " + jamopen_db
		statuspasaran := "OFFLINE"
		statusrevisi := "LOCK"
		if tglskrg > jamopen2 { //EXPIRE
			statuspasaran = "OFFLINE"
			statusrevisi = "LOCK"
		} else {
			if tglskrgend < jamtutup2 { // jikga tgl skrg end dibawah tgltuptup
				statuspasaran = "ONLINE"
			} else {
				flag := _checkpasaranonline(idcomppasaran_db, company)
				if flag {
					if tglskrg >= jamtutup && tglskrg <= jamopen {
						statuspasaran = "OFFLINE"
					} else {
						statuspasaran = "ONLINE"
					}

					if keluarantogel_db != "" {
						if revisi_db < 1 {
							statusrevisi = "OPEN"
						}
					}

					if updatedatekeluaran_db != "" {
						tglupdate, _ := goment.New(updatedatekeluaran_db)
						tglexpirerevisi := tglupdate.Add(30, "minutes").Format("YYYY-MM-DD HH:mm:ss")

						if tglexpirerevisi < tglskrg {
							statusrevisi = "LOCK"
						}
					}
				} else {
					statuspasaran = "ONLINE"
				}
			}

		}
		obj.Idinvoice = strconv.Itoa(idtrxkeluaran)
		obj.TanggalPeriode = datekeluaran_db
		obj.TanggalNext = Get_NextPasaran(company, datekeluaran_db, idcomppasaran_db)
		obj.PeriodeKeluaran = keluaranperiode_db + "-" + idpasarantogel_db
		obj.Keluaran = keluarantogel_db
		obj.Statusrevisi = statusrevisi
		obj.StatusOnlineOffice = statuspasaran
		obj.Create = createkeluaran_db
		obj.CreateDate = createdatekeluaran_db
		obj.Update = updatekeluaran_db
		obj.UpdateDate = updatedatekeluaran_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Fetch_membergroupbynomor(company, typegame, nomortogel string, idtrxkeluaran int) (helpers.Response, error) {
	var obj listMemberByNomor
	var arraobj []listMemberByNomor
	var res helpers.Response
	msg := "Success"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	_, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sql := `SELECT 
		username, 
		nomortogel, typegame, posisitogel, 
		bet, diskon, kei, win 
		FROM ` + tbl_trx_keluarantogel_detail + `  
		WHERE idcompany = ? 
		AND idtrxkeluaran = ? 
		AND typegame = ? 
		AND nomortogel = ? 
	`

	row, err := con.QueryContext(ctx, sql, company, idtrxkeluaran, typegame, nomortogel)

	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			bet_db                                                  int
			diskon_db, kei_db, win_db                               float32
			username_db, nomortogel_db, typegame_db, posisitogel_db string
		)

		err = row.Scan(&username_db, &nomortogel_db, &typegame_db, &posisitogel_db, &bet_db, &diskon_db, &kei_db, &win_db)

		helpers.ErrorCheck(err)

		disc := int(float32(bet_db) * diskon_db)
		discpercen := diskon_db * 100
		kei := int(float32(bet_db) * kei_db)
		keipercen := kei_db * 100
		bayar := bet_db - int(float32(bet_db)*diskon_db) - int(float32(bet_db)*kei_db)
		winhasil := _rumuswinhasil(typegame_db, bayar, bet_db, win_db)

		obj.Member = username_db
		obj.Nomor = nomortogel_db
		obj.Permainan = typegame_db
		obj.Posisitogel = posisitogel_db
		obj.Bet = bet_db
		obj.Disc = disc
		obj.Discpercen = int(discpercen)
		obj.Kei = kei
		obj.Keipercen = int(keipercen)
		obj.Bayar = bayar
		obj.Win = win_db
		obj.Winhasil = winhasil
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Fetch_membergroup(company string, idtrxkeluaran int) (helpers.Response, error) {
	var obj listMember
	var arraobj []listMember
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()

	_, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sql := `SELECT 
		username, 
		count(username) as totalbet, 
		sum(bet-(bet*diskon)-(bet*kei)) as totalbayar,
		sum(cancelbet) as totalcancel,  
		sum(winhasil) as totalwin 
		FROM ` + tbl_trx_keluarantogel_detail + ` 
		WHERE idcompany = ? 
		AND idtrxkeluaran = ? 
		GROUP BY username 
	`

	row, err := con.QueryContext(ctx, sql, company, idtrxkeluaran)

	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			totalbet_db, totalbayar_db, totalwin_db, totalcancel_db float64
			username_db                                             string
		)

		err = row.Scan(
			&username_db,
			&totalbet_db,
			&totalbayar_db,
			&totalcancel_db,
			&totalwin_db)

		helpers.ErrorCheck(err)

		obj.Member = username_db
		obj.Totalbet = int(totalbet_db)
		obj.Totalbayar = int(totalbayar_db)
		obj.Totalcancelbet = int(totalcancel_db)
		obj.Totalwin = int(totalwin_db)
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Fetch_listbet(company, permainan string, idtrxkeluaran int) (helpers.ResponsePeriode, error) {
	var obj periodeBet
	var arraobj []periodeBet
	var res helpers.ResponsePeriode
	ctx := context.Background()
	con := db.CreateCon()
	msg := "failed"
	totalbet := 0
	subtotalbayar := 0
	subtotalwin := 0
	render_page := time.Now()
	_, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sqldetail := `SELECT
					idtrxkeluarandetail , datetimedetail, ipaddress, browsertogel, devicetogel,  username, typegame, nomortogel, 
					bet, diskon, win, kei, statuskeluarandetail,posisitogel, createkeluarandetail, 
					createdatekeluarandetail, updatekeluarandetail, updatedatekeluarandetail 
					FROM ` + tbl_trx_keluarantogel_detail + ` 
					WHERE idcompany = ? 
					AND idtrxkeluaran = ? 
					AND typegame = ? 
					ORDER BY datetimedetail DESC 
				`
	row, err := con.QueryContext(ctx, sqldetail, company, idtrxkeluaran, permainan)

	helpers.ErrorCheck(err)

	for row.Next() {

		var (
			idtrxkeluarandetail_db                                                                                                                              int
			datetimedetail_db, ipaddresss_db, username_db, typegame_db, nomortogel_db, browsertogel_db, devicetogel_db                                          string
			statuskeluarandetail_db, posisitogel_db, createkeluarandetail_db, createdatekeluarandetail_db, updatekeluarandetail_db, updatedatekeluarandetail_db string
			diskon_db, win_db, kei_db, bet_db                                                                                                                   float32
		)

		err = row.Scan(
			&idtrxkeluarandetail_db,
			&datetimedetail_db, &ipaddresss_db, &browsertogel_db, &devicetogel_db, &username_db, &typegame_db, &nomortogel_db,
			&bet_db, &diskon_db, &win_db, &kei_db, &statuskeluarandetail_db, &posisitogel_db, &createkeluarandetail_db,
			&createdatekeluarandetail_db, &updatekeluarandetail_db, &updatedatekeluarandetail_db)

		helpers.ErrorCheck(err)
		if statuskeluarandetail_db != "CANCEL" {
			totalbet += 1
			diskonpercen := diskon_db * 100
			diskonbet := int(float32(bet_db) * diskon_db)
			keipercen := kei_db * 100
			keibet := int(float32(bet_db) * kei_db)
			bayar := int(bet_db) - int(float32(bet_db)*diskon_db) - int(float32(bet_db)*kei_db)
			subtotalbayar = subtotalbayar + bayar
			winhasil := _rumuswinhasil(typegame_db, bayar, int(bet_db), win_db)
			totalwin := 0

			status_css := ""
			switch statuskeluarandetail_db {
			case "RUNNING":
				totalwin = 0
				status_css = config.STATUS_RUNNING
			case "WINNER":
				totalwin = winhasil
				subtotalwin = subtotalwin + winhasil
				status_css = config.STATUS_COMPLETE
			case "LOSE":
				totalwin = 0
				status_css = config.STATUS_CANCEL
			case "CANCEL":
				totalwin = 0
				status_css = config.STATUS_CANCELBET
			}

			obj.Bet_id = idtrxkeluarandetail_db
			obj.Bet_datetime = datetimedetail_db
			obj.Bet_ipaddress = ipaddresss_db
			obj.Bet_device = devicetogel_db
			obj.Bet_timezone = browsertogel_db
			obj.Bet_username = username_db
			obj.Bet_typegame = typegame_db
			obj.Bet_nomortogel = nomortogel_db
			obj.Bet_posisitogel = posisitogel_db
			obj.Bet_bet = int(bet_db)
			obj.Bet_diskon = diskonbet
			obj.Bet_diskonpercen = diskonpercen
			obj.Bet_kei = keibet
			obj.Bet_keipercen = keipercen
			obj.Bet_bayar = bayar
			obj.Bet_win = win_db
			obj.Bet_totalwin = totalwin
			obj.Bet_status = statuskeluarandetail_db
			obj.Bet_statuscss = status_css
			obj.Bet_create = createkeluarandetail_db
			obj.Bet_createDate = createdatekeluarandetail_db
			obj.Bet_update = updatekeluarandetail_db
			obj.Bet_updateDate = updatedatekeluarandetail_db
			arraobj = append(arraobj, obj)
			msg = "Success"
		}
	}
	defer row.Close()
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()
	res.Totalbet = totalbet
	res.Subtotal = subtotalbayar
	res.Subtotalwin = subtotalwin

	return res, nil
}
func Fetch_listbetbystatus(company, status string, idtrxkeluaran int) (helpers.ResponsePeriode, error) {
	var obj periodeBet
	var arraobj []periodeBet
	var res helpers.ResponsePeriode
	ctx := context.Background()
	con := db.CreateCon()
	msg := "failed"
	totalbet := 0
	subtotalbayar := 0
	subtotalwin := 0
	render_page := time.Now()
	_, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sqldetail := `SELECT
					idtrxkeluarandetail , datetimedetail, ipaddress, browsertogel, devicetogel,  username, typegame, nomortogel, 
					bet, diskon, win, kei, statuskeluarandetail,posisitogel ,createkeluarandetail, 
					createdatekeluarandetail, updatekeluarandetail, updatedatekeluarandetail 
					FROM ` + tbl_trx_keluarantogel_detail + ` 
					WHERE idcompany = ? 
					AND idtrxkeluaran = ? 
					AND statuskeluarandetail = ? 
					ORDER BY datetimedetail DESC 
				`
	row, err := con.QueryContext(ctx, sqldetail, company, idtrxkeluaran, status)

	helpers.ErrorCheck(err)

	for row.Next() {
		totalbet += 1
		var (
			idtrxkeluarandetail_db                                                                                                                              int
			datetimedetail_db, ipaddresss_db, username_db, typegame_db, nomortogel_db, browsertogel_db, devicetogel_db                                          string
			statuskeluarandetail_db, posisitogel_db, createkeluarandetail_db, createdatekeluarandetail_db, updatekeluarandetail_db, updatedatekeluarandetail_db string
			diskon_db, win_db, kei_db, bet_db                                                                                                                   float32
		)

		err = row.Scan(
			&idtrxkeluarandetail_db,
			&datetimedetail_db, &ipaddresss_db, &browsertogel_db, &devicetogel_db, &username_db, &typegame_db, &nomortogel_db,
			&bet_db, &diskon_db, &win_db, &kei_db, &statuskeluarandetail_db, &posisitogel_db, &createkeluarandetail_db,
			&createdatekeluarandetail_db, &updatekeluarandetail_db, &updatedatekeluarandetail_db)

		helpers.ErrorCheck(err)

		diskonpercen := diskon_db * 100
		diskonbet := int(float32(bet_db) * diskon_db)
		keipercen := kei_db * 100
		keibet := int(float32(bet_db) * kei_db)
		bayar := int(bet_db) - int(float32(bet_db)*diskon_db) - int(float32(bet_db)*kei_db)
		subtotalbayar = subtotalbayar + bayar
		winhasil := _rumuswinhasil(typegame_db, bayar, int(bet_db), win_db)
		totalwin := 0

		status_css := ""
		switch statuskeluarandetail_db {
		case "RUNNING":
			totalwin = 0
			status_css = config.STATUS_RUNNING
		case "WINNER":
			totalwin = winhasil
			subtotalwin = subtotalwin + winhasil
			status_css = config.STATUS_COMPLETE
		case "LOSE":
			totalwin = 0
			status_css = config.STATUS_CANCEL
		case "CANCEL":
			totalwin = 0
			status_css = config.STATUS_CANCELBET
		}

		obj.Bet_id = idtrxkeluarandetail_db
		obj.Bet_datetime = datetimedetail_db
		obj.Bet_ipaddress = ipaddresss_db
		obj.Bet_device = devicetogel_db
		obj.Bet_timezone = browsertogel_db
		obj.Bet_username = username_db
		obj.Bet_typegame = typegame_db
		obj.Bet_nomortogel = nomortogel_db
		obj.Bet_posisitogel = posisitogel_db
		obj.Bet_bet = int(bet_db)
		obj.Bet_diskon = diskonbet
		obj.Bet_diskonpercen = diskonpercen
		obj.Bet_kei = keibet
		obj.Bet_keipercen = keipercen
		obj.Bet_bayar = bayar
		obj.Bet_win = win_db
		obj.Bet_totalwin = totalwin
		obj.Bet_status = statuskeluarandetail_db
		obj.Bet_statuscss = status_css
		obj.Bet_create = createkeluarandetail_db
		obj.Bet_createDate = createdatekeluarandetail_db
		obj.Bet_update = updatekeluarandetail_db
		obj.Bet_updateDate = updatedatekeluarandetail_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()
	res.Totalbet = totalbet
	res.Subtotal = subtotalbayar
	res.Subtotalwin = subtotalwin

	return res, nil
}
func Fetch_listbetbyusername(company, username string, idtrxkeluaran int) (helpers.ResponsePeriode, error) {
	var obj periodeBet
	var arraobj []periodeBet
	var res helpers.ResponsePeriode
	ctx := context.Background()
	con := db.CreateCon()
	msg := "failed"
	totalbet := 0
	subtotalbayar := 0
	subtotalwin := 0
	render_page := time.Now()
	_, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sqldetail := `SELECT
					idtrxkeluarandetail , datetimedetail, ipaddress, browsertogel, devicetogel,  username, typegame, nomortogel, 
					bet, diskon, win, kei, statuskeluarandetail, posisitogel, createkeluarandetail, 
					createdatekeluarandetail, updatekeluarandetail, updatedatekeluarandetail 
					FROM ` + tbl_trx_keluarantogel_detail + ` 
					WHERE idcompany = ? 
					AND idtrxkeluaran = ? 
					AND username = ? 
					ORDER BY datetimedetail DESC 
				`
	row, err := con.QueryContext(ctx, sqldetail, company, idtrxkeluaran, username)

	helpers.ErrorCheck(err)

	for row.Next() {
		totalbet += 1
		var (
			idtrxkeluarandetail_db                                                                                                                              int
			datetimedetail_db, ipaddresss_db, username_db, typegame_db, nomortogel_db, browsertogel_db, devicetogel_db                                          string
			statuskeluarandetail_db, posisitogel_db, createkeluarandetail_db, createdatekeluarandetail_db, updatekeluarandetail_db, updatedatekeluarandetail_db string
			diskon_db, kei_db                                                                                                                                   float64
			win_db, bet_db                                                                                                                                      float32
		)

		err = row.Scan(
			&idtrxkeluarandetail_db,
			&datetimedetail_db, &ipaddresss_db, &browsertogel_db, &devicetogel_db, &username_db, &typegame_db, &nomortogel_db,
			&bet_db, &diskon_db, &win_db, &kei_db, &statuskeluarandetail_db, &posisitogel_db, &createkeluarandetail_db,
			&createdatekeluarandetail_db, &updatekeluarandetail_db, &updatedatekeluarandetail_db)

		helpers.ErrorCheck(err)

		diskonpercen := diskon_db * 100
		diskonbet := math.Ceil(float64(bet_db) * diskon_db)
		keipercen := kei_db * 100
		keibet := math.Ceil(float64(bet_db) * kei_db)
		bayar := int(bet_db) - int(diskonbet) - int(keibet)
		subtotalbayar = subtotalbayar + bayar
		winhasil := _rumuswinhasil(typegame_db, bayar, int(bet_db), win_db)
		totalwin := 0
		status_css := ""
		switch statuskeluarandetail_db {
		case "RUNNING":
			totalwin = 0
			status_css = config.STATUS_RUNNING
		case "WINNER":
			totalwin = winhasil
			subtotalwin = subtotalwin + winhasil
			status_css = config.STATUS_COMPLETE
		case "LOSE":
			totalwin = 0
			status_css = config.STATUS_CANCEL
		case "CANCEL":
			totalwin = 0
			status_css = config.STATUS_CANCELBET
		}

		obj.Bet_id = idtrxkeluarandetail_db
		obj.Bet_datetime = datetimedetail_db
		obj.Bet_ipaddress = ipaddresss_db
		obj.Bet_device = devicetogel_db
		obj.Bet_timezone = browsertogel_db
		obj.Bet_username = username_db
		obj.Bet_typegame = typegame_db
		obj.Bet_nomortogel = nomortogel_db
		obj.Bet_posisitogel = posisitogel_db
		obj.Bet_bet = int(bet_db)
		obj.Bet_diskon = int(diskonbet)
		obj.Bet_diskonpercen = float32(diskonpercen)
		obj.Bet_kei = int(keibet)
		obj.Bet_keipercen = float32(keipercen)
		obj.Bet_bayar = bayar
		obj.Bet_win = win_db
		obj.Bet_totalwin = totalwin
		obj.Bet_status = statuskeluarandetail_db
		obj.Bet_statuscss = status_css
		obj.Bet_create = createkeluarandetail_db
		obj.Bet_createDate = createdatekeluarandetail_db
		obj.Bet_update = updatekeluarandetail_db
		obj.Bet_updateDate = updatedatekeluarandetail_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()
	res.Totalbet = totalbet
	res.Subtotal = subtotalbayar
	res.Subtotalwin = subtotalwin

	return res, nil
}
func Fetch_listpasaran(company string) (helpers.Response, error) {
	var obj listPasaran
	var arraobj []listPasaran
	var res helpers.Response
	ctx := context.Background()
	con := db.CreateCon()
	msg := "failed"

	render_page := time.Now()
	sql_listpasaran := `SELECT
					A.idcomppasaran , A.idpasarantogel , B.nmpasarantogel 
					FROM ` + config.DB_tbl_mst_company_game_pasaran + ` as A  
					JOIN ` + config.DB_tbl_mst_pasaran_togel + ` as B ON B.idpasarantogel = A.idpasarantogel 
					WHERE A.idcompany = ? 
					AND A.statuspasaranactive = 'Y' 
					ORDER BY B.nmpasarantogel ASC  
				`
	row, err := con.QueryContext(ctx, sql_listpasaran, company)

	helpers.ErrorCheck(err)

	for row.Next() {
		var (
			idcomppasaran_db                     int
			idpasarantogel_db, nmpasarantogel_db string
		)

		err = row.Scan(
			&idcomppasaran_db,
			&idpasarantogel_db, &nmpasarantogel_db)

		helpers.ErrorCheck(err)

		obj.Pasaran_idcomp = idcomppasaran_db
		obj.Pasaran_code = idpasarantogel_db
		obj.Pasaran_name = nmpasarantogel_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Fetch_listprediksi(company, nomorkeluaran string, idcomppasaran int) (helpers.ResponsePeriode, error) {
	var obj listPrediksi
	var arraobj []listPrediksi
	var res helpers.ResponsePeriode
	ctx := context.Background()
	con := db.CreateCon()
	msg := "failed"
	totalbet := 0
	subtotal := 0
	subtotalwin := 0
	render_page := time.Now()
	tbl_trx_keluarantogel, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sql_listprediksi := `SELECT
					A.idtrxkeluaran, A.idtrxkeluarandetail, A.datetimedetail , A.username , A.typegame, A.nomortogel, A.posisitogel, 
					A.bet, A.diskon, A.kei, A.win  
					FROM ` + tbl_trx_keluarantogel_detail + ` as A  
					JOIN ` + tbl_trx_keluarantogel + ` as B ON B.idtrxkeluaran = A.idtrxkeluaran  
					WHERE B.idcompany = ? 
					AND B.idcomppasaran = ?  
					AND B.keluarantogel = '' 
					ORDER BY A.datetimedetail DESC   
				`

	row, err := con.QueryContext(ctx, sql_listprediksi, company, idcomppasaran)
	helpers.ErrorCheck(err)

	for row.Next() {
		var (
			bet_db                                                                                                               int
			diskon_db, kei_db, win_db                                                                                            float32
			idtrxkeluaran_db, idtrxkeluarandetail_db, datetimedetail_db, username_db, typegame_db, nomortogel_db, posisitogel_db string
		)

		err = row.Scan(
			&idtrxkeluaran_db, &idtrxkeluarandetail_db, &datetimedetail_db, &username_db, &typegame_db, &nomortogel_db, &posisitogel_db,
			&bet_db, &diskon_db, &kei_db, &win_db)

		helpers.ErrorCheck(err)
		diskonpercen := diskon_db * 100
		diskonbet := int(float32(bet_db) * diskon_db)
		keipercen := kei_db * 100
		keibet := int(float32(bet_db) * kei_db)

		bayar := bet_db - int(float32(bet_db)*diskon_db) - int(float32(bet_db)*kei_db)
		subtotal = subtotal + bayar
		statuskeluarandetail, winrumus := _rumusTogel(nomorkeluaran, typegame_db, nomortogel_db, posisitogel_db, company, "N", idcomppasaran, 0)
		// statuskeluarandetail, _ := _rumusTogel(keluarantogel, typegame_db, nomortogel_db, posisitogel_db, company, "Y", idcomppasaran, idtrxkeluarandetail_db)
		var winfixed float32 = 0
		winhasil := 0
		if winrumus == 0 {
			winrumus = win_db
		}
		if typegame_db == "COLOK_BEBAS" || typegame_db == "COLOK_MACAU" || typegame_db == "COLOK_NAGA" {
			winhasil = _rumuswinhasil(typegame_db, bayar, bet_db, winrumus)
			winfixed = winrumus
		} else {
			winhasil = _rumuswinhasil(typegame_db, bayar, bet_db, winrumus)
			winfixed = winrumus
		}
		status_css := ""

		switch statuskeluarandetail {
		case "WINNER":
			totalbet = totalbet + 1
			subtotalwin = subtotalwin + winhasil
			status_css = "background:#8BC34A;color:black;font-weight:bold;"
		case "LOSE":
			status_css = "background:#E91E63;font-size:12px;font-weight:bold;color:white;"
		}
		if statuskeluarandetail == "WINNER" {
			obj.Prediksi_invoice = idtrxkeluaran_db
			obj.Prediksi_code = idtrxkeluarandetail_db
			obj.Prediksi_tanggal = datetimedetail_db
			obj.Prediksi_username = username_db
			obj.Prediksi_permainan = typegame_db
			obj.Prediksi_nomor = nomortogel_db
			obj.Prediksi_posisitogel = posisitogel_db
			obj.Prediksi_bet = bet_db
			obj.Prediksi_diskon = diskonbet
			obj.Prediksi_diskonpercen = diskonpercen
			obj.Prediksi_kei = keibet
			obj.Prediksi_keipercen = keipercen
			obj.Prediksi_bayar = bayar
			obj.Prediksi_win = winfixed
			obj.Prediksi_totalwin = winhasil
			obj.Prediksi_status = statuskeluarandetail
			obj.Prediksi_statuscss = status_css
			arraobj = append(arraobj, obj)
		}
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Totalbet = totalbet
	res.Subtotal = subtotalwin
	res.Subtotalwin = subtotal - subtotalwin
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Fetch_listbettable(company string, idtrxkeluaran int) (helpers.Response, error) {
	var obj listbetTable
	var arraobj []listbetTable
	var res helpers.Response
	render_page := time.Now()
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	_, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sql := `SELECT 
		A.typegame 
		FROM ` + tbl_trx_keluarantogel_detail + ` AS A 
		JOIN ` + config.DB_tbl_mst_permainan + ` AS B ON B.idpermainan = A.typegame 
		WHERE A.idcompany = ? 
		AND A.idtrxkeluaran = ? 
		GROUP BY A.typegame ORDER BY B.display ASC 
	`
	row, err := con.QueryContext(ctx, sql, company, idtrxkeluaran)
	helpers.ErrorCheck(err)
	for row.Next() {
		var typegame_db string

		err = row.Scan(&typegame_db)
		helpers.ErrorCheck(err)
		obj.Permainan = typegame_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()
	return res, nil
}
func Fetch_bettable(company, permainan string, idtrxkeluaran int) (helpers.Response, error) {
	var objdetail ListBet
	var arraobjdetail []ListBet
	var res helpers.Response
	render_page := time.Now()
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	_, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sql_betgroup := `SELECT 
			nomortogel, COUNT(username) as totalmember, SUM(bet) as totalbet 
			FROM ` + tbl_trx_keluarantogel_detail + `  
			WHERE idcompany = ? 
			AND idtrxkeluaran = ? 
			AND typegame = ? 
			GROUP BY nomortogel  
			ORDER BY totalbet DESC, totalmember DESC 
		`
	row_betgroup, err_betgroup := con.QueryContext(ctx, sql_betgroup, company, idtrxkeluaran, permainan)
	helpers.ErrorCheck(err_betgroup)
	for row_betgroup.Next() {
		var (
			nomortogel_db  string
			totalmember_db int
			totalbet_db    float32
		)

		err := row_betgroup.Scan(&nomortogel_db, &totalmember_db, &totalbet_db)
		helpers.ErrorCheck(err)

		objdetail.Nomortogel = nomortogel_db
		objdetail.Totalmember = totalmember_db
		objdetail.Totalbet = int(totalbet_db)
		arraobjdetail = append(arraobjdetail, objdetail)
		msg = "Success"
	}
	defer row_betgroup.Close()
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobjdetail
	res.Time = time.Since(render_page).String()
	return res, nil
}

type senderJobssaveperiode struct {
	Idtrxkeluaran string
	Company       string
	Idcomppasaran string
	Keluarantogel string
	Agent         string
}

func Save_Periode(agent, company string, idtrxkeluaran int, keluarantogel string) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	msg := "Failed"
	idcomppasaran := 0
	datekeluaran := ""
	render_page := time.Now()
	tbl_trx_keluarantogel, _, _ := Get_mappingdatabase(company)

	temp_nomor := len(keluarantogel)

	if temp_nomor == 4 {
		flag = true
	}
	if flag { // wajib 4 digit
		flag = false
		sql := `SELECT 
			idcomppasaran, datekeluaran   
			FROM ` + tbl_trx_keluarantogel + `   
			WHERE idcompany = ? 
			AND idtrxkeluaran = ? 
		`
		err := con.QueryRowContext(ctx, sql, company, idtrxkeluaran).Scan(&idcomppasaran, &datekeluaran)

		helpers.ErrorCheck(err)
		if idcomppasaran > 0 {
			//UPDATE PARENT
			sql_updateparent := `
				UPDATE 
				` + tbl_trx_keluarantogel + `   
				SET keluarantogel=?,  
				updatekeluaran=?, updatedatekeluaran=? 
				WHERE idtrxkeluaran=? AND idcompany=? 
				`
			flag_updateparent, msg_updateparent := Exec_SQL(sql_updateparent, tbl_trx_keluarantogel, "UPDATE",
				keluarantogel, agent,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				idtrxkeluaran, company)

			if flag_updateparent {
				msg = "Succes"
				log.Println(msg_updateparent)
				flag = true
				log.Printf("Update Parent tbl_trx_keluarantogel : %d\n", idtrxkeluaran)
				idpasarantogel, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
				nmpasarantogel := Pasaranmaster_id(idpasarantogel, "nmpasarantogel")
				noteafter := ""
				noteafter += "INVOICE - " + strconv.Itoa(idtrxkeluaran) + "<br />"
				noteafter += "PASARAN : " + nmpasarantogel + "<br />"
				noteafter += "KELUARAN : " + keluarantogel
				Insert_log(company, agent, "PERIODE", "UPDATE KELUARAN", "", noteafter)

				totals_bet := _togel_bet_SUM_RUNNING(idtrxkeluaran, company)

				if totals_bet > 0 {
					//rabbitmq
					var obj_sender senderJobssaveperiode
					var arraobj_sender []senderJobssaveperiode
					AMPQ := os.Getenv("AMQP_SERVER_URL")
					conn, err := amqp.Dial(AMPQ)
					failOnError(err, "Failed to connect to RabbitMQ")
					defer conn.Close()

					ch, err := conn.Channel()
					failOnError(err, "Failed to open a channel")
					defer ch.Close()

					q, err := ch.QueueDeclare(
						"agensaveperiode", // name
						false,             // durable
						false,             // delete when unused
						false,             // exclusive
						false,             // no-wait
						nil,               // arguments
					)
					failOnError(err, "Failed to declare a queue")

					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()

					obj_sender.Idtrxkeluaran = strconv.Itoa(idtrxkeluaran)
					obj_sender.Company = company
					obj_sender.Idcomppasaran = strconv.Itoa(idcomppasaran)
					obj_sender.Agent = agent
					obj_sender.Keluarantogel = keluarantogel
					arraobj_sender = append(arraobj_sender, obj_sender)
					body, _ := json.Marshal(arraobj_sender)

					err = ch.PublishWithContext(ctx,
						"",     // exchange
						q.Name, // routing key
						false,  // mandatory
						false,  // immediate
						amqp.Publishing{
							ContentType: "text/plain",
							Body:        []byte(body),
						})
					failOnError(err, "Failed to publish a message")
					log.Printf(" [x] Sent %s\n", body)
				}

				//NEW PASARAN
				year := tglnow.Format("YYYY")
				month := tglnow.Format("MM")
				field_col := tbl_trx_keluarantogel + year + month
				idkeluaran_counter := Get_counter(field_col)
				idkeluaran := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idkeluaran_counter)
				field_col = company + "_" + idpasarantogel + "_" + year
				idperiode_counter := Get_counter(field_col)

				sql_insert := `
					insert into
					` + tbl_trx_keluarantogel + ` (
						idtrxkeluaran, yearmonth, idcomppasaran,
						idcompany, keluaranperiode, datekeluaran,
						createkeluaran, createdatekeluaran
					) values (
						?, ?, ?,
						?, ?, ?,
						?, ?
					)
				`
				flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_trx_keluarantogel, "UPDATE",
					idkeluaran,
					tglnow.Format("YYYY-MM"),
					idcomppasaran,
					company,
					idperiode_counter,
					Get_NextPasaran(company, datekeluaran, idcomppasaran),
					agent,
					tglnow.Format("YYYY-MM-DD HH:mm:ss"))

				if flag_insert {
					msg = "Succes"
					log.Println(msg_insert)
					log.Println("Data Berhasil di save")
				}
			}

		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}

func Save_PeriodeNew(agent, company string, idcomppasaran int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	con := db.CreateCon()
	ctx := context.Background()

	flag := false
	render_page := time.Now()
	idtrxkeluaran := ""
	msg := "Failed"

	tbl_trx_keluarantogel, _, _ := Get_mappingdatabase(company)

	sql_select := `SELECT 
		idtrxkeluaran   
		FROM ` + tbl_trx_keluarantogel + `   
		WHERE idcompany = ? 
		AND idcomppasaran = ? 
		AND keluarantogel = "" 
		ORDER BY idtrxkeluaran DESC LIMIT 1 
	`
	row_select := con.QueryRowContext(ctx, sql_select, company, idcomppasaran)
	switch err_select := row_select.Scan(&idtrxkeluaran); err_select {
	case sql.ErrNoRows:
		flag = false
		msg = "Cannot Insert"
	case nil:
		flag = true
	default:
		helpers.ErrorCheck(err_select)
	}

	if !flag {
		//NEW PASARAN
		sql_insert := `
			INSERT INTO 
			` + tbl_trx_keluarantogel + ` (
				idtrxkeluaran, yearmonth, idcomppasaran, 
				idcompany, keluaranperiode, datekeluaran, 
				createkeluaran, createdatekeluaran 
			) VALUES (
				?, ?, ?,
				?, ?, ?,
				?, ?
			)
		`
		idpasarantogel, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
		year := tglnow.Format("YYYY")
		month := tglnow.Format("MM")
		field_col := tbl_trx_keluarantogel + year + month
		idkeluaran_counter := Get_counter(field_col)
		idkeluaran := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idkeluaran_counter)
		field_col = company + "_" + idpasarantogel + "_" + year
		idperiode_counter := Get_counter(field_col)
		flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_trx_keluarantogel, "INSERT",
			idkeluaran,
			tglnow.Format("YYYY-MM"),
			idcomppasaran,
			company,
			idperiode_counter,
			Get_NextPasaran(company, tglnow.Format("YYYY-MM-DD"), idcomppasaran),
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
			log.Println(msg_insert)
			flag = true

			nmpasarantogel := Pasaranmaster_id(idpasarantogel, "nmpasarantogel")

			noteafter := ""
			noteafter += "INVOICE - " + idkeluaran + "<br />"
			noteafter += "PASARAN : " + nmpasarantogel
			Insert_log(company, agent, "PERIODE", "NEW PASARAN MANUAL", "", noteafter)
		} else {
			log.Println(msg_insert)
		}

	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_PeriodeRevisi(agent, company, msgrevisi string, idtrxkeluaran int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	msg := "Failed"
	flag := false
	revisi := 0
	idcomppasaran := 0
	tbl_trx_keluarantogel, tbl_trx_keluarantogel_detail, tbl_trx_keluarantogel_member := Get_mappingdatabase(company)

	sql_select := `SELECT 
		revisi, idcomppasaran    
		FROM ` + tbl_trx_keluarantogel + `   
		WHERE idcompany = ? 
		AND idtrxkeluaran = ? 
		AND keluarantogel != "" 
		ORDER BY idtrxkeluaran DESC LIMIT 1 
	`
	row_select := con.QueryRowContext(ctx, sql_select, company, idtrxkeluaran)
	switch err_select := row_select.Scan(&revisi, &idcomppasaran); err_select {
	case sql.ErrNoRows:
		msg = "Cannot Update"
	case nil:
		flag = true
	default:
		msg = err_select.Error()
	}
	if flag {
		stmt_keluarantogel_delete, e_keluarantogel_delete := con.PrepareContext(ctx, `
				DELETE FROM  
				`+tbl_trx_keluarantogel+`   
				WHERE idcomppasaran=? AND idcompany=? AND keluarantogel="" 
		`)

		helpers.ErrorCheck(e_keluarantogel_delete)
		rec_keluarantogel_delete, e_rec_keluarantogel_delete := stmt_keluarantogel_delete.ExecContext(ctx, idcomppasaran, company)
		helpers.ErrorCheck(e_rec_keluarantogel_delete)

		affect_keluarantogel_delete, err_affer_keluarantogel_delete := rec_keluarantogel_delete.RowsAffected()
		helpers.ErrorCheck(err_affer_keluarantogel_delete)

		defer stmt_keluarantogel_delete.Close()
		if affect_keluarantogel_delete > 0 {
			flag = true
			log.Printf("Delete tbl_trx_keluarantogel : %d\n", idtrxkeluaran)
		} else {
			flag = false
			log.Println("Delete tbl_trx_keluarantogel failed")
		}
		if flag {
			revisi = revisi + 1
			//UPDATE PARENT
			stmt_keluarantogel, e := con.PrepareContext(ctx, `
				UPDATE 
				`+tbl_trx_keluarantogel+`   
				SET keluarantogel=?, revisi=?, noterevisi=?, total_member=?, 
				total_bet=?, total_outstanding=?, total_win=?, total_lose=?, winlose=?, total_cancel=?, 
				updatekeluaran=?, updatedatekeluaran=? 
				WHERE idtrxkeluaran=? AND idcompany=? 
			`)
			helpers.ErrorCheck(e)
			rec_keluarantogel, e_keluarantogel := stmt_keluarantogel.ExecContext(ctx,
				"", revisi, msgrevisi, 0, 0, 0, 0, 0, 0, 0,
				agent, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				idtrxkeluaran, company)
			helpers.ErrorCheck(e_keluarantogel)

			a_keluarantogel, e_keluarantogel := rec_keluarantogel.RowsAffected()
			helpers.ErrorCheck(e_keluarantogel)

			defer stmt_keluarantogel.Close()
			if a_keluarantogel > 0 {
				flag = true
				log.Printf("Update Parent tbl_trx_keluarantogel : %d\n", idtrxkeluaran)
				idpasarantogel, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
				nmpasarantogel := Pasaranmaster_id(idpasarantogel, "nmpasarantogel")
				noteafter := ""
				noteafter += "INVOICE - " + strconv.Itoa(idtrxkeluaran) + "<br />"
				noteafter += "PASARAN : " + nmpasarantogel + "<br />"
				noteafter += "REASON : " + msgrevisi
				Insert_log(company, agent, "PERIODE", "REVISI INVOICE", "", noteafter)
			} else {
				flag = false
				log.Println("Update tbl_trx_keluarantogel failed")
			}
			if flag {
				//UPDATE CHILD
				stmt_keluarantogeldetail, e_detail := con.PrepareContext(ctx, `
					UPDATE 
					`+tbl_trx_keluarantogel_detail+`   
					SET statuskeluarandetail=?, winhasil=?, cancelbet=?,  
					updatekeluarandetail=?, updatedatekeluarandetail=? 
					WHERE idtrxkeluaran=? AND idcompany=? 
				`)

				helpers.ErrorCheck(e_detail)
				rec_keluarantogeldetail, e_keluarantogeldetail := stmt_keluarantogeldetail.ExecContext(ctx,
					"RUNNING", 0, 0,
					agent, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
					idtrxkeluaran, company)
				helpers.ErrorCheck(e_keluarantogeldetail)

				affect_keluarantogeldetail, err_affer_keluarantogeldetail := rec_keluarantogeldetail.RowsAffected()
				helpers.ErrorCheck(err_affer_keluarantogeldetail)

				defer stmt_keluarantogeldetail.Close()
				if affect_keluarantogeldetail > 0 {
					flag = true
					log.Printf("Update Parent tbl_trx_keluarantogel_detail : %d\n", idtrxkeluaran)
				} else {
					flag = false
					log.Println("Update tbl_trx_keluarantogel_detail failed")
				}
			}
			if flag {
				//DELETE MEMBER
				stmt_keluarantogelmember, e_member := con.PrepareContext(ctx, `
					DELETE FROM  
					`+tbl_trx_keluarantogel_member+`   
					WHERE idtrxkeluaran=? AND idcompany=? 
				`)

				helpers.ErrorCheck(e_member)
				rec_keluarantogelmember, e_keluarantogelmember := stmt_keluarantogelmember.ExecContext(ctx, idtrxkeluaran, company)
				helpers.ErrorCheck(e_keluarantogelmember)

				affect_keluarantogelmember, err_affer_keluarantogelmember := rec_keluarantogelmember.RowsAffected()
				helpers.ErrorCheck(err_affer_keluarantogelmember)

				defer stmt_keluarantogelmember.Close()
				if affect_keluarantogelmember > 0 {
					flag = true
					msg = "Success"
					log.Printf("Delete tbl_trx_keluarantogel_member : %d\n", idtrxkeluaran)
				} else {
					flag = false
					log.Println("Delete tbl_trx_keluarantogel_member failed")
				}
			}
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	}

	return res, nil
}
func Cancelbet_Periode(agent, company string, idtrxkeluaran, idtrxkeluarandetail int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	msg := "Failed"
	flag := false
	idcomppasaran := 0
	tbl_trx_keluarantogel, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sql_select := `SELECT 
		idcomppasaran    
		FROM ` + tbl_trx_keluarantogel + `   
		WHERE idcompany = ? 
		AND idtrxkeluaran = ? 
		AND keluarantogel = "" 
		ORDER BY idtrxkeluaran DESC LIMIT 1 
	`
	row_select := con.QueryRowContext(ctx, sql_select, company, idtrxkeluaran)
	switch err_select := row_select.Scan(&idcomppasaran); err_select {
	case sql.ErrNoRows:
		msg = "Cannot Update"
	case nil:
		flag = true
	default:
		helpers.ErrorCheck(err_select)
	}

	if flag {
		stmt_keluarantogeldetail, e := con.PrepareContext(ctx, `
			UPDATE 
			`+tbl_trx_keluarantogel_detail+`   
			SET statuskeluarandetail=?, 
			updatekeluarandetail=?, updatedatekeluarandetail=? 
			WHERE idtrxkeluarandetail =? AND idtrxkeluaran=? AND idcompany=? 
		`)
		helpers.ErrorCheck(e)
		rec_keluarantogeldetail, e_keluarantogeldetail := stmt_keluarantogeldetail.ExecContext(ctx,
			"CANCEL", agent, tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idtrxkeluarandetail, idtrxkeluaran, company)
		helpers.ErrorCheck(e_keluarantogeldetail)

		a_keluarantogeldetail, e_keluarantogel := rec_keluarantogeldetail.RowsAffected()
		helpers.ErrorCheck(e_keluarantogel)

		defer stmt_keluarantogeldetail.Close()
		if a_keluarantogeldetail > 0 {
			flag = true
			msg = "Success"
			log.Printf("Update Detail tbl_trx_keluarantogel_detail : %d\n", idtrxkeluaran)

			idpasarantogel, _ := Pasaran_id(idcomppasaran, company, "idpasarantogel")
			nmpasarantogel := Pasaranmaster_id(idpasarantogel, "nmpasarantogel")
			noteafter := ""
			noteafter += "INVOICE - " + strconv.Itoa(idtrxkeluaran) + "<br />"
			noteafter += "INVOICE BET - " + strconv.Itoa(idtrxkeluarandetail) + "<br />"
			noteafter += "PASARAN : " + nmpasarantogel
			Insert_log(company, agent, "PERIODE", "CANCEL BET", "", noteafter)
		} else {
			flag = false
			log.Println("Update tbl_trx_keluarantogel_detail failed")
		}

	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	}

	return res, nil
}
func _togel_bayar_SUM(idtrxkeluaran int, company string) int {
	con := db.CreateCon()
	ctx := context.Background()

	_, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sql_bayar := `SELECT 
		sum(bet-(bet*diskon)-(bet*kei)) as total  
		FROM ` + tbl_trx_keluarantogel_detail + `   
		WHERE idcompany = ? 
		AND idtrxkeluaran = ? 
	`
	var total_db sql.NullFloat64
	total := 0
	row := con.QueryRowContext(ctx, sql_bayar, company, idtrxkeluaran)
	switch e := row.Scan(&total_db); e {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		log.Println(total_db)
	default:
		log.Panic(e)
	}
	if total_db.Valid {
		total = int(total_db.Float64)
	}
	return total
}
func _togel_bet_SUM_RUNNING(idtrxkeluaran int, company string) int {
	con := db.CreateCon()
	ctx := context.Background()

	_, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sql_bet := `SELECT 
		count(idtrxkeluarandetail) as total  
		FROM ` + tbl_trx_keluarantogel_detail + `   
		WHERE idcompany = ? 
		AND idtrxkeluaran = ? 
		AND statuskeluarandetail = "RUNNING" 
	`
	var total_db sql.NullInt32
	total := 0

	row := con.QueryRowContext(ctx, sql_bet, company, idtrxkeluaran)
	switch e := row.Scan(&total_db); e {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		log.Println(total_db)
	default:
		log.Panic(e)
	}
	if total_db.Valid {
		total = int(total_db.Int32)
	}
	return total
}
func _togel_bet_SUM(idtrxkeluaran int, company string) int {
	con := db.CreateCon()
	ctx := context.Background()

	_, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sql_bet := `SELECT 
		count(idtrxkeluarandetail) as total  
		FROM ` + tbl_trx_keluarantogel_detail + `   
		WHERE idcompany = ? 
		AND idtrxkeluaran = ? 
	`
	var total_db sql.NullInt32
	total := 0

	row := con.QueryRowContext(ctx, sql_bet, company, idtrxkeluaran)
	switch e := row.Scan(&total_db); e {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		log.Println(total_db)
	default:
		log.Panic(e)
	}
	if total_db.Valid {
		total = int(total_db.Int32)
	}
	return total
}
func _togel_member_COUNT(idtrxkeluaran int, company string) int {
	con := db.CreateCon()
	ctx := context.Background()
	total := 0

	_, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	sql := `SELECT 
		username
		FROM ` + tbl_trx_keluarantogel_detail + `  
		WHERE idcompany = ? 
		AND idtrxkeluaran = ? 
		GROUP BY username 
	`

	row, err := con.QueryContext(ctx, sql, company, idtrxkeluaran)

	helpers.ErrorCheck(err)
	for row.Next() {
		total = total + 1
		var username_db string

		err = row.Scan(
			&username_db)

		helpers.ErrorCheck(err)

	}
	defer row.Close()

	return total
}

func _rumuswinhasil(permainan string, bayar int, bet int, win float32) int {
	winhasil := 0
	if permainan == "50_50_UMUM" || permainan == "50_50_SPECIAL" ||
		permainan == "50_50_KOMBINASI" || permainan == "DASAR" || permainan == "COLOK_BEBAS" ||
		permainan == "COLOK_MACAU" || permainan == "COLOK_NAGA" || permainan == "COLOK_JITU" {

		winhasil = bayar + int(float32(bet)*win)
	} else {
		winhasil = int(float32(bet) * win)
	}
	return winhasil
}
func _rumusTogel(angka, tipe, nomorkeluaran, posisitogel, company, simpandb string, idcomppasaran, idtrxkeluarandetail int) (string, float32) {
	var result string = "LOSE"
	var win float32 = 0

	_, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)

	temp := angka
	temp4d := string([]byte(temp)[0]) + string([]byte(temp)[1]) + string([]byte(temp)[2]) + string([]byte(temp)[3])
	temp3d := string([]byte(temp)[1]) + string([]byte(temp)[2]) + string([]byte(temp)[3])
	temp3dd := string([]byte(temp)[0]) + string([]byte(temp)[1]) + string([]byte(temp)[2])
	temp2d := string([]byte(temp)[2]) + string([]byte(temp)[3])
	temp2dd := string([]byte(temp)[0]) + string([]byte(temp)[1])
	temp2dt := string([]byte(temp)[1]) + string([]byte(temp)[2])

	var temp4d_arr []string
	var temp3d_arr []string
	var temp3dd_arr []string
	var temp2d_arr []string
	var temp2dd_arr []string
	var temp2dt_arr []string

	switch tipe {
	case "4D":
		if temp4d == nomorkeluaran {
			result = "WINNER"
		} else {
			if posisitogel == "BB" {
				flag_bb_4D := false
				for a, _ := range temp4d {
					for b, _ := range temp4d {
						for c, _ := range temp4d {
							for d, _ := range temp4d {
								if a != b && a != c && a != d {
									if b != c && b != d {
										if c != d {
											temp_loop := string([]byte(temp4d)[a]) + string([]byte(temp4d)[b]) + string([]byte(temp4d)[c]) + string([]byte(temp4d)[d])
											if temp4d != temp_loop {
												temp4d_arr = append(temp4d_arr, temp_loop)
											}
											temp_loop = ""
										}
									}
								}
							}
						}
					}
				}
				removeDuplicateValuesSlice := _removeDuplicateValues(temp4d_arr)
				for a, _ := range removeDuplicateValuesSlice {
					if removeDuplicateValuesSlice[a] == nomorkeluaran {
						result = "WINNER"
						flag_bb_4D = true
					}

				}
				if flag_bb_4D {
					_, win_db := Pasaran_id(idcomppasaran, company, "1_win4dbb")
					win = win_db
					if simpandb == "Y" {
						_updatevaluewinbytipe(tbl_trx_keluarantogel_detail, win_db, idtrxkeluarandetail)
					}
				}
			}
		}
	case "3D":
		if temp3d == nomorkeluaran {
			result = "WINNER"
		} else {
			if posisitogel == "BB" {
				flag_bb_3D := false
				for a, _ := range temp3d {
					for b, _ := range temp3d {
						for c, _ := range temp3d {
							if a != b && a != c {
								if b != c {
									temp_loop := string([]byte(temp3d)[a]) + string([]byte(temp3d)[b]) + string([]byte(temp3d)[c])
									if temp3d != temp_loop {
										temp3d_arr = append(temp3d_arr, temp_loop)
									}
									temp_loop = ""
								}
							}
						}
					}
				}
				removeDuplicateValuesSlice := _removeDuplicateValues(temp3d_arr)
				for a, _ := range removeDuplicateValuesSlice {
					if removeDuplicateValuesSlice[a] == nomorkeluaran {
						result = "WINNER"
						flag_bb_3D = true
					}

				}
				if flag_bb_3D {
					_, win_db := Pasaran_id(idcomppasaran, company, "1_win3dbb")
					win = win_db
					if simpandb == "Y" {
						_updatevaluewinbytipe(tbl_trx_keluarantogel_detail, win_db, idtrxkeluarandetail)
					}
				}
			}
		}
	case "3DD":
		if temp3dd == nomorkeluaran {
			result = "WINNER"
		} else {
			if posisitogel == "BB" {
				flag_bb_3DD := false
				for a, _ := range temp3dd {
					for b, _ := range temp3dd {
						for c, _ := range temp3dd {
							if a != b && a != c {
								if b != c {
									temp_loop := string([]byte(temp3dd)[a]) + string([]byte(temp3dd)[b]) + string([]byte(temp3dd)[c])
									if temp3dd != temp_loop {
										temp3dd_arr = append(temp3dd_arr, temp_loop)
									}
									temp_loop = ""
								}
							}
						}
					}
				}
				for a, _ := range temp3dd_arr {
					if temp3dd_arr[a] == nomorkeluaran {
						result = "WINNER"
						flag_bb_3DD = true
					}

				}
				if flag_bb_3DD {
					_, win_db := Pasaran_id(idcomppasaran, company, "1_win3ddbb")
					win = win_db
					if simpandb == "Y" {
						_updatevaluewinbytipe(tbl_trx_keluarantogel_detail, win_db, idtrxkeluarandetail)
					}
				}
			}
		}
	case "2D":
		if temp2d == nomorkeluaran {
			result = "WINNER"
		} else {
			if posisitogel == "BB" {
				flag_bb_2D := false
				for a, _ := range temp2d {
					for b, _ := range temp2d {
						if a != b {
							temp_loop := string([]byte(temp2d)[a]) + string([]byte(temp2d)[b])
							if temp2d != temp_loop {
								temp2d_arr = append(temp2d_arr, temp_loop)
							}
							temp_loop = ""
						}
					}
				}
				removeDuplicateValuesSlice := _removeDuplicateValues(temp2d_arr)
				for a, _ := range removeDuplicateValuesSlice {
					if removeDuplicateValuesSlice[a] == nomorkeluaran {
						result = "WINNER"
						flag_bb_2D = true
					}

				}
				if flag_bb_2D {
					_, win_db := Pasaran_id(idcomppasaran, company, "1_win2dbb")
					win = win_db
					if simpandb == "Y" {
						_updatevaluewinbytipe(tbl_trx_keluarantogel_detail, win_db, idtrxkeluarandetail)
					}
				}
			}
		}
	case "2DD":
		if temp2dd == nomorkeluaran {
			result = "WINNER"
		} else {
			if posisitogel == "BB" {
				flag_bb_2DD := false
				for a, _ := range temp2dd {
					for b, _ := range temp2dd {
						if a != b {
							temp_loop := string([]byte(temp2dd)[a]) + string([]byte(temp2dd)[b])
							if temp2dd != temp_loop {
								temp2dd_arr = append(temp2dd_arr, temp_loop)
							}
							temp_loop = ""
						}
					}
				}
				removeDuplicateValuesSlice := _removeDuplicateValues(temp2dd_arr)
				for a, _ := range removeDuplicateValuesSlice {
					if removeDuplicateValuesSlice[a] == nomorkeluaran {
						result = "WINNER"
						flag_bb_2DD = true
					}

				}
				if flag_bb_2DD {
					_, win_db := Pasaran_id(idcomppasaran, company, "1_win2ddbb")
					win = win_db
					if simpandb == "Y" {
						_updatevaluewinbytipe(tbl_trx_keluarantogel_detail, win_db, idtrxkeluarandetail)
					}
				}
			}
		}
	case "2DT":
		if temp2dt == nomorkeluaran {
			result = "WINNER"
		} else {
			if posisitogel == "BB" {
				flag_bb_2DT := false
				for a, _ := range temp2dt {
					for b, _ := range temp2dt {
						if a != b {
							temp_loop := string([]byte(temp2dt)[a]) + string([]byte(temp2dt)[b])
							if temp2dt != temp_loop {
								temp2dt_arr = append(temp2dt_arr, temp_loop)
							}
							temp_loop = ""
						}
					}
				}
				removeDuplicateValuesSlice := _removeDuplicateValues(temp2dt_arr)
				for a, _ := range removeDuplicateValuesSlice {
					if removeDuplicateValuesSlice[a] == nomorkeluaran {
						result = "WINNER"
						flag_bb_2DT = true
					}

				}
				if flag_bb_2DT {
					_, win_db := Pasaran_id(idcomppasaran, company, "1_win2dtbb")
					win = win_db
					if simpandb == "Y" {
						_updatevaluewinbytipe(tbl_trx_keluarantogel_detail, win_db, idtrxkeluarandetail)
					}
				}
			}
		}
	case "COLOK_BEBAS":
		flag := false
		count := 0
		for i := 0; i < len(temp); i++ {
			if string([]byte(temp)[i]) == nomorkeluaran {
				flag = true
				count = count + 1
			}
		}
		if flag {
			_, win_db := Pasaran_id(idcomppasaran, company, "2_win")
			if count == 1 {
				win = win_db
			}
			if count == 2 {
				win = win_db * 2
			}
			if count == 3 {
				win = win_db * 3
			}
			if count == 4 {
				win = win_db * 3
			}
			fmt.Println(win)

			if simpandb == "Y" {
				//UPDATE WIN DETAIL BET
				_updatevaluewinbytipe(tbl_trx_keluarantogel_detail, win, idtrxkeluarandetail)
			}
			result = "WINNER"
		}
	case "COLOK_MACAU":
		flag_1 := false
		flag_2 := false
		count_1 := 0
		count_2 := 0
		totalcount := 0
		var win float32 = 0
		for i := 0; i < len(temp); i++ {
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[0]) {
				flag_1 = true
				count_1 = count_1 + 1
			}
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[1]) {
				flag_2 = true
				count_2 = count_2 + 1
			}
		}
		if flag_1 && flag_2 {
			totalcount = count_1 + count_2
			if totalcount == 2 {
				_, win = Pasaran_id(idcomppasaran, company, "3_win2digit")
			}
			if totalcount == 3 {
				_, win = Pasaran_id(idcomppasaran, company, "3_win3digit")
			}
			if totalcount == 4 {
				_, win = Pasaran_id(idcomppasaran, company, "3_win4digit")
			}
			if simpandb == "Y" {
				//UPDATE WIN DETAIL BET
				_updatevaluewinbytipe(tbl_trx_keluarantogel_detail, win, idtrxkeluarandetail)
			}
			result = "WINNER"
		}
	case "COLOK_NAGA":
		flag_1 := false
		flag_2 := false
		flag_3 := false
		count_1 := 0
		count_2 := 0
		count_3 := 0
		totalcount := 0
		var win float32 = 0
		for i := 0; i < len(temp); i++ {
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[0]) {
				flag_1 = true
				count_1 = count_1 + 1
			}
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[1]) {
				flag_2 = true
				count_2 = count_2 + 1
			}
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[2]) {
				flag_3 = true
				count_3 = count_3 + 1
			}
		}
		if flag_1 && flag_2 {
			if flag_3 {
				totalcount = count_1 + count_2 + count_3
				log.Println("Total Count Colok Naga :", totalcount)

				if totalcount == 3 {
					_, win = Pasaran_id(idcomppasaran, company, "4_win3digit")
				}
				if totalcount == 4 {
					_, win = Pasaran_id(idcomppasaran, company, "4_win4digit")
				}
				log.Println("WIN COLOK NAGA :", win)
				if simpandb == "Y" {
					//UPDATE WIN DETAIL BET
					_updatevaluewinbytipe(tbl_trx_keluarantogel_detail, win, idtrxkeluarandetail)
				}
				result = "WINNER"
			}
		}
	case "COLOK_JITU":
		flag := false
		as := string([]byte(temp)[0]) + "_AS"
		kop := string([]byte(temp)[1]) + "_KOP"
		kepala := string([]byte(temp)[2]) + "_KEPALA"
		ekor := string([]byte(temp)[3]) + "_EKOR"

		if as == nomorkeluaran {
			flag = true
		}
		if kop == nomorkeluaran {
			flag = true
		}
		if kepala == nomorkeluaran {
			flag = true
		}
		if ekor == nomorkeluaran {
			flag = true
		}
		if flag {
			result = "WINNER"
		}
	case "50_50_UMUM":
		flag := false
		data := []string{}
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])
		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)
		dasar, _ := strconv.Atoi(kepala + ekor)
		//BESARKECIL
		if kepala_2 <= 4 {
			data = append(data, "KECIL")
		} else {
			data = append(data, "BESAR")
		}
		//GENAPGANJIL
		if ekor_2%2 == 0 {
			data = append(data, "GENAP")
		} else {
			data = append(data, "GANJIL")
		}
		log.Printf("DASAR : %d", dasar)
		//TEPITENGAH
		if dasar >= 0 && dasar <= 24 {
			data = append(data, "TEPI")
		}
		if dasar >= 25 && dasar <= 74 {
			data = append(data, "TENGAH")
		}
		if dasar >= 75 && dasar <= 99 {
			data = append(data, "TEPI")
		}
		for i := 0; i < len(data); i++ {
			if data[i] == nomorkeluaran {
				flag = true
			}
		}
		if flag {
			result = "WINNER"
		}
		fmt.Println(data)
	case "50_50_SPECIAL":
		flag := false
		as := string([]byte(temp)[0])
		kop := string([]byte(temp)[1])
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])

		as_2, _ := strconv.Atoi(as)
		kop_2, _ := strconv.Atoi(kop)
		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)
		//AS - BESARKECIL == GENAPGANJIL
		if as_2 <= 4 {
			if nomorkeluaran == "AS_KECIL" {
				flag = true
			}
		} else {
			if nomorkeluaran == "AS_BESAR" {
				flag = true
			}
		}
		if as_2%2 == 0 {
			if nomorkeluaran == "AS_GENAP" {
				flag = true
			}
		} else {
			if nomorkeluaran == "AS_GANJIL" {
				flag = true
			}
		}

		//KOP - BESARKECIL == GENAPGANJIL
		if kop_2 <= 4 {
			if nomorkeluaran == "KOP_KECIL" {
				flag = true
			}
		} else {
			if nomorkeluaran == "KOP_BESAR" {
				flag = true
			}
		}
		if kop_2%2 == 0 {
			if nomorkeluaran == "KOP_GENAP" {
				flag = true
			}
		} else {
			if nomorkeluaran == "KOP_GANJIL" {
				flag = true
			}
		}

		//KEPALA - BESARKECIL == GENAPGANJIL
		if kepala_2 <= 4 {
			if nomorkeluaran == "KEPALA_KECIL" {
				flag = true
			}
		} else {
			if nomorkeluaran == "KEPALA_BESAR" {
				flag = true
			}
		}
		if kepala_2%2 == 0 {
			if nomorkeluaran == "KEPALA_GENAP" {
				flag = true
			}
		} else {
			if nomorkeluaran == "KEPALA_GANJIL" {
				flag = true
			}
		}

		//EKOR - BESARKECIL == GENAPGANJIL
		if ekor_2 <= 4 {
			if nomorkeluaran == "EKOR_KECIL" {
				flag = true
			}
		} else {
			if nomorkeluaran == "EKOR_BESAR" {
				flag = true
			}
		}
		if ekor_2%2 == 0 {
			if nomorkeluaran == "EKOR_GENAP" {
				flag = true
			}
		} else {
			if nomorkeluaran == "EKOR_GANJIL" {
				flag = true
			}
		}

		if flag {
			result = "WINNER"
		}
	case "50_50_KOMBINASI":
		flag := false
		data_1 := ""
		data_2 := ""
		data_3 := ""
		data_4 := ""
		depan := ""
		tengah := ""
		belakang := ""
		depan_1 := ""
		tengah_1 := ""
		belakang_1 := ""
		as := string([]byte(temp)[0])
		kop := string([]byte(temp)[1])
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])

		as_2, _ := strconv.Atoi(as)
		kop_2, _ := strconv.Atoi(kop)
		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)

		if as_2%2 == 0 {
			data_1 = "GENAP"
		} else {
			data_1 = "GANJIL"
		}
		if kop_2%2 == 0 {
			data_2 = "GENAP"
		} else {
			data_2 = "GANJIL"
		}
		if kepala_2%2 == 0 {
			data_3 = "GENAP"
		} else {
			data_3 = "GANJIL"
		}
		if ekor_2%2 == 0 {
			data_4 = "GENAP"
		} else {
			data_4 = "GANJIL"
		}
		depan = data_1 + "-" + data_2
		tengah = data_2 + "-" + data_3
		belakang = data_3 + "-" + data_4

		if depan == "GENAP-GANJIL" || depan == "GANJIL-GENAP" {
			depan = "DEPAN_STEREO"
		} else {
			depan = "DEPAN_MONO"
		}
		if tengah == "GENAP-GANJIL" || tengah == "GANJIL-GENAP" {
			tengah = "TENGAH_STEREO"
		} else {
			tengah = "TENGAH_MONO"
		}
		if belakang == "GENAP-GANJIL" || belakang == "GANJIL-GENAP" {
			belakang = "BELAKANG_STEREO"
		} else {
			belakang = "BELAKANG_MONO"
		}
		if as_2 < kop_2 {
			depan_1 = "DEPAN_KEMBANG"
		}
		if as_2 > kop_2 {
			depan_1 = "DEPAN_KEMPIS"
		}
		if as_2 == kop_2 {
			depan_1 = "DEPAN_KEMBAR"
		}
		if kop_2 < kepala_2 {
			tengah_1 = "TENGAH_KEMBANG"
		}
		if kop_2 > kepala_2 {
			tengah_1 = "TENGAH_KEMPIS"
		}
		if kop_2 == kepala_2 {
			tengah_1 = "TENGAH_KEMBAR"
		}
		if kepala_2 < ekor_2 {
			belakang_1 = "BELAKANG_KEMBANG"
		}
		if kepala_2 > ekor_2 {
			belakang_1 = "BELAKANG_KEMPIS"
		}
		if kepala_2 == ekor_2 {
			belakang_1 = "BELAKANG_KEMBAR"
		}

		if depan == nomorkeluaran {
			flag = true
		}
		if tengah == nomorkeluaran {
			flag = true
		}
		if belakang == nomorkeluaran {
			flag = true
		}
		if depan_1 == nomorkeluaran {
			flag = true
		}
		if tengah_1 == nomorkeluaran {
			flag = true
		}
		if belakang_1 == nomorkeluaran {
			flag = true
		}

		if flag {
			result = "WINNER"
		}
	case "MACAU_KOMBINASI":
		flag := false
		data_1 := ""
		data_2 := ""
		data_3 := ""
		data_4 := ""
		depan := ""
		tengah := ""
		tengah2 := ""
		belakang := ""

		as := string([]byte(temp)[0])
		kop := string([]byte(temp)[1])
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])

		as_2, _ := strconv.Atoi(as)
		kop_2, _ := strconv.Atoi(kop)
		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)

		if as_2 <= 4 {
			data_1 = "KECIL"
		} else {
			data_1 = "BESAR"
		}
		if kop_2%2 == 0 {
			data_2 = "GENAP"
		} else {
			data_2 = "GANJIL"
		}
		if kepala_2 <= 4 {
			data_3 = "KECIL"
		} else {
			data_3 = "BESAR"
		}
		if ekor_2%2 == 0 {
			data_4 = "GENAP"
		} else {
			data_4 = "GANJIL"
		}

		depan = "DEPAN_" + data_1 + "_" + data_2
		tengah = "TENGAH_" + data_2 + "_" + data_3
		tengah2 = "TENGAH_" + data_3 + "_" + data_2
		belakang = "BELAKANG_" + data_3 + "_" + data_4

		if depan == nomorkeluaran {
			flag = true
		}
		if tengah == nomorkeluaran {
			flag = true
		}
		if tengah2 == nomorkeluaran {
			flag = true
		}
		if belakang == nomorkeluaran {
			flag = true
		}

		if flag {
			result = "WINNER"
		}
	case "DASAR":
		flag := false
		data_1 := ""
		data_2 := ""

		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])

		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)

		dasar := kepala_2 + ekor_2

		if dasar > 9 {
			temp2 := strconv.Itoa(dasar) //int to string
			temp21 := string([]byte(temp2)[0])
			temp22 := string([]byte(temp2)[1])

			temp21_2, _ := strconv.Atoi(temp21)
			temp22_2, _ := strconv.Atoi(temp22)
			dasar = temp21_2 + temp22_2
		}
		if dasar <= 4 {
			data_1 = "KECIL"
		} else {
			data_1 = "BESAR"
		}
		if dasar%2 == 0 {
			data_2 = "GENAP"
		} else {
			data_2 = "GANJIL"
		}

		if data_1 == nomorkeluaran {
			flag = true
		}
		if data_2 == nomorkeluaran {
			flag = true
		}

		if flag {
			result = "WINNER"
		}
	case "SHIO":
		flag := false

		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])
		data := _tableshio(kepala + ekor)

		if data == nomorkeluaran {
			flag = true
		}

		if flag {
			result = "WINNER"
		}
	}
	return result, win
}
func _removeDuplicateValues(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func _updatevaluewinbytipe(nmtable string, win float32, idtrxkeluarandetail int) {
	tglnow, _ := goment.New()
	sql_update := `
		UPDATE 
		` + nmtable + `     
		SET win=?, 
		updatekeluarandetail=?, updatedatekeluarandetail=? 
		WHERE idtrxkeluarandetail=? 
	`
	flag_update, msg_update := Exec_SQL(sql_update, nmtable, "UPDATE",
		win,
		"SYSTEM",
		tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idtrxkeluarandetail)
	if !flag_update {
		log.Println(msg_update)
	}
}
func _tableshio(shiodata string) string {
	log.Printf("Shio : %s", shiodata)

	tglnow, _ := goment.New()
	yearnow := tglnow.Format("YYYY")
	log.Println(yearnow)
	result := ""
	switch yearnow {
	case "2022":
		harimau := []string{"01", "13", "25", "37", "49", "61", "73", "85", "97"}
		kerbau := []string{"02", "14", "26", "38", "50", "62", "74", "86", "98"}
		tikus := []string{"03", "15", "27", "39", "51", "63", "75", "87", "99"}
		babi := []string{"04", "16", "28", "40", "52", "64", "76", "88", "00"}
		anjing := []string{"05", "17", "29", "41", "53", "65", "77", "89", ""}
		ayam := []string{"06", "18", "30", "42", "54", "66", "78", "90", ""}
		monyet := []string{"07", "19", "31", "43", "55", "67", "79", "91", ""}
		kambing := []string{"08", "20", "32", "44", "56", "68", "80", "92", ""}
		kuda := []string{"09", "21", "33", "45", "57", "69", "81", "93", ""}
		ular := []string{"10", "22", "34", "46", "58", "70", "82", "94", ""}
		naga := []string{"11", "23", "35", "47", "59", "71", "83", "95", ""}
		kelinci := []string{"12", "24", "36", "48", "60", "72", "84", "96", ""}
		for i := 0; i < len(babi); i++ {
			if shiodata == babi[i] {
				result = "BABI"
			}
		}
		for i := 0; i < len(ular); i++ {
			if shiodata == ular[i] {
				result = "ULAR"
			}
		}
		for i := 0; i < len(anjing); i++ {
			if shiodata == anjing[i] {
				result = "ANJING"
			}
		}
		for i := 0; i < len(ayam); i++ {
			if shiodata == ayam[i] {
				result = "AYAM"
			}
		}
		for i := 0; i < len(monyet); i++ {
			if shiodata == monyet[i] {
				result = "MONYET"
			}
		}
		for i := 0; i < len(kambing); i++ {
			if shiodata == kambing[i] {
				result = "KAMBING"
			}
		}
		for i := 0; i < len(kuda); i++ {
			if shiodata == kuda[i] {
				result = "KUDA"
			}
		}
		for i := 0; i < len(naga); i++ {
			if shiodata == naga[i] {
				result = "NAGA"
			}
		}
		for i := 0; i < len(kelinci); i++ {
			if shiodata == kelinci[i] {
				result = "KELINCI"
			}
		}
		for i := 0; i < len(harimau); i++ {
			if shiodata == harimau[i] {
				result = "HARIMAU"
			}
		}
		for i := 0; i < len(kerbau); i++ {
			if shiodata == kerbau[i] {
				result = "KERBAU"
			}
		}
		for i := 0; i < len(tikus); i++ {
			if shiodata == tikus[i] {
				result = "TIKUS"
			}
		}
	case "2023":
		kelinci := []string{"01", "13", "25", "37", "49", "61", "73", "85", "97"}
		harimau := []string{"02", "14", "26", "38", "50", "62", "74", "86", "98"}
		kerbau := []string{"03", "15", "27", "39", "51", "63", "75", "87", "99"}
		tikus := []string{"04", "16", "28", "40", "52", "64", "76", "88", "00"}
		babi := []string{"05", "17", "29", "41", "53", "65", "77", "89", ""}
		anjing := []string{"06", "18", "30", "42", "54", "66", "78", "90", ""}
		ayam := []string{"07", "19", "31", "43", "55", "67", "79", "91", ""}
		monyet := []string{"08", "20", "32", "44", "56", "68", "80", "92", ""}
		kambing := []string{"09", "21", "33", "45", "57", "69", "81", "93", ""}
		kuda := []string{"10", "22", "34", "46", "58", "70", "82", "94", ""}
		ular := []string{"11", "23", "35", "47", "59", "71", "83", "95", ""}
		naga := []string{"12", "24", "36", "48", "60", "72", "84", "96", ""}
		for i := 0; i < len(babi); i++ {
			if shiodata == babi[i] {
				result = "BABI"
			}
		}
		for i := 0; i < len(ular); i++ {
			if shiodata == ular[i] {
				result = "ULAR"
			}
		}
		for i := 0; i < len(anjing); i++ {
			if shiodata == anjing[i] {
				result = "ANJING"
			}
		}
		for i := 0; i < len(ayam); i++ {
			if shiodata == ayam[i] {
				result = "AYAM"
			}
		}
		for i := 0; i < len(monyet); i++ {
			if shiodata == monyet[i] {
				result = "MONYET"
			}
		}
		for i := 0; i < len(kambing); i++ {
			if shiodata == kambing[i] {
				result = "KAMBING"
			}
		}
		for i := 0; i < len(kuda); i++ {
			if shiodata == kuda[i] {
				result = "KUDA"
			}
		}
		for i := 0; i < len(naga); i++ {
			if shiodata == naga[i] {
				result = "NAGA"
			}
		}
		for i := 0; i < len(kelinci); i++ {
			if shiodata == kelinci[i] {
				result = "KELINCI"
			}
		}
		for i := 0; i < len(harimau); i++ {
			if shiodata == harimau[i] {
				result = "HARIMAU"
			}
		}
		for i := 0; i < len(kerbau); i++ {
			if shiodata == kerbau[i] {
				result = "KERBAU"
			}
		}
		for i := 0; i < len(tikus); i++ {
			if shiodata == tikus[i] {
				result = "TIKUS"
			}
		}
	case "2024":
		naga := []string{"01", "13", "25", "37", "49", "61", "73", "85", "97"}
		kelinci := []string{"02", "14", "26", "38", "50", "62", "74", "86", "98"}
		harimau := []string{"03", "15", "27", "39", "51", "63", "75", "87", "99"}
		kerbau := []string{"04", "16", "28", "40", "52", "64", "76", "88", "00"}
		tikus := []string{"05", "17", "29", "41", "53", "65", "77", "89", ""}
		babi := []string{"06", "18", "30", "42", "54", "66", "78", "90", ""}
		anjing := []string{"07", "19", "31", "43", "55", "67", "79", "91", ""}
		ayam := []string{"08", "20", "32", "44", "56", "68", "80", "92", ""}
		monyet := []string{"09", "21", "33", "45", "57", "69", "81", "93", ""}
		kambing := []string{"10", "22", "34", "46", "58", "70", "82", "94", ""}
		kuda := []string{"11", "23", "35", "47", "59", "71", "83", "95", ""}
		ular := []string{"12", "24", "36", "48", "60", "72", "84", "96", ""}
		for i := 0; i < len(babi); i++ {
			if shiodata == babi[i] {
				result = "BABI"
			}
		}
		for i := 0; i < len(ular); i++ {
			if shiodata == ular[i] {
				result = "ULAR"
			}
		}
		for i := 0; i < len(anjing); i++ {
			if shiodata == anjing[i] {
				result = "ANJING"
			}
		}
		for i := 0; i < len(ayam); i++ {
			if shiodata == ayam[i] {
				result = "AYAM"
			}
		}
		for i := 0; i < len(monyet); i++ {
			if shiodata == monyet[i] {
				result = "MONYET"
			}
		}
		for i := 0; i < len(kambing); i++ {
			if shiodata == kambing[i] {
				result = "KAMBING"
			}
		}
		for i := 0; i < len(kuda); i++ {
			if shiodata == kuda[i] {
				result = "KUDA"
			}
		}
		for i := 0; i < len(naga); i++ {
			if shiodata == naga[i] {
				result = "NAGA"
			}
		}
		for i := 0; i < len(kelinci); i++ {
			if shiodata == kelinci[i] {
				result = "KELINCI"
			}
		}
		for i := 0; i < len(harimau); i++ {
			if shiodata == harimau[i] {
				result = "HARIMAU"
			}
		}
		for i := 0; i < len(kerbau); i++ {
			if shiodata == kerbau[i] {
				result = "KERBAU"
			}
		}
		for i := 0; i < len(tikus); i++ {
			if shiodata == tikus[i] {
				result = "TIKUS"
			}
		}
	}

	return result
}
func Pasaran_id(idcomppasaran int, company, tipecolumn string) (string, float32) {
	con := db.CreateCon()
	ctx := context.Background()
	var result string = ""
	var result_number float32 = 0
	sql_pasaran := `SELECT 
		idpasarantogel , 
		1_win4dbb, 1_win3dbb, 1_win3ddbb, 1_win2dbb, 1_win2ddbb, 1_win2dtbb, 
		2_win as win_cbebas, 3_win2digit as win2_cmacau, 
		3_win3digit as win3_cmacau, 3_win4digit as win4_cmacau, 
		4_win3digit as win3_cnaga, 4_win4digit as win4_cnaga 
		FROM ` + config.DB_tbl_mst_company_game_pasaran + `  
		WHERE idcomppasaran  = ? 
		AND idcompany = ? 
	`
	var (
		idpasarantogel_db                                                         string
		win4dbb_db, win3dbb_db, win3ddbb_db, win2dbb_db, win2ddbb_db, win2dtbb_db float32
		win_cbebas_db, win2_cmacau_db, win3_cmacau_db, win4_cmacau_db             float32
		win3_cnaga_db, win4_cnaga_db                                              float32
	)
	rows := con.QueryRowContext(ctx, sql_pasaran, idcomppasaran, company)
	switch err := rows.Scan(
		&idpasarantogel_db,
		&win4dbb_db, &win3dbb_db, &win3ddbb_db, &win2dbb_db, &win2ddbb_db, &win2dtbb_db,
		&win_cbebas_db, &win2_cmacau_db, &win3_cmacau_db, &win4_cmacau_db,
		&win3_cnaga_db, &win4_cnaga_db); err {
	case sql.ErrNoRows:
		result = ""
	case nil:
		switch tipecolumn {
		case "idpasarantogel":
			result = idpasarantogel_db
		case "1_win4dbb":
			result_number = win4dbb_db
		case "1_win3dbb":
			result_number = win3dbb_db
		case "1_win3ddbb":
			result_number = win3ddbb_db
		case "1_win2dbb":
			result_number = win2dbb_db
		case "1_win2ddbb":
			result_number = win2ddbb_db
		case "1_win2dtbb":
			result_number = win2dtbb_db
		case "2_win":
			result_number = win_cbebas_db
		case "3_win2digit":
			result_number = win2_cmacau_db
		case "3_win3digit":
			result_number = win3_cmacau_db
		case "3_win4digit":
			result_number = win4_cmacau_db
		case "4_win3digit":
			result_number = win3_cnaga_db
		case "4_win4digit":
			result_number = win4_cnaga_db
		}
	default:
		helpers.ErrorCheck(err)
	}
	return result, result_number
}
func Pasaranmaster_id(pasarancode, tipecolumn string) string {
	con := db.CreateCon()
	ctx := context.Background()
	result := ""
	sql_pasaran := `SELECT 
		tipepasaran,nmpasarantogel 
		FROM ` + config.DB_tbl_mst_pasaran_togel + `  
		WHERE idpasarantogel = ? 
	`
	var (
		tipepasaran_db, nmpasarantogel_db string
	)
	rows := con.QueryRowContext(ctx, sql_pasaran, pasarancode)
	switch err := rows.Scan(&tipepasaran_db, &nmpasarantogel_db); err {
	case sql.ErrNoRows:

	case nil:
		switch tipecolumn {
		case "tipepasaran":
			result = tipepasaran_db
		case "nmpasarantogel":
			result = nmpasarantogel_db
		}
	default:
		helpers.ErrorCheck(err)
	}
	return result
}

func _checkpasaranonline(idcomppasaran int, company string) bool {
	var myDays = []string{"minggu", "senin", "selasa", "rabu", "kamis", "jumat", "sabtu"}
	flag := false

	con := db.CreateCon()
	ctx := context.Background()

	tglnow, _ := goment.New()
	daynow := tglnow.Format("d")
	intVar, _ := strconv.ParseInt(daynow, 0, 8)
	daynowhari := myDays[intVar]
	haripasaran := ""
	sqlpasaranonline := `
			SELECT
				haripasaran
			FROM ` + config.DB_tbl_mst_company_game_pasaran_offline + ` 
			WHERE idcomppasaran = ?
			AND idcompany = ? 
			AND haripasaran = ? 
		`

	errpasaranonline := con.QueryRowContext(ctx, sqlpasaranonline, idcomppasaran, company, daynowhari).Scan(&haripasaran)

	if errpasaranonline != sql.ErrNoRows {
		flag = true
	}
	log.Println(flag)
	return flag
}
