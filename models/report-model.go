package models

import (
	"context"
	"math"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/db"
	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

type winlose struct {
	Report_client_username string  `json:"report_client_username"`
	Report_client_turnover int     `json:"report_client_turnover"`
	Report_client_winlose  int     `json:"report_client_winlose"`
	Report_agent_winlose   float64 `json:"report_agent_winlose"`
}

func Fetch_winlose(company, start, end string) (helpers.ResponseReportWinlose, error) {
	var obj winlose
	var arraobj []winlose
	var res helpers.ResponseReportWinlose
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	render_page := time.Now()
	tglnow_start, _ := goment.New(start)
	tglnow_end, _ := goment.New(end)

	tbl_trx_keluarantogel, _, tbl_trx_keluarantogel_member := Get_mappingdatabase(company)
	subtotal_turnover := 0
	subtotal_winlose := 0
	var subtotal_winlose_agent float64 = 0
	sql_winlose := `SELECT 
			A.username , SUM(A.totalbayar) as turnover, SUM(A.totalwin-A.totalbayar) as winlose
			FROM ` + tbl_trx_keluarantogel_member + ` as A 
			JOIN ` + tbl_trx_keluarantogel + ` as B ON B.idtrxkeluaran  = A.idtrxkeluaran  
			WHERE B.idcompany = ? 
			AND B.keluarantogel != '' 
			AND B.datekeluaran >= ? 
			AND B.datekeluaran <= ? 
			GROUP BY A.username 
		`
	row, err := con.QueryContext(ctx, sql_winlose, company, tglnow_start.Format("YYYY-MM-DD"), tglnow_end.Format("YYYY-MM-DD"))
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			turnover_db, winlose_db int
			username_db             string
		)

		err = row.Scan(&username_db, &turnover_db, &winlose_db)
		helpers.ErrorCheck(err)
		subtotal_turnover = subtotal_turnover + turnover_db
		subtotal_winlose = subtotal_winlose + winlose_db
		subtotal_winlose_agent = subtotal_winlose_agent + math.Abs(float64(winlose_db))
		obj.Report_client_username = username_db
		obj.Report_client_turnover = turnover_db
		obj.Report_client_winlose = winlose_db
		obj.Report_agent_winlose = math.Abs(float64(winlose_db))
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Subtotalturnover = subtotal_turnover
	res.Subtotalwinlose = subtotal_winlose
	res.Subtotalwinlose_company = int(subtotal_winlose_agent)
	res.Time = time.Since(render_page).String()
	return res, nil
}
