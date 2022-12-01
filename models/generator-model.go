package models

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/apibackend_go/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
	amqp "github.com/rabbitmq/amqp091-go"
)

type senderJobs struct {
	Totaldata     int
	Idtrxkeluaran string
	Company       string
	Record        interface{}
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

	year := tglnow.Format("YY")
	month := tglnow.Format("MM")
	field_column_counter := fieldtable + tglnow.Format("YYYY") + month
	counter_before, counter_after := Get_counterbooking(field_column_counter, totalmember*totaljob)

	AMPQ := os.Getenv("AMQP_SERVER_URL")
	conn, err := amqp.Dial(AMPQ)
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
		obj_sender.Idtrxkeluaran = idtrxkeluaran
		obj_sender.Company = company
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

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
