package models

import (
	"context"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/config"
	"bitbucket.org/isbtotogroup/apibackend_go/db"
	"bitbucket.org/isbtotogroup/apibackend_go/entities"
	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/gofiber/fiber/v2"
)

func Fetch_loghome(company string) (helpers.Response, error) {
	var obj entities.Model_log
	var arraobj []entities.Model_log
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql := `SELECT 
			idlog, datetimelog, username, pagelog,  
			tipelog, noteafter   
			FROM ` + config.DB_tbl_trx_log + ` 
			WHERE idcompany = ? 
			ORDER BY datetimelog DESC LIMIT 300  
		`

	row, err := con.QueryContext(ctx, sql, company)

	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idlog_db                                                          int
			datetimelog_db, username_db, pagelog_db, tipelog_db, noteafter_db string
		)

		err = row.Scan(
			&idlog_db,
			&datetimelog_db,
			&username_db,
			&pagelog_db,
			&tipelog_db,
			&noteafter_db)

		helpers.ErrorCheck(err)

		obj.Log_id = idlog_db
		obj.Log_datetime = datetimelog_db
		obj.Log_username = username_db
		obj.Log_page = pagelog_db
		obj.Log_tipe = tipelog_db
		obj.Log_note = noteafter_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
