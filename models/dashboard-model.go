package models

import (
	"context"
	"database/sql"
	"log"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/config"
	"bitbucket.org/isbtotogroup/apibackend_go/db"
	"bitbucket.org/isbtotogroup/apibackend_go/entities"
	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

type dashboardparent struct {
	Pasaranname   string      `json:"pasaran_name"`
	Pasarandetail interface{} `json:"pasaran_detial"`
}
type dashboarddetail struct {
	Pasaranwinlose int
}

func Fetch_dashboardwinlose(company, year string) (helpers.Response, error) {
	var obj entities.Model_dashboardwinlose
	var arraobj []entities.Model_dashboardwinlose
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	log.Println(company)
	log.Println(year)
	sql_periode := `SELECT 
			winlosecomp 
			FROM ` + config.DB_tbl_trx_company_invoice + ` 
			WHERE yearinvoice = ? 
			AND idcompany = ?  
			ORDER BY periodeinvoice ASC  
		`
	row, err := con.QueryContext(ctx, sql_periode, year, company)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			winlosecomp_db int
		)
		err = row.Scan(&winlosecomp_db)
		helpers.ErrorCheck(err)
		obj.Winlose = winlosecomp_db
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

func Fetch_dashboard(company string) (helpers.Response, error) {
	var obj dashboardparent
	var arraobj []dashboardparent
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()

	tglnow, _ := goment.New()

	sql_periode := `SELECT 
			A.idcomppasaran , A.idpasarantogel, B.nmpasarantogel 
			FROM ` + config.DB_tbl_mst_company_game_pasaran + ` as A
			JOIN ` + config.DB_tbl_mst_pasaran_togel + ` as B ON B.idpasarantogel  = A.idpasarantogel  
			WHERE A.idcompany = ? 
			AND A.statuspasaranactive = 'Y' 
			ORDER BY B.nmpasarantogel DESC 
		`
	row, err := con.QueryContext(ctx, sql_periode, company)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcomppasaran_db                     int
			idpasarantogel_db, nmpasarantogel_db string
		)

		err = row.Scan(&idcomppasaran_db, &idpasarantogel_db, &nmpasarantogel_db)
		helpers.ErrorCheck(err)
		var objdetail dashboarddetail
		var arraobjdetail []dashboarddetail
		for i := 1; i < 13; i++ {
			month_name := ""
			month_number := ""
			switch i {
			case 1:
				month_name = "JAN"
				month_number = "01"
			case 2:
				month_name = "FEB"
				month_number = "02"
			case 3:
				month_name = "MAR"
				month_number = "03"
			case 4:
				month_name = "APR"
				month_number = "04"
			case 5:
				month_name = "MAY"
				month_number = "05"
			case 6:
				month_name = "JUN"
				month_number = "06"
			case 7:
				month_name = "JUL"
				month_number = "07"
			case 8:
				month_name = "AUG"
				month_number = "08"
			case 9:
				month_name = "SEP"
				month_number = "09"
			case 10:
				month_name = "OCT"
				month_number = "10"
			case 11:
				month_name = "NOV"
				month_number = "11"
			case 12:
				month_name = "DEC"
				month_number = "12"
			}
			start := tglnow.Format("YYYY-") + month_number + "-" + "01"
			end := tglnow.Format("YYYY-") + month_number + "-" + helpers.GetEndRangeDate(month_name)
			var winlose int = _winlose(company, start, end, idcomppasaran_db)
			objdetail.Pasaranwinlose = winlose
			arraobjdetail = append(arraobjdetail, objdetail)
			msg = "Success"
		}
		obj.Pasaranname = nmpasarantogel_db
		obj.Pasarandetail = arraobjdetail
		arraobj = append(arraobj, obj)
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()
	return res, nil
}

func _winlose(company, start, end string, idcomppasaran int) int {
	con := db.CreateCon()
	ctx := context.Background()
	var winlose float64 = 0
	tbl_trx_keluarantogel, _, _ := Get_mappingdatabase(company)

	sql_keluaran := `SELECT
		COALESCE(SUM(total_outstanding-total_cancel-winlose),0 )  as winlose
		FROM ` + tbl_trx_keluarantogel + `  
		WHERE idcompany = ? 
		AND idcomppasaran = ? 
		AND datekeluaran >= ? 
		AND datekeluaran <= ? 
		AND keluarantogel != ''  
	`
	row := con.QueryRowContext(ctx, sql_keluaran, company, idcomppasaran, start, end)
	switch e := row.Scan(&winlose); e {
	case sql.ErrNoRows:

	case nil:

	default:
		panic(e)
	}
	return int(winlose)
}
