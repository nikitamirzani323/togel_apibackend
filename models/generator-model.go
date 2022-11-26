package models

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

type generatorJobs struct {
	Idtrxkeluaran  string
	Datetimedetail string
	Company        string
	Username       string
	Nomortogel     string
	create         string
	createDate     string
}
type generatorResult struct {
	Idtrxkeluarandetail string
	Message             string
	Status              string
}

func Save_Generator(agent, company, idtrxkeluaran string, totalmember, totalrow int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"

	render_page := time.Now()
	flag := false

	tbl_trx_keluarantogel, tbl_trx_keluarantogel_detail, _ := Get_mappingdatabase(company)
	flag = CheckDB(tbl_trx_keluarantogel, "idtrxkeluaran", idtrxkeluaran)

	if flag {

		go _runner_worker(tbl_trx_keluarantogel_detail, agent, company, idtrxkeluaran, totalmember, totalrow)

		msg = "Success"

	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()
	return res, nil
}
func _runner_worker(fieldtable, agent, company, idtrxkeluaran string, totalmember, totaljob int) {
	tglnow, _ := goment.New()
	runtime.GOMAXPROCS(8)
	render_page := time.Now()
	totalWorker := 100
	total_bet := totalmember * totaljob
	buffer_bet := total_bet + 1
	jobs_bet := make(chan generatorJobs, buffer_bet)
	results_bet := make(chan generatorResult, buffer_bet)
	wg := &sync.WaitGroup{}
	for w := 0; w < totalWorker; w++ {
		wg.Add(1)
		go _doJobInsertTransaksi(fieldtable, jobs_bet, results_bet, wg)
	}
	for x := 0; x < totalmember; x++ {
		username_member := "developer_" + strconv.Itoa(x)
		for i := 0; i < totaljob; i++ {
			prize_4D := helpers.GenerateNumber(4)

			jobs_bet <- generatorJobs{
				Idtrxkeluaran:  idtrxkeluaran,
				Datetimedetail: tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				Company:        company,
				Username:       username_member,
				Nomortogel:     prize_4D,
				create:         agent,
				createDate:     tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			}
		}
	}
	close(jobs_bet)

	for a := 0; a < total_bet; a++ { //RESULT
		flag_result := <-results_bet
		if flag_result.Status == "Failed" {
			log.Printf("ID : %s, Message: %s", flag_result.Idtrxkeluarandetail, flag_result.Message)
		}
		// log.Printf("ID : %s, Message: %s", flag_result.Idtrxkeluarandetail, flag_result.Message)

	}
	close(results_bet)
	wg.Wait()
	log.Println("Selesai Sudah")
	log.Println("TIME JOBS: ", time.Since(render_page).String())
}
func _doJobInsertTransaksi(fieldtable string, jobs <-chan generatorJobs, results chan<- generatorResult, wg *sync.WaitGroup) {
	tglnow, _ := goment.New()
	for capture := range jobs {
		for {
			var outerError error
			func(outerError *error) {
				defer func() {
					if err := recover(); err != nil {
						*outerError = fmt.Errorf("%v", err)
					}
				}()

				sql_insert := `
					insert into
					` + fieldtable + ` (
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
				year := tglnow.Format("YY")
				month := tglnow.Format("MM")
				field_column_counter := fieldtable + tglnow.Format("YYYY") + month
				idrecord_counter := Get_counter(field_column_counter)

				idrecord_counter2 := strconv.Itoa(idrecord_counter)
				idrecord := string(year) + string(month) + idrecord_counter2
				flag_insert, msg_insert := Exec_SQL(sql_insert, fieldtable, "INSERT",
					idrecord, capture.Idtrxkeluaran, capture.Datetimedetail,
					"127.0.0.1", capture.Company, capture.Username, "4D", capture.Nomortogel, "FULL", 100, 0, 4000, 0,
					"ASIA/JAKARTA", "WEBSITE", "RUNNING",
					capture.create, capture.createDate)

				if !flag_insert {

					results <- generatorResult{Idtrxkeluarandetail: idrecord, Message: msg_insert, Status: "Failed"}
				} else {
					results <- generatorResult{Idtrxkeluarandetail: idrecord, Message: "Tidak ada masalah", Status: "Success"}
				}

			}(&outerError)
			if outerError == nil {
				break
			}
		}
	}
	wg.Done()
}
