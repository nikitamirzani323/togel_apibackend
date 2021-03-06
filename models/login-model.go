package models

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"bitbucket.org/isbtotogroup/apibackend_go/config"
	"bitbucket.org/isbtotogroup/apibackend_go/db"
	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/nleeper/goment"
)

type MLogin struct {
}

func Login_Model(username, password, ipaddress, timezone string) (bool, string, string, int, error) {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	tglnow, _ := goment.New()
	var passwordDB, idcompanyDB, typeadminDB string
	var idruleadminDB int
	sql_select := `
			SELECT
			password_comp, idcompany, typeadmin, idruleadmin   
			FROM ` + config.DB_tbl_mst_company_admin + ` 
			WHERE username_comp = ?
			AND status_comp = "ACTIVE" 
		`

	row := con.QueryRowContext(ctx, sql_select, username)
	switch e := row.Scan(&passwordDB, &idcompanyDB, &typeadminDB, &idruleadminDB); e {
	case sql.ErrNoRows:
		return false, "", "", 0, errors.New("Username and Password Not Found")
	case nil:
		flag = true
	default:
		return false, "", "", 0, errors.New("Username and Password Not Found")
	}

	hashpass := helpers.HashPasswordMD5(password)
	log.Println("Password : " + hashpass)
	log.Println("Hash : " + passwordDB)
	if hashpass != passwordDB {
		return false, "", "", 0, nil
	}
	if typeadminDB != "MASTER" {
		flag = CheckDBTwoField(config.DB_tbl_mst_company_admin_iplist, "username_comp", username, "iplist", ipaddress)
		if !flag {
			noteafter := ""
			noteafter += "DATE : " + tglnow.Format("YYYY-MM-DD HH:mm:ss") + " \n"
			noteafter += "IP : " + ipaddress + " - Ipaddress is not register"
			Insert_log(idcompanyDB, username, "LOGIN", "FAILED", "", noteafter)

			return false, "", "", 0, errors.New("Ipaddress is not register")
		}
	}
	if flag {
		sql_update := `
			UPDATE ` + config.DB_tbl_mst_company_admin + ` 
			SET lastlogin_comp=?, lastipaddres_comp=? , lasttimezone_comp=?, 
			updatecomp_admin=?,  updatedatecomp_admin=?  
			WHERE username_comp = ? 
			AND status_comp = "ACTIVE" 
		`
		rows_update, err_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(err_update)
		lastlogin := tglnow.Format("YYYY-MM-DD HH:mm:ss")
		res_update, err_update := rows_update.ExecContext(ctx,
			lastlogin,
			ipaddress,
			timezone,
			username,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			username)
		helpers.ErrorCheck(err_update)
		update, e := res_update.RowsAffected()
		helpers.ErrorCheck(e)
		if update > 0 {
			flag = true
			log.Println("Data Berhasil di save")

			noteafter := ""
			noteafter += "DATE : " + lastlogin + " \n"
			noteafter += "IP : " + ipaddress
			Insert_log(idcompanyDB, username, "LOGIN", "UPDATE", "", noteafter)
		}
	}

	return true, idcompanyDB, typeadminDB, idruleadminDB, nil
}
