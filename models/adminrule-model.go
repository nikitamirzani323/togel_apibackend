package models

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/config"
	"bitbucket.org/isbtotogroup/apibackend_go/db"
	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

type adminHomeRule struct {
	No   int    `json:"adminrule_no"`
	Id   int    `json:"adminrule_id"`
	Nama string `json:"adminrule_nama"`
}
type adminEditRule struct {
	No     int    `json:"adminrule_no"`
	Id     int    `json:"adminrule_id"`
	Nama   string `json:"adminrule_nama"`
	Rule   string `json:"adminrule_rule"`
	Create string `json:"adminrule_create"`
	Update string `json:"adminrule_update"`
}

func Fetch_adminruleHome(company string) (helpers.Response, error) {
	var obj adminHomeRule
	var arraobj []adminHomeRule
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql := `SELECT 
			idruleadmin,nmruleadmin  
			FROM ` + config.DB_tbl_mst_company_admin_rule + ` 
			WHERE idcompany = ? 
			ORDER BY nmruleadmin ASC  
		`
	row, err := con.QueryContext(ctx, sql, company)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			idruleadmin_db int
			nmruleadmin_db string
		)

		err = row.Scan(&idruleadmin_db, &nmruleadmin_db)

		helpers.ErrorCheck(err)

		obj.No = no
		obj.Id = idruleadmin_db
		obj.Nama = nmruleadmin_db
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
func Fetch_adminruleDetail(company string, idruleadmin int) (helpers.Response, error) {
	var obj adminEditRule
	var arraobj []adminEditRule
	var res helpers.Response
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql := `SELECT 
			idruleadmin,nmruleadmin, ruleadmin, 
			createruleadmin, createdateruleadmin,    
			updateruleadmin, updatedateruleadmin    
			FROM ` + config.DB_tbl_mst_company_admin_rule + ` 
			WHERE idcompany = ? 
			AND idruleadmin = ? 
		`
	row, err := con.QueryContext(ctx, sql, company, idruleadmin)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			idruleadmin_db                                                           int
			nmruleadmin_db, ruleadmin_db, createruleadmin_db, createdateruleadmin_db string
			updateruleadmin_db, updatedateruleadmin_db                               string
		)

		err = row.Scan(&idruleadmin_db, &nmruleadmin_db, &ruleadmin_db,
			&createruleadmin_db, &createdateruleadmin_db,
			&updateruleadmin_db, &updatedateruleadmin_db)

		helpers.ErrorCheck(err)
		create := createruleadmin_db + ", " + createdateruleadmin_db
		update := updateruleadmin_db + ", " + updatedateruleadmin_db
		if updateruleadmin_db == "" {
			update = ""
		}
		obj.No = no
		obj.Id = idruleadmin_db
		obj.Nama = nmruleadmin_db
		obj.Rule = ruleadmin_db
		obj.Create = create
		obj.Update = update
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
func Save_Adminrule(agent, company, sData, nama string, idruleadmin int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	ctx := context.Background()
	con := db.CreateCon()
	flag := false

	if sData == "New" {
		sql_insert := `
			insert into
			` + config.DB_tbl_mst_company_admin_rule + ` (
				idruleadmin , idcompany, nmruleadmin, 
				createruleadmin, createdateruleadmin
			) values (
				?, ?, ?, 
				?, ?
			)
		`
		stmt_newpasaran, e_newpasaran := con.PrepareContext(ctx, sql_insert)
		helpers.ErrorCheck(e_newpasaran)
		defer stmt_newpasaran.Close()
		field_column := config.DB_tbl_mst_company_admin_rule + tglnow.Format("YYYY")
		res_newpasaran, e_newpasaran := stmt_newpasaran.ExecContext(
			ctx,
			tglnow.Format("YY")+strconv.Itoa(Get_counter(field_column)),
			company,
			nama,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"))
		helpers.ErrorCheck(e_newpasaran)
		insert, e := res_newpasaran.RowsAffected()
		helpers.ErrorCheck(e)
		if insert > 0 {
			flag = true
			log.Println("Data Berhasil di save")

			noteafter := ""
			noteafter += "NAME RULE - " + nama
			Insert_log(company, agent, "ADMIN RULE", "NEW RULE", "", noteafter)

		}
	} else {
		sql_update := `
				UPDATE 
				` + config.DB_tbl_mst_company_admin_rule + `   
				SET nmruleadmin=?, 
				updateruleadmin=?, updatedateruleadmin=? 
				WHERE idruleadmin=? AND idcompany=? 
			`
		stmt_update, e := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(e)
		rec_update, e_update := stmt_update.ExecContext(
			ctx,
			nama,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idruleadmin,
			company)
		helpers.ErrorCheck(e_update)

		update_admin, e_admin := rec_update.RowsAffected()
		helpers.ErrorCheck(e_admin)

		defer stmt_update.Close()
		if update_admin > 0 {
			flag = true
			log.Printf("Update tbl_mst_company_admin_rule Success : %d\n", idruleadmin)

			noteafter := ""
			noteafter += "NAME RULE - " + nama
			Insert_log(company, agent, "ADMIN RULE", "UPDATE RULE", "", noteafter)
		} else {
			log.Println("Update tbl_mst_company_admin_rule failed")
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = "Success"
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = "Failed"
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func Save_Adminruleconf(agent, company, sData, ruleadmin string, idruleadmin int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	ctx := context.Background()
	con := db.CreateCon()
	flag := false

	sql_update := `
				UPDATE 
				` + config.DB_tbl_mst_company_admin_rule + `   
				SET ruleadmin=?, 
				updateruleadmin=?, updatedateruleadmin=? 
				WHERE idruleadmin=? AND idcompany=? 
			`
	stmt_update, e := con.PrepareContext(ctx, sql_update)
	helpers.ErrorCheck(e)
	rec_update, e_update := stmt_update.ExecContext(
		ctx,
		ruleadmin,
		agent,
		tglnow.Format("YYYY-MM-DD HH:mm:ss"),
		idruleadmin,
		company)
	helpers.ErrorCheck(e_update)

	update_admin, e_admin := rec_update.RowsAffected()
	helpers.ErrorCheck(e_admin)

	defer stmt_update.Close()
	if update_admin > 0 {
		flag = true
		log.Printf("Update tbl_mst_company_admin_rule Success : %d\n", idruleadmin)
	} else {
		log.Println("Update tbl_mst_company_admin_rule failed")
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = "Success"
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = "Failed"
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func _adminrule(idruleadmin int, company, tipecolumn string) string {
	con := db.CreateCon()
	ctx := context.Background()
	var result string = ""
	sql_select := `SELECT 
		nmruleadmin 
		FROM ` + config.DB_tbl_mst_company_admin_rule + `  
		WHERE idruleadmin  = ? 
		AND idcompany = ? 
	`
	var (
		nmruleadmin_db string
	)
	rows := con.QueryRowContext(ctx, sql_select, idruleadmin, company)
	switch err := rows.Scan(&nmruleadmin_db); err {
	case sql.ErrNoRows:
		result = ""
	case nil:
		switch tipecolumn {
		case "nmruleadmin":
			result = nmruleadmin_db
		}
	default:
		helpers.ErrorCheck(err)
	}
	return result
}
