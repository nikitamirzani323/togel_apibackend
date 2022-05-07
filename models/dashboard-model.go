package models

import (
	"context"
	"database/sql"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/config"
	"bitbucket.org/isbtotogroup/apibackend_go/db"
	"bitbucket.org/isbtotogroup/apibackend_go/entities"
	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/gofiber/fiber/v2"
)

type dashboardparent struct {
	Pasaranname   string      `json:"pasaran_name"`
	Pasarandetail interface{} `json:"pasaran_detial"`
}
type dashboarddetail struct {
	Pasaranwinlose int
}

func Fetch_dashboardwinlose(company, year string) (helpers.Response, error) {
	var obj entities.Model_dashboardwinlose_parent
	var arraobj []entities.Model_dashboardwinlose_parent
	var res helpers.Response
	msg := "Data Not Found"
	render_page := time.Now()

	var objdetail entities.Model_dashboardwinlose_child
	var arraobjdetail []entities.Model_dashboardwinlose_child

	for i := 1; i < 13; i++ {
		periode := ""
		switch i {
		case 1:
			periode = year + "-01"
		case 2:
			periode = year + "-02"
		case 3:
			periode = year + "-03"
		case 4:
			periode = year + "-04"
		case 5:
			periode = year + "-05"
		case 6:
			periode = year + "-06"
		case 7:
			periode = year + "-07"
		case 8:
			periode = year + "-08"
		case 9:
			periode = year + "-09"
		case 10:
			periode = year + "-10"
		case 11:
			periode = year + "-11"
		case 12:
			periode = year + "-12"
		}

		var winlose int = _invoicewinlose_id(company, year, periode)
		objdetail.Dashboardwinlose_winlose = winlose
		arraobjdetail = append(arraobjdetail, objdetail)
		msg = "Success"
	}

	obj.Dashboardwinlose_nmagen = _company(company)
	obj.Dashboardwinlose_detail = arraobjdetail
	arraobj = append(arraobj, obj)

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()
	return res, nil
}

func Fetch_dashboard(company, year string) (helpers.Response, error) {
	var obj entities.Model_dashboardagenpasaranwinlose_parent
	var arraobj []entities.Model_dashboardagenpasaranwinlose_parent
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()

	sql_periode := `SELECT 
			A.idcomppasaran , A.idpasarantogel, B.nmpasarantogel 
			FROM ` + config.DB_tbl_mst_company_game_pasaran + ` as A
			JOIN ` + config.DB_tbl_mst_pasaran_togel + ` as B ON B.idpasarantogel  = A.idpasarantogel  
			WHERE A.idcompany = ? 
			AND A.statuspasaranactive = 'Y' 
			ORDER BY B.nmpasarantogel ASC  
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
		var objdetail entities.Model_dashboardagenpasaranwinlose_child
		var arraobjdetail []entities.Model_dashboardagenpasaranwinlose_child
		for i := 1; i < 13; i++ {
			periode := ""
			idinvoice := 0
			switch i {
			case 1:
				periode = year + "-01"
			case 2:
				periode = year + "-02"
			case 3:
				periode = year + "-03"
			case 4:
				periode = year + "-04"
			case 5:
				periode = year + "-05"
			case 6:
				periode = year + "-06"
			case 7:
				periode = year + "-07"
			case 8:
				periode = year + "-08"
			case 9:
				periode = year + "-09"
			case 10:
				periode = year + "-10"
			case 11:
				periode = year + "-11"
			case 12:
				periode = year + "-12"
			}
			idinvoice = _invoicewinlose_getidinvoice(company, year, periode)
			var winlose int = _invoicewinlosepasaran_id(idinvoice, idcomppasaran_db)
			objdetail.Dashboardagenpasaran_winlose = winlose
			arraobjdetail = append(arraobjdetail, objdetail)
			msg = "Success"
		}
		obj.Dashboardagenpasaran_nmpasaran = nmpasarantogel_db
		obj.Dashboardagenpasaran_detail = arraobjdetail
		arraobj = append(arraobj, obj)
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()
	return res, nil
}

func _invoicewinlose_getidinvoice(company, year, periode string) int {
	con := db.CreateCon()
	ctx := context.Background()
	result := 0
	sql_select := `SELECT 
		idcompinvoice  
		FROM ` + config.DB_tbl_trx_company_invoice + `  
		WHERE yearinvoice = ? 
		AND idcompany = ? 
		AND periodeinvoice = ? 
	`
	var (
		idcompinvoice_db int
	)
	rows := con.QueryRowContext(ctx, sql_select, year, company, periode)
	switch err := rows.Scan(&idcompinvoice_db); err {
	case sql.ErrNoRows:

	case nil:
		result = idcompinvoice_db
	default:
		helpers.ErrorCheck(err)
	}
	return result
}
func _invoicewinlosepasaran_id(idcompinvoice, idcomppasaran int) int {
	con := db.CreateCon()
	ctx := context.Background()
	result := 0
	sql_select := `SELECT 
		winlosecomppasaran  
		FROM ` + config.DB_tbl_trx_company_invoice_detail + `   
		WHERE idcompinvoice = ? 
		AND idcomppasaran = ? 
	`
	var (
		winlosecomp_db int
	)
	rows := con.QueryRowContext(ctx, sql_select, idcompinvoice, idcomppasaran)
	switch err := rows.Scan(&winlosecomp_db); err {
	case sql.ErrNoRows:

	case nil:
		result = winlosecomp_db
	default:
		helpers.ErrorCheck(err)
	}
	return result
}
func _invoicewinlose_id(company, year, periode string) int {
	con := db.CreateCon()
	ctx := context.Background()
	result := 0
	sql_select := `SELECT 
		winlosecomp  
		FROM ` + config.DB_tbl_trx_company_invoice + `  
		WHERE yearinvoice = ? 
		AND idcompany = ? 
		AND periodeinvoice = ? 
	`
	var (
		winlosecomp_db int
	)
	rows := con.QueryRowContext(ctx, sql_select, year, company, periode)
	switch err := rows.Scan(&winlosecomp_db); err {
	case sql.ErrNoRows:

	case nil:
		result = winlosecomp_db
	default:
		helpers.ErrorCheck(err)
	}
	return result
}
func _company(company string) string {
	con := db.CreateCon()
	ctx := context.Background()
	result := ""
	sql_select := `SELECT
		nmcompany
		FROM ` + config.DB_tbl_mst_company + `  
		WHERE idcompany = ?  
	`
	row := con.QueryRowContext(ctx, sql_select, company)
	switch e := row.Scan(&result); e {
	case sql.ErrNoRows:

	case nil:

	default:
		panic(e)
	}
	return result
}
