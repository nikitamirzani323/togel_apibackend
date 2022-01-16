package models

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/config"
	"bitbucket.org/isbtotogroup/apibackend_go/db"
	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

type adminHome struct {
	No            int    `json:"admin_no"`
	Username      string `json:"admin_username"`
	Nama          string `json:"admin_nama"`
	Tipeadmin     string `json:"admin_tipe"`
	Rule          string `json:"admin_rule"`
	Joindate      string `json:"admin_joindate"`
	Timezone      string `json:"admin_timezone"`
	Lastlogin     string `json:"admin_lastlogin"`
	LastIpaddress string `json:"admin_lastipaddres"`
	Status        string `json:"admin_status"`
}
type adminDetail struct {
	Nama        string `json:"admin_nama"`
	Tipe        string `json:"admin_type"`
	Idruleadmin int    `json:"admin_idrule"`
	Status      string `json:"admin_status"`
	Create      string `json:"admin_create"`
	Update      string `json:"admin_update"`
}
type adminDetail_rule struct {
	Idrule int    `json:"adminrule_idruleadmin"`
	Name   string `json:"adminrule_name"`
}
type adminDetail_iplist struct {
	Idcompiplist int    `json:"adminiplist_idcompiplist"`
	Iplist       string `json:"adminiplist_iplist"`
}

