package models

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
	amqp "github.com/rabbitmq/amqp091-go"
)

type senderJobs struct {
	Totaldata int
	Record    interface{}
}
type generatorJobs struct {
	Idtrxkeluarandetail string
	Idtrxkeluaran       string
	Datetimedetail      string
	Company             string
	Username            string
	Nomortogel          string
	create              string
	createDate          string
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
	year := tglnow.Format("YY")
	month := tglnow.Format("MM")
	field_column_counter := fieldtable + tglnow.Format("YYYY") + month
	counter_before, counter_after := Get_counterbooking(field_column_counter, totalmember*totaljob)

	conn, err := amqp.Dial("amqp://guest:guest@157.230.255.100:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"generator", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for x := 0; x < totalmember; x++ {
		var obj_sender senderJobs
		var arraobj_sender []senderJobs
		var obj generatorJobs
		var arraobj []generatorJobs
		temp_total_sender := 0
		username_member := "developer_" + strconv.Itoa(x)
		for i := 0; i < totaljob; i++ {
			temp_total_sender = temp_total_sender + 1
			idrecord_counter2 := strconv.Itoa(counter_before)
			idrecord := string(year) + string(month) + idrecord_counter2

			prize_4D := helpers.GenerateNumber(4)

			jobs_bet <- generatorJobs{
				Idtrxkeluarandetail: idrecord,
				Idtrxkeluaran:       idtrxkeluaran,
				Datetimedetail:      tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				Company:             company,
				Username:            username_member,
				Nomortogel:          prize_4D,
				create:              agent,
				createDate:          tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			}
			obj.Idtrxkeluarandetail = idrecord
			obj.Idtrxkeluaran = idtrxkeluaran
			obj.Datetimedetail = tglnow.Format("YYYY-MM-DD HH:mm:ss")
			obj.Company = company
			obj.Username = username_member
			obj.Nomortogel = prize_4D
			obj.create = agent
			obj.createDate = tglnow.Format("YYYY-MM-DD HH:mm:ss")
			arraobj = append(arraobj, obj)
			if counter_before > counter_after {
				log.Println("Counter Troubke")
			}
			// log.Println(idrecord)
			counter_before = counter_before + 1
		}
		obj_sender.Totaldata = temp_total_sender
		obj_sender.Record = arraobj
		arraobj_sender = append(arraobj_sender, obj_sender)
		body, _ := json.Marshal(arraobj_sender)

		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", body)
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
	noinvoice, _ := strconv.Atoi(idtrxkeluaran)
	_deleteredis_generator(company, noinvoice)
}
func _doJobInsertTransaksi(fieldtable string, jobs <-chan generatorJobs, results chan<- generatorResult, wg *sync.WaitGroup) {

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
				flag_insert, msg_insert := Exec_SQL(sql_insert, fieldtable, "INSERT",
					capture.Idtrxkeluarandetail, capture.Idtrxkeluaran, capture.Datetimedetail,
					"127.0.0.1", capture.Company, capture.Username, "4D", capture.Nomortogel, "FULL", 100, 0, 4000, 0,
					"ASIA/JAKARTA", "WEBSITE", "RUNNING",
					capture.create, capture.createDate)

				if !flag_insert {

					results <- generatorResult{Idtrxkeluarandetail: capture.Idtrxkeluarandetail, Message: msg_insert, Status: "Failed"}
				} else {
					results <- generatorResult{Idtrxkeluarandetail: capture.Idtrxkeluarandetail, Message: "Tidak ada masalah", Status: "Success"}
				}

			}(&outerError)
			if outerError == nil {
				break
			}
		}
	}
	wg.Done()
}
func _deleteredis_generator(company string, idtrxkeluaran int) {
	const Fieldperiode_home_redis = "LISTPERIODE_AGENT_"
	log.Println("REDIS DELETE")
	log.Println("COMPANY :", company)
	log.Println("INVOICE :", idtrxkeluaran)

	//AGEN
	field_home_redis := Fieldperiode_home_redis + strings.ToLower(company)
	val_homeredis := helpers.DeleteRedis(field_home_redis)
	log.Printf("Redis Delete AGEN - PERIODE HOME : %d", val_homeredis)

	field_homedetail_redis := Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran)
	val_homedetailredis := helpers.DeleteRedis(field_homedetail_redis)
	log.Printf("%s\n", field_homedetail_redis)
	log.Printf("Redis Delete AGEN - PERIODE DETAIL : %d", val_homedetailredis)

	field_homedetail_listmember_redis := Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTMEMBER"
	val_homedetaillistmember_redis := helpers.DeleteRedis(field_homedetail_listmember_redis)
	log.Printf("%s\n", field_homedetail_listmember_redis)
	log.Printf("Redis Delete AGEN - PERIODE DETAIL LISTMEMBER : %d", val_homedetaillistmember_redis)

	field_homedetail_listbettable_redis := Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTBETTABLE"
	val_homedetaillistbettable_redis := helpers.DeleteRedis(field_homedetail_listbettable_redis)
	log.Printf("Redis Delete AGEN - PERIODE DETAIL LISTBETTABLE : %d", val_homedetaillistbettable_redis)

	log_redis := "LISTLOG_AGENT_" + strings.ToLower(company)
	val_agent_redis := helpers.DeleteRedis(log_redis)
	log.Printf("Redis Delete AGEN - LOG status: %d", val_agent_redis)

	val_agent_dashboard := helpers.DeleteRedis("DASHBOARD_CHART_AGENT_" + strings.ToLower(company))
	log.Printf("Redis Delete AGENT DASHBOARD status: %d", val_agent_dashboard)

	val_agent_dashboard_pasaranhome := helpers.DeleteRedis("LISTDASHBOARDPASARAN_AGENT_" + strings.ToLower(company))
	log.Printf("Redis Delete AGENT DASHBOARD PASARAN status: %d", val_agent_dashboard_pasaranhome)

	val_agent4d := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_4D")
	val_agent3d := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_3D")
	val_agent2d := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_2D")
	val_agent2dd := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_2DD")
	val_agent2dt := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_2DT")
	val_agentcolokbebas := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_COLOK_BEBAS")
	val_agentcolokmacau := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_COLOK_MACAU")
	val_agentcoloknaga := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_COLOK_NAGA")
	val_agentcolokjitu := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_COLOK_JITU")
	val_agent5050umum := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_50_50_UMUM")
	val_agent5050special := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_50_50_SPECIAL")
	val_agent5050kombinasi := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_50_50_KOMBINASI")
	val_agentmacaukombinasi := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_MACAU_KOMBINASI")
	val_agentdasar := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_DASAR")
	val_agentshio := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTPERMAINAN_SHIO")
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 4D: %d", val_agent4d)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 3D: %d", val_agent3d)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 2D: %d", val_agent2d)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 2DD: %d", val_agent2dd)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 2DT: %d", val_agent2dt)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET COLOK BEBAS: %d", val_agentcolokbebas)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET COLOK MACAU: %d", val_agentcolokmacau)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET COLOK NAGA: %d", val_agentcoloknaga)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET COLOK JITU: %d", val_agentcolokjitu)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 5050UMUM: %d", val_agent5050umum)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 5050SPECIAL: %d", val_agent5050special)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET 5050KOMBINASI: %d", val_agent5050kombinasi)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET MACAU KOMBINASI: %d", val_agentmacaukombinasi)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET DASAR: %d", val_agentdasar)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET SHIO: %d", val_agentshio)
	val_agentall := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTBET_all")
	val_agentwinner := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTBET_winner")
	val_agentcancel := helpers.DeleteRedis(Fieldperiode_home_redis + strings.ToLower(company) + "_INVOICE_" + strconv.Itoa(idtrxkeluaran) + "_LISTBET_cancel")
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS ALL: %d", val_agentall)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS WINNER: %d", val_agentwinner)
	log.Printf("Redis Delete AGENT PERIODE DETAIL LIST BET STATUS CANCEL: %d", val_agentcancel)

}
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
