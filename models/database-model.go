package models

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"strings"
	s "strings"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/config"
	"bitbucket.org/isbtotogroup/apibackend_go/db"
	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/nleeper/goment"
)

func Get_counter(field_column string) int {
	con := db.CreateCon()
	ctx := context.Background()
	idrecord_counter := 0
	field_redis := field_column + "_counter"
	resultredis, flagredis := helpers.GetRedis(field_redis)
	log.Println(flagredis)
	log.Println(resultredis)
	sqlcounter := `SELECT 
					counter 
					FROM ` + config.DB_tbl_counter + ` 
					WHERE nmcounter = ? 
				`
	var counter int = 0
	row := con.QueryRowContext(ctx, sqlcounter, field_column)
	switch e := row.Scan(&counter); e {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		// log.Println(counter)
	default:
		panic(e)
	}
	if counter > 0 {
		idrecord_counter = int(counter) + 1
		sql_update := `UPDATE ` + config.DB_tbl_counter + ` SET counter=? WHERE nmcounter=? `
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_counter, "UPDATE", idrecord_counter, field_column)
		if !flag_update {
			log.Println(msg_update)
		} else {
			helpers.SetRedis(field_redis, idrecord_counter, time.Hour*5)
		}
	} else {
		idrecord_counter = 1
		sql_insert := `insert into ` + config.DB_tbl_counter + ` (nmcounter, counter) values (?, ?) `
		flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_counter, "INSERT", field_column, idrecord_counter)
		if !flag_insert {
			log.Println(msg_insert)
		} else {
			helpers.SetRedis(field_redis, idrecord_counter, time.Hour*5)
		}

	}
	return idrecord_counter
}
func Get_counterbooking(field_column string, booking int) (int, int) {
	con := db.CreateCon()
	ctx := context.Background()
	idrecord_counter := 0
	counter_before := 0

	sqlcounter := `SELECT 
					counter 
					FROM ` + config.DB_tbl_counter + ` 
					WHERE nmcounter = ? 
				`
	var counter int = 0
	row := con.QueryRowContext(ctx, sqlcounter, field_column)
	switch e := row.Scan(&counter); e {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		// log.Println(counter)
	default:
		panic(e)
	}
	if counter > 0 {
		counter_before = int(counter)
		idrecord_counter = int(counter) + booking
		sql_update := `UPDATE ` + config.DB_tbl_counter + ` SET counter=? WHERE nmcounter=? `
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_counter, "UPDATE", idrecord_counter, field_column)
		if !flag_update {
			log.Println(msg_update)
		}
	} else {
		idrecord_counter = booking
		sql_insert := `insert into ` + config.DB_tbl_counter + ` (nmcounter, counter) values (?, ?) `
		flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_counter, "INSERT", field_column, idrecord_counter)
		if !flag_insert {
			log.Println(msg_insert)
		}

	}
	return counter_before + 1, idrecord_counter
}
func Get_listitemsearch(data, pemisah, search string) bool {
	flag := false
	temp := s.Split(data, pemisah)
	for i := 0; i < len(temp); i++ {
		if temp[i] == search {
			flag = true
			break
		}
	}
	return flag
}
func CheckDB(table, field, value string) bool {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	sql_db := `SELECT 
					` + field + ` 
					FROM ` + table + ` 
					WHERE ` + field + ` = ? 
				`
	row := con.QueryRowContext(ctx, sql_db, value)
	switch e := row.Scan(&field); e {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
		flag = false
	case nil:
		flag = true
	default:
		panic(e)
	}
	return flag
}
func CheckDBTwoField(table, field_1, value_1, field_2, value_2 string) bool {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	sql_db := `SELECT 
					` + field_1 + ` 
					FROM ` + table + ` 
					WHERE ` + field_1 + ` = ? 
					AND ` + field_2 + ` = ? 
				`
	log.Println(sql_db)
	row := con.QueryRowContext(ctx, sql_db, value_1, value_2)
	switch e := row.Scan(&field_1); e {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
		flag = false
	case nil:
		flag = true
	default:
		flag = false
	}
	return flag
}
func Get_document() int {
	con := db.CreateCon()
	ctx := context.Background()
	var iddoc int = 0
	tglnow, _ := goment.New()
	starttgl := tglnow.Format("YYYY") + "-01-01 00:00:00"
	endtgl := tglnow.Format("YYYY") + "-12-31 23:59:59"
	flag := false
	sql_doc := `SELECT 
		iddoctrans   
		FROM ` + config.DB_tbl_mst_doctransaksi + `   
		WHERE start_doctrans <= ? 
		AND end_doctrans >= ? 
	`
	row := con.QueryRowContext(ctx, sql_doc, starttgl, endtgl)
	switch e := row.Scan(&iddoc); e {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
		flag = true
	case nil:
		// log.Println(iddoc)
	default:
		panic(e)
	}

	if flag {
		stmt, e := con.PrepareContext(ctx, `
			insert into `+config.DB_tbl_mst_doctransaksi+` (iddoctrans, nmperiode, start_doctrans, end_doctrans, createdoctrans, createdatetimedoctrans) 
			values (?, ?, ?, ?, ?, ?)
		`)
		helpers.ErrorCheck(e)
		res, e := stmt.ExecContext(ctx,
			tglnow.Format("YYYY"),
			"PERIODE "+tglnow.Format("YYYY"),
			starttgl,
			endtgl,
			"SERVER",
			tglnow.Format("YYYY-MM-DD HH:mm:ss"))
		helpers.ErrorCheck(e)
		id, e := res.RowsAffected()
		helpers.ErrorCheck(e)
		if id > 0 {
			log.Println("Insert id", id)
			log.Println("NEW")
			year, _ := strconv.Atoi(tglnow.Format("YYYY"))
			iddoc = year
		} else {
			panic("ERROR CREATE DOCUMENT TRANSAKSI")
		}

	}

	return iddoc
}
func Get_mappingdatabase(company string) (string, string, string) {
	tbl_trx_keluarantogel := "db_tot_" + strings.ToLower(company) + ".tbl_trx_keluarantogel"
	tbl_trx_keluarantogel_detail := "db_tot_" + strings.ToLower(company) + ".tbl_trx_keluarantogel_detail"
	tbl_trx_keluarantogel_member := "db_tot_" + strings.ToLower(company) + ".tbl_trx_keluarantogel_member"

	return tbl_trx_keluarantogel, tbl_trx_keluarantogel_detail, tbl_trx_keluarantogel_member
}
func Get_NextPasaran(company, tglskrg string, idcomppasaran int) string {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false

	nexttgl := ""
	idcomppasaranoff := 0
	for i := 1; i < 7; i++ {
		tglnow, _ := goment.New(tglskrg)
		temp := tglnow.Add(i, "days").Format("YYYY-MM-DD")
		day := tglnow.Format("ddd")
		hari := ""
		switch day {
		case "Sun":
			hari = "minggu"
		case "Mon":
			hari = "senin"
		case "Tue":
			hari = "selasa"
		case "Wed":
			hari = "rabu"
		case "Thu":
			hari = "kamis"
		case "Fri":
			hari = "jumat"
		case "Sat":
			hari = "sabtu"
		}

		sql_doc := `SELECT
			idcomppasaranoff
			FROM ` + config.DB_tbl_mst_company_game_pasaran_offline + ` 
			WHERE idcompany = ?
			AND idcomppasaran = ?
			AND haripasaran = ?
		`

		row := con.QueryRowContext(ctx, sql_doc, company, idcomppasaran, hari)
		switch e := row.Scan(&idcomppasaranoff); e {
		case sql.ErrNoRows:
			flag = false
		case nil:
			flag = true

		default:
			panic(e)
		}
		if flag {
			nexttgl = temp
			break
		}
	}

	return nexttgl
}
func Get_Company(company string) string {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	result := ""
	codecomp := ""

	sql_company := `SELECT
		codecomp 
		FROM ` + config.DB_tbl_mst_company + `  
		WHERE idcompany = ? 
	`
	row := con.QueryRowContext(ctx, sql_company, company)
	switch e := row.Scan(&codecomp); e {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true

	default:
		panic(e)
	}
	if flag {
		result = codecomp
	}
	return result
}
func Get_CompanyPasaran(company, tipe string, idcomppasaran int) string {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	result := ""
	temp := ""
	idpasarantogel := ""

	sql_select := `SELECT
		idpasarantogel 
		FROM ` + config.DB_tbl_mst_company_game_pasaran + `  
		WHERE idcompany = ? 
		AND idcomppasaran  = ? 
	`
	row := con.QueryRowContext(ctx, sql_select, company, idcomppasaran)
	switch e := row.Scan(&idpasarantogel); e {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true

	default:
		panic(e)
	}
	if flag {
		switch tipe {
		case "idpasarantogel":
			temp = idpasarantogel
		}
		result = temp
	}
	return result
}
func Get_AdminRule(company, tipe string, idruleadmin int) string {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	result := ""
	nmruleadmin := ""
	ruleadmin := ""

	sql_company := `SELECT
		nmruleadmin, ruleadmin  
		FROM ` + config.DB_tbl_mst_company_admin_rule + `  
		WHERE idcompany = ? 
		AND idruleadmin = ? 
	`
	row := con.QueryRowContext(ctx, sql_company, company, idruleadmin)
	switch e := row.Scan(&nmruleadmin, &ruleadmin); e {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true

	default:
		panic(e)
	}
	if flag {
		switch tipe {
		case "nmruleadmin":
			result = nmruleadmin
		case "ruleadmin":
			result = ruleadmin
		}
	}
	return result
}
func Get_Admin(company, username string) (string, int) {
	con := db.CreateCon()
	ctx := context.Background()
	typeadmin := ""
	idruleadmin := 0
	sql_admin := `SELECT
		typeadmin, idruleadmin  
		FROM ` + config.DB_tbl_mst_company_admin + `  
		WHERE idcompany = ? 
		AND username_comp  = ? 
	`
	row := con.QueryRowContext(ctx, sql_admin, company, username)
	switch e := row.Scan(&typeadmin, &idruleadmin); e {
	case sql.ErrNoRows:

	case nil:

	default:
		panic(e)
	}

	return typeadmin, idruleadmin
}
func Get_OnlinePasaran(company string, idcomppasaran int, hari, tipe string) bool {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false

	switch tipe {
	case "total_pasaran":
		idcomppasaranoff := ""
		sql_select := `SELECT
			idcomppasaranoff 
			FROM ` + config.DB_tbl_mst_company_game_pasaran_offline + `  
			WHERE idcompany = ? 
			AND idcomppasaran = ?
		`
		row := con.QueryRowContext(ctx, sql_select, company, idcomppasaran)
		switch e := row.Scan(&idcomppasaranoff); e {
		case sql.ErrNoRows:
			flag = false
		case nil:
			flag = true
		default:
			flag = false
		}
	case "hari":
		idcomppasaranoff := ""
		sql_select := `SELECT
			idcomppasaranoff 
			FROM ` + config.DB_tbl_mst_company_game_pasaran_offline + `  
			WHERE idcompany = ? 
			AND idcomppasaran = ? 
			AND haripasaran = ? 
		`
		row := con.QueryRowContext(ctx, sql_select, company, idcomppasaran, hari)
		switch e := row.Scan(&idcomppasaranoff); e {
		case sql.ErrNoRows:
			flag = false
		case nil:
			flag = true
		default:
			flag = false
		}
	default:
		flag = false
	}

	return flag
}
func Get_Trxkeluaran(company, tipe string, idtrxkeluaran int) string {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	result := ""
	temp := ""
	idcomppasaran := 0
	tbl_trx_keluarantogel, _, _ := Get_mappingdatabase(company)
	sql_select := `SELECT
		idcomppasaran 
		FROM ` + tbl_trx_keluarantogel + `  
		WHERE idcompany = ? 
		AND idtrxkeluaran  = ? 
	`
	row := con.QueryRowContext(ctx, sql_select, company, idtrxkeluaran)
	switch e := row.Scan(&idcomppasaran); e {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true

	default:
		panic(e)
	}
	if flag {
		switch tipe {
		case "idcomppasaran":
			temp = strconv.Itoa(idcomppasaran)
		}
		result = temp
	}
	return result
}
func Insert_log(idcompany, username, page, tipe, notebefore, noteafter string) {
	tglnow, _ := goment.New()
	sql_insert := `
		INSERT INTO 
		` + config.DB_tbl_trx_log + ` (
			idlog, datetimelog, yearlog, 
			idcompany, username, pagelog, tipelog,
			notebefore, noteafter 
		) VALUES (
			?, ?, ?,
			?, ?, ?, ?, 
			?, ?
		)
	`

	year := tglnow.Format("YYYY")
	month := tglnow.Format("MM")
	field_col := config.DB_tbl_trx_log + year + month
	idlog_counter := Get_counter(field_col)
	idlog := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idlog_counter)
	flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_trx_log, "INSERT",
		idlog, tglnow.Format("YYYY-MM-DD HH:mm:ss"), year,
		idcompany, username, page, tipe, notebefore, noteafter)
	if flag_insert {
		log.Println(msg_insert)
	} else {
		log.Println(msg_insert)
	}

}
func Exec_SQL(sql, table, action string, args ...interface{}) (bool, string) {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	msg := ""
	stmt_exec, e_exec := con.PrepareContext(ctx, sql)
	helpers.ErrorCheck(e_exec)
	defer stmt_exec.Close()
	rec_exec, e_exec := stmt_exec.ExecContext(ctx, args...)

	helpers.ErrorCheck(e_exec)
	exec, e := rec_exec.RowsAffected()
	helpers.ErrorCheck(e)
	if exec > 0 {
		flag = true
		msg = "Data " + table + " Berhasil di " + action
	} else {
		msg = "Data " + table + " Failed di " + action
	}
	return flag, msg
}
