package models

import (
	"context"
	"database/sql"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/config"
	"bitbucket.org/isbtotogroup/apibackend_go/db"
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
			month_number := ""
			switch i {
			case 1:
				month_number = "01"
			case 2:
				month_number = "02"
			case 3:
				month_number = "03"
			case 4:
				month_number = "04"
			case 5:
				month_number = "05"
			case 6:
				month_number = "06"
			case 7:
				month_number = "07"
			case 8:
				month_number = "08"
			case 9:
				month_number = "09"
			case 10:
				month_number = "10"
			case 11:
				month_number = "11"
			case 12:
				month_number = "12"
			}
			start := tglnow.Format("YYYY-") + month_number + "-" + "01"
			end := tglnow.Format("YYYY-") + month_number + "-" + helpers.GetEndRangeDate("JAN")
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
		COALESCE(SUM(total_outstanding-winlose),0 )  as winlose
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