func Fetch_adminHome(company string) (helpers.ResponseAdminManagement, error) {
	var obj adminHome
	var arraobj []adminHome
	var res helpers.ResponseAdminManagement
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql := `SELECT 
			idruleadmin, username_comp, nama_comp, typeadmin,  
			status_comp, lasttimezone_comp, lastlogin_comp, lastipaddres_comp, createdatecomp_admin  
			FROM ` + config.DB_tbl_mst_company_admin + ` 
			WHERE idcompany = ? 
			ORDER BY lastlogin_comp DESC 
		`

	row, err := con.QueryContext(ctx, sql, company)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			idruleadmin_db                                                                                         int
			username_comp_db, nama_comp_db, typeadmin_db                                                           string
			status_comp_db, lasttimezone_comp_db, lastlogin_comp_db, lastipaddres_comp_db, createdatecomp_admin_db string
		)

		err = row.Scan(
			&idruleadmin_db,
			&username_comp_db,
			&nama_comp_db,
			&typeadmin_db,
			&status_comp_db,
			&lasttimezone_comp_db,
			&lastlogin_comp_db,
			&lastipaddres_comp_db,
			&createdatecomp_admin_db)

		helpers.ErrorCheck(err)

		obj.No = no
		obj.Username = username_comp_db
		obj.Nama = nama_comp_db
		obj.Tipeadmin = typeadmin_db
		obj.Rule = Get_AdminRule(company, "nmruleadmin", idruleadmin_db)
		obj.Joindate = createdatecomp_admin_db
		obj.Timezone = lasttimezone_comp_db
		obj.Lastlogin = lastlogin_comp_db
		obj.LastIpaddress = lastipaddres_comp_db
		obj.Status = status_comp_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	var objRule adminDetail_rule
	var arraobjRule []adminDetail_rule
	sql_listrule := `SELECT 
		idruleadmin , nmruleadmin	
		FROM ` + config.DB_tbl_mst_company_admin_rule + ` 
		WHERE idcompany = ? 
	`
	row_listrule, err_listrule := con.QueryContext(ctx, sql_listrule, company)

	helpers.ErrorCheck(err_listrule)
	for row_listrule.Next() {
		var (
			idruleadmin_db int
			nmruleadmin_db string
		)

		err = row_listrule.Scan(&idruleadmin_db, &nmruleadmin_db)

		helpers.ErrorCheck(err)

		objRule.Idrule = idruleadmin_db
		objRule.Name = nmruleadmin_db
		arraobjRule = append(arraobjRule, objRule)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listruleadmin = arraobjRule
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_adminDetail(company, username string) (helpers.ResponseAdminManagement, error) {
	var obj adminDetail
	var arraobj []adminDetail
	var res helpers.ResponseAdminManagement
	msg := "Success"
	ctx := context.Background()
	con := db.CreateCon()
	render_page := time.Now()

	sql := `SELECT 
		nama_comp ,typeadmin, idruleadmin, status_comp, createcomp_admin, 
		createdatecomp_admin, updatecomp_admin, updatedatecomp_admin 
		FROM ` + config.DB_tbl_mst_company_admin + ` 
		WHERE idcompany = ? 
		AND username_comp  = ? 
	`
	var (
		idruleadmin_db                                                                                                                         int
		nama_comp_db, typeadmin_db, status_comp_db, createcomp_admin_db, createdatecomp_admin_db, updatecomp_admin_db, updatedatecomp_admin_db string
	)
	err := con.QueryRowContext(ctx, sql, company, username).Scan(
		&nama_comp_db, &typeadmin_db, &idruleadmin_db, &status_comp_db, &createcomp_admin_db, &createdatecomp_admin_db, &updatecomp_admin_db, &updatedatecomp_admin_db)
	helpers.ErrorCheck(err)

	var objRule adminDetail_rule
	var arraobjRule []adminDetail_rule
	sql_listrule := `SELECT 
		idruleadmin , nmruleadmin	
		FROM ` + config.DB_tbl_mst_company_admin_rule + ` 
		WHERE idcompany = ? 
	`
	row, err := con.QueryContext(ctx, sql_listrule, company)

	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idruleadmin_db int
			nmruleadmin_db string
		)

		err = row.Scan(&idruleadmin_db, &nmruleadmin_db)

		helpers.ErrorCheck(err)

		objRule.Idrule = idruleadmin_db
		objRule.Name = nmruleadmin_db
		arraobjRule = append(arraobjRule, objRule)
		msg = "Success"
	}

	var objIplist adminDetail_iplist
	var arraobjIplist []adminDetail_iplist
	sql_listiplist := `SELECT 
		idcompiplist  , iplist	
		FROM ` + config.DB_tbl_mst_company_admin_iplist + ` 
		WHERE username_comp = ?  
	`
	row_listiplist, err_listiplist := con.QueryContext(ctx, sql_listiplist, username)

	helpers.ErrorCheck(err_listiplist)
	for row_listiplist.Next() {
		var (
			idcompiplist_db int
			iplist_db       string
		)

		err = row_listiplist.Scan(&idcompiplist_db, &iplist_db)

		helpers.ErrorCheck(err)

		objIplist.Idcompiplist = idcompiplist_db
		objIplist.Iplist = iplist_db
		arraobjIplist = append(arraobjIplist, objIplist)
		msg = "Success"
	}

	obj.Nama = nama_comp_db
	obj.Tipe = typeadmin_db
	obj.Idruleadmin = idruleadmin_db
	obj.Status = status_comp_db
	obj.Create = createcomp_admin_db + " - " + createdatecomp_admin_db
	obj.Update = updatecomp_admin_db + " - " + updatedatecomp_admin_db
	arraobj = append(arraobj, obj)

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listruleadmin = arraobjRule
	res.Listiplist = arraobjIplist
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_Admin(agent, company, sData, username, password, nama, status string, idruleadmin int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	ctx := context.Background()
	con := db.CreateCon()
	flag := false

	if sData == "New" {
		sql_insert := `
			insert into
			` + config.DB_tbl_mst_company_admin + ` (
				username_comp , password_comp, idcompany, nama_comp, status_comp, idruleadmin, 
				createcomp_admin, createdatecomp_admin
			) values (
				?, ?, ?, ?, ?, ?, 
				?, ?
			)
		`
		stmt_newpasaran, e_newpasaran := con.PrepareContext(ctx, sql_insert)
		helpers.ErrorCheck(e_newpasaran)
		defer stmt_newpasaran.Close()
		hashpass := helpers.HashPasswordMD5(password)
		res_newpasaran, e_newpasaran := stmt_newpasaran.ExecContext(
			ctx,
			username,
			hashpass,
			company,
			nama,
			"ACTIVE",
			idruleadmin,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"))
		helpers.ErrorCheck(e_newpasaran)
		insert, e := res_newpasaran.RowsAffected()
		helpers.ErrorCheck(e)
		if insert > 0 {
			flag = true
			log.Println("Data Berhasil di save")
		}
	} else {
		if password == "" {
			sql_update := `
				UPDATE 
				` + config.DB_tbl_mst_company_admin + `  
				SET nama_comp=?, status_comp=?, idruleadmin=?,  
				updatecomp_admin=?, updatedatecomp_admin=? 
				WHERE username_comp=? AND idcompany=? 
			`
			stmt_admin, e := con.PrepareContext(ctx, sql_update)
			helpers.ErrorCheck(e)
			rec_admin, e_admin := stmt_admin.ExecContext(
				ctx,
				nama,
				status,
				idruleadmin,
				agent,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				username,
				company)
			helpers.ErrorCheck(e_admin)

			update_admin, e_admin := rec_admin.RowsAffected()
			helpers.ErrorCheck(e_admin)

			defer stmt_admin.Close()
			if update_admin > 0 {
				flag = true
				log.Printf("Update tbl_mst_company_admin Success : %s\n", username)
			} else {
				log.Println("Update tbl_mst_company_admin failed")
			}
		} else {
			hashpass := helpers.HashPasswordMD5(password)
			sql_update2 := `
				UPDATE 
				` + config.DB_tbl_mst_company_admin + `   
				SET nama_comp=?, password_comp=?, status_comp=?, idruleadmin=?, 
				updatecomp_admin=?, updatedatecomp_admin=? 
				WHERE username_comp=? AND idcompany=? 
			`
			stmt_admin, e := con.PrepareContext(ctx, sql_update2)
			helpers.ErrorCheck(e)
			rec_admin, e_admin := stmt_admin.ExecContext(
				ctx,
				nama,
				hashpass,
				status,
				idruleadmin,
				agent,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				username,
				company)
			helpers.ErrorCheck(e_admin)

			update_admin, e_admin := rec_admin.RowsAffected()
			helpers.ErrorCheck(e_admin)

			defer stmt_admin.Close()
			if update_admin > 0 {
				flag = true
				log.Printf("Update tbl_mst_company_admin Success : %s\n", username)
			} else {
				log.Println("Update tbl_mst_company_admin failed")
			}
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
func Save_AdminIplist(agent, company, sData, username, ipaddress string) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	ctx := context.Background()
	con := db.CreateCon()
	flag := false

	if sData == "New" {
		sql_insert := `
			insert into
			` + config.DB_tbl_mst_company_admin_iplist + ` (
				idcompiplist  , username_comp, iplist, 
				createcompiplist, createdatecompiplist
			) values (
				?, ?, ?, 
				?, ?
			)
		`
		stmt_insert, e_insert := con.PrepareContext(ctx, sql_insert)
		helpers.ErrorCheck(e_insert)
		defer stmt_insert.Close()
		idrecord := Get_counter("tbl_mst_company_admin_list" + tglnow.Format("YYYY"))
		res_newpasaran, e_newpasaran := stmt_insert.ExecContext(
			ctx,
			tglnow.Format("YY")+strconv.Itoa(idrecord),
			username,
			ipaddress,
			agent,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"))
		helpers.ErrorCheck(e_newpasaran)
		insert, e := res_newpasaran.RowsAffected()
		helpers.ErrorCheck(e)
		if insert > 0 {
			flag = true
			log.Println("Data Berhasil di save")
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
func Delete_AdminIplist(username string, idcompiplist int) (helpers.Response, error) {
	var res helpers.Response
	tglnow, _ := goment.New()
	ctx := context.Background()
	con := db.CreateCon()

	sql_delete := `
		DELETE FROM
		` + config.DB_tbl_mst_company_admin_iplist + ` 
		WHERE idcompiplist = ? 
		AND username_comp = ?
	`
	stmt, e := con.PrepareContext(ctx, sql_delete)
	helpers.ErrorCheck(e)
	rec, e := stmt.ExecContext(ctx, idcompiplist, username)
	helpers.ErrorCheck(e)
	delete, err_delete := rec.RowsAffected()
	helpers.ErrorCheck(err_delete)
	fmt.Printf("The last delete row id: %d\n", delete)
	defer stmt.Close()
	if delete > 0 {
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
