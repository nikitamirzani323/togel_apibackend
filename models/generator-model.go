package models

import (
	"log"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

func Save_Generator(agent, company string, idtrxkeluaran int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false
	tbl_trx_keluarantogel, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)
	for i := 0; i <= 600; i++ {
		flag = CheckDB(tbl_trx_keluarantogel, "idtrxkeluaran", strconv.Itoa(idtrxkeluaran))
		if !flag {
			sql_insert := `
				insert into
				` + tbl_trx_keluarantogel_detail + ` (
					idtrxkeluarandetail, idtrxkeluaran, datetimedetail,
					ipaddress, idcompany, username, typegame, nomortogel,posisitogel, bet,
					diskon, win, kei, browsertogel, devicetogel, statuskeluarandetail, 
					createkeluarandetail, createdatekeluarandetail
				) values (
					?, ?, ?, 
					?, ?, ?, ?, ?, ?,?, 
					?, ?, ?, ?, ?, ?,
					?, ?
				)
			`
			prize_4D := helpers.GenerateNumber(4)

			year := tglnow.Format("YY")
			month := tglnow.Format("MM")
			field_column_counter := tbl_trx_keluarantogel_detail + tglnow.Format("YYYY") + month
			idrecord_counter := Get_counter(field_column_counter)

			idrecord_counter2 := strconv.Itoa(idrecord_counter)
			idrecord := string(year) + string(month) + idrecord_counter2

			flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_trx_keluarantogel_detail, "INSERT",
				idrecord, idtrxkeluaran,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				"127.0.0.1", company, "robot", "4D", prize_4D, "FULL", 100, 0, 4000, 0,
				"ASIA/JAKARTA", "WEBSITE", "RUNNING",
				agent, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if !flag_insert {
				log.Println(msg_insert)

			}
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
